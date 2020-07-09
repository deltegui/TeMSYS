class AdminController < ApplicationController
  before_action :athorized
  before_action :require_admin

  private

  def require_admin
    unless current_user.present? && current_user.is_admin?
      redirect_to '/dashboard/status'
    end
  end
end