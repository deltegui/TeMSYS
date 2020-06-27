require 'test_helper'

class SapiRepositoryTest < ActiveSupport::TestCase
  test "you can get report types" do
    expected = ['temperature', 'humidity']
    sapi = create_mocked_sapi endpoint: '/reporttypes/all', returns: expected
    assert_equal expected, sapi.report_types
  end

  test "you can create a report type" do
    report_name = 'something'
    sapi = create_mocked_sapi endpoint: "/reporttypes/create/#{report_name}", method: :post
    assert_nil sapi.create_report_type(report_name)
  end

  test "you can get active sensors" do
    expected_sensors = [
      {
        name: 'Bombona',
        connection: {
          type: 'http',
          value: '192.168.1.1',
        },
        updateInterval: 2,
        deleted: false,
        supportedReports: ['temperature']
      },
      {
        name: 'Hello',
        connection: {
          type: 'http',
          value: '192.168.1.2',
        },
        updateInterval: 5,
        deleted: false,
        supportedReports: ['humidity']
      },
    ]
    sapi = create_mocked_sapi endpoint: '/sensors', returns: expected_sensors
    assert_equal expected_sensors, sapi.active_sensors
  end

  test "you can get deleted sensors" do
    expected_sensors = [
      {
        name: 'Bombona',
        connection: {
          type: 'http',
          value: '192.168.1.1',
        },
        updateInterval: 2,
        deleted: true,
        supportedReports: ['temperature']
      },
      {
        name: 'Hello',
        connection: {
          type: 'http',
          value: '192.168.1.2',
        },
        updateInterval: 5,
        deleted: true,
        supportedReports: ['humidity']
      },
    ]
    sapi = create_mocked_sapi endpoint: "/sensors?deleted=true", returns: expected_sensors
    assert_equal expected_sensors, sapi.deleted_sensors
  end

  test "you can know all state" do
    expected_reports = [
      {
        type: "temperature",
        sensor: "salon",
        date: "2020-06-20T00:26:35.069761+02:00",
        value: 27
      },
      {
        type: "humidity",
        sensor: "salon",
        date: "2020-06-20T00:26:35.069762+02:00",
        value: 18
      }
    ]
    sapi = create_mocked_sapi endpoint: '/sensors/now', returns: expected_reports
    assert_equal expected_reports, sapi.all_state
  end

  test "you can create a sensor" do
    sensor = {
      name: 'Hello',
      connection: {
        type: 'http',
        value: '192.168.1.2',
      },
      updateInterval: 5,
      deleted: true,
      supportedReports: ['humidity']
    }
    requester = mock()
    requester.
      expects(:post).
      with("/sensor", sensor).
      returns(sensor)
    sapi = SapiRepository.new requester
    assert_equal sensor, sapi.create_sensor(sensor)
  end

  test "you can get a existing sensor" do
    sensor = {
      name: 'Hello',
      connection: {
        type: 'http',
        value: '192.168.1.2',
      },
      updateInterval: 5,
      deleted: true,
      supportedReports: ['humidity']
    }
    sapi = create_mocked_sapi endpoint: "/sensor/#{sensor[:name]}", returns: sensor
    assert_equal sensor, sapi.sensor(sensor[:name])
  end

  test "you can delete a sensor" do
    name = 'my sensor'
    requester = mock()
    sapi = create_mocked_sapi endpoint: "/sensor/#{name}", method: :delete
    assert_nil sapi.delete_sensor(name)
  end

  test "you can update sensor" do
    sensor = {
      name: 'Hello',
      connection: {
        type: 'http',
        value: '192.168.1.2',
      },
      updateInterval: 5,
      deleted: true,
      supportedReports: ['humidity']
    }
    requester = mock()
    requester.
      expects(:post).
      with('/sensor', sensor).
      returns(sensor)
    sapi = SapiRepository.new requester
    assert_equal sensor, sapi.update_sensor(sensor)
  end

  test "you can get one sensor state" do
    name = 'name'
    expected_reports = [
      {
        type: "temperature",
        sensor: "salon",
        date: "2020-06-20T00:26:35.069761+02:00",
        value: 27
      },
      {
        type: "humidity",
        sensor: "salon",
        date: "2020-06-20T00:26:35.069762+02:00",
        value: 18
      }
    ]
    sapi = create_mocked_sapi endpoint: "/sensor/#{name}/now", returns: expected_reports
    assert_equal expected_reports, sapi.sensor_state(name)
  end

  test "you can get all reports" do
    expected_reports = [
      {
        type: "temperature",
        sensor: "salon",
        date: "2020-06-20T00:26:35.069761+02:00",
        value: 27
      },
      {
        type: "humidity",
        sensor: "salon",
        date: "2020-06-20T00:26:35.069762+02:00",
        value: 18
      }
    ]
    sapi = create_mocked_sapi endpoint: '/report', returns: expected_reports
    assert_equal expected_reports, sapi.report
  end

  test "you can get reports between" do
    expected_reports = [
      {
        type: "temperature",
        sensor: "salon",
        date: "2020-06-20T00:26:35.069761+02:00",
        value: 27
      },
      {
        type: "humidity",
        sensor: "salon",
        date: "2020-06-20T00:26:35.069762+02:00",
        value: 18
      }
    ]
    between = {
      from: '2020-06-20T00:26:35.069761+02:00',
      to: '2020-06-20T00:26:35.069761+02:00',
    }
    sapi = create_mocked_sapi endpoint: "/report/dates?from=#{between['from']}&to=#{between['to']}", returns: expected_reports
    assert_equal expected_reports, sapi.report(between)
  end

  def create_mocked_sapi(endpoint:, method: :get, returns: nil)
    requester = mock()
    requester.
      expects(method).
      with(endpoint).
      returns(returns)
    sapi = SapiRepository.new requester
  end
end