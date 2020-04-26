FROM golang:1.14.1-alpine3.11 AS build

# ARG http_proxy=http://proxy.com:8080
# ARG https_proxy=http://proxy.com:8080
ARG NAME=rest-go

WORKDIR ${GOPATH}/src/github.com/ks6088ts/${NAME}
ADD . .
RUN go build \
    -ldflags="-s -w" \
    -o /outputs/cli \
    ./cmd/${NAME}

# ---

FROM alpine:3.11.5
WORKDIR /app
COPY --from=build /outputs/cli .
CMD [ "./cli --help" ]
