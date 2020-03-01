# \ApplicationsRemoteAppsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsRemoteAppsBulkDelete**](ApplicationsRemoteAppsApi.md#ApplicationsRemoteAppsBulkDelete) | **Delete** /applications/remote-apps/ | applications_remote-apps_bulk_delete
[**ApplicationsRemoteAppsBulkUpdate**](ApplicationsRemoteAppsApi.md#ApplicationsRemoteAppsBulkUpdate) | **Put** /applications/remote-apps/ | applications_remote-apps_bulk_update
[**ApplicationsRemoteAppsConnectionInfoRead**](ApplicationsRemoteAppsApi.md#ApplicationsRemoteAppsConnectionInfoRead) | **Get** /applications/remote-apps/{id}/connection-info/ | applications_remote-apps_connection-info_read
[**ApplicationsRemoteAppsCreate**](ApplicationsRemoteAppsApi.md#ApplicationsRemoteAppsCreate) | **Post** /applications/remote-apps/ | applications_remote-apps_create
[**ApplicationsRemoteAppsDelete**](ApplicationsRemoteAppsApi.md#ApplicationsRemoteAppsDelete) | **Delete** /applications/remote-apps/{id}/ | applications_remote-apps_delete
[**ApplicationsRemoteAppsList**](ApplicationsRemoteAppsApi.md#ApplicationsRemoteAppsList) | **Get** /applications/remote-apps/ | applications_remote-apps_list
[**ApplicationsRemoteAppsPartialBulkUpdate**](ApplicationsRemoteAppsApi.md#ApplicationsRemoteAppsPartialBulkUpdate) | **Patch** /applications/remote-apps/ | applications_remote-apps_partial_bulk_update
[**ApplicationsRemoteAppsPartialUpdate**](ApplicationsRemoteAppsApi.md#ApplicationsRemoteAppsPartialUpdate) | **Patch** /applications/remote-apps/{id}/ | applications_remote-apps_partial_update
[**ApplicationsRemoteAppsRead**](ApplicationsRemoteAppsApi.md#ApplicationsRemoteAppsRead) | **Get** /applications/remote-apps/{id}/ | applications_remote-apps_read
[**ApplicationsRemoteAppsUpdate**](ApplicationsRemoteAppsApi.md#ApplicationsRemoteAppsUpdate) | **Put** /applications/remote-apps/{id}/ | applications_remote-apps_update


# **ApplicationsRemoteAppsBulkDelete**
> ApplicationsRemoteAppsBulkDelete(ctx, optional)
applications_remote-apps_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationsRemoteAppsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsRemoteAppsBulkDeleteOpts struct

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

# **ApplicationsRemoteAppsBulkUpdate**
> RemoteApp ApplicationsRemoteAppsBulkUpdate(ctx, data)
applications_remote-apps_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RemoteApp**](RemoteApp.md)|  | 

### Return type

[**RemoteApp**](RemoteApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsRemoteAppsConnectionInfoRead**
> RemoteAppConnectionInfo ApplicationsRemoteAppsConnectionInfoRead(ctx, id)
applications_remote-apps_connection-info_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**RemoteAppConnectionInfo**](RemoteAppConnectionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsRemoteAppsCreate**
> RemoteApp ApplicationsRemoteAppsCreate(ctx, data)
applications_remote-apps_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RemoteApp**](RemoteApp.md)|  | 

### Return type

[**RemoteApp**](RemoteApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsRemoteAppsDelete**
> ApplicationsRemoteAppsDelete(ctx, id)
applications_remote-apps_delete



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

# **ApplicationsRemoteAppsList**
> InlineResponse2001 ApplicationsRemoteAppsList(ctx, optional)
applications_remote-apps_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationsRemoteAppsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsRemoteAppsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsRemoteAppsPartialBulkUpdate**
> RemoteApp ApplicationsRemoteAppsPartialBulkUpdate(ctx, data)
applications_remote-apps_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RemoteApp**](RemoteApp.md)|  | 

### Return type

[**RemoteApp**](RemoteApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsRemoteAppsPartialUpdate**
> RemoteApp ApplicationsRemoteAppsPartialUpdate(ctx, id, data)
applications_remote-apps_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**RemoteApp**](RemoteApp.md)|  | 

### Return type

[**RemoteApp**](RemoteApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsRemoteAppsRead**
> RemoteApp ApplicationsRemoteAppsRead(ctx, id)
applications_remote-apps_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**RemoteApp**](RemoteApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsRemoteAppsUpdate**
> RemoteApp ApplicationsRemoteAppsUpdate(ctx, id, data)
applications_remote-apps_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**RemoteApp**](RemoteApp.md)|  | 

### Return type

[**RemoteApp**](RemoteApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

