# \SubscriptionsAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangesList**](SubscriptionsAPI.md#ChangesList) | **Get** /api/v1/changes | Changes since a cursor
[**SubscriptionsCreate**](SubscriptionsAPI.md#SubscriptionsCreate) | **Post** /api/v1/subscriptions | Follow an entity or a dataset query
[**SubscriptionsDelete**](SubscriptionsAPI.md#SubscriptionsDelete) | **Delete** /api/v1/subscriptions/{id} | Stop following
[**SubscriptionsList**](SubscriptionsAPI.md#SubscriptionsList) | **Get** /api/v1/subscriptions | List your subscriptions



## ChangesList

> ChangesList200Response ChangesList(ctx).Cursor(cursor).Limit(limit).Execute()

Changes since a cursor

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
	cursor := "eyJ2IjoxLCJvIjoiMjAyNi0xMS0wMlQwNjowMjowMC4wMDAwMDBaIiwiaSI6OTAwMn0" // string |  (optional)
	limit := int32(2) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.ChangesList(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.ChangesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangesList`: ChangesList200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.ChangesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

[**ChangesList200Response**](ChangesList200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsCreate

> SubscriptionsCreate201Response SubscriptionsCreate(ctx).SubscriptionsCreateRequest(subscriptionsCreateRequest).Execute()

Follow an entity or a dataset query

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
	subscriptionsCreateRequest := *openapiclient.NewSubscriptionsCreateRequest("ObjectType_example", "EntityId_example") // SubscriptionsCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsCreate(context.Background()).SubscriptionsCreateRequest(subscriptionsCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsCreate`: SubscriptionsCreate201Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionsCreateRequest** | [**SubscriptionsCreateRequest**](SubscriptionsCreateRequest.md) |  | 

### Return type

[**SubscriptionsCreate201Response**](SubscriptionsCreate201Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsDelete

> SubscriptionsDelete200Response SubscriptionsDelete(ctx, id).Execute()

Stop following

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
	id := int32(4242) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsDelete`: SubscriptionsDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionsDelete200Response**](SubscriptionsDelete200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsList

> SubscriptionsList200Response SubscriptionsList(ctx).Execute()

List your subscriptions

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsList`: SubscriptionsList200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsListRequest struct via the builder pattern


### Return type

[**SubscriptionsList200Response**](SubscriptionsList200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

