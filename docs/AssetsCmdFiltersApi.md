# \AssetsCmdFiltersApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsCmdFiltersBulkDelete**](AssetsCmdFiltersApi.md#AssetsCmdFiltersBulkDelete) | **Delete** /assets/cmd-filters/ | assets_cmd-filters_bulk_delete
[**AssetsCmdFiltersBulkUpdate**](AssetsCmdFiltersApi.md#AssetsCmdFiltersBulkUpdate) | **Put** /assets/cmd-filters/ | assets_cmd-filters_bulk_update
[**AssetsCmdFiltersCreate**](AssetsCmdFiltersApi.md#AssetsCmdFiltersCreate) | **Post** /assets/cmd-filters/ | assets_cmd-filters_create
[**AssetsCmdFiltersDelete**](AssetsCmdFiltersApi.md#AssetsCmdFiltersDelete) | **Delete** /assets/cmd-filters/{id}/ | assets_cmd-filters_delete
[**AssetsCmdFiltersList**](AssetsCmdFiltersApi.md#AssetsCmdFiltersList) | **Get** /assets/cmd-filters/ | assets_cmd-filters_list
[**AssetsCmdFiltersPartialBulkUpdate**](AssetsCmdFiltersApi.md#AssetsCmdFiltersPartialBulkUpdate) | **Patch** /assets/cmd-filters/ | assets_cmd-filters_partial_bulk_update
[**AssetsCmdFiltersPartialUpdate**](AssetsCmdFiltersApi.md#AssetsCmdFiltersPartialUpdate) | **Patch** /assets/cmd-filters/{id}/ | assets_cmd-filters_partial_update
[**AssetsCmdFiltersRead**](AssetsCmdFiltersApi.md#AssetsCmdFiltersRead) | **Get** /assets/cmd-filters/{id}/ | assets_cmd-filters_read
[**AssetsCmdFiltersRulesCreate**](AssetsCmdFiltersApi.md#AssetsCmdFiltersRulesCreate) | **Post** /assets/cmd-filters/{filter_pk}/rules/ | assets_cmd-filters_rules_create
[**AssetsCmdFiltersRulesDelete**](AssetsCmdFiltersApi.md#AssetsCmdFiltersRulesDelete) | **Delete** /assets/cmd-filters/{filter_pk}/rules/{id}/ | assets_cmd-filters_rules_delete
[**AssetsCmdFiltersRulesList**](AssetsCmdFiltersApi.md#AssetsCmdFiltersRulesList) | **Get** /assets/cmd-filters/{filter_pk}/rules/ | assets_cmd-filters_rules_list
[**AssetsCmdFiltersRulesPartialUpdate**](AssetsCmdFiltersApi.md#AssetsCmdFiltersRulesPartialUpdate) | **Patch** /assets/cmd-filters/{filter_pk}/rules/{id}/ | assets_cmd-filters_rules_partial_update
[**AssetsCmdFiltersRulesRead**](AssetsCmdFiltersApi.md#AssetsCmdFiltersRulesRead) | **Get** /assets/cmd-filters/{filter_pk}/rules/{id}/ | assets_cmd-filters_rules_read
[**AssetsCmdFiltersRulesUpdate**](AssetsCmdFiltersApi.md#AssetsCmdFiltersRulesUpdate) | **Put** /assets/cmd-filters/{filter_pk}/rules/{id}/ | assets_cmd-filters_rules_update
[**AssetsCmdFiltersUpdate**](AssetsCmdFiltersApi.md#AssetsCmdFiltersUpdate) | **Put** /assets/cmd-filters/{id}/ | assets_cmd-filters_update


# **AssetsCmdFiltersBulkDelete**
> AssetsCmdFiltersBulkDelete(ctx, optional)
assets_cmd-filters_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsCmdFiltersBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsCmdFiltersBulkDeleteOpts struct

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

# **AssetsCmdFiltersBulkUpdate**
> CommandFilter AssetsCmdFiltersBulkUpdate(ctx, data)
assets_cmd-filters_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CommandFilter**](CommandFilter.md)|  | 

### Return type

[**CommandFilter**](CommandFilter.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsCmdFiltersCreate**
> CommandFilter AssetsCmdFiltersCreate(ctx, data)
assets_cmd-filters_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CommandFilter**](CommandFilter.md)|  | 

### Return type

[**CommandFilter**](CommandFilter.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsCmdFiltersDelete**
> AssetsCmdFiltersDelete(ctx, id)
assets_cmd-filters_delete



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

# **AssetsCmdFiltersList**
> InlineResponse2007 AssetsCmdFiltersList(ctx, optional)
assets_cmd-filters_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsCmdFiltersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsCmdFiltersListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsCmdFiltersPartialBulkUpdate**
> CommandFilter AssetsCmdFiltersPartialBulkUpdate(ctx, data)
assets_cmd-filters_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CommandFilter**](CommandFilter.md)|  | 

### Return type

[**CommandFilter**](CommandFilter.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsCmdFiltersPartialUpdate**
> CommandFilter AssetsCmdFiltersPartialUpdate(ctx, id, data)
assets_cmd-filters_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**CommandFilter**](CommandFilter.md)|  | 

### Return type

[**CommandFilter**](CommandFilter.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsCmdFiltersRead**
> CommandFilter AssetsCmdFiltersRead(ctx, id)
assets_cmd-filters_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**CommandFilter**](CommandFilter.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsCmdFiltersRulesCreate**
> CommandFilterRule AssetsCmdFiltersRulesCreate(ctx, filterPk, data)
assets_cmd-filters_rules_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filterPk** | **string**|  | 
  **data** | [**CommandFilterRule**](CommandFilterRule.md)|  | 

### Return type

[**CommandFilterRule**](CommandFilterRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsCmdFiltersRulesDelete**
> AssetsCmdFiltersRulesDelete(ctx, filterPk, id)
assets_cmd-filters_rules_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filterPk** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsCmdFiltersRulesList**
> InlineResponse2008 AssetsCmdFiltersRulesList(ctx, filterPk, optional)
assets_cmd-filters_rules_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filterPk** | **string**|  | 
 **optional** | ***AssetsCmdFiltersRulesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsCmdFiltersRulesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsCmdFiltersRulesPartialUpdate**
> CommandFilterRule AssetsCmdFiltersRulesPartialUpdate(ctx, filterPk, id, data)
assets_cmd-filters_rules_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filterPk** | **string**|  | 
  **id** | **string**|  | 
  **data** | [**CommandFilterRule**](CommandFilterRule.md)|  | 

### Return type

[**CommandFilterRule**](CommandFilterRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsCmdFiltersRulesRead**
> CommandFilterRule AssetsCmdFiltersRulesRead(ctx, filterPk, id)
assets_cmd-filters_rules_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filterPk** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**CommandFilterRule**](CommandFilterRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsCmdFiltersRulesUpdate**
> CommandFilterRule AssetsCmdFiltersRulesUpdate(ctx, filterPk, id, data)
assets_cmd-filters_rules_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filterPk** | **string**|  | 
  **id** | **string**|  | 
  **data** | [**CommandFilterRule**](CommandFilterRule.md)|  | 

### Return type

[**CommandFilterRule**](CommandFilterRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsCmdFiltersUpdate**
> CommandFilter AssetsCmdFiltersUpdate(ctx, id, data)
assets_cmd-filters_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**CommandFilter**](CommandFilter.md)|  | 

### Return type

[**CommandFilter**](CommandFilter.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

