# \AssetsSystemUsersAssetsRelationsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsSystemUsersAssetsRelationsBulkDelete**](AssetsSystemUsersAssetsRelationsApi.md#AssetsSystemUsersAssetsRelationsBulkDelete) | **Delete** /assets/system-users-assets-relations/ | assets_system-users-assets-relations_bulk_delete
[**AssetsSystemUsersAssetsRelationsBulkUpdate**](AssetsSystemUsersAssetsRelationsApi.md#AssetsSystemUsersAssetsRelationsBulkUpdate) | **Put** /assets/system-users-assets-relations/ | assets_system-users-assets-relations_bulk_update
[**AssetsSystemUsersAssetsRelationsCreate**](AssetsSystemUsersAssetsRelationsApi.md#AssetsSystemUsersAssetsRelationsCreate) | **Post** /assets/system-users-assets-relations/ | assets_system-users-assets-relations_create
[**AssetsSystemUsersAssetsRelationsDelete**](AssetsSystemUsersAssetsRelationsApi.md#AssetsSystemUsersAssetsRelationsDelete) | **Delete** /assets/system-users-assets-relations/{id}/ | assets_system-users-assets-relations_delete
[**AssetsSystemUsersAssetsRelationsList**](AssetsSystemUsersAssetsRelationsApi.md#AssetsSystemUsersAssetsRelationsList) | **Get** /assets/system-users-assets-relations/ | assets_system-users-assets-relations_list
[**AssetsSystemUsersAssetsRelationsPartialBulkUpdate**](AssetsSystemUsersAssetsRelationsApi.md#AssetsSystemUsersAssetsRelationsPartialBulkUpdate) | **Patch** /assets/system-users-assets-relations/ | assets_system-users-assets-relations_partial_bulk_update
[**AssetsSystemUsersAssetsRelationsPartialUpdate**](AssetsSystemUsersAssetsRelationsApi.md#AssetsSystemUsersAssetsRelationsPartialUpdate) | **Patch** /assets/system-users-assets-relations/{id}/ | assets_system-users-assets-relations_partial_update
[**AssetsSystemUsersAssetsRelationsRead**](AssetsSystemUsersAssetsRelationsApi.md#AssetsSystemUsersAssetsRelationsRead) | **Get** /assets/system-users-assets-relations/{id}/ | assets_system-users-assets-relations_read
[**AssetsSystemUsersAssetsRelationsUpdate**](AssetsSystemUsersAssetsRelationsApi.md#AssetsSystemUsersAssetsRelationsUpdate) | **Put** /assets/system-users-assets-relations/{id}/ | assets_system-users-assets-relations_update


# **AssetsSystemUsersAssetsRelationsBulkDelete**
> AssetsSystemUsersAssetsRelationsBulkDelete(ctx, optional)
assets_system-users-assets-relations_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsSystemUsersAssetsRelationsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsSystemUsersAssetsRelationsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **asset** | **optional.String**|  | 
 **systemuser** | **optional.String**|  | 
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

# **AssetsSystemUsersAssetsRelationsBulkUpdate**
> SystemUserAssetRelation AssetsSystemUsersAssetsRelationsBulkUpdate(ctx, data)
assets_system-users-assets-relations_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SystemUserAssetRelation**](SystemUserAssetRelation.md)|  | 

### Return type

[**SystemUserAssetRelation**](SystemUserAssetRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersAssetsRelationsCreate**
> SystemUserAssetRelation AssetsSystemUsersAssetsRelationsCreate(ctx, data)
assets_system-users-assets-relations_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SystemUserAssetRelation**](SystemUserAssetRelation.md)|  | 

### Return type

[**SystemUserAssetRelation**](SystemUserAssetRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersAssetsRelationsDelete**
> AssetsSystemUsersAssetsRelationsDelete(ctx, id)
assets_system-users-assets-relations_delete



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

# **AssetsSystemUsersAssetsRelationsList**
> InlineResponse20017 AssetsSystemUsersAssetsRelationsList(ctx, optional)
assets_system-users-assets-relations_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsSystemUsersAssetsRelationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsSystemUsersAssetsRelationsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **asset** | **optional.String**|  | 
 **systemuser** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersAssetsRelationsPartialBulkUpdate**
> SystemUserAssetRelation AssetsSystemUsersAssetsRelationsPartialBulkUpdate(ctx, data)
assets_system-users-assets-relations_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SystemUserAssetRelation**](SystemUserAssetRelation.md)|  | 

### Return type

[**SystemUserAssetRelation**](SystemUserAssetRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersAssetsRelationsPartialUpdate**
> SystemUserAssetRelation AssetsSystemUsersAssetsRelationsPartialUpdate(ctx, id, data)
assets_system-users-assets-relations_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**SystemUserAssetRelation**](SystemUserAssetRelation.md)|  | 

### Return type

[**SystemUserAssetRelation**](SystemUserAssetRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersAssetsRelationsRead**
> SystemUserAssetRelation AssetsSystemUsersAssetsRelationsRead(ctx, id)
assets_system-users-assets-relations_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**SystemUserAssetRelation**](SystemUserAssetRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersAssetsRelationsUpdate**
> SystemUserAssetRelation AssetsSystemUsersAssetsRelationsUpdate(ctx, id, data)
assets_system-users-assets-relations_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**SystemUserAssetRelation**](SystemUserAssetRelation.md)|  | 

### Return type

[**SystemUserAssetRelation**](SystemUserAssetRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

