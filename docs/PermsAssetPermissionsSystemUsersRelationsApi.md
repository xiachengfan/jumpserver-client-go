# \PermsAssetPermissionsSystemUsersRelationsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermsAssetPermissionsSystemUsersRelationsBulkDelete**](PermsAssetPermissionsSystemUsersRelationsApi.md#PermsAssetPermissionsSystemUsersRelationsBulkDelete) | **Delete** /perms/asset-permissions-system-users-relations/ | perms_asset-permissions-system-users-relations_bulk_delete
[**PermsAssetPermissionsSystemUsersRelationsBulkUpdate**](PermsAssetPermissionsSystemUsersRelationsApi.md#PermsAssetPermissionsSystemUsersRelationsBulkUpdate) | **Put** /perms/asset-permissions-system-users-relations/ | perms_asset-permissions-system-users-relations_bulk_update
[**PermsAssetPermissionsSystemUsersRelationsCreate**](PermsAssetPermissionsSystemUsersRelationsApi.md#PermsAssetPermissionsSystemUsersRelationsCreate) | **Post** /perms/asset-permissions-system-users-relations/ | perms_asset-permissions-system-users-relations_create
[**PermsAssetPermissionsSystemUsersRelationsDelete**](PermsAssetPermissionsSystemUsersRelationsApi.md#PermsAssetPermissionsSystemUsersRelationsDelete) | **Delete** /perms/asset-permissions-system-users-relations/{id}/ | perms_asset-permissions-system-users-relations_delete
[**PermsAssetPermissionsSystemUsersRelationsList**](PermsAssetPermissionsSystemUsersRelationsApi.md#PermsAssetPermissionsSystemUsersRelationsList) | **Get** /perms/asset-permissions-system-users-relations/ | perms_asset-permissions-system-users-relations_list
[**PermsAssetPermissionsSystemUsersRelationsPartialBulkUpdate**](PermsAssetPermissionsSystemUsersRelationsApi.md#PermsAssetPermissionsSystemUsersRelationsPartialBulkUpdate) | **Patch** /perms/asset-permissions-system-users-relations/ | perms_asset-permissions-system-users-relations_partial_bulk_update
[**PermsAssetPermissionsSystemUsersRelationsPartialUpdate**](PermsAssetPermissionsSystemUsersRelationsApi.md#PermsAssetPermissionsSystemUsersRelationsPartialUpdate) | **Patch** /perms/asset-permissions-system-users-relations/{id}/ | perms_asset-permissions-system-users-relations_partial_update
[**PermsAssetPermissionsSystemUsersRelationsRead**](PermsAssetPermissionsSystemUsersRelationsApi.md#PermsAssetPermissionsSystemUsersRelationsRead) | **Get** /perms/asset-permissions-system-users-relations/{id}/ | perms_asset-permissions-system-users-relations_read
[**PermsAssetPermissionsSystemUsersRelationsUpdate**](PermsAssetPermissionsSystemUsersRelationsApi.md#PermsAssetPermissionsSystemUsersRelationsUpdate) | **Put** /perms/asset-permissions-system-users-relations/{id}/ | perms_asset-permissions-system-users-relations_update


# **PermsAssetPermissionsSystemUsersRelationsBulkDelete**
> PermsAssetPermissionsSystemUsersRelationsBulkDelete(ctx, optional)
perms_asset-permissions-system-users-relations_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsAssetPermissionsSystemUsersRelationsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsAssetPermissionsSystemUsersRelationsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **systemuser** | **optional.String**|  | 
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

# **PermsAssetPermissionsSystemUsersRelationsBulkUpdate**
> AssetPermissionSystemUserRelation PermsAssetPermissionsSystemUsersRelationsBulkUpdate(ctx, data)
perms_asset-permissions-system-users-relations_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionSystemUserRelation**](AssetPermissionSystemUserRelation.md)|  | 

### Return type

[**AssetPermissionSystemUserRelation**](AssetPermissionSystemUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsSystemUsersRelationsCreate**
> AssetPermissionSystemUserRelation PermsAssetPermissionsSystemUsersRelationsCreate(ctx, data)
perms_asset-permissions-system-users-relations_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionSystemUserRelation**](AssetPermissionSystemUserRelation.md)|  | 

### Return type

[**AssetPermissionSystemUserRelation**](AssetPermissionSystemUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsSystemUsersRelationsDelete**
> PermsAssetPermissionsSystemUsersRelationsDelete(ctx, id)
perms_asset-permissions-system-users-relations_delete



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

# **PermsAssetPermissionsSystemUsersRelationsList**
> InlineResponse20032 PermsAssetPermissionsSystemUsersRelationsList(ctx, optional)
perms_asset-permissions-system-users-relations_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsAssetPermissionsSystemUsersRelationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsAssetPermissionsSystemUsersRelationsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **systemuser** | **optional.String**|  | 
 **assetpermission** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsSystemUsersRelationsPartialBulkUpdate**
> AssetPermissionSystemUserRelation PermsAssetPermissionsSystemUsersRelationsPartialBulkUpdate(ctx, data)
perms_asset-permissions-system-users-relations_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionSystemUserRelation**](AssetPermissionSystemUserRelation.md)|  | 

### Return type

[**AssetPermissionSystemUserRelation**](AssetPermissionSystemUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsSystemUsersRelationsPartialUpdate**
> AssetPermissionSystemUserRelation PermsAssetPermissionsSystemUsersRelationsPartialUpdate(ctx, id, data)
perms_asset-permissions-system-users-relations_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AssetPermissionSystemUserRelation**](AssetPermissionSystemUserRelation.md)|  | 

### Return type

[**AssetPermissionSystemUserRelation**](AssetPermissionSystemUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsSystemUsersRelationsRead**
> AssetPermissionSystemUserRelation PermsAssetPermissionsSystemUsersRelationsRead(ctx, id)
perms_asset-permissions-system-users-relations_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AssetPermissionSystemUserRelation**](AssetPermissionSystemUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsSystemUsersRelationsUpdate**
> AssetPermissionSystemUserRelation PermsAssetPermissionsSystemUsersRelationsUpdate(ctx, id, data)
perms_asset-permissions-system-users-relations_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AssetPermissionSystemUserRelation**](AssetPermissionSystemUserRelation.md)|  | 

### Return type

[**AssetPermissionSystemUserRelation**](AssetPermissionSystemUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

