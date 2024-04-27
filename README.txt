The project spawns 2 microservices:
1. PostgreSQL - database for data storage
2. Hornet - web server for CRUD operations


Pre-requisites:
1. Docker and docker-compose utility must be installed on the setup.
2. Internet connection to download docker images


To start the project run below command from the directory docker-compose.yml is present:
	docker-compose up -d 


To interact with the server-
1. get host IP of the machine running the microservices:
	 ifconfig|less
2. The server listens on port 8080. One can send request using curl or Postman.
	ex: curl -k http://<ip>:8080/api/resource/


The server exposes 5 API endpoints:
1. POST /api/resource - Create a new resource
2. GET /api/resource/{id} - Get a specific resource by ID
3. PUT /api/resource/{id} - Update an existing resource by ID
4. DELETE /api/resource/{id} - Delete a resource by ID
5. GET /api/resource/ - List all resources


The payload request of API endpoint is a string which will be stored into DB alongwith an ID field that will be the primary key:
ex:
{
	"payload": <string>
}

