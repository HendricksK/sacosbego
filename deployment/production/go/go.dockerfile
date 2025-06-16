FROM golang:latest 

RUN apt update -y && apt install supervisor -y

RUN mkdir -p /var/log/supervisor

WORKDIR /app

COPY /app/. .

COPY /proc/. /etc/supervisor/conf.d/

EXPOSE 9000

# CMD ["/usr/bin/supervisord"]

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/supervisord.conf"]

