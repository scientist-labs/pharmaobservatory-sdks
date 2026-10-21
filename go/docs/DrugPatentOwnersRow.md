# DrugPatentOwnersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**OwnerOrgId** | Pointer to **NullableString** | Legal-entity level: the owning organization record as filed. | [optional] 
**OwnerName** | Pointer to **NullableString** | The owner&#39;s name. | [optional] 
**OwnerCompanyId** | Pointer to **NullableString** | Company level: the owner&#39;s company (null when the owner is not in the spine). | [optional] 
**OwnerCompanyName** | Pointer to **NullableString** | Display name of the owner&#39;s company. | [optional] 
**OwnerGroupId** | Pointer to **NullableString** | Group level: the owner&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**OwnerGroupName** | Pointer to **NullableString** | Display name of the owner&#39;s group. | [optional] 
**OrgInSpine** | Pointer to **NullableBool** | True when the owner is an organization record in the spine. | [optional] 
**NPatents** | Pointer to **NullableInt32** | Distinct declared patents the owner holds for the drug. | [optional] 
**EarliestExpiry** | Pointer to **NullableString** | The earliest declared expiry. | [optional] 
**LatestExpiry** | Pointer to **NullableString** | The latest declared expiry. | [optional] 
**BestMatchConfidence** | Pointer to **NullableString** | The best assignee-to-organization match confidence (high, med, low). | [optional] 
**LinkBases** | Pointer to **[]string** | The records that tie the patents to the drug (orange_book, uspto_pte). | [optional] 

## Methods

### NewDrugPatentOwnersRow

`func NewDrugPatentOwnersRow() *DrugPatentOwnersRow`

NewDrugPatentOwnersRow instantiates a new DrugPatentOwnersRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugPatentOwnersRowWithDefaults

`func NewDrugPatentOwnersRowWithDefaults() *DrugPatentOwnersRow`

NewDrugPatentOwnersRowWithDefaults instantiates a new DrugPatentOwnersRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugPatentOwnersRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugPatentOwnersRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugPatentOwnersRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugPatentOwnersRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugPatentOwnersRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugPatentOwnersRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetOwnerOrgId

`func (o *DrugPatentOwnersRow) GetOwnerOrgId() string`

GetOwnerOrgId returns the OwnerOrgId field if non-nil, zero value otherwise.

### GetOwnerOrgIdOk

`func (o *DrugPatentOwnersRow) GetOwnerOrgIdOk() (*string, bool)`

GetOwnerOrgIdOk returns a tuple with the OwnerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrgId

`func (o *DrugPatentOwnersRow) SetOwnerOrgId(v string)`

SetOwnerOrgId sets OwnerOrgId field to given value.

### HasOwnerOrgId

`func (o *DrugPatentOwnersRow) HasOwnerOrgId() bool`

HasOwnerOrgId returns a boolean if a field has been set.

### SetOwnerOrgIdNil

`func (o *DrugPatentOwnersRow) SetOwnerOrgIdNil(b bool)`

 SetOwnerOrgIdNil sets the value for OwnerOrgId to be an explicit nil

### UnsetOwnerOrgId
`func (o *DrugPatentOwnersRow) UnsetOwnerOrgId()`

UnsetOwnerOrgId ensures that no value is present for OwnerOrgId, not even an explicit nil
### GetOwnerName

`func (o *DrugPatentOwnersRow) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *DrugPatentOwnersRow) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *DrugPatentOwnersRow) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *DrugPatentOwnersRow) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### SetOwnerNameNil

`func (o *DrugPatentOwnersRow) SetOwnerNameNil(b bool)`

 SetOwnerNameNil sets the value for OwnerName to be an explicit nil

### UnsetOwnerName
`func (o *DrugPatentOwnersRow) UnsetOwnerName()`

UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil
### GetOwnerCompanyId

`func (o *DrugPatentOwnersRow) GetOwnerCompanyId() string`

GetOwnerCompanyId returns the OwnerCompanyId field if non-nil, zero value otherwise.

### GetOwnerCompanyIdOk

`func (o *DrugPatentOwnersRow) GetOwnerCompanyIdOk() (*string, bool)`

GetOwnerCompanyIdOk returns a tuple with the OwnerCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerCompanyId

`func (o *DrugPatentOwnersRow) SetOwnerCompanyId(v string)`

SetOwnerCompanyId sets OwnerCompanyId field to given value.

### HasOwnerCompanyId

`func (o *DrugPatentOwnersRow) HasOwnerCompanyId() bool`

HasOwnerCompanyId returns a boolean if a field has been set.

### SetOwnerCompanyIdNil

`func (o *DrugPatentOwnersRow) SetOwnerCompanyIdNil(b bool)`

 SetOwnerCompanyIdNil sets the value for OwnerCompanyId to be an explicit nil

### UnsetOwnerCompanyId
`func (o *DrugPatentOwnersRow) UnsetOwnerCompanyId()`

UnsetOwnerCompanyId ensures that no value is present for OwnerCompanyId, not even an explicit nil
### GetOwnerCompanyName

`func (o *DrugPatentOwnersRow) GetOwnerCompanyName() string`

GetOwnerCompanyName returns the OwnerCompanyName field if non-nil, zero value otherwise.

### GetOwnerCompanyNameOk

`func (o *DrugPatentOwnersRow) GetOwnerCompanyNameOk() (*string, bool)`

GetOwnerCompanyNameOk returns a tuple with the OwnerCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerCompanyName

`func (o *DrugPatentOwnersRow) SetOwnerCompanyName(v string)`

SetOwnerCompanyName sets OwnerCompanyName field to given value.

### HasOwnerCompanyName

`func (o *DrugPatentOwnersRow) HasOwnerCompanyName() bool`

HasOwnerCompanyName returns a boolean if a field has been set.

### SetOwnerCompanyNameNil

`func (o *DrugPatentOwnersRow) SetOwnerCompanyNameNil(b bool)`

 SetOwnerCompanyNameNil sets the value for OwnerCompanyName to be an explicit nil

### UnsetOwnerCompanyName
`func (o *DrugPatentOwnersRow) UnsetOwnerCompanyName()`

UnsetOwnerCompanyName ensures that no value is present for OwnerCompanyName, not even an explicit nil
### GetOwnerGroupId

`func (o *DrugPatentOwnersRow) GetOwnerGroupId() string`

GetOwnerGroupId returns the OwnerGroupId field if non-nil, zero value otherwise.

### GetOwnerGroupIdOk

`func (o *DrugPatentOwnersRow) GetOwnerGroupIdOk() (*string, bool)`

GetOwnerGroupIdOk returns a tuple with the OwnerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroupId

`func (o *DrugPatentOwnersRow) SetOwnerGroupId(v string)`

SetOwnerGroupId sets OwnerGroupId field to given value.

### HasOwnerGroupId

`func (o *DrugPatentOwnersRow) HasOwnerGroupId() bool`

HasOwnerGroupId returns a boolean if a field has been set.

### SetOwnerGroupIdNil

`func (o *DrugPatentOwnersRow) SetOwnerGroupIdNil(b bool)`

 SetOwnerGroupIdNil sets the value for OwnerGroupId to be an explicit nil

### UnsetOwnerGroupId
`func (o *DrugPatentOwnersRow) UnsetOwnerGroupId()`

UnsetOwnerGroupId ensures that no value is present for OwnerGroupId, not even an explicit nil
### GetOwnerGroupName

`func (o *DrugPatentOwnersRow) GetOwnerGroupName() string`

GetOwnerGroupName returns the OwnerGroupName field if non-nil, zero value otherwise.

### GetOwnerGroupNameOk

`func (o *DrugPatentOwnersRow) GetOwnerGroupNameOk() (*string, bool)`

GetOwnerGroupNameOk returns a tuple with the OwnerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroupName

`func (o *DrugPatentOwnersRow) SetOwnerGroupName(v string)`

SetOwnerGroupName sets OwnerGroupName field to given value.

### HasOwnerGroupName

`func (o *DrugPatentOwnersRow) HasOwnerGroupName() bool`

HasOwnerGroupName returns a boolean if a field has been set.

### SetOwnerGroupNameNil

`func (o *DrugPatentOwnersRow) SetOwnerGroupNameNil(b bool)`

 SetOwnerGroupNameNil sets the value for OwnerGroupName to be an explicit nil

### UnsetOwnerGroupName
`func (o *DrugPatentOwnersRow) UnsetOwnerGroupName()`

UnsetOwnerGroupName ensures that no value is present for OwnerGroupName, not even an explicit nil
### GetOrgInSpine

`func (o *DrugPatentOwnersRow) GetOrgInSpine() bool`

GetOrgInSpine returns the OrgInSpine field if non-nil, zero value otherwise.

### GetOrgInSpineOk

`func (o *DrugPatentOwnersRow) GetOrgInSpineOk() (*bool, bool)`

GetOrgInSpineOk returns a tuple with the OrgInSpine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgInSpine

`func (o *DrugPatentOwnersRow) SetOrgInSpine(v bool)`

SetOrgInSpine sets OrgInSpine field to given value.

### HasOrgInSpine

`func (o *DrugPatentOwnersRow) HasOrgInSpine() bool`

HasOrgInSpine returns a boolean if a field has been set.

### SetOrgInSpineNil

`func (o *DrugPatentOwnersRow) SetOrgInSpineNil(b bool)`

 SetOrgInSpineNil sets the value for OrgInSpine to be an explicit nil

### UnsetOrgInSpine
`func (o *DrugPatentOwnersRow) UnsetOrgInSpine()`

UnsetOrgInSpine ensures that no value is present for OrgInSpine, not even an explicit nil
### GetNPatents

`func (o *DrugPatentOwnersRow) GetNPatents() int32`

GetNPatents returns the NPatents field if non-nil, zero value otherwise.

### GetNPatentsOk

`func (o *DrugPatentOwnersRow) GetNPatentsOk() (*int32, bool)`

GetNPatentsOk returns a tuple with the NPatents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPatents

`func (o *DrugPatentOwnersRow) SetNPatents(v int32)`

SetNPatents sets NPatents field to given value.

### HasNPatents

`func (o *DrugPatentOwnersRow) HasNPatents() bool`

HasNPatents returns a boolean if a field has been set.

### SetNPatentsNil

`func (o *DrugPatentOwnersRow) SetNPatentsNil(b bool)`

 SetNPatentsNil sets the value for NPatents to be an explicit nil

### UnsetNPatents
`func (o *DrugPatentOwnersRow) UnsetNPatents()`

UnsetNPatents ensures that no value is present for NPatents, not even an explicit nil
### GetEarliestExpiry

`func (o *DrugPatentOwnersRow) GetEarliestExpiry() string`

GetEarliestExpiry returns the EarliestExpiry field if non-nil, zero value otherwise.

### GetEarliestExpiryOk

`func (o *DrugPatentOwnersRow) GetEarliestExpiryOk() (*string, bool)`

GetEarliestExpiryOk returns a tuple with the EarliestExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestExpiry

`func (o *DrugPatentOwnersRow) SetEarliestExpiry(v string)`

SetEarliestExpiry sets EarliestExpiry field to given value.

### HasEarliestExpiry

`func (o *DrugPatentOwnersRow) HasEarliestExpiry() bool`

HasEarliestExpiry returns a boolean if a field has been set.

### SetEarliestExpiryNil

`func (o *DrugPatentOwnersRow) SetEarliestExpiryNil(b bool)`

 SetEarliestExpiryNil sets the value for EarliestExpiry to be an explicit nil

### UnsetEarliestExpiry
`func (o *DrugPatentOwnersRow) UnsetEarliestExpiry()`

UnsetEarliestExpiry ensures that no value is present for EarliestExpiry, not even an explicit nil
### GetLatestExpiry

`func (o *DrugPatentOwnersRow) GetLatestExpiry() string`

GetLatestExpiry returns the LatestExpiry field if non-nil, zero value otherwise.

### GetLatestExpiryOk

`func (o *DrugPatentOwnersRow) GetLatestExpiryOk() (*string, bool)`

GetLatestExpiryOk returns a tuple with the LatestExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestExpiry

`func (o *DrugPatentOwnersRow) SetLatestExpiry(v string)`

SetLatestExpiry sets LatestExpiry field to given value.

### HasLatestExpiry

`func (o *DrugPatentOwnersRow) HasLatestExpiry() bool`

HasLatestExpiry returns a boolean if a field has been set.

### SetLatestExpiryNil

`func (o *DrugPatentOwnersRow) SetLatestExpiryNil(b bool)`

 SetLatestExpiryNil sets the value for LatestExpiry to be an explicit nil

### UnsetLatestExpiry
`func (o *DrugPatentOwnersRow) UnsetLatestExpiry()`

UnsetLatestExpiry ensures that no value is present for LatestExpiry, not even an explicit nil
### GetBestMatchConfidence

`func (o *DrugPatentOwnersRow) GetBestMatchConfidence() string`

GetBestMatchConfidence returns the BestMatchConfidence field if non-nil, zero value otherwise.

### GetBestMatchConfidenceOk

`func (o *DrugPatentOwnersRow) GetBestMatchConfidenceOk() (*string, bool)`

GetBestMatchConfidenceOk returns a tuple with the BestMatchConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestMatchConfidence

`func (o *DrugPatentOwnersRow) SetBestMatchConfidence(v string)`

SetBestMatchConfidence sets BestMatchConfidence field to given value.

### HasBestMatchConfidence

`func (o *DrugPatentOwnersRow) HasBestMatchConfidence() bool`

HasBestMatchConfidence returns a boolean if a field has been set.

### SetBestMatchConfidenceNil

`func (o *DrugPatentOwnersRow) SetBestMatchConfidenceNil(b bool)`

 SetBestMatchConfidenceNil sets the value for BestMatchConfidence to be an explicit nil

### UnsetBestMatchConfidence
`func (o *DrugPatentOwnersRow) UnsetBestMatchConfidence()`

UnsetBestMatchConfidence ensures that no value is present for BestMatchConfidence, not even an explicit nil
### GetLinkBases

`func (o *DrugPatentOwnersRow) GetLinkBases() []string`

GetLinkBases returns the LinkBases field if non-nil, zero value otherwise.

### GetLinkBasesOk

`func (o *DrugPatentOwnersRow) GetLinkBasesOk() (*[]string, bool)`

GetLinkBasesOk returns a tuple with the LinkBases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkBases

`func (o *DrugPatentOwnersRow) SetLinkBases(v []string)`

SetLinkBases sets LinkBases field to given value.

### HasLinkBases

`func (o *DrugPatentOwnersRow) HasLinkBases() bool`

HasLinkBases returns a boolean if a field has been set.

### SetLinkBasesNil

`func (o *DrugPatentOwnersRow) SetLinkBasesNil(b bool)`

 SetLinkBasesNil sets the value for LinkBases to be an explicit nil

### UnsetLinkBases
`func (o *DrugPatentOwnersRow) UnsetLinkBases()`

UnsetLinkBases ensures that no value is present for LinkBases, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


