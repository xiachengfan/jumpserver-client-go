# \UsersUsersGroupsRelationsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersUsersGroupsRelationsBulkDelete**](UsersUsersGroupsRelationsApi.md#UsersUsersGroupsRelationsBulkDelete) | **Delete** /users/users-groups-relations/ | users_users-groups-relations_bulk_delete
[**UsersUsersGroupsRelationsBulkUpdate**](UsersUsersGroupsRelationsApi.md#UsersUsersGroupsRelationsBulkUpdate) | **Put** /users/users-groups-relations/ | users_users-groups-relations_bulk_update
[**UsersUsersGroupsRelationsCreate**](UsersUsersGroupsRelationsApi.md#UsersUsersGroupsRelationsCreate) | **Post** /users/users-groups-relations/ | users_users-groups-relations_create
[**UsersUsersGroupsRelationsDelete**](UsersUsersGroupsRelationsApi.md#UsersUsersGroupsRelationsDelete) | **Delete** /users/users-groups-relations/{id}/ | users_users-groups-relations_delete
[**UsersUsersGroupsRelationsList**](UsersUsersGroupsRelationsApi.md#UsersUsersGroupsRelationsList) | **Get** /users/users-groups-relations/ | users_users-groups-relations_list
[**UsersUsersGroupsRelationsPartialBulkUpdate**](UsersUsersGroupsRelationsApi.md#UsersUsersGroupsRelationsPartialBulkUpdate) | **Patch** /users/users-groups-relations/ | users_users-groups-relations_partial_bulk_update
[**UsersUsersGroupsRelationsPartialUpdate**](UsersUsersGroupsRelationsApi.md#UsersUsersGroupsRelationsPartialUpdate) | **Patch** /users/users-groups-relations/{id}/ | users_users-groups-relations_partial_update
[**UsersUsersGroupsRelationsRead**](UsersUsersGroupsRelationsApi.md#UsersUsersGroupsRelationsRead) | **Get** /users/users-groups-relations/{id}/ | users_users-groups-relations_read
[**UsersUsersGroupsRelationsUpdate**](UsersUsersGroupsRelationsApi.md#UsersUsersGroupsRelationsUpdate) | **Put** /users/users-groups-relations/{id}/ | users_users-groups-relations_update


# **UsersUsersGroupsRelationsBulkDelete**
> UsersUsersGroupsRelationsBulkDelete(ctx, optional)
users_users-groups-relations_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersUsersGroupsRelationsBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersUsersGroupsRelationsBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **optional.String**|  | 
 **usergroup** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersGroupsRelationsBulkUpdate**
> UserUserGroupRelation UsersUsersGroupsRelationsBulkUpdate(ctx, data)
users_users-groups-relations_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**UserUserGroupRelation**](UserUserGroupRelation.md)|  | 

### Return type

[**UserUserGroupRelation**](UserUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersGroupsRelationsCreate**
> UserUserGroupRelation UsersUsersGroupsRelationsCreate(ctx, data)
users_users-groups-relations_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**UserUserGroupRelation**](UserUserGroupRelation.md)|  | 

### Return type

[**UserUserGroupRelation**](UserUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersGroupsRelationsDelete**
> UsersUsersGroupsRelationsDelete(ctx, id)
users_users-groups-relations_delete



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

# **UsersUsersGroupsRelationsList**
> InlineResponse20061 UsersUsersGroupsRelationsList(ctx, optional)
users_users-groups-relations_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersUsersGroupsRelationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersUsersGroupsRelationsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **optional.String**|  | 
 **usergroup** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20061**](inline_response_200_61.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersGroupsRelationsPartialBulkUpdate**
> UserUserGroupRelation UsersUsersGroupsRelationsPartialBulkUpdate(ctx, data)
users_users-groups-relations_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**UserUserGroupRelation**](UserUserGroupRelation.md)|  | 

### Return type

[**UserUserGroupRelation**](UserUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersGroupsRelationsPartialUpdate**
> UserUserGroupRelation UsersUsersGroupsRelationsPartialUpdate(ctx, id, data)
users_users-groups-relations_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**UserUserGroupRelation**](UserUserGroupRelation.md)|  | 

### Return type

[**UserUserGroupRelation**](UserUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersGroupsRelationsRead**
> UserUserGroupRelation UsersUsersGroupsRelationsRead(ctx, id)
users_users-groups-relations_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**UserUserGroupRelation**](UserUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersGroupsRelationsUpdate**
> UserUserGroupRelation UsersUsersGroupsRelationsUpdate(ctx, id, data)
users_users-groups-relations_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**UserUserGroupRelation**](UserUserGroupRelation.md)|  | 

### Return type

[**UserUserGroupRelation**](UserUserGroupRelation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

