# \AssetsDomainsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsDomainsBulkDelete**](AssetsDomainsApi.md#AssetsDomainsBulkDelete) | **Delete** /assets/domains/ | assets_domains_bulk_delete
[**AssetsDomainsBulkUpdate**](AssetsDomainsApi.md#AssetsDomainsBulkUpdate) | **Put** /assets/domains/ | assets_domains_bulk_update
[**AssetsDomainsCreate**](AssetsDomainsApi.md#AssetsDomainsCreate) | **Post** /assets/domains/ | assets_domains_create
[**AssetsDomainsDelete**](AssetsDomainsApi.md#AssetsDomainsDelete) | **Delete** /assets/domains/{id}/ | assets_domains_delete
[**AssetsDomainsList**](AssetsDomainsApi.md#AssetsDomainsList) | **Get** /assets/domains/ | assets_domains_list
[**AssetsDomainsPartialBulkUpdate**](AssetsDomainsApi.md#AssetsDomainsPartialBulkUpdate) | **Patch** /assets/domains/ | assets_domains_partial_bulk_update
[**AssetsDomainsPartialUpdate**](AssetsDomainsApi.md#AssetsDomainsPartialUpdate) | **Patch** /assets/domains/{id}/ | assets_domains_partial_update
[**AssetsDomainsRead**](AssetsDomainsApi.md#AssetsDomainsRead) | **Get** /assets/domains/{id}/ | assets_domains_read
[**AssetsDomainsUpdate**](AssetsDomainsApi.md#AssetsDomainsUpdate) | **Put** /assets/domains/{id}/ | assets_domains_update


# **AssetsDomainsBulkDelete**
> AssetsDomainsBulkDelete(ctx, optional)
assets_domains_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsDomainsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsDomainsBulkDeleteOpts struct

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

# **AssetsDomainsBulkUpdate**
> Domain AssetsDomainsBulkUpdate(ctx, data)
assets_domains_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Domain**](Domain.md)|  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsDomainsCreate**
> Domain AssetsDomainsCreate(ctx, data)
assets_domains_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Domain**](Domain.md)|  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsDomainsDelete**
> AssetsDomainsDelete(ctx, id)
assets_domains_delete



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

# **AssetsDomainsList**
> InlineResponse2009 AssetsDomainsList(ctx, optional)
assets_domains_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsDomainsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsDomainsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsDomainsPartialBulkUpdate**
> Domain AssetsDomainsPartialBulkUpdate(ctx, data)
assets_domains_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Domain**](Domain.md)|  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsDomainsPartialUpdate**
> Domain AssetsDomainsPartialUpdate(ctx, id, data)
assets_domains_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**Domain**](Domain.md)|  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsDomainsRead**
> Domain AssetsDomainsRead(ctx, id)
assets_domains_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsDomainsUpdate**
> Domain AssetsDomainsUpdate(ctx, id, data)
assets_domains_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**Domain**](Domain.md)|  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

