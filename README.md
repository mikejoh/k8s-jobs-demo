# Kubernetes Jobs

_Created for work related testing of Kubernetes Jobs_

## Components
* Kubernetes Deployment and different types of job manifests
* `job-logger`, small web server written in Golang that basically logs different messages depending on which endpoint/path you access with `GET` 
* `job-sender`, small Golang program that sends a HTTP requests that are based on environment variables exposed to the Pod (in this case)

## How-to
1. `docker login` to your Docker Hub account
2. `make dockerize`, this will build and push the two images needed (`job-sender` and `job-logger`)
3. Reference your Docker Hub account and image names in all Kubernetes manifests. Or keep the images provided by me.
4. `kubectl apply` the Deployment, make sure that the `job-logger` web server starts and works as expected.
5. `kubectl apply` one job manifest at a time, tail the logs of the `job-logger` so that you can see the results. 

The end result will be a number of log messages logged by `job-logger` for every Kubernetes Job run and the combination of `completions` and `parallellism` settings on the Job.