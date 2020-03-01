# \PermsUserGroupsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermsUserGroupsAssetsList**](PermsUserGroupsApi.md#PermsUserGroupsAssetsList) | **Get** /perms/user-groups/{id}/assets/ | perms_user-groups_assets_list
[**PermsUserGroupsAssetsSystemUsersList**](PermsUserGroupsApi.md#PermsUserGroupsAssetsSystemUsersList) | **Get** /perms/user-groups/{id}/assets/{asset_id}/system-users/ | perms_user-groups_assets_system-users_list
[**PermsUserGroupsDatabaseAppsList**](PermsUserGroupsApi.md#PermsUserGroupsDatabaseAppsList) | **Get** /perms/user-groups/{id}/database-apps/ | perms_user-groups_database-apps_list
[**PermsUserGroupsNodesAssetsList**](PermsUserGroupsApi.md#PermsUserGroupsNodesAssetsList) | **Get** /perms/user-groups/{id}/nodes/{node_id}/assets/ | perms_user-groups_nodes_assets_list
[**PermsUserGroupsNodesChildrenList**](PermsUserGroupsApi.md#PermsUserGroupsNodesChildrenList) | **Get** /perms/user-groups/{id}/nodes/children/ | perms_user-groups_nodes_children_list
[**PermsUserGroupsNodesChildrenTreeList**](PermsUserGroupsApi.md#PermsUserGroupsNodesChildrenTreeList) | **Get** /perms/user-groups/{id}/nodes/children/tree/ | perms_user-groups_nodes_children_tree_list
[**PermsUserGroupsNodesList**](PermsUserGroupsApi.md#PermsUserGroupsNodesList) | **Get** /perms/user-groups/{id}/nodes/ | perms_user-groups_nodes_list
[**PermsUserGroupsRemoteAppsList**](PermsUserGroupsApi.md#PermsUserGroupsRemoteAppsList) | **Get** /perms/user-groups/{id}/remote-apps/ | perms_user-groups_remote-apps_list


# **PermsUserGroupsAssetsList**
> InlineResponse20046 PermsUserGroupsAssetsList(ctx, id, optional)
perms_user-groups_assets_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***PermsUserGroupsAssetsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsUserGroupsAssetsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostname** | **optional.String**|  | 
 **ip** | **optional.String**|  | 
 **id2** | **optional.String**|  | 
 **comment** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsUserGroupsAssetsSystemUsersList**
> InlineResponse20047 PermsUserGroupsAssetsSystemUsersList(ctx, assetId, id, optional)
perms_user-groups_assets_system-users_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assetId** | **string**|  | 
  **id** | **string**|  | 
 **optional** | ***PermsUserGroupsAssetsSystemUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsUserGroupsAssetsSystemUsersListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsUserGroupsDatabaseAppsList**
> InlineResponse200 PermsUserGroupsDatabaseAppsList(ctx, id, optional)
perms_user-groups_database-apps_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***PermsUserGroupsDatabaseAppsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsUserGroupsDatabaseAppsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsUserGroupsNodesAssetsList**
> InlineResponse20046 PermsUserGroupsNodesAssetsList(ctx, id, nodeId, optional)
perms_user-groups_nodes_assets_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **nodeId** | **string**|  | 
 **optional** | ***PermsUserGroupsNodesAssetsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsUserGroupsNodesAssetsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsUserGroupsNodesChildrenList**
> InlineResponse20048 PermsUserGroupsNodesChildrenList(ctx, id, optional)
perms_user-groups_nodes_children_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***PermsUserGroupsNodesChildrenListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsUserGroupsNodesChildrenListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsUserGroupsNodesChildrenTreeList**
> InlineResponse20049 PermsUserGroupsNodesChildrenTreeList(ctx, id, optional)
perms_user-groups_nodes_children_tree_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***PermsUserGroupsNodesChildrenTreeListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsUserGroupsNodesChildrenTreeListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20049**](inline_response_200_49.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsUserGroupsNodesList**
> InlineResponse20048 PermsUserGroupsNodesList(ctx, id, optional)
perms_user-groups_nodes_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***PermsUserGroupsNodesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsUserGroupsNodesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermsUserGroupsRemoteAppsList**
> InlineResponse2001 PermsUserGroupsRemoteAppsList(ctx, id, optional)
perms_user-groups_remote-apps_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***PermsUserGroupsRemoteAppsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermsUserGroupsRemoteAppsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

