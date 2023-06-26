# HTTP echo

## Building

```shell
docker buildx create --name mybuilder --driver docker-container --bootstrap
docker buildx use mybuilder
docker buildx build --platform linux/amd64,linux/arm64 -t mhalbritter/http-echo:latest -t mhalbritter/http-echo:1 --push .
```
