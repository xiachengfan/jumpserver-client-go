# \AssetsAdminUsersApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAdminUsersAssetsList**](AssetsAdminUsersApi.md#AssetsAdminUsersAssetsList) | **Get** /assets/admin-users/{id}/assets/ | assets_admin-users_assets_list
[**AssetsAdminUsersAuthPartialUpdate**](AssetsAdminUsersApi.md#AssetsAdminUsersAuthPartialUpdate) | **Patch** /assets/admin-users/{id}/auth/ | assets_admin-users_auth_partial_update
[**AssetsAdminUsersAuthUpdate**](AssetsAdminUsersApi.md#AssetsAdminUsersAuthUpdate) | **Put** /assets/admin-users/{id}/auth/ | assets_admin-users_auth_update
[**AssetsAdminUsersBulkDelete**](AssetsAdminUsersApi.md#AssetsAdminUsersBulkDelete) | **Delete** /assets/admin-users/ | assets_admin-users_bulk_delete
[**AssetsAdminUsersBulkUpdate**](AssetsAdminUsersApi.md#AssetsAdminUsersBulkUpdate) | **Put** /assets/admin-users/ | assets_admin-users_bulk_update
[**AssetsAdminUsersConnectiveRead**](AssetsAdminUsersApi.md#AssetsAdminUsersConnectiveRead) | **Get** /assets/admin-users/{id}/connective/ | assets_admin-users_connective_read
[**AssetsAdminUsersCreate**](AssetsAdminUsersApi.md#AssetsAdminUsersCreate) | **Post** /assets/admin-users/ | assets_admin-users_create
[**AssetsAdminUsersDelete**](AssetsAdminUsersApi.md#AssetsAdminUsersDelete) | **Delete** /assets/admin-users/{id}/ | assets_admin-users_delete
[**AssetsAdminUsersList**](AssetsAdminUsersApi.md#AssetsAdminUsersList) | **Get** /assets/admin-users/ | assets_admin-users_list
[**AssetsAdminUsersNodesPartialUpdate**](AssetsAdminUsersApi.md#AssetsAdminUsersNodesPartialUpdate) | **Patch** /assets/admin-users/{id}/nodes/ | assets_admin-users_nodes_partial_update
[**AssetsAdminUsersNodesUpdate**](AssetsAdminUsersApi.md#AssetsAdminUsersNodesUpdate) | **Put** /assets/admin-users/{id}/nodes/ | assets_admin-users_nodes_update
[**AssetsAdminUsersPartialBulkUpdate**](AssetsAdminUsersApi.md#AssetsAdminUsersPartialBulkUpdate) | **Patch** /assets/admin-users/ | assets_admin-users_partial_bulk_update
[**AssetsAdminUsersPartialUpdate**](AssetsAdminUsersApi.md#AssetsAdminUsersPartialUpdate) | **Patch** /assets/admin-users/{id}/ | assets_admin-users_partial_update
[**AssetsAdminUsersRead**](AssetsAdminUsersApi.md#AssetsAdminUsersRead) | **Get** /assets/admin-users/{id}/ | assets_admin-users_read
[**AssetsAdminUsersUpdate**](AssetsAdminUsersApi.md#AssetsAdminUsersUpdate) | **Put** /assets/admin-users/{id}/ | assets_admin-users_update


# **AssetsAdminUsersAssetsList**
> InlineResponse2003 AssetsAdminUsersAssetsList(ctx, id, optional)
assets_admin-users_assets_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AssetsAdminUsersAssetsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsAdminUsersAssetsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAdminUsersAuthPartialUpdate**
> AdminUserAuth AssetsAdminUsersAuthPartialUpdate(ctx, id, data)
assets_admin-users_auth_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AdminUserAuth**](AdminUserAuth.md)|  | 

### Return type

[**AdminUserAuth**](AdminUserAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAdminUsersAuthUpdate**
> AdminUserAuth AssetsAdminUsersAuthUpdate(ctx, id, data)
assets_admin-users_auth_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AdminUserAuth**](AdminUserAuth.md)|  | 

### Return type

[**AdminUserAuth**](AdminUserAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAdminUsersBulkDelete**
> AssetsAdminUsersBulkDelete(ctx, optional)
assets_admin-users_bulk_delete

Admin user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsAdminUsersBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsAdminUsersBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **username** | **optional.String**|  | 
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

# **AssetsAdminUsersBulkUpdate**
> AdminUser AssetsAdminUsersBulkUpdate(ctx, data)
assets_admin-users_bulk_update

Admin user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdminUser**](AdminUser.md)|  | 

### Return type

[**AdminUser**](AdminUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAdminUsersConnectiveRead**
> TaskId AssetsAdminUsersConnectiveRead(ctx, id)
assets_admin-users_connective_read

Test asset admin user assets_connectivity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**TaskId**](TaskID.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAdminUsersCreate**
> AdminUser AssetsAdminUsersCreate(ctx, data)
assets_admin-users_create

Admin user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdminUser**](AdminUser.md)|  | 

### Return type

[**AdminUser**](AdminUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAdminUsersDelete**
> AssetsAdminUsersDelete(ctx, id)
assets_admin-users_delete

Admin user api set, for add,delete,update,list,retrieve resource

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

# **AssetsAdminUsersList**
> InlineResponse2002 AssetsAdminUsersList(ctx, optional)
assets_admin-users_list

Admin user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsAdminUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsAdminUsersListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **username** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAdminUsersNodesPartialUpdate**
> ReplaceNodeAdminUser AssetsAdminUsersNodesPartialUpdate(ctx, id, data)
assets_admin-users_nodes_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**ReplaceNodeAdminUser**](ReplaceNodeAdminUser.md)|  | 

### Return type

[**ReplaceNodeAdminUser**](ReplaceNodeAdminUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAdminUsersNodesUpdate**
> ReplaceNodeAdminUser AssetsAdminUsersNodesUpdate(ctx, id, data)
assets_admin-users_nodes_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**ReplaceNodeAdminUser**](ReplaceNodeAdminUser.md)|  | 

### Return type

[**ReplaceNodeAdminUser**](ReplaceNodeAdminUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAdminUsersPartialBulkUpdate**
> AdminUser AssetsAdminUsersPartialBulkUpdate(ctx, data)
assets_admin-users_partial_bulk_update

Admin user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdminUser**](AdminUser.md)|  | 

### Return type

[**AdminUser**](AdminUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAdminUsersPartialUpdate**
> AdminUser AssetsAdminUsersPartialUpdate(ctx, id, data)
assets_admin-users_partial_update

Admin user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AdminUser**](AdminUser.md)|  | 

### Return type

[**AdminUser**](AdminUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAdminUsersRead**
> AdminUser AssetsAdminUsersRead(ctx, id)
assets_admin-users_read

Admin user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AdminUser**](AdminUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsAdminUsersUpdate**
> AdminUser AssetsAdminUsersUpdate(ctx, id, data)
assets_admin-users_update

Admin user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AdminUser**](AdminUser.md)|  | 

### Return type

[**AdminUser**](AdminUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

