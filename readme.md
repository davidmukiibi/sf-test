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


To setup and run the whole setup, setup github secrets below:

- `DEFAULT_IMAGE` this is the image the application deployment will launch with
- `DEPLOYMENT_NAME` the name of your deployment in the k8s deployment spec
- `DOCKER_EMAIL` the docker email of your docker account
- `DOCKER_PASSWORD` the docker password of the same account
- `DOCKER_USERNAME` the docker username of the same account
- `GKE_CLUSTER` the name of your gke cluster
- `GKE_LOCATION` the region in which your cluster resides
- `GKE_PROJECT` the gcp project name
- `GKE_SA_KEY` the serviceaccount json key
- `TFE_TOKEN` the terraform cloud token

After that, head over to terraform cloud at `app.terraform.io` and add this secret:

- `GOOGLE_CREDENTIALS` the serviceaccount json key

With the above secrets set, ideally you should head into the `terraform-configs` and make a change, be it adding an empty new line and commit and push. This will trigger the build for terraform and will build or modify your GKE infrastructure. But as pointed out above, the json key gave me hard time with formatting as i saved it in terraform cloud.

When the infrastructure has been setup successfully, `cd` into `k8s-configs` and make a change in there too, be it an empty new line. This will trigger the github actions workflow that will create/update the kubernetes resources decalred in your `.yaml` config files with your GKE cluster.

When the deployments are up and running, simply run `kubectl get services`, since it's in the default namespace, it will simply return the load balanced service with the external IP and port that can be used to access the deployment.

In the app, i have the path `/scalingfunds` setup to display a pleasant message.
