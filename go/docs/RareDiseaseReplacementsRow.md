# RareDiseaseReplacementsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplacementAssertionId** | Pointer to **NullableString** | The identifier of the replacement assertion. | [optional] 
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**PreferredName** | Pointer to **NullableString** | The deprecated disease&#39;s preferred English name. | [optional] 
**ReplacementOrphaCode** | Pointer to **NullableString** | The ORPHA code of the concept that replaces the deprecated one. | [optional] 
**ReplacementPreferredName** | Pointer to **NullableString** | The replacement concept&#39;s preferred English name. | [optional] 
**ReplacementStatus** | Pointer to **NullableString** | The replacement concept&#39;s lifecycle state. | [optional] 
**Predicate** | Pointer to **NullableString** | moved_to when Orphanet moved the concept to the replacement, referred_to when it referred the concept there. | [optional] 
**AssertionSource** | Pointer to **NullableString** | The Orphanet file the assertion was read from (ordo or orphadata_product1). | [optional] 
**SourceRelease** | Pointer to **NullableString** | The Orphadata Science release the row was read from. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The URL of the source file the row was read from. | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence the assertion&#39;s file is published under. | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRareDiseaseReplacementsRow

`func NewRareDiseaseReplacementsRow() *RareDiseaseReplacementsRow`

NewRareDiseaseReplacementsRow instantiates a new RareDiseaseReplacementsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseaseReplacementsRowWithDefaults

`func NewRareDiseaseReplacementsRowWithDefaults() *RareDiseaseReplacementsRow`

NewRareDiseaseReplacementsRowWithDefaults instantiates a new RareDiseaseReplacementsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplacementAssertionId

`func (o *RareDiseaseReplacementsRow) GetReplacementAssertionId() string`

GetReplacementAssertionId returns the ReplacementAssertionId field if non-nil, zero value otherwise.

### GetReplacementAssertionIdOk

`func (o *RareDiseaseReplacementsRow) GetReplacementAssertionIdOk() (*string, bool)`

GetReplacementAssertionIdOk returns a tuple with the ReplacementAssertionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementAssertionId

`func (o *RareDiseaseReplacementsRow) SetReplacementAssertionId(v string)`

SetReplacementAssertionId sets ReplacementAssertionId field to given value.

### HasReplacementAssertionId

`func (o *RareDiseaseReplacementsRow) HasReplacementAssertionId() bool`

HasReplacementAssertionId returns a boolean if a field has been set.

### SetReplacementAssertionIdNil

`func (o *RareDiseaseReplacementsRow) SetReplacementAssertionIdNil(b bool)`

 SetReplacementAssertionIdNil sets the value for ReplacementAssertionId to be an explicit nil

### UnsetReplacementAssertionId
`func (o *RareDiseaseReplacementsRow) UnsetReplacementAssertionId()`

UnsetReplacementAssertionId ensures that no value is present for ReplacementAssertionId, not even an explicit nil
### GetOrphaCode

`func (o *RareDiseaseReplacementsRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseaseReplacementsRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseaseReplacementsRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseaseReplacementsRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseaseReplacementsRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseaseReplacementsRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetPreferredName

`func (o *RareDiseaseReplacementsRow) GetPreferredName() string`

GetPreferredName returns the PreferredName field if non-nil, zero value otherwise.

### GetPreferredNameOk

`func (o *RareDiseaseReplacementsRow) GetPreferredNameOk() (*string, bool)`

GetPreferredNameOk returns a tuple with the PreferredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredName

`func (o *RareDiseaseReplacementsRow) SetPreferredName(v string)`

SetPreferredName sets PreferredName field to given value.

### HasPreferredName

`func (o *RareDiseaseReplacementsRow) HasPreferredName() bool`

HasPreferredName returns a boolean if a field has been set.

### SetPreferredNameNil

`func (o *RareDiseaseReplacementsRow) SetPreferredNameNil(b bool)`

 SetPreferredNameNil sets the value for PreferredName to be an explicit nil

### UnsetPreferredName
`func (o *RareDiseaseReplacementsRow) UnsetPreferredName()`

UnsetPreferredName ensures that no value is present for PreferredName, not even an explicit nil
### GetReplacementOrphaCode

`func (o *RareDiseaseReplacementsRow) GetReplacementOrphaCode() string`

GetReplacementOrphaCode returns the ReplacementOrphaCode field if non-nil, zero value otherwise.

### GetReplacementOrphaCodeOk

`func (o *RareDiseaseReplacementsRow) GetReplacementOrphaCodeOk() (*string, bool)`

GetReplacementOrphaCodeOk returns a tuple with the ReplacementOrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementOrphaCode

`func (o *RareDiseaseReplacementsRow) SetReplacementOrphaCode(v string)`

SetReplacementOrphaCode sets ReplacementOrphaCode field to given value.

### HasReplacementOrphaCode

`func (o *RareDiseaseReplacementsRow) HasReplacementOrphaCode() bool`

HasReplacementOrphaCode returns a boolean if a field has been set.

### SetReplacementOrphaCodeNil

`func (o *RareDiseaseReplacementsRow) SetReplacementOrphaCodeNil(b bool)`

 SetReplacementOrphaCodeNil sets the value for ReplacementOrphaCode to be an explicit nil

### UnsetReplacementOrphaCode
`func (o *RareDiseaseReplacementsRow) UnsetReplacementOrphaCode()`

UnsetReplacementOrphaCode ensures that no value is present for ReplacementOrphaCode, not even an explicit nil
### GetReplacementPreferredName

`func (o *RareDiseaseReplacementsRow) GetReplacementPreferredName() string`

GetReplacementPreferredName returns the ReplacementPreferredName field if non-nil, zero value otherwise.

### GetReplacementPreferredNameOk

`func (o *RareDiseaseReplacementsRow) GetReplacementPreferredNameOk() (*string, bool)`

GetReplacementPreferredNameOk returns a tuple with the ReplacementPreferredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementPreferredName

`func (o *RareDiseaseReplacementsRow) SetReplacementPreferredName(v string)`

SetReplacementPreferredName sets ReplacementPreferredName field to given value.

### HasReplacementPreferredName

`func (o *RareDiseaseReplacementsRow) HasReplacementPreferredName() bool`

HasReplacementPreferredName returns a boolean if a field has been set.

### SetReplacementPreferredNameNil

`func (o *RareDiseaseReplacementsRow) SetReplacementPreferredNameNil(b bool)`

 SetReplacementPreferredNameNil sets the value for ReplacementPreferredName to be an explicit nil

### UnsetReplacementPreferredName
`func (o *RareDiseaseReplacementsRow) UnsetReplacementPreferredName()`

UnsetReplacementPreferredName ensures that no value is present for ReplacementPreferredName, not even an explicit nil
### GetReplacementStatus

`func (o *RareDiseaseReplacementsRow) GetReplacementStatus() string`

GetReplacementStatus returns the ReplacementStatus field if non-nil, zero value otherwise.

### GetReplacementStatusOk

`func (o *RareDiseaseReplacementsRow) GetReplacementStatusOk() (*string, bool)`

GetReplacementStatusOk returns a tuple with the ReplacementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementStatus

`func (o *RareDiseaseReplacementsRow) SetReplacementStatus(v string)`

SetReplacementStatus sets ReplacementStatus field to given value.

### HasReplacementStatus

`func (o *RareDiseaseReplacementsRow) HasReplacementStatus() bool`

HasReplacementStatus returns a boolean if a field has been set.

### SetReplacementStatusNil

`func (o *RareDiseaseReplacementsRow) SetReplacementStatusNil(b bool)`

 SetReplacementStatusNil sets the value for ReplacementStatus to be an explicit nil

### UnsetReplacementStatus
`func (o *RareDiseaseReplacementsRow) UnsetReplacementStatus()`

UnsetReplacementStatus ensures that no value is present for ReplacementStatus, not even an explicit nil
### GetPredicate

`func (o *RareDiseaseReplacementsRow) GetPredicate() string`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *RareDiseaseReplacementsRow) GetPredicateOk() (*string, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *RareDiseaseReplacementsRow) SetPredicate(v string)`

SetPredicate sets Predicate field to given value.

### HasPredicate

`func (o *RareDiseaseReplacementsRow) HasPredicate() bool`

HasPredicate returns a boolean if a field has been set.

### SetPredicateNil

`func (o *RareDiseaseReplacementsRow) SetPredicateNil(b bool)`

 SetPredicateNil sets the value for Predicate to be an explicit nil

### UnsetPredicate
`func (o *RareDiseaseReplacementsRow) UnsetPredicate()`

UnsetPredicate ensures that no value is present for Predicate, not even an explicit nil
### GetAssertionSource

`func (o *RareDiseaseReplacementsRow) GetAssertionSource() string`

GetAssertionSource returns the AssertionSource field if non-nil, zero value otherwise.

### GetAssertionSourceOk

`func (o *RareDiseaseReplacementsRow) GetAssertionSourceOk() (*string, bool)`

GetAssertionSourceOk returns a tuple with the AssertionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionSource

`func (o *RareDiseaseReplacementsRow) SetAssertionSource(v string)`

SetAssertionSource sets AssertionSource field to given value.

### HasAssertionSource

`func (o *RareDiseaseReplacementsRow) HasAssertionSource() bool`

HasAssertionSource returns a boolean if a field has been set.

### SetAssertionSourceNil

`func (o *RareDiseaseReplacementsRow) SetAssertionSourceNil(b bool)`

 SetAssertionSourceNil sets the value for AssertionSource to be an explicit nil

### UnsetAssertionSource
`func (o *RareDiseaseReplacementsRow) UnsetAssertionSource()`

UnsetAssertionSource ensures that no value is present for AssertionSource, not even an explicit nil
### GetSourceRelease

`func (o *RareDiseaseReplacementsRow) GetSourceRelease() string`

GetSourceRelease returns the SourceRelease field if non-nil, zero value otherwise.

### GetSourceReleaseOk

`func (o *RareDiseaseReplacementsRow) GetSourceReleaseOk() (*string, bool)`

GetSourceReleaseOk returns a tuple with the SourceRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRelease

`func (o *RareDiseaseReplacementsRow) SetSourceRelease(v string)`

SetSourceRelease sets SourceRelease field to given value.

### HasSourceRelease

`func (o *RareDiseaseReplacementsRow) HasSourceRelease() bool`

HasSourceRelease returns a boolean if a field has been set.

### SetSourceReleaseNil

`func (o *RareDiseaseReplacementsRow) SetSourceReleaseNil(b bool)`

 SetSourceReleaseNil sets the value for SourceRelease to be an explicit nil

### UnsetSourceRelease
`func (o *RareDiseaseReplacementsRow) UnsetSourceRelease()`

UnsetSourceRelease ensures that no value is present for SourceRelease, not even an explicit nil
### GetSourceUrl

`func (o *RareDiseaseReplacementsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *RareDiseaseReplacementsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *RareDiseaseReplacementsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *RareDiseaseReplacementsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *RareDiseaseReplacementsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *RareDiseaseReplacementsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetLicenseName

`func (o *RareDiseaseReplacementsRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *RareDiseaseReplacementsRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *RareDiseaseReplacementsRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *RareDiseaseReplacementsRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *RareDiseaseReplacementsRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *RareDiseaseReplacementsRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *RareDiseaseReplacementsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RareDiseaseReplacementsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RareDiseaseReplacementsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RareDiseaseReplacementsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RareDiseaseReplacementsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RareDiseaseReplacementsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


