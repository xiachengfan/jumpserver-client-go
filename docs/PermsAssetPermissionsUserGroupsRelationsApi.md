# \PermsAssetPermissionsUserGroupsRelationsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermsAssetPermissionsUserGroupsRelationsBulkDelete**](PermsAssetPermissionsUserGroupsRelationsApi.md#PermsAssetPermissionsUserGroupsRelationsBulkDelete) | **Delete** /perms/asset-permissions-user-groups-relations/ | perms_asset-permissions-user-groups-relations_bulk_delete
[**PermsAssetPermissionsUserGroupsRelationsBulkUpdate**](PermsAssetPermissionsUserGroupsRelationsApi.md#PermsAssetPermissionsUserGroupsRelationsBulkUpdate) | **Put** /perms/asset-permissions-user-groups-relations/ | perms_asset-permissions-user-groups-relations_bulk_update
[**PermsAssetPermissionsUserGroupsRelationsCreate**](PermsAssetPermissionsUserGroupsRelationsApi.md#PermsAssetPermissionsUserGroupsRelationsCreate) | **Post** /perms/asset-permissions-user-groups-relations/ | perms_asset-permissions-user-groups-relations_create
[**PermsAssetPermissionsUserGroupsRelationsDelete**](PermsAssetPermissionsUserGroupsRelationsApi.md#PermsAssetPermissionsUserGroupsRelationsDelete) | **Delete** /perms/asset-permissions-user-groups-relations/{id}/ | perms_asset-permissions-user-groups-relations_delete
[**PermsAssetPermissionsUserGroupsRelationsList**](PermsAssetPermissionsUserGroupsRelationsApi.md#PermsAssetPermissionsUserGroupsRelationsList) | **Get** /perms/asset-permissions-user-groups-relations/ | perms_asset-permissions-user-groups-relations_list
[**PermsAssetPermissionsUserGroupsRelationsPartialBulkUpdate**](PermsAssetPermissionsUserGroupsRelationsApi.md#PermsAssetPermissionsUserGroupsRelationsPartialBulkUpdate) | **Patch** /perms/asset-permissions-user-groups-relations/ | perms_asset-permissions-user-groups-relations_partial_bulk_update
[**PermsAssetPermissionsUserGroupsRelationsPartialUpdate**](PermsAssetPermissionsUserGroupsRelationsApi.md#PermsAssetPermissionsUserGroupsRelationsPartialUpdate) | **Patch** /perms/asset-permissions-user-groups-relations/{id}/ | perms_asset-permissions-user-groups-relations_partial_update
[**PermsAssetPermissionsUserGroupsRelationsRead**](PermsAssetPermissionsUserGroupsRelationsApi.md#PermsAssetPermissionsUserGroupsRelationsRead) | **Get** /perms/asset-permissions-user-groups-relations/{id}/ | perms_asset-permissions-user-groups-relations_read
[**PermsAssetPermissionsUserGroupsRelationsUpdate**](PermsAssetPermissionsUserGroupsRelationsApi.md#PermsAssetPermissionsUserGroupsRelationsUpdate) | **Put** /perms/asset-permissions-user-groups-relations/{id}/ | perms_asset-permissions-user-groups-relations_update


# **PermsAssetPermissionsUserGroupsRelationsBulkDelete**
> PermsAssetPermissionsUserGroupsRelationsBulkDelete(ctx, optional)
perms_asset-permissions-user-groups-relations_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsAssetPermissionsUserGroupsRelationsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsAssetPermissionsUserGroupsRelationsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **usergroup** | **optional.String**|  | 
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

# **PermsAssetPermissionsUserGroupsRelationsBulkUpdate**
> AssetPermissionUserGroupRelation PermsAssetPermissionsUserGroupsRelationsBulkUpdate(ctx, data)
perms_asset-permissions-user-groups-relations_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionUserGroupRelation**](AssetPermissionUserGroupRelation.md)|  | 

### Return type

[**AssetPermissionUserGroupRelation**](AssetPermissionUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUserGroupsRelationsCreate**
> AssetPermissionUserGroupRelation PermsAssetPermissionsUserGroupsRelationsCreate(ctx, data)
perms_asset-permissions-user-groups-relations_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionUserGroupRelation**](AssetPermissionUserGroupRelation.md)|  | 

### Return type

[**AssetPermissionUserGroupRelation**](AssetPermissionUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUserGroupsRelationsDelete**
> PermsAssetPermissionsUserGroupsRelationsDelete(ctx, id)
perms_asset-permissions-user-groups-relations_delete



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

# **PermsAssetPermissionsUserGroupsRelationsList**
> InlineResponse20033 PermsAssetPermissionsUserGroupsRelationsList(ctx, optional)
perms_asset-permissions-user-groups-relations_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsAssetPermissionsUserGroupsRelationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsAssetPermissionsUserGroupsRelationsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **usergroup** | **optional.String**|  | 
 **assetpermission** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUserGroupsRelationsPartialBulkUpdate**
> AssetPermissionUserGroupRelation PermsAssetPermissionsUserGroupsRelationsPartialBulkUpdate(ctx, data)
perms_asset-permissions-user-groups-relations_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionUserGroupRelation**](AssetPermissionUserGroupRelation.md)|  | 

### Return type

[**AssetPermissionUserGroupRelation**](AssetPermissionUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUserGroupsRelationsPartialUpdate**
> AssetPermissionUserGroupRelation PermsAssetPermissionsUserGroupsRelationsPartialUpdate(ctx, id, data)
perms_asset-permissions-user-groups-relations_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AssetPermissionUserGroupRelation**](AssetPermissionUserGroupRelation.md)|  | 

### Return type

[**AssetPermissionUserGroupRelation**](AssetPermissionUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUserGroupsRelationsRead**
> AssetPermissionUserGroupRelation PermsAssetPermissionsUserGroupsRelationsRead(ctx, id)
perms_asset-permissions-user-groups-relations_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AssetPermissionUserGroupRelation**](AssetPermissionUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsUserGroupsRelationsUpdate**
> AssetPermissionUserGroupRelation PermsAssetPermissionsUserGroupsRelationsUpdate(ctx, id, data)
perms_asset-permissions-user-groups-relations_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AssetPermissionUserGroupRelation**](AssetPermissionUserGroupRelation.md)|  | 

### Return type

[**AssetPermissionUserGroupRelation**](AssetPermissionUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

