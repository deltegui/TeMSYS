namespace :user do
  desc "Create admin user for your system"
  task create: :environment do
    STDOUT.puts "Enter user name:"
    username = STDIN.gets.squish
    STDOUT.puts "Enter password:"
    password = STDIN.gets.squish
    user = User.new username: username, password: password
    user.save
  end

end
