# Order Journal

## Steps for development of Order API and setting it up on AWS using docker

## Setup mongodb cluster
```
Create replica set of 3 mongodb cluster
```

## Write GO APIs for Orders for - 
```
1. Order Status by OrderId
2. Order Status by UserId
3. Post Order
4. Get all Order details
```

## Create Docker Image
```
docker build -t saumil0257/order:v1.0
push docker image to AWS
docker push saumil0257/order:v1.0
```

## set up docker image on AWS
```
1. Launch Amazon Linux AMI
2. Configure Security group
3. Install docker on amazon linux ami
4. Pull docker image from docker hub
5. Run docker conatiner
```

## Auto scale order docker instance

## Set up application load balancer
```
1. Create application load balancer
2. Create taregt group for Order instance
3. Add targets in target group
4. Configure paths to redirect to order instance
```
