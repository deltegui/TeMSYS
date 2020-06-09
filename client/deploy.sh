basedir="/home/diego/TempClient"
echo "Deploying into $1 with username $2..."
ssh -t $2@$1 "cd $basedir && sudo docker-compose down --rmi local"
ssh -t $2@$1 "rm -rf $basedir"
scp -rp ./dist/* $2@$1:$basedir
ssh -t $2@$1 "cd $basedir && sudo docker-compose up --build -d"
