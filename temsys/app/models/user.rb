class User < ApplicationRecord
  ADMIN = 'admin'
  ANALYST = 'analyst'

  has_secure_password

  def is_admin?
    self.role == User::ADMIN
  end
end
