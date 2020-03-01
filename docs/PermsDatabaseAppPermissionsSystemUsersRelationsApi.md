# \PermsDatabaseAppPermissionsSystemUsersRelationsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermsDatabaseAppPermissionsSystemUsersRelationsBulkDelete**](PermsDatabaseAppPermissionsSystemUsersRelationsApi.md#PermsDatabaseAppPermissionsSystemUsersRelationsBulkDelete) | **Delete** /perms/database-app-permissions-system-users-relations/ | perms_database-app-permissions-system-users-relations_bulk_delete
[**PermsDatabaseAppPermissionsSystemUsersRelationsBulkUpdate**](PermsDatabaseAppPermissionsSystemUsersRelationsApi.md#PermsDatabaseAppPermissionsSystemUsersRelationsBulkUpdate) | **Put** /perms/database-app-permissions-system-users-relations/ | perms_database-app-permissions-system-users-relations_bulk_update
[**PermsDatabaseAppPermissionsSystemUsersRelationsCreate**](PermsDatabaseAppPermissionsSystemUsersRelationsApi.md#PermsDatabaseAppPermissionsSystemUsersRelationsCreate) | **Post** /perms/database-app-permissions-system-users-relations/ | perms_database-app-permissions-system-users-relations_create
[**PermsDatabaseAppPermissionsSystemUsersRelationsDelete**](PermsDatabaseAppPermissionsSystemUsersRelationsApi.md#PermsDatabaseAppPermissionsSystemUsersRelationsDelete) | **Delete** /perms/database-app-permissions-system-users-relations/{id}/ | perms_database-app-permissions-system-users-relations_delete
[**PermsDatabaseAppPermissionsSystemUsersRelationsList**](PermsDatabaseAppPermissionsSystemUsersRelationsApi.md#PermsDatabaseAppPermissionsSystemUsersRelationsList) | **Get** /perms/database-app-permissions-system-users-relations/ | perms_database-app-permissions-system-users-relations_list
[**PermsDatabaseAppPermissionsSystemUsersRelationsPartialBulkUpdate**](PermsDatabaseAppPermissionsSystemUsersRelationsApi.md#PermsDatabaseAppPermissionsSystemUsersRelationsPartialBulkUpdate) | **Patch** /perms/database-app-permissions-system-users-relations/ | perms_database-app-permissions-system-users-relations_partial_bulk_update
[**PermsDatabaseAppPermissionsSystemUsersRelationsPartialUpdate**](PermsDatabaseAppPermissionsSystemUsersRelationsApi.md#PermsDatabaseAppPermissionsSystemUsersRelationsPartialUpdate) | **Patch** /perms/database-app-permissions-system-users-relations/{id}/ | perms_database-app-permissions-system-users-relations_partial_update
[**PermsDatabaseAppPermissionsSystemUsersRelationsRead**](PermsDatabaseAppPermissionsSystemUsersRelationsApi.md#PermsDatabaseAppPermissionsSystemUsersRelationsRead) | **Get** /perms/database-app-permissions-system-users-relations/{id}/ | perms_database-app-permissions-system-users-relations_read
[**PermsDatabaseAppPermissionsSystemUsersRelationsUpdate**](PermsDatabaseAppPermissionsSystemUsersRelationsApi.md#PermsDatabaseAppPermissionsSystemUsersRelationsUpdate) | **Put** /perms/database-app-permissions-system-users-relations/{id}/ | perms_database-app-permissions-system-users-relations_update


# **PermsDatabaseAppPermissionsSystemUsersRelationsBulkDelete**
> PermsDatabaseAppPermissionsSystemUsersRelationsBulkDelete(ctx, optional)
perms_database-app-permissions-system-users-relations_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsDatabaseAppPermissionsSystemUsersRelationsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsDatabaseAppPermissionsSystemUsersRelationsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **systemuser** | **optional.String**|  | 
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

# **PermsDatabaseAppPermissionsSystemUsersRelationsBulkUpdate**
> DatabaseAppPermissionSystemUserRelation PermsDatabaseAppPermissionsSystemUsersRelationsBulkUpdate(ctx, data)
perms_database-app-permissions-system-users-relations_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermissionSystemUserRelation**](DatabaseAppPermissionSystemUserRelation.md)|  | 

### Return type

[**DatabaseAppPermissionSystemUserRelation**](DatabaseAppPermissionSystemUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsSystemUsersRelationsCreate**
> DatabaseAppPermissionSystemUserRelation PermsDatabaseAppPermissionsSystemUsersRelationsCreate(ctx, data)
perms_database-app-permissions-system-users-relations_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermissionSystemUserRelation**](DatabaseAppPermissionSystemUserRelation.md)|  | 

### Return type

[**DatabaseAppPermissionSystemUserRelation**](DatabaseAppPermissionSystemUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsSystemUsersRelationsDelete**
> PermsDatabaseAppPermissionsSystemUsersRelationsDelete(ctx, id)
perms_database-app-permissions-system-users-relations_delete



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

# **PermsDatabaseAppPermissionsSystemUsersRelationsList**
> InlineResponse20039 PermsDatabaseAppPermissionsSystemUsersRelationsList(ctx, optional)
perms_database-app-permissions-system-users-relations_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsDatabaseAppPermissionsSystemUsersRelationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsDatabaseAppPermissionsSystemUsersRelationsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **systemuser** | **optional.String**|  | 
 **databaseapppermission** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsSystemUsersRelationsPartialBulkUpdate**
> DatabaseAppPermissionSystemUserRelation PermsDatabaseAppPermissionsSystemUsersRelationsPartialBulkUpdate(ctx, data)
perms_database-app-permissions-system-users-relations_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermissionSystemUserRelation**](DatabaseAppPermissionSystemUserRelation.md)|  | 

### Return type

[**DatabaseAppPermissionSystemUserRelation**](DatabaseAppPermissionSystemUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsSystemUsersRelationsPartialUpdate**
> DatabaseAppPermissionSystemUserRelation PermsDatabaseAppPermissionsSystemUsersRelationsPartialUpdate(ctx, id, data)
perms_database-app-permissions-system-users-relations_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**DatabaseAppPermissionSystemUserRelation**](DatabaseAppPermissionSystemUserRelation.md)|  | 

### Return type

[**DatabaseAppPermissionSystemUserRelation**](DatabaseAppPermissionSystemUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsSystemUsersRelationsRead**
> DatabaseAppPermissionSystemUserRelation PermsDatabaseAppPermissionsSystemUsersRelationsRead(ctx, id)
perms_database-app-permissions-system-users-relations_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DatabaseAppPermissionSystemUserRelation**](DatabaseAppPermissionSystemUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsSystemUsersRelationsUpdate**
> DatabaseAppPermissionSystemUserRelation PermsDatabaseAppPermissionsSystemUsersRelationsUpdate(ctx, id, data)
perms_database-app-permissions-system-users-relations_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**DatabaseAppPermissionSystemUserRelation**](DatabaseAppPermissionSystemUserRelation.md)|  | 

### Return type

[**DatabaseAppPermissionSystemUserRelation**](DatabaseAppPermissionSystemUserRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

