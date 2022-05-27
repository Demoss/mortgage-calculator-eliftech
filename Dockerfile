FROM golang:1.17-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-postgres.sh

RUN go mod download

RUN go build -o calc ./cmd/main.go

CMD ["./calc"]