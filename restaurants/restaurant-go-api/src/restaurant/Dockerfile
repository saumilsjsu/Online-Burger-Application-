FROM golang:latest 

WORKDIR /go/src/app

ENV "SERVER"="10.0.4.181" "DATABASE"="burger" "COLLECTION"="order" "ADMIN_DATABASE"="admin" "USERNAME"="mongo-admin" "PASSWORD"="cmpe281"

EXPOSE 8001

COPY . .

RUN cd /go/src/app

RUN go get -d -v

RUN go build

ENTRYPOINT ["./app"]

