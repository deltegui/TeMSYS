import os
from datetime import datetime, timedelta
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
        self.file_name = name

    def save(self, watt_report):
        with open(self.file_name, "r+") as file:
            file.truncate(0)
        with open(self.file_name, "w") as file:
            date_str = watt_report.time.strftime(WattCache.DATE_FORMAT)
            file.write('{};{}'.format(date_str, watt_report.watts))

    def load(self):
        with open(self.file_name, "r+") as file:
            raw = file.read()
            if (len(raw) == 0):
                return None
            [raw_date, raw_watts] = raw.split(';')
            date = datetime.strptime(raw_date, WattCache.DATE_FORMAT)
            return WattReport(raw_watts, time=date)


class Sensor:
    def __init__(self):
        self.connection = Iber()
        self.connection.login(os.environ.get('IBER_USER'), os.environ.get('IBER_PASS'))

    def read(self):
        watt = self.connection.watthourmeter()
        return WattReport(watt, datetime.now())


app = Flask(__name__)
sensor = Sensor()
cache = WattCache()

@app.route('/')
def serve_data():
    report = cache.load()
    if report is None or report.is_old():
        report = sensor.read()
        cache.save(report)
    return {
        "watts": report.watts,
    }
