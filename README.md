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

## Run app locally

### Start DB

- Update env to use local db path at tutorial-simple-bank/app.env
- `docker run --name simple-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres`

### Run test

- `make postgres`
- `make test`

## Run app with Docker

- Build first time: `docker compose build -d`
- Start docker: `docker compose up -d`
- Stop docker: `docker compose down`
