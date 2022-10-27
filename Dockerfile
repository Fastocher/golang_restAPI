FROM golang:alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

COPY init.sql /docker-entrypoint-initdb.d/10-init.sql

# RUN apt-get update
# RUN apt-get -y install postgresql-client

RUN apk update
RUN apk add postgresql-client

RUN chmod +x wait-for-postgres.sh


RUN go mod download
RUN go build -o restapp ./cmd/main.go

CMD ["./restapp"]
