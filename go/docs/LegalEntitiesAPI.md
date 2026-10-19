# \LegalEntitiesAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LegalEntitiesCompanyPrograms**](LegalEntitiesAPI.md#LegalEntitiesCompanyPrograms) | **Get** /api/v1/legal_entities/{id}/company_programs | Company programs
[**LegalEntitiesDrugPatentCliff**](LegalEntitiesAPI.md#LegalEntitiesDrugPatentCliff) | **Get** /api/v1/legal_entities/{id}/drug_patent_cliff | Patent-cliff exposure
[**LegalEntitiesGet**](LegalEntitiesAPI.md#LegalEntitiesGet) | **Get** /api/v1/legal_entities/{id} | One legal entity
[**LegalEntitiesInstitutionalHoldings**](LegalEntitiesAPI.md#LegalEntitiesInstitutionalHoldings) | **Get** /api/v1/legal_entities/{id}/institutional_holdings | Institutional holdings (SEC Form 13F)
[**LegalEntitiesList**](LegalEntitiesAPI.md#LegalEntitiesList) | **Get** /api/v1/legal_entities | List legal entities
[**LegalEntitiesPdufaDates**](LegalEntitiesAPI.md#LegalEntitiesPdufaDates) | **Get** /api/v1/legal_entities/{id}/pdufa_dates | FDA decision dates (PDUFA) and catalysts
[**LegalEntitiesSponsorRecruitmentByPhase**](LegalEntitiesAPI.md#LegalEntitiesSponsorRecruitmentByPhase) | **Get** /api/v1/legal_entities/{id}/sponsor_recruitment_by_phase | Trial recruitment duration by sponsor and phase
[**LegalEntitiesSponsorRecruitmentDuration**](LegalEntitiesAPI.md#LegalEntitiesSponsorRecruitmentDuration) | **Get** /api/v1/legal_entities/{id}/sponsor_recruitment_duration | Trial recruitment duration by sponsor
[**LegalEntitiesSponsorTrialStartForecast**](LegalEntitiesAPI.md#LegalEntitiesSponsorTrialStartForecast) | **Get** /api/v1/legal_entities/{id}/sponsor_trial_start_forecast | Trial-start forecast by sponsor
[**LegalEntitiesSponsorTrials**](LegalEntitiesAPI.md#LegalEntitiesSponsorTrials) | **Get** /api/v1/legal_entities/{id}/sponsor_trials | Sponsor trials



## LegalEntitiesCompanyPrograms

> CompanyProgramsResponse LegalEntitiesCompanyPrograms(ctx, id).Limit(limit).Sort(sort).Execute()

Company programs

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
	id := "edgar:0001114448" // string | 
	limit := int32(2) // int32 |  (optional)
	sort := "max_phase" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalEntitiesAPI.LegalEntitiesCompanyPrograms(context.Background(), id).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalEntitiesAPI.LegalEntitiesCompanyPrograms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LegalEntitiesCompanyPrograms`: CompanyProgramsResponse
	fmt.Fprintf(os.Stdout, "Response from `LegalEntitiesAPI.LegalEntitiesCompanyPrograms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLegalEntitiesCompanyProgramsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**CompanyProgramsResponse**](CompanyProgramsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LegalEntitiesDrugPatentCliff

> DrugPatentCliffResponse LegalEntitiesDrugPatentCliff(ctx, id).Limit(limit).Sort(sort).Execute()

Patent-cliff exposure

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
	id := "edgar:0001114448" // string | 
	limit := int32(2) // int32 |  (optional)
	sort := "cliff_date" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalEntitiesAPI.LegalEntitiesDrugPatentCliff(context.Background(), id).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalEntitiesAPI.LegalEntitiesDrugPatentCliff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LegalEntitiesDrugPatentCliff`: DrugPatentCliffResponse
	fmt.Fprintf(os.Stdout, "Response from `LegalEntitiesAPI.LegalEntitiesDrugPatentCliff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLegalEntitiesDrugPatentCliffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**DrugPatentCliffResponse**](DrugPatentCliffResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LegalEntitiesGet

> LegalEntitiesRecord LegalEntitiesGet(ctx, id).Execute()

One legal entity

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
	id := "edgar:0001114448" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalEntitiesAPI.LegalEntitiesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalEntitiesAPI.LegalEntitiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LegalEntitiesGet`: LegalEntitiesRecord
	fmt.Fprintf(os.Stdout, "Response from `LegalEntitiesAPI.LegalEntitiesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLegalEntitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LegalEntitiesRecord**](LegalEntitiesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LegalEntitiesInstitutionalHoldings

> InstitutionalHoldingsResponse LegalEntitiesInstitutionalHoldings(ctx, id).Limit(limit).Sort(sort).Execute()

Institutional holdings (SEC Form 13F)

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
	id := "edgar:0001114448" // string | 
	limit := int32(2) // int32 |  (optional)
	sort := "reported_value_amount" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalEntitiesAPI.LegalEntitiesInstitutionalHoldings(context.Background(), id).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalEntitiesAPI.LegalEntitiesInstitutionalHoldings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LegalEntitiesInstitutionalHoldings`: InstitutionalHoldingsResponse
	fmt.Fprintf(os.Stdout, "Response from `LegalEntitiesAPI.LegalEntitiesInstitutionalHoldings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLegalEntitiesInstitutionalHoldingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**InstitutionalHoldingsResponse**](InstitutionalHoldingsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LegalEntitiesList

> LegalEntitiesResponse LegalEntitiesList(ctx).Limit(limit).Sort(sort).Execute()

List legal entities

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
	limit := int32(2) // int32 |  (optional)
	sort := "activity_score" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalEntitiesAPI.LegalEntitiesList(context.Background()).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalEntitiesAPI.LegalEntitiesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LegalEntitiesList`: LegalEntitiesResponse
	fmt.Fprintf(os.Stdout, "Response from `LegalEntitiesAPI.LegalEntitiesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLegalEntitiesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**LegalEntitiesResponse**](LegalEntitiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LegalEntitiesPdufaDates

> PdufaDatesResponse LegalEntitiesPdufaDates(ctx, id).Limit(limit).Sort(sort).Execute()

FDA decision dates (PDUFA) and catalysts

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
	id := "edgar:0001114448" // string | 
	limit := int32(2) // int32 |  (optional)
	sort := "decision_date" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalEntitiesAPI.LegalEntitiesPdufaDates(context.Background(), id).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalEntitiesAPI.LegalEntitiesPdufaDates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LegalEntitiesPdufaDates`: PdufaDatesResponse
	fmt.Fprintf(os.Stdout, "Response from `LegalEntitiesAPI.LegalEntitiesPdufaDates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLegalEntitiesPdufaDatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**PdufaDatesResponse**](PdufaDatesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LegalEntitiesSponsorRecruitmentByPhase

> SponsorRecruitmentByPhaseResponse LegalEntitiesSponsorRecruitmentByPhase(ctx, id).Limit(limit).Sort(sort).Execute()

Trial recruitment duration by sponsor and phase

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
	id := "edgar:0001114448" // string | 
	limit := int32(2) // int32 |  (optional)
	sort := "sponsor_org_id" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalEntitiesAPI.LegalEntitiesSponsorRecruitmentByPhase(context.Background(), id).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalEntitiesAPI.LegalEntitiesSponsorRecruitmentByPhase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LegalEntitiesSponsorRecruitmentByPhase`: SponsorRecruitmentByPhaseResponse
	fmt.Fprintf(os.Stdout, "Response from `LegalEntitiesAPI.LegalEntitiesSponsorRecruitmentByPhase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLegalEntitiesSponsorRecruitmentByPhaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**SponsorRecruitmentByPhaseResponse**](SponsorRecruitmentByPhaseResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LegalEntitiesSponsorRecruitmentDuration

> SponsorRecruitmentDurationResponse LegalEntitiesSponsorRecruitmentDuration(ctx, id).Limit(limit).Sort(sort).Execute()

Trial recruitment duration by sponsor

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
	id := "edgar:0001114448" // string | 
	limit := int32(2) // int32 |  (optional)
	sort := "sponsor_org_id" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalEntitiesAPI.LegalEntitiesSponsorRecruitmentDuration(context.Background(), id).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalEntitiesAPI.LegalEntitiesSponsorRecruitmentDuration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LegalEntitiesSponsorRecruitmentDuration`: SponsorRecruitmentDurationResponse
	fmt.Fprintf(os.Stdout, "Response from `LegalEntitiesAPI.LegalEntitiesSponsorRecruitmentDuration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLegalEntitiesSponsorRecruitmentDurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**SponsorRecruitmentDurationResponse**](SponsorRecruitmentDurationResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LegalEntitiesSponsorTrialStartForecast

> SponsorTrialStartForecastResponse LegalEntitiesSponsorTrialStartForecast(ctx, id).Limit(limit).Sort(sort).Execute()

Trial-start forecast by sponsor

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
	id := "edgar:0001114448" // string | 
	limit := int32(2) // int32 |  (optional)
	sort := "sponsor_org_id" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalEntitiesAPI.LegalEntitiesSponsorTrialStartForecast(context.Background(), id).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalEntitiesAPI.LegalEntitiesSponsorTrialStartForecast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LegalEntitiesSponsorTrialStartForecast`: SponsorTrialStartForecastResponse
	fmt.Fprintf(os.Stdout, "Response from `LegalEntitiesAPI.LegalEntitiesSponsorTrialStartForecast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLegalEntitiesSponsorTrialStartForecastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**SponsorTrialStartForecastResponse**](SponsorTrialStartForecastResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LegalEntitiesSponsorTrials

> SponsorTrialsResponse LegalEntitiesSponsorTrials(ctx, id).Limit(limit).Sort(sort).Execute()

Sponsor trials

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
	id := "edgar:0001114448" // string | 
	limit := int32(2) // int32 |  (optional)
	sort := "start_date" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalEntitiesAPI.LegalEntitiesSponsorTrials(context.Background(), id).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalEntitiesAPI.LegalEntitiesSponsorTrials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LegalEntitiesSponsorTrials`: SponsorTrialsResponse
	fmt.Fprintf(os.Stdout, "Response from `LegalEntitiesAPI.LegalEntitiesSponsorTrials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLegalEntitiesSponsorTrialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**SponsorTrialsResponse**](SponsorTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

