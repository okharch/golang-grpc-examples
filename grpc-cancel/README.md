# gRPC Time Streaming Example with Go

This example demonstrates the proper method to halt a streaming service from the client side using gRPC in Go. The key to achieving this is through context cancellation. By canceling the context from the client side, the server is notified and can appropriately terminate the stream.

In this particular example, the server streams the current time to the client. After a period of 5 seconds, the client initiates context cancellation, which subsequently stops the time streaming from the server. This showcases an elegant way to manage control flow between client and server in real-time streaming scenarios.

## Getting Started

### Prerequisites

- Go 1.14 or higher
- Protocol Buffers v3
- gRPC-Go

### Installing

1. Install Protocol Buffers v3 and Go plugin:

    ```shell
    $ go get -u google.golang.org/protobuf/cmd/protoc-gen-go
    $ go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
    $ export PATH="$PATH:$(go env GOPATH)/bin"
    ```

2. Generate Go code from the protobuf file:

    ```shell
    $ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative time_stream.proto
    ```

3. Install the Go dependencies:

    ```shell
    $ go mod tidy
    ```

### Running

1. Start the server:

    ```shell
    $ go run server.go
    ```

2. In a new terminal window, run the client:

    ```shell
    $ go run client.go
    ```

The client will print the time received from the server every second. After 5 seconds, the client will cancel the context and stop receiving the time.
