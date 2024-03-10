[![Go Reference](https://pkg.go.dev/badge/github.com/rezakhademix/govalidator.svg)](https://pkg.go.dev/github.com/rezakhademix/govalidator) [![Go Report Card](https://goreportcard.com/badge/github.com/rezakhademix/govalidator)](https://goreportcard.com/report/github.com/rezakhademix/govalidator) [![codecov](https://codecov.io/gh/rezakhademix/govalidator/graph/badge.svg?token=BDWNVIC670)](https://codecov.io/gh/rezakhademix/govalidator) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) 

## govalidator

This is a Golang Validator package without any type assertion or reflection that provides data validation.

## Requirements

Go 1.22 or above.

## Getting Started

GoValidator includes a set of validation rules and a handy `check()` method for defining any custom rule and You can use it to describe how a value should be considered valid.

### Installation

Run the following command to install the package:

```
go get github.com/rezakhademix/govalidator
```

## govalidator examples:

1. simple:
   ```go
       type User struct {
           Name string `json:"name"`
           Age unit    `json:"age"`
       }

       var user User
       _ := c.ShouldBind(&user)  // error ignored for simplicity

       v := govalidator.New()

       v.RequiredInt(user.Age, "age").        // age can not be null or 0
         MinInt(int(user.Age), 18, "age")     // minimum value for age must be 18
         RequiredString(user.Name, "name")    // name can not be null or ""
         MaxString(user.Name, 50, "name")     // maximum allowed charactars for name field is 50

         if v.IsFailed() {
             return v.Errors()  // will return failed validation errors
         }
   ```
2. with custom field names and messages:

   ```go
    type User struct {
        name string `json:"name"`
    }

    var user User
    _ := c.ShouldBind(&user)  // error ignored for simplicity

    v := govalidator.New()

    v.MaxString(user.name, "first_name", "please fill first_name field") // with custom field name and custom validation message

    if v.IsFailed() {
        return v.Errors()
    }
   ```

3. advanced usage:
   You can define any custom rules or any flexible rule that does not exist in default govalidator package. Simply use `check()` method to define your desired data validations:

```go
    type Profile struct {
       Name   string
       Age    int
       Score  int
       Status []string
    }

   var profile Profile

   // after filling profile struct data with binding or other methods

   v := govalidator.New()

   v.Check(profile.Name != "", "name", "name is required")  // check is a method to define rule as first parameter and then pass field name and validation error message

   v.Check(profile.Age > 18, "age", "age must be greater than 18")


   // we just need to pass a bool as a rule
   // `checkScore()` is method that checks score validation rules and returns a bool
   v.Check(checkScore(), "score", "score must")

   // More Complex rules:
   // define a custom `In()` method thats returns a bool as result, e.g:
   func In[T comparable](value T, permittedValues ...T) bool {
	    for i := range permittedValues {
		    if value == permittedValues[i] {
			    return true
		    }
	    }

	    return false
   }

   statuses := []string{"active", "inactive", "pending"}

   v.check(In[ProfileStatuses](profile.Status, statuses), "status", "status must be in: ...")

   if v.IsFailed() {
       return v.Errors()
   }
```

more informations and cool things comming soon.
