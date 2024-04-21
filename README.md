# RESTFUL-API
A RESTful API built with python(Flask + MongoDB | Django + Mysql) and go(Gin + Mysql |
Fiber + MongoDB) running in docker, allowing users to perform CRUD operations on a
database through HTTP requests
### prerequisites
- [Docker](https://www.docker.com/get-started): Docker is required to run the application in a container.
### EndPoints
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
    "age": "22"
}
```
4. **PUT /update/ID**
* Desc: Update the user identified by their `ID`
* input: name and/or age that you want to modify
```
{
    "name": "israel",
    "age": "22"
}
```
5. **DELETE /delete/ID**
  * Desc: Delete the user identified by their `ID`
# Go(Gin + Mysql)
* git clone https://github.com/RaelzeraXD/restful-api
* cd restful-api/gin
* docker-compose up
# Go(Fiber + MongoDB)
* git clone https://github.com/RaelzeraXD/restful-api
* cd restful-api/fiber
* docker-compose up
# Python(Django + Mysql)
* git clone https://github.com/RaelzeraXD/restful-api
* cd restful-api/flask
* docker-compose up
# Python(Flask + MongoDB)
* git clone https://github.com/RaelzeraXD/restful-api
* cd restful-api/flask
* docker-compose up
