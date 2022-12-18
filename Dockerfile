FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY wait-for-postgres.sh ./
RUN chmod +x wait-for-postgres.sh

RUN apt-get update
RUN apt-get -y install postgresql-client

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/app/main.go
RUN go build -v -o /usr/local/bin/mig ./migrations/*.go

CMD ["app"]