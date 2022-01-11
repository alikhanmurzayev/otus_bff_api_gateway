# HW BFF. Api Gateway

### Clone the repo:
```bash
git clone https://github.com/alikhanmurzayev/otus_bff_api_gateway.git && cd otus_bff_api_gateway
```

### start minikube
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

### Setup ambassador:
Obtain a new licence https://www.getambassador.io/aes-community-license-renewal/
```bash
helm install -n ambassador --set licenseKey.value=eyJhbGciOiJQUzUxMiIsInR5cCI6IkpXVCJ9.eyJsaWNlbnNlX2tleV92ZXJzaW9uIjoidjIiLCJjdXN0b21lcl9pZCI6ImplZG93YXQ3MDlAdmViMzQuY29tLTE2NDE3NDM0ODUiLCJjdXN0b21lcl9lbWFpbCI6ImplZG93YXQ3MDlAdmViMzQuY29tIiwiZW5hYmxlZF9mZWF0dXJlcyI6WyIiLCJmaWx0ZXIiLCJyYXRlbGltaXQiLCJ0cmFmZmljIiwiZGV2cG9ydGFsIl0sImVuZm9yY2VkX2xpbWl0cyI6W3sibCI6ImRldnBvcnRhbC1zZXJ2aWNlcyIsInYiOjV9LHsibCI6InJhdGVsaW1pdC1zZXJ2aWNlIiwidiI6NX0seyJsIjoiYXV0aGZpbHRlci1zZXJ2aWNlIiwidiI6NX0seyJsIjoidHJhZmZpYy11c2VycyIsInYiOjV9XSwibWV0YWRhdGEiOnt9LCJleHAiOjE2NzMyNzk0ODUsImlhdCI6MTY0MTc0MzQ4NSwibmJmIjoxNjQxNzQzNDg1fQ.BhQi32pQiNR9KWO8iPblts39iU2asq7x7yrWHpIV4_n7wdpVrDHTjAmRRowu_FDGXMkCrUVfDJuNheMrmHjMoA1avCKYlL9E-xv1oQwrg6kFuTaS3xHL6rP9VEX7aawobtoybOpXZiorr0W3W6hvER1yihwbxjCW1dcpYjD1lZ9I-5qd9na1fWrKD1L37Oxm5kTixPKF7usbAfxpZN-cqVRPIpivHsqsPbsh-AB44hPWDKnnh3qZwUQ39Y0lw3z1R7cWrqet4G_yYcIU-KQZFQVXRMlL9FCCwgS9g_xHe8Xx65k3xxXIjKt-RKwZuRc7q6S-ACeWPNwcfdvW0vVH4w \
-f k8s/ambassador.yaml ambassador datawire/ambassador
```

Ensure that ambassador is up
```bash
kubectl get po -n ambassador
```

### Setup routes and auth rules
```bash
kubectl apply -f k8s/ambassador-routes.yaml -f k8s/ambassador-auth.yaml
```

### Run test:
```bash
newman run --global-var "baseUrl=$(minikube service list|grep 'ambassador'|grep 'http'|grep -Eo 'http://[^ >]+'|head -1)" OTUS.API_Gateway.postman_collection.json
```
