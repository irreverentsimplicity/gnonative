# Go server

This Go executable run a Gnomobile service.## Run the service

`go run ./goserver uds`

or

`go run ./goserver tcp`

The gRPC server prints the TCP address/socket path it listens to:

`server UDS path: xxx` or `server TCP address: xxx`

To close it, press Ctrl+C in the terminal.

## Usage

`go run ./goserver`