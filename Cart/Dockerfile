FROM golang:latest
EXPOSE 8000

WORKDIR /go/src/app

ENV "MONGO_SERVER"="mongodb" "MONGO_DATABASE"="burger" "MONGO_COLLECTION"="Users" "MONGO_ADMIN_DATABASE"="admin" "MONGO_USERNAME"="mongo-admin" "MONGO_PASS"="cmpe281"

COPY . .

RUN go get -d -v

RUN /go/src/app

RUN go build