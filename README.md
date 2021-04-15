# Gostorage

#### Performs tasks:
1. Uploading files;
1. Deleting files;
1. Serving uploaded files.

#### Some information about service:
* Swagger Open API documentation: `host:port/docs/index.html`;
* Static files serving: `host:port/static/...`
* Makefile for fast using commands: `./Makefile`.

#### Dependencies:
* gorilla/mux - as request router and dispatcher for matching incoming requests to their respective handler;
* Golangcilint - a set of linters for writing good code in Go;

#### Configuration:
* Environment variables:
```bash
CONFIG_PATH=/usr/src/app/configs/dev.toml
```

* Configuration .toml file:
```toml
bind_host = "localhost:2045/storage" # Server host
bind_addr = ":8080" # Server port
log_level = "debug" # Logging level
write_timeout = 15 # Request write timeout (seconds)
read_timeout = 15 # Request read timeout (seconds)
idle_timeout = 60 # Request idle timeout (seconds)

# Paths for saving uploaded files
file_path = '/static/'
```
