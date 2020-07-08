class SensorManagementController < ApplicationController

  def initialize
    @sapi = SapiRepository::from_config
    @reports = @sapi.report_types
  end

  def index
    raw_sensors = self.merge_sensors @sapi.active_sensors, @sapi.deleted_sensors
    @sensors = self.enum_to_model raw_sensors
    render :index, layout: 'dashboard'
  end

  def update_sensor
    puts @sapi.update_sensor self.model_to_sensor(params)
    redirect_to '/dashboard/admin/sensor'
  end

  def show_create_sensor
    render :show_create_sensor, layout: 'dashboard'
  end

  def create_sensor
    @sapi.create_sensor self.model_to_sensor(params)
    redirect_to '/dashboard/admin/sensor'
  end

  private

  def model_to_sensor(model)
    supported_reports = @reports.filter { |r| model.keys.include? r }
    {
      name: params[:name],
      connection: {
        type: 'http',
        value: params[:ip],
      },
      updateInterval: params[:update_interval].to_i,
      deleted: params[:delete] == '1' ? true : false,
      supportedReports: supported_reports,
    }
  end

  def enum_to_model(enum)
    enum.map { |sensor| self.sensor_to_model sensor }
  end

  def merge_sensors(first, second)
    [first, second].lazy.flat_map(&:lazy)
  end

  def sensor_to_model(sensor)
    {
      name: sensor['name'],
      ip: sensor['connection']['value'],
      update_interval: sensor['updateInterval'],
      delete: sensor['deleted'],
      supported_reports: self.generate_supported(sensor),
    }
  end

  def generate_supported(sensor)
    @reports.map do |report|
      is_supported = sensor['supportedReports'].include? report
      {
        type: report,
        supported: is_supported,
      }
    end
  end
end
