# golang-watchlist-presentation

## Running the slideshow
Make sure you have Go installed (see below).

If you don't have `present` on your `$PATH`, run
```bash
make install-present
```

Then run
```bash
make
```

## Install Golang
```bash
# install go
brew install go
export GOPATH=$HOME/golang
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```
