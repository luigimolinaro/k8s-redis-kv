# DevOps task

## The task

### install.sh

This bash script install 
* Kubernetes with kind
* Build the "web services written in go"
* Run the "web service" on port 8080


  -------------------
  |  web service go | --> K8s cluster --> Redis 
  ------------------


## Next step 
* Create registry
* Create a container (I written a primitiva in Dockerfile) 
* Use a pods for the web service
* Expose webservice port on a LOADBALANCER service

## How to test 

### POST
curl -H "Content-Type: application/json" -d '{
        "name": "Luigi",
       "surname": "Molinaro"
       }' http://localhost:8080/posts

### GET
curl http://localhost:8080/posts 

## GET SINGLE ID
curl http://localhost:8080/posts/1

## GET ALL ID's
curl http://localhost:8080/posts/ 
