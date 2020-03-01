# \ApplicationsDatabaseAppsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsDatabaseAppsBulkDelete**](ApplicationsDatabaseAppsApi.md#ApplicationsDatabaseAppsBulkDelete) | **Delete** /applications/database-apps/ | applications_database-apps_bulk_delete
[**ApplicationsDatabaseAppsBulkUpdate**](ApplicationsDatabaseAppsApi.md#ApplicationsDatabaseAppsBulkUpdate) | **Put** /applications/database-apps/ | applications_database-apps_bulk_update
[**ApplicationsDatabaseAppsCreate**](ApplicationsDatabaseAppsApi.md#ApplicationsDatabaseAppsCreate) | **Post** /applications/database-apps/ | applications_database-apps_create
[**ApplicationsDatabaseAppsDelete**](ApplicationsDatabaseAppsApi.md#ApplicationsDatabaseAppsDelete) | **Delete** /applications/database-apps/{id}/ | applications_database-apps_delete
[**ApplicationsDatabaseAppsList**](ApplicationsDatabaseAppsApi.md#ApplicationsDatabaseAppsList) | **Get** /applications/database-apps/ | applications_database-apps_list
[**ApplicationsDatabaseAppsPartialBulkUpdate**](ApplicationsDatabaseAppsApi.md#ApplicationsDatabaseAppsPartialBulkUpdate) | **Patch** /applications/database-apps/ | applications_database-apps_partial_bulk_update
[**ApplicationsDatabaseAppsPartialUpdate**](ApplicationsDatabaseAppsApi.md#ApplicationsDatabaseAppsPartialUpdate) | **Patch** /applications/database-apps/{id}/ | applications_database-apps_partial_update
[**ApplicationsDatabaseAppsRead**](ApplicationsDatabaseAppsApi.md#ApplicationsDatabaseAppsRead) | **Get** /applications/database-apps/{id}/ | applications_database-apps_read
[**ApplicationsDatabaseAppsUpdate**](ApplicationsDatabaseAppsApi.md#ApplicationsDatabaseAppsUpdate) | **Put** /applications/database-apps/{id}/ | applications_database-apps_update


# **ApplicationsDatabaseAppsBulkDelete**
> ApplicationsDatabaseAppsBulkDelete(ctx, optional)
applications_database-apps_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationsDatabaseAppsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsDatabaseAppsBulkDeleteOpts struct

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

# **ApplicationsDatabaseAppsBulkUpdate**
> DatabaseApp ApplicationsDatabaseAppsBulkUpdate(ctx, data)
applications_database-apps_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseApp**](DatabaseApp.md)|  | 

### Return type

[**DatabaseApp**](DatabaseApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsDatabaseAppsCreate**
> DatabaseApp ApplicationsDatabaseAppsCreate(ctx, data)
applications_database-apps_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseApp**](DatabaseApp.md)|  | 

### Return type

[**DatabaseApp**](DatabaseApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsDatabaseAppsDelete**
> ApplicationsDatabaseAppsDelete(ctx, id)
applications_database-apps_delete



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

# **ApplicationsDatabaseAppsList**
> InlineResponse200 ApplicationsDatabaseAppsList(ctx, optional)
applications_database-apps_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationsDatabaseAppsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsDatabaseAppsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsDatabaseAppsPartialBulkUpdate**
> DatabaseApp ApplicationsDatabaseAppsPartialBulkUpdate(ctx, data)
applications_database-apps_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DatabaseApp**](DatabaseApp.md)|  | 

### Return type

[**DatabaseApp**](DatabaseApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsDatabaseAppsPartialUpdate**
> DatabaseApp ApplicationsDatabaseAppsPartialUpdate(ctx, id, data)
applications_database-apps_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**DatabaseApp**](DatabaseApp.md)|  | 

### Return type

[**DatabaseApp**](DatabaseApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsDatabaseAppsRead**
> DatabaseApp ApplicationsDatabaseAppsRead(ctx, id)
applications_database-apps_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DatabaseApp**](DatabaseApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsDatabaseAppsUpdate**
> DatabaseApp ApplicationsDatabaseAppsUpdate(ctx, id, data)
applications_database-apps_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**DatabaseApp**](DatabaseApp.md)|  | 

### Return type

[**DatabaseApp**](DatabaseApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

