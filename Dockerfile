FROM golang:latest

RUN apt update
RUN mkdir /app
WORKDIR /app
ADD . /app

CMD ["go", "run", "main.go"]