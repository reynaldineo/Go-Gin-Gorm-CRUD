# GO Gin GORM CRUD
Simple CRUD in Go using Gin Framework, GORM for ORM and postgresql for the database
## Initial Set Up

### Init go module

```
go mod init github.com/reynaldineo/Go-Gin-Gorm-CRUD`
```

Get dependecies, CompileDaemon for auto run and build

```
go get github.com/githubnemo/CompileDaemon
```

Install, so it can be use in terminal/CLI

```
go install github.com/githubnemo/CompileDaemon
```

Get dependencies, godotenv for procces .env files

```
go get github.com/joho/godotenv
```

### Get Gin Web Framework for Golang

```
go get -u github.com/gin-gonic/gin
```

### Get GORM, ORM library for Golang

```
go get -u gorm.io/gorm
```

### Get Database driver for postgres

```
go get -u gorm.io/driver/postgres
```

## Run CompileDaemon for auto run and build

```
CompileDaemon -command="./go-gin-gorm-crud"
```
