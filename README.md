# Backend-Master-Class-Golang-Postgres-Kubernetes-gRPC-

### Github lecture

https://github.com/techschool/simplebank

API commit lecture:
https://github.com/techschool/simplebank/commits/master/?before=931b0d98159595bac39f569a4a0cef2235c750de+105

Operation commit lecture:
https://github.com/techschool/simplebank/commits/master/?before=931b0d98159595bac39f569a4a0cef2235c750de+70

### Docs

DB diagram:
https://dbdiagram.io/d/%5Btutorial%5DSimple-bank-65ffa82dae072629ceccb8e2

### Run Docker

`docker run --name simple-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres`

### Run test

- `make postgres`
- `make test`
