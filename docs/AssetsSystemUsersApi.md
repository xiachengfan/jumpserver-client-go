# \AssetsSystemUsersApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsSystemUsersAssetsAuthInfoRead**](AssetsSystemUsersApi.md#AssetsSystemUsersAssetsAuthInfoRead) | **Get** /assets/system-users/{id}/assets/{aid}/auth-info/ | assets_system-users_assets_auth-info_read
[**AssetsSystemUsersAssetsList**](AssetsSystemUsersApi.md#AssetsSystemUsersAssetsList) | **Get** /assets/system-users/{id}/assets/ | assets_system-users_assets_list
[**AssetsSystemUsersAssetsPushRead**](AssetsSystemUsersApi.md#AssetsSystemUsersAssetsPushRead) | **Get** /assets/system-users/{id}/assets/{aid}/push/ | assets_system-users_assets_push_read
[**AssetsSystemUsersAssetsTestRead**](AssetsSystemUsersApi.md#AssetsSystemUsersAssetsTestRead) | **Get** /assets/system-users/{id}/assets/{aid}/test/ | assets_system-users_assets_test_read
[**AssetsSystemUsersAuthInfoDelete**](AssetsSystemUsersApi.md#AssetsSystemUsersAuthInfoDelete) | **Delete** /assets/system-users/{id}/auth-info/ | assets_system-users_auth-info_delete
[**AssetsSystemUsersAuthInfoPartialUpdate**](AssetsSystemUsersApi.md#AssetsSystemUsersAuthInfoPartialUpdate) | **Patch** /assets/system-users/{id}/auth-info/ | assets_system-users_auth-info_partial_update
[**AssetsSystemUsersAuthInfoRead**](AssetsSystemUsersApi.md#AssetsSystemUsersAuthInfoRead) | **Get** /assets/system-users/{id}/auth-info/ | assets_system-users_auth-info_read
[**AssetsSystemUsersAuthInfoUpdate**](AssetsSystemUsersApi.md#AssetsSystemUsersAuthInfoUpdate) | **Put** /assets/system-users/{id}/auth-info/ | assets_system-users_auth-info_update
[**AssetsSystemUsersBulkDelete**](AssetsSystemUsersApi.md#AssetsSystemUsersBulkDelete) | **Delete** /assets/system-users/ | assets_system-users_bulk_delete
[**AssetsSystemUsersBulkUpdate**](AssetsSystemUsersApi.md#AssetsSystemUsersBulkUpdate) | **Put** /assets/system-users/ | assets_system-users_bulk_update
[**AssetsSystemUsersCmdFilterRulesList**](AssetsSystemUsersApi.md#AssetsSystemUsersCmdFilterRulesList) | **Get** /assets/system-users/{id}/cmd-filter-rules/ | assets_system-users_cmd-filter-rules_list
[**AssetsSystemUsersConnectiveRead**](AssetsSystemUsersApi.md#AssetsSystemUsersConnectiveRead) | **Get** /assets/system-users/{id}/connective/ | assets_system-users_connective_read
[**AssetsSystemUsersCreate**](AssetsSystemUsersApi.md#AssetsSystemUsersCreate) | **Post** /assets/system-users/ | assets_system-users_create
[**AssetsSystemUsersDelete**](AssetsSystemUsersApi.md#AssetsSystemUsersDelete) | **Delete** /assets/system-users/{id}/ | assets_system-users_delete
[**AssetsSystemUsersList**](AssetsSystemUsersApi.md#AssetsSystemUsersList) | **Get** /assets/system-users/ | assets_system-users_list
[**AssetsSystemUsersPartialBulkUpdate**](AssetsSystemUsersApi.md#AssetsSystemUsersPartialBulkUpdate) | **Patch** /assets/system-users/ | assets_system-users_partial_bulk_update
[**AssetsSystemUsersPartialUpdate**](AssetsSystemUsersApi.md#AssetsSystemUsersPartialUpdate) | **Patch** /assets/system-users/{id}/ | assets_system-users_partial_update
[**AssetsSystemUsersPushRead**](AssetsSystemUsersApi.md#AssetsSystemUsersPushRead) | **Get** /assets/system-users/{id}/push/ | assets_system-users_push_read
[**AssetsSystemUsersRead**](AssetsSystemUsersApi.md#AssetsSystemUsersRead) | **Get** /assets/system-users/{id}/ | assets_system-users_read
[**AssetsSystemUsersUpdate**](AssetsSystemUsersApi.md#AssetsSystemUsersUpdate) | **Put** /assets/system-users/{id}/ | assets_system-users_update


# **AssetsSystemUsersAssetsAuthInfoRead**
> SystemUserAuth AssetsSystemUsersAssetsAuthInfoRead(ctx, aid, id)
assets_system-users_assets_auth-info_read

Get system user with asset auth info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aid** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**SystemUserAuth**](SystemUserAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersAssetsList**
> InlineResponse2003 AssetsSystemUsersAssetsList(ctx, id, optional)
assets_system-users_assets_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AssetsSystemUsersAssetsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsSystemUsersAssetsListOpts struct

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

# **AssetsSystemUsersAssetsPushRead**
> TaskId AssetsSystemUsersAssetsPushRead(ctx, aid, id)
assets_system-users_assets_push_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aid** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**TaskId**](TaskID.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersAssetsTestRead**
> TaskId AssetsSystemUsersAssetsTestRead(ctx, aid, id)
assets_system-users_assets_test_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aid** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**TaskId**](TaskID.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersAuthInfoDelete**
> AssetsSystemUsersAuthInfoDelete(ctx, id)
assets_system-users_auth-info_delete

Get system user auth info

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

# **AssetsSystemUsersAuthInfoPartialUpdate**
> SystemUserAuth AssetsSystemUsersAuthInfoPartialUpdate(ctx, id, data)
assets_system-users_auth-info_partial_update

Get system user auth info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**SystemUserAuth**](SystemUserAuth.md)|  | 

### Return type

[**SystemUserAuth**](SystemUserAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersAuthInfoRead**
> SystemUserAuth AssetsSystemUsersAuthInfoRead(ctx, id)
assets_system-users_auth-info_read

Get system user auth info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**SystemUserAuth**](SystemUserAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersAuthInfoUpdate**
> SystemUserAuth AssetsSystemUsersAuthInfoUpdate(ctx, id, data)
assets_system-users_auth-info_update

Get system user auth info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**SystemUserAuth**](SystemUserAuth.md)|  | 

### Return type

[**SystemUserAuth**](SystemUserAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersBulkDelete**
> AssetsSystemUsersBulkDelete(ctx, optional)
assets_system-users_bulk_delete

System user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsSystemUsersBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsSystemUsersBulkDeleteOpts struct

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

# **AssetsSystemUsersBulkUpdate**
> SystemUser AssetsSystemUsersBulkUpdate(ctx, data)
assets_system-users_bulk_update

System user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SystemUser**](SystemUser.md)|  | 

### Return type

[**SystemUser**](SystemUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersCmdFilterRulesList**
> InlineResponse2008 AssetsSystemUsersCmdFilterRulesList(ctx, id, optional)
assets_system-users_cmd-filter-rules_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AssetsSystemUsersCmdFilterRulesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsSystemUsersCmdFilterRulesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersConnectiveRead**
> CeleryTask AssetsSystemUsersConnectiveRead(ctx, id)
assets_system-users_connective_read

Push system user to cluster assets api

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**CeleryTask**](CeleryTask.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersCreate**
> SystemUser AssetsSystemUsersCreate(ctx, data)
assets_system-users_create

System user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SystemUser**](SystemUser.md)|  | 

### Return type

[**SystemUser**](SystemUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersDelete**
> AssetsSystemUsersDelete(ctx, id)
assets_system-users_delete

System user api set, for add,delete,update,list,retrieve resource

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

# **AssetsSystemUsersList**
> InlineResponse20019 AssetsSystemUsersList(ctx, optional)
assets_system-users_list

System user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsSystemUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsSystemUsersListOpts struct

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

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersPartialBulkUpdate**
> SystemUser AssetsSystemUsersPartialBulkUpdate(ctx, data)
assets_system-users_partial_bulk_update

System user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SystemUser**](SystemUser.md)|  | 

### Return type

[**SystemUser**](SystemUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersPartialUpdate**
> SystemUser AssetsSystemUsersPartialUpdate(ctx, id, data)
assets_system-users_partial_update

System user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**SystemUser**](SystemUser.md)|  | 

### Return type

[**SystemUser**](SystemUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersPushRead**
> CeleryTask AssetsSystemUsersPushRead(ctx, id)
assets_system-users_push_read

Push system user to cluster assets api

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**CeleryTask**](CeleryTask.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersRead**
> SystemUser AssetsSystemUsersRead(ctx, id)
assets_system-users_read

System user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**SystemUser**](SystemUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersUpdate**
> SystemUser AssetsSystemUsersUpdate(ctx, id, data)
assets_system-users_update

System user api set, for add,delete,update,list,retrieve resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**SystemUser**](SystemUser.md)|  | 

### Return type

[**SystemUser**](SystemUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

