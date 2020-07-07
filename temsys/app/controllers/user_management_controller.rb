class UserManagementController < ApplicationController
  before_action :athorized

  def index
    render :index, layout: 'dashboard'
  end
end
