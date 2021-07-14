# ServicesUsersConversationsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceUserConversation**](ServicesUsersConversationsApi.md#DeleteServiceUserConversation) | **Delete** /v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid} | 
[**FetchServiceUserConversation**](ServicesUsersConversationsApi.md#FetchServiceUserConversation) | **Get** /v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid} | 
[**ListServiceUserConversation**](ServicesUsersConversationsApi.md#ListServiceUserConversation) | **Get** /v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations | 
[**UpdateServiceUserConversation**](ServicesUsersConversationsApi.md#UpdateServiceUserConversation) | **Post** /v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid} | 



## DeleteServiceUserConversation

> DeleteServiceUserConversation(ctx, ChatServiceSidUserSidConversationSid)



Delete a specific User Conversation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**UserSid** | **string** | The unique SID identifier of the [User resource](https://www.twilio.com/docs/conversations/api/user-resource). This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource.
**ConversationSid** | **string** | The unique SID identifier of the Conversation. This value can be either the &#x60;sid&#x60; or the &#x60;unique_name&#x60; of the [Conversation resource](https://www.twilio.com/docs/conversations/api/conversation-resource).

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceUserConversationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceUserConversation

> ConversationsV1ServiceServiceUserServiceUserConversation FetchServiceUserConversation(ctx, ChatServiceSidUserSidConversationSid)



Fetch a specific User Conversation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**UserSid** | **string** | The unique SID identifier of the [User resource](https://www.twilio.com/docs/conversations/api/user-resource). This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource.
**ConversationSid** | **string** | The unique SID identifier of the Conversation. This value can be either the &#x60;sid&#x60; or the &#x60;unique_name&#x60; of the [Conversation resource](https://www.twilio.com/docs/conversations/api/conversation-resource).

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceUserConversationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceServiceUserServiceUserConversation**](ConversationsV1ServiceServiceUserServiceUserConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceUserConversation

> ListServiceUserConversationResponse ListServiceUserConversation(ctx, ChatServiceSidUserSidoptional)



Retrieve a list of all User Conversations for the User.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**UserSid** | **string** | The unique SID identifier of the [User resource](https://www.twilio.com/docs/conversations/api/user-resource). This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceUserConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceUserConversationResponse**](ListServiceUserConversationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceUserConversation

> ConversationsV1ServiceServiceUserServiceUserConversation UpdateServiceUserConversation(ctx, ChatServiceSidUserSidConversationSidoptional)



Update a specific User Conversation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**UserSid** | **string** | The unique SID identifier of the [User resource](https://www.twilio.com/docs/conversations/api/user-resource). This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource.
**ConversationSid** | **string** | The unique SID identifier of the Conversation. This value can be either the &#x60;sid&#x60; or the &#x60;unique_name&#x60; of the [Conversation resource](https://www.twilio.com/docs/conversations/api/conversation-resource).

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceUserConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**LastReadMessageIndex** | **int** | The index of the last Message in the Conversation that the Participant has read.
**LastReadTimestamp** | **time.Time** | The date of the last message read in conversation by the user, given in ISO 8601 format.
**NotificationLevel** | **string** | The Notification Level of this User Conversation. One of &#x60;default&#x60; or &#x60;muted&#x60;.

### Return type

[**ConversationsV1ServiceServiceUserServiceUserConversation**](ConversationsV1ServiceServiceUserServiceUserConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
