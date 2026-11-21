# \RwdUsNetworksAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RwdUsNetworksGet**](RwdUsNetworksAPI.md#RwdUsNetworksGet) | **Get** /api/v1/rwd_us_networks/{id} | One us real-world data network
[**RwdUsNetworksList**](RwdUsNetworksAPI.md#RwdUsNetworksList) | **Get** /api/v1/rwd_us_networks | List us real-world data networks
[**RwdUsNetworksRwdUsNetworkMemberships**](RwdUsNetworksAPI.md#RwdUsNetworksRwdUsNetworkMemberships) | **Get** /api/v1/rwd_us_networks/{id}/rwd_us_network_memberships | US real-world data network membership (N3C)
[**RwdUsNetworksRwdUsNetworkProjects**](RwdUsNetworksAPI.md#RwdUsNetworksRwdUsNetworkProjects) | **Get** /api/v1/rwd_us_networks/{id}/rwd_us_network_projects | US real-world data network projects (N3C)



## RwdUsNetworksGet

> RwdUsNetworksRecord RwdUsNetworksGet(ctx, id).Execute()

One us real-world data network

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
	id := "n3c" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdUsNetworksAPI.RwdUsNetworksGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdUsNetworksAPI.RwdUsNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdUsNetworksGet`: RwdUsNetworksRecord
	fmt.Fprintf(os.Stdout, "Response from `RwdUsNetworksAPI.RwdUsNetworksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRwdUsNetworksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RwdUsNetworksRecord**](RwdUsNetworksRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RwdUsNetworksList

> RwdUsNetworksResponse RwdUsNetworksList(ctx).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

List us real-world data networks

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
	resp, r, err := apiClient.RwdUsNetworksAPI.RwdUsNetworksList(context.Background()).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdUsNetworksAPI.RwdUsNetworksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdUsNetworksList`: RwdUsNetworksResponse
	fmt.Fprintf(os.Stdout, "Response from `RwdUsNetworksAPI.RwdUsNetworksList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRwdUsNetworksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**RwdUsNetworksResponse**](RwdUsNetworksResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RwdUsNetworksRwdUsNetworkMemberships

> RwdUsNetworkMembershipsResponse RwdUsNetworksRwdUsNetworkMemberships(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

US real-world data network membership (N3C)

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
	id := "n3c" // string | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdUsNetworksAPI.RwdUsNetworksRwdUsNetworkMemberships(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdUsNetworksAPI.RwdUsNetworksRwdUsNetworkMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdUsNetworksRwdUsNetworkMemberships`: RwdUsNetworkMembershipsResponse
	fmt.Fprintf(os.Stdout, "Response from `RwdUsNetworksAPI.RwdUsNetworksRwdUsNetworkMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRwdUsNetworksRwdUsNetworkMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**RwdUsNetworkMembershipsResponse**](RwdUsNetworkMembershipsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RwdUsNetworksRwdUsNetworkProjects

> RwdUsNetworkProjectsResponse RwdUsNetworksRwdUsNetworkProjects(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

US real-world data network projects (N3C)

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
	id := "n3c" // string | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdUsNetworksAPI.RwdUsNetworksRwdUsNetworkProjects(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdUsNetworksAPI.RwdUsNetworksRwdUsNetworkProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdUsNetworksRwdUsNetworkProjects`: RwdUsNetworkProjectsResponse
	fmt.Fprintf(os.Stdout, "Response from `RwdUsNetworksAPI.RwdUsNetworksRwdUsNetworkProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRwdUsNetworksRwdUsNetworkProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**RwdUsNetworkProjectsResponse**](RwdUsNetworkProjectsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

