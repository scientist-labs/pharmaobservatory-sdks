# \RwdNetworksAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RwdNetworksGet**](RwdNetworksAPI.md#RwdNetworksGet) | **Get** /api/v1/rwd_networks/{id} | One real-world data register network
[**RwdNetworksList**](RwdNetworksAPI.md#RwdNetworksList) | **Get** /api/v1/rwd_networks | List real-world data register networks



## RwdNetworksGet

> RwdNetworksRecord RwdNetworksGet(ctx, id).Execute()

One real-world data register network

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/scientist-labs/pharmaobservatory-sdks/go"
)

func main() {
	id := int32(900000101) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdNetworksAPI.RwdNetworksGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdNetworksAPI.RwdNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdNetworksGet`: RwdNetworksRecord
	fmt.Fprintf(os.Stdout, "Response from `RwdNetworksAPI.RwdNetworksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRwdNetworksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RwdNetworksRecord**](RwdNetworksRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RwdNetworksList

> RwdNetworksResponse RwdNetworksList(ctx).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

List real-world data register networks

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/scientist-labs/pharmaobservatory-sdks/go"
)

func main() {
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdNetworksAPI.RwdNetworksList(context.Background()).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdNetworksAPI.RwdNetworksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdNetworksList`: RwdNetworksResponse
	fmt.Fprintf(os.Stdout, "Response from `RwdNetworksAPI.RwdNetworksList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRwdNetworksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**RwdNetworksResponse**](RwdNetworksResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

