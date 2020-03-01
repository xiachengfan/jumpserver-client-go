# \TerminalTerminalsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TerminalTerminalsAccessKeyList**](TerminalTerminalsApi.md#TerminalTerminalsAccessKeyList) | **Get** /terminal/terminals/{terminal}/access-key/ | terminal_terminals_access-key_list
[**TerminalTerminalsConfigList**](TerminalTerminalsApi.md#TerminalTerminalsConfigList) | **Get** /terminal/terminals/config/ | terminal_terminals_config_list
[**TerminalTerminalsCreate**](TerminalTerminalsApi.md#TerminalTerminalsCreate) | **Post** /terminal/terminals/ | terminal_terminals_create
[**TerminalTerminalsDelete**](TerminalTerminalsApi.md#TerminalTerminalsDelete) | **Delete** /terminal/terminals/{id}/ | terminal_terminals_delete
[**TerminalTerminalsList**](TerminalTerminalsApi.md#TerminalTerminalsList) | **Get** /terminal/terminals/ | terminal_terminals_list
[**TerminalTerminalsPartialUpdate**](TerminalTerminalsApi.md#TerminalTerminalsPartialUpdate) | **Patch** /terminal/terminals/{id}/ | terminal_terminals_partial_update
[**TerminalTerminalsRead**](TerminalTerminalsApi.md#TerminalTerminalsRead) | **Get** /terminal/terminals/{id}/ | terminal_terminals_read
[**TerminalTerminalsSessionsBulkDelete**](TerminalTerminalsApi.md#TerminalTerminalsSessionsBulkDelete) | **Delete** /terminal/terminals/{terminal}/sessions/ | terminal_terminals_sessions_bulk_delete
[**TerminalTerminalsSessionsBulkUpdate**](TerminalTerminalsApi.md#TerminalTerminalsSessionsBulkUpdate) | **Put** /terminal/terminals/{terminal}/sessions/ | terminal_terminals_sessions_bulk_update
[**TerminalTerminalsSessionsCreate**](TerminalTerminalsApi.md#TerminalTerminalsSessionsCreate) | **Post** /terminal/terminals/{terminal}/sessions/ | terminal_terminals_sessions_create
[**TerminalTerminalsSessionsDelete**](TerminalTerminalsApi.md#TerminalTerminalsSessionsDelete) | **Delete** /terminal/terminals/{terminal}/sessions/{id}/ | terminal_terminals_sessions_delete
[**TerminalTerminalsSessionsList**](TerminalTerminalsApi.md#TerminalTerminalsSessionsList) | **Get** /terminal/terminals/{terminal}/sessions/ | terminal_terminals_sessions_list
[**TerminalTerminalsSessionsPartialBulkUpdate**](TerminalTerminalsApi.md#TerminalTerminalsSessionsPartialBulkUpdate) | **Patch** /terminal/terminals/{terminal}/sessions/ | terminal_terminals_sessions_partial_bulk_update
[**TerminalTerminalsSessionsPartialUpdate**](TerminalTerminalsApi.md#TerminalTerminalsSessionsPartialUpdate) | **Patch** /terminal/terminals/{terminal}/sessions/{id}/ | terminal_terminals_sessions_partial_update
[**TerminalTerminalsSessionsRead**](TerminalTerminalsApi.md#TerminalTerminalsSessionsRead) | **Get** /terminal/terminals/{terminal}/sessions/{id}/ | terminal_terminals_sessions_read
[**TerminalTerminalsSessionsUpdate**](TerminalTerminalsApi.md#TerminalTerminalsSessionsUpdate) | **Put** /terminal/terminals/{terminal}/sessions/{id}/ | terminal_terminals_sessions_update
[**TerminalTerminalsStatusCreate**](TerminalTerminalsApi.md#TerminalTerminalsStatusCreate) | **Post** /terminal/terminals/{terminal}/status/ | terminal_terminals_status_create
[**TerminalTerminalsStatusDelete**](TerminalTerminalsApi.md#TerminalTerminalsStatusDelete) | **Delete** /terminal/terminals/{terminal}/status/{id}/ | terminal_terminals_status_delete
[**TerminalTerminalsStatusList**](TerminalTerminalsApi.md#TerminalTerminalsStatusList) | **Get** /terminal/terminals/{terminal}/status/ | terminal_terminals_status_list
[**TerminalTerminalsStatusPartialUpdate**](TerminalTerminalsApi.md#TerminalTerminalsStatusPartialUpdate) | **Patch** /terminal/terminals/{terminal}/status/{id}/ | terminal_terminals_status_partial_update
[**TerminalTerminalsStatusRead**](TerminalTerminalsApi.md#TerminalTerminalsStatusRead) | **Get** /terminal/terminals/{terminal}/status/{id}/ | terminal_terminals_status_read
[**TerminalTerminalsStatusUpdate**](TerminalTerminalsApi.md#TerminalTerminalsStatusUpdate) | **Put** /terminal/terminals/{terminal}/status/{id}/ | terminal_terminals_status_update
[**TerminalTerminalsUpdate**](TerminalTerminalsApi.md#TerminalTerminalsUpdate) | **Put** /terminal/terminals/{id}/ | terminal_terminals_update


# **TerminalTerminalsAccessKeyList**
> TerminalTerminalsAccessKeyList(ctx, terminal)
terminal_terminals_access-key_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **terminal** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsConfigList**
> TerminalTerminalsConfigList(ctx, )
terminal_terminals_config_list



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

# **TerminalTerminalsCreate**
> Terminal TerminalTerminalsCreate(ctx, data)
terminal_terminals_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Terminal**](Terminal.md)|  | 

### Return type

[**Terminal**](Terminal.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsDelete**
> TerminalTerminalsDelete(ctx, id)
terminal_terminals_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this terminal. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsList**
> InlineResponse20058 TerminalTerminalsList(ctx, optional)
terminal_terminals_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TerminalTerminalsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerminalTerminalsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20058**](inline_response_200_58.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsPartialUpdate**
> Terminal TerminalTerminalsPartialUpdate(ctx, id, data)
terminal_terminals_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this terminal. | 
  **data** | [**Terminal**](Terminal.md)|  | 

### Return type

[**Terminal**](Terminal.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsRead**
> Terminal TerminalTerminalsRead(ctx, id)
terminal_terminals_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this terminal. | 

### Return type

[**Terminal**](Terminal.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsSessionsBulkDelete**
> TerminalTerminalsSessionsBulkDelete(ctx, terminal, optional)
terminal_terminals_sessions_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **terminal** | **string**|  | 
 **optional** | ***TerminalTerminalsSessionsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerminalTerminalsSessionsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | **optional.String**|  | 
 **asset** | **optional.String**|  | 
 **systemUser** | **optional.String**|  | 
 **remoteAddr** | **optional.String**|  | 
 **protocol** | **optional.String**|  | 
 **terminal2** | **optional.String**|  | 
 **isFinished** | **optional.String**|  | 
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

# **TerminalTerminalsSessionsBulkUpdate**
> Session TerminalTerminalsSessionsBulkUpdate(ctx, terminal, data)
terminal_terminals_sessions_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **terminal** | **string**|  | 
  **data** | [**Session**](Session.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsSessionsCreate**
> Session TerminalTerminalsSessionsCreate(ctx, terminal, data)
terminal_terminals_sessions_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **terminal** | **string**|  | 
  **data** | [**Session**](Session.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsSessionsDelete**
> TerminalTerminalsSessionsDelete(ctx, id, terminal)
terminal_terminals_sessions_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **terminal** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsSessionsList**
> InlineResponse20056 TerminalTerminalsSessionsList(ctx, terminal, optional)
terminal_terminals_sessions_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **terminal** | **string**|  | 
 **optional** | ***TerminalTerminalsSessionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerminalTerminalsSessionsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | **optional.String**|  | 
 **asset** | **optional.String**|  | 
 **systemUser** | **optional.String**|  | 
 **remoteAddr** | **optional.String**|  | 
 **protocol** | **optional.String**|  | 
 **terminal2** | **optional.String**|  | 
 **isFinished** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20056**](inline_response_200_56.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsSessionsPartialBulkUpdate**
> Session TerminalTerminalsSessionsPartialBulkUpdate(ctx, terminal, data)
terminal_terminals_sessions_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **terminal** | **string**|  | 
  **data** | [**Session**](Session.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsSessionsPartialUpdate**
> Session TerminalTerminalsSessionsPartialUpdate(ctx, id, terminal, data)
terminal_terminals_sessions_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **terminal** | **string**|  | 
  **data** | [**Session**](Session.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsSessionsRead**
> Session TerminalTerminalsSessionsRead(ctx, id, terminal)
terminal_terminals_sessions_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **terminal** | **string**|  | 

### Return type

[**Session**](Session.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsSessionsUpdate**
> Session TerminalTerminalsSessionsUpdate(ctx, id, terminal, data)
terminal_terminals_sessions_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **terminal** | **string**|  | 
  **data** | [**Session**](Session.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsStatusCreate**
> Status TerminalTerminalsStatusCreate(ctx, terminal, data)
terminal_terminals_status_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **terminal** | **string**|  | 
  **data** | [**Status**](Status.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsStatusDelete**
> TerminalTerminalsStatusDelete(ctx, id, terminal)
terminal_terminals_status_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this status. | 
  **terminal** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsStatusList**
> InlineResponse20057 TerminalTerminalsStatusList(ctx, terminal, optional)
terminal_terminals_status_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **terminal** | **string**|  | 
 **optional** | ***TerminalTerminalsStatusListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerminalTerminalsStatusListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20057**](inline_response_200_57.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsStatusPartialUpdate**
> Status TerminalTerminalsStatusPartialUpdate(ctx, id, terminal, data)
terminal_terminals_status_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this status. | 
  **terminal** | **string**|  | 
  **data** | [**Status**](Status.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsStatusRead**
> Status TerminalTerminalsStatusRead(ctx, id, terminal)
terminal_terminals_status_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this status. | 
  **terminal** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsStatusUpdate**
> Status TerminalTerminalsStatusUpdate(ctx, id, terminal, data)
terminal_terminals_status_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this status. | 
  **terminal** | **string**|  | 
  **data** | [**Status**](Status.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalTerminalsUpdate**
> Terminal TerminalTerminalsUpdate(ctx, id, data)
terminal_terminals_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this terminal. | 
  **data** | [**Terminal**](Terminal.md)|  | 

### Return type

[**Terminal**](Terminal.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

