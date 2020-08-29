FROM golang:alpine

RUN apk update && apk add git

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o aqilliz_app

ENTRYPOINT [ "./aqilliz_app" ]