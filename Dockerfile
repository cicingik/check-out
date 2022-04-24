FROM golang:1.16.15-alpine3.14 as builder
ARG COMMIT_MSG
ARG COMMIT
ARG GIT_TAG
ARG BUILD_TAG
ARG APP_ENTRYPOINT=./cmd/checkout

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY ./go.mod ./go.sum /app/

RUN go mod download

COPY . /app
RUN set -exuo pipefail \
    && go build -a -i \
        -ldflags "-w -s \
            -X \"github.com/cicingik/check-out.CommitMsg=${COMMIT_MSG}\" \
            -X \"github.com/cicingik/check-out.CommitHash=${COMMIT}\" \
            -X \"github.com/cicingik/check-out.AppVersion=${GIT_TAG}\" \
            -X \"github.com/cicingik/check-out.ReleaseVersion=${BUILD_TAG}\"" \
        -o checkout \
        ${APP_ENTRYPOINT}

FROM alpine:3.14
LABEL author="dany <danysatyanegara@gmail.com>" \
      version="v1.1.0-beta" \
      description="docker image for check-out"
ARG CO_HTTP_PORT=2777
ENV CO_HTTP_PORT_BIND=$CO_HTTP_PORT

WORKDIR /app/
COPY --from=builder /app/checkout .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY data /app/data

EXPOSE ${CO_HTTP_PORT_BIND}
CMD [ "/app/checkout" ]