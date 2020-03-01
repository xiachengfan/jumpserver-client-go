# \AssetsAssetUsersApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAssetUsersAuthInfoRead**](AssetsAssetUsersApi.md#AssetsAssetUsersAuthInfoRead) | **Get** /assets/asset-users/auth-info/ | assets_asset-users_auth-info_read
[**AssetsAssetUsersCreate**](AssetsAssetUsersApi.md#AssetsAssetUsersCreate) | **Post** /assets/asset-users/ | assets_asset-users_create
[**AssetsAssetUsersList**](AssetsAssetUsersApi.md#AssetsAssetUsersList) | **Get** /assets/asset-users/ | assets_asset-users_list
[**AssetsAssetUsersRead**](AssetsAssetUsersApi.md#AssetsAssetUsersRead) | **Get** /assets/asset-users/{id}/ | assets_asset-users_read
[**AssetsAssetUsersTestConnectiveRead**](AssetsAssetUsersApi.md#AssetsAssetUsersTestConnectiveRead) | **Get** /assets/asset-users/test-connective/ | assets_asset-users_test-connective_read


# **AssetsAssetUsersAuthInfoRead**
> AssetUserAuthInfo AssetsAssetUsersAuthInfoRead(ctx, )
assets_asset-users_auth-info_read



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AssetUserAuthInfo**](AssetUserAuthInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetUsersCreate**
> AssetUser AssetsAssetUsersCreate(ctx, data)
assets_asset-users_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetUser**](AssetUser.md)|  | 

### Return type

[**AssetUser**](AssetUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetUsersList**
> InlineResponse2005 AssetsAssetUsersList(ctx, optional)
assets_asset-users_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsAssetUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsAssetUsersListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetUsersRead**
> AssetUser AssetsAssetUsersRead(ctx, id)
assets_asset-users_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AssetUser**](AssetUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAssetUsersTestConnectiveRead**
> TaskId AssetsAssetUsersTestConnectiveRead(ctx, )
assets_asset-users_test-connective_read

Test asset users connective

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TaskId**](TaskID.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

