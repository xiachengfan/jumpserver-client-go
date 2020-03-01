# \OpsHistoryApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpsHistoryCreate**](OpsHistoryApi.md#OpsHistoryCreate) | **Post** /ops/history/ | ops_history_create
[**OpsHistoryDelete**](OpsHistoryApi.md#OpsHistoryDelete) | **Delete** /ops/history/{id}/ | ops_history_delete
[**OpsHistoryList**](OpsHistoryApi.md#OpsHistoryList) | **Get** /ops/history/ | ops_history_list
[**OpsHistoryPartialUpdate**](OpsHistoryApi.md#OpsHistoryPartialUpdate) | **Patch** /ops/history/{id}/ | ops_history_partial_update
[**OpsHistoryRead**](OpsHistoryApi.md#OpsHistoryRead) | **Get** /ops/history/{id}/ | ops_history_read
[**OpsHistoryUpdate**](OpsHistoryApi.md#OpsHistoryUpdate) | **Put** /ops/history/{id}/ | ops_history_update


# **OpsHistoryCreate**
> AdHocRunHistory OpsHistoryCreate(ctx, data)
ops_history_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdHocRunHistory**](AdHocRunHistory.md)|  | 

### Return type

[**AdHocRunHistory**](AdHocRunHistory.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsHistoryDelete**
> OpsHistoryDelete(ctx, id)
ops_history_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this ad hoc run history. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsHistoryList**
> InlineResponse20025 OpsHistoryList(ctx, optional)
ops_history_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpsHistoryListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpsHistoryListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsHistoryPartialUpdate**
> AdHocRunHistory OpsHistoryPartialUpdate(ctx, id, data)
ops_history_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this ad hoc run history. | 
  **data** | [**AdHocRunHistory**](AdHocRunHistory.md)|  | 

### Return type

[**AdHocRunHistory**](AdHocRunHistory.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsHistoryRead**
> AdHocRunHistory OpsHistoryRead(ctx, id)
ops_history_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this ad hoc run history. | 

### Return type

[**AdHocRunHistory**](AdHocRunHistory.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsHistoryUpdate**
> AdHocRunHistory OpsHistoryUpdate(ctx, id, data)
ops_history_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this ad hoc run history. | 
  **data** | [**AdHocRunHistory**](AdHocRunHistory.md)|  | 

### Return type

[**AdHocRunHistory**](AdHocRunHistory.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

