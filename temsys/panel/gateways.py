import http
from .entities import Sensor


class Sapi:
    def __init__(self, url):
        self.__client = http.client.HTTPConnection(url)

    def get_all_status(self):
        return self.__make_request(endpoint='/sensors/all/now')

    def get_sensor_status(self, sensor):
        return self.__make_request(endpoint=f'/sensor/{sensor.name}/now')

    def get_all_report_types(self):
        return self.__make_request(endpoint='/report/types/all')

    def create_report_type(self, name):
        self.__make_request(method='POST', endpoint=f'/report/types/create/{name}')

    def get_all_sensors(self):
        response = self.__make_request(endpoint='/sensors')
        return Sensor.from_array(response)

    def __make_request(self, method='GET', endpoint='/'):
        self.__client.request(method, endpoint)
        res = self.__client.getresponse()
        return res.read().decode()
