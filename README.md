[![Go Reference](https://pkg.go.dev/badge/github.com/rezakhademix/govalidator.svg)](https://pkg.go.dev/github.com/rezakhademix/govalidator) [![Go Report Card](https://goreportcard.com/badge/github.com/rezakhademix/govalidator)](https://goreportcard.com/report/github.com/rezakhademix/govalidator) [![codecov](https://codecov.io/gh/rezakhademix/govalidator/graph/badge.svg?token=BDWNVIC670)](https://codecov.io/gh/rezakhademix/govalidator) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) 

# GoValidator
GoValidator is a data validation package that can sanitize and validate your data to ensure its safety and integrity as much as possible.

Our goal is to avoid using any type assertion or reflection for simplicity.
We would be delighted if you were interested in helping us improve the GoValidator package. Feel free to make your pull request.

# Getting Started
Go Validator includes a set of validation rules and a handy `check()` method for defining any custom rule.

## Installation
Run the following command to install the package:
```
go get github.com/rezakhademix/govalidator
```
## Import
Import package in your code-base by using:
```
import validator "github.com/rezakhademix/govalidator"
```

# Documentation
GoValidator package has the following features and method to use.
Each validation rule in GoValidator has it's own default message, e.g: `required` rule for a required `name` has the default message of: `name is required` but you can define your custom field name and message.

### Methods (also known as validation rules)
---

| Method | Description |
| - | - |
| RequiredInt       | `RequiredInt` checks if an integer value is provided or not.|
| RequiredFloat     | `RequiredFloat` checks if a float value is provided or not.|
| RequiredString    | `RequiredString` checks if a string value is empty or not.|
| RequiredSlice     | `RequiredSlice` checks if a slice has any value or not.|
| BetweenInt        | `BetweenInt` checks whether value falls within the specified range or not.|
| BetweenFloat      | `BetweenFloat` checks whether value falls within the specified range or not.|
| Date              | `Date` checks value to be a valid, non-relative date.|
| Email             | `Email` checks value to match `EmailRegex` regular expression.|
| Exists            | `Exists` checks given value exists in given table or not.|
| NotExists         | `NotExists` checks given value doesn't exist in given table.|
| LenString         | `LenString` checks length of given string is equal to given size or not.|
| LenInt            | `LenInt` checks length of given integer is equal to given size or not.|
| LenSlice          | `LenSlice` checks length of given slice is equal to given size or not.|
| MaxInt            | `MaxInt` checks given integer is less than or equal given max value.|
| MaxFloat          | `MaxFloat` checks given float is less than or equal given max value.|
| MaxString         | `MaxString` checks length of given string is less than or equal given max value.|
| MinInt            | `MinInt` checks given integer is greater than or equal given min value.|
| MinFloat          | `MinFloat` checks given float is greater than or equal given min value.|
| MinString         | `MinString` checks length of given string is greater than or equal given min value.|
| RegexMatches      | `RegexMatches` checks given string matches given regular expression pattern.|
| UUID              | `UUID` checks given value is a valid universally unique identifier (UUID).|
| When              | `When` will execute given closure if given condition is true.|

### Functions (other common validation rules)
---
| Method | Description |
| - | - |
| In        | `In` checks given value is included in the provided list of values.|
| Unique    | `Unique` checks whether values in the provided slice are unique.|


---
### Examples:
---
1. simple:
    ```go
        type User struct {
            Name string `json:"name"`
            Age unit    `json:"age"`
        }

        var user User
        _ := c.Bind(&user)  // error ignored for simplicity

        v := govalidator.New()

        v.RequiredInt(user.Age, "age", "").         // age can not be null or 0
            MinInt(int(user.Age), 18, "age", "")    // minimum value for age is 18

        v.RequiredString(user.Name, "name", "")     // name can not be null, "" or " "
            MaxString(user.Name, 50, "name", "")    // maximum acceptable length for name field is 50

        if v.IsFailed() {
            return v.Errors()  // will return failed validation error messages
        }
   ```

2. with custom field names and messages:
    ```go
        type User struct {
            Name string `json:"name"`
        }

        var user User
        _ := c.ShouldBind(&user)  // error ignored for simplicity

        v := govalidator.New()

        v.MaxString(user.Name, "first_name", "please fill first_name field") // with custom field name and custom validation error message

        if v.IsFailed() {
            return v.Errors()
         }
   ```

3. more complex with custom validation rule: You can define any custom rules or any flexible rule that does not exist in default govalidator package. Simply use check() method to define your desired data validations:

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
        // `checkScore()` is your custom method that returns a bool and can be used as a rule
        v.Check(checkScore(), "score", "score must")

        // using `In` Generic rule:
        statuses := []string{"active", "inactive", "pending"}

        v.Check(validator.In[ProfileStatuses](profile.Status, statuses), "status", "status must be in:...")

        if v.IsFailed() {
            return v.Errors()
        }
    ```

---
### Benchmarks
---
TBD