class ApplicationController < ActionController::Base

  def is_logged?
    session[:user_id] != nil
  end

  def athorized
    redirect_to '/' unless is_logged?
  end
end
