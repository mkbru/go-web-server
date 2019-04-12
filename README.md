# go-web-server

## Instructions

Write a Dockerfile that compiles and then runs the web server in this repository.  

Write it as if the container will be used in production.  

-----

## Notes
This Go webserver compiles with Go version 1.12+.  

It exposes a web server on port 8080 and logs to STDOUT.  The port is configurable by setting the PORT environment variable.  

It only has one endpoint that returns status code 200: `/health`. All other requests will return 404: "404 page not found". 


## Instructions for running the docker container

1. Make sure Docker is running
2. Clone this repo `git clone https://github.com/mkbru/fullscript.git`
3. Create the image `docker build -t go-web-server .`
4. Check to make sure the image was created `docker images | grep go-web-server`
5. Run the container `docker run -p :8080 go-web-server`
6. `curl localhost:8080/health`

