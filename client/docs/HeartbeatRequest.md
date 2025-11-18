# HeartbeatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SentAt** | **time.Time** |  | 

## Methods

### NewHeartbeatRequest

`func NewHeartbeatRequest(sentAt time.Time, ) *HeartbeatRequest`

NewHeartbeatRequest instantiates a new HeartbeatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeartbeatRequestWithDefaults

`func NewHeartbeatRequestWithDefaults() *HeartbeatRequest`

NewHeartbeatRequestWithDefaults instantiates a new HeartbeatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSentAt

`func (o *HeartbeatRequest) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *HeartbeatRequest) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *HeartbeatRequest) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


