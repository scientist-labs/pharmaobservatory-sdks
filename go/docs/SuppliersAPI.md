# \SuppliersAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SuppliersGet**](SuppliersAPI.md#SuppliersGet) | **Get** /api/v1/suppliers/{id} | One supplier
[**SuppliersList**](SuppliersAPI.md#SuppliersList) | **Get** /api/v1/suppliers | List suppliers
[**SuppliersProviderBiomarkerCapability**](SuppliersAPI.md#SuppliersProviderBiomarkerCapability) | **Get** /api/v1/suppliers/{id}/provider_biomarker_capability | Supplier biomarker capability
[**SuppliersProviderBiomarkerCapabilitySummary**](SuppliersAPI.md#SuppliersProviderBiomarkerCapabilitySummary) | **Get** /api/v1/suppliers/{id}/provider_biomarker_capability_summary | Supplier biomarker capability, summarised
[**SuppliersProviderCliaLicensure**](SuppliersAPI.md#SuppliersProviderCliaLicensure) | **Get** /api/v1/suppliers/{id}/provider_clia_licensure | Supplier CLIA laboratory licensure
[**SuppliersProviderRwdDelivery**](SuppliersAPI.md#SuppliersProviderRwdDelivery) | **Get** /api/v1/suppliers/{id}/provider_rwd_delivery | RX provider real-world data study delivery track record (HMA-EMA catalogue)
[**SuppliersProviderRwdDeliveryStudy**](SuppliersAPI.md#SuppliersProviderRwdDeliveryStudy) | **Get** /api/v1/suppliers/{id}/provider_rwd_delivery_study | RX provider real-world data study final-report dates below the floor (HMA-EMA catalogue)
[**SuppliersProviderRwdFootprint**](SuppliersAPI.md#SuppliersProviderRwdFootprint) | **Get** /api/v1/suppliers/{id}/provider_rwd_footprint | Supplier real-world data footprint (HMA-EMA catalogue)
[**SuppliersRwdUsNetworkMemberships**](SuppliersAPI.md#SuppliersRwdUsNetworkMemberships) | **Get** /api/v1/suppliers/{id}/rwd_us_network_memberships | US real-world data network membership (N3C)
[**SuppliersSupplierRwdDeclarations**](SuppliersAPI.md#SuppliersSupplierRwdDeclarations) | **Get** /api/v1/suppliers/{id}/supplier_rwd_declarations | Supplier declared real-world data facts
[**SuppliersSupplierServices**](SuppliersAPI.md#SuppliersSupplierServices) | **Get** /api/v1/suppliers/{id}/supplier_services | Supplier services
[**SuppliersSupplierSites**](SuppliersAPI.md#SuppliersSupplierSites) | **Get** /api/v1/suppliers/{id}/supplier_sites | Supplier sites



## SuppliersGet

> SuppliersRecord SuppliersGet(ctx, id).Execute()

One supplier

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
	id := int32(990502) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppliersAPI.SuppliersGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppliersAPI.SuppliersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppliersGet`: SuppliersRecord
	fmt.Fprintf(os.Stdout, "Response from `SuppliersAPI.SuppliersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppliersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuppliersRecord**](SuppliersRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppliersList

> SuppliersResponse SuppliersList(ctx).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

List suppliers

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
	resp, r, err := apiClient.SuppliersAPI.SuppliersList(context.Background()).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppliersAPI.SuppliersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppliersList`: SuppliersResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppliersAPI.SuppliersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuppliersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**SuppliersResponse**](SuppliersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppliersProviderBiomarkerCapability

> ProviderBiomarkerCapabilityResponse SuppliersProviderBiomarkerCapability(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

Supplier biomarker capability

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
	id := int32(990502) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppliersAPI.SuppliersProviderBiomarkerCapability(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppliersAPI.SuppliersProviderBiomarkerCapability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppliersProviderBiomarkerCapability`: ProviderBiomarkerCapabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppliersAPI.SuppliersProviderBiomarkerCapability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppliersProviderBiomarkerCapabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**ProviderBiomarkerCapabilityResponse**](ProviderBiomarkerCapabilityResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppliersProviderBiomarkerCapabilitySummary

> ProviderBiomarkerCapabilitySummaryResponse SuppliersProviderBiomarkerCapabilitySummary(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

Supplier biomarker capability, summarised

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
	id := int32(990502) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppliersAPI.SuppliersProviderBiomarkerCapabilitySummary(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppliersAPI.SuppliersProviderBiomarkerCapabilitySummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppliersProviderBiomarkerCapabilitySummary`: ProviderBiomarkerCapabilitySummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppliersAPI.SuppliersProviderBiomarkerCapabilitySummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppliersProviderBiomarkerCapabilitySummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**ProviderBiomarkerCapabilitySummaryResponse**](ProviderBiomarkerCapabilitySummaryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppliersProviderCliaLicensure

> ProviderCliaLicensureResponse SuppliersProviderCliaLicensure(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

Supplier CLIA laboratory licensure

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
	id := int32(990502) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppliersAPI.SuppliersProviderCliaLicensure(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppliersAPI.SuppliersProviderCliaLicensure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppliersProviderCliaLicensure`: ProviderCliaLicensureResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppliersAPI.SuppliersProviderCliaLicensure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppliersProviderCliaLicensureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**ProviderCliaLicensureResponse**](ProviderCliaLicensureResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppliersProviderRwdDelivery

> ProviderRwdDeliveryResponse SuppliersProviderRwdDelivery(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

RX provider real-world data study delivery track record (HMA-EMA catalogue)

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
	id := int32(990502) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppliersAPI.SuppliersProviderRwdDelivery(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppliersAPI.SuppliersProviderRwdDelivery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppliersProviderRwdDelivery`: ProviderRwdDeliveryResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppliersAPI.SuppliersProviderRwdDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppliersProviderRwdDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**ProviderRwdDeliveryResponse**](ProviderRwdDeliveryResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppliersProviderRwdDeliveryStudy

> ProviderRwdDeliveryStudyResponse SuppliersProviderRwdDeliveryStudy(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

RX provider real-world data study final-report dates below the floor (HMA-EMA catalogue)

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
	id := int32(990502) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppliersAPI.SuppliersProviderRwdDeliveryStudy(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppliersAPI.SuppliersProviderRwdDeliveryStudy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppliersProviderRwdDeliveryStudy`: ProviderRwdDeliveryStudyResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppliersAPI.SuppliersProviderRwdDeliveryStudy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppliersProviderRwdDeliveryStudyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**ProviderRwdDeliveryStudyResponse**](ProviderRwdDeliveryStudyResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppliersProviderRwdFootprint

> ProviderRwdFootprintResponse SuppliersProviderRwdFootprint(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

Supplier real-world data footprint (HMA-EMA catalogue)

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
	id := int32(990502) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppliersAPI.SuppliersProviderRwdFootprint(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppliersAPI.SuppliersProviderRwdFootprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppliersProviderRwdFootprint`: ProviderRwdFootprintResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppliersAPI.SuppliersProviderRwdFootprint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppliersProviderRwdFootprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**ProviderRwdFootprintResponse**](ProviderRwdFootprintResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppliersRwdUsNetworkMemberships

> RwdUsNetworkMembershipsResponse SuppliersRwdUsNetworkMemberships(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

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
	id := int32(990502) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppliersAPI.SuppliersRwdUsNetworkMemberships(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppliersAPI.SuppliersRwdUsNetworkMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppliersRwdUsNetworkMemberships`: RwdUsNetworkMembershipsResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppliersAPI.SuppliersRwdUsNetworkMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppliersRwdUsNetworkMembershipsRequest struct via the builder pattern


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


## SuppliersSupplierRwdDeclarations

> SupplierRwdDeclarationsResponse SuppliersSupplierRwdDeclarations(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

Supplier declared real-world data facts

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
	id := int32(990502) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppliersAPI.SuppliersSupplierRwdDeclarations(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppliersAPI.SuppliersSupplierRwdDeclarations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppliersSupplierRwdDeclarations`: SupplierRwdDeclarationsResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppliersAPI.SuppliersSupplierRwdDeclarations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppliersSupplierRwdDeclarationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**SupplierRwdDeclarationsResponse**](SupplierRwdDeclarationsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppliersSupplierServices

> SupplierServicesResponse SuppliersSupplierServices(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

Supplier services

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
	id := int32(990502) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppliersAPI.SuppliersSupplierServices(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppliersAPI.SuppliersSupplierServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppliersSupplierServices`: SupplierServicesResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppliersAPI.SuppliersSupplierServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppliersSupplierServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**SupplierServicesResponse**](SupplierServicesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppliersSupplierSites

> SupplierSitesResponse SuppliersSupplierSites(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

Supplier sites

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
	id := int32(990502) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppliersAPI.SuppliersSupplierSites(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppliersAPI.SuppliersSupplierSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppliersSupplierSites`: SupplierSitesResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppliersAPI.SuppliersSupplierSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppliersSupplierSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**SupplierSitesResponse**](SupplierSitesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

