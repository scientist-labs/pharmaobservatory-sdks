# DrugVariantsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**MemberId** | Pointer to **NullableString** | The id of the folded drug record: the drug_id other datasets carry. | [optional] 
**MemberName** | Pointer to **NullableString** | The folded record&#39;s display name. | [optional] 
**Role** | Pointer to **NullableString** | single for a one-to-one fold; component_of_compound for a combination arm listing the molecule. | [optional] 
**Method** | Pointer to **NullableString** | How the fold was made (for example rxnorm_synonym, unii_synonym, brand, override, registered_substance). | [optional] 
**Tier** | Pointer to **NullableInt32** | The resolver tier that made the fold (1 to 4); 4 is the model backstop. | [optional] 
**Confidence** | Pointer to **NullableFloat64** | The resolver&#39;s confidence for the method, from 0 to 1. | [optional] 
**MemberTrials** | Pointer to **NullableInt32** | Registered trials of the member record. | [optional] 
**MemberMaxPhase** | Pointer to **NullableInt32** | The member&#39;s highest development phase. | [optional] 
**MemberApproved** | Pointer to **NullableBool** | True when the member is approved somewhere. | [optional] 
**MemberBiologic** | Pointer to **NullableBool** | True when the member is a biologic substance. | [optional] 

## Methods

### NewDrugVariantsRow

`func NewDrugVariantsRow() *DrugVariantsRow`

NewDrugVariantsRow instantiates a new DrugVariantsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugVariantsRowWithDefaults

`func NewDrugVariantsRowWithDefaults() *DrugVariantsRow`

NewDrugVariantsRowWithDefaults instantiates a new DrugVariantsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugVariantsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugVariantsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugVariantsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugVariantsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugVariantsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugVariantsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetMemberId

`func (o *DrugVariantsRow) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *DrugVariantsRow) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *DrugVariantsRow) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *DrugVariantsRow) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### SetMemberIdNil

`func (o *DrugVariantsRow) SetMemberIdNil(b bool)`

 SetMemberIdNil sets the value for MemberId to be an explicit nil

### UnsetMemberId
`func (o *DrugVariantsRow) UnsetMemberId()`

UnsetMemberId ensures that no value is present for MemberId, not even an explicit nil
### GetMemberName

`func (o *DrugVariantsRow) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *DrugVariantsRow) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *DrugVariantsRow) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.

### HasMemberName

`func (o *DrugVariantsRow) HasMemberName() bool`

HasMemberName returns a boolean if a field has been set.

### SetMemberNameNil

`func (o *DrugVariantsRow) SetMemberNameNil(b bool)`

 SetMemberNameNil sets the value for MemberName to be an explicit nil

### UnsetMemberName
`func (o *DrugVariantsRow) UnsetMemberName()`

UnsetMemberName ensures that no value is present for MemberName, not even an explicit nil
### GetRole

`func (o *DrugVariantsRow) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DrugVariantsRow) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DrugVariantsRow) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *DrugVariantsRow) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *DrugVariantsRow) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *DrugVariantsRow) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetMethod

`func (o *DrugVariantsRow) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *DrugVariantsRow) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *DrugVariantsRow) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *DrugVariantsRow) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *DrugVariantsRow) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *DrugVariantsRow) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetTier

`func (o *DrugVariantsRow) GetTier() int32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *DrugVariantsRow) GetTierOk() (*int32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *DrugVariantsRow) SetTier(v int32)`

SetTier sets Tier field to given value.

### HasTier

`func (o *DrugVariantsRow) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *DrugVariantsRow) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *DrugVariantsRow) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil
### GetConfidence

`func (o *DrugVariantsRow) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *DrugVariantsRow) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *DrugVariantsRow) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *DrugVariantsRow) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *DrugVariantsRow) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *DrugVariantsRow) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetMemberTrials

`func (o *DrugVariantsRow) GetMemberTrials() int32`

GetMemberTrials returns the MemberTrials field if non-nil, zero value otherwise.

### GetMemberTrialsOk

`func (o *DrugVariantsRow) GetMemberTrialsOk() (*int32, bool)`

GetMemberTrialsOk returns a tuple with the MemberTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberTrials

`func (o *DrugVariantsRow) SetMemberTrials(v int32)`

SetMemberTrials sets MemberTrials field to given value.

### HasMemberTrials

`func (o *DrugVariantsRow) HasMemberTrials() bool`

HasMemberTrials returns a boolean if a field has been set.

### SetMemberTrialsNil

`func (o *DrugVariantsRow) SetMemberTrialsNil(b bool)`

 SetMemberTrialsNil sets the value for MemberTrials to be an explicit nil

### UnsetMemberTrials
`func (o *DrugVariantsRow) UnsetMemberTrials()`

UnsetMemberTrials ensures that no value is present for MemberTrials, not even an explicit nil
### GetMemberMaxPhase

`func (o *DrugVariantsRow) GetMemberMaxPhase() int32`

GetMemberMaxPhase returns the MemberMaxPhase field if non-nil, zero value otherwise.

### GetMemberMaxPhaseOk

`func (o *DrugVariantsRow) GetMemberMaxPhaseOk() (*int32, bool)`

GetMemberMaxPhaseOk returns a tuple with the MemberMaxPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberMaxPhase

`func (o *DrugVariantsRow) SetMemberMaxPhase(v int32)`

SetMemberMaxPhase sets MemberMaxPhase field to given value.

### HasMemberMaxPhase

`func (o *DrugVariantsRow) HasMemberMaxPhase() bool`

HasMemberMaxPhase returns a boolean if a field has been set.

### SetMemberMaxPhaseNil

`func (o *DrugVariantsRow) SetMemberMaxPhaseNil(b bool)`

 SetMemberMaxPhaseNil sets the value for MemberMaxPhase to be an explicit nil

### UnsetMemberMaxPhase
`func (o *DrugVariantsRow) UnsetMemberMaxPhase()`

UnsetMemberMaxPhase ensures that no value is present for MemberMaxPhase, not even an explicit nil
### GetMemberApproved

`func (o *DrugVariantsRow) GetMemberApproved() bool`

GetMemberApproved returns the MemberApproved field if non-nil, zero value otherwise.

### GetMemberApprovedOk

`func (o *DrugVariantsRow) GetMemberApprovedOk() (*bool, bool)`

GetMemberApprovedOk returns a tuple with the MemberApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberApproved

`func (o *DrugVariantsRow) SetMemberApproved(v bool)`

SetMemberApproved sets MemberApproved field to given value.

### HasMemberApproved

`func (o *DrugVariantsRow) HasMemberApproved() bool`

HasMemberApproved returns a boolean if a field has been set.

### SetMemberApprovedNil

`func (o *DrugVariantsRow) SetMemberApprovedNil(b bool)`

 SetMemberApprovedNil sets the value for MemberApproved to be an explicit nil

### UnsetMemberApproved
`func (o *DrugVariantsRow) UnsetMemberApproved()`

UnsetMemberApproved ensures that no value is present for MemberApproved, not even an explicit nil
### GetMemberBiologic

`func (o *DrugVariantsRow) GetMemberBiologic() bool`

GetMemberBiologic returns the MemberBiologic field if non-nil, zero value otherwise.

### GetMemberBiologicOk

`func (o *DrugVariantsRow) GetMemberBiologicOk() (*bool, bool)`

GetMemberBiologicOk returns a tuple with the MemberBiologic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberBiologic

`func (o *DrugVariantsRow) SetMemberBiologic(v bool)`

SetMemberBiologic sets MemberBiologic field to given value.

### HasMemberBiologic

`func (o *DrugVariantsRow) HasMemberBiologic() bool`

HasMemberBiologic returns a boolean if a field has been set.

### SetMemberBiologicNil

`func (o *DrugVariantsRow) SetMemberBiologicNil(b bool)`

 SetMemberBiologicNil sets the value for MemberBiologic to be an explicit nil

### UnsetMemberBiologic
`func (o *DrugVariantsRow) UnsetMemberBiologic()`

UnsetMemberBiologic ensures that no value is present for MemberBiologic, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


