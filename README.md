Repo for practicing creating a REST API in Golang

"Let’s go with a basic spec. Rest api for creating, managing and updating user"

Command Line Examples:

GET
"curl http://localhost:8080/persons --header "Content-Type: application/json" --request "GET""

GET BY FIRSTNAME
"curl http://localhost:8080/persons/:firstname --header "Content-Type: application/json" --request "GET""

POST
"curl http://localhost:8080/persons --include --header "Content-Type: application/json" --request "POST" --data "{\"firstname\": \"Steve\",\"lastname\": \"Smith\",\"age\": 35}""

PATCH (Works by id)
"curl http://localhost:8080/persons/:id --include --header "Content-Type: application/json" --request "PATCH" --data "{\"firstname\": \"Steve\",\"lastname\": \"Jones\",\"age\": 35}"  "
