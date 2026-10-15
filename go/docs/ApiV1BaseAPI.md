# \ApiV1BaseAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1WhoamiGet**](ApiV1BaseAPI.md#ApiV1WhoamiGet) | **Get** /api/v1/whoami | whoami



## ApiV1WhoamiGet

> ApiV1WhoamiGet200Response ApiV1WhoamiGet(ctx).Execute()

whoami

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
	resp, r, err := apiClient.ApiV1BaseAPI.ApiV1WhoamiGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1BaseAPI.ApiV1WhoamiGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1WhoamiGet`: ApiV1WhoamiGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1BaseAPI.ApiV1WhoamiGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WhoamiGetRequest struct via the builder pattern


### Return type

[**ApiV1WhoamiGet200Response**](ApiV1WhoamiGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

