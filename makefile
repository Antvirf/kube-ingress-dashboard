build:
	docker build . -t test --platform=linux/arm64/v8 --file Dockerfile-m1

run:
	docker run --platform=linux/arm64/v8 -p 8080:8080 test:latest