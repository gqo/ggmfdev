# ggmf.dev

Graeme's personal website, duh.

# Setup

[This](https://go.dev/doc/manage-install) page and [this](https://go.dev/doc/install) page are the basis for these go install recommendations.

Download the latest Linux Go tarball. At the time of writing, it's version `1.27.2`. Run the following:

```
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.27.2.linux-amd64.tar.gz
``` 

Then add this info to your `~/.bashrc` file:
```
# go
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH/bin
```

Then install the version of go running in production:
```
go get golang.org/dl/go1.16.2
go install golang.org/dl/go1.16.2
go1.16.2 download
```

Then run the server (points at `:8080` by default) from root using this command:
```
go1.16.2 run cmd/ggmfdev/ggmf.go
```

# Deployment

If you're me in the future or some weirdo other than me who wants to host my website, my recommendations for simple deployment are:

Using [Caddy](https://caddyserver.com/) (specifically `v2.6.4`), run the following command:
```
caddy reverse-proxy --from https://ggmf.dev --to localhost:8080
```

Might need to run as `sudo` for binding permissions. Might also need to kill `nginx` or `caddy` itself if they're already running as `systemctl` services.

Then, in a separate terminal session (try `screen`), run the website as normal.
