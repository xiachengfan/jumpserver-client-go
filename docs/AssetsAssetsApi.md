# \AssetsAssetsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAssetsAliveRead**](AssetsAssetsApi.md#AssetsAssetsAliveRead) | **Get** /assets/assets/{id}/alive/ | assets_assets_alive_read
[**AssetsAssetsBulkDelete**](AssetsAssetsApi.md#AssetsAssetsBulkDelete) | **Delete** /assets/assets/ | assets_assets_bulk_delete
[**AssetsAssetsBulkUpdate**](AssetsAssetsApi.md#AssetsAssetsBulkUpdate) | **Put** /assets/assets/ | assets_assets_bulk_update
[**AssetsAssetsCreate**](AssetsAssetsApi.md#AssetsAssetsCreate) | **Post** /assets/assets/ | assets_assets_create
[**AssetsAssetsDelete**](AssetsAssetsApi.md#AssetsAssetsDelete) | **Delete** /assets/assets/{id}/ | assets_assets_delete
[**AssetsAssetsGatewayRead**](AssetsAssetsApi.md#AssetsAssetsGatewayRead) | **Get** /assets/assets/{id}/gateway/ | assets_assets_gateway_read
[**AssetsAssetsList**](AssetsAssetsApi.md#AssetsAssetsList) | **Get** /assets/assets/ | assets_assets_list
[**AssetsAssetsPartialBulkUpdate**](AssetsAssetsApi.md#AssetsAssetsPartialBulkUpdate) | **Patch** /assets/assets/ | assets_assets_partial_bulk_update
[**AssetsAssetsPartialUpdate**](AssetsAssetsApi.md#AssetsAssetsPartialUpdate) | **Patch** /assets/assets/{id}/ | assets_assets_partial_update
[**AssetsAssetsPlatformRead**](AssetsAssetsApi.md#AssetsAssetsPlatformRead) | **Get** /assets/assets/{id}/platform/ | assets_assets_platform_read
[**AssetsAssetsRead**](AssetsAssetsApi.md#AssetsAssetsRead) | **Get** /assets/assets/{id}/ | assets_assets_read
[**AssetsAssetsRefreshRead**](AssetsAssetsApi.md#AssetsAssetsRefreshRead) | **Get** /assets/assets/{id}/refresh/ | assets_assets_refresh_read
[**AssetsAssetsUpdate**](AssetsAssetsApi.md#AssetsAssetsUpdate) | **Put** /assets/assets/{id}/ | assets_assets_update


# **AssetsAssetsAliveRead**
> TaskId AssetsAssetsAliveRead(ctx, id)
assets_assets_alive_read

Test asset admin user assets_connectivity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**TaskId**](TaskID.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetsBulkDelete**
> AssetsAssetsBulkDelete(ctx, optional)
assets_assets_bulk_delete

API endpoint that allows Asset to be viewed or edited.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsAssetsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsAssetsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostname** | **optional.String**|  | 
 **ip** | **optional.String**|  | 
 **systemuserId** | **optional.String**|  | 
 **adminUserId** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **node** | **optional.String**|  | 
 **all** | **optional.String**|  | 
 **label** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetsBulkUpdate**
> Asset AssetsAssetsBulkUpdate(ctx, data)
assets_assets_bulk_update

API endpoint that allows Asset to be viewed or edited.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Asset**](Asset.md)|  | 

### Return type

[**Asset**](Asset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetsCreate**
> Asset AssetsAssetsCreate(ctx, data)
assets_assets_create

API endpoint that allows Asset to be viewed or edited.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Asset**](Asset.md)|  | 

### Return type

[**Asset**](Asset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetsDelete**
> AssetsAssetsDelete(ctx, id)
assets_assets_delete

API endpoint that allows Asset to be viewed or edited.

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

# **AssetsAssetsGatewayRead**
> GatewayWithAuth AssetsAssetsGatewayRead(ctx, id)
assets_assets_gateway_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**GatewayWithAuth**](GatewayWithAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetsList**
> InlineResponse2006 AssetsAssetsList(ctx, optional)
assets_assets_list

API endpoint that allows Asset to be viewed or edited.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsAssetsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsAssetsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostname** | **optional.String**|  | 
 **ip** | **optional.String**|  | 
 **systemuserId** | **optional.String**|  | 
 **adminUserId** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **node** | **optional.String**|  | 
 **all** | **optional.String**|  | 
 **label** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetsPartialBulkUpdate**
> Asset AssetsAssetsPartialBulkUpdate(ctx, data)
assets_assets_partial_bulk_update

API endpoint that allows Asset to be viewed or edited.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Asset**](Asset.md)|  | 

### Return type

[**Asset**](Asset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetsPartialUpdate**
> Asset AssetsAssetsPartialUpdate(ctx, id, data)
assets_assets_partial_update

API endpoint that allows Asset to be viewed or edited.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**Asset**](Asset.md)|  | 

### Return type

[**Asset**](Asset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetsPlatformRead**
> Platform AssetsAssetsPlatformRead(ctx, id)
assets_assets_platform_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this 系统平台. | 

### Return type

[**Platform**](Platform.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetsRead**
> Asset AssetsAssetsRead(ctx, id)
assets_assets_read

API endpoint that allows Asset to be viewed or edited.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Asset**](Asset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetsRefreshRead**
> Asset AssetsAssetsRefreshRead(ctx, id)
assets_assets_refresh_read

Refresh asset hardware info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Asset**](Asset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetsUpdate**
> Asset AssetsAssetsUpdate(ctx, id, data)
assets_assets_update

API endpoint that allows Asset to be viewed or edited.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**Asset**](Asset.md)|  | 

### Return type

[**Asset**](Asset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

