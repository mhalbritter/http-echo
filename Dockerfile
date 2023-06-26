FROM golang:alpine AS builder
ARG TARGETOS
ARG TARGETARCH
ARG USER=user
ARG UID=10001
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
ARG UID=10001
ENV HTTP_ECH0_PORT=80
ENV HTTP_ECH0_TLS_PORT=443
EXPOSE ${HTTP_ECH0_PORT}
EXPOSE ${HTTP_ECH0_TLS_PORT}
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
USER user:user
WORKDIR /go/bin/
COPY --chown=${UID}:${UID} server.crt server.crt
COPY --chown=${UID}:${UID} server.key server.key
COPY --from=builder --chown=${UID}:${UID} /go/bin/http-echo http-echo
ENTRYPOINT ["/go/bin/http-echo"]
