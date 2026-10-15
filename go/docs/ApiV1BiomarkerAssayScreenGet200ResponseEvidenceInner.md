# ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BiomarkerAssayEvidenceId** | **string** |  | 
**Pmid** | **string** |  | 
**GeneSymbolRaw** | **NullableString** |  | 
**BiomarkerId** | **NullableString** |  | 
**GeneResolutionStatus** | **string** |  | 
**GeneResolutionMethod** | **NullableString** |  | 
**AssayName** | **NullableString** |  | 
**SpecimenType** | **NullableString** |  | 
**ComparatorMethod** | **NullableString** |  | 
**ReportedPerformance** | **NullableString** |  | 
**ValidationTypeMentioned** | **string** |  | 
**AssayProviderCompany** | **[]string** |  | 
**FundingOrganizations** | **[]string** |  | 
**EvidenceQuote** | **[]string** |  | 
**Confidence** | **string** |  | 
**SourceRef** | [**ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInnerSourceRef**](ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInnerSourceRef.md) |  | 

## Methods

### NewApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner

`func NewApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner(biomarkerAssayEvidenceId string, pmid string, geneSymbolRaw NullableString, biomarkerId NullableString, geneResolutionStatus string, geneResolutionMethod NullableString, assayName NullableString, specimenType NullableString, comparatorMethod NullableString, reportedPerformance NullableString, validationTypeMentioned string, assayProviderCompany []string, fundingOrganizations []string, evidenceQuote []string, confidence string, sourceRef ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInnerSourceRef, ) *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner`

NewApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner instantiates a new ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1BiomarkerAssayScreenGet200ResponseEvidenceInnerWithDefaults

`func NewApiV1BiomarkerAssayScreenGet200ResponseEvidenceInnerWithDefaults() *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner`

NewApiV1BiomarkerAssayScreenGet200ResponseEvidenceInnerWithDefaults instantiates a new ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiomarkerAssayEvidenceId

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetBiomarkerAssayEvidenceId() string`

GetBiomarkerAssayEvidenceId returns the BiomarkerAssayEvidenceId field if non-nil, zero value otherwise.

### GetBiomarkerAssayEvidenceIdOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetBiomarkerAssayEvidenceIdOk() (*string, bool)`

GetBiomarkerAssayEvidenceIdOk returns a tuple with the BiomarkerAssayEvidenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerAssayEvidenceId

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetBiomarkerAssayEvidenceId(v string)`

SetBiomarkerAssayEvidenceId sets BiomarkerAssayEvidenceId field to given value.


### GetPmid

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetPmid() string`

GetPmid returns the Pmid field if non-nil, zero value otherwise.

### GetPmidOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetPmidOk() (*string, bool)`

GetPmidOk returns a tuple with the Pmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmid

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetPmid(v string)`

SetPmid sets Pmid field to given value.


### GetGeneSymbolRaw

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetGeneSymbolRaw() string`

GetGeneSymbolRaw returns the GeneSymbolRaw field if non-nil, zero value otherwise.

### GetGeneSymbolRawOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetGeneSymbolRawOk() (*string, bool)`

GetGeneSymbolRawOk returns a tuple with the GeneSymbolRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbolRaw

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetGeneSymbolRaw(v string)`

SetGeneSymbolRaw sets GeneSymbolRaw field to given value.


### SetGeneSymbolRawNil

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetGeneSymbolRawNil(b bool)`

 SetGeneSymbolRawNil sets the value for GeneSymbolRaw to be an explicit nil

### UnsetGeneSymbolRaw
`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) UnsetGeneSymbolRaw()`

UnsetGeneSymbolRaw ensures that no value is present for GeneSymbolRaw, not even an explicit nil
### GetBiomarkerId

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetBiomarkerId() string`

GetBiomarkerId returns the BiomarkerId field if non-nil, zero value otherwise.

### GetBiomarkerIdOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetBiomarkerIdOk() (*string, bool)`

GetBiomarkerIdOk returns a tuple with the BiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerId

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetBiomarkerId(v string)`

SetBiomarkerId sets BiomarkerId field to given value.


### SetBiomarkerIdNil

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetBiomarkerIdNil(b bool)`

 SetBiomarkerIdNil sets the value for BiomarkerId to be an explicit nil

### UnsetBiomarkerId
`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) UnsetBiomarkerId()`

UnsetBiomarkerId ensures that no value is present for BiomarkerId, not even an explicit nil
### GetGeneResolutionStatus

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetGeneResolutionStatus() string`

GetGeneResolutionStatus returns the GeneResolutionStatus field if non-nil, zero value otherwise.

### GetGeneResolutionStatusOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetGeneResolutionStatusOk() (*string, bool)`

GetGeneResolutionStatusOk returns a tuple with the GeneResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneResolutionStatus

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetGeneResolutionStatus(v string)`

SetGeneResolutionStatus sets GeneResolutionStatus field to given value.


### GetGeneResolutionMethod

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetGeneResolutionMethod() string`

GetGeneResolutionMethod returns the GeneResolutionMethod field if non-nil, zero value otherwise.

### GetGeneResolutionMethodOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetGeneResolutionMethodOk() (*string, bool)`

GetGeneResolutionMethodOk returns a tuple with the GeneResolutionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneResolutionMethod

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetGeneResolutionMethod(v string)`

SetGeneResolutionMethod sets GeneResolutionMethod field to given value.


### SetGeneResolutionMethodNil

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetGeneResolutionMethodNil(b bool)`

 SetGeneResolutionMethodNil sets the value for GeneResolutionMethod to be an explicit nil

### UnsetGeneResolutionMethod
`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) UnsetGeneResolutionMethod()`

UnsetGeneResolutionMethod ensures that no value is present for GeneResolutionMethod, not even an explicit nil
### GetAssayName

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetAssayName() string`

GetAssayName returns the AssayName field if non-nil, zero value otherwise.

### GetAssayNameOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetAssayNameOk() (*string, bool)`

GetAssayNameOk returns a tuple with the AssayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayName

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetAssayName(v string)`

SetAssayName sets AssayName field to given value.


### SetAssayNameNil

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetAssayNameNil(b bool)`

 SetAssayNameNil sets the value for AssayName to be an explicit nil

### UnsetAssayName
`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) UnsetAssayName()`

UnsetAssayName ensures that no value is present for AssayName, not even an explicit nil
### GetSpecimenType

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetSpecimenType() string`

GetSpecimenType returns the SpecimenType field if non-nil, zero value otherwise.

### GetSpecimenTypeOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetSpecimenTypeOk() (*string, bool)`

GetSpecimenTypeOk returns a tuple with the SpecimenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecimenType

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetSpecimenType(v string)`

SetSpecimenType sets SpecimenType field to given value.


### SetSpecimenTypeNil

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetSpecimenTypeNil(b bool)`

 SetSpecimenTypeNil sets the value for SpecimenType to be an explicit nil

### UnsetSpecimenType
`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) UnsetSpecimenType()`

UnsetSpecimenType ensures that no value is present for SpecimenType, not even an explicit nil
### GetComparatorMethod

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetComparatorMethod() string`

GetComparatorMethod returns the ComparatorMethod field if non-nil, zero value otherwise.

### GetComparatorMethodOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetComparatorMethodOk() (*string, bool)`

GetComparatorMethodOk returns a tuple with the ComparatorMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparatorMethod

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetComparatorMethod(v string)`

SetComparatorMethod sets ComparatorMethod field to given value.


### SetComparatorMethodNil

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetComparatorMethodNil(b bool)`

 SetComparatorMethodNil sets the value for ComparatorMethod to be an explicit nil

### UnsetComparatorMethod
`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) UnsetComparatorMethod()`

UnsetComparatorMethod ensures that no value is present for ComparatorMethod, not even an explicit nil
### GetReportedPerformance

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetReportedPerformance() string`

GetReportedPerformance returns the ReportedPerformance field if non-nil, zero value otherwise.

### GetReportedPerformanceOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetReportedPerformanceOk() (*string, bool)`

GetReportedPerformanceOk returns a tuple with the ReportedPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedPerformance

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetReportedPerformance(v string)`

SetReportedPerformance sets ReportedPerformance field to given value.


### SetReportedPerformanceNil

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetReportedPerformanceNil(b bool)`

 SetReportedPerformanceNil sets the value for ReportedPerformance to be an explicit nil

### UnsetReportedPerformance
`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) UnsetReportedPerformance()`

UnsetReportedPerformance ensures that no value is present for ReportedPerformance, not even an explicit nil
### GetValidationTypeMentioned

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetValidationTypeMentioned() string`

GetValidationTypeMentioned returns the ValidationTypeMentioned field if non-nil, zero value otherwise.

### GetValidationTypeMentionedOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetValidationTypeMentionedOk() (*string, bool)`

GetValidationTypeMentionedOk returns a tuple with the ValidationTypeMentioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationTypeMentioned

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetValidationTypeMentioned(v string)`

SetValidationTypeMentioned sets ValidationTypeMentioned field to given value.


### GetAssayProviderCompany

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetAssayProviderCompany() []string`

GetAssayProviderCompany returns the AssayProviderCompany field if non-nil, zero value otherwise.

### GetAssayProviderCompanyOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetAssayProviderCompanyOk() (*[]string, bool)`

GetAssayProviderCompanyOk returns a tuple with the AssayProviderCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayProviderCompany

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetAssayProviderCompany(v []string)`

SetAssayProviderCompany sets AssayProviderCompany field to given value.


### SetAssayProviderCompanyNil

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetAssayProviderCompanyNil(b bool)`

 SetAssayProviderCompanyNil sets the value for AssayProviderCompany to be an explicit nil

### UnsetAssayProviderCompany
`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) UnsetAssayProviderCompany()`

UnsetAssayProviderCompany ensures that no value is present for AssayProviderCompany, not even an explicit nil
### GetFundingOrganizations

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetFundingOrganizations() []string`

GetFundingOrganizations returns the FundingOrganizations field if non-nil, zero value otherwise.

### GetFundingOrganizationsOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetFundingOrganizationsOk() (*[]string, bool)`

GetFundingOrganizationsOk returns a tuple with the FundingOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingOrganizations

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetFundingOrganizations(v []string)`

SetFundingOrganizations sets FundingOrganizations field to given value.


### SetFundingOrganizationsNil

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetFundingOrganizationsNil(b bool)`

 SetFundingOrganizationsNil sets the value for FundingOrganizations to be an explicit nil

### UnsetFundingOrganizations
`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) UnsetFundingOrganizations()`

UnsetFundingOrganizations ensures that no value is present for FundingOrganizations, not even an explicit nil
### GetEvidenceQuote

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetEvidenceQuote() []string`

GetEvidenceQuote returns the EvidenceQuote field if non-nil, zero value otherwise.

### GetEvidenceQuoteOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetEvidenceQuoteOk() (*[]string, bool)`

GetEvidenceQuoteOk returns a tuple with the EvidenceQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceQuote

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetEvidenceQuote(v []string)`

SetEvidenceQuote sets EvidenceQuote field to given value.


### GetConfidence

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.


### GetSourceRef

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetSourceRef() ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInnerSourceRef`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) GetSourceRefOk() (*ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInnerSourceRef, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner) SetSourceRef(v ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInnerSourceRef)`

SetSourceRef sets SourceRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


