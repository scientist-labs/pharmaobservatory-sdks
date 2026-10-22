# \RwdStudiesAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RwdStudiesGet**](RwdStudiesAPI.md#RwdStudiesGet) | **Get** /api/v1/rwd_studies/{id} | One real-world data study
[**RwdStudiesList**](RwdStudiesAPI.md#RwdStudiesList) | **Get** /api/v1/rwd_studies | List real-world data studies
[**RwdStudiesRwdConductorStudies**](RwdStudiesAPI.md#RwdStudiesRwdConductorStudies) | **Get** /api/v1/rwd_studies/{id}/rwd_conductor_studies | Real-world data study conductor studies (HMA-EMA catalogue)
[**RwdStudiesRwdStudyConditions**](RwdStudiesAPI.md#RwdStudiesRwdStudyConditions) | **Get** /api/v1/rwd_studies/{id}/rwd_study_conditions | Registered real-world data study conditions
[**RwdStudiesRwdStudyDataSources**](RwdStudiesAPI.md#RwdStudiesRwdStudyDataSources) | **Get** /api/v1/rwd_studies/{id}/rwd_study_data_sources | Registered real-world data study data sources
[**RwdStudiesRwdStudyDrugs**](RwdStudiesAPI.md#RwdStudiesRwdStudyDrugs) | **Get** /api/v1/rwd_studies/{id}/rwd_study_drugs | Registered real-world data study drugs
[**RwdStudiesRwdStudyMilestones**](RwdStudiesAPI.md#RwdStudiesRwdStudyMilestones) | **Get** /api/v1/rwd_studies/{id}/rwd_study_milestones | Registered real-world data study milestones



## RwdStudiesGet

> RwdStudiesRecord RwdStudiesGet(ctx, id).Execute()

One real-world data study

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
	id := int32(900000301) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdStudiesAPI.RwdStudiesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdStudiesAPI.RwdStudiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdStudiesGet`: RwdStudiesRecord
	fmt.Fprintf(os.Stdout, "Response from `RwdStudiesAPI.RwdStudiesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRwdStudiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RwdStudiesRecord**](RwdStudiesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RwdStudiesList

> RwdStudiesResponse RwdStudiesList(ctx).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

List real-world data studies

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
	resp, r, err := apiClient.RwdStudiesAPI.RwdStudiesList(context.Background()).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdStudiesAPI.RwdStudiesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdStudiesList`: RwdStudiesResponse
	fmt.Fprintf(os.Stdout, "Response from `RwdStudiesAPI.RwdStudiesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRwdStudiesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**RwdStudiesResponse**](RwdStudiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RwdStudiesRwdConductorStudies

> RwdConductorStudiesResponse RwdStudiesRwdConductorStudies(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

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
	id := int32(900000301) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdStudiesAPI.RwdStudiesRwdConductorStudies(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdStudiesAPI.RwdStudiesRwdConductorStudies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdStudiesRwdConductorStudies`: RwdConductorStudiesResponse
	fmt.Fprintf(os.Stdout, "Response from `RwdStudiesAPI.RwdStudiesRwdConductorStudies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRwdStudiesRwdConductorStudiesRequest struct via the builder pattern


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


## RwdStudiesRwdStudyConditions

> RwdStudyConditionsResponse RwdStudiesRwdStudyConditions(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

Registered real-world data study conditions

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
	id := int32(900000301) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdStudiesAPI.RwdStudiesRwdStudyConditions(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdStudiesAPI.RwdStudiesRwdStudyConditions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdStudiesRwdStudyConditions`: RwdStudyConditionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RwdStudiesAPI.RwdStudiesRwdStudyConditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRwdStudiesRwdStudyConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**RwdStudyConditionsResponse**](RwdStudyConditionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RwdStudiesRwdStudyDataSources

> RwdStudyDataSourcesResponse RwdStudiesRwdStudyDataSources(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

Registered real-world data study data sources

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
	id := int32(900000301) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdStudiesAPI.RwdStudiesRwdStudyDataSources(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdStudiesAPI.RwdStudiesRwdStudyDataSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdStudiesRwdStudyDataSources`: RwdStudyDataSourcesResponse
	fmt.Fprintf(os.Stdout, "Response from `RwdStudiesAPI.RwdStudiesRwdStudyDataSources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRwdStudiesRwdStudyDataSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**RwdStudyDataSourcesResponse**](RwdStudyDataSourcesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RwdStudiesRwdStudyDrugs

> RwdStudyDrugsResponse RwdStudiesRwdStudyDrugs(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

Registered real-world data study drugs

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
	id := int32(900000301) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdStudiesAPI.RwdStudiesRwdStudyDrugs(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdStudiesAPI.RwdStudiesRwdStudyDrugs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdStudiesRwdStudyDrugs`: RwdStudyDrugsResponse
	fmt.Fprintf(os.Stdout, "Response from `RwdStudiesAPI.RwdStudiesRwdStudyDrugs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRwdStudiesRwdStudyDrugsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**RwdStudyDrugsResponse**](RwdStudyDrugsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RwdStudiesRwdStudyMilestones

> RwdStudyMilestonesResponse RwdStudiesRwdStudyMilestones(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()

Registered real-world data study milestones

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
	id := int32(900000301) // int32 | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. (optional)
	limit := int32(56) // int32 | Rows per page, up to 500 (default 50). (optional)
	sort := "sort_example" // string | Comma separated columns, a leading minus for descending: sort=a,-b. (optional)
	groupBy := "groupBy_example" // string | A dimension to group by; a date takes :year, :quarter or :month. (optional)
	measure := "measure_example" // string | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RwdStudiesAPI.RwdStudiesRwdStudyMilestones(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).GroupBy(groupBy).Measure(measure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RwdStudiesAPI.RwdStudiesRwdStudyMilestones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RwdStudiesRwdStudyMilestones`: RwdStudyMilestonesResponse
	fmt.Fprintf(os.Stdout, "Response from `RwdStudiesAPI.RwdStudiesRwdStudyMilestones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRwdStudiesRwdStudyMilestonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. Opaque; it is bound to the filters and the warehouse build that issued it. | 
 **limit** | **int32** | Rows per page, up to 500 (default 50). | 
 **sort** | **string** | Comma separated columns, a leading minus for descending: sort&#x3D;a,-b. | 
 **groupBy** | **string** | A dimension to group by; a date takes :year, :quarter or :month. | 
 **measure** | **string** | count, sum(col), avg(col), min(col), max(col), median(col) or count_distinct(col). | 

### Return type

[**RwdStudyMilestonesResponse**](RwdStudyMilestonesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

