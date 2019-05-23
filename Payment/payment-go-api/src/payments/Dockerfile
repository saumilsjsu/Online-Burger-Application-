FROM golang:latest 

WORKDIR /go/src/app

ENV "MONGO_SERVER"="10.0.4.181" "MONGO_DATABASE"="burger" "MONGO_COLLECTION"="payment" "ADMIN_DATABASE"="admin" "MONGO_USERNAME"="mongo-admin" "MONGO_PASS"="cmpe281"

EXPOSE 8002

COPY . .

RUN cd /go/src/app

RUN go get -d -v

RUN go build

ENTRYPOINT ["./app"]