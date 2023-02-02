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
└── cmd
|   ├── migration
|   |   └── migration.go
|   └── main.go
├── config
|   ├── db
|   |   ├── db.go
|   |   └── postgresql.go
|   ├── cors.go
|   └── routers.go
├── internal
|   ├── contoller
|   |   ├── article
|   |   │   └── article.go
|   |   ├── root
|   |   │   └── root.go
|   |   └── user
|   |       └── auth.go
|   ├── domain
|   │   ├── article
|   │   │   ├── model
|   │   │   │   ├── article_request.go
|   │   │   │   ├── article_response.go
|   │   │   │   └── article.go
|   │   │   ├── repository
|   │   │   │   └── article.go
|   │   │   └── service
|   │   │       └── article.go
|   │   └── user
|   │       ├── model
|   │       │   ├── user_request.go
|   │       │   ├── user_response.go
|   │       │   └── user.go
|   │       ├── repository
|   │       │   └── user.go
|   │       └── service
|   │           └── user.go
|   ├── middleware
|   │   └── auth.go
|   └── router
|       └── main.go
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
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
└── README.md
```

## How To
### Running The App
- First get the dependencies with this command:
```shell
go mod download && go mod verify
```

- Copy the `.env.example` to `.env` with run this command:
```shell
make env
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