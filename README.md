# HTTP echo

* [Docker Hub page](https://hub.docker.com/r/mhalbritter/http-echo)

## Running

```shell
go run ./...
```

You can use the environment variable `HTTP_ECH0_PORT` to set the HTTP port, and the environment variable `HTTP_ECH0_TLS_PORT` to set the HTTPS port. 

## Building

### Docker image

```shell
export VERSION=1
docker build -t mhalbritter/http-echo:latest -t mhalbritter/http-echo:"$VERSION" .
```

### Docker image with multi-platform support

Prepare the environment:

```shell
docker buildx create --name mybuilder --driver docker-container --bootstrap
docker buildx use mybuilder
```

Build the image and push it:

```shell
export VERSION=1
docker buildx build --platform linux/amd64,linux/arm64 -t mhalbritter/http-echo:latest -t mhalbritter/http-echo:"$VERSION" --push .
```
