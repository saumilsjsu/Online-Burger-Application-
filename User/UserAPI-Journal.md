# User Journal

## Steps for development of User API and setting it up on AWS using docker

## Setup mongodb cluster
```
Create replica set of 3 mongodb cluster
```

## Write GO APIs for Users for - 
```
1. User signup
2. User signin
3. User update
4. Delete User
5. List Users
```

## Create Docker Image
```
docker build -t sarthaksugandhi9/user:v1.0
push docker image to AWS
docker push sarthaksugandhi9/user:v1.0
```

## set up docker image on AWS
```
1. Launch Amazon Linux AMI
2. Configure Security group
3. Install docker on amazon linux ami
4. Pull docker image from docker hub
5. Run docker conatiner
```

## Auto scale user docker instance

## Set up application load balancer
```
1. Create application load balancer
2. Create taregt group for User instance
3. Add targets in target group
4. Configure paths to redirect to user instance
```
