# Backend-Master-Class-Golang-Postgres-Kubernetes-gRPC-

### Github lecture

https://github.com/techschool/simplebank

API commit lecture:
https://github.com/techschool/simplebank/commits/master/?before=931b0d98159595bac39f569a4a0cef2235c750de+105

Operation commit lecture:
https://github.com/techschool/simplebank/commits/master/?before=931b0d98159595bac39f569a4a0cef2235c750de+70

# About App

### Docs

DB diagram:
https://dbdiagram.io/d/%5Btutorial%5DSimple-bank-65ffa82dae072629ceccb8e2

## Get start

- Install packages: `go mod tidy`

## Run app locally

### Start DB

- Update env to use local db path at tutorial-simple-bank/app.env
- `docker run --name simple-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres`
- `make server`

### Run test

- `make postgres`
- `make test`

## Run app with Docker

- Update env to use docker db path at tutorial-simple-bank/app.env
- Build first time: `docker compose build -d`
- Start docker: `docker compose up -d`
- Stop docker: `docker compose down`

## Migrate DB

- Create migration: `migrate create -ext sql -dir db/migration -seq example_name`
- Migrate up: `make migrateup`
- Migrate down: `make migratedown`

## RPC

- Generate proto: `make proto`
- Create Grpc once, can run both rpc and http: https://github.com/grpc-ecosystem/grpc-gateway
