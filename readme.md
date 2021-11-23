# Scaling Funds on GKE

A simple "Hello scaling funds" application written in Golang and Deployed on GKE

I deployed this app on GKE as instructed with the help of terraform.

Dockerized tghe application and pushed the docker image to my docker repository.

Deployed the appliction with the help of kubectl.

Used github actions as the CI/CD tool of choice.

After the application is deployed, fetch the k8s services and use the loadbalancer IP and port number to see the contents of the page. A simple json payload.

## Gotchas

- I did not use a private cluster and a better load balancer. To make it better, i'd use a fully fledged google load balancer, or any other of my choice as they are now supported in the current version of kubernetes.
- I used docker registry simply because the service account file was giving me a lot of json debugging issues in the deployment spec as well as terraform cloud. Ran away from them as i couldn't debug them in time.
- I added the back bone infrastructure code along with the applicaiton code and differentiate them with 2 different workflow files which are triggered when certain files are changed. Ideally the back bone infra should reside in its own repo and maintained separately.
- Currently the app is not running because i tore down the cluster for cloud cost reasons.

Terraform configurations files can be found in the folder `./terraform-configs` and the kubernetes deployment files can be found in the folder `./k8s-configs` and the workflow file in `./.github/workflows`
