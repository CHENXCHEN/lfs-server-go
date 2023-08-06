LFS Server Go
======
## Introduction
This is based on [git-lfs / lfs-test-server](https://github.com/git-lfs/lfs-test-server).

[lfs]: https://github.com/github/git-lfs
[api]: https://github.com/github/git-lfs/tree/master/docs/api#readme

LFS Server Go is an server that implements the [Git LFS API][api]. It
is intended to be used for testing the [Git LFS][lfs] client and is not in a
production ready state.

LFS Server Go is written in Go, with pre-compiled binaries available for Mac,
Windows, Linux, and FreeBSD.

See [CONTRIBUTING.md](CONTRIBUTING.md) for info on working on LFS Server Go and
sending patches.

## Different With lfs-test-server
ChangeLog:
+ Isolation (based on user/repo)
  + Data directories are isolated by user/repo
  + Object ID is isolated by user/repo
  + Lock is isolated by user/repo
+ Config File Support
  + Use the `LFS_SERVER_GO_CONFIG` environment variable to specify the configuration file, by default is `config.ini`

## Installing

Use the Go installer:

```
  $ go install github.com/CHENXCHEN/lfs-server-go
```


## Building

To build from source, use the Go tools:

```
  $ go get github.com/CHENXCHEN/lfs-server-go
```


## Running

Running the binary will start an LFS server on `localhost:8080` by default.
There are few things that can be configured via `config.ini` file:

```ini
[Main]
; Port to listen on
; The host used when the server generates URLs, default: "localhost:8080"
Listen = tcp://:8080
; Host address - used for downloading
; The host used when the server generates URLs, default: "localhost:8080"
Host = 127.0.0.1:8080
; consider lfs-server-go may behind a reverse proxy
; ExtOrigin =
; login for the admin user
; An administrator username, default: not set
AdminUser = admin
; An administrator password, default: not set
AdminPass = admin
; path to ssl certificate
;Cert = somekey.crt
; path to ssl key
;Key = somekey.key
Scheme = http
; Should the contents be public?
Public = true
; Database Configuration
; The database file the server uses to store meta information, default: "lfs.db"
MetaDB = lfs.db
; Content Store Configuration
; The path where LFS files are store, default: "lfs-content"
ContentPath = lfs_content

; Tus Configuration
; set to 'true' to enable tusd (tus.io) resumable upload server; tusd must be on PATH, installed separately
UseTus = false
; The host used to start the tusd upload server, default "localhost:1080"
TusHost = localhost:1080
```

If the `ADMINUSER` and `ADMINPASS` variables are set, a
rudimentary admin interface can be accessed via
`http://$LFS_HOST/mgmt`. Here you can add and remove users, which must
be done before you can use the server with the client.  If either of
these variables are not set (which is the default), the administrative
interface is disabled.

To use the LFS Server Go with the Git LFS client, configure it in the repository's `.lfsconfig`:


```
  [lfs]
    url = "http://localhost:8080/user/repo"

```

HTTPS:

NOTE: If using https with a self signed cert also disable cert checking in the client repo.

```
  [lfs]
    url = "https://localhost:8080/user/repo"

  [http]
    sslverify = false

```


An example usage:


Generate a key pair
```
openssl req -x509 -sha256 -nodes -days 2100 -newkey rsa:2048 -keyout mine.key -out mine.crt
```

Download Config file And Modified

```shell
wget https://raw.githubusercontent.com/CHENXCHEN/lfs-server-go/main/config.example.ini -O config.ini
```

Make yourself a run script

```
#!/bin/bash

set -eu
set -o pipefail

LFS_SERVER_GO_CONFIG=config.ini ./lfs-server-go
```

Build the server

```
go build

```

Run

```
bash run.sh

```

Check the managment page

browser: https://localhost:8080/mgmt


