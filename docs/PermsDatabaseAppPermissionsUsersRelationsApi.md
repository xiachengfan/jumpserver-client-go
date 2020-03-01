# \PermsDatabaseAppPermissionsUsersRelationsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermsDatabaseAppPermissionsUsersRelationsBulkDelete**](PermsDatabaseAppPermissionsUsersRelationsApi.md#PermsDatabaseAppPermissionsUsersRelationsBulkDelete) | **Delete** /perms/database-app-permissions-users-relations/ | perms_database-app-permissions-users-relations_bulk_delete
[**PermsDatabaseAppPermissionsUsersRelationsBulkUpdate**](PermsDatabaseAppPermissionsUsersRelationsApi.md#PermsDatabaseAppPermissionsUsersRelationsBulkUpdate) | **Put** /perms/database-app-permissions-users-relations/ | perms_database-app-permissions-users-relations_bulk_update
[**PermsDatabaseAppPermissionsUsersRelationsCreate**](PermsDatabaseAppPermissionsUsersRelationsApi.md#PermsDatabaseAppPermissionsUsersRelationsCreate) | **Post** /perms/database-app-permissions-users-relations/ | perms_database-app-permissions-users-relations_create
[**PermsDatabaseAppPermissionsUsersRelationsDelete**](PermsDatabaseAppPermissionsUsersRelationsApi.md#PermsDatabaseAppPermissionsUsersRelationsDelete) | **Delete** /perms/database-app-permissions-users-relations/{id}/ | perms_database-app-permissions-users-relations_delete
[**PermsDatabaseAppPermissionsUsersRelationsList**](PermsDatabaseAppPermissionsUsersRelationsApi.md#PermsDatabaseAppPermissionsUsersRelationsList) | **Get** /perms/database-app-permissions-users-relations/ | perms_database-app-permissions-users-relations_list
[**PermsDatabaseAppPermissionsUsersRelationsPartialBulkUpdate**](PermsDatabaseAppPermissionsUsersRelationsApi.md#PermsDatabaseAppPermissionsUsersRelationsPartialBulkUpdate) | **Patch** /perms/database-app-permissions-users-relations/ | perms_database-app-permissions-users-relations_partial_bulk_update
[**PermsDatabaseAppPermissionsUsersRelationsPartialUpdate**](PermsDatabaseAppPermissionsUsersRelationsApi.md#PermsDatabaseAppPermissionsUsersRelationsPartialUpdate) | **Patch** /perms/database-app-permissions-users-relations/{id}/ | perms_database-app-permissions-users-relations_partial_update
[**PermsDatabaseAppPermissionsUsersRelationsRead**](PermsDatabaseAppPermissionsUsersRelationsApi.md#PermsDatabaseAppPermissionsUsersRelationsRead) | **Get** /perms/database-app-permissions-users-relations/{id}/ | perms_database-app-permissions-users-relations_read
[**PermsDatabaseAppPermissionsUsersRelationsUpdate**](PermsDatabaseAppPermissionsUsersRelationsApi.md#PermsDatabaseAppPermissionsUsersRelationsUpdate) | **Put** /perms/database-app-permissions-users-relations/{id}/ | perms_database-app-permissions-users-relations_update


# **PermsDatabaseAppPermissionsUsersRelationsBulkDelete**
> PermsDatabaseAppPermissionsUsersRelationsBulkDelete(ctx, optional)
perms_database-app-permissions-users-relations_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsDatabaseAppPermissionsUsersRelationsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsDatabaseAppPermissionsUsersRelationsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **user** | **optional.String**|  | 
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

# **PermsDatabaseAppPermissionsUsersRelationsBulkUpdate**
> DatabaseAppPermissionUserRelation PermsDatabaseAppPermissionsUsersRelationsBulkUpdate(ctx, data)
perms_database-app-permissions-users-relations_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermissionUserRelation**](DatabaseAppPermissionUserRelation.md)|  | 

### Return type

[**DatabaseAppPermissionUserRelation**](DatabaseAppPermissionUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsUsersRelationsCreate**
> DatabaseAppPermissionUserRelation PermsDatabaseAppPermissionsUsersRelationsCreate(ctx, data)
perms_database-app-permissions-users-relations_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermissionUserRelation**](DatabaseAppPermissionUserRelation.md)|  | 

### Return type

[**DatabaseAppPermissionUserRelation**](DatabaseAppPermissionUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsUsersRelationsDelete**
> PermsDatabaseAppPermissionsUsersRelationsDelete(ctx, id)
perms_database-app-permissions-users-relations_delete



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

# **PermsDatabaseAppPermissionsUsersRelationsList**
> InlineResponse20041 PermsDatabaseAppPermissionsUsersRelationsList(ctx, optional)
perms_database-app-permissions-users-relations_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsDatabaseAppPermissionsUsersRelationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsDatabaseAppPermissionsUsersRelationsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **user** | **optional.String**|  | 
 **databaseapppermission** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsUsersRelationsPartialBulkUpdate**
> DatabaseAppPermissionUserRelation PermsDatabaseAppPermissionsUsersRelationsPartialBulkUpdate(ctx, data)
perms_database-app-permissions-users-relations_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermissionUserRelation**](DatabaseAppPermissionUserRelation.md)|  | 

### Return type

[**DatabaseAppPermissionUserRelation**](DatabaseAppPermissionUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsUsersRelationsPartialUpdate**
> DatabaseAppPermissionUserRelation PermsDatabaseAppPermissionsUsersRelationsPartialUpdate(ctx, id, data)
perms_database-app-permissions-users-relations_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**DatabaseAppPermissionUserRelation**](DatabaseAppPermissionUserRelation.md)|  | 

### Return type

[**DatabaseAppPermissionUserRelation**](DatabaseAppPermissionUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsUsersRelationsRead**
> DatabaseAppPermissionUserRelation PermsDatabaseAppPermissionsUsersRelationsRead(ctx, id)
perms_database-app-permissions-users-relations_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DatabaseAppPermissionUserRelation**](DatabaseAppPermissionUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsUsersRelationsUpdate**
> DatabaseAppPermissionUserRelation PermsDatabaseAppPermissionsUsersRelationsUpdate(ctx, id, data)
perms_database-app-permissions-users-relations_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**DatabaseAppPermissionUserRelation**](DatabaseAppPermissionUserRelation.md)|  | 

### Return type

[**DatabaseAppPermissionUserRelation**](DatabaseAppPermissionUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

