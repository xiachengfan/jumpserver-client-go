# \AuthenticationAuthApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationAuthCreate**](AuthenticationAuthApi.md#AuthenticationAuthCreate) | **Post** /authentication/auth/ | authentication_auth_create


# **AuthenticationAuthCreate**
> BearerToken AuthenticationAuthCreate(ctx, data)
authentication_auth_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BearerToken**](BearerToken.md)|  | 

### Return type

[**BearerToken**](BearerToken.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

