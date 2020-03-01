# \TicketsTicketsApi

All URIs are relative to *http://jumpserver.test.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TicketsTicketsCommentsCreate**](TicketsTicketsApi.md#TicketsTicketsCommentsCreate) | **Post** /tickets/tickets/{ticket_id}/comments/ | tickets_tickets_comments_create
[**TicketsTicketsCommentsList**](TicketsTicketsApi.md#TicketsTicketsCommentsList) | **Get** /tickets/tickets/{ticket_id}/comments/ | tickets_tickets_comments_list
[**TicketsTicketsCommentsRead**](TicketsTicketsApi.md#TicketsTicketsCommentsRead) | **Get** /tickets/tickets/{ticket_id}/comments/{id}/ | tickets_tickets_comments_read
[**TicketsTicketsCreate**](TicketsTicketsApi.md#TicketsTicketsCreate) | **Post** /tickets/tickets/ | tickets_tickets_create
[**TicketsTicketsDelete**](TicketsTicketsApi.md#TicketsTicketsDelete) | **Delete** /tickets/tickets/{id}/ | tickets_tickets_delete
[**TicketsTicketsList**](TicketsTicketsApi.md#TicketsTicketsList) | **Get** /tickets/tickets/ | tickets_tickets_list
[**TicketsTicketsPartialUpdate**](TicketsTicketsApi.md#TicketsTicketsPartialUpdate) | **Patch** /tickets/tickets/{id}/ | tickets_tickets_partial_update
[**TicketsTicketsRead**](TicketsTicketsApi.md#TicketsTicketsRead) | **Get** /tickets/tickets/{id}/ | tickets_tickets_read
[**TicketsTicketsUpdate**](TicketsTicketsApi.md#TicketsTicketsUpdate) | **Put** /tickets/tickets/{id}/ | tickets_tickets_update


# **TicketsTicketsCommentsCreate**
> TicketsTicketsCommentsCreate(ctx, ticketId)
tickets_tickets_comments_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ticketId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TicketsTicketsCommentsList**
> TicketsTicketsCommentsList(ctx, ticketId, optional)
tickets_tickets_comments_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ticketId** | **string**|  | 
 **optional** | ***TicketsTicketsCommentsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TicketsTicketsCommentsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TicketsTicketsCommentsRead**
> TicketsTicketsCommentsRead(ctx, id, ticketId)
tickets_tickets_comments_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **ticketId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TicketsTicketsCreate**
> Ticket TicketsTicketsCreate(ctx, data)
tickets_tickets_create



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Ticket**](Ticket.md)|  | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TicketsTicketsDelete**
> TicketsTicketsDelete(ctx, id)
tickets_tickets_delete



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this ticket. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TicketsTicketsList**
> InlineResponse20059 TicketsTicketsList(ctx, optional)
tickets_tickets_list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TicketsTicketsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TicketsTicketsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**|  | 
 **title** | **optional.String**|  | 
 **action** | **optional.String**|  | 
 **userDisplay** | **optional.String**|  | 
 **search** | **optional.String**| A search term. | 
 **order** | **optional.String**| Which field to use when ordering the results. | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20059**](inline_response_200_59.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TicketsTicketsPartialUpdate**
> Ticket TicketsTicketsPartialUpdate(ctx, id, data)
tickets_tickets_partial_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this ticket. | 
  **data** | [**Ticket**](Ticket.md)|  | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TicketsTicketsRead**
> Ticket TicketsTicketsRead(ctx, id)
tickets_tickets_read



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this ticket. | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TicketsTicketsUpdate**
> Ticket TicketsTicketsUpdate(ctx, id, data)
tickets_tickets_update



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this ticket. | 
  **data** | [**Ticket**](Ticket.md)|  | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/csv, */*
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

