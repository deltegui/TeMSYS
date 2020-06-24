require 'net/http'
require 'json'

class SapiRepository

  def initialize(url)
    @base_url = url
  end

  def report_types
    self.get '/reporttypes/all'
  end

  def create_report_type(name)
    self.post "/reporttypes/create/#{name}"
  end

  def active_sensors

  end

  def deleted_sensors

  end

  def active_state

  end

  def create_sensor(sensor)

  end

  def sensor(name)

  end

  def delete_sensor(name)

  end

  def update_sensor(sensor)

  end

  def sensor_state(sensor)

  end

  def report(between)

  end

  private

  def make_uri(endpoint)
    URI.parse "#{@base_url}#{endpoint}"
  end

  def get(endpoint)
    res = Net::HTTP:get(self.make_uri(endpoint))
    JSON[res.body]
  end

  def post(endpoint, data = nil)
    data ||= ""
    res = Net:HTTP::post(self.make_uri(endpoint), data, 'Content-Type' => 'application/json')
    JSON[res.body]
  end

  def delete(endpoint)
    res = Net::HTTP::delete(self.make_uri(endpoint))
    JSON[res.body]
  end
end
