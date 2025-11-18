# UploadStatsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SentAt** | **time.Time** |  | 
**UploadTime** | **int32** | the number of nanoseconds it took to upload a video | 

## Methods

### NewUploadStatsRequest

`func NewUploadStatsRequest(sentAt time.Time, uploadTime int32, ) *UploadStatsRequest`

NewUploadStatsRequest instantiates a new UploadStatsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadStatsRequestWithDefaults

`func NewUploadStatsRequestWithDefaults() *UploadStatsRequest`

NewUploadStatsRequestWithDefaults instantiates a new UploadStatsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSentAt

`func (o *UploadStatsRequest) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *UploadStatsRequest) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *UploadStatsRequest) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.


### GetUploadTime

`func (o *UploadStatsRequest) GetUploadTime() int32`

GetUploadTime returns the UploadTime field if non-nil, zero value otherwise.

### GetUploadTimeOk

`func (o *UploadStatsRequest) GetUploadTimeOk() (*int32, bool)`

GetUploadTimeOk returns a tuple with the UploadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTime

`func (o *UploadStatsRequest) SetUploadTime(v int32)`

SetUploadTime sets UploadTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


