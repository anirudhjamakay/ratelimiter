#!/bin/bash

set -e  # stop on error

APP_NAME="ratelimiter"
IMAGE_NAME="ratelimiter:latest"
DEPLOYMENT_NAME="ratelimiter"
NAMESPACE="default"

GOOS=linux GOARCH=amd64 go build -o app ./cmd/server
docker build -t $IMAGE_NAME .


CONTEXT=$(kubectl config current-context)

if [[ $CONTEXT == *"minikube"* ]]; then
    minikube image load $IMAGE_NAME
elif [[ $CONTEXT == *"kind"* ]]; then
    kind load docker-image $IMAGE_NAME
else
    echo "Docker Desktop Kubernetes (no extra load needed)"
fi

kubectl rollout restart deployment/$DEPLOYMENT_NAME -n $NAMESPACE

kubectl rollout status deployment/$DEPLOYMENT_NAME -n $NAMESPACE

kubectl logs -f deployment/$DEPLOYMENT_NAME -n $NAMESPACE
