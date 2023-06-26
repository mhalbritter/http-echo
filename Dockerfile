FROM golang:alpine AS builder
ARG TARGETOS
ARG TARGETARCH
ENV USER=user
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR $GOPATH/src/github.com/mhalbritter/http-echo/
COPY . .
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-w -s" -o /go/bin/http-echo

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/http-echo /go/bin/http-echo
USER user:user
EXPOSE 8080
ENTRYPOINT ["/go/bin/http-echo"]
