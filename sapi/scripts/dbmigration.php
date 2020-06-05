<?php
require 'config.php';

class Migrator {

    private $pdoIN;

    private $pdoOUT;

    private $readerFunc;

    private $writerFunc;

    public function __construct(string $user, string $password, string $dbin, string $dbout) {
        try {
            $this->pdoIN = new PDO("mysql:host=127.0.0.1;dbname={$dbin}", $user, $password);
            $this->pdoOUT = new PDO("mysql:host=127.0.0.1;dbname={$dbout}", $user, $password);
        } catch(PDOException $e) {
            error_log($e);
            die();
        }
    }

    public function usingAsReader(callable $readerFunc): Migrator {
        $this->readerFunc = $readerFunc;
        return $this;
    }

    public function usingAsWriter(callable $writerFunc): Migrator {
        $this->writerFunc = $writerFunc;
        return $this;
    }

    public function migrate() {
        $data = call_user_func($this->readerFunc, $this->pdoIN);
        foreach($data as $entry) {
            echo '.';
            call_user_func($this->writerFunc, $this->pdoOUT, $entry);
        }
    }

    public function resetDatabase(bool $mode) {
        $varcharLong = $mode ? 155 : 255;
        $this->pdoOUT->exec('DROP TABLE IF EXISTS USED_REPORT_TYPES;');
        $this->pdoOUT->exec('DROP TABLE IF EXISTS REPORTS');
        $this->pdoOUT->exec('DROP TABLE IF EXISTS SENSORS');
        $this->pdoOUT->exec('DROP TABLE IF EXISTS REPORT_TYPES');
        $this->pdoOUT->exec("CREATE TABLE REPORT_TYPES (
            TYPE_NAME VARCHAR({$varcharLong}) PRIMARY KEY
        )");
        $this->pdoOUT->exec("CREATE TABLE SENSORS (
            NAME VARCHAR({$varcharLong}) PRIMARY KEY,
            CONNTYPE VARCHAR({$varcharLong}) NOT NULL,
            CONNVALUE VARCHAR({$varcharLong}) NOT NULL,
            UPDATE_INTERVAL INTEGER NOT NULL,
            DELETED BOOLEAN NOT NULL DEFAULT 0
        )");
        $this->pdoOUT->exec("CREATE TABLE USED_REPORT_TYPES (
            SENSOR VARCHAR({$varcharLong}) NOT NULL,
            REPORT_TYPE VARCHAR({$varcharLong}) NOT NULL,
            ADD_DATE DATETIME(3)
        )");
        $this->pdoOUT->exec("CREATE TABLE REPORTS (
            ID INTEGER PRIMARY KEY AUTO_INCREMENT,
            SENSOR VARCHAR({$varcharLong}),
            TYPE VARCHAR({$varcharLong}) NOT NULL,
            VALUE FLOAT NOT NULL,
            REPORT_DATE DATETIME(3)
        )");
        $this->pdoOUT->exec("ALTER TABLE REPORTS ADD CONSTRAINT FOREIGN KEY (SENSOR) REFERENCES SENSORS(NAME)");
        $this->pdoOUT->exec("ALTER TABLE REPORTS ADD CONSTRAINT FOREIGN KEY (TYPE) REFERENCES REPORT_TYPES(TYPE_NAME)");
        $this->pdoOUT->exec("ALTER TABLE USED_REPORT_TYPES ADD CONSTRAINT PRIMARY KEY (SENSOR, REPORT_TYPE)");
        $this->pdoOUT->exec("ALTER TABLE USED_REPORT_TYPES ADD CONSTRAINT FOREIGN KEY (SENSOR) REFERENCES SENSORS(NAME)");
        $this->pdoOUT->exec("ALTER TABLE USED_REPORT_TYPES ADD CONSTRAINT FOREIGN KEY (REPORT_TYPE) REFERENCES REPORT_TYPES(TYPE_NAME)");
    }
}

function sensorReader($pdo) {
    $statement = $pdo->prepare('SELECT
            NAME,
            CONNTYPE,
            CONNVALUE,
            SUPPORTED_REPORT_TYPES,
            DELETED
        FROM
            SENSORS');
    $statement->execute();
    return $statement->fetchAll();
}

function sensorWriter($pdo, $entry) {
    echo "Migrating:";
    print_r($entry);
    $s = $pdo->prepare("INSERT INTO SENSORS
            (NAME,
            CONNTYPE,
            CONNVALUE,
            UPDATE_INTERVAL,
            DELETED)
        VALUES
            (:NAME,
            :CONNTYPE,
            :CONNVALUE,
            60,
            :DELETED)");
    $s->bindParam(':NAME', $entry['NAME']);
    $s->bindParam(':CONNTYPE', $entry['CONNTYPE']);
    $s->bindParam(':CONNVALUE', $entry['CONNVALUE']);
    $s->bindParam(':DELETED', $entry['DELETED']);
    if(!$s->execute()) {
        print "FAIL: \n" . $s->errorInfo()[2] . "\n";
    }
    $reports = explode(';', $entry['SUPPORTED_REPORT_TYPES']);
    foreach($reports as $r) {
        echo "Insert report type $r\n";
        $reportStmt = $pdo->prepare("INSERT INTO REPORT_TYPES VALUES (:R)");
        $reportStmt->bindParam(':R', $r);
        $reportStmt->execute();
        $used = $pdo->prepare("INSERT INTO USED_REPORT_TYPES (SENSOR, REPORT_TYPE, ADD_DATE) VALUES (:NAME, :TYPE, NOW())");
        $used->bindParam(':NAME', $entry['NAME']);
        $used->bindParam(':TYPE', $r);
        if(!$used->execute()) {
            print "FAIL: \n";
            print_r($s->errorInfo());
            die();
        }
    }
}

function reportReader($pdo) {
    $statement = $pdo->prepare("SELECT
            SENSOR,
            TYPE,
            VALUE,
            REPORT_DATE
        FROM
            REPORTS");
    $statement->execute();
    $data = $statement->fetchAll();
    echo "Reports to migrate: " . count($data) . "\n";
    return $data;
}

function reportWriter($pdo, $entry) {
    $s = $pdo->prepare("INSERT INTO REPORTS
        (SENSOR,
        TYPE,
        VALUE,
        REPORT_DATE)
    VALUES
        (:SENSOR,
        :TYPE,
        :VALUE,
        :DATE)");
    $s->bindParam(':SENSOR', $entry['SENSOR']);
    $s->bindParam(':TYPE', $entry['TYPE']);
    $s->bindParam(':VALUE', $entry['VALUE']);
    $s->bindParam(':DATE', $entry['REPORT_DATE']);
    if(!$s->execute()) {
        print "FAIL: \n" . $s->errorInfo()[2] . "\n";
    }
}

$user = $config['db']['user'];
$pass = $config['db']['password'];
$dbin = $argv[1];
$dbout = $argv[2];
$reset = $argv[3] ?? '';
$shortMode = isset($argv[4]) && $argv[4] === 'short';
$migrator = new Migrator($user, $pass, $dbin, $dbout);
echo "Migrating from {$dbin} to {$dbout}:\n";
if($reset === 'reset') {
    print "Reset database...\n";
    $migrator->resetDatabase($shortMode);
}
echo "Migrating Sensors...\n";
$migrator
    ->usingAsReader('sensorReader')
    ->usingAsWriter('sensorWriter')
    ->migrate();
echo "Migrating Readers...\n";
$migrator
    ->usingAsReader('reportReader')
    ->usingAsWriter('reportWriter')
    ->migrate();
echo "DONE\n";