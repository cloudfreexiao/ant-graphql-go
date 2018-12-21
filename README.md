# an-graphql-go

GraphQL API server in golang to get linux system info.

## Requirement

* golang installed v>=1.11
* add $GOPATH/bin to $PATH

```
export PATH=$PATH:$GOPATH/bin
```

* install go packages

```
go get github.com/msteinert/pam
go get github.com/spf13/cobra
```

* [option] install httpie

## Usage

* install go-bindata

```
make setup
```

* run server

```
make run
```

* build static binary

```
make
```

## Daemon Flags

```
Flags:
  -d, --debug          debug mode
      --disable-auth   disable auth middleware
  -h, --help           help for graphql-server
  -p, --port int       port number (default 9527)
```

** graphiql is only available when disable auth middleware **

## Controllers

### Ping

To check whether server is alive

```
curl http://localhost:9527/ping
```

### Login

Get auth token

```
http -v --json POST localhost:9527/login name=kevin passwd=somepassword
```

### Refresh Token

Refresh token before token expire

```
http -v -f POST localhost:9527/refresh_token "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDQ2MDkwNDAsIm5hbWUiOiJyb290Iiwib3JpZ19pYXQiOjE1NDQ2MDU0NDB9.a1DZTB17HXJrC
```

### Graphql

#### User

Get user

```
query {
    user(uid: "0") {
        uid
        gid
        name
        home
        groups {
            gid
            name
        }

    }
}
```

Get user by user name

```
query {
    userByName(name: "root") {
        uid
        gid
        name
        home
        groups {
            gid
            name
        }

    }
}
```

Get users

```
query {
    users {
        uid
        gid
        name
        home
        groups {
            gid
            name
        }

    }
}
```

#### CPU

Get CPU

```
query {
    cpu(id:"cpu1") {
        id
        total
        user
        system
        idle
        iowait
    }
}
```

Get CPUs

```
query {
    cpus {
        id
        total
        user
        system
        idle
        iowait
    }
}
```

#### Memory

Get Memory

```
query {
    memory {
        total
        free
        used
        shared
        buffer
        cache
        swap
    }
}
```

#### Network Iface

Get Iface

```
query {
    iface(name: "eno1") {
        name
        mac
        addrv4 {
            ip
            mask
        }
        addrv6 {
            ip
            mask
        }
        mtu
        rx
        tx
    }
}
```

Get Ifaces

```
query {
    ifaces {
        name
        mac
        addrv4 {
            ip
            mask
        }
        addrv6 {
            ip
            mask
        }
        mtu
        rx
        tx
    }
}
```

#### Service

Get service

```
query {
    service (name: "smb") {
        name
        mainPID
        activeState
        unitFileState
    }
}
```

Start service

```
mutation {
	startService (name: "smb") {
        name
        mainPID
        activeState
        unitFileState
    }
}
```

Stop service

```
mutation {
    stopService (name: "smb") {
        name
        mainPID
        activeState
        unitFileState
    }
}
```

Enable service

```
mutation {
    enableService (name: "smb") {
        name
        mainPID
        activeState
        unitFileState
    }
}
```

Disable service

```
mutation {
    disableService (name: "smb") {
        name
        mainPID
        activeState
        unitFileState
    }
}
```

## Query using curl

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "{ user(uid:\"0\") { uid, gid, name, home, groups { gid, name } } }"}' localhost:9527/graphql
```

## Todo
- netstat -anp | grep 9527 | wc -l
- ss -anp sport = :2057 | wc -l
- ps -elf | grep nginx | grep -w '20:58' | grep -v grep | wc -l
- netstat -anp | grep 9527
- 
- [x] Add mongodb ([mongodb-driver](https://github.com/mongodb/mongo-go-driver))
- [x] Add redis ([redis](https://github.com/go-redis/redis))
