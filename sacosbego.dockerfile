FROM alpine:latest

RUN mkdir /app

COPY main /app

CMD ["app/main"]