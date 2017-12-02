FROM ubuntu:14.04
MAINTAINER "srikantha.muvva@gmail.com"

RUN apt-get update && apt-get upgrade -y
# Install golang
RUN apt-get install golang-go -y

# Add webserver binary
RUN mkdir -p /opt/gowebsrv/bin
ADD webserv /opt/gowebsrv/bin

# Expose port 8080
EXPOSE 8080

ENTRYPOINT /opt/gowebsrv/bin/webserv
