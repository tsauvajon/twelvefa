# Use a tagged version instead of latest to avoid regressions
FROM golang:1.12.7-buster as builder

# Use alpine for faster pulls on slow connections - like mine
# temporarily not using alpine to test Go modules
# FROM golang:1.12.7-alpine3.10 as builder

# Copy the local package files to the container's workspace.
WORKDIR /root/twelvefa/
COPY . .

# Set a location to install dependencies
ENV GOBIN=/go/bin/

# Get the dependencies
RUN ./install.sh

# Re generate the proto files
RUN ./generate.sh calc

# Build the app for alpine, executable name: cli
RUN cd cli && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o cli

# Use alpine for faster pulls and a smaller image
# Use a multi-stage Dockerfile to avoid anything unecessary in the final image
FROM golang:1.12.7-alpine3.10

WORKDIR /root/

# Store the CircleCI build number in the environment for easier debugging
ARG CIRCLECI_BUILD_NUMBER
ENV CIRCLECI_BUILD_NUMBER=${CIRCLECI_BUILD_NUMBER}

# Copy the executable from the previous stage
COPY --from=builder /root/twelvefa/cli/cli /root