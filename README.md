# Project go

Web service for encoding passwords over HTTP.

## Getting Started

To build and run project:
```
	go build
	./<executable>
```

### Prerequisites

Go Tools


### Issuing Request over HTTP

Hash and Encode Passwords
Encodes passwords with SHA512 and returns a base64 encoded string of the SHA512 hash

Example:

```
curl --data password=angryMonkey http://localhost:8080/hash
```

Return Value

```
{"password":"ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="}
```

Graceful Shutdown
The server rejects new requests and waits for any pending/in-flight work to finish before exiting.

Example:

```
curl --data {} http://localhost:8080/shutdown
```

Return Value

```
{"message":"shutting down server...."}
```

Statistics
Provides basic information about the password hashes seen so far such as total number of hashes and the average time to encode those hashes in microseconds.

Example:
```
curl http://localhost:8080/stats
```

Return Value

```
{"Total":4,"Average":9.5}
```

## Running the tests
Tests issue all the different requests supported.  The tests also issue requests that should return an error.

How to run tests:

```
go test
```

## Assumptions and Improvements
All services return an encoded JSON reponse. Stats for the hashes is persistent even after shutting down the server.  Stats are stored in a JSON file.  This can be improved upon by writting to a database.  A database was not added because I felt that was out of the scope of this project.