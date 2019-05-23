FROM golang:latest 

WORKDIR /go/src/app

ENV "SERVER"="mongodb" "DATABASE"="burger" "COLLECTION"="Users" "ADMIN_DATABASE"="admin" "USERNAME"="mongo-admin" "PASSWORD"="cmpe281"

EXPOSE 8003

COPY . .

RUN cd /go/src/app

RUN go get -d -v

RUN go build

ENTRYPOINT ["./app"]