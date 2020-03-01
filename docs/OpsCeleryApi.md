# \OpsCeleryApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpsCeleryPeriodTasksList**](OpsCeleryApi.md#OpsCeleryPeriodTasksList) | **Get** /ops/celery/period-tasks/ | ops_celery_period-tasks_list
[**OpsCeleryPeriodTasksPartialUpdate**](OpsCeleryApi.md#OpsCeleryPeriodTasksPartialUpdate) | **Patch** /ops/celery/period-tasks/{id}/ | ops_celery_period-tasks_partial_update
[**OpsCeleryPeriodTasksRead**](OpsCeleryApi.md#OpsCeleryPeriodTasksRead) | **Get** /ops/celery/period-tasks/{id}/ | ops_celery_period-tasks_read
[**OpsCeleryTaskLogRead**](OpsCeleryApi.md#OpsCeleryTaskLogRead) | **Get** /ops/celery/task/{id}/log/ | ops_celery_task_log_read
[**OpsCeleryTaskResultRead**](OpsCeleryApi.md#OpsCeleryTaskResultRead) | **Get** /ops/celery/task/{id}/result/ | ops_celery_task_result_read


# **OpsCeleryPeriodTasksList**
> InlineResponse20023 OpsCeleryPeriodTasksList(ctx, optional)
ops_celery_period-tasks_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpsCeleryPeriodTasksListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpsCeleryPeriodTasksListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsCeleryPeriodTasksPartialUpdate**
> CeleryPeriodTask OpsCeleryPeriodTasksPartialUpdate(ctx, id, data)
ops_celery_period-tasks_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this periodic task. | 
  **data** | [**CeleryPeriodTask**](CeleryPeriodTask.md)|  | 

### Return type

[**CeleryPeriodTask**](CeleryPeriodTask.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsCeleryPeriodTasksRead**
> CeleryPeriodTask OpsCeleryPeriodTasksRead(ctx, id)
ops_celery_period-tasks_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this periodic task. | 

### Return type

[**CeleryPeriodTask**](CeleryPeriodTask.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsCeleryTaskLogRead**
> Output OpsCeleryTaskLogRead(ctx, id)
ops_celery_task_log_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Output**](Output.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpsCeleryTaskResultRead**
> CeleryResult OpsCeleryTaskResultRead(ctx, id)
ops_celery_task_result_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**CeleryResult**](CeleryResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

