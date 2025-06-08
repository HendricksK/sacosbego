FROM golang:latest

WORKDIR /

# RUN go mod download

RUN go build -o /app/main

EXPOSE 9000

CMD ["/app/main"]