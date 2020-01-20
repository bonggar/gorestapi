# GO REST API

GO REST API written in Go (Golang) powered by popular Gin Web Framework and Gorm ORM library.


# Dependencies

- Gin: github.com/gin-gonic/contrib,
- Gorm: github.com/gin-gonic/gin.
- Godotenv: github.com/joho/godotenv
- Go-Sqlite3: github.com/mattn/go-sqlite3
- Govalidator v9: gopkg.in/go-playground/validator.v9


# Features
- Config Environment
- RESTFull Endpoint
- ORM Support MySQL, PostgreSQL, SQL Server, SQLite.
- Auto Migrate
- Properly Formatted Error Request Validation
- Gracefully Shutdown
- Unit Test
- ReactJS Frontend CRUD Demo

# SPECS

This project created to serve REST API with Golang as backend and CRUD with ReactJS as frontend.

## How To

> go run main.go
> go build
> go install

## Config

Just copy .env.example file to .env
- Application name, default: "My App"
> APP_NAME="GO - REST- API"
- HTTP PORT, default: 8080
> APP_PORT=3000
- Gin Mode, default: debug
> GIN_MODE=release
- Debugging executed queries, default: false
> DB_DEBUG=true
- Database name, default: mydata
> DB_NAME=gorestdb

## Routes

Base URL :
> http://localhost:8080/

List of endpoint :
| Method         |URI                          |Purpose                         |
|----------------|-----------------------------|-----------------------------|
|GET             |`'/api/v1/users'`            |Get collection of user       |
|GET             |`'/api/v1/users/:id'`        |Get user by ID               |
|POST            |`'/api/v1/users'`            |Create user                  |
|PUT             |`'/api/v1/users/:id'`        |Update user by ID            |
|DELETE          |`'/api/v1/users/:id'`        |Delete user by ID            |
|OPTIONS         |`'/api/v1/users'`            |Support for CORS             |
|OPTIONS         |`'/api/v1/users/:id'`        |Support for CORS             |

## Response

* **Response Body:**

  ```javascript
    {
      success: bool,
      message: string,
      data: payload,
      error: errors,
    }
  ```
 
	 Payload : mixed of single object or collections of object or null
 	 Errors : must be collection of object or null
 	 
* **Validation Error :**

  ```javascript
    {
      id: string,
      value: string,
      caused: string,
      message: string,
    }
  ```
  
  * **Example Code 200 OK:**

  ```javascript
    {
      success: true,
      message: "User has been created successfully",
      data: {
	      id: 1,
	      created_at:  "2020-01-20T06:56:40.211826979+07:00",
	      updated_at:  "2020-01-20T06:56:40.211826979+07:00",
	      name:  "Alan Smith",
	      email:  "alan@123.com",
	      phone:  "0811122225555",
	      dob:  "2000-12-01",
	      gender:  "m",
	      address:  "New York",
      },
      error: null
    }
  ```


	* **Example Code 422 Unprocessable Entity:**

  ```javascript
    {
      success: false,
      message: "Could not create user",
      data: null,
      error: [
	      {
		      id: name,
		      value: "",
		      caused: "Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag",
		      message: "Required",
	      },
	      {
		      id: email,
		      value: "alan@123.",
		      caused: "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag" tag",
		      message: "Invalid email format",
	      },
      ]
    }
  ```



## Unit Test

Run unit test :
> go test -v ./...


## Authors

* **Bonggar Situmorang** - *Initial work* - [bonggar](https://github.com/bonggar)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

