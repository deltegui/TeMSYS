class SessionsController < ApplicationController

  def index
    redirect_to '/dashboard/status' if is_logged?
  end

  def login_user
    return redirect_to '/dashboard/status' if is_logged?
    username, password = params.values_at(:username, :password)
    unless helpers.login(username, password)
      render :index
      return
    end
    redirect_to '/dashboard/status'
  end

  def logout
    session.delete :user_id
    redirect_to '/'
  end

end
