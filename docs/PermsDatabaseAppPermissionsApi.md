# \PermsDatabaseAppPermissionsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermsDatabaseAppPermissionsBulkDelete**](PermsDatabaseAppPermissionsApi.md#PermsDatabaseAppPermissionsBulkDelete) | **Delete** /perms/database-app-permissions/ | perms_database-app-permissions_bulk_delete
[**PermsDatabaseAppPermissionsBulkUpdate**](PermsDatabaseAppPermissionsApi.md#PermsDatabaseAppPermissionsBulkUpdate) | **Put** /perms/database-app-permissions/ | perms_database-app-permissions_bulk_update
[**PermsDatabaseAppPermissionsCreate**](PermsDatabaseAppPermissionsApi.md#PermsDatabaseAppPermissionsCreate) | **Post** /perms/database-app-permissions/ | perms_database-app-permissions_create
[**PermsDatabaseAppPermissionsDatabaseAppsAllList**](PermsDatabaseAppPermissionsApi.md#PermsDatabaseAppPermissionsDatabaseAppsAllList) | **Get** /perms/database-app-permissions/{id}/database-apps/all/ | perms_database-app-permissions_database-apps_all_list
[**PermsDatabaseAppPermissionsDelete**](PermsDatabaseAppPermissionsApi.md#PermsDatabaseAppPermissionsDelete) | **Delete** /perms/database-app-permissions/{id}/ | perms_database-app-permissions_delete
[**PermsDatabaseAppPermissionsList**](PermsDatabaseAppPermissionsApi.md#PermsDatabaseAppPermissionsList) | **Get** /perms/database-app-permissions/ | perms_database-app-permissions_list
[**PermsDatabaseAppPermissionsPartialBulkUpdate**](PermsDatabaseAppPermissionsApi.md#PermsDatabaseAppPermissionsPartialBulkUpdate) | **Patch** /perms/database-app-permissions/ | perms_database-app-permissions_partial_bulk_update
[**PermsDatabaseAppPermissionsPartialUpdate**](PermsDatabaseAppPermissionsApi.md#PermsDatabaseAppPermissionsPartialUpdate) | **Patch** /perms/database-app-permissions/{id}/ | perms_database-app-permissions_partial_update
[**PermsDatabaseAppPermissionsRead**](PermsDatabaseAppPermissionsApi.md#PermsDatabaseAppPermissionsRead) | **Get** /perms/database-app-permissions/{id}/ | perms_database-app-permissions_read
[**PermsDatabaseAppPermissionsUpdate**](PermsDatabaseAppPermissionsApi.md#PermsDatabaseAppPermissionsUpdate) | **Put** /perms/database-app-permissions/{id}/ | perms_database-app-permissions_update
[**PermsDatabaseAppPermissionsUserValidateList**](PermsDatabaseAppPermissionsApi.md#PermsDatabaseAppPermissionsUserValidateList) | **Get** /perms/database-app-permissions/user/validate/ | perms_database-app-permissions_user_validate_list
[**PermsDatabaseAppPermissionsUsersAllList**](PermsDatabaseAppPermissionsApi.md#PermsDatabaseAppPermissionsUsersAllList) | **Get** /perms/database-app-permissions/{id}/users/all/ | perms_database-app-permissions_users_all_list


# **PermsDatabaseAppPermissionsBulkDelete**
> PermsDatabaseAppPermissionsBulkDelete(ctx, optional)
perms_database-app-permissions_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsDatabaseAppPermissionsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsDatabaseAppPermissionsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
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

# **PermsDatabaseAppPermissionsBulkUpdate**
> DatabaseAppPermission PermsDatabaseAppPermissionsBulkUpdate(ctx, data)
perms_database-app-permissions_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermission**](DatabaseAppPermission.md)|  | 

### Return type

[**DatabaseAppPermission**](DatabaseAppPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsCreate**
> DatabaseAppPermission PermsDatabaseAppPermissionsCreate(ctx, data)
perms_database-app-permissions_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermission**](DatabaseAppPermission.md)|  | 

### Return type

[**DatabaseAppPermission**](DatabaseAppPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsDatabaseAppsAllList**
> InlineResponse20043 PermsDatabaseAppPermissionsDatabaseAppsAllList(ctx, id, optional)
perms_database-app-permissions_database-apps_all_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***PermsDatabaseAppPermissionsDatabaseAppsAllListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsDatabaseAppPermissionsDatabaseAppsAllListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsDelete**
> PermsDatabaseAppPermissionsDelete(ctx, id)
perms_database-app-permissions_delete



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

# **PermsDatabaseAppPermissionsList**
> InlineResponse20042 PermsDatabaseAppPermissionsList(ctx, optional)
perms_database-app-permissions_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsDatabaseAppPermissionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsDatabaseAppPermissionsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20042**](inline_response_200_42.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsPartialBulkUpdate**
> DatabaseAppPermission PermsDatabaseAppPermissionsPartialBulkUpdate(ctx, data)
perms_database-app-permissions_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseAppPermission**](DatabaseAppPermission.md)|  | 

### Return type

[**DatabaseAppPermission**](DatabaseAppPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsPartialUpdate**
> DatabaseAppPermission PermsDatabaseAppPermissionsPartialUpdate(ctx, id, data)
perms_database-app-permissions_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**DatabaseAppPermission**](DatabaseAppPermission.md)|  | 

### Return type

[**DatabaseAppPermission**](DatabaseAppPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsRead**
> DatabaseAppPermission PermsDatabaseAppPermissionsRead(ctx, id)
perms_database-app-permissions_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DatabaseAppPermission**](DatabaseAppPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsUpdate**
> DatabaseAppPermission PermsDatabaseAppPermissionsUpdate(ctx, id, data)
perms_database-app-permissions_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**DatabaseAppPermission**](DatabaseAppPermission.md)|  | 

### Return type

[**DatabaseAppPermission**](DatabaseAppPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsDatabaseAppPermissionsUserValidateList**
> PermsDatabaseAppPermissionsUserValidateList(ctx, )
perms_database-app-permissions_user_validate_list



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

# **PermsDatabaseAppPermissionsUsersAllList**
> InlineResponse20044 PermsDatabaseAppPermissionsUsersAllList(ctx, id, optional)
perms_database-app-permissions_users_all_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***PermsDatabaseAppPermissionsUsersAllListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsDatabaseAppPermissionsUsersAllListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20044**](inline_response_200_44.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

