# CreateSinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A human readable description for the Sink **This value should not contain PII.** | 
**SinkConfiguration** | [**map[string]interface{}**](.md) | The information required for Twilio to connect to the provided Sink encoded as JSON. | 
**SinkType** | **string** | The Sink type. Can only be \&quot;kinesis\&quot; or \&quot;webhook\&quot; currently. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

