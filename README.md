# Go Rest API Documentation

## Table Of Content
- [Prerequisite](#prerequisites)
  - [Spesification](#spesification)
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