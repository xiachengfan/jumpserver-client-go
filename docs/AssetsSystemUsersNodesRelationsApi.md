# \AssetsSystemUsersNodesRelationsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsSystemUsersNodesRelationsBulkDelete**](AssetsSystemUsersNodesRelationsApi.md#AssetsSystemUsersNodesRelationsBulkDelete) | **Delete** /assets/system-users-nodes-relations/ | assets_system-users-nodes-relations_bulk_delete
[**AssetsSystemUsersNodesRelationsBulkUpdate**](AssetsSystemUsersNodesRelationsApi.md#AssetsSystemUsersNodesRelationsBulkUpdate) | **Put** /assets/system-users-nodes-relations/ | assets_system-users-nodes-relations_bulk_update
[**AssetsSystemUsersNodesRelationsCreate**](AssetsSystemUsersNodesRelationsApi.md#AssetsSystemUsersNodesRelationsCreate) | **Post** /assets/system-users-nodes-relations/ | assets_system-users-nodes-relations_create
[**AssetsSystemUsersNodesRelationsDelete**](AssetsSystemUsersNodesRelationsApi.md#AssetsSystemUsersNodesRelationsDelete) | **Delete** /assets/system-users-nodes-relations/{id}/ | assets_system-users-nodes-relations_delete
[**AssetsSystemUsersNodesRelationsList**](AssetsSystemUsersNodesRelationsApi.md#AssetsSystemUsersNodesRelationsList) | **Get** /assets/system-users-nodes-relations/ | assets_system-users-nodes-relations_list
[**AssetsSystemUsersNodesRelationsPartialBulkUpdate**](AssetsSystemUsersNodesRelationsApi.md#AssetsSystemUsersNodesRelationsPartialBulkUpdate) | **Patch** /assets/system-users-nodes-relations/ | assets_system-users-nodes-relations_partial_bulk_update
[**AssetsSystemUsersNodesRelationsPartialUpdate**](AssetsSystemUsersNodesRelationsApi.md#AssetsSystemUsersNodesRelationsPartialUpdate) | **Patch** /assets/system-users-nodes-relations/{id}/ | assets_system-users-nodes-relations_partial_update
[**AssetsSystemUsersNodesRelationsRead**](AssetsSystemUsersNodesRelationsApi.md#AssetsSystemUsersNodesRelationsRead) | **Get** /assets/system-users-nodes-relations/{id}/ | assets_system-users-nodes-relations_read
[**AssetsSystemUsersNodesRelationsUpdate**](AssetsSystemUsersNodesRelationsApi.md#AssetsSystemUsersNodesRelationsUpdate) | **Put** /assets/system-users-nodes-relations/{id}/ | assets_system-users-nodes-relations_update


# **AssetsSystemUsersNodesRelationsBulkDelete**
> AssetsSystemUsersNodesRelationsBulkDelete(ctx, optional)
assets_system-users-nodes-relations_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsSystemUsersNodesRelationsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsSystemUsersNodesRelationsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **node** | **optional.String**|  | 
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

# **AssetsSystemUsersNodesRelationsBulkUpdate**
> SystemUserNodeRelation AssetsSystemUsersNodesRelationsBulkUpdate(ctx, data)
assets_system-users-nodes-relations_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SystemUserNodeRelation**](SystemUserNodeRelation.md)|  | 

### Return type

[**SystemUserNodeRelation**](SystemUserNodeRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersNodesRelationsCreate**
> SystemUserNodeRelation AssetsSystemUsersNodesRelationsCreate(ctx, data)
assets_system-users-nodes-relations_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SystemUserNodeRelation**](SystemUserNodeRelation.md)|  | 

### Return type

[**SystemUserNodeRelation**](SystemUserNodeRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersNodesRelationsDelete**
> AssetsSystemUsersNodesRelationsDelete(ctx, id)
assets_system-users-nodes-relations_delete



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

# **AssetsSystemUsersNodesRelationsList**
> InlineResponse20018 AssetsSystemUsersNodesRelationsList(ctx, optional)
assets_system-users-nodes-relations_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsSystemUsersNodesRelationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsSystemUsersNodesRelationsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float32**|  | 
 **node** | **optional.String**|  | 
 **systemuser** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersNodesRelationsPartialBulkUpdate**
> SystemUserNodeRelation AssetsSystemUsersNodesRelationsPartialBulkUpdate(ctx, data)
assets_system-users-nodes-relations_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SystemUserNodeRelation**](SystemUserNodeRelation.md)|  | 

### Return type

[**SystemUserNodeRelation**](SystemUserNodeRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersNodesRelationsPartialUpdate**
> SystemUserNodeRelation AssetsSystemUsersNodesRelationsPartialUpdate(ctx, id, data)
assets_system-users-nodes-relations_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**SystemUserNodeRelation**](SystemUserNodeRelation.md)|  | 

### Return type

[**SystemUserNodeRelation**](SystemUserNodeRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersNodesRelationsRead**
> SystemUserNodeRelation AssetsSystemUsersNodesRelationsRead(ctx, id)
assets_system-users-nodes-relations_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**SystemUserNodeRelation**](SystemUserNodeRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsSystemUsersNodesRelationsUpdate**
> SystemUserNodeRelation AssetsSystemUsersNodesRelationsUpdate(ctx, id, data)
assets_system-users-nodes-relations_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**SystemUserNodeRelation**](SystemUserNodeRelation.md)|  | 

### Return type

[**SystemUserNodeRelation**](SystemUserNodeRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

