# Kubernetes Golang Demo

Just an example I wanted to play with to try out some new tools and areas of
tooling that have changed significantly since last I used them. This is a very
simple golang http server. I mostly wanted to try out ko as the image builder
without requiring a Dockerfile for the project, kind as the deployment target
for a local kubernetes cluster (also running cloud-provider-kind to give a way
to access the service without needing to port forward) and the kustomize tool
now built directly into kubectl to deploy it to the local cluster.

## Running

When I do this I push to my personal dockerhub account so that the kubernetes
nodes can pull the images. You'll definitely have to change this to run it if
you want to try it. But here's an example anyway.

* make sure you're logged into dockerhub to be able to push images
* start up a kind cluster
* run the cloud-provider-kind program (I do this in another terminal)
* KO_DOCKER_REPO=mikerowehl/kdemo ko build --bare --tags=latest --platform=linux/amd64
* cd base
* kubectl apply -k .

When the service comes up with an external IP, you should be able to hit the
service on that IP. Note though that the kind cloud provider doesn't work
like most ingress controllers. It binds to the service port.

There's nothing special here, just wanted a working sample to start so that
I can pick up some other tools and experiment. Pushing just so I have a simple
containerized app to in a git repo to work with.
