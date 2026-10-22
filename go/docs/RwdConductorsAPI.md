# \RwdConductorsAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RwdConductorsGet**](RwdConductorsAPI.md#RwdConductorsGet) | **Get** /api/v1/rwd_conductors/{id} | One real-world data study conductor
[**RwdConductorsList**](RwdConductorsAPI.md#RwdConductorsList) | **Get** /api/v1/rwd_conductors | List real-world data study conductors
[**RwdConductorsRwdConductorDelivery**](RwdConductorsAPI.md#RwdConductorsRwdConductorDelivery) | **Get** /api/v1/rwd_conductors/{id}/rwd_conductor_delivery | Real-world data study conductor delivery by mandate (HMA-EMA catalogue)
[**RwdConductorsRwdConductorFootprint**](RwdConductorsAPI.md#RwdConductorsRwdConductorFootprint) | **Get** /api/v1/rwd_conductors/{id}/rwd_conductor_footprint | Real-world data study conductor footprint (HMA-EMA catalogue)
[**RwdConductorsRwdConductorStudies**](RwdConductorsAPI.md#RwdConductorsRwdConductorStudies) | **Get** /api/v1/rwd_conductors/{id}/rwd_conductor_studies | Real-world data study conductor studies (HMA-EMA catalogue)



## RwdConductorsGet

> RwdConductorsRecord RwdConductorsGet(ctx, id).Execute()

One real-world data study conductor

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
	id := "seedmark-consulting" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdConductorsAPI.RwdConductorsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdConductorsAPI.RwdConductorsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdConductorsGet`: RwdConductorsRecord
	fmt.Fprintf(os.Stdout, "Response from `RwdConductorsAPI.RwdConductorsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRwdConductorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RwdConductorsRecord**](RwdConductorsRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RwdConductorsList

> RwdConductorsResponse RwdConductorsList(ctx).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

List real-world data study conductors

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
	resp, r, err := apiClient.RwdConductorsAPI.RwdConductorsList(context.Background()).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdConductorsAPI.RwdConductorsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdConductorsList`: RwdConductorsResponse
	fmt.Fprintf(os.Stdout, "Response from `RwdConductorsAPI.RwdConductorsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRwdConductorsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**RwdConductorsResponse**](RwdConductorsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RwdConductorsRwdConductorDelivery

> RwdConductorDeliveryResponse RwdConductorsRwdConductorDelivery(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

Real-world data study conductor delivery by mandate (HMA-EMA catalogue)

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
	id := "seedmark-consulting" // string | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdConductorsAPI.RwdConductorsRwdConductorDelivery(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdConductorsAPI.RwdConductorsRwdConductorDelivery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdConductorsRwdConductorDelivery`: RwdConductorDeliveryResponse
	fmt.Fprintf(os.Stdout, "Response from `RwdConductorsAPI.RwdConductorsRwdConductorDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRwdConductorsRwdConductorDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**RwdConductorDeliveryResponse**](RwdConductorDeliveryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RwdConductorsRwdConductorFootprint

> RwdConductorFootprintResponse RwdConductorsRwdConductorFootprint(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

Real-world data study conductor footprint (HMA-EMA catalogue)

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
	id := "seedmark-consulting" // string | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdConductorsAPI.RwdConductorsRwdConductorFootprint(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdConductorsAPI.RwdConductorsRwdConductorFootprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdConductorsRwdConductorFootprint`: RwdConductorFootprintResponse
	fmt.Fprintf(os.Stdout, "Response from `RwdConductorsAPI.RwdConductorsRwdConductorFootprint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRwdConductorsRwdConductorFootprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**RwdConductorFootprintResponse**](RwdConductorFootprintResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RwdConductorsRwdConductorStudies

> RwdConductorStudiesResponse RwdConductorsRwdConductorStudies(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

Real-world data study conductor studies (HMA-EMA catalogue)

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
	id := "seedmark-consulting" // string | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdConductorsAPI.RwdConductorsRwdConductorStudies(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdConductorsAPI.RwdConductorsRwdConductorStudies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdConductorsRwdConductorStudies`: RwdConductorStudiesResponse
	fmt.Fprintf(os.Stdout, "Response from `RwdConductorsAPI.RwdConductorsRwdConductorStudies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRwdConductorsRwdConductorStudiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**RwdConductorStudiesResponse**](RwdConductorStudiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

