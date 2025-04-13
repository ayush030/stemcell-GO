![1_WNAfe8HdCISV0MFnK5cnag](https://github.com/user-attachments/assets/8a61a207-6c8d-4ad4-934b-e4e9f6d6b77f)
This document specifies guideines to start stemcell-GO which is a ready to deploy REST API server application. 

The intent of this repository is to provide a framework to write GO server application which only requires to replace business logic, skipping all the hassle to create other environment.

The project spawns 2 microservices:
1. PostgreSQL - database for data storage
2. Hornet - web server for CRUD operations


Pre-requisites:
1. Internet connection to download dependencies.


To start the server-
1. on Ubuntu machine, run script `run-ubuntu.sh`
2. for other OSs, run below command from the project directory where docker-compose.yml is present:
	`docker-compose up -d`  

NOTE: On host machines with OSs other then Ubuntu, docker and docker-compose utility must be already installed on it.


To interact with the server-
1. get server host IP of the machine running the microservices:
	 `ifconfig|less`
2. The server listens on port 8080. One can send request using curl or Postman.  
ex: `curl -k http://localhost:8080/api/resource/`


The server exposes 5 API endpoints:
1. POST /api/resource - Create a new resource
2. GET /api/resource/{id} - Get a specific resource by ID
3. PUT /api/resource/{id} - Update an existing resource by ID
4. DELETE /api/resource/{id} - Delete a resource by ID
5. GET /api/resource/ - List all resources


The payload request of API endpoint is a string which will be stored into DB alongwith an ID field that will be the primary key.  
ex:  `{"payload": "payload_string"}`

The ansible-playbook directory contains playbooks to deploy a running stemcell-GO server on a EC2 instance. Refer to its README.txt for more info.

The exposed endpoints can be tested via importing postman collection provided in test/postman-collections directory.

Go Linters:

The project uses `golangci-lint` as linter library.
1. To install run  
`curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.0.2`  
											OR    
`go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.0.2`
2. To check the version of linter run `golangci-lint --version` on your console. The utility should be present in $GOPATH/bin directory
3. To run the linters, run `make lint` on your console.