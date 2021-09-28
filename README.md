# Discounts Applier
This is a REST API endpoint that applies some discounts to a given list of products.

## Get Started 
### Requirements
To run the application it is necessary to have the port:8080 available and count with docker-compose installed. Also, to run the unitary tests it is necessary to count with GO.

### Running tests
The project consists of unitary tests as well as integration tests. The results of these can be seen in the actions tab on Github.

#### Unitary tests
To run the unitary tests it is necessary to have GO installed and from the project root run the command:
```shell
go test ./...
```

#### Integration Tests
Docker-compose is used for the integration tests so one requirement is to have this tool installed. 
In the project root we find the file `docker-compose-test.yml`.
To run it we can execute:
```shell
docker-compose -f docker-compose-test.yml up --build
```
The data that is contained in the DB at the moment of running the tests is found in the file `data.json`.

#### Single Command
To make the task of running the tests easier, the project has the script `test.sh`. This same command will run the unitary tests, then the integration tests and will stop the running containers.

### Running the api
In order to serve the API in our computer we also use docker compose by running this command:
```shell
docker-compose up --build
```
For more convenience the command `run.sh` was included.

## Functioning
The API consists of a single endpoint:
```http request
GET /products
```
The one returns a list of products(max 5) applying certain discount rules.

### Filters
The API counts with a filter by category and by "price less than...". These are combinable with each other.

### Filter by category
It is specified by the query param `category` as follows:
```http request
GET /products?category={category}
```

### Filter by price less than...
In order to filter by prices the following query param is sent `priceLessThan`:
```http request
GET /products?priceLessThan={price}
```

## Design
The solution was implemented by GO following package driven design.

### Code structure
The code of the application itself is found in the homonym subdirectory to the project `discounts-applier`. Inside the cmd folder, there are 2 applications divided, `api` and `writer`. The second one only used to save the data in the DB. These two packages hold the logic related to the application interface. 
The business logic is inside `internal`.


### Mocks & Stubs
#### Mocks
To develop the unit tests, we always tried to mock up the next layer. For this purpose, the mck subpackage was used within testify package, also used thanks to its variety of "assert" functions used for the tests.

#### Stubs 
In some cases, mocking a dependency was not easy since it was called through some package function, to solve that a series of "stubs" were implemented.
These consist in instead of calling the function directly, I have a variable with said function as a value. The stubs allow me to temporarily change the value of this variable for another that complies with the same interface.

### Testing DB Client ??
Undoubtedly the struct that brought the most problems to be tested was the repository, since the Mongo client does not present an easy way to be tested. Due to this functionality, Interfaces and structs were generated that wrap the types provided by Mongo in order to mock them up.

### Subcommands
The same executable `` 'api'`` is the one that serves the application, or runs the integration tests.
