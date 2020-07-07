namespace :user do
  desc "Create admin user for your system"
  task create: :environment do
    STDOUT.puts "Enter user name:"
    username = STDIN.gets.squish
    STDOUT.puts "Enter password:"
    password = STDIN.gets.squish
    STDOUT.puts "Is admin?"
    admin_response = STDIN.gets.squish
    role = admin_response == 'y' ? User::ADMIN : User::ANALYST
    user = User.new username: username, password: password, role: role
    user.save
  end

end
