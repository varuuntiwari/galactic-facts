FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /galactic-facts-api

EXPOSE 8080

CMD [ "/galactic-facts-api" ]