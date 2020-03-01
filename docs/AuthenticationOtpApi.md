# \AuthenticationOtpApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationOtpVerifyCreate**](AuthenticationOtpApi.md#AuthenticationOtpVerifyCreate) | **Post** /authentication/otp/verify/ | authentication_otp_verify_create


# **AuthenticationOtpVerifyCreate**
> OtpVerify AuthenticationOtpVerifyCreate(ctx, data)
authentication_otp_verify_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**OtpVerify**](OtpVerify.md)|  | 

### Return type

[**OtpVerify**](OtpVerify.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

