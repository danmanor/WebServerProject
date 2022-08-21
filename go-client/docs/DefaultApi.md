# {{classname}}

All URIs are relative to *http://127.0.0.1:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUser**](DefaultApi.md#AddUser) | **Post** /api/users | Create a new user in the system
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /api/users/{id} | Deletes a user from the system
[**FindUsers**](DefaultApi.md#FindUsers) | **Get** /api/users | Retrieves all the users from the system

# **AddUser**
> User AddUser(ctx, body)
Create a new user in the system

Create a new user in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NewUser**](NewUser.md)| Create a new user in the system | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> User DeleteUser(ctx, id)
Deletes a user from the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| user id to delete | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUsers**
> []User FindUsers(ctx, )
Retrieves all the users from the system

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

