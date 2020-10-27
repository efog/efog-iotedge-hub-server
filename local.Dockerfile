FROM golang:1.15.3-buster as build
ENV BACKEND_PORT=56789
ENV BACKEND_HOST=localhost
ENV FRONTEND_HOST=12345
ENV FRONTEND_PORT=12345
EXPOSE 12345
EXPOSE 56789
RUN apt update && apt upgrade -y
RUN apt install libzmq3-dev pkg-config -y
WORKDIR /go/src/server
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["efog-iotedge-hub-server"]