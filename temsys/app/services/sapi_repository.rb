require 'net/http'
require 'json'

class SapiRepository

  def initialize(requester)
    @requester = requester
  end

  def self.build(url)
    new HttpRequester.new(url)
  end

  def self.from_config
    SapiRepository.build Rails.application.config.sapi.url
  end

  def report_types
    @requester.get '/reporttypes/all'
  end

  def create_report_type(name)
    @requester.post "/reporttypes/create/#{name}"
  end

  def active_sensors
    @requester.get "/sensors"
  end

  def deleted_sensors
    @requester.get "/sensors?deleted=true"
  end

  def all_state
    @requester.get "/sensors/now"
  end

  def create_sensor(sensor)
    @requester.post "/sensor", sensor
  end

  def sensor(name)
    @requester.get "/sensor/#{name}"
  end

  def delete_sensor(name)
    @requester.delete "/sensor/#{name}"
  end

  def update_sensor(sensor)
    @requester.post "/sensor/update", sensor
  end

  def sensor_state(name)
    @requester.get "/sensor/#{name}/now"
  end

  def report(between = nil)
    if between
      @requester.get "/report/dates?from=#{between['from']}&to=#{between['to']}"
    else
      @requester.get "/report"
    end
  end
end
