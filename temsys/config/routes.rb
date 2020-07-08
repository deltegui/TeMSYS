Rails.application.routes.draw do
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
  root 'sessions#index'
  post '/login', to: 'sessions#login_user'
  get '/logout', to: 'sessions#logout'

  scope '/dashboard' do
    get '/status', to: 'status#index'
  end

  scope '/dashboard/admin/user' do
    get '', to: 'user_management#index'
    get '/create', to: 'user_management#show_create_user'
    post '/create', to: 'user_management#create_user'
    post '/delete', to: 'user_management#delete_user'
    post '/update', to: 'user_management#update_user'
  end

  scope '/dashboard/admin/sensor' do
    get '', to: 'sensor_management#index'
    post '/update', to: 'sensor_management#update_sensor'
    get '/create', to: 'sensor_management#show_create_sensor'
    post '/create', to: 'sensor_management#create_sensor'
  end
end
