# Sample gRPC Project

# Requirement
* Golang latest version
* Protocol Buffer Compiler
* gRPC Golang plugin

# How to setup
1. Run this command for install all dependencies
```
make init
```

2. Run this command for generate gRPC code
```
make generate
```

3. Run this command for run server
```
cd server && go run *.go
```

4. Run this command for run client
```
cd client && go run main.go
```
