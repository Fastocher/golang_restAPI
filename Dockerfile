FROM golang:1.19.2

RUN go version
ENV GOPATH=/

COPY ./ ./

COPY init.sql /docker-entrypoint-initdb.d/10-init.sql

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh


RUN go mod download
RUN go build -o restapp ./cmd/main.go

CMD ["./restapp"]
