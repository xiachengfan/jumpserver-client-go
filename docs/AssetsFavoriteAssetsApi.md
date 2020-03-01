# \AssetsFavoriteAssetsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsFavoriteAssetsBulkDelete**](AssetsFavoriteAssetsApi.md#AssetsFavoriteAssetsBulkDelete) | **Delete** /assets/favorite-assets/ | assets_favorite-assets_bulk_delete
[**AssetsFavoriteAssetsBulkUpdate**](AssetsFavoriteAssetsApi.md#AssetsFavoriteAssetsBulkUpdate) | **Put** /assets/favorite-assets/ | assets_favorite-assets_bulk_update
[**AssetsFavoriteAssetsCreate**](AssetsFavoriteAssetsApi.md#AssetsFavoriteAssetsCreate) | **Post** /assets/favorite-assets/ | assets_favorite-assets_create
[**AssetsFavoriteAssetsDelete**](AssetsFavoriteAssetsApi.md#AssetsFavoriteAssetsDelete) | **Delete** /assets/favorite-assets/{id}/ | assets_favorite-assets_delete
[**AssetsFavoriteAssetsList**](AssetsFavoriteAssetsApi.md#AssetsFavoriteAssetsList) | **Get** /assets/favorite-assets/ | assets_favorite-assets_list
[**AssetsFavoriteAssetsPartialBulkUpdate**](AssetsFavoriteAssetsApi.md#AssetsFavoriteAssetsPartialBulkUpdate) | **Patch** /assets/favorite-assets/ | assets_favorite-assets_partial_bulk_update
[**AssetsFavoriteAssetsPartialUpdate**](AssetsFavoriteAssetsApi.md#AssetsFavoriteAssetsPartialUpdate) | **Patch** /assets/favorite-assets/{id}/ | assets_favorite-assets_partial_update
[**AssetsFavoriteAssetsRead**](AssetsFavoriteAssetsApi.md#AssetsFavoriteAssetsRead) | **Get** /assets/favorite-assets/{id}/ | assets_favorite-assets_read
[**AssetsFavoriteAssetsUpdate**](AssetsFavoriteAssetsApi.md#AssetsFavoriteAssetsUpdate) | **Put** /assets/favorite-assets/{id}/ | assets_favorite-assets_update


# **AssetsFavoriteAssetsBulkDelete**
> AssetsFavoriteAssetsBulkDelete(ctx, optional)
assets_favorite-assets_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsFavoriteAssetsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsFavoriteAssetsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsFavoriteAssetsBulkUpdate**
> FavoriteAsset AssetsFavoriteAssetsBulkUpdate(ctx, data)
assets_favorite-assets_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**FavoriteAsset**](FavoriteAsset.md)|  | 

### Return type

[**FavoriteAsset**](FavoriteAsset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsFavoriteAssetsCreate**
> FavoriteAsset AssetsFavoriteAssetsCreate(ctx, data)
assets_favorite-assets_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**FavoriteAsset**](FavoriteAsset.md)|  | 

### Return type

[**FavoriteAsset**](FavoriteAsset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsFavoriteAssetsDelete**
> AssetsFavoriteAssetsDelete(ctx, id)
assets_favorite-assets_delete



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

# **AssetsFavoriteAssetsList**
> InlineResponse20010 AssetsFavoriteAssetsList(ctx, optional)
assets_favorite-assets_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsFavoriteAssetsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsFavoriteAssetsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsFavoriteAssetsPartialBulkUpdate**
> FavoriteAsset AssetsFavoriteAssetsPartialBulkUpdate(ctx, data)
assets_favorite-assets_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**FavoriteAsset**](FavoriteAsset.md)|  | 

### Return type

[**FavoriteAsset**](FavoriteAsset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsFavoriteAssetsPartialUpdate**
> FavoriteAsset AssetsFavoriteAssetsPartialUpdate(ctx, id, data)
assets_favorite-assets_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**FavoriteAsset**](FavoriteAsset.md)|  | 

### Return type

[**FavoriteAsset**](FavoriteAsset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsFavoriteAssetsRead**
> FavoriteAsset AssetsFavoriteAssetsRead(ctx, id)
assets_favorite-assets_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**FavoriteAsset**](FavoriteAsset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsFavoriteAssetsUpdate**
> FavoriteAsset AssetsFavoriteAssetsUpdate(ctx, id, data)
assets_favorite-assets_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**FavoriteAsset**](FavoriteAsset.md)|  | 

### Return type

[**FavoriteAsset**](FavoriteAsset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

