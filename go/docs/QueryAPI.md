# \QueryAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryResult**](QueryAPI.md#QueryResult) | **Get** /api/v1/results/{result_id} | Page a stored result
[**QueryRun**](QueryAPI.md#QueryRun) | **Post** /api/v1/query | Run a SQL query



## QueryResult

> ResultPage QueryResult(ctx, resultId).Cursor(cursor).Execute()

Page a stored result

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
	resultId := "5d3f6a1e-0c1b-4f7e-9a52-2b8c1d4e6f70" // string | 
	cursor := "eyJvIjoyMDB9" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryAPI.QueryResult(context.Background(), resultId).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryAPI.QueryResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryResult`: ResultPage
	fmt.Fprintf(os.Stdout, "Response from `QueryAPI.QueryResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** |  | 

### Return type

[**ResultPage**](ResultPage.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryRun

> QueryResponse QueryRun(ctx).QueryRunRequest(queryRunRequest).Execute()

Run a SQL query

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
	queryRunRequest := *openapiclient.NewQueryRunRequest("Sql_example") // QueryRunRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryAPI.QueryRun(context.Background()).QueryRunRequest(queryRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryAPI.QueryRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryRun`: QueryResponse
	fmt.Fprintf(os.Stdout, "Response from `QueryAPI.QueryRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryRunRequest** | [**QueryRunRequest**](QueryRunRequest.md) |  | 

### Return type

[**QueryResponse**](QueryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

