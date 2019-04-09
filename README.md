# go-web-server

This Go webserver compiles with Go version 1.12+.  

It exposes a web server on port 8080 and logs to STDOUT.  The port is configurable by setting the PORT environment variable.  

It only has one endpoint that returns status code 200: `/health`. All other requests will return 404: "404 page not found".  