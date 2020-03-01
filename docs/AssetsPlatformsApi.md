# \AssetsPlatformsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsPlatformsCreate**](AssetsPlatformsApi.md#AssetsPlatformsCreate) | **Post** /assets/platforms/ | assets_platforms_create
[**AssetsPlatformsDelete**](AssetsPlatformsApi.md#AssetsPlatformsDelete) | **Delete** /assets/platforms/{id}/ | assets_platforms_delete
[**AssetsPlatformsList**](AssetsPlatformsApi.md#AssetsPlatformsList) | **Get** /assets/platforms/ | assets_platforms_list
[**AssetsPlatformsPartialUpdate**](AssetsPlatformsApi.md#AssetsPlatformsPartialUpdate) | **Patch** /assets/platforms/{id}/ | assets_platforms_partial_update
[**AssetsPlatformsRead**](AssetsPlatformsApi.md#AssetsPlatformsRead) | **Get** /assets/platforms/{id}/ | assets_platforms_read
[**AssetsPlatformsUpdate**](AssetsPlatformsApi.md#AssetsPlatformsUpdate) | **Put** /assets/platforms/{id}/ | assets_platforms_update


# **AssetsPlatformsCreate**
> Platform AssetsPlatformsCreate(ctx, data)
assets_platforms_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Platform**](Platform.md)|  | 

### Return type

[**Platform**](Platform.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsPlatformsDelete**
> AssetsPlatformsDelete(ctx, id)
assets_platforms_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this 系统平台. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsPlatformsList**
> InlineResponse20016 AssetsPlatformsList(ctx, optional)
assets_platforms_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsPlatformsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsPlatformsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **base** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsPlatformsPartialUpdate**
> Platform AssetsPlatformsPartialUpdate(ctx, id, data)
assets_platforms_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this 系统平台. | 
  **data** | [**Platform**](Platform.md)|  | 

### Return type

[**Platform**](Platform.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsPlatformsRead**
> Platform AssetsPlatformsRead(ctx, id)
assets_platforms_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this 系统平台. | 

### Return type

[**Platform**](Platform.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsPlatformsUpdate**
> Platform AssetsPlatformsUpdate(ctx, id, data)
assets_platforms_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this 系统平台. | 
  **data** | [**Platform**](Platform.md)|  | 

### Return type

[**Platform**](Platform.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

