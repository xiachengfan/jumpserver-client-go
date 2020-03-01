# \AssetsNodesApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsNodesAssetsAddPartialUpdate**](AssetsNodesApi.md#AssetsNodesAssetsAddPartialUpdate) | **Patch** /assets/nodes/{id}/assets/add/ | assets_nodes_assets_add_partial_update
[**AssetsNodesAssetsAddUpdate**](AssetsNodesApi.md#AssetsNodesAssetsAddUpdate) | **Put** /assets/nodes/{id}/assets/add/ | assets_nodes_assets_add_update
[**AssetsNodesAssetsList**](AssetsNodesApi.md#AssetsNodesAssetsList) | **Get** /assets/nodes/{id}/assets/ | assets_nodes_assets_list
[**AssetsNodesAssetsRemovePartialUpdate**](AssetsNodesApi.md#AssetsNodesAssetsRemovePartialUpdate) | **Patch** /assets/nodes/{id}/assets/remove/ | assets_nodes_assets_remove_partial_update
[**AssetsNodesAssetsRemoveUpdate**](AssetsNodesApi.md#AssetsNodesAssetsRemoveUpdate) | **Put** /assets/nodes/{id}/assets/remove/ | assets_nodes_assets_remove_update
[**AssetsNodesAssetsReplacePartialUpdate**](AssetsNodesApi.md#AssetsNodesAssetsReplacePartialUpdate) | **Patch** /assets/nodes/{id}/assets/replace/ | assets_nodes_assets_replace_partial_update
[**AssetsNodesAssetsReplaceUpdate**](AssetsNodesApi.md#AssetsNodesAssetsReplaceUpdate) | **Put** /assets/nodes/{id}/assets/replace/ | assets_nodes_assets_replace_update
[**AssetsNodesCacheDelete**](AssetsNodesApi.md#AssetsNodesCacheDelete) | **Delete** /assets/nodes/cache/ | assets_nodes_cache_delete
[**AssetsNodesCacheList**](AssetsNodesApi.md#AssetsNodesCacheList) | **Get** /assets/nodes/cache/ | assets_nodes_cache_list
[**AssetsNodesChildrenAddPartialUpdate**](AssetsNodesApi.md#AssetsNodesChildrenAddPartialUpdate) | **Patch** /assets/nodes/{id}/children/add/ | assets_nodes_children_add_partial_update
[**AssetsNodesChildrenAddUpdate**](AssetsNodesApi.md#AssetsNodesChildrenAddUpdate) | **Put** /assets/nodes/{id}/children/add/ | assets_nodes_children_add_update
[**AssetsNodesChildrenCreate**](AssetsNodesApi.md#AssetsNodesChildrenCreate) | **Post** /assets/nodes/{id}/children/ | assets_nodes_children_create
[**AssetsNodesChildrenList**](AssetsNodesApi.md#AssetsNodesChildrenList) | **Get** /assets/nodes/{id}/children/ | assets_nodes_children_list
[**AssetsNodesChildrenTreeList**](AssetsNodesApi.md#AssetsNodesChildrenTreeList) | **Get** /assets/nodes/children/tree/ | assets_nodes_children_tree_list
[**AssetsNodesCreate**](AssetsNodesApi.md#AssetsNodesCreate) | **Post** /assets/nodes/ | assets_nodes_create
[**AssetsNodesDelete**](AssetsNodesApi.md#AssetsNodesDelete) | **Delete** /assets/nodes/{id}/ | assets_nodes_delete
[**AssetsNodesList**](AssetsNodesApi.md#AssetsNodesList) | **Get** /assets/nodes/ | assets_nodes_list
[**AssetsNodesPartialUpdate**](AssetsNodesApi.md#AssetsNodesPartialUpdate) | **Patch** /assets/nodes/{id}/ | assets_nodes_partial_update
[**AssetsNodesRead**](AssetsNodesApi.md#AssetsNodesRead) | **Get** /assets/nodes/{id}/ | assets_nodes_read
[**AssetsNodesRefreshHardwareInfoList**](AssetsNodesApi.md#AssetsNodesRefreshHardwareInfoList) | **Get** /assets/nodes/{id}/refresh-hardware-info/ | assets_nodes_refresh-hardware-info_list
[**AssetsNodesRootChildrenCreate**](AssetsNodesApi.md#AssetsNodesRootChildrenCreate) | **Post** /assets/nodes/children/ | assets_nodes_root_children_create
[**AssetsNodesRootChildrenList**](AssetsNodesApi.md#AssetsNodesRootChildrenList) | **Get** /assets/nodes/children/ | assets_nodes_root_children_list
[**AssetsNodesTestConnectiveList**](AssetsNodesApi.md#AssetsNodesTestConnectiveList) | **Get** /assets/nodes/{id}/test-connective/ | assets_nodes_test-connective_list
[**AssetsNodesTreeList**](AssetsNodesApi.md#AssetsNodesTreeList) | **Get** /assets/nodes/tree/ | assets_nodes_tree_list
[**AssetsNodesUpdate**](AssetsNodesApi.md#AssetsNodesUpdate) | **Put** /assets/nodes/{id}/ | assets_nodes_update


# **AssetsNodesAssetsAddPartialUpdate**
> NodeAssets AssetsNodesAssetsAddPartialUpdate(ctx, id, data)
assets_nodes_assets_add_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**NodeAssets**](NodeAssets.md)|  | 

### Return type

[**NodeAssets**](NodeAssets.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesAssetsAddUpdate**
> NodeAssets AssetsNodesAssetsAddUpdate(ctx, id, data)
assets_nodes_assets_add_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**NodeAssets**](NodeAssets.md)|  | 

### Return type

[**NodeAssets**](NodeAssets.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesAssetsList**
> InlineResponse2006 AssetsNodesAssetsList(ctx, id, optional)
assets_nodes_assets_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AssetsNodesAssetsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsNodesAssetsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesAssetsRemovePartialUpdate**
> NodeAssets AssetsNodesAssetsRemovePartialUpdate(ctx, id, data)
assets_nodes_assets_remove_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**NodeAssets**](NodeAssets.md)|  | 

### Return type

[**NodeAssets**](NodeAssets.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesAssetsRemoveUpdate**
> NodeAssets AssetsNodesAssetsRemoveUpdate(ctx, id, data)
assets_nodes_assets_remove_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**NodeAssets**](NodeAssets.md)|  | 

### Return type

[**NodeAssets**](NodeAssets.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesAssetsReplacePartialUpdate**
> NodeAssets AssetsNodesAssetsReplacePartialUpdate(ctx, id, data)
assets_nodes_assets_replace_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**NodeAssets**](NodeAssets.md)|  | 

### Return type

[**NodeAssets**](NodeAssets.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesAssetsReplaceUpdate**
> NodeAssets AssetsNodesAssetsReplaceUpdate(ctx, id, data)
assets_nodes_assets_replace_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**NodeAssets**](NodeAssets.md)|  | 

### Return type

[**NodeAssets**](NodeAssets.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesCacheDelete**
> AssetsNodesCacheDelete(ctx, )
assets_nodes_cache_delete



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesCacheList**
> AssetsNodesCacheList(ctx, )
assets_nodes_cache_list



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesChildrenAddPartialUpdate**
> NodeAddChildren AssetsNodesChildrenAddPartialUpdate(ctx, id, data)
assets_nodes_children_add_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**NodeAddChildren**](NodeAddChildren.md)|  | 

### Return type

[**NodeAddChildren**](NodeAddChildren.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesChildrenAddUpdate**
> NodeAddChildren AssetsNodesChildrenAddUpdate(ctx, id, data)
assets_nodes_children_add_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**NodeAddChildren**](NodeAddChildren.md)|  | 

### Return type

[**NodeAddChildren**](NodeAddChildren.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesChildrenCreate**
> Node AssetsNodesChildrenCreate(ctx, id, data)
assets_nodes_children_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**Node**](Node.md)|  | 

### Return type

[**Node**](Node.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesChildrenList**
> InlineResponse20014 AssetsNodesChildrenList(ctx, id, optional)
assets_nodes_children_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AssetsNodesChildrenListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsNodesChildrenListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesChildrenTreeList**
> InlineResponse20015 AssetsNodesChildrenTreeList(ctx, optional)
assets_nodes_children_tree_list

节点子节点作为树返回， [   {     \"id\": \"\",     \"name\": \"\",     \"pId\": \"\",     \"meta\": \"\"   } ]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsNodesChildrenTreeListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsNodesChildrenTreeListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesCreate**
> Node AssetsNodesCreate(ctx, data)
assets_nodes_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Node**](Node.md)|  | 

### Return type

[**Node**](Node.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesDelete**
> AssetsNodesDelete(ctx, id)
assets_nodes_delete



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

# **AssetsNodesList**
> InlineResponse20014 AssetsNodesList(ctx, optional)
assets_nodes_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsNodesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsNodesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **value** | **optional.String**|  | 
 **key** | **optional.String**|  | 
 **id** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesPartialUpdate**
> Node AssetsNodesPartialUpdate(ctx, id, data)
assets_nodes_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**Node**](Node.md)|  | 

### Return type

[**Node**](Node.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesRead**
> Node AssetsNodesRead(ctx, id)
assets_nodes_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Node**](Node.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesRefreshHardwareInfoList**
> AssetsNodesRefreshHardwareInfoList(ctx, id)
assets_nodes_refresh-hardware-info_list



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

# **AssetsNodesRootChildrenCreate**
> Node AssetsNodesRootChildrenCreate(ctx, data)
assets_nodes_root_children_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Node**](Node.md)|  | 

### Return type

[**Node**](Node.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesRootChildrenList**
> InlineResponse20014 AssetsNodesRootChildrenList(ctx, optional)
assets_nodes_root_children_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsNodesRootChildrenListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsNodesRootChildrenListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesTestConnectiveList**
> AssetsNodesTestConnectiveList(ctx, id)
assets_nodes_test-connective_list



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

# **AssetsNodesTreeList**
> InlineResponse20015 AssetsNodesTreeList(ctx, optional)
assets_nodes_tree_list

获取节点列表树 [   {     \"id\": \"\",     \"name\": \"\",     \"pId\": \"\",     \"meta\": \"\"   } ]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsNodesTreeListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsNodesTreeListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetsNodesUpdate**
> Node AssetsNodesUpdate(ctx, id, data)
assets_nodes_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**Node**](Node.md)|  | 

### Return type

[**Node**](Node.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

