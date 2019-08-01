# Interview with Ori (ori.co) — Technical Task

## Summary

This repository contains a microservices and a CLI to use this service, following
the [12factor](https://12factor.net/) principles.

The CLI allows you to connect to a gRPC server: `cli connect $address`. Once
successfully connected, you enter the interactive mode where you can enter 5
commands:
- `exit` to quit the interactive mode
- `add x y` to add two integers together: `add 2 5     // 7`
- `max x y` to return the max of two integers: `max 2 5     // 5`
- `np x [y [z...]]` to return the nth prime numbers: `np 2 12      // 3 37`
- `help` to display a help message

More about the nth prime number: the first prime numbers are 2 3 5 7.
Therefore, the 1st prime number is 2, the 2nd prime number is 3, the 3rd is 5 and
so on. The command allows you to fetch any number of nth prime numbers at a time.

## Repository structure

`./calc` contains the 'business implementation' - in reality, returning basic values
from inputs. It also contains the protobuf definition and the generated protobuf
`.pb.go` files.
I chose to use [Protocol buffers](https://developers.google.com/protocol-buffers/),
because it plays well with gRPC and allowed me to develop a client/server
prototype quickly.

The root repository contains the repository CI configuration (scripts, Docker, gitignore...)
as well as the server code. The server API is defined in `server.go` and the service
entrypoint is in `main.go`.

The CLI can be found in the eponym folder: `./cli`. The client API is defined in
`client.go`, the interactive commands are isolated in `./commands` and the logic
is in the entrypoint, `cli.go`.

## Run using Docker Compose

```
docker-compose up -d twelvefa # backend
docker-compose run calcli     # cli
./cli help
./cli connect twelvefa:80
> max 2 6
> np 9999 1234 1 2 3 4
> exit
```

## Run from source

```
# install the dependencies
./install.sh

# regenerate the protobuf files
./generate.sh calc

# test
go test . ./calc

# benchmark
go test ./calc -bench=.

# build
go build

# test the client
PORT=3001 ./twelvefa & # run the service in the background to test the client
TWELVEFA_ADDRESS=:3001 go test -v ./cli
pkill ./twelvefa

# run the backend
PORT=3000 ./twelvefa

# open a new terminal
cd ./cli
go build
./cli connect :3000
> max 2 6
> np 9999 1234 1 2 3 4
> exit
```

## Deploy

Most of the steps for creating the GCP project are described in `init-gcp.sh`.  
This file has not been tested, consider running each command manually.

After the project is correctly configured, the service is deployed on GKE
by CircleCI.

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

Q: **Prove how it fits and uses the best cloud native understanding**

A: This applications leverages the Cloud by abstracting all of the infrastructure
and networking, allowing this repository to focus on the actual code. The
infrastructure, network. build steps are not manually configured and deployed,
but merely described through configuration files (Docker/Kubernetes/Terraform...).

The different parts of the application are loosely coupled, and are easy to
replace, update or horizontally scale by simply adding more machines (note: at a
bigger scale, other bottlenecks will appear, such a monitoring/databases).

## Event Store

Q: **How would you expand on this service to allow for the use of an eventstore?**

A: Instead of having queries sent directly from a CLI to a backend, we could use an
event store. Each command would create a new event. For example, running the command
`max 2 99` would create a new event that would look something like:  
`
  {
    type: command
    id: 1
    command: max
    parameters: {
      a: 2
      b: 99
    }
  }
`  
The first available backend would then execute the command, and send another event:  
`
  {
    type: result
    command_id: 1
    result: 99
  }
`  
The CLI would then receive this event by looking for results for the command_id `1`,
and display the result: `99`.

With the Pub Sub (Publisher Subscriber) architecture, a publisher publishes on one
or several topics, and the subscribers subscribe to any topic they want.  
In this example, CLIs would publish on the `command` topic and subcribe to the
`result` topic. The backends would do the opposite.

To go further, we could imagine a logging backing that would subscribe to both
topics and log all of the events, commands and results.

## External Access

Q: **How would this service be accessed and used from an external client from
the cluster?**

A: To access the service from an external client from the cluster, the schema
would be basically the same for any Cloud Provider:
- create a Public IP address
- associate it with the Kubernetes cluster
- configure an Ingress to redirect the traffic to the service

Example ingress.yaml:
```
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: twelvefa-ingress
spec:
  backend:
    serviceName: twelvefa
    servicePort: 80
```

## Configuration

Environment variables to set.

### In CircleCI

Example configuration:
- `GOOGLE_CIRCLECI`: GCP CircleCI service account credentials (json content)
- `GOOGLE_TERRRAFORM`: GCP Terraform service account credentials (json content)
- `GOOGLE_CLUSTER_NAME`: `twelvefa-gke-cluster`
- `GOOGLE_COMPUTE_ZONE`: `us-east1-c`
- `GOOGLE_REGION`: `us-east1`
- `GOOGLE_PROJECT_ID`: `ori-tsauvajon`

### Local environment

Example configuration:

```
export PROJECT_ID=ori-tsauvajon
export TF_CREDS=./creds/terraform.json
export CIRCLECI_CREDS=./creds/circleci.json
export PORT=3000 # what port will the service run on locally?
export GOOGLE_COMPUTE_ZONE=us-east1-c
export GOOGLE_COMPUTE_REGION=us-east1
```

## Production

`kubectly get services`

NAME         | TYPE       | CLUSTER-IP     | EXTERNAL-IP |  PORT(S)  |  AGE
-------------|------------|----------------|-------------|-----------|-----
kubernetes   | ClusterIP  | 10.11.240.1    | \<none>      |  443/TCP  |  29m
twelvefa     | ClusterIP  | 10.11.242.192  | \<none>      |  3000/TCP |  16m

`kubectly get pods`

NAME                       | READY  | STATUS   | RESTARTS  | AGE
---------------------------|--------|----------|-----------|-------
calcli-7484f77f44-jzz7k    | 1/1    | Running  | 0         | 15m
calcli-7484f77f44-rd4nl    | 1/1    | Running  | 0         | 15m
twelvefa-7bb6f58659-2lpjb  | 1/1    | Running  | 0         | 15m
twelvefa-7bb6f58659-8q2r4  | 1/1    | Running  | 0         | 15m
twelvefa-7bb6f58659-n22vn  | 1/1    | Running  | 0         | 15m

## Next steps

- Add end to end tests (run a server, run commands in the CLI, check results)
- Automate project creation, API enabling... with Terraform
- Cache the dependencies for faster Docker builds (both locally and in the CI)
- Use a configuration manager: use Vault or KMS to manage configs
- Create a staging environment to allow for better end-to-end testing before deployment
- Think about rate limiting at some point
- Improve logging: write more logs, e.g. log errors
- Improve monitoring: use New Relic, Data Dog or grafana to create useful dashboards

## Sources

- 12factor: https://12factor.net/
- Go Modules: https://github.com/golang/go/wiki/Modules
- Protobuf:
  - Documentation: https://developers.google.com/protocol-buffers/
  - Tutorial: https://medium.com/@shijuvar/building-high-performance-apis-in-go-using-grpc-and-protocol-buffers-2eda5b80771b
- Testing: https://golang.org/pkg/testing/
- NthPrimesNumber:
  - Euler challenge: https://projecteuler.net/problem=7
  - Sieve of Erathostenes: https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
- Event Sourcing: https://martinfowler.com/eaaDev/EventSourcing.html
- Google Container Registry: https://cloud.google.com/container-registry/docs/
- Terraform: https://www.terraform.io/docs/providers/google
