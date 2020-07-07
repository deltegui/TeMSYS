require 'test_helper'

class SessionsControllerTest < ActionDispatch::IntegrationTest
  test "index renders index page" do
    get "/"
    assert_response :success
  end

  test "index redirects to dashboard if a user is logged in" do
    sign_in_transaction do
      get "/"
      assert_response :redirect
    end
  end

  test "logout always logs out a user and redirects" do
    sign_in
    get "/logout"
    assert_response :redirect
  end

end
