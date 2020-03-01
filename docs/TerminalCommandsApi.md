# \TerminalCommandsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TerminalCommandsCreate**](TerminalCommandsApi.md#TerminalCommandsCreate) | **Post** /terminal/commands/ | terminal_commands_create
[**TerminalCommandsDelete**](TerminalCommandsApi.md#TerminalCommandsDelete) | **Delete** /terminal/commands/{id}/ | terminal_commands_delete
[**TerminalCommandsExportList**](TerminalCommandsApi.md#TerminalCommandsExportList) | **Get** /terminal/commands/export/ | terminal_commands_export_list
[**TerminalCommandsList**](TerminalCommandsApi.md#TerminalCommandsList) | **Get** /terminal/commands/ | terminal_commands_list
[**TerminalCommandsPartialUpdate**](TerminalCommandsApi.md#TerminalCommandsPartialUpdate) | **Patch** /terminal/commands/{id}/ | terminal_commands_partial_update
[**TerminalCommandsRead**](TerminalCommandsApi.md#TerminalCommandsRead) | **Get** /terminal/commands/{id}/ | terminal_commands_read
[**TerminalCommandsUpdate**](TerminalCommandsApi.md#TerminalCommandsUpdate) | **Put** /terminal/commands/{id}/ | terminal_commands_update


# **TerminalCommandsCreate**
> SessionCommand TerminalCommandsCreate(ctx, data)
terminal_commands_create

接受app发送来的command log, 格式如下 {     \"user\": \"admin\",     \"asset\": \"localhost\",     \"system_user\": \"web\",     \"session\": \"xxxxxx\",     \"input\": \"whoami\",     \"output\": \"d2hvbWFp\",  # base64.b64encode(s)     \"timestamp\": 1485238673.0 }

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SessionCommand**](SessionCommand.md)|  | 

### Return type

[**SessionCommand**](SessionCommand.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalCommandsDelete**
> TerminalCommandsDelete(ctx, id)
terminal_commands_delete

接受app发送来的command log, 格式如下 {     \"user\": \"admin\",     \"asset\": \"localhost\",     \"system_user\": \"web\",     \"session\": \"xxxxxx\",     \"input\": \"whoami\",     \"output\": \"d2hvbWFp\",  # base64.b64encode(s)     \"timestamp\": 1485238673.0 }

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

# **TerminalCommandsExportList**
> InlineResponse20054 TerminalCommandsExportList(ctx, optional)
terminal_commands_export_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TerminalCommandsExportListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerminalCommandsExportListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **optional.String**|  | 
 **systemUser** | **optional.String**|  | 
 **user** | **optional.String**|  | 
 **session** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20054**](inline_response_200_54.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalCommandsList**
> InlineResponse20054 TerminalCommandsList(ctx, optional)
terminal_commands_list

接受app发送来的command log, 格式如下 {     \"user\": \"admin\",     \"asset\": \"localhost\",     \"system_user\": \"web\",     \"session\": \"xxxxxx\",     \"input\": \"whoami\",     \"output\": \"d2hvbWFp\",  # base64.b64encode(s)     \"timestamp\": 1485238673.0 }

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TerminalCommandsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerminalCommandsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **optional.String**|  | 
 **systemUser** | **optional.String**|  | 
 **user** | **optional.String**|  | 
 **session** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20054**](inline_response_200_54.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalCommandsPartialUpdate**
> SessionCommand TerminalCommandsPartialUpdate(ctx, id, data)
terminal_commands_partial_update

接受app发送来的command log, 格式如下 {     \"user\": \"admin\",     \"asset\": \"localhost\",     \"system_user\": \"web\",     \"session\": \"xxxxxx\",     \"input\": \"whoami\",     \"output\": \"d2hvbWFp\",  # base64.b64encode(s)     \"timestamp\": 1485238673.0 }

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**SessionCommand**](SessionCommand.md)|  | 

### Return type

[**SessionCommand**](SessionCommand.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalCommandsRead**
> SessionCommand TerminalCommandsRead(ctx, id)
terminal_commands_read

接受app发送来的command log, 格式如下 {     \"user\": \"admin\",     \"asset\": \"localhost\",     \"system_user\": \"web\",     \"session\": \"xxxxxx\",     \"input\": \"whoami\",     \"output\": \"d2hvbWFp\",  # base64.b64encode(s)     \"timestamp\": 1485238673.0 }

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**SessionCommand**](SessionCommand.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminalCommandsUpdate**
> SessionCommand TerminalCommandsUpdate(ctx, id, data)
terminal_commands_update

接受app发送来的command log, 格式如下 {     \"user\": \"admin\",     \"asset\": \"localhost\",     \"system_user\": \"web\",     \"session\": \"xxxxxx\",     \"input\": \"whoami\",     \"output\": \"d2hvbWFp\",  # base64.b64encode(s)     \"timestamp\": 1485238673.0 }

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**SessionCommand**](SessionCommand.md)|  | 

### Return type

[**SessionCommand**](SessionCommand.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

