# \ApplicationApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServiceApp**](ApplicationApi.md#AddServiceApp) | **Post** /application/{appId}/service | Add a service
[**CreateApp**](ApplicationApi.md#CreateApp) | **Post** /application | Create an application
[**DeleteApp**](ApplicationApi.md#DeleteApp) | **Delete** /application/{appId} | Delete an application
[**DeleteAppService**](ApplicationApi.md#DeleteAppService) | **Delete** /application/{appId}/service/{serviceId} | Delete a service of an application
[**GetApp**](ApplicationApi.md#GetApp) | **Get** /application/{appId} | Get an application.
[**GetAppServiceStatus**](ApplicationApi.md#GetAppServiceStatus) | **Get** /application/{appId}/service/{serviceId} | Get the deployment status of a service of an application
[**GetAppServices**](ApplicationApi.md#GetAppServices) | **Get** /application/{appId}/service | Get the services of an Application
[**ListApp**](ApplicationApi.md#ListApp) | **Get** /application | Retrieve the deployed applications
[**UpdateApp**](ApplicationApi.md#UpdateApp) | **Put** /application/{appId} | Update an application


# **AddServiceApp**
> AddServiceApp($appId, $service)

Add a service




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string**| ID of the app to update | 
 **service** | [**ServiceName**](ServiceName.md)| The service to add. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateApp**
> CreateApp($application)

Create an application

Should include 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | [**ApplicationCreation**](ApplicationCreation.md)| The application to create. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApp**
> DeleteApp($appId)

Delete an application




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string**| ID of the app to update | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAppService**
> DeleteAppService($appId, $serviceId)

Delete a service of an application




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string**| ID of the app to update | 
 **serviceId** | **string**| ID of the app to update | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApp**
> ApplicationStatus GetApp($appId)

Get an application.




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string**| ID of the app to update | 

### Return type

[**ApplicationStatus**](ApplicationStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppServiceStatus**
> ServiceStatus GetAppServiceStatus($appId, $serviceId)

Get the deployment status of a service of an application




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string**| ID of the app to update | 
 **serviceId** | **string**| ID of the app to update | 

### Return type

[**ServiceStatus**](ServiceStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppServices**
> []ServiceStatus GetAppServices($appId)

Get the services of an Application




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string**| ID of the app to update | 

### Return type

[**[]ServiceStatus**](ServiceStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApp**
> []ApplicationStatus ListApp()

Retrieve the deployed applications

Should include 


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]ApplicationStatus**](ApplicationStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApp**
> UpdateApp($appId, $application)

Update an application




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string**| ID of the app to update | 
 **application** | [**ApplicationCreation**](ApplicationCreation.md)| The application to create. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

