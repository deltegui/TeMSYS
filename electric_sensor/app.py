import os
from datetime import datetime, timedelta
from threading import Thread, Lock
from flask import Flask
from oligo import Iber


class WattReport:
    def __init__(self, watts, time):
        self.watts = float(watts)
        self.time = time

    def is_old(self):
        twentyMinutesAgo = datetime.now() - timedelta(minutes=20)
        return self.time < twentyMinutesAgo


class WattCache:
    DATE_FORMAT = "%m/%d/%Y %H:%M:%S"

    def __init__(self, name = './cache'):
        self.mutex = Lock()
        self.file_name = name
        self.create()

    def create(self):
        open(self.file_name, "a+").close()

    def save(self, watt_report):
        self.mutex.acquire()
        with open(self.file_name, "r+") as file:
            file.truncate(0)
        with open(self.file_name, "w") as file:
            date_str = watt_report.time.strftime(WattCache.DATE_FORMAT)
            file.write('{};{}'.format(date_str, watt_report.watts))
        self.mutex.release()

    def load(self):
        self.mutex.acquire()
        with open(self.file_name, "r+") as file:
            raw = file.read()
            if (len(raw) == 0):
                self.mutex.release()
                return None
            [raw_date, raw_watts] = raw.split(';')
            date = datetime.strptime(raw_date, WattCache.DATE_FORMAT)
            self.mutex.release()
            return WattReport(raw_watts, time=date)


class Sensor:
    def __init__(self):
        self.mutex = Lock()
        self.connection = Iber()
        self.connection.login(os.environ.get('IBER_USER'), os.environ.get('IBER_PASS'))
        self.reading = False

    def read(self):
        self.mutex.acquire()
        if self.reading:
            self.mutex.release()
            return
        self.reading = True
        self.mutex.release()

        watt = self.connection.watthourmeter()
        report = WattReport(watt, datetime.now())

        self.mutex.acquire()
        self.reading = False
        self.mutex.release()
        return report


app = Flask(__name__)
sensor = Sensor()
cache = WattCache()

def refresh_data():
    report = sensor.read()
    if report is None:
        return WattReport(0, datetime.now())
    cache.save(report)
    return report

@app.route('/')
def serve_data():
    report = cache.load()
    if report is None:
        report = refresh_data()
    if report.is_old():
        Thread(target=refresh_data).start()
    return {
        "watts": report.watts,
    }
