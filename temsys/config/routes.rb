Rails.application.routes.draw do
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
  root 'sessions#index'
  post '/login', to: 'sessions#login_user'
  get '/logout', to: 'sessions#logout'

  scope '/dashboard' do
    get '/status', to: 'status#index'
  end

  scope '/dashboard/admin' do
    get '/user', to: 'user_management#index'
    get '/user/create', to: 'user_management#show_create_user'
    post '/user/create', to: 'user_management#create_user'
    post '/user/delete', to: 'user_management#delete_user'
    post '/user/update', to: 'user_management#update_user'
  end
end
