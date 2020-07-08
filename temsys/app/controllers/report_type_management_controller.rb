class ReportTypeManagementController < ApplicationController

  def initialize
    @sapi = SapiRepository::from_config
  end

  def index
    @report_types = @sapi.report_types
    render :index, layout: 'dashboard'
  end

  def create_report_type
    @sapi.create_report_type params[:new_report_type]
    redirect_to '/dashboard/admin/report_types'
  end
end
