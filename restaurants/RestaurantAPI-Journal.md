# Restaurant Journal

## Steps for development of Restaurant API and setting it up on AWS using docker

## Setup mongodb cluster
```
Create mongodb replica set of 3 instances
```

## Write GO APIs for Users for - 
```
1. List All Restaurants
2. Insert new Restaurant
3. Update Restaurant
4. Delete Restaurant
5. List Restaurant by Zipcode
```

## Create Docker Image
```
docker build -t sarthaksugandhi9/restaurant:v1.0
push docker image to AWS
docker push sarthaksugandhi9/restaurant:v1.0
```

## set up docker image on AWS
```
1. Launch Amazon Linux AMI
2. Configure Security group
3. Install docker on amazon linux ami
4. Pull docker image from docker hub
5. Run docker conatiner
```

## Auto scale restaurant docker instance

## Set up application load balancer
```
1. Create application load balancer
2. Create target group for Restaurant instance
3. Add targets in target group
4. Configure paths to redirect to restaurant instance
```
