# \AssetsGatewaysApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsGatewaysBulkDelete**](AssetsGatewaysApi.md#AssetsGatewaysBulkDelete) | **Delete** /assets/gateways/ | assets_gateways_bulk_delete
[**AssetsGatewaysBulkUpdate**](AssetsGatewaysApi.md#AssetsGatewaysBulkUpdate) | **Put** /assets/gateways/ | assets_gateways_bulk_update
[**AssetsGatewaysCreate**](AssetsGatewaysApi.md#AssetsGatewaysCreate) | **Post** /assets/gateways/ | assets_gateways_create
[**AssetsGatewaysDelete**](AssetsGatewaysApi.md#AssetsGatewaysDelete) | **Delete** /assets/gateways/{id}/ | assets_gateways_delete
[**AssetsGatewaysList**](AssetsGatewaysApi.md#AssetsGatewaysList) | **Get** /assets/gateways/ | assets_gateways_list
[**AssetsGatewaysPartialBulkUpdate**](AssetsGatewaysApi.md#AssetsGatewaysPartialBulkUpdate) | **Patch** /assets/gateways/ | assets_gateways_partial_bulk_update
[**AssetsGatewaysPartialUpdate**](AssetsGatewaysApi.md#AssetsGatewaysPartialUpdate) | **Patch** /assets/gateways/{id}/ | assets_gateways_partial_update
[**AssetsGatewaysRead**](AssetsGatewaysApi.md#AssetsGatewaysRead) | **Get** /assets/gateways/{id}/ | assets_gateways_read
[**AssetsGatewaysTestConnectiveCreate**](AssetsGatewaysApi.md#AssetsGatewaysTestConnectiveCreate) | **Post** /assets/gateways/{id}/test-connective/ | assets_gateways_test-connective_create
[**AssetsGatewaysUpdate**](AssetsGatewaysApi.md#AssetsGatewaysUpdate) | **Put** /assets/gateways/{id}/ | assets_gateways_update


# **AssetsGatewaysBulkDelete**
> AssetsGatewaysBulkDelete(ctx, optional)
assets_gateways_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsGatewaysBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsGatewaysBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **username** | **optional.String**|  | 
 **ip** | **optional.String**|  | 
 **domain** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsGatewaysBulkUpdate**
> Gateway AssetsGatewaysBulkUpdate(ctx, data)
assets_gateways_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Gateway**](Gateway.md)|  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsGatewaysCreate**
> Gateway AssetsGatewaysCreate(ctx, data)
assets_gateways_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Gateway**](Gateway.md)|  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsGatewaysDelete**
> AssetsGatewaysDelete(ctx, id)
assets_gateways_delete



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

# **AssetsGatewaysList**
> InlineResponse20011 AssetsGatewaysList(ctx, optional)
assets_gateways_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsGatewaysListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsGatewaysListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **username** | **optional.String**|  | 
 **ip** | **optional.String**|  | 
 **domain** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsGatewaysPartialBulkUpdate**
> Gateway AssetsGatewaysPartialBulkUpdate(ctx, data)
assets_gateways_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Gateway**](Gateway.md)|  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsGatewaysPartialUpdate**
> Gateway AssetsGatewaysPartialUpdate(ctx, id, data)
assets_gateways_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**Gateway**](Gateway.md)|  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsGatewaysRead**
> Gateway AssetsGatewaysRead(ctx, id)
assets_gateways_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsGatewaysTestConnectiveCreate**
> AssetsGatewaysTestConnectiveCreate(ctx, id)
assets_gateways_test-connective_create



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

# **AssetsGatewaysUpdate**
> Gateway AssetsGatewaysUpdate(ctx, id, data)
assets_gateways_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**Gateway**](Gateway.md)|  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

