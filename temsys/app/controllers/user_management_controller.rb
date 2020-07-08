class UserManagementController < ApplicationController
  before_action :athorized

  def index
    @users = User.all
    render :index, layout: 'dashboard'
  end

  def update_user
    data = self.handle_user_params
    user = User.find_by username: data[:username]
    user.update password: data[:password], role: data[:role] if user
    redirect_to '/dashboard/admin/user'
  end

  def delete_user
    user = User.find_by username: params[:user]
    user.destroy if user
    redirect_to '/dashboard/admin/user'
  end

  def show_create_user
    @user = User.new
    render :show_create_user, layout: 'dashboard'
  end

  def create_user
    unless User.create(self.handle_user_params)
      redirect_to '/dashboard/admin/user/create'
      return
    end
    redirect_to '/dashboard/admin/user'
  end

  private

  def handle_user_params
    {
      username: params[:user][:username],
      password: params[:user][:password],
      role: params[:user][:role],
    }
  end
end
