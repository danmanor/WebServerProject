# Go App With Swagger

A simple web server for managing users

## Overview
This server was generated by the [swagger-codegen]
(https://github.com/swagger-api/swagger-codegen) project.  
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.  
-

To see how to make this your own, look here:

[README](https://github.com/swagger-api/swagger-codegen/blob/master/README.md)

- API version: 1.0
- Build date: 2022-08-21T08:24:57.730Z[GMT]


### Running the app with minikube
To run the server, follow these simple steps:

pull the app-k8s folder

```
minikube start
kubectl create secret generic postgres-password --from-literal POSTGRES_PASSWORD=<password>
kubectl apply -f app-k8s
```

pull app-client
then run inside:

```
go run main.go
```

## Documentation for API Endpoints

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**AddUser**](docs/DefaultApi.md#adduser) | **Post** /api/users | Create a new user in the system
*DefaultApi* | [**DeleteUser**](docs/DefaultApi.md#deleteuser) | **Delete** /api/users/{id} | Deletes a user from the system
*DefaultApi* | [**FindUsers**](docs/DefaultApi.md#findusers) | **Get** /api/users | Retrieves all the users from the system

## Documentation For Models

 - [ModelError](docs/ModelError.md)
 - [NewUser](docs/NewUser.md)
 - [User](docs/User.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

dmanor@redhat.com


