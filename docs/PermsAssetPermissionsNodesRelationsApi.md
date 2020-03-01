# \PermsAssetPermissionsNodesRelationsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermsAssetPermissionsNodesRelationsBulkDelete**](PermsAssetPermissionsNodesRelationsApi.md#PermsAssetPermissionsNodesRelationsBulkDelete) | **Delete** /perms/asset-permissions-nodes-relations/ | perms_asset-permissions-nodes-relations_bulk_delete
[**PermsAssetPermissionsNodesRelationsBulkUpdate**](PermsAssetPermissionsNodesRelationsApi.md#PermsAssetPermissionsNodesRelationsBulkUpdate) | **Put** /perms/asset-permissions-nodes-relations/ | perms_asset-permissions-nodes-relations_bulk_update
[**PermsAssetPermissionsNodesRelationsCreate**](PermsAssetPermissionsNodesRelationsApi.md#PermsAssetPermissionsNodesRelationsCreate) | **Post** /perms/asset-permissions-nodes-relations/ | perms_asset-permissions-nodes-relations_create
[**PermsAssetPermissionsNodesRelationsDelete**](PermsAssetPermissionsNodesRelationsApi.md#PermsAssetPermissionsNodesRelationsDelete) | **Delete** /perms/asset-permissions-nodes-relations/{id}/ | perms_asset-permissions-nodes-relations_delete
[**PermsAssetPermissionsNodesRelationsList**](PermsAssetPermissionsNodesRelationsApi.md#PermsAssetPermissionsNodesRelationsList) | **Get** /perms/asset-permissions-nodes-relations/ | perms_asset-permissions-nodes-relations_list
[**PermsAssetPermissionsNodesRelationsPartialBulkUpdate**](PermsAssetPermissionsNodesRelationsApi.md#PermsAssetPermissionsNodesRelationsPartialBulkUpdate) | **Patch** /perms/asset-permissions-nodes-relations/ | perms_asset-permissions-nodes-relations_partial_bulk_update
[**PermsAssetPermissionsNodesRelationsPartialUpdate**](PermsAssetPermissionsNodesRelationsApi.md#PermsAssetPermissionsNodesRelationsPartialUpdate) | **Patch** /perms/asset-permissions-nodes-relations/{id}/ | perms_asset-permissions-nodes-relations_partial_update
[**PermsAssetPermissionsNodesRelationsRead**](PermsAssetPermissionsNodesRelationsApi.md#PermsAssetPermissionsNodesRelationsRead) | **Get** /perms/asset-permissions-nodes-relations/{id}/ | perms_asset-permissions-nodes-relations_read
[**PermsAssetPermissionsNodesRelationsUpdate**](PermsAssetPermissionsNodesRelationsApi.md#PermsAssetPermissionsNodesRelationsUpdate) | **Put** /perms/asset-permissions-nodes-relations/{id}/ | perms_asset-permissions-nodes-relations_update


# **PermsAssetPermissionsNodesRelationsBulkDelete**
> PermsAssetPermissionsNodesRelationsBulkDelete(ctx, optional)
perms_asset-permissions-nodes-relations_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsAssetPermissionsNodesRelationsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsAssetPermissionsNodesRelationsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **node** | **optional.String**|  | 
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

# **PermsAssetPermissionsNodesRelationsBulkUpdate**
> AssetPermissionNodeRelation PermsAssetPermissionsNodesRelationsBulkUpdate(ctx, data)
perms_asset-permissions-nodes-relations_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionNodeRelation**](AssetPermissionNodeRelation.md)|  | 

### Return type

[**AssetPermissionNodeRelation**](AssetPermissionNodeRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsNodesRelationsCreate**
> AssetPermissionNodeRelation PermsAssetPermissionsNodesRelationsCreate(ctx, data)
perms_asset-permissions-nodes-relations_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionNodeRelation**](AssetPermissionNodeRelation.md)|  | 

### Return type

[**AssetPermissionNodeRelation**](AssetPermissionNodeRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsNodesRelationsDelete**
> PermsAssetPermissionsNodesRelationsDelete(ctx, id)
perms_asset-permissions-nodes-relations_delete



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

# **PermsAssetPermissionsNodesRelationsList**
> InlineResponse20031 PermsAssetPermissionsNodesRelationsList(ctx, optional)
perms_asset-permissions-nodes-relations_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsAssetPermissionsNodesRelationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsAssetPermissionsNodesRelationsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **node** | **optional.String**|  | 
 **assetpermission** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsNodesRelationsPartialBulkUpdate**
> AssetPermissionNodeRelation PermsAssetPermissionsNodesRelationsPartialBulkUpdate(ctx, data)
perms_asset-permissions-nodes-relations_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionNodeRelation**](AssetPermissionNodeRelation.md)|  | 

### Return type

[**AssetPermissionNodeRelation**](AssetPermissionNodeRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsNodesRelationsPartialUpdate**
> AssetPermissionNodeRelation PermsAssetPermissionsNodesRelationsPartialUpdate(ctx, id, data)
perms_asset-permissions-nodes-relations_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AssetPermissionNodeRelation**](AssetPermissionNodeRelation.md)|  | 

### Return type

[**AssetPermissionNodeRelation**](AssetPermissionNodeRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsNodesRelationsRead**
> AssetPermissionNodeRelation PermsAssetPermissionsNodesRelationsRead(ctx, id)
perms_asset-permissions-nodes-relations_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AssetPermissionNodeRelation**](AssetPermissionNodeRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsAssetPermissionsNodesRelationsUpdate**
> AssetPermissionNodeRelation PermsAssetPermissionsNodesRelationsUpdate(ctx, id, data)
perms_asset-permissions-nodes-relations_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AssetPermissionNodeRelation**](AssetPermissionNodeRelation.md)|  | 

### Return type

[**AssetPermissionNodeRelation**](AssetPermissionNodeRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

