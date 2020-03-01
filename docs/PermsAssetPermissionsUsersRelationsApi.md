# \PermsAssetPermissionsUsersRelationsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermsAssetPermissionsUsersRelationsBulkDelete**](PermsAssetPermissionsUsersRelationsApi.md#PermsAssetPermissionsUsersRelationsBulkDelete) | **Delete** /perms/asset-permissions-users-relations/ | perms_asset-permissions-users-relations_bulk_delete
[**PermsAssetPermissionsUsersRelationsBulkUpdate**](PermsAssetPermissionsUsersRelationsApi.md#PermsAssetPermissionsUsersRelationsBulkUpdate) | **Put** /perms/asset-permissions-users-relations/ | perms_asset-permissions-users-relations_bulk_update
[**PermsAssetPermissionsUsersRelationsCreate**](PermsAssetPermissionsUsersRelationsApi.md#PermsAssetPermissionsUsersRelationsCreate) | **Post** /perms/asset-permissions-users-relations/ | perms_asset-permissions-users-relations_create
[**PermsAssetPermissionsUsersRelationsDelete**](PermsAssetPermissionsUsersRelationsApi.md#PermsAssetPermissionsUsersRelationsDelete) | **Delete** /perms/asset-permissions-users-relations/{id}/ | perms_asset-permissions-users-relations_delete
[**PermsAssetPermissionsUsersRelationsList**](PermsAssetPermissionsUsersRelationsApi.md#PermsAssetPermissionsUsersRelationsList) | **Get** /perms/asset-permissions-users-relations/ | perms_asset-permissions-users-relations_list
[**PermsAssetPermissionsUsersRelationsPartialBulkUpdate**](PermsAssetPermissionsUsersRelationsApi.md#PermsAssetPermissionsUsersRelationsPartialBulkUpdate) | **Patch** /perms/asset-permissions-users-relations/ | perms_asset-permissions-users-relations_partial_bulk_update
[**PermsAssetPermissionsUsersRelationsPartialUpdate**](PermsAssetPermissionsUsersRelationsApi.md#PermsAssetPermissionsUsersRelationsPartialUpdate) | **Patch** /perms/asset-permissions-users-relations/{id}/ | perms_asset-permissions-users-relations_partial_update
[**PermsAssetPermissionsUsersRelationsRead**](PermsAssetPermissionsUsersRelationsApi.md#PermsAssetPermissionsUsersRelationsRead) | **Get** /perms/asset-permissions-users-relations/{id}/ | perms_asset-permissions-users-relations_read
[**PermsAssetPermissionsUsersRelationsUpdate**](PermsAssetPermissionsUsersRelationsApi.md#PermsAssetPermissionsUsersRelationsUpdate) | **Put** /perms/asset-permissions-users-relations/{id}/ | perms_asset-permissions-users-relations_update


# **PermsAssetPermissionsUsersRelationsBulkDelete**
> PermsAssetPermissionsUsersRelationsBulkDelete(ctx, optional)
perms_asset-permissions-users-relations_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsAssetPermissionsUsersRelationsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsAssetPermissionsUsersRelationsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **user** | **optional.String**|  | 
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

# **PermsAssetPermissionsUsersRelationsBulkUpdate**
> AssetPermissionUserRelation PermsAssetPermissionsUsersRelationsBulkUpdate(ctx, data)
perms_asset-permissions-users-relations_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionUserRelation**](AssetPermissionUserRelation.md)|  | 

### Return type

[**AssetPermissionUserRelation**](AssetPermissionUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUsersRelationsCreate**
> AssetPermissionUserRelation PermsAssetPermissionsUsersRelationsCreate(ctx, data)
perms_asset-permissions-users-relations_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionUserRelation**](AssetPermissionUserRelation.md)|  | 

### Return type

[**AssetPermissionUserRelation**](AssetPermissionUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUsersRelationsDelete**
> PermsAssetPermissionsUsersRelationsDelete(ctx, id)
perms_asset-permissions-users-relations_delete



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

# **PermsAssetPermissionsUsersRelationsList**
> InlineResponse20034 PermsAssetPermissionsUsersRelationsList(ctx, optional)
perms_asset-permissions-users-relations_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsAssetPermissionsUsersRelationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsAssetPermissionsUsersRelationsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **user** | **optional.String**|  | 
 **assetpermission** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUsersRelationsPartialBulkUpdate**
> AssetPermissionUserRelation PermsAssetPermissionsUsersRelationsPartialBulkUpdate(ctx, data)
perms_asset-permissions-users-relations_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionUserRelation**](AssetPermissionUserRelation.md)|  | 

### Return type

[**AssetPermissionUserRelation**](AssetPermissionUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUsersRelationsPartialUpdate**
> AssetPermissionUserRelation PermsAssetPermissionsUsersRelationsPartialUpdate(ctx, id, data)
perms_asset-permissions-users-relations_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AssetPermissionUserRelation**](AssetPermissionUserRelation.md)|  | 

### Return type

[**AssetPermissionUserRelation**](AssetPermissionUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUsersRelationsRead**
> AssetPermissionUserRelation PermsAssetPermissionsUsersRelationsRead(ctx, id)
perms_asset-permissions-users-relations_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AssetPermissionUserRelation**](AssetPermissionUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUsersRelationsUpdate**
> AssetPermissionUserRelation PermsAssetPermissionsUsersRelationsUpdate(ctx, id, data)
perms_asset-permissions-users-relations_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AssetPermissionUserRelation**](AssetPermissionUserRelation.md)|  | 

### Return type

[**AssetPermissionUserRelation**](AssetPermissionUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

