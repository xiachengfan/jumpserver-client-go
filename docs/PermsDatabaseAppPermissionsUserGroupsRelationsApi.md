# \PermsDatabaseAppPermissionsUserGroupsRelationsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermsDatabaseAppPermissionsUserGroupsRelationsBulkDelete**](PermsDatabaseAppPermissionsUserGroupsRelationsApi.md#PermsDatabaseAppPermissionsUserGroupsRelationsBulkDelete) | **Delete** /perms/database-app-permissions-user-groups-relations/ | perms_database-app-permissions-user-groups-relations_bulk_delete
[**PermsDatabaseAppPermissionsUserGroupsRelationsBulkUpdate**](PermsDatabaseAppPermissionsUserGroupsRelationsApi.md#PermsDatabaseAppPermissionsUserGroupsRelationsBulkUpdate) | **Put** /perms/database-app-permissions-user-groups-relations/ | perms_database-app-permissions-user-groups-relations_bulk_update
[**PermsDatabaseAppPermissionsUserGroupsRelationsCreate**](PermsDatabaseAppPermissionsUserGroupsRelationsApi.md#PermsDatabaseAppPermissionsUserGroupsRelationsCreate) | **Post** /perms/database-app-permissions-user-groups-relations/ | perms_database-app-permissions-user-groups-relations_create
[**PermsDatabaseAppPermissionsUserGroupsRelationsDelete**](PermsDatabaseAppPermissionsUserGroupsRelationsApi.md#PermsDatabaseAppPermissionsUserGroupsRelationsDelete) | **Delete** /perms/database-app-permissions-user-groups-relations/{id}/ | perms_database-app-permissions-user-groups-relations_delete
[**PermsDatabaseAppPermissionsUserGroupsRelationsList**](PermsDatabaseAppPermissionsUserGroupsRelationsApi.md#PermsDatabaseAppPermissionsUserGroupsRelationsList) | **Get** /perms/database-app-permissions-user-groups-relations/ | perms_database-app-permissions-user-groups-relations_list
[**PermsDatabaseAppPermissionsUserGroupsRelationsPartialBulkUpdate**](PermsDatabaseAppPermissionsUserGroupsRelationsApi.md#PermsDatabaseAppPermissionsUserGroupsRelationsPartialBulkUpdate) | **Patch** /perms/database-app-permissions-user-groups-relations/ | perms_database-app-permissions-user-groups-relations_partial_bulk_update
[**PermsDatabaseAppPermissionsUserGroupsRelationsPartialUpdate**](PermsDatabaseAppPermissionsUserGroupsRelationsApi.md#PermsDatabaseAppPermissionsUserGroupsRelationsPartialUpdate) | **Patch** /perms/database-app-permissions-user-groups-relations/{id}/ | perms_database-app-permissions-user-groups-relations_partial_update
[**PermsDatabaseAppPermissionsUserGroupsRelationsRead**](PermsDatabaseAppPermissionsUserGroupsRelationsApi.md#PermsDatabaseAppPermissionsUserGroupsRelationsRead) | **Get** /perms/database-app-permissions-user-groups-relations/{id}/ | perms_database-app-permissions-user-groups-relations_read
[**PermsDatabaseAppPermissionsUserGroupsRelationsUpdate**](PermsDatabaseAppPermissionsUserGroupsRelationsApi.md#PermsDatabaseAppPermissionsUserGroupsRelationsUpdate) | **Put** /perms/database-app-permissions-user-groups-relations/{id}/ | perms_database-app-permissions-user-groups-relations_update


# **PermsDatabaseAppPermissionsUserGroupsRelationsBulkDelete**
> PermsDatabaseAppPermissionsUserGroupsRelationsBulkDelete(ctx, optional)
perms_database-app-permissions-user-groups-relations_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsDatabaseAppPermissionsUserGroupsRelationsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsDatabaseAppPermissionsUserGroupsRelationsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **usergroup** | **optional.String**|  | 
 **databaseapppermission** | **optional.String**|  | 
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

# **PermsDatabaseAppPermissionsUserGroupsRelationsBulkUpdate**
> DatabaseAppPermissionUserGroupRelation PermsDatabaseAppPermissionsUserGroupsRelationsBulkUpdate(ctx, data)
perms_database-app-permissions-user-groups-relations_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermissionUserGroupRelation**](DatabaseAppPermissionUserGroupRelation.md)|  | 

### Return type

[**DatabaseAppPermissionUserGroupRelation**](DatabaseAppPermissionUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsUserGroupsRelationsCreate**
> DatabaseAppPermissionUserGroupRelation PermsDatabaseAppPermissionsUserGroupsRelationsCreate(ctx, data)
perms_database-app-permissions-user-groups-relations_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermissionUserGroupRelation**](DatabaseAppPermissionUserGroupRelation.md)|  | 

### Return type

[**DatabaseAppPermissionUserGroupRelation**](DatabaseAppPermissionUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsUserGroupsRelationsDelete**
> PermsDatabaseAppPermissionsUserGroupsRelationsDelete(ctx, id)
perms_database-app-permissions-user-groups-relations_delete



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

# **PermsDatabaseAppPermissionsUserGroupsRelationsList**
> InlineResponse20040 PermsDatabaseAppPermissionsUserGroupsRelationsList(ctx, optional)
perms_database-app-permissions-user-groups-relations_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsDatabaseAppPermissionsUserGroupsRelationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsDatabaseAppPermissionsUserGroupsRelationsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **usergroup** | **optional.String**|  | 
 **databaseapppermission** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsUserGroupsRelationsPartialBulkUpdate**
> DatabaseAppPermissionUserGroupRelation PermsDatabaseAppPermissionsUserGroupsRelationsPartialBulkUpdate(ctx, data)
perms_database-app-permissions-user-groups-relations_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermissionUserGroupRelation**](DatabaseAppPermissionUserGroupRelation.md)|  | 

### Return type

[**DatabaseAppPermissionUserGroupRelation**](DatabaseAppPermissionUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsUserGroupsRelationsPartialUpdate**
> DatabaseAppPermissionUserGroupRelation PermsDatabaseAppPermissionsUserGroupsRelationsPartialUpdate(ctx, id, data)
perms_database-app-permissions-user-groups-relations_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**DatabaseAppPermissionUserGroupRelation**](DatabaseAppPermissionUserGroupRelation.md)|  | 

### Return type

[**DatabaseAppPermissionUserGroupRelation**](DatabaseAppPermissionUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsUserGroupsRelationsRead**
> DatabaseAppPermissionUserGroupRelation PermsDatabaseAppPermissionsUserGroupsRelationsRead(ctx, id)
perms_database-app-permissions-user-groups-relations_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DatabaseAppPermissionUserGroupRelation**](DatabaseAppPermissionUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsUserGroupsRelationsUpdate**
> DatabaseAppPermissionUserGroupRelation PermsDatabaseAppPermissionsUserGroupsRelationsUpdate(ctx, id, data)
perms_database-app-permissions-user-groups-relations_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**DatabaseAppPermissionUserGroupRelation**](DatabaseAppPermissionUserGroupRelation.md)|  | 

### Return type

[**DatabaseAppPermissionUserGroupRelation**](DatabaseAppPermissionUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

