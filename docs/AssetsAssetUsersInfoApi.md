# \AssetsAssetUsersInfoApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAssetUsersInfoList**](AssetsAssetUsersInfoApi.md#AssetsAssetUsersInfoList) | **Get** /assets/asset-users-info/ | assets_asset-users-info_list
[**AssetsAssetUsersInfoRead**](AssetsAssetUsersInfoApi.md#AssetsAssetUsersInfoRead) | **Get** /assets/asset-users-info/{id}/ | assets_asset-users-info_read


# **AssetsAssetUsersInfoList**
> InlineResponse2004 AssetsAssetUsersInfoList(ctx, optional)
assets_asset-users-info_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsAssetUsersInfoListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsAssetUsersInfoListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetUsersInfoRead**
> AssetUserExport AssetsAssetUsersInfoRead(ctx, id)
assets_asset-users-info_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AssetUserExport**](AssetUserExport.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

