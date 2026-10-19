# DrugPatentCliffRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrugId** | Pointer to **NullableString** | The drug key. | [optional] 
**DrugName** | Pointer to **NullableString** | The drug&#39;s name. | [optional] 
**PharmClass** | Pointer to **NullableString** | The drug&#39;s pharmacologic class. | [optional] 
**AtcClass** | Pointer to **NullableString** | The drug&#39;s ATC class. | [optional] 
**Brand** | Pointer to **NullableString** | The brand name. | [optional] 
**CandidateBrands** | Pointer to **NullableString** | Other candidate brand names, comma separated. | [optional] 
**CliffDate** | Pointer to **NullableString** | The date exclusivity is expected to end (loss of exclusivity); always in the future when the dataset is built. | [optional] 
**CliffYear** | Pointer to **NullableInt32** | The year of cliff_date. | [optional] 
**CliffBasis** | Pointer to **NullableString** | Which regime set the date: curated_loe, orange_book_patent, pte_expiry or bpcia_exclusivity. | [optional] 
**CliffSource** | Pointer to **NullableString** | Source of a curated date (null on the other bases). | [optional] 
**CliffConfidence** | Pointer to **NullableString** | Confidence of a curated date (null on the other bases). | [optional] 
**HasGeneric** | Pointer to **NullableBool** | True when a generic is already marketed. | [optional] 
**BiosimilarAvailable** | Pointer to **NullableBool** | True when a biosimilar is already marketed. | [optional] 
**ApplicantName** | Pointer to **NullableString** | The brand or originator applicant as the Orange Book or Purple Book names it. | [optional] 
**ApplicantOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the applicant resolved to (null when unresolved). | [optional] 
**ApplicantCompanyId** | Pointer to **NullableString** | Company level: the company the applicant belongs to (null when unresolved). | [optional] 
**ApplicantCompanyName** | Pointer to **NullableString** | Display name of the applicant&#39;s company. | [optional] 
**ApplicantGroupId** | Pointer to **NullableString** | Group level: the applicant&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**ApplicantGroupName** | Pointer to **NullableString** | Display name of the applicant&#39;s group. | [optional] 
**ApplicantMatchScope** | Pointer to **NullableString** | How the applicant name matched an organization record. | [optional] 
**MedicareSpend** | Pointer to **NullableFloat32** | Latest-year Medicare Part B and D spend for the drug, in US dollars. | [optional] 
**MedicareSpendYear** | Pointer to **NullableInt32** | The year of medicare_spend. | [optional] 
**DisclosedRevenueUsdMillions** | Pointer to **NullableFloat32** | Latest full-year disclosed worldwide revenue, in US millions. | [optional] 
**DisclosedRevenueYear** | Pointer to **NullableInt32** | The fiscal year of the disclosed revenue. | [optional] 
**SizeUsd** | Pointer to **NullableFloat32** | The single size axis in US dollars: disclosed revenue when present, else Medicare spend. | [optional] 
**SizeBasis** | Pointer to **NullableString** | Which source size_usd used: disclosed_worldwide_revenue, medicare_spend or none. | [optional] 

## Methods

### NewDrugPatentCliffRow

`func NewDrugPatentCliffRow() *DrugPatentCliffRow`

NewDrugPatentCliffRow instantiates a new DrugPatentCliffRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugPatentCliffRowWithDefaults

`func NewDrugPatentCliffRowWithDefaults() *DrugPatentCliffRow`

NewDrugPatentCliffRowWithDefaults instantiates a new DrugPatentCliffRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrugId

`func (o *DrugPatentCliffRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *DrugPatentCliffRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *DrugPatentCliffRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *DrugPatentCliffRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *DrugPatentCliffRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *DrugPatentCliffRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetDrugName

`func (o *DrugPatentCliffRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *DrugPatentCliffRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *DrugPatentCliffRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *DrugPatentCliffRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *DrugPatentCliffRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *DrugPatentCliffRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetPharmClass

`func (o *DrugPatentCliffRow) GetPharmClass() string`

GetPharmClass returns the PharmClass field if non-nil, zero value otherwise.

### GetPharmClassOk

`func (o *DrugPatentCliffRow) GetPharmClassOk() (*string, bool)`

GetPharmClassOk returns a tuple with the PharmClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPharmClass

`func (o *DrugPatentCliffRow) SetPharmClass(v string)`

SetPharmClass sets PharmClass field to given value.

### HasPharmClass

`func (o *DrugPatentCliffRow) HasPharmClass() bool`

HasPharmClass returns a boolean if a field has been set.

### SetPharmClassNil

`func (o *DrugPatentCliffRow) SetPharmClassNil(b bool)`

 SetPharmClassNil sets the value for PharmClass to be an explicit nil

### UnsetPharmClass
`func (o *DrugPatentCliffRow) UnsetPharmClass()`

UnsetPharmClass ensures that no value is present for PharmClass, not even an explicit nil
### GetAtcClass

`func (o *DrugPatentCliffRow) GetAtcClass() string`

GetAtcClass returns the AtcClass field if non-nil, zero value otherwise.

### GetAtcClassOk

`func (o *DrugPatentCliffRow) GetAtcClassOk() (*string, bool)`

GetAtcClassOk returns a tuple with the AtcClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcClass

`func (o *DrugPatentCliffRow) SetAtcClass(v string)`

SetAtcClass sets AtcClass field to given value.

### HasAtcClass

`func (o *DrugPatentCliffRow) HasAtcClass() bool`

HasAtcClass returns a boolean if a field has been set.

### SetAtcClassNil

`func (o *DrugPatentCliffRow) SetAtcClassNil(b bool)`

 SetAtcClassNil sets the value for AtcClass to be an explicit nil

### UnsetAtcClass
`func (o *DrugPatentCliffRow) UnsetAtcClass()`

UnsetAtcClass ensures that no value is present for AtcClass, not even an explicit nil
### GetBrand

`func (o *DrugPatentCliffRow) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *DrugPatentCliffRow) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *DrugPatentCliffRow) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *DrugPatentCliffRow) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *DrugPatentCliffRow) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *DrugPatentCliffRow) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetCandidateBrands

`func (o *DrugPatentCliffRow) GetCandidateBrands() string`

GetCandidateBrands returns the CandidateBrands field if non-nil, zero value otherwise.

### GetCandidateBrandsOk

`func (o *DrugPatentCliffRow) GetCandidateBrandsOk() (*string, bool)`

GetCandidateBrandsOk returns a tuple with the CandidateBrands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateBrands

`func (o *DrugPatentCliffRow) SetCandidateBrands(v string)`

SetCandidateBrands sets CandidateBrands field to given value.

### HasCandidateBrands

`func (o *DrugPatentCliffRow) HasCandidateBrands() bool`

HasCandidateBrands returns a boolean if a field has been set.

### SetCandidateBrandsNil

`func (o *DrugPatentCliffRow) SetCandidateBrandsNil(b bool)`

 SetCandidateBrandsNil sets the value for CandidateBrands to be an explicit nil

### UnsetCandidateBrands
`func (o *DrugPatentCliffRow) UnsetCandidateBrands()`

UnsetCandidateBrands ensures that no value is present for CandidateBrands, not even an explicit nil
### GetCliffDate

`func (o *DrugPatentCliffRow) GetCliffDate() string`

GetCliffDate returns the CliffDate field if non-nil, zero value otherwise.

### GetCliffDateOk

`func (o *DrugPatentCliffRow) GetCliffDateOk() (*string, bool)`

GetCliffDateOk returns a tuple with the CliffDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliffDate

`func (o *DrugPatentCliffRow) SetCliffDate(v string)`

SetCliffDate sets CliffDate field to given value.

### HasCliffDate

`func (o *DrugPatentCliffRow) HasCliffDate() bool`

HasCliffDate returns a boolean if a field has been set.

### SetCliffDateNil

`func (o *DrugPatentCliffRow) SetCliffDateNil(b bool)`

 SetCliffDateNil sets the value for CliffDate to be an explicit nil

### UnsetCliffDate
`func (o *DrugPatentCliffRow) UnsetCliffDate()`

UnsetCliffDate ensures that no value is present for CliffDate, not even an explicit nil
### GetCliffYear

`func (o *DrugPatentCliffRow) GetCliffYear() int32`

GetCliffYear returns the CliffYear field if non-nil, zero value otherwise.

### GetCliffYearOk

`func (o *DrugPatentCliffRow) GetCliffYearOk() (*int32, bool)`

GetCliffYearOk returns a tuple with the CliffYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliffYear

`func (o *DrugPatentCliffRow) SetCliffYear(v int32)`

SetCliffYear sets CliffYear field to given value.

### HasCliffYear

`func (o *DrugPatentCliffRow) HasCliffYear() bool`

HasCliffYear returns a boolean if a field has been set.

### SetCliffYearNil

`func (o *DrugPatentCliffRow) SetCliffYearNil(b bool)`

 SetCliffYearNil sets the value for CliffYear to be an explicit nil

### UnsetCliffYear
`func (o *DrugPatentCliffRow) UnsetCliffYear()`

UnsetCliffYear ensures that no value is present for CliffYear, not even an explicit nil
### GetCliffBasis

`func (o *DrugPatentCliffRow) GetCliffBasis() string`

GetCliffBasis returns the CliffBasis field if non-nil, zero value otherwise.

### GetCliffBasisOk

`func (o *DrugPatentCliffRow) GetCliffBasisOk() (*string, bool)`

GetCliffBasisOk returns a tuple with the CliffBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliffBasis

`func (o *DrugPatentCliffRow) SetCliffBasis(v string)`

SetCliffBasis sets CliffBasis field to given value.

### HasCliffBasis

`func (o *DrugPatentCliffRow) HasCliffBasis() bool`

HasCliffBasis returns a boolean if a field has been set.

### SetCliffBasisNil

`func (o *DrugPatentCliffRow) SetCliffBasisNil(b bool)`

 SetCliffBasisNil sets the value for CliffBasis to be an explicit nil

### UnsetCliffBasis
`func (o *DrugPatentCliffRow) UnsetCliffBasis()`

UnsetCliffBasis ensures that no value is present for CliffBasis, not even an explicit nil
### GetCliffSource

`func (o *DrugPatentCliffRow) GetCliffSource() string`

GetCliffSource returns the CliffSource field if non-nil, zero value otherwise.

### GetCliffSourceOk

`func (o *DrugPatentCliffRow) GetCliffSourceOk() (*string, bool)`

GetCliffSourceOk returns a tuple with the CliffSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliffSource

`func (o *DrugPatentCliffRow) SetCliffSource(v string)`

SetCliffSource sets CliffSource field to given value.

### HasCliffSource

`func (o *DrugPatentCliffRow) HasCliffSource() bool`

HasCliffSource returns a boolean if a field has been set.

### SetCliffSourceNil

`func (o *DrugPatentCliffRow) SetCliffSourceNil(b bool)`

 SetCliffSourceNil sets the value for CliffSource to be an explicit nil

### UnsetCliffSource
`func (o *DrugPatentCliffRow) UnsetCliffSource()`

UnsetCliffSource ensures that no value is present for CliffSource, not even an explicit nil
### GetCliffConfidence

`func (o *DrugPatentCliffRow) GetCliffConfidence() string`

GetCliffConfidence returns the CliffConfidence field if non-nil, zero value otherwise.

### GetCliffConfidenceOk

`func (o *DrugPatentCliffRow) GetCliffConfidenceOk() (*string, bool)`

GetCliffConfidenceOk returns a tuple with the CliffConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliffConfidence

`func (o *DrugPatentCliffRow) SetCliffConfidence(v string)`

SetCliffConfidence sets CliffConfidence field to given value.

### HasCliffConfidence

`func (o *DrugPatentCliffRow) HasCliffConfidence() bool`

HasCliffConfidence returns a boolean if a field has been set.

### SetCliffConfidenceNil

`func (o *DrugPatentCliffRow) SetCliffConfidenceNil(b bool)`

 SetCliffConfidenceNil sets the value for CliffConfidence to be an explicit nil

### UnsetCliffConfidence
`func (o *DrugPatentCliffRow) UnsetCliffConfidence()`

UnsetCliffConfidence ensures that no value is present for CliffConfidence, not even an explicit nil
### GetHasGeneric

`func (o *DrugPatentCliffRow) GetHasGeneric() bool`

GetHasGeneric returns the HasGeneric field if non-nil, zero value otherwise.

### GetHasGenericOk

`func (o *DrugPatentCliffRow) GetHasGenericOk() (*bool, bool)`

GetHasGenericOk returns a tuple with the HasGeneric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGeneric

`func (o *DrugPatentCliffRow) SetHasGeneric(v bool)`

SetHasGeneric sets HasGeneric field to given value.

### HasHasGeneric

`func (o *DrugPatentCliffRow) HasHasGeneric() bool`

HasHasGeneric returns a boolean if a field has been set.

### SetHasGenericNil

`func (o *DrugPatentCliffRow) SetHasGenericNil(b bool)`

 SetHasGenericNil sets the value for HasGeneric to be an explicit nil

### UnsetHasGeneric
`func (o *DrugPatentCliffRow) UnsetHasGeneric()`

UnsetHasGeneric ensures that no value is present for HasGeneric, not even an explicit nil
### GetBiosimilarAvailable

`func (o *DrugPatentCliffRow) GetBiosimilarAvailable() bool`

GetBiosimilarAvailable returns the BiosimilarAvailable field if non-nil, zero value otherwise.

### GetBiosimilarAvailableOk

`func (o *DrugPatentCliffRow) GetBiosimilarAvailableOk() (*bool, bool)`

GetBiosimilarAvailableOk returns a tuple with the BiosimilarAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosimilarAvailable

`func (o *DrugPatentCliffRow) SetBiosimilarAvailable(v bool)`

SetBiosimilarAvailable sets BiosimilarAvailable field to given value.

### HasBiosimilarAvailable

`func (o *DrugPatentCliffRow) HasBiosimilarAvailable() bool`

HasBiosimilarAvailable returns a boolean if a field has been set.

### SetBiosimilarAvailableNil

`func (o *DrugPatentCliffRow) SetBiosimilarAvailableNil(b bool)`

 SetBiosimilarAvailableNil sets the value for BiosimilarAvailable to be an explicit nil

### UnsetBiosimilarAvailable
`func (o *DrugPatentCliffRow) UnsetBiosimilarAvailable()`

UnsetBiosimilarAvailable ensures that no value is present for BiosimilarAvailable, not even an explicit nil
### GetApplicantName

`func (o *DrugPatentCliffRow) GetApplicantName() string`

GetApplicantName returns the ApplicantName field if non-nil, zero value otherwise.

### GetApplicantNameOk

`func (o *DrugPatentCliffRow) GetApplicantNameOk() (*string, bool)`

GetApplicantNameOk returns a tuple with the ApplicantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantName

`func (o *DrugPatentCliffRow) SetApplicantName(v string)`

SetApplicantName sets ApplicantName field to given value.

### HasApplicantName

`func (o *DrugPatentCliffRow) HasApplicantName() bool`

HasApplicantName returns a boolean if a field has been set.

### SetApplicantNameNil

`func (o *DrugPatentCliffRow) SetApplicantNameNil(b bool)`

 SetApplicantNameNil sets the value for ApplicantName to be an explicit nil

### UnsetApplicantName
`func (o *DrugPatentCliffRow) UnsetApplicantName()`

UnsetApplicantName ensures that no value is present for ApplicantName, not even an explicit nil
### GetApplicantOrgId

`func (o *DrugPatentCliffRow) GetApplicantOrgId() string`

GetApplicantOrgId returns the ApplicantOrgId field if non-nil, zero value otherwise.

### GetApplicantOrgIdOk

`func (o *DrugPatentCliffRow) GetApplicantOrgIdOk() (*string, bool)`

GetApplicantOrgIdOk returns a tuple with the ApplicantOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantOrgId

`func (o *DrugPatentCliffRow) SetApplicantOrgId(v string)`

SetApplicantOrgId sets ApplicantOrgId field to given value.

### HasApplicantOrgId

`func (o *DrugPatentCliffRow) HasApplicantOrgId() bool`

HasApplicantOrgId returns a boolean if a field has been set.

### SetApplicantOrgIdNil

`func (o *DrugPatentCliffRow) SetApplicantOrgIdNil(b bool)`

 SetApplicantOrgIdNil sets the value for ApplicantOrgId to be an explicit nil

### UnsetApplicantOrgId
`func (o *DrugPatentCliffRow) UnsetApplicantOrgId()`

UnsetApplicantOrgId ensures that no value is present for ApplicantOrgId, not even an explicit nil
### GetApplicantCompanyId

`func (o *DrugPatentCliffRow) GetApplicantCompanyId() string`

GetApplicantCompanyId returns the ApplicantCompanyId field if non-nil, zero value otherwise.

### GetApplicantCompanyIdOk

`func (o *DrugPatentCliffRow) GetApplicantCompanyIdOk() (*string, bool)`

GetApplicantCompanyIdOk returns a tuple with the ApplicantCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantCompanyId

`func (o *DrugPatentCliffRow) SetApplicantCompanyId(v string)`

SetApplicantCompanyId sets ApplicantCompanyId field to given value.

### HasApplicantCompanyId

`func (o *DrugPatentCliffRow) HasApplicantCompanyId() bool`

HasApplicantCompanyId returns a boolean if a field has been set.

### SetApplicantCompanyIdNil

`func (o *DrugPatentCliffRow) SetApplicantCompanyIdNil(b bool)`

 SetApplicantCompanyIdNil sets the value for ApplicantCompanyId to be an explicit nil

### UnsetApplicantCompanyId
`func (o *DrugPatentCliffRow) UnsetApplicantCompanyId()`

UnsetApplicantCompanyId ensures that no value is present for ApplicantCompanyId, not even an explicit nil
### GetApplicantCompanyName

`func (o *DrugPatentCliffRow) GetApplicantCompanyName() string`

GetApplicantCompanyName returns the ApplicantCompanyName field if non-nil, zero value otherwise.

### GetApplicantCompanyNameOk

`func (o *DrugPatentCliffRow) GetApplicantCompanyNameOk() (*string, bool)`

GetApplicantCompanyNameOk returns a tuple with the ApplicantCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantCompanyName

`func (o *DrugPatentCliffRow) SetApplicantCompanyName(v string)`

SetApplicantCompanyName sets ApplicantCompanyName field to given value.

### HasApplicantCompanyName

`func (o *DrugPatentCliffRow) HasApplicantCompanyName() bool`

HasApplicantCompanyName returns a boolean if a field has been set.

### SetApplicantCompanyNameNil

`func (o *DrugPatentCliffRow) SetApplicantCompanyNameNil(b bool)`

 SetApplicantCompanyNameNil sets the value for ApplicantCompanyName to be an explicit nil

### UnsetApplicantCompanyName
`func (o *DrugPatentCliffRow) UnsetApplicantCompanyName()`

UnsetApplicantCompanyName ensures that no value is present for ApplicantCompanyName, not even an explicit nil
### GetApplicantGroupId

`func (o *DrugPatentCliffRow) GetApplicantGroupId() string`

GetApplicantGroupId returns the ApplicantGroupId field if non-nil, zero value otherwise.

### GetApplicantGroupIdOk

`func (o *DrugPatentCliffRow) GetApplicantGroupIdOk() (*string, bool)`

GetApplicantGroupIdOk returns a tuple with the ApplicantGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantGroupId

`func (o *DrugPatentCliffRow) SetApplicantGroupId(v string)`

SetApplicantGroupId sets ApplicantGroupId field to given value.

### HasApplicantGroupId

`func (o *DrugPatentCliffRow) HasApplicantGroupId() bool`

HasApplicantGroupId returns a boolean if a field has been set.

### SetApplicantGroupIdNil

`func (o *DrugPatentCliffRow) SetApplicantGroupIdNil(b bool)`

 SetApplicantGroupIdNil sets the value for ApplicantGroupId to be an explicit nil

### UnsetApplicantGroupId
`func (o *DrugPatentCliffRow) UnsetApplicantGroupId()`

UnsetApplicantGroupId ensures that no value is present for ApplicantGroupId, not even an explicit nil
### GetApplicantGroupName

`func (o *DrugPatentCliffRow) GetApplicantGroupName() string`

GetApplicantGroupName returns the ApplicantGroupName field if non-nil, zero value otherwise.

### GetApplicantGroupNameOk

`func (o *DrugPatentCliffRow) GetApplicantGroupNameOk() (*string, bool)`

GetApplicantGroupNameOk returns a tuple with the ApplicantGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantGroupName

`func (o *DrugPatentCliffRow) SetApplicantGroupName(v string)`

SetApplicantGroupName sets ApplicantGroupName field to given value.

### HasApplicantGroupName

`func (o *DrugPatentCliffRow) HasApplicantGroupName() bool`

HasApplicantGroupName returns a boolean if a field has been set.

### SetApplicantGroupNameNil

`func (o *DrugPatentCliffRow) SetApplicantGroupNameNil(b bool)`

 SetApplicantGroupNameNil sets the value for ApplicantGroupName to be an explicit nil

### UnsetApplicantGroupName
`func (o *DrugPatentCliffRow) UnsetApplicantGroupName()`

UnsetApplicantGroupName ensures that no value is present for ApplicantGroupName, not even an explicit nil
### GetApplicantMatchScope

`func (o *DrugPatentCliffRow) GetApplicantMatchScope() string`

GetApplicantMatchScope returns the ApplicantMatchScope field if non-nil, zero value otherwise.

### GetApplicantMatchScopeOk

`func (o *DrugPatentCliffRow) GetApplicantMatchScopeOk() (*string, bool)`

GetApplicantMatchScopeOk returns a tuple with the ApplicantMatchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantMatchScope

`func (o *DrugPatentCliffRow) SetApplicantMatchScope(v string)`

SetApplicantMatchScope sets ApplicantMatchScope field to given value.

### HasApplicantMatchScope

`func (o *DrugPatentCliffRow) HasApplicantMatchScope() bool`

HasApplicantMatchScope returns a boolean if a field has been set.

### SetApplicantMatchScopeNil

`func (o *DrugPatentCliffRow) SetApplicantMatchScopeNil(b bool)`

 SetApplicantMatchScopeNil sets the value for ApplicantMatchScope to be an explicit nil

### UnsetApplicantMatchScope
`func (o *DrugPatentCliffRow) UnsetApplicantMatchScope()`

UnsetApplicantMatchScope ensures that no value is present for ApplicantMatchScope, not even an explicit nil
### GetMedicareSpend

`func (o *DrugPatentCliffRow) GetMedicareSpend() float32`

GetMedicareSpend returns the MedicareSpend field if non-nil, zero value otherwise.

### GetMedicareSpendOk

`func (o *DrugPatentCliffRow) GetMedicareSpendOk() (*float32, bool)`

GetMedicareSpendOk returns a tuple with the MedicareSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedicareSpend

`func (o *DrugPatentCliffRow) SetMedicareSpend(v float32)`

SetMedicareSpend sets MedicareSpend field to given value.

### HasMedicareSpend

`func (o *DrugPatentCliffRow) HasMedicareSpend() bool`

HasMedicareSpend returns a boolean if a field has been set.

### SetMedicareSpendNil

`func (o *DrugPatentCliffRow) SetMedicareSpendNil(b bool)`

 SetMedicareSpendNil sets the value for MedicareSpend to be an explicit nil

### UnsetMedicareSpend
`func (o *DrugPatentCliffRow) UnsetMedicareSpend()`

UnsetMedicareSpend ensures that no value is present for MedicareSpend, not even an explicit nil
### GetMedicareSpendYear

`func (o *DrugPatentCliffRow) GetMedicareSpendYear() int32`

GetMedicareSpendYear returns the MedicareSpendYear field if non-nil, zero value otherwise.

### GetMedicareSpendYearOk

`func (o *DrugPatentCliffRow) GetMedicareSpendYearOk() (*int32, bool)`

GetMedicareSpendYearOk returns a tuple with the MedicareSpendYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedicareSpendYear

`func (o *DrugPatentCliffRow) SetMedicareSpendYear(v int32)`

SetMedicareSpendYear sets MedicareSpendYear field to given value.

### HasMedicareSpendYear

`func (o *DrugPatentCliffRow) HasMedicareSpendYear() bool`

HasMedicareSpendYear returns a boolean if a field has been set.

### SetMedicareSpendYearNil

`func (o *DrugPatentCliffRow) SetMedicareSpendYearNil(b bool)`

 SetMedicareSpendYearNil sets the value for MedicareSpendYear to be an explicit nil

### UnsetMedicareSpendYear
`func (o *DrugPatentCliffRow) UnsetMedicareSpendYear()`

UnsetMedicareSpendYear ensures that no value is present for MedicareSpendYear, not even an explicit nil
### GetDisclosedRevenueUsdMillions

`func (o *DrugPatentCliffRow) GetDisclosedRevenueUsdMillions() float32`

GetDisclosedRevenueUsdMillions returns the DisclosedRevenueUsdMillions field if non-nil, zero value otherwise.

### GetDisclosedRevenueUsdMillionsOk

`func (o *DrugPatentCliffRow) GetDisclosedRevenueUsdMillionsOk() (*float32, bool)`

GetDisclosedRevenueUsdMillionsOk returns a tuple with the DisclosedRevenueUsdMillions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosedRevenueUsdMillions

`func (o *DrugPatentCliffRow) SetDisclosedRevenueUsdMillions(v float32)`

SetDisclosedRevenueUsdMillions sets DisclosedRevenueUsdMillions field to given value.

### HasDisclosedRevenueUsdMillions

`func (o *DrugPatentCliffRow) HasDisclosedRevenueUsdMillions() bool`

HasDisclosedRevenueUsdMillions returns a boolean if a field has been set.

### SetDisclosedRevenueUsdMillionsNil

`func (o *DrugPatentCliffRow) SetDisclosedRevenueUsdMillionsNil(b bool)`

 SetDisclosedRevenueUsdMillionsNil sets the value for DisclosedRevenueUsdMillions to be an explicit nil

### UnsetDisclosedRevenueUsdMillions
`func (o *DrugPatentCliffRow) UnsetDisclosedRevenueUsdMillions()`

UnsetDisclosedRevenueUsdMillions ensures that no value is present for DisclosedRevenueUsdMillions, not even an explicit nil
### GetDisclosedRevenueYear

`func (o *DrugPatentCliffRow) GetDisclosedRevenueYear() int32`

GetDisclosedRevenueYear returns the DisclosedRevenueYear field if non-nil, zero value otherwise.

### GetDisclosedRevenueYearOk

`func (o *DrugPatentCliffRow) GetDisclosedRevenueYearOk() (*int32, bool)`

GetDisclosedRevenueYearOk returns a tuple with the DisclosedRevenueYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosedRevenueYear

`func (o *DrugPatentCliffRow) SetDisclosedRevenueYear(v int32)`

SetDisclosedRevenueYear sets DisclosedRevenueYear field to given value.

### HasDisclosedRevenueYear

`func (o *DrugPatentCliffRow) HasDisclosedRevenueYear() bool`

HasDisclosedRevenueYear returns a boolean if a field has been set.

### SetDisclosedRevenueYearNil

`func (o *DrugPatentCliffRow) SetDisclosedRevenueYearNil(b bool)`

 SetDisclosedRevenueYearNil sets the value for DisclosedRevenueYear to be an explicit nil

### UnsetDisclosedRevenueYear
`func (o *DrugPatentCliffRow) UnsetDisclosedRevenueYear()`

UnsetDisclosedRevenueYear ensures that no value is present for DisclosedRevenueYear, not even an explicit nil
### GetSizeUsd

`func (o *DrugPatentCliffRow) GetSizeUsd() float32`

GetSizeUsd returns the SizeUsd field if non-nil, zero value otherwise.

### GetSizeUsdOk

`func (o *DrugPatentCliffRow) GetSizeUsdOk() (*float32, bool)`

GetSizeUsdOk returns a tuple with the SizeUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeUsd

`func (o *DrugPatentCliffRow) SetSizeUsd(v float32)`

SetSizeUsd sets SizeUsd field to given value.

### HasSizeUsd

`func (o *DrugPatentCliffRow) HasSizeUsd() bool`

HasSizeUsd returns a boolean if a field has been set.

### SetSizeUsdNil

`func (o *DrugPatentCliffRow) SetSizeUsdNil(b bool)`

 SetSizeUsdNil sets the value for SizeUsd to be an explicit nil

### UnsetSizeUsd
`func (o *DrugPatentCliffRow) UnsetSizeUsd()`

UnsetSizeUsd ensures that no value is present for SizeUsd, not even an explicit nil
### GetSizeBasis

`func (o *DrugPatentCliffRow) GetSizeBasis() string`

GetSizeBasis returns the SizeBasis field if non-nil, zero value otherwise.

### GetSizeBasisOk

`func (o *DrugPatentCliffRow) GetSizeBasisOk() (*string, bool)`

GetSizeBasisOk returns a tuple with the SizeBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBasis

`func (o *DrugPatentCliffRow) SetSizeBasis(v string)`

SetSizeBasis sets SizeBasis field to given value.

### HasSizeBasis

`func (o *DrugPatentCliffRow) HasSizeBasis() bool`

HasSizeBasis returns a boolean if a field has been set.

### SetSizeBasisNil

`func (o *DrugPatentCliffRow) SetSizeBasisNil(b bool)`

 SetSizeBasisNil sets the value for SizeBasis to be an explicit nil

### UnsetSizeBasis
`func (o *DrugPatentCliffRow) UnsetSizeBasis()`

UnsetSizeBasis ensures that no value is present for SizeBasis, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


