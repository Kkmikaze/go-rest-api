# Go Rest API Documentation

## Table Of Content
- [Prerequisite](#prerequisites)
  - [Spesification](#spesification)
  - [Layout](#layout)
- [How To](#how-to)
- [References](#references)
  - [GIT style](#git-style)


## Prerequisites
What things you need to setup the application:

### Spesification
This application uses:
- Gin
- Gorm
- PostgreSQL

### Layout
```
.
├── app
|   └── contoller
|       ├── root
|       │   └── root.go
|       └── user
|           └── auth.go
├── config
|   ├── collection
|   │   └── main.go
|   ├── middleware
|   │   └── auth.go
|   ├── cors.go
|   └── routes.go
├── db
|   ├── migrate
|   │   └── migrate.go
|   ├── db.go
|   └── postgresql.go
├── domain
|   └── user
|       ├── model
|       │   ├── user_request.go
|       │   ├── user_response.go
|       │   └── user.go
|       ├── repository
|       │   └── auth.go
|       └── auth.go
├── lib
|   ├── auth
|   │   └── context.go
|   ├── constant
|   │   └── error.go
|   ├── encrypt
|   │   └── encrypt.go
|   ├── env
|   │   └── env.go
|   └── response
|       └── response.go
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md
```

## How To
### Running The App
- First get the dependencies with this command:
```shell
go mod tidy
```

- Copy the `.env.example` to `.env` with run this command:
```shell
cp .env.example .env
```

- for migrate database schema use this command:
```shell
make run migrate
```

- and for running the application can use this command:
```shell
make run
```

## References
### GIT Style
For commit message style or git style guide, use this doc
- [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)