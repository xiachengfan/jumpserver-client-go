# \TerminalCommandStoragesApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TerminalCommandStoragesCreate**](TerminalCommandStoragesApi.md#TerminalCommandStoragesCreate) | **Post** /terminal/command-storages/ | terminal_command-storages_create
[**TerminalCommandStoragesDelete**](TerminalCommandStoragesApi.md#TerminalCommandStoragesDelete) | **Delete** /terminal/command-storages/{id}/ | terminal_command-storages_delete
[**TerminalCommandStoragesList**](TerminalCommandStoragesApi.md#TerminalCommandStoragesList) | **Get** /terminal/command-storages/ | terminal_command-storages_list
[**TerminalCommandStoragesPartialUpdate**](TerminalCommandStoragesApi.md#TerminalCommandStoragesPartialUpdate) | **Patch** /terminal/command-storages/{id}/ | terminal_command-storages_partial_update
[**TerminalCommandStoragesRead**](TerminalCommandStoragesApi.md#TerminalCommandStoragesRead) | **Get** /terminal/command-storages/{id}/ | terminal_command-storages_read
[**TerminalCommandStoragesTestConnectiveRead**](TerminalCommandStoragesApi.md#TerminalCommandStoragesTestConnectiveRead) | **Get** /terminal/command-storages/{id}/test-connective/ | terminal_command-storages_test-connective_read
[**TerminalCommandStoragesUpdate**](TerminalCommandStoragesApi.md#TerminalCommandStoragesUpdate) | **Put** /terminal/command-storages/{id}/ | terminal_command-storages_update


# **TerminalCommandStoragesCreate**
> CommandStorage TerminalCommandStoragesCreate(ctx, data)
terminal_command-storages_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CommandStorage**](CommandStorage.md)|  | 

### Return type

[**CommandStorage**](CommandStorage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalCommandStoragesDelete**
> TerminalCommandStoragesDelete(ctx, id)
terminal_command-storages_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this command storage. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalCommandStoragesList**
> InlineResponse20053 TerminalCommandStoragesList(ctx, optional)
terminal_command-storages_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TerminalCommandStoragesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerminalCommandStoragesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20053**](inline_response_200_53.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalCommandStoragesPartialUpdate**
> CommandStorage TerminalCommandStoragesPartialUpdate(ctx, id, data)
terminal_command-storages_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this command storage. | 
  **data** | [**CommandStorage**](CommandStorage.md)|  | 

### Return type

[**CommandStorage**](CommandStorage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalCommandStoragesRead**
> CommandStorage TerminalCommandStoragesRead(ctx, id)
terminal_command-storages_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this command storage. | 

### Return type

[**CommandStorage**](CommandStorage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalCommandStoragesTestConnectiveRead**
> TerminalCommandStoragesTestConnectiveRead(ctx, id)
terminal_command-storages_test-connective_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this command storage. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalCommandStoragesUpdate**
> CommandStorage TerminalCommandStoragesUpdate(ctx, id, data)
terminal_command-storages_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this command storage. | 
  **data** | [**CommandStorage**](CommandStorage.md)|  | 

### Return type

[**CommandStorage**](CommandStorage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

