# RwdUsNetworkMembershipsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MembershipAssertionId** | Pointer to **NullableString** | The assertion id: a hash of the network, institution, provider uuid and match method. | [optional] 
**NetworkSource** | Pointer to **NullableString** | The source the member rows come from. | [optional] 
**NetworkKey** | Pointer to **NullableString** | The network key, the entry key of the rwd_us_networks object type. | [optional] 
**NetworkName** | Pointer to **NullableString** | The network&#39;s name. | [optional] 
**NetworkOperator** | Pointer to **NullableString** | The operating body. | [optional] 
**MembershipKind** | Pointer to **NullableString** | The strongest agreement the institution holds. membership_kind is the strongest agreement the institution holds (dta_signatory, else dua_signatory, else lhba_signatory), never data contributor. | [optional] 
**HasDta** | Pointer to **NullableBool** | The institution holds at least one published DTA (Data Transfer Agreement) row. | [optional] 
**HasDua** | Pointer to **NullableBool** | The institution holds at least one published DUA (Data Use Agreement) row. | [optional] 
**HasLhba** | Pointer to **NullableBool** | The institution holds at least one agreement published with the type code LHBA, which N3C does not expand. | [optional] 
**HasCurrentDta** | Pointer to **NullableBool** | A DTA row&#39;s published term end is on or after the retrieval date. | [optional] 
**HasCurrentDua** | Pointer to **NullableBool** | A DUA row is within its published term. | [optional] 
**HasCurrentLhba** | Pointer to **NullableBool** | An LHBA row is within its published term. | [optional] 
**FirstExecutedOn** | Pointer to **NullableString** | The earliest published execution date over the institution&#39;s agreement rows. | [optional] 
**LatestTermEndsOn** | Pointer to **NullableString** | The latest published term end over the institution&#39;s rows. | [optional] 
**AgreementCount** | Pointer to **NullableInt32** | Published agreement rows folded into this member. | [optional] 
**InstitutionName** | Pointer to **NullableString** | The institution as N3C publishes it (the ROR name). | [optional] 
**InstitutionNameKey** | Pointer to **NullableString** | The normalized name key, the join key to other name-keyed datasets. | [optional] 
**RorId** | Pointer to **NullableString** | The bare ROR id (Research Organization Registry) N3C published; null when it published a domain instead. | [optional] 
**InstitutionDomain** | Pointer to **NullableString** | The organization&#39;s email domain as published, an organization domain and not a personal address. | [optional] 
**IsLead** | Pointer to **NullableBool** | Null for N3C, which has no lead concept; false would be a false No. | [optional] 
**AsOf** | Pointer to **NullableString** | The retrieval date of the agreement snapshot. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The API endpoint the agreements came from. | [optional] 
**RetrievedAt** | Pointer to **NullableTime** | When the snapshot was retrieved, in UTC. | [optional] 
**OrgId** | Pointer to **NullableString** | The organization spine id the member resolved to, or null. Carried, never minted. | [optional] 
**OrgName** | Pointer to **NullableString** | The spine organization&#39;s name. | [optional] 
**OrgMatchRung** | Pointer to **NullableString** | How the member reached the organization: ror_exact (the source&#39;s own ROR id), xref, match_keys, ambiguous or unresolved. | [optional] 
**OrgMatchTier** | Pointer to **NullableString** | The tier of the organization match (ror_exact is high, stated as the source&#39;s own identifier). | [optional] 
**OrgMatchConfidence** | Pointer to **NullableFloat32** | The organization match confidence (0.95 on ror_exact). | [optional] 
**ProviderId** | Pointer to **NullableString** | The RX supplier id as text, the entry key of the suppliers object type. Null on an unmatched row. | [optional] 
**ProviderUuid** | Pointer to **NullableString** | The Scientist.com public provider uuid. Null on an unmatched row. | [optional] 
**ProviderSlug** | Pointer to **NullableString** | The supplier&#39;s marketplace slug. | [optional] 
**ProviderName** | Pointer to **NullableString** | The supplier&#39;s name. | [optional] 
**MatchMethod** | Pointer to **NullableString** | How the supplier was matched: org_member_exact (through the resolved organization), rx_name_key_exact (an exact name match, not a verified identity) or unmatched (no provider). | [optional] 
**AlsoMatchedByName** | Pointer to **NullableBool** | True on an org_member_exact row whose provider also matches by name; null on an unmatched row. | [optional] 
**MatchConfidence** | Pointer to **NullableFloat32** | The organization match confidence passed through on org_member_exact; null on rx_name_key_exact and unmatched. | [optional] 
**MemberTotal** | Pointer to **NullableInt32** | Members of the network, the register-wide denominator printed beside any count. | [optional] 
**MembersWithCurrentDta** | Pointer to **NullableInt32** | Network members with a DTA within its published term. | [optional] 
**MembersWithCurrentDua** | Pointer to **NullableInt32** | Network members with a DUA within its published term. | [optional] 
**MembersResolvedOrg** | Pointer to **NullableInt32** | Network members resolved to an organization. | [optional] 
**MembersMatchedRx** | Pointer to **NullableInt32** | Network members matched to at least one RX provider. | [optional] 
**Attribution** | Pointer to **NullableString** | The credit line every surface carrying this data prints. | [optional] 

## Methods

### NewRwdUsNetworkMembershipsRow

`func NewRwdUsNetworkMembershipsRow() *RwdUsNetworkMembershipsRow`

NewRwdUsNetworkMembershipsRow instantiates a new RwdUsNetworkMembershipsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdUsNetworkMembershipsRowWithDefaults

`func NewRwdUsNetworkMembershipsRowWithDefaults() *RwdUsNetworkMembershipsRow`

NewRwdUsNetworkMembershipsRowWithDefaults instantiates a new RwdUsNetworkMembershipsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembershipAssertionId

`func (o *RwdUsNetworkMembershipsRow) GetMembershipAssertionId() string`

GetMembershipAssertionId returns the MembershipAssertionId field if non-nil, zero value otherwise.

### GetMembershipAssertionIdOk

`func (o *RwdUsNetworkMembershipsRow) GetMembershipAssertionIdOk() (*string, bool)`

GetMembershipAssertionIdOk returns a tuple with the MembershipAssertionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipAssertionId

`func (o *RwdUsNetworkMembershipsRow) SetMembershipAssertionId(v string)`

SetMembershipAssertionId sets MembershipAssertionId field to given value.

### HasMembershipAssertionId

`func (o *RwdUsNetworkMembershipsRow) HasMembershipAssertionId() bool`

HasMembershipAssertionId returns a boolean if a field has been set.

### SetMembershipAssertionIdNil

`func (o *RwdUsNetworkMembershipsRow) SetMembershipAssertionIdNil(b bool)`

 SetMembershipAssertionIdNil sets the value for MembershipAssertionId to be an explicit nil

### UnsetMembershipAssertionId
`func (o *RwdUsNetworkMembershipsRow) UnsetMembershipAssertionId()`

UnsetMembershipAssertionId ensures that no value is present for MembershipAssertionId, not even an explicit nil
### GetNetworkSource

`func (o *RwdUsNetworkMembershipsRow) GetNetworkSource() string`

GetNetworkSource returns the NetworkSource field if non-nil, zero value otherwise.

### GetNetworkSourceOk

`func (o *RwdUsNetworkMembershipsRow) GetNetworkSourceOk() (*string, bool)`

GetNetworkSourceOk returns a tuple with the NetworkSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSource

`func (o *RwdUsNetworkMembershipsRow) SetNetworkSource(v string)`

SetNetworkSource sets NetworkSource field to given value.

### HasNetworkSource

`func (o *RwdUsNetworkMembershipsRow) HasNetworkSource() bool`

HasNetworkSource returns a boolean if a field has been set.

### SetNetworkSourceNil

`func (o *RwdUsNetworkMembershipsRow) SetNetworkSourceNil(b bool)`

 SetNetworkSourceNil sets the value for NetworkSource to be an explicit nil

### UnsetNetworkSource
`func (o *RwdUsNetworkMembershipsRow) UnsetNetworkSource()`

UnsetNetworkSource ensures that no value is present for NetworkSource, not even an explicit nil
### GetNetworkKey

`func (o *RwdUsNetworkMembershipsRow) GetNetworkKey() string`

GetNetworkKey returns the NetworkKey field if non-nil, zero value otherwise.

### GetNetworkKeyOk

`func (o *RwdUsNetworkMembershipsRow) GetNetworkKeyOk() (*string, bool)`

GetNetworkKeyOk returns a tuple with the NetworkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkKey

`func (o *RwdUsNetworkMembershipsRow) SetNetworkKey(v string)`

SetNetworkKey sets NetworkKey field to given value.

### HasNetworkKey

`func (o *RwdUsNetworkMembershipsRow) HasNetworkKey() bool`

HasNetworkKey returns a boolean if a field has been set.

### SetNetworkKeyNil

`func (o *RwdUsNetworkMembershipsRow) SetNetworkKeyNil(b bool)`

 SetNetworkKeyNil sets the value for NetworkKey to be an explicit nil

### UnsetNetworkKey
`func (o *RwdUsNetworkMembershipsRow) UnsetNetworkKey()`

UnsetNetworkKey ensures that no value is present for NetworkKey, not even an explicit nil
### GetNetworkName

`func (o *RwdUsNetworkMembershipsRow) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *RwdUsNetworkMembershipsRow) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *RwdUsNetworkMembershipsRow) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *RwdUsNetworkMembershipsRow) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### SetNetworkNameNil

`func (o *RwdUsNetworkMembershipsRow) SetNetworkNameNil(b bool)`

 SetNetworkNameNil sets the value for NetworkName to be an explicit nil

### UnsetNetworkName
`func (o *RwdUsNetworkMembershipsRow) UnsetNetworkName()`

UnsetNetworkName ensures that no value is present for NetworkName, not even an explicit nil
### GetNetworkOperator

`func (o *RwdUsNetworkMembershipsRow) GetNetworkOperator() string`

GetNetworkOperator returns the NetworkOperator field if non-nil, zero value otherwise.

### GetNetworkOperatorOk

`func (o *RwdUsNetworkMembershipsRow) GetNetworkOperatorOk() (*string, bool)`

GetNetworkOperatorOk returns a tuple with the NetworkOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOperator

`func (o *RwdUsNetworkMembershipsRow) SetNetworkOperator(v string)`

SetNetworkOperator sets NetworkOperator field to given value.

### HasNetworkOperator

`func (o *RwdUsNetworkMembershipsRow) HasNetworkOperator() bool`

HasNetworkOperator returns a boolean if a field has been set.

### SetNetworkOperatorNil

`func (o *RwdUsNetworkMembershipsRow) SetNetworkOperatorNil(b bool)`

 SetNetworkOperatorNil sets the value for NetworkOperator to be an explicit nil

### UnsetNetworkOperator
`func (o *RwdUsNetworkMembershipsRow) UnsetNetworkOperator()`

UnsetNetworkOperator ensures that no value is present for NetworkOperator, not even an explicit nil
### GetMembershipKind

`func (o *RwdUsNetworkMembershipsRow) GetMembershipKind() string`

GetMembershipKind returns the MembershipKind field if non-nil, zero value otherwise.

### GetMembershipKindOk

`func (o *RwdUsNetworkMembershipsRow) GetMembershipKindOk() (*string, bool)`

GetMembershipKindOk returns a tuple with the MembershipKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipKind

`func (o *RwdUsNetworkMembershipsRow) SetMembershipKind(v string)`

SetMembershipKind sets MembershipKind field to given value.

### HasMembershipKind

`func (o *RwdUsNetworkMembershipsRow) HasMembershipKind() bool`

HasMembershipKind returns a boolean if a field has been set.

### SetMembershipKindNil

`func (o *RwdUsNetworkMembershipsRow) SetMembershipKindNil(b bool)`

 SetMembershipKindNil sets the value for MembershipKind to be an explicit nil

### UnsetMembershipKind
`func (o *RwdUsNetworkMembershipsRow) UnsetMembershipKind()`

UnsetMembershipKind ensures that no value is present for MembershipKind, not even an explicit nil
### GetHasDta

`func (o *RwdUsNetworkMembershipsRow) GetHasDta() bool`

GetHasDta returns the HasDta field if non-nil, zero value otherwise.

### GetHasDtaOk

`func (o *RwdUsNetworkMembershipsRow) GetHasDtaOk() (*bool, bool)`

GetHasDtaOk returns a tuple with the HasDta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDta

`func (o *RwdUsNetworkMembershipsRow) SetHasDta(v bool)`

SetHasDta sets HasDta field to given value.

### HasHasDta

`func (o *RwdUsNetworkMembershipsRow) HasHasDta() bool`

HasHasDta returns a boolean if a field has been set.

### SetHasDtaNil

`func (o *RwdUsNetworkMembershipsRow) SetHasDtaNil(b bool)`

 SetHasDtaNil sets the value for HasDta to be an explicit nil

### UnsetHasDta
`func (o *RwdUsNetworkMembershipsRow) UnsetHasDta()`

UnsetHasDta ensures that no value is present for HasDta, not even an explicit nil
### GetHasDua

`func (o *RwdUsNetworkMembershipsRow) GetHasDua() bool`

GetHasDua returns the HasDua field if non-nil, zero value otherwise.

### GetHasDuaOk

`func (o *RwdUsNetworkMembershipsRow) GetHasDuaOk() (*bool, bool)`

GetHasDuaOk returns a tuple with the HasDua field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDua

`func (o *RwdUsNetworkMembershipsRow) SetHasDua(v bool)`

SetHasDua sets HasDua field to given value.

### HasHasDua

`func (o *RwdUsNetworkMembershipsRow) HasHasDua() bool`

HasHasDua returns a boolean if a field has been set.

### SetHasDuaNil

`func (o *RwdUsNetworkMembershipsRow) SetHasDuaNil(b bool)`

 SetHasDuaNil sets the value for HasDua to be an explicit nil

### UnsetHasDua
`func (o *RwdUsNetworkMembershipsRow) UnsetHasDua()`

UnsetHasDua ensures that no value is present for HasDua, not even an explicit nil
### GetHasLhba

`func (o *RwdUsNetworkMembershipsRow) GetHasLhba() bool`

GetHasLhba returns the HasLhba field if non-nil, zero value otherwise.

### GetHasLhbaOk

`func (o *RwdUsNetworkMembershipsRow) GetHasLhbaOk() (*bool, bool)`

GetHasLhbaOk returns a tuple with the HasLhba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLhba

`func (o *RwdUsNetworkMembershipsRow) SetHasLhba(v bool)`

SetHasLhba sets HasLhba field to given value.

### HasHasLhba

`func (o *RwdUsNetworkMembershipsRow) HasHasLhba() bool`

HasHasLhba returns a boolean if a field has been set.

### SetHasLhbaNil

`func (o *RwdUsNetworkMembershipsRow) SetHasLhbaNil(b bool)`

 SetHasLhbaNil sets the value for HasLhba to be an explicit nil

### UnsetHasLhba
`func (o *RwdUsNetworkMembershipsRow) UnsetHasLhba()`

UnsetHasLhba ensures that no value is present for HasLhba, not even an explicit nil
### GetHasCurrentDta

`func (o *RwdUsNetworkMembershipsRow) GetHasCurrentDta() bool`

GetHasCurrentDta returns the HasCurrentDta field if non-nil, zero value otherwise.

### GetHasCurrentDtaOk

`func (o *RwdUsNetworkMembershipsRow) GetHasCurrentDtaOk() (*bool, bool)`

GetHasCurrentDtaOk returns a tuple with the HasCurrentDta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCurrentDta

`func (o *RwdUsNetworkMembershipsRow) SetHasCurrentDta(v bool)`

SetHasCurrentDta sets HasCurrentDta field to given value.

### HasHasCurrentDta

`func (o *RwdUsNetworkMembershipsRow) HasHasCurrentDta() bool`

HasHasCurrentDta returns a boolean if a field has been set.

### SetHasCurrentDtaNil

`func (o *RwdUsNetworkMembershipsRow) SetHasCurrentDtaNil(b bool)`

 SetHasCurrentDtaNil sets the value for HasCurrentDta to be an explicit nil

### UnsetHasCurrentDta
`func (o *RwdUsNetworkMembershipsRow) UnsetHasCurrentDta()`

UnsetHasCurrentDta ensures that no value is present for HasCurrentDta, not even an explicit nil
### GetHasCurrentDua

`func (o *RwdUsNetworkMembershipsRow) GetHasCurrentDua() bool`

GetHasCurrentDua returns the HasCurrentDua field if non-nil, zero value otherwise.

### GetHasCurrentDuaOk

`func (o *RwdUsNetworkMembershipsRow) GetHasCurrentDuaOk() (*bool, bool)`

GetHasCurrentDuaOk returns a tuple with the HasCurrentDua field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCurrentDua

`func (o *RwdUsNetworkMembershipsRow) SetHasCurrentDua(v bool)`

SetHasCurrentDua sets HasCurrentDua field to given value.

### HasHasCurrentDua

`func (o *RwdUsNetworkMembershipsRow) HasHasCurrentDua() bool`

HasHasCurrentDua returns a boolean if a field has been set.

### SetHasCurrentDuaNil

`func (o *RwdUsNetworkMembershipsRow) SetHasCurrentDuaNil(b bool)`

 SetHasCurrentDuaNil sets the value for HasCurrentDua to be an explicit nil

### UnsetHasCurrentDua
`func (o *RwdUsNetworkMembershipsRow) UnsetHasCurrentDua()`

UnsetHasCurrentDua ensures that no value is present for HasCurrentDua, not even an explicit nil
### GetHasCurrentLhba

`func (o *RwdUsNetworkMembershipsRow) GetHasCurrentLhba() bool`

GetHasCurrentLhba returns the HasCurrentLhba field if non-nil, zero value otherwise.

### GetHasCurrentLhbaOk

`func (o *RwdUsNetworkMembershipsRow) GetHasCurrentLhbaOk() (*bool, bool)`

GetHasCurrentLhbaOk returns a tuple with the HasCurrentLhba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCurrentLhba

`func (o *RwdUsNetworkMembershipsRow) SetHasCurrentLhba(v bool)`

SetHasCurrentLhba sets HasCurrentLhba field to given value.

### HasHasCurrentLhba

`func (o *RwdUsNetworkMembershipsRow) HasHasCurrentLhba() bool`

HasHasCurrentLhba returns a boolean if a field has been set.

### SetHasCurrentLhbaNil

`func (o *RwdUsNetworkMembershipsRow) SetHasCurrentLhbaNil(b bool)`

 SetHasCurrentLhbaNil sets the value for HasCurrentLhba to be an explicit nil

### UnsetHasCurrentLhba
`func (o *RwdUsNetworkMembershipsRow) UnsetHasCurrentLhba()`

UnsetHasCurrentLhba ensures that no value is present for HasCurrentLhba, not even an explicit nil
### GetFirstExecutedOn

`func (o *RwdUsNetworkMembershipsRow) GetFirstExecutedOn() string`

GetFirstExecutedOn returns the FirstExecutedOn field if non-nil, zero value otherwise.

### GetFirstExecutedOnOk

`func (o *RwdUsNetworkMembershipsRow) GetFirstExecutedOnOk() (*string, bool)`

GetFirstExecutedOnOk returns a tuple with the FirstExecutedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstExecutedOn

`func (o *RwdUsNetworkMembershipsRow) SetFirstExecutedOn(v string)`

SetFirstExecutedOn sets FirstExecutedOn field to given value.

### HasFirstExecutedOn

`func (o *RwdUsNetworkMembershipsRow) HasFirstExecutedOn() bool`

HasFirstExecutedOn returns a boolean if a field has been set.

### SetFirstExecutedOnNil

`func (o *RwdUsNetworkMembershipsRow) SetFirstExecutedOnNil(b bool)`

 SetFirstExecutedOnNil sets the value for FirstExecutedOn to be an explicit nil

### UnsetFirstExecutedOn
`func (o *RwdUsNetworkMembershipsRow) UnsetFirstExecutedOn()`

UnsetFirstExecutedOn ensures that no value is present for FirstExecutedOn, not even an explicit nil
### GetLatestTermEndsOn

`func (o *RwdUsNetworkMembershipsRow) GetLatestTermEndsOn() string`

GetLatestTermEndsOn returns the LatestTermEndsOn field if non-nil, zero value otherwise.

### GetLatestTermEndsOnOk

`func (o *RwdUsNetworkMembershipsRow) GetLatestTermEndsOnOk() (*string, bool)`

GetLatestTermEndsOnOk returns a tuple with the LatestTermEndsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTermEndsOn

`func (o *RwdUsNetworkMembershipsRow) SetLatestTermEndsOn(v string)`

SetLatestTermEndsOn sets LatestTermEndsOn field to given value.

### HasLatestTermEndsOn

`func (o *RwdUsNetworkMembershipsRow) HasLatestTermEndsOn() bool`

HasLatestTermEndsOn returns a boolean if a field has been set.

### SetLatestTermEndsOnNil

`func (o *RwdUsNetworkMembershipsRow) SetLatestTermEndsOnNil(b bool)`

 SetLatestTermEndsOnNil sets the value for LatestTermEndsOn to be an explicit nil

### UnsetLatestTermEndsOn
`func (o *RwdUsNetworkMembershipsRow) UnsetLatestTermEndsOn()`

UnsetLatestTermEndsOn ensures that no value is present for LatestTermEndsOn, not even an explicit nil
### GetAgreementCount

`func (o *RwdUsNetworkMembershipsRow) GetAgreementCount() int32`

GetAgreementCount returns the AgreementCount field if non-nil, zero value otherwise.

### GetAgreementCountOk

`func (o *RwdUsNetworkMembershipsRow) GetAgreementCountOk() (*int32, bool)`

GetAgreementCountOk returns a tuple with the AgreementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementCount

`func (o *RwdUsNetworkMembershipsRow) SetAgreementCount(v int32)`

SetAgreementCount sets AgreementCount field to given value.

### HasAgreementCount

`func (o *RwdUsNetworkMembershipsRow) HasAgreementCount() bool`

HasAgreementCount returns a boolean if a field has been set.

### SetAgreementCountNil

`func (o *RwdUsNetworkMembershipsRow) SetAgreementCountNil(b bool)`

 SetAgreementCountNil sets the value for AgreementCount to be an explicit nil

### UnsetAgreementCount
`func (o *RwdUsNetworkMembershipsRow) UnsetAgreementCount()`

UnsetAgreementCount ensures that no value is present for AgreementCount, not even an explicit nil
### GetInstitutionName

`func (o *RwdUsNetworkMembershipsRow) GetInstitutionName() string`

GetInstitutionName returns the InstitutionName field if non-nil, zero value otherwise.

### GetInstitutionNameOk

`func (o *RwdUsNetworkMembershipsRow) GetInstitutionNameOk() (*string, bool)`

GetInstitutionNameOk returns a tuple with the InstitutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionName

`func (o *RwdUsNetworkMembershipsRow) SetInstitutionName(v string)`

SetInstitutionName sets InstitutionName field to given value.

### HasInstitutionName

`func (o *RwdUsNetworkMembershipsRow) HasInstitutionName() bool`

HasInstitutionName returns a boolean if a field has been set.

### SetInstitutionNameNil

`func (o *RwdUsNetworkMembershipsRow) SetInstitutionNameNil(b bool)`

 SetInstitutionNameNil sets the value for InstitutionName to be an explicit nil

### UnsetInstitutionName
`func (o *RwdUsNetworkMembershipsRow) UnsetInstitutionName()`

UnsetInstitutionName ensures that no value is present for InstitutionName, not even an explicit nil
### GetInstitutionNameKey

`func (o *RwdUsNetworkMembershipsRow) GetInstitutionNameKey() string`

GetInstitutionNameKey returns the InstitutionNameKey field if non-nil, zero value otherwise.

### GetInstitutionNameKeyOk

`func (o *RwdUsNetworkMembershipsRow) GetInstitutionNameKeyOk() (*string, bool)`

GetInstitutionNameKeyOk returns a tuple with the InstitutionNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionNameKey

`func (o *RwdUsNetworkMembershipsRow) SetInstitutionNameKey(v string)`

SetInstitutionNameKey sets InstitutionNameKey field to given value.

### HasInstitutionNameKey

`func (o *RwdUsNetworkMembershipsRow) HasInstitutionNameKey() bool`

HasInstitutionNameKey returns a boolean if a field has been set.

### SetInstitutionNameKeyNil

`func (o *RwdUsNetworkMembershipsRow) SetInstitutionNameKeyNil(b bool)`

 SetInstitutionNameKeyNil sets the value for InstitutionNameKey to be an explicit nil

### UnsetInstitutionNameKey
`func (o *RwdUsNetworkMembershipsRow) UnsetInstitutionNameKey()`

UnsetInstitutionNameKey ensures that no value is present for InstitutionNameKey, not even an explicit nil
### GetRorId

`func (o *RwdUsNetworkMembershipsRow) GetRorId() string`

GetRorId returns the RorId field if non-nil, zero value otherwise.

### GetRorIdOk

`func (o *RwdUsNetworkMembershipsRow) GetRorIdOk() (*string, bool)`

GetRorIdOk returns a tuple with the RorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRorId

`func (o *RwdUsNetworkMembershipsRow) SetRorId(v string)`

SetRorId sets RorId field to given value.

### HasRorId

`func (o *RwdUsNetworkMembershipsRow) HasRorId() bool`

HasRorId returns a boolean if a field has been set.

### SetRorIdNil

`func (o *RwdUsNetworkMembershipsRow) SetRorIdNil(b bool)`

 SetRorIdNil sets the value for RorId to be an explicit nil

### UnsetRorId
`func (o *RwdUsNetworkMembershipsRow) UnsetRorId()`

UnsetRorId ensures that no value is present for RorId, not even an explicit nil
### GetInstitutionDomain

`func (o *RwdUsNetworkMembershipsRow) GetInstitutionDomain() string`

GetInstitutionDomain returns the InstitutionDomain field if non-nil, zero value otherwise.

### GetInstitutionDomainOk

`func (o *RwdUsNetworkMembershipsRow) GetInstitutionDomainOk() (*string, bool)`

GetInstitutionDomainOk returns a tuple with the InstitutionDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionDomain

`func (o *RwdUsNetworkMembershipsRow) SetInstitutionDomain(v string)`

SetInstitutionDomain sets InstitutionDomain field to given value.

### HasInstitutionDomain

`func (o *RwdUsNetworkMembershipsRow) HasInstitutionDomain() bool`

HasInstitutionDomain returns a boolean if a field has been set.

### SetInstitutionDomainNil

`func (o *RwdUsNetworkMembershipsRow) SetInstitutionDomainNil(b bool)`

 SetInstitutionDomainNil sets the value for InstitutionDomain to be an explicit nil

### UnsetInstitutionDomain
`func (o *RwdUsNetworkMembershipsRow) UnsetInstitutionDomain()`

UnsetInstitutionDomain ensures that no value is present for InstitutionDomain, not even an explicit nil
### GetIsLead

`func (o *RwdUsNetworkMembershipsRow) GetIsLead() bool`

GetIsLead returns the IsLead field if non-nil, zero value otherwise.

### GetIsLeadOk

`func (o *RwdUsNetworkMembershipsRow) GetIsLeadOk() (*bool, bool)`

GetIsLeadOk returns a tuple with the IsLead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLead

`func (o *RwdUsNetworkMembershipsRow) SetIsLead(v bool)`

SetIsLead sets IsLead field to given value.

### HasIsLead

`func (o *RwdUsNetworkMembershipsRow) HasIsLead() bool`

HasIsLead returns a boolean if a field has been set.

### SetIsLeadNil

`func (o *RwdUsNetworkMembershipsRow) SetIsLeadNil(b bool)`

 SetIsLeadNil sets the value for IsLead to be an explicit nil

### UnsetIsLead
`func (o *RwdUsNetworkMembershipsRow) UnsetIsLead()`

UnsetIsLead ensures that no value is present for IsLead, not even an explicit nil
### GetAsOf

`func (o *RwdUsNetworkMembershipsRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *RwdUsNetworkMembershipsRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *RwdUsNetworkMembershipsRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *RwdUsNetworkMembershipsRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *RwdUsNetworkMembershipsRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *RwdUsNetworkMembershipsRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetSourceUrl

`func (o *RwdUsNetworkMembershipsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *RwdUsNetworkMembershipsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *RwdUsNetworkMembershipsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *RwdUsNetworkMembershipsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *RwdUsNetworkMembershipsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *RwdUsNetworkMembershipsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetRetrievedAt

`func (o *RwdUsNetworkMembershipsRow) GetRetrievedAt() time.Time`

GetRetrievedAt returns the RetrievedAt field if non-nil, zero value otherwise.

### GetRetrievedAtOk

`func (o *RwdUsNetworkMembershipsRow) GetRetrievedAtOk() (*time.Time, bool)`

GetRetrievedAtOk returns a tuple with the RetrievedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievedAt

`func (o *RwdUsNetworkMembershipsRow) SetRetrievedAt(v time.Time)`

SetRetrievedAt sets RetrievedAt field to given value.

### HasRetrievedAt

`func (o *RwdUsNetworkMembershipsRow) HasRetrievedAt() bool`

HasRetrievedAt returns a boolean if a field has been set.

### SetRetrievedAtNil

`func (o *RwdUsNetworkMembershipsRow) SetRetrievedAtNil(b bool)`

 SetRetrievedAtNil sets the value for RetrievedAt to be an explicit nil

### UnsetRetrievedAt
`func (o *RwdUsNetworkMembershipsRow) UnsetRetrievedAt()`

UnsetRetrievedAt ensures that no value is present for RetrievedAt, not even an explicit nil
### GetOrgId

`func (o *RwdUsNetworkMembershipsRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *RwdUsNetworkMembershipsRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *RwdUsNetworkMembershipsRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *RwdUsNetworkMembershipsRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *RwdUsNetworkMembershipsRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *RwdUsNetworkMembershipsRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetOrgName

`func (o *RwdUsNetworkMembershipsRow) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *RwdUsNetworkMembershipsRow) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *RwdUsNetworkMembershipsRow) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *RwdUsNetworkMembershipsRow) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### SetOrgNameNil

`func (o *RwdUsNetworkMembershipsRow) SetOrgNameNil(b bool)`

 SetOrgNameNil sets the value for OrgName to be an explicit nil

### UnsetOrgName
`func (o *RwdUsNetworkMembershipsRow) UnsetOrgName()`

UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil
### GetOrgMatchRung

`func (o *RwdUsNetworkMembershipsRow) GetOrgMatchRung() string`

GetOrgMatchRung returns the OrgMatchRung field if non-nil, zero value otherwise.

### GetOrgMatchRungOk

`func (o *RwdUsNetworkMembershipsRow) GetOrgMatchRungOk() (*string, bool)`

GetOrgMatchRungOk returns a tuple with the OrgMatchRung field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMatchRung

`func (o *RwdUsNetworkMembershipsRow) SetOrgMatchRung(v string)`

SetOrgMatchRung sets OrgMatchRung field to given value.

### HasOrgMatchRung

`func (o *RwdUsNetworkMembershipsRow) HasOrgMatchRung() bool`

HasOrgMatchRung returns a boolean if a field has been set.

### SetOrgMatchRungNil

`func (o *RwdUsNetworkMembershipsRow) SetOrgMatchRungNil(b bool)`

 SetOrgMatchRungNil sets the value for OrgMatchRung to be an explicit nil

### UnsetOrgMatchRung
`func (o *RwdUsNetworkMembershipsRow) UnsetOrgMatchRung()`

UnsetOrgMatchRung ensures that no value is present for OrgMatchRung, not even an explicit nil
### GetOrgMatchTier

`func (o *RwdUsNetworkMembershipsRow) GetOrgMatchTier() string`

GetOrgMatchTier returns the OrgMatchTier field if non-nil, zero value otherwise.

### GetOrgMatchTierOk

`func (o *RwdUsNetworkMembershipsRow) GetOrgMatchTierOk() (*string, bool)`

GetOrgMatchTierOk returns a tuple with the OrgMatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMatchTier

`func (o *RwdUsNetworkMembershipsRow) SetOrgMatchTier(v string)`

SetOrgMatchTier sets OrgMatchTier field to given value.

### HasOrgMatchTier

`func (o *RwdUsNetworkMembershipsRow) HasOrgMatchTier() bool`

HasOrgMatchTier returns a boolean if a field has been set.

### SetOrgMatchTierNil

`func (o *RwdUsNetworkMembershipsRow) SetOrgMatchTierNil(b bool)`

 SetOrgMatchTierNil sets the value for OrgMatchTier to be an explicit nil

### UnsetOrgMatchTier
`func (o *RwdUsNetworkMembershipsRow) UnsetOrgMatchTier()`

UnsetOrgMatchTier ensures that no value is present for OrgMatchTier, not even an explicit nil
### GetOrgMatchConfidence

`func (o *RwdUsNetworkMembershipsRow) GetOrgMatchConfidence() float32`

GetOrgMatchConfidence returns the OrgMatchConfidence field if non-nil, zero value otherwise.

### GetOrgMatchConfidenceOk

`func (o *RwdUsNetworkMembershipsRow) GetOrgMatchConfidenceOk() (*float32, bool)`

GetOrgMatchConfidenceOk returns a tuple with the OrgMatchConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMatchConfidence

`func (o *RwdUsNetworkMembershipsRow) SetOrgMatchConfidence(v float32)`

SetOrgMatchConfidence sets OrgMatchConfidence field to given value.

### HasOrgMatchConfidence

`func (o *RwdUsNetworkMembershipsRow) HasOrgMatchConfidence() bool`

HasOrgMatchConfidence returns a boolean if a field has been set.

### SetOrgMatchConfidenceNil

`func (o *RwdUsNetworkMembershipsRow) SetOrgMatchConfidenceNil(b bool)`

 SetOrgMatchConfidenceNil sets the value for OrgMatchConfidence to be an explicit nil

### UnsetOrgMatchConfidence
`func (o *RwdUsNetworkMembershipsRow) UnsetOrgMatchConfidence()`

UnsetOrgMatchConfidence ensures that no value is present for OrgMatchConfidence, not even an explicit nil
### GetProviderId

`func (o *RwdUsNetworkMembershipsRow) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *RwdUsNetworkMembershipsRow) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *RwdUsNetworkMembershipsRow) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *RwdUsNetworkMembershipsRow) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *RwdUsNetworkMembershipsRow) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *RwdUsNetworkMembershipsRow) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetProviderUuid

`func (o *RwdUsNetworkMembershipsRow) GetProviderUuid() string`

GetProviderUuid returns the ProviderUuid field if non-nil, zero value otherwise.

### GetProviderUuidOk

`func (o *RwdUsNetworkMembershipsRow) GetProviderUuidOk() (*string, bool)`

GetProviderUuidOk returns a tuple with the ProviderUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUuid

`func (o *RwdUsNetworkMembershipsRow) SetProviderUuid(v string)`

SetProviderUuid sets ProviderUuid field to given value.

### HasProviderUuid

`func (o *RwdUsNetworkMembershipsRow) HasProviderUuid() bool`

HasProviderUuid returns a boolean if a field has been set.

### SetProviderUuidNil

`func (o *RwdUsNetworkMembershipsRow) SetProviderUuidNil(b bool)`

 SetProviderUuidNil sets the value for ProviderUuid to be an explicit nil

### UnsetProviderUuid
`func (o *RwdUsNetworkMembershipsRow) UnsetProviderUuid()`

UnsetProviderUuid ensures that no value is present for ProviderUuid, not even an explicit nil
### GetProviderSlug

`func (o *RwdUsNetworkMembershipsRow) GetProviderSlug() string`

GetProviderSlug returns the ProviderSlug field if non-nil, zero value otherwise.

### GetProviderSlugOk

`func (o *RwdUsNetworkMembershipsRow) GetProviderSlugOk() (*string, bool)`

GetProviderSlugOk returns a tuple with the ProviderSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSlug

`func (o *RwdUsNetworkMembershipsRow) SetProviderSlug(v string)`

SetProviderSlug sets ProviderSlug field to given value.

### HasProviderSlug

`func (o *RwdUsNetworkMembershipsRow) HasProviderSlug() bool`

HasProviderSlug returns a boolean if a field has been set.

### SetProviderSlugNil

`func (o *RwdUsNetworkMembershipsRow) SetProviderSlugNil(b bool)`

 SetProviderSlugNil sets the value for ProviderSlug to be an explicit nil

### UnsetProviderSlug
`func (o *RwdUsNetworkMembershipsRow) UnsetProviderSlug()`

UnsetProviderSlug ensures that no value is present for ProviderSlug, not even an explicit nil
### GetProviderName

`func (o *RwdUsNetworkMembershipsRow) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *RwdUsNetworkMembershipsRow) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *RwdUsNetworkMembershipsRow) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *RwdUsNetworkMembershipsRow) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *RwdUsNetworkMembershipsRow) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *RwdUsNetworkMembershipsRow) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetMatchMethod

`func (o *RwdUsNetworkMembershipsRow) GetMatchMethod() string`

GetMatchMethod returns the MatchMethod field if non-nil, zero value otherwise.

### GetMatchMethodOk

`func (o *RwdUsNetworkMembershipsRow) GetMatchMethodOk() (*string, bool)`

GetMatchMethodOk returns a tuple with the MatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMethod

`func (o *RwdUsNetworkMembershipsRow) SetMatchMethod(v string)`

SetMatchMethod sets MatchMethod field to given value.

### HasMatchMethod

`func (o *RwdUsNetworkMembershipsRow) HasMatchMethod() bool`

HasMatchMethod returns a boolean if a field has been set.

### SetMatchMethodNil

`func (o *RwdUsNetworkMembershipsRow) SetMatchMethodNil(b bool)`

 SetMatchMethodNil sets the value for MatchMethod to be an explicit nil

### UnsetMatchMethod
`func (o *RwdUsNetworkMembershipsRow) UnsetMatchMethod()`

UnsetMatchMethod ensures that no value is present for MatchMethod, not even an explicit nil
### GetAlsoMatchedByName

`func (o *RwdUsNetworkMembershipsRow) GetAlsoMatchedByName() bool`

GetAlsoMatchedByName returns the AlsoMatchedByName field if non-nil, zero value otherwise.

### GetAlsoMatchedByNameOk

`func (o *RwdUsNetworkMembershipsRow) GetAlsoMatchedByNameOk() (*bool, bool)`

GetAlsoMatchedByNameOk returns a tuple with the AlsoMatchedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlsoMatchedByName

`func (o *RwdUsNetworkMembershipsRow) SetAlsoMatchedByName(v bool)`

SetAlsoMatchedByName sets AlsoMatchedByName field to given value.

### HasAlsoMatchedByName

`func (o *RwdUsNetworkMembershipsRow) HasAlsoMatchedByName() bool`

HasAlsoMatchedByName returns a boolean if a field has been set.

### SetAlsoMatchedByNameNil

`func (o *RwdUsNetworkMembershipsRow) SetAlsoMatchedByNameNil(b bool)`

 SetAlsoMatchedByNameNil sets the value for AlsoMatchedByName to be an explicit nil

### UnsetAlsoMatchedByName
`func (o *RwdUsNetworkMembershipsRow) UnsetAlsoMatchedByName()`

UnsetAlsoMatchedByName ensures that no value is present for AlsoMatchedByName, not even an explicit nil
### GetMatchConfidence

`func (o *RwdUsNetworkMembershipsRow) GetMatchConfidence() float32`

GetMatchConfidence returns the MatchConfidence field if non-nil, zero value otherwise.

### GetMatchConfidenceOk

`func (o *RwdUsNetworkMembershipsRow) GetMatchConfidenceOk() (*float32, bool)`

GetMatchConfidenceOk returns a tuple with the MatchConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchConfidence

`func (o *RwdUsNetworkMembershipsRow) SetMatchConfidence(v float32)`

SetMatchConfidence sets MatchConfidence field to given value.

### HasMatchConfidence

`func (o *RwdUsNetworkMembershipsRow) HasMatchConfidence() bool`

HasMatchConfidence returns a boolean if a field has been set.

### SetMatchConfidenceNil

`func (o *RwdUsNetworkMembershipsRow) SetMatchConfidenceNil(b bool)`

 SetMatchConfidenceNil sets the value for MatchConfidence to be an explicit nil

### UnsetMatchConfidence
`func (o *RwdUsNetworkMembershipsRow) UnsetMatchConfidence()`

UnsetMatchConfidence ensures that no value is present for MatchConfidence, not even an explicit nil
### GetMemberTotal

`func (o *RwdUsNetworkMembershipsRow) GetMemberTotal() int32`

GetMemberTotal returns the MemberTotal field if non-nil, zero value otherwise.

### GetMemberTotalOk

`func (o *RwdUsNetworkMembershipsRow) GetMemberTotalOk() (*int32, bool)`

GetMemberTotalOk returns a tuple with the MemberTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberTotal

`func (o *RwdUsNetworkMembershipsRow) SetMemberTotal(v int32)`

SetMemberTotal sets MemberTotal field to given value.

### HasMemberTotal

`func (o *RwdUsNetworkMembershipsRow) HasMemberTotal() bool`

HasMemberTotal returns a boolean if a field has been set.

### SetMemberTotalNil

`func (o *RwdUsNetworkMembershipsRow) SetMemberTotalNil(b bool)`

 SetMemberTotalNil sets the value for MemberTotal to be an explicit nil

### UnsetMemberTotal
`func (o *RwdUsNetworkMembershipsRow) UnsetMemberTotal()`

UnsetMemberTotal ensures that no value is present for MemberTotal, not even an explicit nil
### GetMembersWithCurrentDta

`func (o *RwdUsNetworkMembershipsRow) GetMembersWithCurrentDta() int32`

GetMembersWithCurrentDta returns the MembersWithCurrentDta field if non-nil, zero value otherwise.

### GetMembersWithCurrentDtaOk

`func (o *RwdUsNetworkMembershipsRow) GetMembersWithCurrentDtaOk() (*int32, bool)`

GetMembersWithCurrentDtaOk returns a tuple with the MembersWithCurrentDta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersWithCurrentDta

`func (o *RwdUsNetworkMembershipsRow) SetMembersWithCurrentDta(v int32)`

SetMembersWithCurrentDta sets MembersWithCurrentDta field to given value.

### HasMembersWithCurrentDta

`func (o *RwdUsNetworkMembershipsRow) HasMembersWithCurrentDta() bool`

HasMembersWithCurrentDta returns a boolean if a field has been set.

### SetMembersWithCurrentDtaNil

`func (o *RwdUsNetworkMembershipsRow) SetMembersWithCurrentDtaNil(b bool)`

 SetMembersWithCurrentDtaNil sets the value for MembersWithCurrentDta to be an explicit nil

### UnsetMembersWithCurrentDta
`func (o *RwdUsNetworkMembershipsRow) UnsetMembersWithCurrentDta()`

UnsetMembersWithCurrentDta ensures that no value is present for MembersWithCurrentDta, not even an explicit nil
### GetMembersWithCurrentDua

`func (o *RwdUsNetworkMembershipsRow) GetMembersWithCurrentDua() int32`

GetMembersWithCurrentDua returns the MembersWithCurrentDua field if non-nil, zero value otherwise.

### GetMembersWithCurrentDuaOk

`func (o *RwdUsNetworkMembershipsRow) GetMembersWithCurrentDuaOk() (*int32, bool)`

GetMembersWithCurrentDuaOk returns a tuple with the MembersWithCurrentDua field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersWithCurrentDua

`func (o *RwdUsNetworkMembershipsRow) SetMembersWithCurrentDua(v int32)`

SetMembersWithCurrentDua sets MembersWithCurrentDua field to given value.

### HasMembersWithCurrentDua

`func (o *RwdUsNetworkMembershipsRow) HasMembersWithCurrentDua() bool`

HasMembersWithCurrentDua returns a boolean if a field has been set.

### SetMembersWithCurrentDuaNil

`func (o *RwdUsNetworkMembershipsRow) SetMembersWithCurrentDuaNil(b bool)`

 SetMembersWithCurrentDuaNil sets the value for MembersWithCurrentDua to be an explicit nil

### UnsetMembersWithCurrentDua
`func (o *RwdUsNetworkMembershipsRow) UnsetMembersWithCurrentDua()`

UnsetMembersWithCurrentDua ensures that no value is present for MembersWithCurrentDua, not even an explicit nil
### GetMembersResolvedOrg

`func (o *RwdUsNetworkMembershipsRow) GetMembersResolvedOrg() int32`

GetMembersResolvedOrg returns the MembersResolvedOrg field if non-nil, zero value otherwise.

### GetMembersResolvedOrgOk

`func (o *RwdUsNetworkMembershipsRow) GetMembersResolvedOrgOk() (*int32, bool)`

GetMembersResolvedOrgOk returns a tuple with the MembersResolvedOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersResolvedOrg

`func (o *RwdUsNetworkMembershipsRow) SetMembersResolvedOrg(v int32)`

SetMembersResolvedOrg sets MembersResolvedOrg field to given value.

### HasMembersResolvedOrg

`func (o *RwdUsNetworkMembershipsRow) HasMembersResolvedOrg() bool`

HasMembersResolvedOrg returns a boolean if a field has been set.

### SetMembersResolvedOrgNil

`func (o *RwdUsNetworkMembershipsRow) SetMembersResolvedOrgNil(b bool)`

 SetMembersResolvedOrgNil sets the value for MembersResolvedOrg to be an explicit nil

### UnsetMembersResolvedOrg
`func (o *RwdUsNetworkMembershipsRow) UnsetMembersResolvedOrg()`

UnsetMembersResolvedOrg ensures that no value is present for MembersResolvedOrg, not even an explicit nil
### GetMembersMatchedRx

`func (o *RwdUsNetworkMembershipsRow) GetMembersMatchedRx() int32`

GetMembersMatchedRx returns the MembersMatchedRx field if non-nil, zero value otherwise.

### GetMembersMatchedRxOk

`func (o *RwdUsNetworkMembershipsRow) GetMembersMatchedRxOk() (*int32, bool)`

GetMembersMatchedRxOk returns a tuple with the MembersMatchedRx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersMatchedRx

`func (o *RwdUsNetworkMembershipsRow) SetMembersMatchedRx(v int32)`

SetMembersMatchedRx sets MembersMatchedRx field to given value.

### HasMembersMatchedRx

`func (o *RwdUsNetworkMembershipsRow) HasMembersMatchedRx() bool`

HasMembersMatchedRx returns a boolean if a field has been set.

### SetMembersMatchedRxNil

`func (o *RwdUsNetworkMembershipsRow) SetMembersMatchedRxNil(b bool)`

 SetMembersMatchedRxNil sets the value for MembersMatchedRx to be an explicit nil

### UnsetMembersMatchedRx
`func (o *RwdUsNetworkMembershipsRow) UnsetMembersMatchedRx()`

UnsetMembersMatchedRx ensures that no value is present for MembersMatchedRx, not even an explicit nil
### GetAttribution

`func (o *RwdUsNetworkMembershipsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdUsNetworkMembershipsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdUsNetworkMembershipsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdUsNetworkMembershipsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdUsNetworkMembershipsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdUsNetworkMembershipsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


