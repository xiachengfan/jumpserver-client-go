# \OpsCommandExecutionsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpsCommandExecutionsCreate**](OpsCommandExecutionsApi.md#OpsCommandExecutionsCreate) | **Post** /ops/command-executions/ | ops_command-executions_create
[**OpsCommandExecutionsDelete**](OpsCommandExecutionsApi.md#OpsCommandExecutionsDelete) | **Delete** /ops/command-executions/{id}/ | ops_command-executions_delete
[**OpsCommandExecutionsList**](OpsCommandExecutionsApi.md#OpsCommandExecutionsList) | **Get** /ops/command-executions/ | ops_command-executions_list
[**OpsCommandExecutionsPartialUpdate**](OpsCommandExecutionsApi.md#OpsCommandExecutionsPartialUpdate) | **Patch** /ops/command-executions/{id}/ | ops_command-executions_partial_update
[**OpsCommandExecutionsRead**](OpsCommandExecutionsApi.md#OpsCommandExecutionsRead) | **Get** /ops/command-executions/{id}/ | ops_command-executions_read
[**OpsCommandExecutionsUpdate**](OpsCommandExecutionsApi.md#OpsCommandExecutionsUpdate) | **Put** /ops/command-executions/{id}/ | ops_command-executions_update


# **OpsCommandExecutionsCreate**
> CommandExecution OpsCommandExecutionsCreate(ctx, data)
ops_command-executions_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CommandExecution**](CommandExecution.md)|  | 

### Return type

[**CommandExecution**](CommandExecution.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsCommandExecutionsDelete**
> OpsCommandExecutionsDelete(ctx, id)
ops_command-executions_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsCommandExecutionsList**
> InlineResponse20024 OpsCommandExecutionsList(ctx, optional)
ops_command-executions_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpsCommandExecutionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpsCommandExecutionsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsCommandExecutionsPartialUpdate**
> CommandExecution OpsCommandExecutionsPartialUpdate(ctx, id, data)
ops_command-executions_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**CommandExecution**](CommandExecution.md)|  | 

### Return type

[**CommandExecution**](CommandExecution.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsCommandExecutionsRead**
> CommandExecution OpsCommandExecutionsRead(ctx, id)
ops_command-executions_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**CommandExecution**](CommandExecution.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsCommandExecutionsUpdate**
> CommandExecution OpsCommandExecutionsUpdate(ctx, id, data)
ops_command-executions_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**CommandExecution**](CommandExecution.md)|  | 

### Return type

[**CommandExecution**](CommandExecution.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

