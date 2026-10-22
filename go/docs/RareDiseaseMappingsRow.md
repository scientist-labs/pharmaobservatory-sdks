# RareDiseaseMappingsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**XrefAssertionId** | Pointer to **NullableString** | The identifier of the cross-reference assertion. | [optional] 
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**TargetNamespace** | Pointer to **NullableString** | The vocabulary the target code belongs to (ICD-10, ICD-11, MESH, MONDO, UMLS, GARD, ORPHA and others). | [optional] 
**TargetCode** | Pointer to **NullableString** | The code in the target vocabulary. | [optional] 
**TargetCurie** | Pointer to **NullableString** | The target as a prefixed identifier (namespace and code). | [optional] 
**MappingRelation** | Pointer to **NullableString** | The relation the source states from the disease to the target: exact, narrower, broader, undetermined or unspecified. | [optional] 
**AssertionSource** | Pointer to **NullableString** | The file the assertion was read from: ordo, orphadata_product1 or mondo. | [optional] 
**Confidence** | Pointer to **NullableFloat32** | The source&#39;s confidence in the assertion, where it states one. | [optional] 
**ConditionId** | Pointer to **NullableString** | The MeSH descriptor id of the matching condition, on MESH rows; the entry key of the conditions dataset. | [optional] 
**ConditionName** | Pointer to **NullableString** | The matching condition&#39;s name, on MESH rows. | [optional] 
**SourceRelease** | Pointer to **NullableString** | The Orphadata Science release the row was read from. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The URL of the source file the row was read from. | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence the assertion&#39;s file is published under. | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRareDiseaseMappingsRow

`func NewRareDiseaseMappingsRow() *RareDiseaseMappingsRow`

NewRareDiseaseMappingsRow instantiates a new RareDiseaseMappingsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseaseMappingsRowWithDefaults

`func NewRareDiseaseMappingsRowWithDefaults() *RareDiseaseMappingsRow`

NewRareDiseaseMappingsRowWithDefaults instantiates a new RareDiseaseMappingsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXrefAssertionId

`func (o *RareDiseaseMappingsRow) GetXrefAssertionId() string`

GetXrefAssertionId returns the XrefAssertionId field if non-nil, zero value otherwise.

### GetXrefAssertionIdOk

`func (o *RareDiseaseMappingsRow) GetXrefAssertionIdOk() (*string, bool)`

GetXrefAssertionIdOk returns a tuple with the XrefAssertionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXrefAssertionId

`func (o *RareDiseaseMappingsRow) SetXrefAssertionId(v string)`

SetXrefAssertionId sets XrefAssertionId field to given value.

### HasXrefAssertionId

`func (o *RareDiseaseMappingsRow) HasXrefAssertionId() bool`

HasXrefAssertionId returns a boolean if a field has been set.

### SetXrefAssertionIdNil

`func (o *RareDiseaseMappingsRow) SetXrefAssertionIdNil(b bool)`

 SetXrefAssertionIdNil sets the value for XrefAssertionId to be an explicit nil

### UnsetXrefAssertionId
`func (o *RareDiseaseMappingsRow) UnsetXrefAssertionId()`

UnsetXrefAssertionId ensures that no value is present for XrefAssertionId, not even an explicit nil
### GetOrphaCode

`func (o *RareDiseaseMappingsRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseaseMappingsRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseaseMappingsRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseaseMappingsRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseaseMappingsRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseaseMappingsRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetTargetNamespace

`func (o *RareDiseaseMappingsRow) GetTargetNamespace() string`

GetTargetNamespace returns the TargetNamespace field if non-nil, zero value otherwise.

### GetTargetNamespaceOk

`func (o *RareDiseaseMappingsRow) GetTargetNamespaceOk() (*string, bool)`

GetTargetNamespaceOk returns a tuple with the TargetNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNamespace

`func (o *RareDiseaseMappingsRow) SetTargetNamespace(v string)`

SetTargetNamespace sets TargetNamespace field to given value.

### HasTargetNamespace

`func (o *RareDiseaseMappingsRow) HasTargetNamespace() bool`

HasTargetNamespace returns a boolean if a field has been set.

### SetTargetNamespaceNil

`func (o *RareDiseaseMappingsRow) SetTargetNamespaceNil(b bool)`

 SetTargetNamespaceNil sets the value for TargetNamespace to be an explicit nil

### UnsetTargetNamespace
`func (o *RareDiseaseMappingsRow) UnsetTargetNamespace()`

UnsetTargetNamespace ensures that no value is present for TargetNamespace, not even an explicit nil
### GetTargetCode

`func (o *RareDiseaseMappingsRow) GetTargetCode() string`

GetTargetCode returns the TargetCode field if non-nil, zero value otherwise.

### GetTargetCodeOk

`func (o *RareDiseaseMappingsRow) GetTargetCodeOk() (*string, bool)`

GetTargetCodeOk returns a tuple with the TargetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCode

`func (o *RareDiseaseMappingsRow) SetTargetCode(v string)`

SetTargetCode sets TargetCode field to given value.

### HasTargetCode

`func (o *RareDiseaseMappingsRow) HasTargetCode() bool`

HasTargetCode returns a boolean if a field has been set.

### SetTargetCodeNil

`func (o *RareDiseaseMappingsRow) SetTargetCodeNil(b bool)`

 SetTargetCodeNil sets the value for TargetCode to be an explicit nil

### UnsetTargetCode
`func (o *RareDiseaseMappingsRow) UnsetTargetCode()`

UnsetTargetCode ensures that no value is present for TargetCode, not even an explicit nil
### GetTargetCurie

`func (o *RareDiseaseMappingsRow) GetTargetCurie() string`

GetTargetCurie returns the TargetCurie field if non-nil, zero value otherwise.

### GetTargetCurieOk

`func (o *RareDiseaseMappingsRow) GetTargetCurieOk() (*string, bool)`

GetTargetCurieOk returns a tuple with the TargetCurie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCurie

`func (o *RareDiseaseMappingsRow) SetTargetCurie(v string)`

SetTargetCurie sets TargetCurie field to given value.

### HasTargetCurie

`func (o *RareDiseaseMappingsRow) HasTargetCurie() bool`

HasTargetCurie returns a boolean if a field has been set.

### SetTargetCurieNil

`func (o *RareDiseaseMappingsRow) SetTargetCurieNil(b bool)`

 SetTargetCurieNil sets the value for TargetCurie to be an explicit nil

### UnsetTargetCurie
`func (o *RareDiseaseMappingsRow) UnsetTargetCurie()`

UnsetTargetCurie ensures that no value is present for TargetCurie, not even an explicit nil
### GetMappingRelation

`func (o *RareDiseaseMappingsRow) GetMappingRelation() string`

GetMappingRelation returns the MappingRelation field if non-nil, zero value otherwise.

### GetMappingRelationOk

`func (o *RareDiseaseMappingsRow) GetMappingRelationOk() (*string, bool)`

GetMappingRelationOk returns a tuple with the MappingRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingRelation

`func (o *RareDiseaseMappingsRow) SetMappingRelation(v string)`

SetMappingRelation sets MappingRelation field to given value.

### HasMappingRelation

`func (o *RareDiseaseMappingsRow) HasMappingRelation() bool`

HasMappingRelation returns a boolean if a field has been set.

### SetMappingRelationNil

`func (o *RareDiseaseMappingsRow) SetMappingRelationNil(b bool)`

 SetMappingRelationNil sets the value for MappingRelation to be an explicit nil

### UnsetMappingRelation
`func (o *RareDiseaseMappingsRow) UnsetMappingRelation()`

UnsetMappingRelation ensures that no value is present for MappingRelation, not even an explicit nil
### GetAssertionSource

`func (o *RareDiseaseMappingsRow) GetAssertionSource() string`

GetAssertionSource returns the AssertionSource field if non-nil, zero value otherwise.

### GetAssertionSourceOk

`func (o *RareDiseaseMappingsRow) GetAssertionSourceOk() (*string, bool)`

GetAssertionSourceOk returns a tuple with the AssertionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionSource

`func (o *RareDiseaseMappingsRow) SetAssertionSource(v string)`

SetAssertionSource sets AssertionSource field to given value.

### HasAssertionSource

`func (o *RareDiseaseMappingsRow) HasAssertionSource() bool`

HasAssertionSource returns a boolean if a field has been set.

### SetAssertionSourceNil

`func (o *RareDiseaseMappingsRow) SetAssertionSourceNil(b bool)`

 SetAssertionSourceNil sets the value for AssertionSource to be an explicit nil

### UnsetAssertionSource
`func (o *RareDiseaseMappingsRow) UnsetAssertionSource()`

UnsetAssertionSource ensures that no value is present for AssertionSource, not even an explicit nil
### GetConfidence

`func (o *RareDiseaseMappingsRow) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *RareDiseaseMappingsRow) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *RareDiseaseMappingsRow) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *RareDiseaseMappingsRow) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *RareDiseaseMappingsRow) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *RareDiseaseMappingsRow) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetConditionId

`func (o *RareDiseaseMappingsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *RareDiseaseMappingsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *RareDiseaseMappingsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *RareDiseaseMappingsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *RareDiseaseMappingsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *RareDiseaseMappingsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *RareDiseaseMappingsRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *RareDiseaseMappingsRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *RareDiseaseMappingsRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *RareDiseaseMappingsRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *RareDiseaseMappingsRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *RareDiseaseMappingsRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetSourceRelease

`func (o *RareDiseaseMappingsRow) GetSourceRelease() string`

GetSourceRelease returns the SourceRelease field if non-nil, zero value otherwise.

### GetSourceReleaseOk

`func (o *RareDiseaseMappingsRow) GetSourceReleaseOk() (*string, bool)`

GetSourceReleaseOk returns a tuple with the SourceRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRelease

`func (o *RareDiseaseMappingsRow) SetSourceRelease(v string)`

SetSourceRelease sets SourceRelease field to given value.

### HasSourceRelease

`func (o *RareDiseaseMappingsRow) HasSourceRelease() bool`

HasSourceRelease returns a boolean if a field has been set.

### SetSourceReleaseNil

`func (o *RareDiseaseMappingsRow) SetSourceReleaseNil(b bool)`

 SetSourceReleaseNil sets the value for SourceRelease to be an explicit nil

### UnsetSourceRelease
`func (o *RareDiseaseMappingsRow) UnsetSourceRelease()`

UnsetSourceRelease ensures that no value is present for SourceRelease, not even an explicit nil
### GetSourceUrl

`func (o *RareDiseaseMappingsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *RareDiseaseMappingsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *RareDiseaseMappingsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *RareDiseaseMappingsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *RareDiseaseMappingsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *RareDiseaseMappingsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetLicenseName

`func (o *RareDiseaseMappingsRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *RareDiseaseMappingsRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *RareDiseaseMappingsRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *RareDiseaseMappingsRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *RareDiseaseMappingsRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *RareDiseaseMappingsRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *RareDiseaseMappingsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RareDiseaseMappingsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RareDiseaseMappingsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RareDiseaseMappingsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RareDiseaseMappingsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RareDiseaseMappingsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


