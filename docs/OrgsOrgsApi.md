# \OrgsOrgsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgsCreate**](OrgsOrgsApi.md#OrgsOrgsCreate) | **Post** /orgs/orgs/ | orgs_orgs_create
[**OrgsOrgsDelete**](OrgsOrgsApi.md#OrgsOrgsDelete) | **Delete** /orgs/orgs/{id}/ | orgs_orgs_delete
[**OrgsOrgsList**](OrgsOrgsApi.md#OrgsOrgsList) | **Get** /orgs/orgs/ | orgs_orgs_list
[**OrgsOrgsMembershipAdminsCreate**](OrgsOrgsApi.md#OrgsOrgsMembershipAdminsCreate) | **Post** /orgs/orgs/{org_id}/membership/admins/ | orgs_orgs_membership_admins_create
[**OrgsOrgsMembershipAdminsDelete**](OrgsOrgsApi.md#OrgsOrgsMembershipAdminsDelete) | **Delete** /orgs/orgs/{org_id}/membership/admins/{user_id}/ | orgs_orgs_membership_admins_delete
[**OrgsOrgsMembershipAdminsList**](OrgsOrgsApi.md#OrgsOrgsMembershipAdminsList) | **Get** /orgs/orgs/{org_id}/membership/admins/ | orgs_orgs_membership_admins_list
[**OrgsOrgsMembershipAdminsRead**](OrgsOrgsApi.md#OrgsOrgsMembershipAdminsRead) | **Get** /orgs/orgs/{org_id}/membership/admins/{user_id}/ | orgs_orgs_membership_admins_read
[**OrgsOrgsMembershipUsersCreate**](OrgsOrgsApi.md#OrgsOrgsMembershipUsersCreate) | **Post** /orgs/orgs/{org_id}/membership/users/ | orgs_orgs_membership_users_create
[**OrgsOrgsMembershipUsersDelete**](OrgsOrgsApi.md#OrgsOrgsMembershipUsersDelete) | **Delete** /orgs/orgs/{org_id}/membership/users/{user_id}/ | orgs_orgs_membership_users_delete
[**OrgsOrgsMembershipUsersList**](OrgsOrgsApi.md#OrgsOrgsMembershipUsersList) | **Get** /orgs/orgs/{org_id}/membership/users/ | orgs_orgs_membership_users_list
[**OrgsOrgsMembershipUsersRead**](OrgsOrgsApi.md#OrgsOrgsMembershipUsersRead) | **Get** /orgs/orgs/{org_id}/membership/users/{user_id}/ | orgs_orgs_membership_users_read
[**OrgsOrgsPartialUpdate**](OrgsOrgsApi.md#OrgsOrgsPartialUpdate) | **Patch** /orgs/orgs/{id}/ | orgs_orgs_partial_update
[**OrgsOrgsRead**](OrgsOrgsApi.md#OrgsOrgsRead) | **Get** /orgs/orgs/{id}/ | orgs_orgs_read
[**OrgsOrgsUpdate**](OrgsOrgsApi.md#OrgsOrgsUpdate) | **Put** /orgs/orgs/{id}/ | orgs_orgs_update


# **OrgsOrgsCreate**
> Org OrgsOrgsCreate(ctx, data)
orgs_orgs_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Org**](Org.md)|  | 

### Return type

[**Org**](Org.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgsDelete**
> OrgsOrgsDelete(ctx, id)
orgs_orgs_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this 组织. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgsList**
> InlineResponse20027 OrgsOrgsList(ctx, optional)
orgs_orgs_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrgsOrgsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsOrgsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgsMembershipAdminsCreate**
> OrgMembershipAdmin OrgsOrgsMembershipAdminsCreate(ctx, orgId, data)
orgs_orgs_membership_admins_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | **string**|  | 
  **data** | [**OrgMembershipAdmin**](OrgMembershipAdmin.md)|  | 

### Return type

[**OrgMembershipAdmin**](OrgMembershipAdmin.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgsMembershipAdminsDelete**
> OrgsOrgsMembershipAdminsDelete(ctx, orgId, userId)
orgs_orgs_membership_admins_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgsMembershipAdminsList**
> InlineResponse20028 OrgsOrgsMembershipAdminsList(ctx, orgId, optional)
orgs_orgs_membership_admins_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | **string**|  | 
 **optional** | ***OrgsOrgsMembershipAdminsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsOrgsMembershipAdminsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgsMembershipAdminsRead**
> OrgMembershipAdmin OrgsOrgsMembershipAdminsRead(ctx, orgId, userId)
orgs_orgs_membership_admins_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**OrgMembershipAdmin**](OrgMembershipAdmin.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgsMembershipUsersCreate**
> OrgMembershipUser OrgsOrgsMembershipUsersCreate(ctx, orgId, data)
orgs_orgs_membership_users_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | **string**|  | 
  **data** | [**OrgMembershipUser**](OrgMembershipUser.md)|  | 

### Return type

[**OrgMembershipUser**](OrgMembershipUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgsMembershipUsersDelete**
> OrgsOrgsMembershipUsersDelete(ctx, orgId, userId)
orgs_orgs_membership_users_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgsMembershipUsersList**
> InlineResponse20029 OrgsOrgsMembershipUsersList(ctx, orgId, optional)
orgs_orgs_membership_users_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | **string**|  | 
 **optional** | ***OrgsOrgsMembershipUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsOrgsMembershipUsersListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgsMembershipUsersRead**
> OrgMembershipUser OrgsOrgsMembershipUsersRead(ctx, orgId, userId)
orgs_orgs_membership_users_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**OrgMembershipUser**](OrgMembershipUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgsPartialUpdate**
> Org OrgsOrgsPartialUpdate(ctx, id, data)
orgs_orgs_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this 组织. | 
  **data** | [**Org**](Org.md)|  | 

### Return type

[**Org**](Org.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgsRead**
> OrgRead OrgsOrgsRead(ctx, id)
orgs_orgs_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this 组织. | 

### Return type

[**OrgRead**](OrgRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgsUpdate**
> Org OrgsOrgsUpdate(ctx, id, data)
orgs_orgs_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this 组织. | 
  **data** | [**Org**](Org.md)|  | 

### Return type

[**Org**](Org.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

