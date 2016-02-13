# Running the server
To run this server, from the root of this project:
```sh
go run ./src/*.go
```

# Test with go
After setup the server:
```sh
go run ./client/main.go
```

# Test with curl
## POST
```sh
curl https://test161.ops-class.org/api-v1/submissions -X POST -d '{"ID": 124, "name": "name"}'
```

## GET
```sh
curl https://test161.ops-class.org/api-v1/submissions/123
```
