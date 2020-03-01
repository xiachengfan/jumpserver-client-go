# \AuthenticationMfaApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationMfaChallengeCreate**](AuthenticationMfaApi.md#AuthenticationMfaChallengeCreate) | **Post** /authentication/mfa/challenge/ | authentication_mfa_challenge_create


# **AuthenticationMfaChallengeCreate**
> MfaChallenge AuthenticationMfaChallengeCreate(ctx, data)
authentication_mfa_challenge_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MfaChallenge**](MfaChallenge.md)|  | 

### Return type

[**MfaChallenge**](MFAChallenge.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

