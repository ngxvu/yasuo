FROM golang:1.18 as build-env

ARG VERSION=0.0.0

# cache dependencies first
WORKDIR /code
COPY go.mod .
COPY go.sum .
RUN go mod tidy

# lastly copy source, any change in source will not break above cache
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -a -ldflags="-s -w -X main.version=${VERSION}" \
  -o app .

# runtime
FROM alpine:3.16

RUN apk --update add ca-certificates

WORKDIR /code

COPY --from=build-env /code/conf conf
COPY --from=build-env /code/app app
COPY --from=build-env /code/resources resources

CMD ["/code/app"]

