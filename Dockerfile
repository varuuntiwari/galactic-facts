FROM golang:latest

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go mod download
RUN go build -o api . 

CMD [ "/app/api" ]