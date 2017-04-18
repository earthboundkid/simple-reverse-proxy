# simple-reverse-proxy
This is a simple reverse proxy written in Go. The problem it solves is that when developing a web application, one frequently ends up with addresses in one's browser like `local.example.com:8000` instead of the cleaner `local.example.com`. The reason that many web apps run on ports other than 80 is that use of lower numbered ports requires a super-user on Unix systems. To work around this, just run simple-reverse-proxy as super user with `sudo simple-reverse-proxy`.

## Installation
First install [Go](http://golang.org).

If you just want to install the binary to your current directory and don't care about the source code, run

```bash
GOBIN=$(pwd) GOPATH=/tmp/gobuild go get github.com/carlmjohnson/simple-reverse-proxy
```

## Usage
Run `simple-reverse-proxy -h` to see usage help. Remember to use super-user if you want to listen on port 80, the standard web port.
