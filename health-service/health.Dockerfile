FROM golang:1.24.5

WORKDIR /app

COPY . .

RUN go build -o service

EXPOSE 7001

CMD ["./service"]