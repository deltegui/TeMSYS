class StatusController < ApplicationController

  before_action :athorized

  def initialize
    @sapi = SapiRepository::build "http://localhost:8080"
  end

  def show
    @states = @sapi.all_state
  end
end
