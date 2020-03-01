# \PermsRemoteAppPermissionsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermsRemoteAppPermissionsCreate**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsCreate) | **Post** /perms/remote-app-permissions/ | perms_remote-app-permissions_create
[**PermsRemoteAppPermissionsDelete**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsDelete) | **Delete** /perms/remote-app-permissions/{id}/ | perms_remote-app-permissions_delete
[**PermsRemoteAppPermissionsList**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsList) | **Get** /perms/remote-app-permissions/ | perms_remote-app-permissions_list
[**PermsRemoteAppPermissionsPartialUpdate**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsPartialUpdate) | **Patch** /perms/remote-app-permissions/{id}/ | perms_remote-app-permissions_partial_update
[**PermsRemoteAppPermissionsRead**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsRead) | **Get** /perms/remote-app-permissions/{id}/ | perms_remote-app-permissions_read
[**PermsRemoteAppPermissionsRemoteAppsAddPartialUpdate**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsRemoteAppsAddPartialUpdate) | **Patch** /perms/remote-app-permissions/{id}/remote-apps/add/ | perms_remote-app-permissions_remote-apps_add_partial_update
[**PermsRemoteAppPermissionsRemoteAppsAddRead**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsRemoteAppsAddRead) | **Get** /perms/remote-app-permissions/{id}/remote-apps/add/ | perms_remote-app-permissions_remote-apps_add_read
[**PermsRemoteAppPermissionsRemoteAppsAddUpdate**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsRemoteAppsAddUpdate) | **Put** /perms/remote-app-permissions/{id}/remote-apps/add/ | perms_remote-app-permissions_remote-apps_add_update
[**PermsRemoteAppPermissionsRemoteAppsRemovePartialUpdate**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsRemoteAppsRemovePartialUpdate) | **Patch** /perms/remote-app-permissions/{id}/remote-apps/remove/ | perms_remote-app-permissions_remote-apps_remove_partial_update
[**PermsRemoteAppPermissionsRemoteAppsRemoveRead**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsRemoteAppsRemoveRead) | **Get** /perms/remote-app-permissions/{id}/remote-apps/remove/ | perms_remote-app-permissions_remote-apps_remove_read
[**PermsRemoteAppPermissionsRemoteAppsRemoveUpdate**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsRemoteAppsRemoveUpdate) | **Put** /perms/remote-app-permissions/{id}/remote-apps/remove/ | perms_remote-app-permissions_remote-apps_remove_update
[**PermsRemoteAppPermissionsUpdate**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsUpdate) | **Put** /perms/remote-app-permissions/{id}/ | perms_remote-app-permissions_update
[**PermsRemoteAppPermissionsUserValidateList**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsUserValidateList) | **Get** /perms/remote-app-permissions/user/validate/ | perms_remote-app-permissions_user_validate_list
[**PermsRemoteAppPermissionsUsersAddPartialUpdate**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsUsersAddPartialUpdate) | **Patch** /perms/remote-app-permissions/{id}/users/add/ | perms_remote-app-permissions_users_add_partial_update
[**PermsRemoteAppPermissionsUsersAddRead**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsUsersAddRead) | **Get** /perms/remote-app-permissions/{id}/users/add/ | perms_remote-app-permissions_users_add_read
[**PermsRemoteAppPermissionsUsersAddUpdate**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsUsersAddUpdate) | **Put** /perms/remote-app-permissions/{id}/users/add/ | perms_remote-app-permissions_users_add_update
[**PermsRemoteAppPermissionsUsersRemovePartialUpdate**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsUsersRemovePartialUpdate) | **Patch** /perms/remote-app-permissions/{id}/users/remove/ | perms_remote-app-permissions_users_remove_partial_update
[**PermsRemoteAppPermissionsUsersRemoveRead**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsUsersRemoveRead) | **Get** /perms/remote-app-permissions/{id}/users/remove/ | perms_remote-app-permissions_users_remove_read
[**PermsRemoteAppPermissionsUsersRemoveUpdate**](PermsRemoteAppPermissionsApi.md#PermsRemoteAppPermissionsUsersRemoveUpdate) | **Put** /perms/remote-app-permissions/{id}/users/remove/ | perms_remote-app-permissions_users_remove_update


# **PermsRemoteAppPermissionsCreate**
> RemoteAppPermission PermsRemoteAppPermissionsCreate(ctx, data)
perms_remote-app-permissions_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RemoteAppPermission**](RemoteAppPermission.md)|  | 

### Return type

[**RemoteAppPermission**](RemoteAppPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsDelete**
> PermsRemoteAppPermissionsDelete(ctx, id)
perms_remote-app-permissions_delete



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

# **PermsRemoteAppPermissionsList**
> InlineResponse20045 PermsRemoteAppPermissionsList(ctx, optional)
perms_remote-app-permissions_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermsRemoteAppPermissionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsRemoteAppPermissionsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsPartialUpdate**
> RemoteAppPermission PermsRemoteAppPermissionsPartialUpdate(ctx, id, data)
perms_remote-app-permissions_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**RemoteAppPermission**](RemoteAppPermission.md)|  | 

### Return type

[**RemoteAppPermission**](RemoteAppPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsRead**
> RemoteAppPermission PermsRemoteAppPermissionsRead(ctx, id)
perms_remote-app-permissions_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**RemoteAppPermission**](RemoteAppPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsRemoteAppsAddPartialUpdate**
> RemoteAppPermissionUpdateRemoteApp PermsRemoteAppPermissionsRemoteAppsAddPartialUpdate(ctx, id, data)
perms_remote-app-permissions_remote-apps_add_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**RemoteAppPermissionUpdateRemoteApp**](RemoteAppPermissionUpdateRemoteApp.md)|  | 

### Return type

[**RemoteAppPermissionUpdateRemoteApp**](RemoteAppPermissionUpdateRemoteApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsRemoteAppsAddRead**
> RemoteAppPermissionUpdateRemoteApp PermsRemoteAppPermissionsRemoteAppsAddRead(ctx, id)
perms_remote-app-permissions_remote-apps_add_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**RemoteAppPermissionUpdateRemoteApp**](RemoteAppPermissionUpdateRemoteApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsRemoteAppsAddUpdate**
> RemoteAppPermissionUpdateRemoteApp PermsRemoteAppPermissionsRemoteAppsAddUpdate(ctx, id, data)
perms_remote-app-permissions_remote-apps_add_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**RemoteAppPermissionUpdateRemoteApp**](RemoteAppPermissionUpdateRemoteApp.md)|  | 

### Return type

[**RemoteAppPermissionUpdateRemoteApp**](RemoteAppPermissionUpdateRemoteApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsRemoteAppsRemovePartialUpdate**
> RemoteAppPermissionUpdateRemoteApp PermsRemoteAppPermissionsRemoteAppsRemovePartialUpdate(ctx, id, data)
perms_remote-app-permissions_remote-apps_remove_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**RemoteAppPermissionUpdateRemoteApp**](RemoteAppPermissionUpdateRemoteApp.md)|  | 

### Return type

[**RemoteAppPermissionUpdateRemoteApp**](RemoteAppPermissionUpdateRemoteApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsRemoteAppsRemoveRead**
> RemoteAppPermissionUpdateRemoteApp PermsRemoteAppPermissionsRemoteAppsRemoveRead(ctx, id)
perms_remote-app-permissions_remote-apps_remove_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**RemoteAppPermissionUpdateRemoteApp**](RemoteAppPermissionUpdateRemoteApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsRemoteAppsRemoveUpdate**
> RemoteAppPermissionUpdateRemoteApp PermsRemoteAppPermissionsRemoteAppsRemoveUpdate(ctx, id, data)
perms_remote-app-permissions_remote-apps_remove_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**RemoteAppPermissionUpdateRemoteApp**](RemoteAppPermissionUpdateRemoteApp.md)|  | 

### Return type

[**RemoteAppPermissionUpdateRemoteApp**](RemoteAppPermissionUpdateRemoteApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsUpdate**
> RemoteAppPermission PermsRemoteAppPermissionsUpdate(ctx, id, data)
perms_remote-app-permissions_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**RemoteAppPermission**](RemoteAppPermission.md)|  | 

### Return type

[**RemoteAppPermission**](RemoteAppPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsUserValidateList**
> PermsRemoteAppPermissionsUserValidateList(ctx, )
perms_remote-app-permissions_user_validate_list



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

# **PermsRemoteAppPermissionsUsersAddPartialUpdate**
> RemoteAppPermissionUpdateUser PermsRemoteAppPermissionsUsersAddPartialUpdate(ctx, id, data)
perms_remote-app-permissions_users_add_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**RemoteAppPermissionUpdateUser**](RemoteAppPermissionUpdateUser.md)|  | 

### Return type

[**RemoteAppPermissionUpdateUser**](RemoteAppPermissionUpdateUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsUsersAddRead**
> RemoteAppPermissionUpdateUser PermsRemoteAppPermissionsUsersAddRead(ctx, id)
perms_remote-app-permissions_users_add_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**RemoteAppPermissionUpdateUser**](RemoteAppPermissionUpdateUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsUsersAddUpdate**
> RemoteAppPermissionUpdateUser PermsRemoteAppPermissionsUsersAddUpdate(ctx, id, data)
perms_remote-app-permissions_users_add_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**RemoteAppPermissionUpdateUser**](RemoteAppPermissionUpdateUser.md)|  | 

### Return type

[**RemoteAppPermissionUpdateUser**](RemoteAppPermissionUpdateUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsUsersRemovePartialUpdate**
> RemoteAppPermissionUpdateUser PermsRemoteAppPermissionsUsersRemovePartialUpdate(ctx, id, data)
perms_remote-app-permissions_users_remove_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**RemoteAppPermissionUpdateUser**](RemoteAppPermissionUpdateUser.md)|  | 

### Return type

[**RemoteAppPermissionUpdateUser**](RemoteAppPermissionUpdateUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsUsersRemoveRead**
> RemoteAppPermissionUpdateUser PermsRemoteAppPermissionsUsersRemoveRead(ctx, id)
perms_remote-app-permissions_users_remove_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**RemoteAppPermissionUpdateUser**](RemoteAppPermissionUpdateUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsRemoteAppPermissionsUsersRemoveUpdate**
> RemoteAppPermissionUpdateUser PermsRemoteAppPermissionsUsersRemoveUpdate(ctx, id, data)
perms_remote-app-permissions_users_remove_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**RemoteAppPermissionUpdateUser**](RemoteAppPermissionUpdateUser.md)|  | 

### Return type

[**RemoteAppPermissionUpdateUser**](RemoteAppPermissionUpdateUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

