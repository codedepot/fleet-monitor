# GetDeviceStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgUploadTime** | **string** | returned as a time duration string. Eg: 5m10s | 
**Uptime** | **float64** | Uptime as a percentage. eg: 98.999 | 

## Methods

### NewGetDeviceStatsResponse

`func NewGetDeviceStatsResponse(avgUploadTime string, uptime float64, ) *GetDeviceStatsResponse`

NewGetDeviceStatsResponse instantiates a new GetDeviceStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceStatsResponseWithDefaults

`func NewGetDeviceStatsResponseWithDefaults() *GetDeviceStatsResponse`

NewGetDeviceStatsResponseWithDefaults instantiates a new GetDeviceStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgUploadTime

`func (o *GetDeviceStatsResponse) GetAvgUploadTime() string`

GetAvgUploadTime returns the AvgUploadTime field if non-nil, zero value otherwise.

### GetAvgUploadTimeOk

`func (o *GetDeviceStatsResponse) GetAvgUploadTimeOk() (*string, bool)`

GetAvgUploadTimeOk returns a tuple with the AvgUploadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgUploadTime

`func (o *GetDeviceStatsResponse) SetAvgUploadTime(v string)`

SetAvgUploadTime sets AvgUploadTime field to given value.


### GetUptime

`func (o *GetDeviceStatsResponse) GetUptime() float64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *GetDeviceStatsResponse) GetUptimeOk() (*float64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *GetDeviceStatsResponse) SetUptime(v float64)`

SetUptime sets Uptime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


