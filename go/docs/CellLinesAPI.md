# \CellLinesAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CellLinesGet**](CellLinesAPI.md#CellLinesGet) | **Get** /api/v1/cell-lines/{id} | One cell-lines record
[**CellLinesIdentity**](CellLinesAPI.md#CellLinesIdentity) | **Get** /api/v1/cell-lines/{id}/identity | Identity and availability
[**CellLinesList**](CellLinesAPI.md#CellLinesList) | **Get** /api/v1/cell-lines | Cell line identity safety
[**CellLinesSafetyResolution**](CellLinesAPI.md#CellLinesSafetyResolution) | **Get** /api/v1/cell-lines/{id}/safety_resolution | Safety resolution
[**CellLinesStatusCounts**](CellLinesAPI.md#CellLinesStatusCounts) | **Get** /api/v1/cell-lines/status_counts | Cell line safety summary



## CellLinesGet

> CellLinesGet200Response CellLinesGet(ctx, id).Execute()

One cell-lines record

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
	id := "SCI-CONF" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CellLinesAPI.CellLinesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellLinesAPI.CellLinesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CellLinesGet`: CellLinesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CellLinesAPI.CellLinesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCellLinesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CellLinesGet200Response**](CellLinesGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CellLinesIdentity

> CellLinesIdentity200Response CellLinesIdentity(ctx, id).Execute()

Identity and availability

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
	id := "SCI-CONF" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CellLinesAPI.CellLinesIdentity(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellLinesAPI.CellLinesIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CellLinesIdentity`: CellLinesIdentity200Response
	fmt.Fprintf(os.Stdout, "Response from `CellLinesAPI.CellLinesIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCellLinesIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CellLinesIdentity200Response**](CellLinesIdentity200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CellLinesList

> CellLinesList200Response CellLinesList(ctx).Cursor(cursor).Facets(facets).Limit(limit).SafetyStatus(safetyStatus).Execute()

Cell line identity safety

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
	cursor := "eyJwIjoyLCJsIjoyLCJmIjoic3BlYy1maW5nZXJwcmludCJ9" // string |  (optional)
	facets := true // bool |  (optional)
	limit := int32(1) // int32 |  (optional)
	safetyStatus := "confirmed_problematic" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CellLinesAPI.CellLinesList(context.Background()).Cursor(cursor).Facets(facets).Limit(limit).SafetyStatus(safetyStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellLinesAPI.CellLinesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CellLinesList`: CellLinesList200Response
	fmt.Fprintf(os.Stdout, "Response from `CellLinesAPI.CellLinesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCellLinesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  | 
 **facets** | **bool** |  | 
 **limit** | **int32** |  | 
 **safetyStatus** | **string** |  | 

### Return type

[**CellLinesList200Response**](CellLinesList200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CellLinesSafetyResolution

> CellLinesSafetyResolution200Response CellLinesSafetyResolution(ctx, id).Execute()

Safety resolution

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
	id := "SCI-CONF" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CellLinesAPI.CellLinesSafetyResolution(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellLinesAPI.CellLinesSafetyResolution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CellLinesSafetyResolution`: CellLinesSafetyResolution200Response
	fmt.Fprintf(os.Stdout, "Response from `CellLinesAPI.CellLinesSafetyResolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCellLinesSafetyResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CellLinesSafetyResolution200Response**](CellLinesSafetyResolution200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CellLinesStatusCounts

> CellLinesStatusCounts200Response CellLinesStatusCounts(ctx).Species(species).Execute()

Cell line safety summary

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
	species := "Mouse" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CellLinesAPI.CellLinesStatusCounts(context.Background()).Species(species).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellLinesAPI.CellLinesStatusCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CellLinesStatusCounts`: CellLinesStatusCounts200Response
	fmt.Fprintf(os.Stdout, "Response from `CellLinesAPI.CellLinesStatusCounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCellLinesStatusCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **species** | **string** |  | 

### Return type

[**CellLinesStatusCounts200Response**](CellLinesStatusCounts200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

