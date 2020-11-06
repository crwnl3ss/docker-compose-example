FROM golang:1.15.3-alpine3.12

ARG suffix=""

WORKDIR /tmp/
COPY ./src/main.go /tmp
RUN CGO_ENABLED=0 go build -o hello_world &&\
  chmod +x hello_world &&\
  cp hello_world /usr/bin/

ENTRYPOINT /usr/bin/hello_world --suffix=${suffix}
