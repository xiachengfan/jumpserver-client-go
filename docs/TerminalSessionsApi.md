# \TerminalSessionsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TerminalSessionsBulkDelete**](TerminalSessionsApi.md#TerminalSessionsBulkDelete) | **Delete** /terminal/sessions/ | terminal_sessions_bulk_delete
[**TerminalSessionsBulkUpdate**](TerminalSessionsApi.md#TerminalSessionsBulkUpdate) | **Put** /terminal/sessions/ | terminal_sessions_bulk_update
[**TerminalSessionsCreate**](TerminalSessionsApi.md#TerminalSessionsCreate) | **Post** /terminal/sessions/ | terminal_sessions_create
[**TerminalSessionsDelete**](TerminalSessionsApi.md#TerminalSessionsDelete) | **Delete** /terminal/sessions/{id}/ | terminal_sessions_delete
[**TerminalSessionsList**](TerminalSessionsApi.md#TerminalSessionsList) | **Get** /terminal/sessions/ | terminal_sessions_list
[**TerminalSessionsPartialBulkUpdate**](TerminalSessionsApi.md#TerminalSessionsPartialBulkUpdate) | **Patch** /terminal/sessions/ | terminal_sessions_partial_bulk_update
[**TerminalSessionsPartialUpdate**](TerminalSessionsApi.md#TerminalSessionsPartialUpdate) | **Patch** /terminal/sessions/{id}/ | terminal_sessions_partial_update
[**TerminalSessionsRead**](TerminalSessionsApi.md#TerminalSessionsRead) | **Get** /terminal/sessions/{id}/ | terminal_sessions_read
[**TerminalSessionsReplayCreate**](TerminalSessionsApi.md#TerminalSessionsReplayCreate) | **Post** /terminal/sessions/{id}/replay/ | terminal_sessions_replay_create
[**TerminalSessionsReplayRead**](TerminalSessionsApi.md#TerminalSessionsReplayRead) | **Get** /terminal/sessions/{id}/replay/ | terminal_sessions_replay_read
[**TerminalSessionsUpdate**](TerminalSessionsApi.md#TerminalSessionsUpdate) | **Put** /terminal/sessions/{id}/ | terminal_sessions_update


# **TerminalSessionsBulkDelete**
> TerminalSessionsBulkDelete(ctx, optional)
terminal_sessions_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TerminalSessionsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerminalSessionsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **optional.String**|  | 
 **asset** | **optional.String**|  | 
 **systemUser** | **optional.String**|  | 
 **remoteAddr** | **optional.String**|  | 
 **protocol** | **optional.String**|  | 
 **terminal** | **optional.String**|  | 
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

# **TerminalSessionsBulkUpdate**
> Session TerminalSessionsBulkUpdate(ctx, data)
terminal_sessions_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Session**](Session.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalSessionsCreate**
> Session TerminalSessionsCreate(ctx, data)
terminal_sessions_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Session**](Session.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalSessionsDelete**
> TerminalSessionsDelete(ctx, id)
terminal_sessions_delete



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

# **TerminalSessionsList**
> InlineResponse20056 TerminalSessionsList(ctx, optional)
terminal_sessions_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TerminalSessionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerminalSessionsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **optional.String**|  | 
 **asset** | **optional.String**|  | 
 **systemUser** | **optional.String**|  | 
 **remoteAddr** | **optional.String**|  | 
 **protocol** | **optional.String**|  | 
 **terminal** | **optional.String**|  | 
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

# **TerminalSessionsPartialBulkUpdate**
> Session TerminalSessionsPartialBulkUpdate(ctx, data)
terminal_sessions_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Session**](Session.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalSessionsPartialUpdate**
> Session TerminalSessionsPartialUpdate(ctx, id, data)
terminal_sessions_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**Session**](Session.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalSessionsRead**
> Session TerminalSessionsRead(ctx, id)
terminal_sessions_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Session**](Session.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalSessionsReplayCreate**
> TerminalSessionsReplayCreate(ctx, id)
terminal_sessions_replay_create



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

# **TerminalSessionsReplayRead**
> TerminalSessionsReplayRead(ctx, id)
terminal_sessions_replay_read



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

# **TerminalSessionsUpdate**
> Session TerminalSessionsUpdate(ctx, id, data)
terminal_sessions_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**Session**](Session.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

