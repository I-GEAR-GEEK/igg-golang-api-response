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
    status, resData := response.InternalServerError("Update profile error", "error message something", "UPDATE_PROFILE_ERROR")
    return c.Status(status).JSON(resData)
}
```

## Function
```golang
response.Item(item interface{}, message string)
response.Success(item interface{}, message string)
response.Pagination(data Pagination, message string)
response.Accepted(item interface{}, message string)
response.Created(message string, item interface{})
response.Updated(message string, item interface{})
response.Deleted(message string)
response.BadRequest(message string, code string)
response.Unauthorized(message string, code string)
response.Forbidden(message string, code string)
response.NotFound(message string, code string)
response.ValidateFailed(message string, errorMessages interface{})
response.InternalServerError(message string, detail string, code string)
```