#!/bin/bash
#Install local kubernetes using kind https://kind.sigs.k8s.io/

#set -x

#check id root
if [[ $(id -u) -ne 0 ]] ; then echo "Please run as root" ; exit 1 ; fi

kind_create_cluster() {
  #Set-up Kind cluster
  kind create cluster --name $1
  kubectl cluster-info --context kind-$1
}

install_kind() {
  #Linux assumed
  curl -s -Lo ./kind "https://kind.sigs.k8s.io/dl/v0.10.0/kind-$(uname)-amd64"
  chmod +x ./kind
  mv ./kind /usr/local/bin/kind
}

#Pre-requisite check
which go || { echo "Please Install go..." 1>&2 ; exit 1; }
which docker || { echo "Please Install docker..." 1>&2 ; exit 1; }
which kubectl || { echo "Please Install kubectl..." 1>&2 ; exit 1; }
which kind || install_kind 

#Cluster Creation
kind_create_cluster tiko

#k8spart
kubectl apply -Rf k8s
APP_IP=$(kubectl get nodes -o wide | grep -i tiko | grep -Eo '[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}')
sed "s/REDIS_MASTER_SERVICE_HOST:DB_PORT/$APP_IP:30000/g" go-rest/db.go.template > go-rest/db.go
cd go-rest && go build 

#Examples : 

#echo "posting My name"
#
#curl -H "Content-Type: application/json" -d '{
#        "name": "Luigi",
#	"surname": "Molinaro"
#	}' http://localhost:8080/posts
#
#echo "getting my name"
#curl http://localhost:8080/posts | jq

set -
