# \OpsAdhocApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpsAdhocCreate**](OpsAdhocApi.md#OpsAdhocCreate) | **Post** /ops/adhoc/ | ops_adhoc_create
[**OpsAdhocDelete**](OpsAdhocApi.md#OpsAdhocDelete) | **Delete** /ops/adhoc/{id}/ | ops_adhoc_delete
[**OpsAdhocList**](OpsAdhocApi.md#OpsAdhocList) | **Get** /ops/adhoc/ | ops_adhoc_list
[**OpsAdhocPartialUpdate**](OpsAdhocApi.md#OpsAdhocPartialUpdate) | **Patch** /ops/adhoc/{id}/ | ops_adhoc_partial_update
[**OpsAdhocRead**](OpsAdhocApi.md#OpsAdhocRead) | **Get** /ops/adhoc/{id}/ | ops_adhoc_read
[**OpsAdhocUpdate**](OpsAdhocApi.md#OpsAdhocUpdate) | **Put** /ops/adhoc/{id}/ | ops_adhoc_update


# **OpsAdhocCreate**
> AdHoc OpsAdhocCreate(ctx, data)
ops_adhoc_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdHoc**](AdHoc.md)|  | 

### Return type

[**AdHoc**](AdHoc.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsAdhocDelete**
> OpsAdhocDelete(ctx, id)
ops_adhoc_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this ad hoc. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsAdhocList**
> InlineResponse20022 OpsAdhocList(ctx, optional)
ops_adhoc_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpsAdhocListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpsAdhocListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsAdhocPartialUpdate**
> AdHoc OpsAdhocPartialUpdate(ctx, id, data)
ops_adhoc_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this ad hoc. | 
  **data** | [**AdHoc**](AdHoc.md)|  | 

### Return type

[**AdHoc**](AdHoc.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsAdhocRead**
> AdHoc OpsAdhocRead(ctx, id)
ops_adhoc_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this ad hoc. | 

### Return type

[**AdHoc**](AdHoc.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsAdhocUpdate**
> AdHoc OpsAdhocUpdate(ctx, id, data)
ops_adhoc_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this ad hoc. | 
  **data** | [**AdHoc**](AdHoc.md)|  | 

### Return type

[**AdHoc**](AdHoc.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

