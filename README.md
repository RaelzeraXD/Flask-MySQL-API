# RESTful API
This project demonstrates a RESTful API implementation in both Python and Go with MySQL and MongoDB. The API uses JWT tokens for authentication, with two implementations utilizing RabbitMQ as a message broker. Swagger is integrated for API documentation and better visualization of endpoints.
## Prerequisites
- [Docker](https://www.docker.com/get-started): Docker is required to run the application in a container.

## Usage
Once the application is running, you can access the API endpoints as described below. Swagger documentation is available at /swagger or /docs depending on the framework used.
## EndPoints
1. **GET /users**
* Desc: Returns all users
3. **GET /users/ID**
* Desc: Returns a single user
3. **POST /create**
* Desc: Create a new user
* input:
```
{
    "name": "israel",
    "age": 22
}
```
4. **PUT /update/ID**
* Desc: Update the user identified by their `ID`
* input: name and/or age that you want to modify
5. **DELETE /delete/ID**
  * Desc: Delete the user identified by their `ID`

## Installation
## Go(Gin + Mysql)
* git clone https://github.com/RaelzeraXD/api
* cd api/gin
* docker-compose up
## Go(Fiber + MongoDB)
* git clone https://github.com/RaelzeraXD/api
* cd api/fiber
* docker-compose up
## Python(Django + Mysql)
* git clone https://github.com/RaelzeraXD/api
* cd api/django
* docker-compose up
## Python(Flask + MongoDB)
* git clone https://github.com/RaelzeraXD/api
* cd api/flask
* docker-compose up

### Contributing
Contributions are welcome! Please fork the repository and create a pull request with your changes.

### License
This project is licensed under the MIT License.
