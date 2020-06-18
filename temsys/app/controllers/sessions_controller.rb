class SessionsController < ApplicationController
  def index
  end

  def login_user
    user = User.find_by username: params[:username]
    if !user || !user.authenticate(params[:password])
      render :index
      return
    end
    session[:user_id] = user.id
    redirect_to '/dashboard'
  end

  def logout
    session.delete :user_id
    redirect_to '/'
  end
end
