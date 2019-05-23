FROM golang:latest 

WORKDIR /go/src/app

ENV "MONGO_SERVER"="mongodb" "MONGO_DATABASE"="burger" "MONGO_COLLECTION"="Users" "MONGO_ADMIN_DATABASE"="admin" "MONGO_USERNAME"="mongo-admin" "MONGO_PASS"="cmpe281"

EXPOSE 8000

COPY . .

RUN cd /go/src/app

RUN go get -d -v

RUN go build

ENTRYPOINT ["./app"]

