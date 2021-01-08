# Simple Golang dev environment

docker build -rm --pull -t golang --target dev-env .

docker run -it -v ./src:/src
