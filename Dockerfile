FROM golang:latest

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go mod download
RUN go build -o api . 
EXPOSE 8080

CMD [ "/app/api" ]