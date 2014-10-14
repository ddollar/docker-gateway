.PHONY: all build

default: build

build:
	rm -rf build
	gox -osarch="linux/amd64" -output="build/{{.Dir}}"
	docker build -t ddollar/docker-gateway .

release: build
	docker push ddollar/docker-gateway

run: build
	docker run ddollar/docker-gateway
