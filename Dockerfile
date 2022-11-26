FROM golang:1.18-bullseye

# Working Directory
RUN mkdir /src
WORKDIR /src

# Create log directory inside the container for iris and file logging
RUN mkdir -p log
RUN mkdir -p log/filebeat

# Install required files for filebeat and iris
RUN apt-get update && \
    apt-get -y install wget && \
    apt-get -y install bash && \
    apt-get -y install nano && \
    apt-get -y install telnet && \    
    apt-get -y install supervisor

# Copy files
ADD src .
RUN go mod download

# add supervisor conf
ADD conf.d/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Setting up environment variable
ENV TZ="Asia/Tehran"

# Expose port
EXPOSE 8080

# build
RUN go build -o build/ cmd/server/server.go
RUN go build -o build/ cmd/worker/worker.go

# Moved to compose
# ENTRYPOINT ["/usr/bin/supervisord"]