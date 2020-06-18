Rails.application.routes.draw do
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
  root 'sessions#index'
  post '/login', to: 'sessions#login_user'
  get '/logout', to: 'sessions#logout'
end
