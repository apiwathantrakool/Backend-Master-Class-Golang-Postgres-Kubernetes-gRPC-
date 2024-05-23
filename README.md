# Backend-Master-Class-Golang-Postgres-Kubernetes-gRPC-

### Github lecture

https://github.com/techschool/simplebank

Previous commit lecture:
https://github.com/techschool/simplebank/commits/master/?after=931b0d98159595bac39f569a4a0cef2235c750de+104

### Run Docker

`docker run --name simple-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres`

### Run test

- `make postgres`
- `make test`
