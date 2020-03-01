# \AssetsGatheredUsersApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsGatheredUsersCreate**](AssetsGatheredUsersApi.md#AssetsGatheredUsersCreate) | **Post** /assets/gathered-users/ | assets_gathered-users_create
[**AssetsGatheredUsersDelete**](AssetsGatheredUsersApi.md#AssetsGatheredUsersDelete) | **Delete** /assets/gathered-users/{id}/ | assets_gathered-users_delete
[**AssetsGatheredUsersList**](AssetsGatheredUsersApi.md#AssetsGatheredUsersList) | **Get** /assets/gathered-users/ | assets_gathered-users_list
[**AssetsGatheredUsersPartialUpdate**](AssetsGatheredUsersApi.md#AssetsGatheredUsersPartialUpdate) | **Patch** /assets/gathered-users/{id}/ | assets_gathered-users_partial_update
[**AssetsGatheredUsersRead**](AssetsGatheredUsersApi.md#AssetsGatheredUsersRead) | **Get** /assets/gathered-users/{id}/ | assets_gathered-users_read
[**AssetsGatheredUsersUpdate**](AssetsGatheredUsersApi.md#AssetsGatheredUsersUpdate) | **Put** /assets/gathered-users/{id}/ | assets_gathered-users_update


# **AssetsGatheredUsersCreate**
> GatheredUser AssetsGatheredUsersCreate(ctx, data)
assets_gathered-users_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**GatheredUser**](GatheredUser.md)|  | 

### Return type

[**GatheredUser**](GatheredUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsGatheredUsersDelete**
> AssetsGatheredUsersDelete(ctx, id)
assets_gathered-users_delete



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

# **AssetsGatheredUsersList**
> InlineResponse20012 AssetsGatheredUsersList(ctx, optional)
assets_gathered-users_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsGatheredUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsGatheredUsersListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **optional.String**|  | 
 **username** | **optional.String**|  | 
 **present** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **node** | **optional.String**|  | 
 **all** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsGatheredUsersPartialUpdate**
> GatheredUser AssetsGatheredUsersPartialUpdate(ctx, id, data)
assets_gathered-users_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**GatheredUser**](GatheredUser.md)|  | 

### Return type

[**GatheredUser**](GatheredUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsGatheredUsersRead**
> GatheredUser AssetsGatheredUsersRead(ctx, id)
assets_gathered-users_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**GatheredUser**](GatheredUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsGatheredUsersUpdate**
> GatheredUser AssetsGatheredUsersUpdate(ctx, id, data)
assets_gathered-users_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**GatheredUser**](GatheredUser.md)|  | 

### Return type

[**GatheredUser**](GatheredUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

