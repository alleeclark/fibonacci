# Simple GO Lang REST API

> Simple RESTful API to get the sequence of fibonacci values

# Requirements

  - Golang version 1.9+ for building without exe

# Installing

  - Install Dependecies
    ```shell
    $ go get -u github.com/gorilla/mux
    ```

## Quick Start
### On windows start server with 
``` bash
.\main.exe
```

### Run with go installed
``` bash
go run ./main.go
```

## Endpoints

### Post to get Sequence of fibonacci values
``` bash
POST fib
http://localhost:8080/fib
# Request sample
# {
#   "number":45,
# }
```

```
# Running Test

  - Unit test
    ```shell
    $ cd controllers
    $ go test
    ```
