# WOLWEB application

## Prerequirement
#### Dependencies:
```sh
go get
```
#### Make storage directory:
```sh
mkdir -p store
```
#### Fill following env or use default:
```
# optional
PORT = 80
```
```
# optional
NETWORK = 192.168.1.255
```
```
# optional
COOKIE_SECRET = autogen uuid() for start app
```
```
# optional
COOKIE_TTL = 86400 (24h)
```
```
# required
PASSWORD = enter here
```

## Docker run
When start container with `docker run` set option:  
Volume for persistent (optional)
```sh
-v host_path:/app/store
```
