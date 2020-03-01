# \SettingsLdapApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsLdapCacheRefreshRead**](SettingsLdapApi.md#SettingsLdapCacheRefreshRead) | **Get** /settings/ldap/cache/refresh/ | settings_ldap_cache_refresh_read
[**SettingsLdapTestingCreate**](SettingsLdapApi.md#SettingsLdapTestingCreate) | **Post** /settings/ldap/testing/ | settings_ldap_testing_create
[**SettingsLdapUsersImportCreate**](SettingsLdapApi.md#SettingsLdapUsersImportCreate) | **Post** /settings/ldap/users/import/ | settings_ldap_users_import_create
[**SettingsLdapUsersList**](SettingsLdapApi.md#SettingsLdapUsersList) | **Get** /settings/ldap/users/ | settings_ldap_users_list


# **SettingsLdapCacheRefreshRead**
> SettingsLdapCacheRefreshRead(ctx, )
settings_ldap_cache_refresh_read



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

# **SettingsLdapTestingCreate**
> SettingsLdapTestingCreate(ctx, )
settings_ldap_testing_create



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

# **SettingsLdapUsersImportCreate**
> SettingsLdapUsersImportCreate(ctx, )
settings_ldap_users_import_create



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

# **SettingsLdapUsersList**
> InlineResponse20052 SettingsLdapUsersList(ctx, optional)
settings_ldap_users_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettingsLdapUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsLdapUsersListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20052**](inline_response_200_52.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

