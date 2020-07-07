Rails.application.routes.draw do
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
  root 'sessions#index'
  post '/login', to: 'sessions#login_user'
  get '/logout', to: 'sessions#logout'

  scope '/dashboard' do
    get '/status', to: 'status#index'
  end

  scope '/dashboard/admin' do
    get '/users', to: 'user_management#index'
    get '/user/create', to: 'user_management#show_create_user'
  end
end
