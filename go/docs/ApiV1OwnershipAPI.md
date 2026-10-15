# \ApiV1OwnershipAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1InvestorProfileGet**](ApiV1OwnershipAPI.md#ApiV1InvestorProfileGet) | **Get** /api/v1/investor_profile | investor_profile
[**ApiV1OrgFinancingAndOwnershipGet**](ApiV1OwnershipAPI.md#ApiV1OrgFinancingAndOwnershipGet) | **Get** /api/v1/org_financing_and_ownership | org_financing_and_ownership



## ApiV1InvestorProfileGet

> ApiV1InvestorProfileGet200Response ApiV1InvestorProfileGet(ctx).InvestorId(investorId).InvestorName(investorName).Execute()

investor_profile

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
	investorId := "adv:283882" // string |  (optional)
	investorName := "Bioventure Capital Advisers" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1OwnershipAPI.ApiV1InvestorProfileGet(context.Background()).InvestorId(investorId).InvestorName(investorName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1OwnershipAPI.ApiV1InvestorProfileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1InvestorProfileGet`: ApiV1InvestorProfileGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1OwnershipAPI.ApiV1InvestorProfileGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InvestorProfileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investorId** | **string** |  | 
 **investorName** | **string** |  | 

### Return type

[**ApiV1InvestorProfileGet200Response**](ApiV1InvestorProfileGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1OrgFinancingAndOwnershipGet

> ApiV1OrgFinancingAndOwnershipGet200Response ApiV1OrgFinancingAndOwnershipGet(ctx).EdgarCik(edgarCik).OrgId(orgId).OrgName(orgName).RorId(rorId).Execute()

org_financing_and_ownership

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
	edgarCik := int32(1114448) // int32 |  (optional)
	orgId := "edgar:0001048477" // string |  (optional)
	orgName := "Novartis Ag" // string |  (optional)
	rorId := "01hcyya48" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1OwnershipAPI.ApiV1OrgFinancingAndOwnershipGet(context.Background()).EdgarCik(edgarCik).OrgId(orgId).OrgName(orgName).RorId(rorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1OwnershipAPI.ApiV1OrgFinancingAndOwnershipGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1OrgFinancingAndOwnershipGet`: ApiV1OrgFinancingAndOwnershipGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1OwnershipAPI.ApiV1OrgFinancingAndOwnershipGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1OrgFinancingAndOwnershipGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgarCik** | **int32** |  | 
 **orgId** | **string** |  | 
 **orgName** | **string** |  | 
 **rorId** | **string** |  | 

### Return type

[**ApiV1OrgFinancingAndOwnershipGet200Response**](ApiV1OrgFinancingAndOwnershipGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

