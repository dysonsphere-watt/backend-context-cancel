# Backend Context Cancel

A simple project repo for testing different Golang frameworks for support with early request termination.

## Start

```bash
git clone https://github.com/dysonsphere-watt/backend-context-cancel
cd backend-context-cancel
go mod tidy
go run .
```

After starting it, you can make `GET` API requests to ports `9001-9003`.
Each endpoint will have a 10s database call which includes the context propagation.
