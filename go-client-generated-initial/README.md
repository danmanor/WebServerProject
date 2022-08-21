# Go API client for swagger

A simple web server for managing users

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://127.0.0.1:3000*

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