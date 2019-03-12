#!/bin/bash
set -euo pipefail

http2_curl() {
	curl --silent -w "\n%{http_code}\n" --http2-prior-knowledge "$@"
}

echo "# Deleting resources if any"
kubectl delete --filename . &> /dev/null || true
sleep 10

echo -e "# Creating resources"
kubectl apply --filename . > /dev/null
sleep 10

cluster_ip=$(kubectl get service --all-namespaces | grep ingressgateway | tr -s ' ' | cut -d' ' -f5)

echo "# Curling"
echo "## Authority + Content-Type"
http2_curl ${cluster_ip} --header 'Host: hello-http2.example.com' \
                         --header 'Authority: hello-http2.example.com' \
                         --header 'Content-Type: application/json'

echo "## Authority"
http2_curl ${cluster_ip} --header 'Host: hello-http2.example.com' \
                         --header 'Authority: hello-http2.example.com'

echo "## Content-Type"
http2_curl ${cluster_ip} --header 'Host: hello-http2.example.com' \
                         --header 'Content-Type: application/json'
