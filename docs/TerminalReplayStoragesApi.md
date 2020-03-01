# \TerminalReplayStoragesApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TerminalReplayStoragesCreate**](TerminalReplayStoragesApi.md#TerminalReplayStoragesCreate) | **Post** /terminal/replay-storages/ | terminal_replay-storages_create
[**TerminalReplayStoragesDelete**](TerminalReplayStoragesApi.md#TerminalReplayStoragesDelete) | **Delete** /terminal/replay-storages/{id}/ | terminal_replay-storages_delete
[**TerminalReplayStoragesList**](TerminalReplayStoragesApi.md#TerminalReplayStoragesList) | **Get** /terminal/replay-storages/ | terminal_replay-storages_list
[**TerminalReplayStoragesPartialUpdate**](TerminalReplayStoragesApi.md#TerminalReplayStoragesPartialUpdate) | **Patch** /terminal/replay-storages/{id}/ | terminal_replay-storages_partial_update
[**TerminalReplayStoragesRead**](TerminalReplayStoragesApi.md#TerminalReplayStoragesRead) | **Get** /terminal/replay-storages/{id}/ | terminal_replay-storages_read
[**TerminalReplayStoragesTestConnectiveRead**](TerminalReplayStoragesApi.md#TerminalReplayStoragesTestConnectiveRead) | **Get** /terminal/replay-storages/{id}/test-connective/ | terminal_replay-storages_test-connective_read
[**TerminalReplayStoragesUpdate**](TerminalReplayStoragesApi.md#TerminalReplayStoragesUpdate) | **Put** /terminal/replay-storages/{id}/ | terminal_replay-storages_update


# **TerminalReplayStoragesCreate**
> ReplayStorage TerminalReplayStoragesCreate(ctx, data)
terminal_replay-storages_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ReplayStorage**](ReplayStorage.md)|  | 

### Return type

[**ReplayStorage**](ReplayStorage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalReplayStoragesDelete**
> TerminalReplayStoragesDelete(ctx, id)
terminal_replay-storages_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this replay storage. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalReplayStoragesList**
> InlineResponse20055 TerminalReplayStoragesList(ctx, optional)
terminal_replay-storages_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TerminalReplayStoragesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerminalReplayStoragesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20055**](inline_response_200_55.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalReplayStoragesPartialUpdate**
> ReplayStorage TerminalReplayStoragesPartialUpdate(ctx, id, data)
terminal_replay-storages_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this replay storage. | 
  **data** | [**ReplayStorage**](ReplayStorage.md)|  | 

### Return type

[**ReplayStorage**](ReplayStorage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalReplayStoragesRead**
> ReplayStorage TerminalReplayStoragesRead(ctx, id)
terminal_replay-storages_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this replay storage. | 

### Return type

[**ReplayStorage**](ReplayStorage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalReplayStoragesTestConnectiveRead**
> TerminalReplayStoragesTestConnectiveRead(ctx, id)
terminal_replay-storages_test-connective_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this replay storage. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalReplayStoragesUpdate**
> ReplayStorage TerminalReplayStoragesUpdate(ctx, id, data)
terminal_replay-storages_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this replay storage. | 
  **data** | [**ReplayStorage**](ReplayStorage.md)|  | 

### Return type

[**ReplayStorage**](ReplayStorage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

