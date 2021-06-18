FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
RUN mkdir -p $GOPATH/src/api-key-value-service
ADD . $GOPATH/src/api-key-value-service
WORKDIR $GOPATH/src/api-key-value-service

RUN go get -d -v
RUN go build -o api-key-value-service .

# Stage 2
FROM alpine

RUN mkdir /app
COPY --from=builder /go/src/api-key-value-service/api-key-value-service /app/
COPY --from=builder /go/src/api-key-value-service/.env /app/
ARG APP_VERSION
ARG APP_NAME
ARG MODULE_NAME
ENV APP_VERSION=$APP_VERSION
ENV MODULE_NAME = $MODULE_NAME
ENV APP_NAME = $APP_NAME
EXPOSE 8080
WORKDIR /app
CMD ["./api-key-value-service"]