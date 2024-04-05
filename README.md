# RESTFUL-API
A RESTful API built with python(flask + mysql) and go(gin + mysql) running in docker,
allowing users to perform CRUD operations on a database through HTTP requests
### prerequisites
- [Docker](https://www.docker.com/get-started): Docker is required to run the application in a container.
# Go
* git clone https://github.com/RaelzeraXD/restful-api
* cd restful-api/gin
* docker-compose up
### EndPoints
1. **GET /getallusers**
2. **GET /getuserbyid/ID**
* Desc: Get the user identified by their `ID`, which must be replaced with an integer.
3. **POST /createuser**
* Desc: Create a new user
* Body request: "name" "age" "email"
4. **PATCH /updateuser/ID**
* Desc: Update the user identified by their `ID`, which must be replaced with an integer.
* Body request: "name" "age" "email" any param that you want to modify
5. **DELETE /deleteuser/ID**
  * Desc: Delete the user identified by their `ID`, which must be replaced with an integer.

# Python
* git clone https://github.com/RaelzeraXD/restful-api
* cd restful-api/flask
* docker-compose up
* open your browser on localhost:5000
