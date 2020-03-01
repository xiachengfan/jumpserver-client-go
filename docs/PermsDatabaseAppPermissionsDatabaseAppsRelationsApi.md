# \PermsDatabaseAppPermissionsDatabaseAppsRelationsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermsDatabaseAppPermissionsDatabaseAppsRelationsBulkDelete**](PermsDatabaseAppPermissionsDatabaseAppsRelationsApi.md#PermsDatabaseAppPermissionsDatabaseAppsRelationsBulkDelete) | **Delete** /perms/database-app-permissions-database-apps-relations/ | perms_database-app-permissions-database-apps-relations_bulk_delete
[**PermsDatabaseAppPermissionsDatabaseAppsRelationsBulkUpdate**](PermsDatabaseAppPermissionsDatabaseAppsRelationsApi.md#PermsDatabaseAppPermissionsDatabaseAppsRelationsBulkUpdate) | **Put** /perms/database-app-permissions-database-apps-relations/ | perms_database-app-permissions-database-apps-relations_bulk_update
[**PermsDatabaseAppPermissionsDatabaseAppsRelationsCreate**](PermsDatabaseAppPermissionsDatabaseAppsRelationsApi.md#PermsDatabaseAppPermissionsDatabaseAppsRelationsCreate) | **Post** /perms/database-app-permissions-database-apps-relations/ | perms_database-app-permissions-database-apps-relations_create
[**PermsDatabaseAppPermissionsDatabaseAppsRelationsDelete**](PermsDatabaseAppPermissionsDatabaseAppsRelationsApi.md#PermsDatabaseAppPermissionsDatabaseAppsRelationsDelete) | **Delete** /perms/database-app-permissions-database-apps-relations/{id}/ | perms_database-app-permissions-database-apps-relations_delete
[**PermsDatabaseAppPermissionsDatabaseAppsRelationsList**](PermsDatabaseAppPermissionsDatabaseAppsRelationsApi.md#PermsDatabaseAppPermissionsDatabaseAppsRelationsList) | **Get** /perms/database-app-permissions-database-apps-relations/ | perms_database-app-permissions-database-apps-relations_list
[**PermsDatabaseAppPermissionsDatabaseAppsRelationsPartialBulkUpdate**](PermsDatabaseAppPermissionsDatabaseAppsRelationsApi.md#PermsDatabaseAppPermissionsDatabaseAppsRelationsPartialBulkUpdate) | **Patch** /perms/database-app-permissions-database-apps-relations/ | perms_database-app-permissions-database-apps-relations_partial_bulk_update
[**PermsDatabaseAppPermissionsDatabaseAppsRelationsPartialUpdate**](PermsDatabaseAppPermissionsDatabaseAppsRelationsApi.md#PermsDatabaseAppPermissionsDatabaseAppsRelationsPartialUpdate) | **Patch** /perms/database-app-permissions-database-apps-relations/{id}/ | perms_database-app-permissions-database-apps-relations_partial_update
[**PermsDatabaseAppPermissionsDatabaseAppsRelationsRead**](PermsDatabaseAppPermissionsDatabaseAppsRelationsApi.md#PermsDatabaseAppPermissionsDatabaseAppsRelationsRead) | **Get** /perms/database-app-permissions-database-apps-relations/{id}/ | perms_database-app-permissions-database-apps-relations_read
[**PermsDatabaseAppPermissionsDatabaseAppsRelationsUpdate**](PermsDatabaseAppPermissionsDatabaseAppsRelationsApi.md#PermsDatabaseAppPermissionsDatabaseAppsRelationsUpdate) | **Put** /perms/database-app-permissions-database-apps-relations/{id}/ | perms_database-app-permissions-database-apps-relations_update


# **PermsDatabaseAppPermissionsDatabaseAppsRelationsBulkDelete**
> PermsDatabaseAppPermissionsDatabaseAppsRelationsBulkDelete(ctx, optional)
perms_database-app-permissions-database-apps-relations_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsDatabaseAppPermissionsDatabaseAppsRelationsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsDatabaseAppPermissionsDatabaseAppsRelationsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **databaseapp** | **optional.String**|  | 
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

# **PermsDatabaseAppPermissionsDatabaseAppsRelationsBulkUpdate**
> DatabaseAppPermissionDatabaseAppRelation PermsDatabaseAppPermissionsDatabaseAppsRelationsBulkUpdate(ctx, data)
perms_database-app-permissions-database-apps-relations_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermissionDatabaseAppRelation**](DatabaseAppPermissionDatabaseAppRelation.md)|  | 

### Return type

[**DatabaseAppPermissionDatabaseAppRelation**](DatabaseAppPermissionDatabaseAppRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsDatabaseAppsRelationsCreate**
> DatabaseAppPermissionDatabaseAppRelation PermsDatabaseAppPermissionsDatabaseAppsRelationsCreate(ctx, data)
perms_database-app-permissions-database-apps-relations_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermissionDatabaseAppRelation**](DatabaseAppPermissionDatabaseAppRelation.md)|  | 

### Return type

[**DatabaseAppPermissionDatabaseAppRelation**](DatabaseAppPermissionDatabaseAppRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsDatabaseAppsRelationsDelete**
> PermsDatabaseAppPermissionsDatabaseAppsRelationsDelete(ctx, id)
perms_database-app-permissions-database-apps-relations_delete



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

# **PermsDatabaseAppPermissionsDatabaseAppsRelationsList**
> InlineResponse20038 PermsDatabaseAppPermissionsDatabaseAppsRelationsList(ctx, optional)
perms_database-app-permissions-database-apps-relations_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsDatabaseAppPermissionsDatabaseAppsRelationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsDatabaseAppPermissionsDatabaseAppsRelationsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **databaseapp** | **optional.String**|  | 
 **databaseapppermission** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsDatabaseAppsRelationsPartialBulkUpdate**
> DatabaseAppPermissionDatabaseAppRelation PermsDatabaseAppPermissionsDatabaseAppsRelationsPartialBulkUpdate(ctx, data)
perms_database-app-permissions-database-apps-relations_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermissionDatabaseAppRelation**](DatabaseAppPermissionDatabaseAppRelation.md)|  | 

### Return type

[**DatabaseAppPermissionDatabaseAppRelation**](DatabaseAppPermissionDatabaseAppRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsDatabaseAppsRelationsPartialUpdate**
> DatabaseAppPermissionDatabaseAppRelation PermsDatabaseAppPermissionsDatabaseAppsRelationsPartialUpdate(ctx, id, data)
perms_database-app-permissions-database-apps-relations_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**DatabaseAppPermissionDatabaseAppRelation**](DatabaseAppPermissionDatabaseAppRelation.md)|  | 

### Return type

[**DatabaseAppPermissionDatabaseAppRelation**](DatabaseAppPermissionDatabaseAppRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsDatabaseAppsRelationsRead**
> DatabaseAppPermissionDatabaseAppRelation PermsDatabaseAppPermissionsDatabaseAppsRelationsRead(ctx, id)
perms_database-app-permissions-database-apps-relations_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DatabaseAppPermissionDatabaseAppRelation**](DatabaseAppPermissionDatabaseAppRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsDatabaseAppsRelationsUpdate**
> DatabaseAppPermissionDatabaseAppRelation PermsDatabaseAppPermissionsDatabaseAppsRelationsUpdate(ctx, id, data)
perms_database-app-permissions-database-apps-relations_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**DatabaseAppPermissionDatabaseAppRelation**](DatabaseAppPermissionDatabaseAppRelation.md)|  | 

### Return type

[**DatabaseAppPermissionDatabaseAppRelation**](DatabaseAppPermissionDatabaseAppRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

