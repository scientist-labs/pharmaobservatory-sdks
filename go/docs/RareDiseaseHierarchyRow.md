# RareDiseaseHierarchyRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HierarchyAssertionId** | Pointer to **NullableString** | The identifier of the subclass assertion. | [optional] 
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the child disease the assertion is entered on. | [optional] 
**PreferredName** | Pointer to **NullableString** | The child disease&#39;s preferred English name. | [optional] 
**ParentOrphaCode** | Pointer to **NullableString** | The ORPHA code of the direct parent disease. | [optional] 
**ParentPreferredName** | Pointer to **NullableString** | The parent disease&#39;s preferred English name. | [optional] 
**ParentStatus** | Pointer to **NullableString** | The parent concept&#39;s lifecycle state. | [optional] 
**ChildStatus** | Pointer to **NullableString** | The child concept&#39;s lifecycle state. | [optional] 
**Predicate** | Pointer to **NullableString** | The relation asserted from child to parent (subclass_of). | [optional] 
**SourceRelease** | Pointer to **NullableString** | The Orphadata Science release the row was read from. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The URL of the source file the row was read from. | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence of the ontology release. | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRareDiseaseHierarchyRow

`func NewRareDiseaseHierarchyRow() *RareDiseaseHierarchyRow`

NewRareDiseaseHierarchyRow instantiates a new RareDiseaseHierarchyRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseaseHierarchyRowWithDefaults

`func NewRareDiseaseHierarchyRowWithDefaults() *RareDiseaseHierarchyRow`

NewRareDiseaseHierarchyRowWithDefaults instantiates a new RareDiseaseHierarchyRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHierarchyAssertionId

`func (o *RareDiseaseHierarchyRow) GetHierarchyAssertionId() string`

GetHierarchyAssertionId returns the HierarchyAssertionId field if non-nil, zero value otherwise.

### GetHierarchyAssertionIdOk

`func (o *RareDiseaseHierarchyRow) GetHierarchyAssertionIdOk() (*string, bool)`

GetHierarchyAssertionIdOk returns a tuple with the HierarchyAssertionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyAssertionId

`func (o *RareDiseaseHierarchyRow) SetHierarchyAssertionId(v string)`

SetHierarchyAssertionId sets HierarchyAssertionId field to given value.

### HasHierarchyAssertionId

`func (o *RareDiseaseHierarchyRow) HasHierarchyAssertionId() bool`

HasHierarchyAssertionId returns a boolean if a field has been set.

### SetHierarchyAssertionIdNil

`func (o *RareDiseaseHierarchyRow) SetHierarchyAssertionIdNil(b bool)`

 SetHierarchyAssertionIdNil sets the value for HierarchyAssertionId to be an explicit nil

### UnsetHierarchyAssertionId
`func (o *RareDiseaseHierarchyRow) UnsetHierarchyAssertionId()`

UnsetHierarchyAssertionId ensures that no value is present for HierarchyAssertionId, not even an explicit nil
### GetOrphaCode

`func (o *RareDiseaseHierarchyRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseaseHierarchyRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseaseHierarchyRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseaseHierarchyRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseaseHierarchyRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseaseHierarchyRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetPreferredName

`func (o *RareDiseaseHierarchyRow) GetPreferredName() string`

GetPreferredName returns the PreferredName field if non-nil, zero value otherwise.

### GetPreferredNameOk

`func (o *RareDiseaseHierarchyRow) GetPreferredNameOk() (*string, bool)`

GetPreferredNameOk returns a tuple with the PreferredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredName

`func (o *RareDiseaseHierarchyRow) SetPreferredName(v string)`

SetPreferredName sets PreferredName field to given value.

### HasPreferredName

`func (o *RareDiseaseHierarchyRow) HasPreferredName() bool`

HasPreferredName returns a boolean if a field has been set.

### SetPreferredNameNil

`func (o *RareDiseaseHierarchyRow) SetPreferredNameNil(b bool)`

 SetPreferredNameNil sets the value for PreferredName to be an explicit nil

### UnsetPreferredName
`func (o *RareDiseaseHierarchyRow) UnsetPreferredName()`

UnsetPreferredName ensures that no value is present for PreferredName, not even an explicit nil
### GetParentOrphaCode

`func (o *RareDiseaseHierarchyRow) GetParentOrphaCode() string`

GetParentOrphaCode returns the ParentOrphaCode field if non-nil, zero value otherwise.

### GetParentOrphaCodeOk

`func (o *RareDiseaseHierarchyRow) GetParentOrphaCodeOk() (*string, bool)`

GetParentOrphaCodeOk returns a tuple with the ParentOrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrphaCode

`func (o *RareDiseaseHierarchyRow) SetParentOrphaCode(v string)`

SetParentOrphaCode sets ParentOrphaCode field to given value.

### HasParentOrphaCode

`func (o *RareDiseaseHierarchyRow) HasParentOrphaCode() bool`

HasParentOrphaCode returns a boolean if a field has been set.

### SetParentOrphaCodeNil

`func (o *RareDiseaseHierarchyRow) SetParentOrphaCodeNil(b bool)`

 SetParentOrphaCodeNil sets the value for ParentOrphaCode to be an explicit nil

### UnsetParentOrphaCode
`func (o *RareDiseaseHierarchyRow) UnsetParentOrphaCode()`

UnsetParentOrphaCode ensures that no value is present for ParentOrphaCode, not even an explicit nil
### GetParentPreferredName

`func (o *RareDiseaseHierarchyRow) GetParentPreferredName() string`

GetParentPreferredName returns the ParentPreferredName field if non-nil, zero value otherwise.

### GetParentPreferredNameOk

`func (o *RareDiseaseHierarchyRow) GetParentPreferredNameOk() (*string, bool)`

GetParentPreferredNameOk returns a tuple with the ParentPreferredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPreferredName

`func (o *RareDiseaseHierarchyRow) SetParentPreferredName(v string)`

SetParentPreferredName sets ParentPreferredName field to given value.

### HasParentPreferredName

`func (o *RareDiseaseHierarchyRow) HasParentPreferredName() bool`

HasParentPreferredName returns a boolean if a field has been set.

### SetParentPreferredNameNil

`func (o *RareDiseaseHierarchyRow) SetParentPreferredNameNil(b bool)`

 SetParentPreferredNameNil sets the value for ParentPreferredName to be an explicit nil

### UnsetParentPreferredName
`func (o *RareDiseaseHierarchyRow) UnsetParentPreferredName()`

UnsetParentPreferredName ensures that no value is present for ParentPreferredName, not even an explicit nil
### GetParentStatus

`func (o *RareDiseaseHierarchyRow) GetParentStatus() string`

GetParentStatus returns the ParentStatus field if non-nil, zero value otherwise.

### GetParentStatusOk

`func (o *RareDiseaseHierarchyRow) GetParentStatusOk() (*string, bool)`

GetParentStatusOk returns a tuple with the ParentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStatus

`func (o *RareDiseaseHierarchyRow) SetParentStatus(v string)`

SetParentStatus sets ParentStatus field to given value.

### HasParentStatus

`func (o *RareDiseaseHierarchyRow) HasParentStatus() bool`

HasParentStatus returns a boolean if a field has been set.

### SetParentStatusNil

`func (o *RareDiseaseHierarchyRow) SetParentStatusNil(b bool)`

 SetParentStatusNil sets the value for ParentStatus to be an explicit nil

### UnsetParentStatus
`func (o *RareDiseaseHierarchyRow) UnsetParentStatus()`

UnsetParentStatus ensures that no value is present for ParentStatus, not even an explicit nil
### GetChildStatus

`func (o *RareDiseaseHierarchyRow) GetChildStatus() string`

GetChildStatus returns the ChildStatus field if non-nil, zero value otherwise.

### GetChildStatusOk

`func (o *RareDiseaseHierarchyRow) GetChildStatusOk() (*string, bool)`

GetChildStatusOk returns a tuple with the ChildStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildStatus

`func (o *RareDiseaseHierarchyRow) SetChildStatus(v string)`

SetChildStatus sets ChildStatus field to given value.

### HasChildStatus

`func (o *RareDiseaseHierarchyRow) HasChildStatus() bool`

HasChildStatus returns a boolean if a field has been set.

### SetChildStatusNil

`func (o *RareDiseaseHierarchyRow) SetChildStatusNil(b bool)`

 SetChildStatusNil sets the value for ChildStatus to be an explicit nil

### UnsetChildStatus
`func (o *RareDiseaseHierarchyRow) UnsetChildStatus()`

UnsetChildStatus ensures that no value is present for ChildStatus, not even an explicit nil
### GetPredicate

`func (o *RareDiseaseHierarchyRow) GetPredicate() string`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *RareDiseaseHierarchyRow) GetPredicateOk() (*string, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *RareDiseaseHierarchyRow) SetPredicate(v string)`

SetPredicate sets Predicate field to given value.

### HasPredicate

`func (o *RareDiseaseHierarchyRow) HasPredicate() bool`

HasPredicate returns a boolean if a field has been set.

### SetPredicateNil

`func (o *RareDiseaseHierarchyRow) SetPredicateNil(b bool)`

 SetPredicateNil sets the value for Predicate to be an explicit nil

### UnsetPredicate
`func (o *RareDiseaseHierarchyRow) UnsetPredicate()`

UnsetPredicate ensures that no value is present for Predicate, not even an explicit nil
### GetSourceRelease

`func (o *RareDiseaseHierarchyRow) GetSourceRelease() string`

GetSourceRelease returns the SourceRelease field if non-nil, zero value otherwise.

### GetSourceReleaseOk

`func (o *RareDiseaseHierarchyRow) GetSourceReleaseOk() (*string, bool)`

GetSourceReleaseOk returns a tuple with the SourceRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRelease

`func (o *RareDiseaseHierarchyRow) SetSourceRelease(v string)`

SetSourceRelease sets SourceRelease field to given value.

### HasSourceRelease

`func (o *RareDiseaseHierarchyRow) HasSourceRelease() bool`

HasSourceRelease returns a boolean if a field has been set.

### SetSourceReleaseNil

`func (o *RareDiseaseHierarchyRow) SetSourceReleaseNil(b bool)`

 SetSourceReleaseNil sets the value for SourceRelease to be an explicit nil

### UnsetSourceRelease
`func (o *RareDiseaseHierarchyRow) UnsetSourceRelease()`

UnsetSourceRelease ensures that no value is present for SourceRelease, not even an explicit nil
### GetSourceUrl

`func (o *RareDiseaseHierarchyRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *RareDiseaseHierarchyRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *RareDiseaseHierarchyRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *RareDiseaseHierarchyRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *RareDiseaseHierarchyRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *RareDiseaseHierarchyRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetLicenseName

`func (o *RareDiseaseHierarchyRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *RareDiseaseHierarchyRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *RareDiseaseHierarchyRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *RareDiseaseHierarchyRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *RareDiseaseHierarchyRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *RareDiseaseHierarchyRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *RareDiseaseHierarchyRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RareDiseaseHierarchyRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RareDiseaseHierarchyRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RareDiseaseHierarchyRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RareDiseaseHierarchyRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RareDiseaseHierarchyRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


