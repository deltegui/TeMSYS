module SessionsHelper
  def login(username, password)
    user = User.find_by username: username
    unless user && user.authenticate(password)
      return false
    end
    session[:user_id] = user.id
    true
  end

end
