build: build-back public-folder
	cd ./client && npm run build
	cp -a ./client/dist/* ./dist/public
	cp ./config.json ./dist/config.json

build-back:
	npm run build

public-folder:
	mkdir ./dist/public

clean:
	rm -rf ./dist
	rm -rf ./client/dist

deploy: build
	sh ./deploy.sh 192.168.1.25 diego