FROM golang:1.21.0
WORKDIR /app
ENTRYPOINT [ "tail", "-f", "/dev/null" ]