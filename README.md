# token-pooling
Token pooling system

A simple microservice for serving a set of tokens. Tokens resources are pre fetched from db.

## Steps to run
```
cd $GOPATH/src/github.com/arjunksofficial/
git clone https://github.com/arjunksofficial/token-pooling.git
cd $GOPATH/src/github.com/arjunksofficial/token-pooling
```

To install Postgres via docker
```
docker-compose up
```

create database **token** in the postges server via commandline or pgadmin,etc.

Run migration and start server

```
go get bitbucket.org/liamstask/goose/cmd/goose
source config/local.env
cd $GOPATH/src/github.com/arjunksofficial/token-pooling/migration
goose up
cd ..
go test
go run main.go
```

Server has 12 tokens loaded in DB with LRU Cache of limit 12

Check endpoints

[localhost:8080/token](localhost:8080/token)

[localhost:8080/stats](localhost:8080/stats)
