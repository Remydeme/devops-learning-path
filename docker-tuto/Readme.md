# DOCKER TUTORIAL & DEPLOY ON AWS 


## go-server

* 1 - In this project we build a simple golang server 
with one endpoint.

* 2 - We build a Dockerfile

* 3 - Docker build command build an image of our server

* 4 - We use AWS EB to deploy our simple api

[AWS EBS](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/single-container-docker-configuration.html#create_deploy_docker_image_dockerrun)

[Docker tuto](https://docker-curriculum.com)


## Docker command 

### run

* docker run -it name/image sh => Run an image in interactive mode.

* docker run -d name/image     => Run a container in background (d stand for detached)

* docker run -d -P --name alias_name name/image => launch the conatainer in bacground mode and name this container
  alias_name

### network

* docker network ls  => display existing networks 

* docker network inspect network_name => display information about the docker network 

### container 

* docker container ls => display all the running container

* docker container logs name => display all the logs about this conatainer 

### rm or stop 

* docker stop $(docker ps -q) => stop all the running instance 

* docker rm $(docker ps -a -q -f status="exited") => rm all the stopped container 



