# Go-Clean-Architecture  

REST API Server

## Setup

```
docker-compose up -d
```

## Used 

- Architecture : Clean Architecture
- Language : Golang
- Web Framework : Gin
- Relational Database : MySQL
- ORM : GORM
- Development environment : Docker

## For what  

Understand the basics of go and clean architecture.  

## Response example  

#### Create a new user
```
POST /users?Name=piyo

http status 201, [{"ID":"2","Name":"piyo"}]
```

#### Get all users
```
GET /users

http status 200, [{"ID":"1","Name":"hoge"},{"ID":"2","Name":"piyo"}]
```

#### Get a user
```
GET /users/1

http status 200, [{"ID":"1","Name":"hoge"}]
```

#### Update a user
```
PUT /users/2?Name=Apex  

http status 201, [{"ID":"2","Name":"Apex"}]
```

#### Delete a user
```
DELETE /users/2

http status 200, [{"ID":"2","Name":"Apex"}]