# HW BFF. Api Gateway

### Clone the repo:
```bash
git clone https://github.com/alikhanmurzayev/otus_bff_api_gateway.git && cd otus_bff_api_gateway
```

### Start minikube
```bash
minikube start --cpus=4 --memory=4g --vm-driver=virtualbox
```

### Start services:
```bash
kubectl apply \
-f k8s/namespaces.yaml \
-f k8s/config.yaml \
-f k8s/postgres.yaml \
-f k8s/user-app.yaml \
-f k8s/auth-app.yaml
```

Ensure that services are running
```bash
watch kubectl get po -n default
```

### Setup ambassador:
Obtain a new licence https://www.getambassador.io/aes-community-license-renewal/
```bash
helm install -n ambassador --set licenseKey.value=<enter licence key here> \
-f k8s/ambassador.yaml ambassador datawire/ambassador
```

Ensure that ambassador is up
```bash
watch kubectl get po -n ambassador
```

### Setup routes and auth rules
```bash
kubectl apply -f k8s/ambassador-routes.yaml -f k8s/ambassador-auth.yaml
```

### Run test:
```bash
newman run  \
--global-var "baseUrl=$(minikube service list|grep 'ambassador'|grep 'http'|grep -Eo 'http://[^ >]+'|head -1)" \
OTUS.API_Gateway.postman_collection.json
```
