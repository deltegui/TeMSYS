$basedir="/home/diego/TempanalizrAPI"
echo "Deploying into $1 with username $2..."
ssh -t $2@$1 "mv $basedir/temapidata $basedir/.."
ssh -t $2@$1 'sudo docker-compose down --rmi local'
ssh -t $2@$1 'rm -rf '
scp -rp ./dist/* $2@$1:/home/diego/tempanalizr/
ssh -t $2@$1 'cp -a /home/diego/config.json /home/diego/tempanalizr/'