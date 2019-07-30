# Use a tagged version instead of latest to avoid regressions
# Use alpine for faster pulls on slow connections - like mine
FROM golang:1.12.7-alpine3.10

# Copy the local package files to the container's workspace.
WORKDIR /root/twelvefa/
COPY . .

# Get the depencies
RUN go get

# Build the app for alpine, executable name: twelvefa
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o twelvefa

# Best image I found to deploy to GCP
# Use alpine, again, for faster pulls
# Use a multi-stage Dockerfile to avoid anything unecessary in the final image
FROM google/cloud-sdk:255.0.0-alpine

WORKDIR /root/

# Store the CircleCI build number in the environment for easier debugging
ARG CIRCLECI_BUILD_NUMBER
ENV CIRCLECI_BUILD_NUMBER=${CIRCLECI_BUILD_NUMBER}

# Copy the executable from the previous stage
COPY --from=builder /root/twelvefa/twelvefa /root

# Run the app when the container starts
ENTRYPOINT /root/twelvefa

EXPOSE 80