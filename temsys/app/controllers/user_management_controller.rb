class UserManagementController < ApplicationController
  before_action :athorized

  def index
    @users = User.all
    render :index, layout: 'dashboard'
  end

  def show_create_user
    render :show_create_user, layout: 'dashboard'
  end
end
