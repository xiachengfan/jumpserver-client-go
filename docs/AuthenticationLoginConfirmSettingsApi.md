# \AuthenticationLoginConfirmSettingsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationLoginConfirmSettingsPartialUpdate**](AuthenticationLoginConfirmSettingsApi.md#AuthenticationLoginConfirmSettingsPartialUpdate) | **Patch** /authentication/login-confirm-settings/{user_id}/ | authentication_login-confirm-settings_partial_update
[**AuthenticationLoginConfirmSettingsUpdate**](AuthenticationLoginConfirmSettingsApi.md#AuthenticationLoginConfirmSettingsUpdate) | **Put** /authentication/login-confirm-settings/{user_id}/ | authentication_login-confirm-settings_update


# **AuthenticationLoginConfirmSettingsPartialUpdate**
> LoginConfirmSetting AuthenticationLoginConfirmSettingsPartialUpdate(ctx, userId, data)
authentication_login-confirm-settings_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **data** | [**LoginConfirmSetting**](LoginConfirmSetting.md)|  | 

### Return type

[**LoginConfirmSetting**](LoginConfirmSetting.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationLoginConfirmSettingsUpdate**
> LoginConfirmSetting AuthenticationLoginConfirmSettingsUpdate(ctx, userId, data)
authentication_login-confirm-settings_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **data** | [**LoginConfirmSetting**](LoginConfirmSetting.md)|  | 

### Return type

[**LoginConfirmSetting**](LoginConfirmSetting.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

