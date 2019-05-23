# Running User Microservice

## Running user microservice on local environment
### Clone code
```
git clone <git-url>
```

### Set GOPATH variable
```
Use below command to set GOPATH from terminal
export GOPATH="path to project root directory"
```

### Build Application
```
Run - 
1. cd Users/user-go-api/src/users
2. go build
```

##W Run Application
```
Run - ./users
```

## Creating and running user microservice in docker environment
### Set up docker environment
```
1. Install docker (Refer - https://docs.docker.com/glossary/?term=install )
2. Start docker service
```

### Create docker image
```
1. cd User/user-go-api/src/users
2. Edit docker ENV variables with your database hosts and credentials
2. docker build . -t user
3. docker images
```

### Run docker image
```
1. docker run -d --network user --name mongodb mongo (if using docker for mongo database)
1. docker run -d -p 8000:8000 --network user --name user-container user
```
