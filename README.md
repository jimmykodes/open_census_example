# Open Census Example

This is a trivial example of setting up a Go project that uses [Open Census](https://github.com/census-instrumentation/opencensus-go)
for metrics, and a Prometheus exporter to view them.

## Running Locally
All the docker ups and downs can be handled through the makefile.
```shell script
make # this will do all the initial setup and start the docker containers and subscribe to its logs

make rebuild # after making changes to the go code, this will rebuild the go binary & image

make stop # stop the docker containers
make down # stop and remove containers

make logs # subscribe to the container logs
```

## Accessing locally
The _very_ basic API can be accessed at `http://localhost:8080`. The metrics can be accessed at `http://localhost:8081/metrics`.

You will need to send traffic to the API before all the logs will populate on the `/metrics` page.
