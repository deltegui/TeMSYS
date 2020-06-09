class Sensor:
    def __init__(self, name='', report_types=[], ip='192.168.1.1', update_interval=1, supported_reports=[], deleted=False):
        self.name = name
        self.report_types = report_types
        self.update_interval = update_interval
        self.supported_reports = supported_reports
        self.deleted = deleted

    @staticmethod
    def from_array(dictionary):
        self.
