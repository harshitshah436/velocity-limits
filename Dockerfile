FROM golang:1.16.4-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/velocity-limits

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Run Unit Tests
WORKDIR /app/velocity-limits/test
ENV CGO_ENABLED 0
RUN go test -v ./...

# Build the Go app
WORKDIR /app/velocity-limits/cmd/velocity-limits
RUN go build -o ./out/velocity-limits .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./out/velocity-limits"]