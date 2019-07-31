# Interview with Ori (ori.co) — Technical Task

## Summary

[TODO]

## Usage

```
# install the dependencies
./install.sh

# generate the .pb.go files
./generate.sh calc

# test
go test . ./calc

# benchmark
go test ./calc -bench=.

# run a cli
docker-compose up
docker exec -it calcli /bin/bash
```

## 12 Factors

### I. Codebase
*One codebase tracked in revision control, many deploys*

The code is tracked in a source repository, here using git and hosting it on
[GitHub](https://github.com/tsauvajon/twelvefa). Other possibilities: other
git hosting applications (GitLab, BitBucket...), SVN.

### II. Dependencies
*Explicitly declare and isolate dependencies*

Using [Go Modules](https://github.com/golang/go/wiki/Modules), all dependencies
are declared and documented with their version in the `go.mod` file.
Other possibilities: go dep.

### III. Config
*Store config in the environment*

Instead of using configuration files, or even worse, storing the configuration
in the code, environment variables are read when the service is started.
Better alternatives: use a configuration manager, such as Vault, KMS or Chef.

### IV. Backing services
*Treat backing services as attached resources*

Although not used in this particular repository, I could use backing services such
as a relational database (PostgreSQL) a search engine (Elastic Search) or a
key-value store (Redis). Instead of hard-coding the address of these services,
they should be attached and detached as needed, using for example configuration
or a service discovery tool.

### V. Build, release, run
*Strictly separate build and run stages*

This service is first tested and built (using CircleCI), deployed to a container
registry, and finally deployed in the production Kubernetes cluster.

### VI. Processes
*Execute the app as one or more stateless processes*

Using the same input will always produce the same output. There is nothing
it this application that uses any kind of state, for example a session.

### VII. Port binding
*Export services via port binding*

The application runs on the local port defined by `$PORT` (default: 80). Using
Docker, we bind this port to any port we want. For example, if I want to run this
service locally on port 3000, I'll bind the host port 3000 on the container's
port 80: `docker run -p 3000:80`.

### VIII. Concurrency
*Scale out via the process model*

Many nodes are run concurrently, and the queries are load balanced between them,
by Kubernetes. Scaling vertically is possible, by just using bigger machines. A
better way of scaling would be to scale vertically, by adding more nodes (machines)
or more replicas of the same container.

### IX. Disposability
*Maximize robustness with fast startup and graceful shutdown*

Because we don't have a lot of overhead, the startup is quite fast. We just
instantiate a few structs, start listening on the preferred port and the service
is running.
Interruption signals are caught, and we make sure the server is stopped
gracefully before actually exiting the program (see `main.go`)

### X. Dev/prod parity
*Keep development, staging, and production as similar as possible*

Using Docker and Docker-Compose for local development, and Terraform/Kubernetes for
production, I am able to achieve relative similarity in both environments.

Any other environment that would need to be added (next one would be staging),
would be deployed with Terraform and would be the same as production.

### XI. Logs
*Treat logs as event streams*

Logs are written to the standard output, `stderr`. They can be read with
`docker logs ...` for example.

### XII. Admin processes
*Run admin/management tasks as one-off processes*

Deploying/updating/tearing down the infrastructure is done using Terraform.
Therefore, each task is a single command to be run.

Generating the executable file is done through Docker, and is again a one-off
process.

Every other task that would require several actions will be included in a single
executable bash file for that task, that would in turn run all of the actions.

Future tasks could include database migrations, installing multiple dependencies.

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

- Create a secure connection with 
- Cache the dependencies for faster Docker builds
- Improve monitoring: use New Relic, Data Dog or grafana to create useful dashboards
- Improve logging: write more logs, e.g. log errors
- Use a configuration manager: use Vault or KMS to manage configs
- Create a staging environment: duplicate the infrastructure, config and
deployments to allow for better end-to-end testing before deployment
- As the service or the ecosystem grows bigger, consider deployment/testing
strategies that scale better (blue-green, canary...)
- Think about rate limiting at some point

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
