# \AuditsFtpLogsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditsFtpLogsCreate**](AuditsFtpLogsApi.md#AuditsFtpLogsCreate) | **Post** /audits/ftp-logs/ | audits_ftp-logs_create
[**AuditsFtpLogsList**](AuditsFtpLogsApi.md#AuditsFtpLogsList) | **Get** /audits/ftp-logs/ | audits_ftp-logs_list
[**AuditsFtpLogsRead**](AuditsFtpLogsApi.md#AuditsFtpLogsRead) | **Get** /audits/ftp-logs/{id}/ | audits_ftp-logs_read


# **AuditsFtpLogsCreate**
> FtpLog AuditsFtpLogsCreate(ctx, data)
audits_ftp-logs_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**FtpLog**](FtpLog.md)|  | 

### Return type

[**FtpLog**](FTPLog.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuditsFtpLogsList**
> InlineResponse20020 AuditsFtpLogsList(ctx, optional)
audits_ftp-logs_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuditsFtpLogsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuditsFtpLogsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuditsFtpLogsRead**
> FtpLog AuditsFtpLogsRead(ctx, id)
audits_ftp-logs_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**FtpLog**](FTPLog.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

