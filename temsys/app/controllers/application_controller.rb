class ApplicationController < ActionController::Base
  helper_method :current_user

  def is_logged?
    session[:user_id] != nil
  end

  def current_user
    User.find_by_id session[:user_id] if is_logged?
  end

  def athorized
    redirect_to '/' unless is_logged?
  end
end
