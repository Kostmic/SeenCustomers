# SeenCustomers
This solution consists of a frontend(React) + backend(Golang). The backend is a REST API with basic CRUD functionality and the frontend is built with create-react-app as a template

## Setup
### Golang
Go to https://golang.org/dl/ to download the latest golang release.
Once downloaded navigate into the backend subdirectory of this project and run this command:

```go run .```

### React
Download Node from https://nodejs.org/en/download/.
Once downloaded navigate into the frontend subdirectory of this project and run these commands:

```npm install```

```npm start```


Once the backend + frontend are both up and running you will see a table of customers fetched from the backend.

### Manipulating table data
To change the data in the table there is an endpoint that will handle it. Remember to vary the body content and http request type:

Add a new user:

POST(http://localhost:8080/user) - You need to pass a user JSON object (email, phone)


Update an existing user:

PUT(http://localhost:8080/user) - Pass in a JSON ojbect with a user ID

Delete an existing user:

DELETE(http://localhost:8080/user) - Pass in a JSON object with a user ID
