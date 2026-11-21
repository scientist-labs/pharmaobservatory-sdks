# RwdUsNetworksRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkKey** | Pointer to **NullableString** | The network key, the entry key of the rwd_us_networks object type (n3c today). | [optional] 
**NetworkSource** | Pointer to **NullableString** | The source the network rows come from. | [optional] 
**NetworkName** | Pointer to **NullableString** | The network&#39;s name. | [optional] 
**NetworkOperator** | Pointer to **NullableString** | The operating body. | [optional] 
**MemberTotal** | Pointer to **NullableInt32** | Institutions holding at least one published agreement, our count of the register. | [optional] 
**MembersWithDta** | Pointer to **NullableInt32** | Institutions holding a Data Transfer Agreement. | [optional] 
**MembersWithCurrentDta** | Pointer to **NullableInt32** | Institutions with a DTA row whose published term end is on or after the retrieval date. | [optional] 
**MembersWithDua** | Pointer to **NullableInt32** | Institutions holding a Data Use Agreement. | [optional] 
**MembersWithCurrentDua** | Pointer to **NullableInt32** | Institutions with a DUA row within its published term. | [optional] 
**MembersWithLhba** | Pointer to **NullableInt32** | Institutions holding an agreement published with the type code LHBA. | [optional] 
**MembersResolvedOrg** | Pointer to **NullableInt32** | Members resolved to an organization on the spine. | [optional] 
**MembersMatchedRx** | Pointer to **NullableInt32** | Members matched to at least one RX provider, by an exact method visible in rwd_us_network_memberships. | [optional] 
**PublishedDtaCount** | Pointer to **NullableInt32** | N3C&#39;s own published count of DTAs, as N3C labels it. Differs from ours; not reconciled. | [optional] 
**PublishedDuaCount** | Pointer to **NullableInt32** | N3C&#39;s own published count of DUAs. Differs from ours; not reconciled. | [optional] 
**PublishedInstitutionCount** | Pointer to **NullableInt32** | N3C&#39;s own published count of institutions. Differs from ours; not reconciled. | [optional] 
**PublishedProjectCount** | Pointer to **NullableInt32** | N3C&#39;s own published count of projects. Differs from ours; not reconciled. | [optional] 
**AsOf** | Pointer to **NullableString** | The retrieval date of the agreement snapshot. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The API endpoint the agreements came from. | [optional] 
**AgreementsUrl** | Pointer to **NullableString** | The public N3C dashboard. | [optional] 
**SignatoriesPageUrl** | Pointer to **NullableString** | N3C&#39;s HTML signatories page, kept as a documented reconciliation and not ingested. | [optional] 
**RetrievedAt** | Pointer to **NullableTime** | When the snapshot was retrieved, in UTC. | [optional] 
**Attribution** | Pointer to **NullableString** | The credit line every surface carrying this data prints. | [optional] 

## Methods

### NewRwdUsNetworksRow

`func NewRwdUsNetworksRow() *RwdUsNetworksRow`

NewRwdUsNetworksRow instantiates a new RwdUsNetworksRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdUsNetworksRowWithDefaults

`func NewRwdUsNetworksRowWithDefaults() *RwdUsNetworksRow`

NewRwdUsNetworksRowWithDefaults instantiates a new RwdUsNetworksRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkKey

`func (o *RwdUsNetworksRow) GetNetworkKey() string`

GetNetworkKey returns the NetworkKey field if non-nil, zero value otherwise.

### GetNetworkKeyOk

`func (o *RwdUsNetworksRow) GetNetworkKeyOk() (*string, bool)`

GetNetworkKeyOk returns a tuple with the NetworkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkKey

`func (o *RwdUsNetworksRow) SetNetworkKey(v string)`

SetNetworkKey sets NetworkKey field to given value.

### HasNetworkKey

`func (o *RwdUsNetworksRow) HasNetworkKey() bool`

HasNetworkKey returns a boolean if a field has been set.

### SetNetworkKeyNil

`func (o *RwdUsNetworksRow) SetNetworkKeyNil(b bool)`

 SetNetworkKeyNil sets the value for NetworkKey to be an explicit nil

### UnsetNetworkKey
`func (o *RwdUsNetworksRow) UnsetNetworkKey()`

UnsetNetworkKey ensures that no value is present for NetworkKey, not even an explicit nil
### GetNetworkSource

`func (o *RwdUsNetworksRow) GetNetworkSource() string`

GetNetworkSource returns the NetworkSource field if non-nil, zero value otherwise.

### GetNetworkSourceOk

`func (o *RwdUsNetworksRow) GetNetworkSourceOk() (*string, bool)`

GetNetworkSourceOk returns a tuple with the NetworkSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSource

`func (o *RwdUsNetworksRow) SetNetworkSource(v string)`

SetNetworkSource sets NetworkSource field to given value.

### HasNetworkSource

`func (o *RwdUsNetworksRow) HasNetworkSource() bool`

HasNetworkSource returns a boolean if a field has been set.

### SetNetworkSourceNil

`func (o *RwdUsNetworksRow) SetNetworkSourceNil(b bool)`

 SetNetworkSourceNil sets the value for NetworkSource to be an explicit nil

### UnsetNetworkSource
`func (o *RwdUsNetworksRow) UnsetNetworkSource()`

UnsetNetworkSource ensures that no value is present for NetworkSource, not even an explicit nil
### GetNetworkName

`func (o *RwdUsNetworksRow) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *RwdUsNetworksRow) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *RwdUsNetworksRow) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *RwdUsNetworksRow) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### SetNetworkNameNil

`func (o *RwdUsNetworksRow) SetNetworkNameNil(b bool)`

 SetNetworkNameNil sets the value for NetworkName to be an explicit nil

### UnsetNetworkName
`func (o *RwdUsNetworksRow) UnsetNetworkName()`

UnsetNetworkName ensures that no value is present for NetworkName, not even an explicit nil
### GetNetworkOperator

`func (o *RwdUsNetworksRow) GetNetworkOperator() string`

GetNetworkOperator returns the NetworkOperator field if non-nil, zero value otherwise.

### GetNetworkOperatorOk

`func (o *RwdUsNetworksRow) GetNetworkOperatorOk() (*string, bool)`

GetNetworkOperatorOk returns a tuple with the NetworkOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOperator

`func (o *RwdUsNetworksRow) SetNetworkOperator(v string)`

SetNetworkOperator sets NetworkOperator field to given value.

### HasNetworkOperator

`func (o *RwdUsNetworksRow) HasNetworkOperator() bool`

HasNetworkOperator returns a boolean if a field has been set.

### SetNetworkOperatorNil

`func (o *RwdUsNetworksRow) SetNetworkOperatorNil(b bool)`

 SetNetworkOperatorNil sets the value for NetworkOperator to be an explicit nil

### UnsetNetworkOperator
`func (o *RwdUsNetworksRow) UnsetNetworkOperator()`

UnsetNetworkOperator ensures that no value is present for NetworkOperator, not even an explicit nil
### GetMemberTotal

`func (o *RwdUsNetworksRow) GetMemberTotal() int32`

GetMemberTotal returns the MemberTotal field if non-nil, zero value otherwise.

### GetMemberTotalOk

`func (o *RwdUsNetworksRow) GetMemberTotalOk() (*int32, bool)`

GetMemberTotalOk returns a tuple with the MemberTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberTotal

`func (o *RwdUsNetworksRow) SetMemberTotal(v int32)`

SetMemberTotal sets MemberTotal field to given value.

### HasMemberTotal

`func (o *RwdUsNetworksRow) HasMemberTotal() bool`

HasMemberTotal returns a boolean if a field has been set.

### SetMemberTotalNil

`func (o *RwdUsNetworksRow) SetMemberTotalNil(b bool)`

 SetMemberTotalNil sets the value for MemberTotal to be an explicit nil

### UnsetMemberTotal
`func (o *RwdUsNetworksRow) UnsetMemberTotal()`

UnsetMemberTotal ensures that no value is present for MemberTotal, not even an explicit nil
### GetMembersWithDta

`func (o *RwdUsNetworksRow) GetMembersWithDta() int32`

GetMembersWithDta returns the MembersWithDta field if non-nil, zero value otherwise.

### GetMembersWithDtaOk

`func (o *RwdUsNetworksRow) GetMembersWithDtaOk() (*int32, bool)`

GetMembersWithDtaOk returns a tuple with the MembersWithDta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersWithDta

`func (o *RwdUsNetworksRow) SetMembersWithDta(v int32)`

SetMembersWithDta sets MembersWithDta field to given value.

### HasMembersWithDta

`func (o *RwdUsNetworksRow) HasMembersWithDta() bool`

HasMembersWithDta returns a boolean if a field has been set.

### SetMembersWithDtaNil

`func (o *RwdUsNetworksRow) SetMembersWithDtaNil(b bool)`

 SetMembersWithDtaNil sets the value for MembersWithDta to be an explicit nil

### UnsetMembersWithDta
`func (o *RwdUsNetworksRow) UnsetMembersWithDta()`

UnsetMembersWithDta ensures that no value is present for MembersWithDta, not even an explicit nil
### GetMembersWithCurrentDta

`func (o *RwdUsNetworksRow) GetMembersWithCurrentDta() int32`

GetMembersWithCurrentDta returns the MembersWithCurrentDta field if non-nil, zero value otherwise.

### GetMembersWithCurrentDtaOk

`func (o *RwdUsNetworksRow) GetMembersWithCurrentDtaOk() (*int32, bool)`

GetMembersWithCurrentDtaOk returns a tuple with the MembersWithCurrentDta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersWithCurrentDta

`func (o *RwdUsNetworksRow) SetMembersWithCurrentDta(v int32)`

SetMembersWithCurrentDta sets MembersWithCurrentDta field to given value.

### HasMembersWithCurrentDta

`func (o *RwdUsNetworksRow) HasMembersWithCurrentDta() bool`

HasMembersWithCurrentDta returns a boolean if a field has been set.

### SetMembersWithCurrentDtaNil

`func (o *RwdUsNetworksRow) SetMembersWithCurrentDtaNil(b bool)`

 SetMembersWithCurrentDtaNil sets the value for MembersWithCurrentDta to be an explicit nil

### UnsetMembersWithCurrentDta
`func (o *RwdUsNetworksRow) UnsetMembersWithCurrentDta()`

UnsetMembersWithCurrentDta ensures that no value is present for MembersWithCurrentDta, not even an explicit nil
### GetMembersWithDua

`func (o *RwdUsNetworksRow) GetMembersWithDua() int32`

GetMembersWithDua returns the MembersWithDua field if non-nil, zero value otherwise.

### GetMembersWithDuaOk

`func (o *RwdUsNetworksRow) GetMembersWithDuaOk() (*int32, bool)`

GetMembersWithDuaOk returns a tuple with the MembersWithDua field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersWithDua

`func (o *RwdUsNetworksRow) SetMembersWithDua(v int32)`

SetMembersWithDua sets MembersWithDua field to given value.

### HasMembersWithDua

`func (o *RwdUsNetworksRow) HasMembersWithDua() bool`

HasMembersWithDua returns a boolean if a field has been set.

### SetMembersWithDuaNil

`func (o *RwdUsNetworksRow) SetMembersWithDuaNil(b bool)`

 SetMembersWithDuaNil sets the value for MembersWithDua to be an explicit nil

### UnsetMembersWithDua
`func (o *RwdUsNetworksRow) UnsetMembersWithDua()`

UnsetMembersWithDua ensures that no value is present for MembersWithDua, not even an explicit nil
### GetMembersWithCurrentDua

`func (o *RwdUsNetworksRow) GetMembersWithCurrentDua() int32`

GetMembersWithCurrentDua returns the MembersWithCurrentDua field if non-nil, zero value otherwise.

### GetMembersWithCurrentDuaOk

`func (o *RwdUsNetworksRow) GetMembersWithCurrentDuaOk() (*int32, bool)`

GetMembersWithCurrentDuaOk returns a tuple with the MembersWithCurrentDua field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersWithCurrentDua

`func (o *RwdUsNetworksRow) SetMembersWithCurrentDua(v int32)`

SetMembersWithCurrentDua sets MembersWithCurrentDua field to given value.

### HasMembersWithCurrentDua

`func (o *RwdUsNetworksRow) HasMembersWithCurrentDua() bool`

HasMembersWithCurrentDua returns a boolean if a field has been set.

### SetMembersWithCurrentDuaNil

`func (o *RwdUsNetworksRow) SetMembersWithCurrentDuaNil(b bool)`

 SetMembersWithCurrentDuaNil sets the value for MembersWithCurrentDua to be an explicit nil

### UnsetMembersWithCurrentDua
`func (o *RwdUsNetworksRow) UnsetMembersWithCurrentDua()`

UnsetMembersWithCurrentDua ensures that no value is present for MembersWithCurrentDua, not even an explicit nil
### GetMembersWithLhba

`func (o *RwdUsNetworksRow) GetMembersWithLhba() int32`

GetMembersWithLhba returns the MembersWithLhba field if non-nil, zero value otherwise.

### GetMembersWithLhbaOk

`func (o *RwdUsNetworksRow) GetMembersWithLhbaOk() (*int32, bool)`

GetMembersWithLhbaOk returns a tuple with the MembersWithLhba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersWithLhba

`func (o *RwdUsNetworksRow) SetMembersWithLhba(v int32)`

SetMembersWithLhba sets MembersWithLhba field to given value.

### HasMembersWithLhba

`func (o *RwdUsNetworksRow) HasMembersWithLhba() bool`

HasMembersWithLhba returns a boolean if a field has been set.

### SetMembersWithLhbaNil

`func (o *RwdUsNetworksRow) SetMembersWithLhbaNil(b bool)`

 SetMembersWithLhbaNil sets the value for MembersWithLhba to be an explicit nil

### UnsetMembersWithLhba
`func (o *RwdUsNetworksRow) UnsetMembersWithLhba()`

UnsetMembersWithLhba ensures that no value is present for MembersWithLhba, not even an explicit nil
### GetMembersResolvedOrg

`func (o *RwdUsNetworksRow) GetMembersResolvedOrg() int32`

GetMembersResolvedOrg returns the MembersResolvedOrg field if non-nil, zero value otherwise.

### GetMembersResolvedOrgOk

`func (o *RwdUsNetworksRow) GetMembersResolvedOrgOk() (*int32, bool)`

GetMembersResolvedOrgOk returns a tuple with the MembersResolvedOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersResolvedOrg

`func (o *RwdUsNetworksRow) SetMembersResolvedOrg(v int32)`

SetMembersResolvedOrg sets MembersResolvedOrg field to given value.

### HasMembersResolvedOrg

`func (o *RwdUsNetworksRow) HasMembersResolvedOrg() bool`

HasMembersResolvedOrg returns a boolean if a field has been set.

### SetMembersResolvedOrgNil

`func (o *RwdUsNetworksRow) SetMembersResolvedOrgNil(b bool)`

 SetMembersResolvedOrgNil sets the value for MembersResolvedOrg to be an explicit nil

### UnsetMembersResolvedOrg
`func (o *RwdUsNetworksRow) UnsetMembersResolvedOrg()`

UnsetMembersResolvedOrg ensures that no value is present for MembersResolvedOrg, not even an explicit nil
### GetMembersMatchedRx

`func (o *RwdUsNetworksRow) GetMembersMatchedRx() int32`

GetMembersMatchedRx returns the MembersMatchedRx field if non-nil, zero value otherwise.

### GetMembersMatchedRxOk

`func (o *RwdUsNetworksRow) GetMembersMatchedRxOk() (*int32, bool)`

GetMembersMatchedRxOk returns a tuple with the MembersMatchedRx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersMatchedRx

`func (o *RwdUsNetworksRow) SetMembersMatchedRx(v int32)`

SetMembersMatchedRx sets MembersMatchedRx field to given value.

### HasMembersMatchedRx

`func (o *RwdUsNetworksRow) HasMembersMatchedRx() bool`

HasMembersMatchedRx returns a boolean if a field has been set.

### SetMembersMatchedRxNil

`func (o *RwdUsNetworksRow) SetMembersMatchedRxNil(b bool)`

 SetMembersMatchedRxNil sets the value for MembersMatchedRx to be an explicit nil

### UnsetMembersMatchedRx
`func (o *RwdUsNetworksRow) UnsetMembersMatchedRx()`

UnsetMembersMatchedRx ensures that no value is present for MembersMatchedRx, not even an explicit nil
### GetPublishedDtaCount

`func (o *RwdUsNetworksRow) GetPublishedDtaCount() int32`

GetPublishedDtaCount returns the PublishedDtaCount field if non-nil, zero value otherwise.

### GetPublishedDtaCountOk

`func (o *RwdUsNetworksRow) GetPublishedDtaCountOk() (*int32, bool)`

GetPublishedDtaCountOk returns a tuple with the PublishedDtaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDtaCount

`func (o *RwdUsNetworksRow) SetPublishedDtaCount(v int32)`

SetPublishedDtaCount sets PublishedDtaCount field to given value.

### HasPublishedDtaCount

`func (o *RwdUsNetworksRow) HasPublishedDtaCount() bool`

HasPublishedDtaCount returns a boolean if a field has been set.

### SetPublishedDtaCountNil

`func (o *RwdUsNetworksRow) SetPublishedDtaCountNil(b bool)`

 SetPublishedDtaCountNil sets the value for PublishedDtaCount to be an explicit nil

### UnsetPublishedDtaCount
`func (o *RwdUsNetworksRow) UnsetPublishedDtaCount()`

UnsetPublishedDtaCount ensures that no value is present for PublishedDtaCount, not even an explicit nil
### GetPublishedDuaCount

`func (o *RwdUsNetworksRow) GetPublishedDuaCount() int32`

GetPublishedDuaCount returns the PublishedDuaCount field if non-nil, zero value otherwise.

### GetPublishedDuaCountOk

`func (o *RwdUsNetworksRow) GetPublishedDuaCountOk() (*int32, bool)`

GetPublishedDuaCountOk returns a tuple with the PublishedDuaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDuaCount

`func (o *RwdUsNetworksRow) SetPublishedDuaCount(v int32)`

SetPublishedDuaCount sets PublishedDuaCount field to given value.

### HasPublishedDuaCount

`func (o *RwdUsNetworksRow) HasPublishedDuaCount() bool`

HasPublishedDuaCount returns a boolean if a field has been set.

### SetPublishedDuaCountNil

`func (o *RwdUsNetworksRow) SetPublishedDuaCountNil(b bool)`

 SetPublishedDuaCountNil sets the value for PublishedDuaCount to be an explicit nil

### UnsetPublishedDuaCount
`func (o *RwdUsNetworksRow) UnsetPublishedDuaCount()`

UnsetPublishedDuaCount ensures that no value is present for PublishedDuaCount, not even an explicit nil
### GetPublishedInstitutionCount

`func (o *RwdUsNetworksRow) GetPublishedInstitutionCount() int32`

GetPublishedInstitutionCount returns the PublishedInstitutionCount field if non-nil, zero value otherwise.

### GetPublishedInstitutionCountOk

`func (o *RwdUsNetworksRow) GetPublishedInstitutionCountOk() (*int32, bool)`

GetPublishedInstitutionCountOk returns a tuple with the PublishedInstitutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedInstitutionCount

`func (o *RwdUsNetworksRow) SetPublishedInstitutionCount(v int32)`

SetPublishedInstitutionCount sets PublishedInstitutionCount field to given value.

### HasPublishedInstitutionCount

`func (o *RwdUsNetworksRow) HasPublishedInstitutionCount() bool`

HasPublishedInstitutionCount returns a boolean if a field has been set.

### SetPublishedInstitutionCountNil

`func (o *RwdUsNetworksRow) SetPublishedInstitutionCountNil(b bool)`

 SetPublishedInstitutionCountNil sets the value for PublishedInstitutionCount to be an explicit nil

### UnsetPublishedInstitutionCount
`func (o *RwdUsNetworksRow) UnsetPublishedInstitutionCount()`

UnsetPublishedInstitutionCount ensures that no value is present for PublishedInstitutionCount, not even an explicit nil
### GetPublishedProjectCount

`func (o *RwdUsNetworksRow) GetPublishedProjectCount() int32`

GetPublishedProjectCount returns the PublishedProjectCount field if non-nil, zero value otherwise.

### GetPublishedProjectCountOk

`func (o *RwdUsNetworksRow) GetPublishedProjectCountOk() (*int32, bool)`

GetPublishedProjectCountOk returns a tuple with the PublishedProjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedProjectCount

`func (o *RwdUsNetworksRow) SetPublishedProjectCount(v int32)`

SetPublishedProjectCount sets PublishedProjectCount field to given value.

### HasPublishedProjectCount

`func (o *RwdUsNetworksRow) HasPublishedProjectCount() bool`

HasPublishedProjectCount returns a boolean if a field has been set.

### SetPublishedProjectCountNil

`func (o *RwdUsNetworksRow) SetPublishedProjectCountNil(b bool)`

 SetPublishedProjectCountNil sets the value for PublishedProjectCount to be an explicit nil

### UnsetPublishedProjectCount
`func (o *RwdUsNetworksRow) UnsetPublishedProjectCount()`

UnsetPublishedProjectCount ensures that no value is present for PublishedProjectCount, not even an explicit nil
### GetAsOf

`func (o *RwdUsNetworksRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *RwdUsNetworksRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *RwdUsNetworksRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *RwdUsNetworksRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *RwdUsNetworksRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *RwdUsNetworksRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetSourceUrl

`func (o *RwdUsNetworksRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *RwdUsNetworksRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *RwdUsNetworksRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *RwdUsNetworksRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *RwdUsNetworksRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *RwdUsNetworksRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetAgreementsUrl

`func (o *RwdUsNetworksRow) GetAgreementsUrl() string`

GetAgreementsUrl returns the AgreementsUrl field if non-nil, zero value otherwise.

### GetAgreementsUrlOk

`func (o *RwdUsNetworksRow) GetAgreementsUrlOk() (*string, bool)`

GetAgreementsUrlOk returns a tuple with the AgreementsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementsUrl

`func (o *RwdUsNetworksRow) SetAgreementsUrl(v string)`

SetAgreementsUrl sets AgreementsUrl field to given value.

### HasAgreementsUrl

`func (o *RwdUsNetworksRow) HasAgreementsUrl() bool`

HasAgreementsUrl returns a boolean if a field has been set.

### SetAgreementsUrlNil

`func (o *RwdUsNetworksRow) SetAgreementsUrlNil(b bool)`

 SetAgreementsUrlNil sets the value for AgreementsUrl to be an explicit nil

### UnsetAgreementsUrl
`func (o *RwdUsNetworksRow) UnsetAgreementsUrl()`

UnsetAgreementsUrl ensures that no value is present for AgreementsUrl, not even an explicit nil
### GetSignatoriesPageUrl

`func (o *RwdUsNetworksRow) GetSignatoriesPageUrl() string`

GetSignatoriesPageUrl returns the SignatoriesPageUrl field if non-nil, zero value otherwise.

### GetSignatoriesPageUrlOk

`func (o *RwdUsNetworksRow) GetSignatoriesPageUrlOk() (*string, bool)`

GetSignatoriesPageUrlOk returns a tuple with the SignatoriesPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatoriesPageUrl

`func (o *RwdUsNetworksRow) SetSignatoriesPageUrl(v string)`

SetSignatoriesPageUrl sets SignatoriesPageUrl field to given value.

### HasSignatoriesPageUrl

`func (o *RwdUsNetworksRow) HasSignatoriesPageUrl() bool`

HasSignatoriesPageUrl returns a boolean if a field has been set.

### SetSignatoriesPageUrlNil

`func (o *RwdUsNetworksRow) SetSignatoriesPageUrlNil(b bool)`

 SetSignatoriesPageUrlNil sets the value for SignatoriesPageUrl to be an explicit nil

### UnsetSignatoriesPageUrl
`func (o *RwdUsNetworksRow) UnsetSignatoriesPageUrl()`

UnsetSignatoriesPageUrl ensures that no value is present for SignatoriesPageUrl, not even an explicit nil
### GetRetrievedAt

`func (o *RwdUsNetworksRow) GetRetrievedAt() time.Time`

GetRetrievedAt returns the RetrievedAt field if non-nil, zero value otherwise.

### GetRetrievedAtOk

`func (o *RwdUsNetworksRow) GetRetrievedAtOk() (*time.Time, bool)`

GetRetrievedAtOk returns a tuple with the RetrievedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievedAt

`func (o *RwdUsNetworksRow) SetRetrievedAt(v time.Time)`

SetRetrievedAt sets RetrievedAt field to given value.

### HasRetrievedAt

`func (o *RwdUsNetworksRow) HasRetrievedAt() bool`

HasRetrievedAt returns a boolean if a field has been set.

### SetRetrievedAtNil

`func (o *RwdUsNetworksRow) SetRetrievedAtNil(b bool)`

 SetRetrievedAtNil sets the value for RetrievedAt to be an explicit nil

### UnsetRetrievedAt
`func (o *RwdUsNetworksRow) UnsetRetrievedAt()`

UnsetRetrievedAt ensures that no value is present for RetrievedAt, not even an explicit nil
### GetAttribution

`func (o *RwdUsNetworksRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdUsNetworksRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdUsNetworksRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdUsNetworksRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdUsNetworksRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdUsNetworksRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


