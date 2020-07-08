class StatusController < ApplicationController
  before_action :athorized

  def initialize
    @sapi = SapiRepository::from_config
  end

  def index
    states = @sapi.all_state
    sensors = @sapi.active_sensors
    @models = sensors.map do |sensor|
      reports = states.filter { |s| s['sensor'] == sensor['name'] }
      unless reports
        reports = []
      end
      {
        name: sensor['name'],
        ip: sensor['connection']['value'],
        reports: reports,
      }
    end
    render :index, layout: 'dashboard'
  end
end
