# Basic App

## Deploy
```
helm upgrade -i -n basic-app-primary --create-namespace basic-app-primary .

helm upgrade -i -n basic-app-secondary --create-namespace basic-app-secondary .
```

## Validate
```
curl http://basic-app-istio-gateway-basic-app-primary.china-dev01.dev.infra.webex.com/readiness

curl https://basic-app-primary.china-dev01.dev.infra.webex.com/readiness

nslookup ingress0.public.china-dev01.dev.infra.webex.com
```