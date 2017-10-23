# sample-microservice

### Getting Started

To build the code, you must first install

- [Docker](https://www.docker.com/)
- [Go](https://golang.org/)
- [Minikuber](https://github.com/kubernetes/minikube)
- [Virtual Box](https://www.virtualbox.org/)

Or complete use [Tectonic](https://coreos.com/tectonic/)

**Make sure to read the slide for detail instruction (PDF file).**

### build Go file

Before build make sure you docker is up and running.

From command line Run ```./build```, only work on linux (**Sorry MS Window User**,
you need to create batch file or at least window 10 should work with the shell script)

If everything alights the docker image ~14mb as created and store in docker.

**Warning** change the desire version in the script or make it as argument with $1 ;)

### Deployment

Make sure your ```minikube``` is up and running, if not run the following
command line ```minikube start```

**Note** for release production or test environment should use ```deployment.yaml```
and then use command line ```kubectl apply -f deployment.yaml``` instead to deploy or
update deployment version.

In the command line, ```kubectl run sample2 --image=sample2:v1.0.0 --port=8080```

A new deployment should be available on the Kubernetes dashboard.

### Create Service (Cluster and Load Balance)

Run the following command ```kubectl apply -f service.yaml```

### Access Microservice

Run the following command ```minikube service sample2 --url```