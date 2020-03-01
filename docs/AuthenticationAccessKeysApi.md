# \AuthenticationAccessKeysApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationAccessKeysCreate**](AuthenticationAccessKeysApi.md#AuthenticationAccessKeysCreate) | **Post** /authentication/access-keys/ | authentication_access-keys_create
[**AuthenticationAccessKeysDelete**](AuthenticationAccessKeysApi.md#AuthenticationAccessKeysDelete) | **Delete** /authentication/access-keys/{id}/ | authentication_access-keys_delete
[**AuthenticationAccessKeysList**](AuthenticationAccessKeysApi.md#AuthenticationAccessKeysList) | **Get** /authentication/access-keys/ | authentication_access-keys_list
[**AuthenticationAccessKeysPartialUpdate**](AuthenticationAccessKeysApi.md#AuthenticationAccessKeysPartialUpdate) | **Patch** /authentication/access-keys/{id}/ | authentication_access-keys_partial_update
[**AuthenticationAccessKeysRead**](AuthenticationAccessKeysApi.md#AuthenticationAccessKeysRead) | **Get** /authentication/access-keys/{id}/ | authentication_access-keys_read
[**AuthenticationAccessKeysUpdate**](AuthenticationAccessKeysApi.md#AuthenticationAccessKeysUpdate) | **Put** /authentication/access-keys/{id}/ | authentication_access-keys_update


# **AuthenticationAccessKeysCreate**
> AccessKey AuthenticationAccessKeysCreate(ctx, data)
authentication_access-keys_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AccessKey**](AccessKey.md)|  | 

### Return type

[**AccessKey**](AccessKey.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationAccessKeysDelete**
> AuthenticationAccessKeysDelete(ctx, id)
authentication_access-keys_delete



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

# **AuthenticationAccessKeysList**
> InlineResponse20021 AuthenticationAccessKeysList(ctx, optional)
authentication_access-keys_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationAccessKeysListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationAccessKeysListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationAccessKeysPartialUpdate**
> AccessKey AuthenticationAccessKeysPartialUpdate(ctx, id, data)
authentication_access-keys_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AccessKey**](AccessKey.md)|  | 

### Return type

[**AccessKey**](AccessKey.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationAccessKeysRead**
> AccessKey AuthenticationAccessKeysRead(ctx, id)
authentication_access-keys_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AccessKey**](AccessKey.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationAccessKeysUpdate**
> AccessKey AuthenticationAccessKeysUpdate(ctx, id, data)
authentication_access-keys_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AccessKey**](AccessKey.md)|  | 

### Return type

[**AccessKey**](AccessKey.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

