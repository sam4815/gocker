#  gocker

## About

Containers are a fundamental part of my daily workflow, but I have little understanding of what's going on "under the hood". This project is my attempt to demystify containerization by building a (very) simple toy container in Go.

## Installation

Using [Podman](https://podman.io/) on Mac OS X, shell into a privileged Linux container with
```
podman run -ti --privileged -v $(pwd):/app -w /app golang:1.19-alpine /bin/sh
```

Run the Go container using
```
go run main.go run {COMMAND}
```

## Resources
1. https://www.infoq.com/articles/build-a-container-golang/
1. https://faun.pub/deep-into-container-build-your-own-container-with-golang-98ef93f42923

