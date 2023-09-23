**Prerequisites**

    1. Golang (prefer version go1.21.1)
    2. MySQL (mysql Ver 8.0)

**How to run**

    1. create a database
    2. create a table in the database by copying and pasting the ddl in the migrations folder and then selecting the file.up.sql
    3. fill in the data in the table: product, category, product_category 
    4. create a .env file as in the example file .env.example
    5. download all dependencies with the command "go mod vendor"
    6. run the project with the command "go run main.go" in the terminal
    7. hit the endpoint via Postman as in the Swagger documentation

**Note**

    1. Swagger documentation is in the docs folder -> swagger.yaml

**ERD**

[![erd-challenge-test-synapsis.png](https://i.postimg.cc/dQdbSfQm/erd-challenge-test-synapsis.png)](https://postimg.cc/NLsCKpH5)
