# DevOps task

## The task

### install.sh

This bash script install 
* Kubernetes with kind
* Build the "web services written in go"
* Run the "web service" on port 30080 (install.sh echo the right IP)
* Install Graphana and Prometheus

## Next step 

* Makes some security check
* Expose webservice port on a LOADBALANCER service

## How to test 

### POST
curl -H "Content-Type: application/json" -d '{
        "name": "Luigi",
       "surname": "Molinaro"
       }' http://[IP]:8080/posts

### GET
curl http://[IP]:8080/posts 

## GET SINGLE ID
curl http://[IP]:8080/posts/1

## GET ALL ID's
curl http://[IP]:8080/posts/ 
