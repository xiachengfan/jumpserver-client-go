# \UsersUsersApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersUsersBulkDelete**](UsersUsersApi.md#UsersUsersBulkDelete) | **Delete** /users/users/ | users_users_bulk_delete
[**UsersUsersBulkUpdate**](UsersUsersApi.md#UsersUsersBulkUpdate) | **Put** /users/users/ | users_users_bulk_update
[**UsersUsersCreate**](UsersUsersApi.md#UsersUsersCreate) | **Post** /users/users/ | users_users_create
[**UsersUsersDelete**](UsersUsersApi.md#UsersUsersDelete) | **Delete** /users/users/{id}/ | users_users_delete
[**UsersUsersList**](UsersUsersApi.md#UsersUsersList) | **Get** /users/users/ | users_users_list
[**UsersUsersOtpResetRead**](UsersUsersApi.md#UsersUsersOtpResetRead) | **Get** /users/users/{id}/otp/reset/ | users_users_otp_reset_read
[**UsersUsersPartialBulkUpdate**](UsersUsersApi.md#UsersUsersPartialBulkUpdate) | **Patch** /users/users/ | users_users_partial_bulk_update
[**UsersUsersPartialUpdate**](UsersUsersApi.md#UsersUsersPartialUpdate) | **Patch** /users/users/{id}/ | users_users_partial_update
[**UsersUsersPasswordPartialUpdate**](UsersUsersApi.md#UsersUsersPasswordPartialUpdate) | **Patch** /users/users/{id}/password/ | users_users_password_partial_update
[**UsersUsersPasswordRead**](UsersUsersApi.md#UsersUsersPasswordRead) | **Get** /users/users/{id}/password/ | users_users_password_read
[**UsersUsersPasswordResetPartialUpdate**](UsersUsersApi.md#UsersUsersPasswordResetPartialUpdate) | **Patch** /users/users/{id}/password/reset/ | users_users_password_reset_partial_update
[**UsersUsersPasswordResetUpdate**](UsersUsersApi.md#UsersUsersPasswordResetUpdate) | **Put** /users/users/{id}/password/reset/ | users_users_password_reset_update
[**UsersUsersPasswordUpdate**](UsersUsersApi.md#UsersUsersPasswordUpdate) | **Put** /users/users/{id}/password/ | users_users_password_update
[**UsersUsersPubkeyResetPartialUpdate**](UsersUsersApi.md#UsersUsersPubkeyResetPartialUpdate) | **Patch** /users/users/{id}/pubkey/reset/ | users_users_pubkey_reset_partial_update
[**UsersUsersPubkeyResetUpdate**](UsersUsersApi.md#UsersUsersPubkeyResetUpdate) | **Put** /users/users/{id}/pubkey/reset/ | users_users_pubkey_reset_update
[**UsersUsersPubkeyUpdatePartialUpdate**](UsersUsersApi.md#UsersUsersPubkeyUpdatePartialUpdate) | **Patch** /users/users/{id}/pubkey/update/ | users_users_pubkey_update_partial_update
[**UsersUsersPubkeyUpdateUpdate**](UsersUsersApi.md#UsersUsersPubkeyUpdateUpdate) | **Put** /users/users/{id}/pubkey/update/ | users_users_pubkey_update_update
[**UsersUsersRead**](UsersUsersApi.md#UsersUsersRead) | **Get** /users/users/{id}/ | users_users_read
[**UsersUsersUnblockPartialUpdate**](UsersUsersApi.md#UsersUsersUnblockPartialUpdate) | **Patch** /users/users/{id}/unblock/ | users_users_unblock_partial_update
[**UsersUsersUnblockUpdate**](UsersUsersApi.md#UsersUsersUnblockUpdate) | **Put** /users/users/{id}/unblock/ | users_users_unblock_update
[**UsersUsersUpdate**](UsersUsersApi.md#UsersUsersUpdate) | **Put** /users/users/{id}/ | users_users_update


# **UsersUsersBulkDelete**
> UsersUsersBulkDelete(ctx, optional)
users_users_bulk_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersUsersBulkDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersUsersBulkDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**|  | 
 **email** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **id** | **optional.String**|  | 
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

# **UsersUsersBulkUpdate**
> User UsersUsersBulkUpdate(ctx, data)
users_users_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersCreate**
> User UsersUsersCreate(ctx, data)
users_users_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersDelete**
> UsersUsersDelete(ctx, id)
users_users_delete



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

# **UsersUsersList**
> InlineResponse20062 UsersUsersList(ctx, optional)
users_users_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersUsersListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**|  | 
 **email** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **id** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **spm** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20062**](inline_response_200_62.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersOtpResetRead**
> ResetOtp UsersUsersOtpResetRead(ctx, id)
users_users_otp_reset_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ResetOtp**](ResetOTP.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersPartialBulkUpdate**
> User UsersUsersPartialBulkUpdate(ctx, data)
users_users_partial_bulk_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersPartialUpdate**
> User UsersUsersPartialUpdate(ctx, id, data)
users_users_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersPasswordPartialUpdate**
> ChangeUserPassword UsersUsersPasswordPartialUpdate(ctx, id, data)
users_users_password_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**ChangeUserPassword**](ChangeUserPassword.md)|  | 

### Return type

[**ChangeUserPassword**](ChangeUserPassword.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersPasswordRead**
> ChangeUserPassword UsersUsersPasswordRead(ctx, id)
users_users_password_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ChangeUserPassword**](ChangeUserPassword.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersPasswordResetPartialUpdate**
> User UsersUsersPasswordResetPartialUpdate(ctx, id, data)
users_users_password_reset_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this 用户. | 
  **data** | [**User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersPasswordResetUpdate**
> User UsersUsersPasswordResetUpdate(ctx, id, data)
users_users_password_reset_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this 用户. | 
  **data** | [**User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersPasswordUpdate**
> ChangeUserPassword UsersUsersPasswordUpdate(ctx, id, data)
users_users_password_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**ChangeUserPassword**](ChangeUserPassword.md)|  | 

### Return type

[**ChangeUserPassword**](ChangeUserPassword.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersPubkeyResetPartialUpdate**
> User UsersUsersPubkeyResetPartialUpdate(ctx, id, data)
users_users_pubkey_reset_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersPubkeyResetUpdate**
> User UsersUsersPubkeyResetUpdate(ctx, id, data)
users_users_pubkey_reset_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersPubkeyUpdatePartialUpdate**
> UserPkUpdate UsersUsersPubkeyUpdatePartialUpdate(ctx, id, data)
users_users_pubkey_update_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**UserPkUpdate**](UserPkUpdate.md)|  | 

### Return type

[**UserPkUpdate**](UserPKUpdate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersPubkeyUpdateUpdate**
> UserPkUpdate UsersUsersPubkeyUpdateUpdate(ctx, id, data)
users_users_pubkey_update_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**UserPkUpdate**](UserPkUpdate.md)|  | 

### Return type

[**UserPkUpdate**](UserPKUpdate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersRead**
> User UsersUsersRead(ctx, id)
users_users_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersUnblockPartialUpdate**
> User UsersUsersUnblockPartialUpdate(ctx, id, data)
users_users_unblock_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersUnblockUpdate**
> User UsersUsersUnblockUpdate(ctx, id, data)
users_users_unblock_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersUpdate**
> User UsersUsersUpdate(ctx, id, data)
users_users_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

