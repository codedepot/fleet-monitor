# \DefaultAPI

All URIs are relative to *http://127.0.0.1:6733/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesDeviceIdHeartbeatPost**](DefaultAPI.md#DevicesDeviceIdHeartbeatPost) | **Post** /devices/{device_id}/heartbeat | 
[**DevicesDeviceIdStatsGet**](DefaultAPI.md#DevicesDeviceIdStatsGet) | **Get** /devices/{device_id}/stats | 
[**DevicesDeviceIdStatsPost**](DefaultAPI.md#DevicesDeviceIdStatsPost) | **Post** /devices/{device_id}/stats | 



## DevicesDeviceIdHeartbeatPost

> DevicesDeviceIdHeartbeatPost(ctx, deviceId).HeartbeatRequest(heartbeatRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/codedepot/fleet-monitor/client"
)

func main() {
	deviceId := "deviceId_example" // string | ID of a device to register heartbeat with
	heartbeatRequest := *openapiclient.NewHeartbeatRequest(time.Now()) // HeartbeatRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DevicesDeviceIdHeartbeatPost(context.Background(), deviceId).HeartbeatRequest(heartbeatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DevicesDeviceIdHeartbeatPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | ID of a device to register heartbeat with | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceIdHeartbeatPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **heartbeatRequest** | [**HeartbeatRequest**](HeartbeatRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesDeviceIdStatsGet

> GetDeviceStatsResponse DevicesDeviceIdStatsGet(ctx, deviceId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/codedepot/fleet-monitor/client"
)

func main() {
	deviceId := "deviceId_example" // string | ID of a device to register heartbeat with

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DevicesDeviceIdStatsGet(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DevicesDeviceIdStatsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesDeviceIdStatsGet`: GetDeviceStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DevicesDeviceIdStatsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | ID of a device to register heartbeat with | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceIdStatsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceStatsResponse**](GetDeviceStatsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesDeviceIdStatsPost

> DevicesDeviceIdStatsPost(ctx, deviceId).UploadStatsRequest(uploadStatsRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/codedepot/fleet-monitor/client"
)

func main() {
	deviceId := "deviceId_example" // string | ID of a device to register heartbeat with
	uploadStatsRequest := *openapiclient.NewUploadStatsRequest(time.Now(), int32(123)) // UploadStatsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DevicesDeviceIdStatsPost(context.Background(), deviceId).UploadStatsRequest(uploadStatsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DevicesDeviceIdStatsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | ID of a device to register heartbeat with | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceIdStatsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uploadStatsRequest** | [**UploadStatsRequest**](UploadStatsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

