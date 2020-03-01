# \TerminalStatusApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TerminalStatusCreate**](TerminalStatusApi.md#TerminalStatusCreate) | **Post** /terminal/status/ | terminal_status_create
[**TerminalStatusDelete**](TerminalStatusApi.md#TerminalStatusDelete) | **Delete** /terminal/status/{id}/ | terminal_status_delete
[**TerminalStatusList**](TerminalStatusApi.md#TerminalStatusList) | **Get** /terminal/status/ | terminal_status_list
[**TerminalStatusPartialUpdate**](TerminalStatusApi.md#TerminalStatusPartialUpdate) | **Patch** /terminal/status/{id}/ | terminal_status_partial_update
[**TerminalStatusRead**](TerminalStatusApi.md#TerminalStatusRead) | **Get** /terminal/status/{id}/ | terminal_status_read
[**TerminalStatusUpdate**](TerminalStatusApi.md#TerminalStatusUpdate) | **Put** /terminal/status/{id}/ | terminal_status_update


# **TerminalStatusCreate**
> Status TerminalStatusCreate(ctx, data)
terminal_status_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Status**](Status.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalStatusDelete**
> TerminalStatusDelete(ctx, id)
terminal_status_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this status. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalStatusList**
> InlineResponse20057 TerminalStatusList(ctx, optional)
terminal_status_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TerminalStatusListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerminalStatusListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20057**](inline_response_200_57.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalStatusPartialUpdate**
> Status TerminalStatusPartialUpdate(ctx, id, data)
terminal_status_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this status. | 
  **data** | [**Status**](Status.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalStatusRead**
> Status TerminalStatusRead(ctx, id)
terminal_status_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this status. | 

### Return type

[**Status**](Status.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalStatusUpdate**
> Status TerminalStatusUpdate(ctx, id, data)
terminal_status_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this status. | 
  **data** | [**Status**](Status.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

