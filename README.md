# simple-reverse-proxy
This is a simple reverse proxy written in Go. The problem it solves is that when developing a web application, one frequently ends up with addresses in one's browser like `local.example.com:8000` instead of the cleaner `local.example.com`. The reason that many web apps run on ports other than 80 is that use of lower numbered ports requires a super-user on Unix systems. To work around this, just run simple-reverse-proxy as super user with `sudo simple-reverse-proxy`.

##Installation
First install [Go](http://golang.org) and set your `GOPATH` environmental variable to the directory you would like the project saved in. Then run `go install github.com/carlmjohnson/simple-reverse-proxy`. The binary will be installed in `$GOPATH/bin`.

##Usage
Run `simple-reverse-proxy -h` to see usage help. Remember to use super-user if you want to listen on port 80, the standard web port.
