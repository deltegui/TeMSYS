class User < ApplicationRecord
  ADMIN = 'admin'
  ANALYST = 'analyst'

  has_secure_password

  validates :username, presence: true
  validates :password, presence: true
  validates :role, inclusion: { in: [User::ADMIN, User::ANALYST] }

  def is_admin?
    self.role == User::ADMIN
  end

  def self.available_roles
    [
      User::ADMIN,
      User::ANALYST,
    ]
  end
end
