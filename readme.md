# SafeBoda on Minikube

A simple "Hello World" application written in Golang and Deployed on Minikube

## Step 1: Start Minikube

Run the command below in the your terminal:

```docker
minikube start && eval $(minikube docker-env)
```

## Step 2: Dockerize

Run the command below in the same directory that has the "dockerfile"

```docker
docker build -t mukiibi/safeboda:v10 .
```

## Step 3: Deploy to Minikube

```kubernetes
kubectl create secrets.yaml
kubectl create services.yaml
kubectl create ingress.yaml
kubectl create deployments.yaml
```

## SSL certificate
I generated self signed certificates for the domain "go-safeboda.info" using this site:

```
https://www.selfsignedcertificate.com/
```

I later encoded them using base64 with command:

```bash
cat go-safeboda.info.cert | base64
cat go-safeboda.info.key | base64
```

I then copied the results to the secrets file appropriately.

And Viola, you have your hello world application hosted on Minikube as I would in production.

In your web browser, visit:
```
go-safeboda.info
```
or
```
go-safeboda.info/safe
```
to see the magic!