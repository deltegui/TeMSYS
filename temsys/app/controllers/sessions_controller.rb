class SessionsController < ApplicationController
  def index
    if helpers.is_logged?
      redirect_to '/dashboard'
    end
  end

  def login_user
    if helpers.is_logged?
      redirect_to '/dashboard'
    end
    username, password = params.values_at(:username, :password)
    unless helpers.login(username, password)
      render :index
      return
    end
    redirect_to '/dashboard'
  end

  def logout
    session.delete :user_id
    redirect_to '/'
  end

end
