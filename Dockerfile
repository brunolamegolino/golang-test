FROM golang:latest

WORKDIR /go

USER 10

CMD [ "./command.sh" ]