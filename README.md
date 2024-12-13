[![Go Reference](https://pkg.go.dev/badge/github.com/rezakhademix/govalidator.svg)](https://pkg.go.dev/github.com/rezakhademix/govalidator) [![Go Report Card](https://goreportcard.com/badge/github.com/rezakhademix/govalidator)](https://goreportcard.com/report/github.com/rezakhademix/govalidator) [![codecov](https://codecov.io/gh/rezakhademix/govalidator/graph/badge.svg?token=BDWNVIC670)](https://codecov.io/gh/rezakhademix/govalidator) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

# GoValidator
GoValidator is a data validation package that can sanitize and validate your data to ensure its safety and integrity as much as possible. GoValidator will not use struct tags as validation and there is no need of them.

Our goal is to avoid using any type assertion or reflection for simplicity.
We would be delighted if you were interested in helping us improve the GoValidator package. Feel free to make your pull request.

# Getting Started
Go Validator includes a set of validation rules and a handy `CustomRule()` method to define any custom **rule**.

## Installation
Run the following command to install the package:
```
go get github.com/rezakhademix/govalidator/v2
```
## Import
Import package in your code-base by using:
```
import validator "github.com/rezakhademix/govalidator/v2"
```

# Documentation
GoValidator package has the following features and method to use.
Each validation rule in GoValidator has it's own default message, e.g: `required` rule for a required `name` has the default message of: `name is required` but you can define your custom field name and message.

### Methods (also known as validation rules)
---

| Method           | Description                                                                                                         |
| ---------------- | ------------------------------------------------------------------------------------------------------------------- |
| RequiredInt      | `RequiredInt` checks if an integer value is provided or not.                                                        |
| RequiredFloat    | `RequiredFloat` checks if a float value is provided or not.                                                         |
| RequiredString   | `RequiredString` checks if a string value is empty or not.                                                          |
| RequiredSlice    | `RequiredSlice` checks if a slice has any value or not.                                                             |
| BetweenInt       | `BetweenInt` checks whether value falls within the specified range or not.                                          |
| BetweenFloat     | `BetweenFloat` checks whether value falls within the specified range or not.                                        |
| BetweenString    | `BetweenString` checks whether string length falls within the specified range or not.                               |
| Date             | `Date` checks value to be a valid, non-relative date with given layout.                                             |
| Email            | `Email` checks value to match `EmailRegex` regular expression.                                                      |
| Exists           | `Exists` checks given value exists in given table or not.                                                           |
| NotExists        | `NotExists` checks given value doesn't exist in given table, except.                                                |
| ExistsExceptSelf | `ExistsExceptSelf` checks given value exists in a specified column of a table, excluding a row identified by selfID |
| LenString        | `LenString` checks length of given string is equal to given size or not.                                            |
| LenInt           | `LenInt` checks length of given integer is equal to given size or not.                                              |
| LenSlice         | `LenSlice` checks length of given slice is equal to given size or not.                                              |
| MaxInt           | `MaxInt` checks given integer is less than or equal given max value.                                                |
| MaxFloat         | `MaxFloat` checks given float is less than or equal given max value.                                                |
| MaxString        | `MaxString` checks length of given string is less than or equal given max value.                                    |
| MinInt           | `MinInt` checks given integer is greater than or equal given min value.                                             |
| MinFloat         | `MinFloat` checks given float is greater than or equal given min value.                                             |
| MinString        | `MinString` checks length of given string is greater than or equal given min value.                                 |
| RegexMatches     | `RegexMatches` checks given string matches given regular expression pattern.                                        |
| UUID             | `UUID` checks given value is a valid universally unique identifier (UUID).                                          |
| When             | `When` will execute given closure if given condition is true.                                                       |
| URL              | `URL` will check given string is a valid URL or not.                                                                |
| Before           | `Before` will check if given time instant t is before given time instant u.                                         |
| After            | `After` will check if given time instant t is after given time instant u.                                           |
| IP4              | `IP4` will check whether given string is a valid IPV4.                                                              |
| DefaultInt       | `DefaultInt` sets a default value for any pointer to an int that is passed.                                         |
| DefaultFloat     | `DefaultFloat` sets a default value for any pointer to a float that is passed.                                      |
| DefaultString    | `DefaultString` sets a default value for any pointer to a string that is passed.                                    |
| IsJSON           | `IsJSON` will check if given string is a valid JSON.                                                                |
| NumericString    | `NumericString` will check if given value is a valid string of numbers.                                             |
| MAC              | `MAC` will check if given value is a valid MAC address.                                                             |
| Time             | `Time` will check if given value is a non-relative time with given layout.                                          |
| CustomRule       | `CustomRule` is a dynamic method to define any custom validation rule.                                              |

### Functions (other common validation rules)
---
| Method | Description                                                         |
| ------ | ------------------------------------------------------------------- |
| In     | `In` checks given value is included in the provided list of values. |
| Unique | `Unique` checks whether values in the provided slice are unique.    |


---
### Examples:
---
1. Simple usage: ([Go Playground](https://go.dev/play/p/qevHb1KqOHf))
    ```go
        type User struct {
            Name string `json:"name"`
            Age unit    `json:"age"`
        }

        var user User
       
        v := govalidator.New() // be sure to import govalidator/v2

        v.RequiredInt(user.Age, "age", "").         // age can not be null or 0
            MinInt(user.Age, 18, "age", "")         // minimum value for age is 18

        v.RequiredString(user.Name, "name", "")     // name can not be null, "" or " "
            MaxString(user.Name, 50, "name", "")    // maximum acceptable length for name field is 50

        if v.IsFailed() {
            return v.Errors()  // will return failed validation error messages
        }
   ```

2. With custom field names and messages:
    ```go
        type User struct {
            Name string `json:"name"`
        }

        var user User

        v := govalidator.New()

        v.RequiredString(user.Name, "first_name", "please fill first_name field") // with custom field name and custom validation error message

        if v.IsFailed() {
            return v.Errors()
         }
   ```

3. Advanced usage with any custom validation rule: You can define any custom rules or any flexible rule that does not exist in default govalidator package. Simply use `CustomRule()` method to define your desired data validations:

    ```go
        type Profile struct {
            Name   string
            Age    int
            Score  int
            Status []string
        }

        var profile Profile

        // after filling profile struct

        v := govalidator.New()

        v.CustomRule(profile.Name != "", "name", "name is required")  // CustomRule is a method to define a custom rule as first parameter and then pass field name and validation error message

        v.CustomRule(profile.Age > 18, "age", "age must be greater than 18")

        // we just need to pass a bool as a rule
        // `checkScore()` is your custom method that returns a bool and can be used as a rule
        v.CustomRule(checkScore(), "score", "score must")

        // using `In` Generic rule:
        statuses := []string{"active", "inactive", "pending"}

        v.CustomRule(validator.In[ProfileStatuses](profile.Status, statuses), "status", "status must be in:...")

        if v.IsFailed() {
            return v.Errors()
        }
    ```

4. Deep Dive:

    ```go
        type CategoryCreateReq struct {
	        ParentID    *int   `json:"parent_id"`
	        Name        string `json:"name"`
	        Description string `json:"description"`
	        Status      int    `json:"status"`
            Meta        string `json:"meta"`
        }


        var categoryCreateReq CategoryCreateReq

        // after filling CategoryCreateReq struct with binding or other methods

        v := govalidator.New().WithRepo(validatorRepo)  // be sure to import govalidator/v2
		    
		ok := v.
			RequiredString(req.Name, "name", msgCategoryNameRequired).
			MaxString(req.Name, nameMaxLength, "name", msgCategoryNameMaxLength).
			MinString(req.Name, nameMinLength, "name", msgCategoryNameMinLength).
			NotExists(req.Name, "categories", "name", "name", msgCategoryNameAlreadyExists). // ensure the value of req.Name does not exist in the "name" column of the "categories" table in the database
		  	MaxString(req.Description, descriptionMaxLength, "description", msgDescriptionMaxLength).
		  	MinInt(req.Status, minStatus, "status", msgMinCategoryStatusIsWrong).
		  	MaxInt(req.Status, maxStatus, "status", msgMaxCategoryStatusIsWrong).
    		IsJSON(req.Meta, "meta", msgMetaMustBeJSON).
		  	When(req.ParentID != nil, func() {  
                		v.Exists(*req.ParentID, "categories", "id", "parent_id", msgCategoryParentIDNotExist)   // checks if the value of req.ParentID exists in the "id" column of the "categories" table in the database
			}).
		  	IsPassed()

        if !ok {
          return v.Errors()
        }
    ```

---
### Benchmarks
I wrote the sample struct below as a DTO to run benchmarks. You can view full benchmark code on [golang/playground](https://go.dev/play/p/n-Wo5vOGvHB) however, as you probably already know, benchmarks cannot be executed directly on the playground `playground` and we have to use a code editor to run them.

```go
type UserCreateReq struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	PhoneNumber   string `json:"phone_number"`
	Email         string `json:"email,omitempty"`
	FatherName    string `json:"father_name"`
	CertificateID string `json:"certificate_id"`
	BirthDate     string `json:"birth_date"`
	CompanyID     int    `json:"company_id"`
	Gender        int8   `json:"gender"`
}
```

| **Library**       | **Operations/sec (ns/op)** | **Memory Allocations (B/op)** | **Allocations/op** |
| ----------------- | -------------------------- | ----------------------------- | ------------------ |
| `govalidator`     | 870.1 ns/op                | 394 B/op                      | 22 allocs/op       |
| `go-playground`   | 17750.6 ns/op              | 20158 B/op                    | 284 alloc/op       |
| `ozzo-validation` | 4597 ns/op                 | 7192 B/op                     | 126 allocs/op      |

### Notes
These benchmarks were conducted on an **Apple M3 Pro** CPU.
****