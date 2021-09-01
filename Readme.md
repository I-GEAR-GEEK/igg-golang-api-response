# IGG Golang Standard API Response Format

## Install
```command
go get github.com/igeargeek/igg-golang-api-response
```

## How to use

```golang
package main

import "github.com/igeargeek/igg-golang-api-response/response"

func Simple(c context) error {
    status, resData := response.InternalServerError("")
    return c.Status(status).JSON(resData)
}
```

## Function
```golang
response.Unauthorized(message string)
response.NotFound(message string)
response.BadRequest(message string)
response.UnProcessableEntity(message string)
response.ValidateFailed(errorMessages interface{}, message string)
response.InternalServerError(message string)
response.Created(item interface{}, message string)
response.Updated(item interface{}, message string)
response.Deleted(message string)
response.Item(item interface{}, message string) // response get by id
response.Collection(items interface{}, message string) // response get all
response.Success(item interface{}, message string)// response normal
response.Pagination(data Pagination, message string)
```