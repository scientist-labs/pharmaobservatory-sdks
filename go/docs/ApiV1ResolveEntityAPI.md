# \ApiV1ResolveEntityAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ResolveEntityGet**](ApiV1ResolveEntityAPI.md#ApiV1ResolveEntityGet) | **Get** /api/v1/resolve_entity | Resolve a free-text or coded entity mention to a canonical Pharma Observatory ID



## ApiV1ResolveEntityGet

> ApiV1ResolveEntityGet200Response ApiV1ResolveEntityGet(ctx).Query(query).RawKey(rawKey).SourceType(sourceType).Spine(spine).Execute()

Resolve a free-text or coded entity mention to a canonical Pharma Observatory ID

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
	query := "Anchorloss Biotherapeutics" // string |  (optional)
	rawKey := "biomarin" // string |  (optional)
	sourceType := "aact_sponsor" // string |  (optional)
	spine := "institution" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1ResolveEntityAPI.ApiV1ResolveEntityGet(context.Background()).Query(query).RawKey(rawKey).SourceType(sourceType).Spine(spine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1ResolveEntityAPI.ApiV1ResolveEntityGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ResolveEntityGet`: ApiV1ResolveEntityGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1ResolveEntityAPI.ApiV1ResolveEntityGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ResolveEntityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **rawKey** | **string** |  | 
 **sourceType** | **string** |  | 
 **spine** | **string** |  | 

### Return type

[**ApiV1ResolveEntityGet200Response**](ApiV1ResolveEntityGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

