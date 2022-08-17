# kubernetes deployment detail

gcloud init

gcloud container clusters get-credentials <cluster-name> --region <region-name> --project <project-name>

e.g.

    gcloud container clusters get-credentials test-tweet-service --region us-central1 --project shekhar-sandbox-1


# create pod

kubectl create -f tweet-service-pod.yaml

# create service

kubectl create -f tweet-service-service.yaml


# delete service

kubectl delete service tweet-service

# delete pod

kubectl delete pod tweet-service


