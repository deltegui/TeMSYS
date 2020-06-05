build:
	mkdir ./build
	go build -o ./build/fynd ./main.go

clean:
	rm -rf ./build

watch:
	reflex -r '(.go|.html)' -s -- sh -c 'go run ./main.go'