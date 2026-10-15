# \ApiV1SiteRecordAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1SiteRecordGet**](ApiV1SiteRecordAPI.md#ApiV1SiteRecordGet) | **Get** /api/v1/site_record | Look up a clinical trial site&#39;s full record, including hosted trials and named investigators



## ApiV1SiteRecordGet

> ApiV1SiteRecordGet200Response ApiV1SiteRecordGet(ctx).InvestigatorId(investigatorId).Limit(limit).OrgId(orgId).SiteId(siteId).Execute()

Look up a clinical trial site's full record, including hosted trials and named investigators

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
	investigatorId := "aact:e7c2d751838e731e6898645b06a70eb1" // string |  (optional)
	limit := int32(1) // int32 |  (optional)
	orgId := "edgar:0001060955" // string |  (optional)
	siteId := "87a040a71a6f82289add8872820314eb" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1SiteRecordAPI.ApiV1SiteRecordGet(context.Background()).InvestigatorId(investigatorId).Limit(limit).OrgId(orgId).SiteId(siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1SiteRecordAPI.ApiV1SiteRecordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1SiteRecordGet`: ApiV1SiteRecordGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1SiteRecordAPI.ApiV1SiteRecordGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1SiteRecordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investigatorId** | **string** |  | 
 **limit** | **int32** |  | 
 **orgId** | **string** |  | 
 **siteId** | **string** |  | 

### Return type

[**ApiV1SiteRecordGet200Response**](ApiV1SiteRecordGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

