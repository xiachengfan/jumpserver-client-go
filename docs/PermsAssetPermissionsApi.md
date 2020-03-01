# \PermsAssetPermissionsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermsAssetPermissionsAssetsAllList**](PermsAssetPermissionsApi.md#PermsAssetPermissionsAssetsAllList) | **Get** /perms/asset-permissions/{id}/assets/all/ | perms_asset-permissions_assets_all_list
[**PermsAssetPermissionsCacheRefreshRead**](PermsAssetPermissionsApi.md#PermsAssetPermissionsCacheRefreshRead) | **Get** /perms/asset-permissions/cache/refresh/ | perms_asset-permissions_cache_refresh_read
[**PermsAssetPermissionsCreate**](PermsAssetPermissionsApi.md#PermsAssetPermissionsCreate) | **Post** /perms/asset-permissions/ | perms_asset-permissions_create
[**PermsAssetPermissionsDelete**](PermsAssetPermissionsApi.md#PermsAssetPermissionsDelete) | **Delete** /perms/asset-permissions/{id}/ | perms_asset-permissions_delete
[**PermsAssetPermissionsList**](PermsAssetPermissionsApi.md#PermsAssetPermissionsList) | **Get** /perms/asset-permissions/ | perms_asset-permissions_list
[**PermsAssetPermissionsPartialUpdate**](PermsAssetPermissionsApi.md#PermsAssetPermissionsPartialUpdate) | **Patch** /perms/asset-permissions/{id}/ | perms_asset-permissions_partial_update
[**PermsAssetPermissionsRead**](PermsAssetPermissionsApi.md#PermsAssetPermissionsRead) | **Get** /perms/asset-permissions/{id}/ | perms_asset-permissions_read
[**PermsAssetPermissionsUpdate**](PermsAssetPermissionsApi.md#PermsAssetPermissionsUpdate) | **Put** /perms/asset-permissions/{id}/ | perms_asset-permissions_update
[**PermsAssetPermissionsUserActionsRead**](PermsAssetPermissionsApi.md#PermsAssetPermissionsUserActionsRead) | **Get** /perms/asset-permissions/user/actions/ | perms_asset-permissions_user_actions_read
[**PermsAssetPermissionsUserValidateList**](PermsAssetPermissionsApi.md#PermsAssetPermissionsUserValidateList) | **Get** /perms/asset-permissions/user/validate/ | perms_asset-permissions_user_validate_list
[**PermsAssetPermissionsUsersAllList**](PermsAssetPermissionsApi.md#PermsAssetPermissionsUsersAllList) | **Get** /perms/asset-permissions/{id}/users/all/ | perms_asset-permissions_users_all_list


# **PermsAssetPermissionsAssetsAllList**
> InlineResponse20036 PermsAssetPermissionsAssetsAllList(ctx, id, optional)
perms_asset-permissions_assets_all_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***PermsAssetPermissionsAssetsAllListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsAssetPermissionsAssetsAllListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsCacheRefreshRead**
> PermsAssetPermissionsCacheRefreshRead(ctx, )
perms_asset-permissions_cache_refresh_read



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsCreate**
> AssetPermissionCreateUpdate PermsAssetPermissionsCreate(ctx, data)
perms_asset-permissions_create

资产授权列表的增删改查api

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionCreateUpdate**](AssetPermissionCreateUpdate.md)|  | 

### Return type

[**AssetPermissionCreateUpdate**](AssetPermissionCreateUpdate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsDelete**
> PermsAssetPermissionsDelete(ctx, id)
perms_asset-permissions_delete

资产授权列表的增删改查api

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

# **PermsAssetPermissionsList**
> InlineResponse20035 PermsAssetPermissionsList(ctx, optional)
perms_asset-permissions_list

资产授权列表的增删改查api

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsAssetPermissionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsAssetPermissionsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsPartialUpdate**
> AssetPermissionCreateUpdate PermsAssetPermissionsPartialUpdate(ctx, id, data)
perms_asset-permissions_partial_update

资产授权列表的增删改查api

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AssetPermissionCreateUpdate**](AssetPermissionCreateUpdate.md)|  | 

### Return type

[**AssetPermissionCreateUpdate**](AssetPermissionCreateUpdate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsRead**
> AssetPermissionCreateUpdate PermsAssetPermissionsRead(ctx, id)
perms_asset-permissions_read

资产授权列表的增删改查api

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AssetPermissionCreateUpdate**](AssetPermissionCreateUpdate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUpdate**
> AssetPermissionCreateUpdate PermsAssetPermissionsUpdate(ctx, id, data)
perms_asset-permissions_update

资产授权列表的增删改查api

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AssetPermissionCreateUpdate**](AssetPermissionCreateUpdate.md)|  | 

### Return type

[**AssetPermissionCreateUpdate**](AssetPermissionCreateUpdate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUserActionsRead**
> Actions PermsAssetPermissionsUserActionsRead(ctx, )
perms_asset-permissions_user_actions_read



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Actions**](Actions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUserValidateList**
> PermsAssetPermissionsUserValidateList(ctx, )
perms_asset-permissions_user_validate_list



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUsersAllList**
> InlineResponse20037 PermsAssetPermissionsUsersAllList(ctx, id, optional)
perms_asset-permissions_users_all_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***PermsAssetPermissionsUsersAllListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsAssetPermissionsUsersAllListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

