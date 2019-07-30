# Interview with Ori (ori.co) — Technical Task

## Summary

[TODO]

## Usage

### Test

`go test ./calc`

### Benchmark

`go test ./calc -bench=.`

### Generate the .pb.go files

`./generate.sh calc`

### Debug

`docker-compose up`

## 12 Factors

### I. Codebase
*One codebase tracked in revision control, many deploys*

The code is tracked on [GitHub](https://github.com/tsauvajon/12fa). All deploys
are made from this single source.

### II. Dependencies
*Explicitly declare and isolate dependencies*

Using [Go Modules](https://github.com/golang/go/wiki/Modules), all dependencies
are declared and documented with their version in the `go.mod` file.

### III. Config
*Store config in the environment*

Instead of using configuration files, or even worse, storing the configuration
in the code, environment variables are read when the service is started.

### IV. Backing services
*Treat backing services as attached resources*

[TODO]

### V. Build, release, run
*Strictly separate build and run stages*

This service is first tested and built (using CircleCI), deployed to a container
registry, and finally deployed in the production kubernetes cluster.

### VI. Processes
*Execute the app as one or more stateless processes*

[TODO]

### VII. Port binding
*Export services via port binding*

The application runs on the local port defined by `$PORT` (default: 80). Using
Docker, we bind this port to any port we want. For example, if I want to run this
service locally on port 3000, I'll bind the host port 3000 on the container's
port 80: `docker run -p 3000:80`

### VIII. Concurrency
*Scale out via the process model*

[TODO]

### IX. Disposability
*Maximize robustness with fast startup and graceful shutdown*

[TODO]

### X. Dev/prod parity
*Keep development, staging, and production as similar as possible*

Using Docker and Docker-Compose for local development, and kubernetes for
production, I am able to achieve relative similarity in both environments.

Any other environment that would need to be added (next one would be staging),
would use kubernetes as well.

### XI. Logs
*Treat logs as event streams*

[TODO]

### XII. Admin processes
*Run admin/management tasks as one-off processes*

[TODO]

## Cloud Native

## Event Store

Q: **How would you expand on this service to allow for the use of an eventstore?**

A:

## External Access

Q: **How would this service be accessed and used from an external client from
the cluster?**

A: To access the service from an external client from the cluster, I would have
to add a Reverse Proxy. I could be using any web server or proxy, but my first
choice would be nginx as it is lightweight, easy to configure and plays well
in a microservices environment.
The schema would be basically the same for any Cloud Provider:
- create a Public IP address
- associate it with the Kubernetes cluster
- create a rule/security group to allow tcp 80 and/or tcp 443 through (or any other port used)
- redirect the traffic from this IP to the reverse proxy service (e.g. using an Ingress)
- add a rule in the reverse proxy to forward requests and responses between the
client and the internal service; any load-balancing rule would be useful here

This way, there is no direct access to the internal service, and the only open
route is through the reverse proxy.

## Next steps

[TODO]


## Sources used

- 12factor: https://12factor.net/
- Go Modules: https://github.com/golang/go/wiki/Modules
- Protobuf:
  - Documentation: https://developers.google.com/protocol-buffers/
  - Tutorial: https://medium.com/@shijuvar/building-high-performance-apis-in-go-using-grpc-and-protocol-buffers-2eda5b80771b
- Testing: https://golang.org/pkg/testing/
- NthPrimesNumber:
  - Euler challenge: https://projecteuler.net/problem=7
  - Sieve of Erathostenes: https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
