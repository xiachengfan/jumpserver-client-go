# \PermsAssetPermissionsAssetsRelationsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermsAssetPermissionsAssetsRelationsBulkDelete**](PermsAssetPermissionsAssetsRelationsApi.md#PermsAssetPermissionsAssetsRelationsBulkDelete) | **Delete** /perms/asset-permissions-assets-relations/ | perms_asset-permissions-assets-relations_bulk_delete
[**PermsAssetPermissionsAssetsRelationsBulkUpdate**](PermsAssetPermissionsAssetsRelationsApi.md#PermsAssetPermissionsAssetsRelationsBulkUpdate) | **Put** /perms/asset-permissions-assets-relations/ | perms_asset-permissions-assets-relations_bulk_update
[**PermsAssetPermissionsAssetsRelationsCreate**](PermsAssetPermissionsAssetsRelationsApi.md#PermsAssetPermissionsAssetsRelationsCreate) | **Post** /perms/asset-permissions-assets-relations/ | perms_asset-permissions-assets-relations_create
[**PermsAssetPermissionsAssetsRelationsDelete**](PermsAssetPermissionsAssetsRelationsApi.md#PermsAssetPermissionsAssetsRelationsDelete) | **Delete** /perms/asset-permissions-assets-relations/{id}/ | perms_asset-permissions-assets-relations_delete
[**PermsAssetPermissionsAssetsRelationsList**](PermsAssetPermissionsAssetsRelationsApi.md#PermsAssetPermissionsAssetsRelationsList) | **Get** /perms/asset-permissions-assets-relations/ | perms_asset-permissions-assets-relations_list
[**PermsAssetPermissionsAssetsRelationsPartialBulkUpdate**](PermsAssetPermissionsAssetsRelationsApi.md#PermsAssetPermissionsAssetsRelationsPartialBulkUpdate) | **Patch** /perms/asset-permissions-assets-relations/ | perms_asset-permissions-assets-relations_partial_bulk_update
[**PermsAssetPermissionsAssetsRelationsPartialUpdate**](PermsAssetPermissionsAssetsRelationsApi.md#PermsAssetPermissionsAssetsRelationsPartialUpdate) | **Patch** /perms/asset-permissions-assets-relations/{id}/ | perms_asset-permissions-assets-relations_partial_update
[**PermsAssetPermissionsAssetsRelationsRead**](PermsAssetPermissionsAssetsRelationsApi.md#PermsAssetPermissionsAssetsRelationsRead) | **Get** /perms/asset-permissions-assets-relations/{id}/ | perms_asset-permissions-assets-relations_read
[**PermsAssetPermissionsAssetsRelationsUpdate**](PermsAssetPermissionsAssetsRelationsApi.md#PermsAssetPermissionsAssetsRelationsUpdate) | **Put** /perms/asset-permissions-assets-relations/{id}/ | perms_asset-permissions-assets-relations_update


# **PermsAssetPermissionsAssetsRelationsBulkDelete**
> PermsAssetPermissionsAssetsRelationsBulkDelete(ctx, optional)
perms_asset-permissions-assets-relations_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsAssetPermissionsAssetsRelationsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsAssetPermissionsAssetsRelationsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **asset** | **optional.String**|  | 
 **assetpermission** | **optional.String**|  | 
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

# **PermsAssetPermissionsAssetsRelationsBulkUpdate**
> AssetPermissionAssetRelation PermsAssetPermissionsAssetsRelationsBulkUpdate(ctx, data)
perms_asset-permissions-assets-relations_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionAssetRelation**](AssetPermissionAssetRelation.md)|  | 

### Return type

[**AssetPermissionAssetRelation**](AssetPermissionAssetRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsAssetsRelationsCreate**
> AssetPermissionAssetRelation PermsAssetPermissionsAssetsRelationsCreate(ctx, data)
perms_asset-permissions-assets-relations_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionAssetRelation**](AssetPermissionAssetRelation.md)|  | 

### Return type

[**AssetPermissionAssetRelation**](AssetPermissionAssetRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsAssetsRelationsDelete**
> PermsAssetPermissionsAssetsRelationsDelete(ctx, id)
perms_asset-permissions-assets-relations_delete



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

# **PermsAssetPermissionsAssetsRelationsList**
> InlineResponse20030 PermsAssetPermissionsAssetsRelationsList(ctx, optional)
perms_asset-permissions-assets-relations_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsAssetPermissionsAssetsRelationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsAssetPermissionsAssetsRelationsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **asset** | **optional.String**|  | 
 **assetpermission** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsAssetsRelationsPartialBulkUpdate**
> AssetPermissionAssetRelation PermsAssetPermissionsAssetsRelationsPartialBulkUpdate(ctx, data)
perms_asset-permissions-assets-relations_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionAssetRelation**](AssetPermissionAssetRelation.md)|  | 

### Return type

[**AssetPermissionAssetRelation**](AssetPermissionAssetRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsAssetsRelationsPartialUpdate**
> AssetPermissionAssetRelation PermsAssetPermissionsAssetsRelationsPartialUpdate(ctx, id, data)
perms_asset-permissions-assets-relations_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AssetPermissionAssetRelation**](AssetPermissionAssetRelation.md)|  | 

### Return type

[**AssetPermissionAssetRelation**](AssetPermissionAssetRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsAssetsRelationsRead**
> AssetPermissionAssetRelation PermsAssetPermissionsAssetsRelationsRead(ctx, id)
perms_asset-permissions-assets-relations_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AssetPermissionAssetRelation**](AssetPermissionAssetRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsAssetsRelationsUpdate**
> AssetPermissionAssetRelation PermsAssetPermissionsAssetsRelationsUpdate(ctx, id, data)
perms_asset-permissions-assets-relations_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AssetPermissionAssetRelation**](AssetPermissionAssetRelation.md)|  | 

### Return type

[**AssetPermissionAssetRelation**](AssetPermissionAssetRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

