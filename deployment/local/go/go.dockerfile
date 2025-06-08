FROM golang:latest 

WORKDIR /app

COPY /app/. .

RUN go get

RUN go install 
    
RUN go build -o main 

EXPOSE 9000

CMD ["/app/main"]