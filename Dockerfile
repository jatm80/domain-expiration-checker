FROM golang:alpine AS build-env
RUN apk add git
ARG version=0.0.0
WORKDIR /app
COPY . .
RUN go build -o /go/bin ./...

FROM alpine:latest
ARG version
ENV APP_VERSION=$version
WORKDIR /app
COPY --from=build-env /go/bin/domain-expiration-checker .
ENTRYPOINT ["./domain-expiration-checker"]