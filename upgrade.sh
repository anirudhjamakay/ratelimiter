#!/bin/bash

set -e

APP_NAME="ratelimiter"
IMAGE_NAME="ratelimiter:latest"
DEPLOYMENT_NAME="ratelimiter"
NAMESPACE="default"
LOCAL_PORT=8080
REMOTE_PORT=8080

GOOS=linux GOARCH=amd64 go build -o app ./cmd/server
docker build -t $IMAGE_NAME .


CONTEXT=$(kubectl config current-context)

if [[ $CONTEXT == *"minikube"* ]]; then
    minikube image load $IMAGE_NAME
elif [[ $CONTEXT == *"kind"* ]]; then
    kind load docker-image $IMAGE_NAME
else
    echo "Using Docker Desktop Kubernetes"
fi

kubectl rollout restart deployment/$DEPLOYMENT_NAME -n $NAMESPACE

kubectl rollout status deployment/$DEPLOYMENT_NAME -n $NAMESPACE

kubectl port-forward deployment/$DEPLOYMENT_NAME $LOCAL_PORT:$REMOTE_PORT -n $NAMESPACE > portforward.log 2>&1 &

sleep 2

echo "Port-forward running"
echo ""
echo "Access your API at:"
echo "http://localhost:$LOCAL_PORT/hello"
echo ""

kubectl logs -f deployment/$DEPLOYMENT_NAME -n $NAMESPACE