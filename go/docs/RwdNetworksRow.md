# RwdNetworksRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **NullableString** | The register&#39;s own Network ID. The entry key of the rwd_networks object type. | [optional] 
**NetworkName** | Pointer to **NullableString** | The network&#39;s name as the register spells it. | [optional] 
**NetworkNameKey** | Pointer to **NullableString** | The normalised name key the study and member edges join on. | [optional] 
**RegisterState** | Pointer to **NullableString** | self_registered_profile when the entry declares at least one of description, countries, therapeutic areas or funding source, catalogue_name_entry otherwise. is_encepp_partner is null on a catalogue name entry. | [optional] 
**IsSelfRegisteredProfile** | Pointer to **NullableBool** | True when registration_state is self_registered_profile. | [optional] 
**DeclaredProfileFields** | Pointer to **NullableInt32** | How many of the declarable profile fields the entry answered. | [optional] 
**DeclarableProfileFields** | Pointer to **NullableInt32** | How many profile fields the register offers; the denominator of declared_profile_fields. | [optional] 
**NetworkCountries** | Pointer to **NullableString** | The country text the network declares, as typed. | [optional] 
**PrimaryTherapeuticDiseaseAreas** | Pointer to **NullableString** | The primary therapeutic and disease areas the network declares, as typed; blank on most networks, so a facet over it would remove the largest networks. | [optional] 
**HasDeclaredTherapeuticArea** | Pointer to **NullableBool** | Whether the network declares a therapeutic area at all (the blank is the registrant&#39;s silence). | [optional] 
**SourcesOfFunding** | Pointer to **NullableString** | The sources of funding the network declares, as typed. | [optional] 
**NetworkWebsite** | Pointer to **NullableString** | The network&#39;s website as the register holds it. | [optional] 
**CatalogueUrl** | Pointer to **NullableString** | The address of this entry on the HMA-EMA catalogue site (the register&#39;s URL LINK). | [optional] 
**IsEnceppPartner** | Pointer to **NullableBool** | Whether the entry declares an ENCePP partnership; null on a catalogue name entry, where the export writes a default No that is not an answer. | [optional] 
**DeclaredMemberCount** | Pointer to **NullableInt32** | How many member institutions the network&#39;s own list names; zero says the network named nobody, not that it has no members. | [optional] 
**DeclaredMemberNameKeys** | Pointer to **[]string** | The normalised name keys of the declared members. | [optional] 
**DeclaredMemberNames** | Pointer to **[]string** | The names of the declared members as the register spells them. | [optional] 
**ListsAnyMember** | Pointer to **NullableBool** | Whether the network&#39;s own list names at least one member. | [optional] 
**DeliveredStudiesLifetime** | Pointer to **NullableInt32** | Distinct registered studies that named the network. Declared membership and delivered work are different facts and never share a column. | [optional] 
**DeliveredStudiesLast5Years** | Pointer to **NullableInt32** | Delivered studies that started within five years of the snapshot. | [optional] 
**DeliveredStudiesOngoing** | Pointer to **NullableInt32** | Delivered studies whose status is Ongoing. | [optional] 
**DeliveredStudiesPlanned** | Pointer to **NullableInt32** | Delivered studies whose status is Planned. | [optional] 
**MostRecentStudyStartOn** | Pointer to **NullableString** | The latest actual start date among the studies it conducted (or delivered, for a network). | [optional] 
**RecordUpdatedAt** | Pointer to **NullableTime** | When the registrant last edited this catalogue entry. A different clock from last_refreshed_on: it dates the web form, not the database. | [optional] 
**RecordUpdatedIsMassEdit** | Pointer to **NullableBool** | True when the record&#39;s last update date is a mass-edit date: one shared by an outsized part of the register, so it dates a migration and not the registrant. | [optional] 
**RecordUpdatedAtAttested** | Pointer to **NullableTime** | record_updated_at with the mass-edit dates removed (null for a mass-edit date or no timestamp): the column to sort a freshness question on. | [optional] 
**RecordUpdatedDateShare** | Pointer to **NullableFloat32** | The share of the register whose last update carries the same date, on every row so a mass-edit date is visible. | [optional] 
**RecordFirstPublishedAt** | Pointer to **NullableTime** | When the catalogue entry was first published. | [optional] 
**IsOutdated** | Pointer to **NullableBool** | EMA&#39;s own Outdated flag, carried through untouched; the register flags none of the data sources, so read false as EMA&#39;s answer and not as a currency finding. | [optional] 
**NetworkRegisterTotal** | Pointer to **NullableInt32** | The number of networks in the whole network register at the snapshot. | [optional] 
**NetworksListingAnyMember** | Pointer to **NullableInt32** | How many register networks list at least one member. | [optional] 
**MembershipEdgeTotal** | Pointer to **NullableInt32** | How many network-to-institution membership edges the member lists hold. | [optional] 
**InstitutionsInAnyMemberList** | Pointer to **NullableInt32** | How many register institutions at least one network member list names. | [optional] 
**InstitutionNameKeyTotal** | Pointer to **NullableInt32** | How many distinct institution name keys the register holds. | [optional] 
**NetworksDeclaringATherapeuticArea** | Pointer to **NullableInt32** | How many register networks declare a therapeutic area. | [optional] 
**NetworksWithAnyDeliveredStudy** | Pointer to **NullableInt32** | How many register networks at least one registered study named. | [optional] 
**SelfRegisteredProfileTotal** | Pointer to **NullableInt32** | How many register entries are self-registered profiles. | [optional] 
**CatalogueNameEntryTotal** | Pointer to **NullableInt32** | How many register entries are catalogue name entries. | [optional] 
**SelfRegisteredProfileEnceppYesTotal** | Pointer to **NullableInt32** | How many self-registered networks declare an ENCePP partnership. | [optional] 
**CatalogueNameEntryEnceppYesTotal** | Pointer to **NullableInt32** | How many catalogue name entries declare one (the mart nulls the declaration there, so this is the false-No check). | [optional] 
**RegisterSourceCount** | Pointer to **NullableInt32** | The number of data sources in the whole data source register at the snapshot: the denominator every source count is over. Not a count of real-world data sources in the world. | [optional] 
**StudyRegisterTotal** | Pointer to **NullableInt32** | The number of studies in the whole study register at the snapshot. | [optional] 
**StudiesNamingARegisteredSource** | Pointer to **NullableInt32** | How many registered studies name at least one data source that is itself registered; the denominator under study_count. | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdNetworksRow

`func NewRwdNetworksRow() *RwdNetworksRow`

NewRwdNetworksRow instantiates a new RwdNetworksRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdNetworksRowWithDefaults

`func NewRwdNetworksRowWithDefaults() *RwdNetworksRow`

NewRwdNetworksRowWithDefaults instantiates a new RwdNetworksRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *RwdNetworksRow) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *RwdNetworksRow) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *RwdNetworksRow) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *RwdNetworksRow) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### SetNetworkIdNil

`func (o *RwdNetworksRow) SetNetworkIdNil(b bool)`

 SetNetworkIdNil sets the value for NetworkId to be an explicit nil

### UnsetNetworkId
`func (o *RwdNetworksRow) UnsetNetworkId()`

UnsetNetworkId ensures that no value is present for NetworkId, not even an explicit nil
### GetNetworkName

`func (o *RwdNetworksRow) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *RwdNetworksRow) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *RwdNetworksRow) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *RwdNetworksRow) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### SetNetworkNameNil

`func (o *RwdNetworksRow) SetNetworkNameNil(b bool)`

 SetNetworkNameNil sets the value for NetworkName to be an explicit nil

### UnsetNetworkName
`func (o *RwdNetworksRow) UnsetNetworkName()`

UnsetNetworkName ensures that no value is present for NetworkName, not even an explicit nil
### GetNetworkNameKey

`func (o *RwdNetworksRow) GetNetworkNameKey() string`

GetNetworkNameKey returns the NetworkNameKey field if non-nil, zero value otherwise.

### GetNetworkNameKeyOk

`func (o *RwdNetworksRow) GetNetworkNameKeyOk() (*string, bool)`

GetNetworkNameKeyOk returns a tuple with the NetworkNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkNameKey

`func (o *RwdNetworksRow) SetNetworkNameKey(v string)`

SetNetworkNameKey sets NetworkNameKey field to given value.

### HasNetworkNameKey

`func (o *RwdNetworksRow) HasNetworkNameKey() bool`

HasNetworkNameKey returns a boolean if a field has been set.

### SetNetworkNameKeyNil

`func (o *RwdNetworksRow) SetNetworkNameKeyNil(b bool)`

 SetNetworkNameKeyNil sets the value for NetworkNameKey to be an explicit nil

### UnsetNetworkNameKey
`func (o *RwdNetworksRow) UnsetNetworkNameKey()`

UnsetNetworkNameKey ensures that no value is present for NetworkNameKey, not even an explicit nil
### GetRegisterState

`func (o *RwdNetworksRow) GetRegisterState() string`

GetRegisterState returns the RegisterState field if non-nil, zero value otherwise.

### GetRegisterStateOk

`func (o *RwdNetworksRow) GetRegisterStateOk() (*string, bool)`

GetRegisterStateOk returns a tuple with the RegisterState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterState

`func (o *RwdNetworksRow) SetRegisterState(v string)`

SetRegisterState sets RegisterState field to given value.

### HasRegisterState

`func (o *RwdNetworksRow) HasRegisterState() bool`

HasRegisterState returns a boolean if a field has been set.

### SetRegisterStateNil

`func (o *RwdNetworksRow) SetRegisterStateNil(b bool)`

 SetRegisterStateNil sets the value for RegisterState to be an explicit nil

### UnsetRegisterState
`func (o *RwdNetworksRow) UnsetRegisterState()`

UnsetRegisterState ensures that no value is present for RegisterState, not even an explicit nil
### GetIsSelfRegisteredProfile

`func (o *RwdNetworksRow) GetIsSelfRegisteredProfile() bool`

GetIsSelfRegisteredProfile returns the IsSelfRegisteredProfile field if non-nil, zero value otherwise.

### GetIsSelfRegisteredProfileOk

`func (o *RwdNetworksRow) GetIsSelfRegisteredProfileOk() (*bool, bool)`

GetIsSelfRegisteredProfileOk returns a tuple with the IsSelfRegisteredProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfRegisteredProfile

`func (o *RwdNetworksRow) SetIsSelfRegisteredProfile(v bool)`

SetIsSelfRegisteredProfile sets IsSelfRegisteredProfile field to given value.

### HasIsSelfRegisteredProfile

`func (o *RwdNetworksRow) HasIsSelfRegisteredProfile() bool`

HasIsSelfRegisteredProfile returns a boolean if a field has been set.

### SetIsSelfRegisteredProfileNil

`func (o *RwdNetworksRow) SetIsSelfRegisteredProfileNil(b bool)`

 SetIsSelfRegisteredProfileNil sets the value for IsSelfRegisteredProfile to be an explicit nil

### UnsetIsSelfRegisteredProfile
`func (o *RwdNetworksRow) UnsetIsSelfRegisteredProfile()`

UnsetIsSelfRegisteredProfile ensures that no value is present for IsSelfRegisteredProfile, not even an explicit nil
### GetDeclaredProfileFields

`func (o *RwdNetworksRow) GetDeclaredProfileFields() int32`

GetDeclaredProfileFields returns the DeclaredProfileFields field if non-nil, zero value otherwise.

### GetDeclaredProfileFieldsOk

`func (o *RwdNetworksRow) GetDeclaredProfileFieldsOk() (*int32, bool)`

GetDeclaredProfileFieldsOk returns a tuple with the DeclaredProfileFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredProfileFields

`func (o *RwdNetworksRow) SetDeclaredProfileFields(v int32)`

SetDeclaredProfileFields sets DeclaredProfileFields field to given value.

### HasDeclaredProfileFields

`func (o *RwdNetworksRow) HasDeclaredProfileFields() bool`

HasDeclaredProfileFields returns a boolean if a field has been set.

### SetDeclaredProfileFieldsNil

`func (o *RwdNetworksRow) SetDeclaredProfileFieldsNil(b bool)`

 SetDeclaredProfileFieldsNil sets the value for DeclaredProfileFields to be an explicit nil

### UnsetDeclaredProfileFields
`func (o *RwdNetworksRow) UnsetDeclaredProfileFields()`

UnsetDeclaredProfileFields ensures that no value is present for DeclaredProfileFields, not even an explicit nil
### GetDeclarableProfileFields

`func (o *RwdNetworksRow) GetDeclarableProfileFields() int32`

GetDeclarableProfileFields returns the DeclarableProfileFields field if non-nil, zero value otherwise.

### GetDeclarableProfileFieldsOk

`func (o *RwdNetworksRow) GetDeclarableProfileFieldsOk() (*int32, bool)`

GetDeclarableProfileFieldsOk returns a tuple with the DeclarableProfileFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarableProfileFields

`func (o *RwdNetworksRow) SetDeclarableProfileFields(v int32)`

SetDeclarableProfileFields sets DeclarableProfileFields field to given value.

### HasDeclarableProfileFields

`func (o *RwdNetworksRow) HasDeclarableProfileFields() bool`

HasDeclarableProfileFields returns a boolean if a field has been set.

### SetDeclarableProfileFieldsNil

`func (o *RwdNetworksRow) SetDeclarableProfileFieldsNil(b bool)`

 SetDeclarableProfileFieldsNil sets the value for DeclarableProfileFields to be an explicit nil

### UnsetDeclarableProfileFields
`func (o *RwdNetworksRow) UnsetDeclarableProfileFields()`

UnsetDeclarableProfileFields ensures that no value is present for DeclarableProfileFields, not even an explicit nil
### GetNetworkCountries

`func (o *RwdNetworksRow) GetNetworkCountries() string`

GetNetworkCountries returns the NetworkCountries field if non-nil, zero value otherwise.

### GetNetworkCountriesOk

`func (o *RwdNetworksRow) GetNetworkCountriesOk() (*string, bool)`

GetNetworkCountriesOk returns a tuple with the NetworkCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCountries

`func (o *RwdNetworksRow) SetNetworkCountries(v string)`

SetNetworkCountries sets NetworkCountries field to given value.

### HasNetworkCountries

`func (o *RwdNetworksRow) HasNetworkCountries() bool`

HasNetworkCountries returns a boolean if a field has been set.

### SetNetworkCountriesNil

`func (o *RwdNetworksRow) SetNetworkCountriesNil(b bool)`

 SetNetworkCountriesNil sets the value for NetworkCountries to be an explicit nil

### UnsetNetworkCountries
`func (o *RwdNetworksRow) UnsetNetworkCountries()`

UnsetNetworkCountries ensures that no value is present for NetworkCountries, not even an explicit nil
### GetPrimaryTherapeuticDiseaseAreas

`func (o *RwdNetworksRow) GetPrimaryTherapeuticDiseaseAreas() string`

GetPrimaryTherapeuticDiseaseAreas returns the PrimaryTherapeuticDiseaseAreas field if non-nil, zero value otherwise.

### GetPrimaryTherapeuticDiseaseAreasOk

`func (o *RwdNetworksRow) GetPrimaryTherapeuticDiseaseAreasOk() (*string, bool)`

GetPrimaryTherapeuticDiseaseAreasOk returns a tuple with the PrimaryTherapeuticDiseaseAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTherapeuticDiseaseAreas

`func (o *RwdNetworksRow) SetPrimaryTherapeuticDiseaseAreas(v string)`

SetPrimaryTherapeuticDiseaseAreas sets PrimaryTherapeuticDiseaseAreas field to given value.

### HasPrimaryTherapeuticDiseaseAreas

`func (o *RwdNetworksRow) HasPrimaryTherapeuticDiseaseAreas() bool`

HasPrimaryTherapeuticDiseaseAreas returns a boolean if a field has been set.

### SetPrimaryTherapeuticDiseaseAreasNil

`func (o *RwdNetworksRow) SetPrimaryTherapeuticDiseaseAreasNil(b bool)`

 SetPrimaryTherapeuticDiseaseAreasNil sets the value for PrimaryTherapeuticDiseaseAreas to be an explicit nil

### UnsetPrimaryTherapeuticDiseaseAreas
`func (o *RwdNetworksRow) UnsetPrimaryTherapeuticDiseaseAreas()`

UnsetPrimaryTherapeuticDiseaseAreas ensures that no value is present for PrimaryTherapeuticDiseaseAreas, not even an explicit nil
### GetHasDeclaredTherapeuticArea

`func (o *RwdNetworksRow) GetHasDeclaredTherapeuticArea() bool`

GetHasDeclaredTherapeuticArea returns the HasDeclaredTherapeuticArea field if non-nil, zero value otherwise.

### GetHasDeclaredTherapeuticAreaOk

`func (o *RwdNetworksRow) GetHasDeclaredTherapeuticAreaOk() (*bool, bool)`

GetHasDeclaredTherapeuticAreaOk returns a tuple with the HasDeclaredTherapeuticArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDeclaredTherapeuticArea

`func (o *RwdNetworksRow) SetHasDeclaredTherapeuticArea(v bool)`

SetHasDeclaredTherapeuticArea sets HasDeclaredTherapeuticArea field to given value.

### HasHasDeclaredTherapeuticArea

`func (o *RwdNetworksRow) HasHasDeclaredTherapeuticArea() bool`

HasHasDeclaredTherapeuticArea returns a boolean if a field has been set.

### SetHasDeclaredTherapeuticAreaNil

`func (o *RwdNetworksRow) SetHasDeclaredTherapeuticAreaNil(b bool)`

 SetHasDeclaredTherapeuticAreaNil sets the value for HasDeclaredTherapeuticArea to be an explicit nil

### UnsetHasDeclaredTherapeuticArea
`func (o *RwdNetworksRow) UnsetHasDeclaredTherapeuticArea()`

UnsetHasDeclaredTherapeuticArea ensures that no value is present for HasDeclaredTherapeuticArea, not even an explicit nil
### GetSourcesOfFunding

`func (o *RwdNetworksRow) GetSourcesOfFunding() string`

GetSourcesOfFunding returns the SourcesOfFunding field if non-nil, zero value otherwise.

### GetSourcesOfFundingOk

`func (o *RwdNetworksRow) GetSourcesOfFundingOk() (*string, bool)`

GetSourcesOfFundingOk returns a tuple with the SourcesOfFunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesOfFunding

`func (o *RwdNetworksRow) SetSourcesOfFunding(v string)`

SetSourcesOfFunding sets SourcesOfFunding field to given value.

### HasSourcesOfFunding

`func (o *RwdNetworksRow) HasSourcesOfFunding() bool`

HasSourcesOfFunding returns a boolean if a field has been set.

### SetSourcesOfFundingNil

`func (o *RwdNetworksRow) SetSourcesOfFundingNil(b bool)`

 SetSourcesOfFundingNil sets the value for SourcesOfFunding to be an explicit nil

### UnsetSourcesOfFunding
`func (o *RwdNetworksRow) UnsetSourcesOfFunding()`

UnsetSourcesOfFunding ensures that no value is present for SourcesOfFunding, not even an explicit nil
### GetNetworkWebsite

`func (o *RwdNetworksRow) GetNetworkWebsite() string`

GetNetworkWebsite returns the NetworkWebsite field if non-nil, zero value otherwise.

### GetNetworkWebsiteOk

`func (o *RwdNetworksRow) GetNetworkWebsiteOk() (*string, bool)`

GetNetworkWebsiteOk returns a tuple with the NetworkWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkWebsite

`func (o *RwdNetworksRow) SetNetworkWebsite(v string)`

SetNetworkWebsite sets NetworkWebsite field to given value.

### HasNetworkWebsite

`func (o *RwdNetworksRow) HasNetworkWebsite() bool`

HasNetworkWebsite returns a boolean if a field has been set.

### SetNetworkWebsiteNil

`func (o *RwdNetworksRow) SetNetworkWebsiteNil(b bool)`

 SetNetworkWebsiteNil sets the value for NetworkWebsite to be an explicit nil

### UnsetNetworkWebsite
`func (o *RwdNetworksRow) UnsetNetworkWebsite()`

UnsetNetworkWebsite ensures that no value is present for NetworkWebsite, not even an explicit nil
### GetCatalogueUrl

`func (o *RwdNetworksRow) GetCatalogueUrl() string`

GetCatalogueUrl returns the CatalogueUrl field if non-nil, zero value otherwise.

### GetCatalogueUrlOk

`func (o *RwdNetworksRow) GetCatalogueUrlOk() (*string, bool)`

GetCatalogueUrlOk returns a tuple with the CatalogueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueUrl

`func (o *RwdNetworksRow) SetCatalogueUrl(v string)`

SetCatalogueUrl sets CatalogueUrl field to given value.

### HasCatalogueUrl

`func (o *RwdNetworksRow) HasCatalogueUrl() bool`

HasCatalogueUrl returns a boolean if a field has been set.

### SetCatalogueUrlNil

`func (o *RwdNetworksRow) SetCatalogueUrlNil(b bool)`

 SetCatalogueUrlNil sets the value for CatalogueUrl to be an explicit nil

### UnsetCatalogueUrl
`func (o *RwdNetworksRow) UnsetCatalogueUrl()`

UnsetCatalogueUrl ensures that no value is present for CatalogueUrl, not even an explicit nil
### GetIsEnceppPartner

`func (o *RwdNetworksRow) GetIsEnceppPartner() bool`

GetIsEnceppPartner returns the IsEnceppPartner field if non-nil, zero value otherwise.

### GetIsEnceppPartnerOk

`func (o *RwdNetworksRow) GetIsEnceppPartnerOk() (*bool, bool)`

GetIsEnceppPartnerOk returns a tuple with the IsEnceppPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnceppPartner

`func (o *RwdNetworksRow) SetIsEnceppPartner(v bool)`

SetIsEnceppPartner sets IsEnceppPartner field to given value.

### HasIsEnceppPartner

`func (o *RwdNetworksRow) HasIsEnceppPartner() bool`

HasIsEnceppPartner returns a boolean if a field has been set.

### SetIsEnceppPartnerNil

`func (o *RwdNetworksRow) SetIsEnceppPartnerNil(b bool)`

 SetIsEnceppPartnerNil sets the value for IsEnceppPartner to be an explicit nil

### UnsetIsEnceppPartner
`func (o *RwdNetworksRow) UnsetIsEnceppPartner()`

UnsetIsEnceppPartner ensures that no value is present for IsEnceppPartner, not even an explicit nil
### GetDeclaredMemberCount

`func (o *RwdNetworksRow) GetDeclaredMemberCount() int32`

GetDeclaredMemberCount returns the DeclaredMemberCount field if non-nil, zero value otherwise.

### GetDeclaredMemberCountOk

`func (o *RwdNetworksRow) GetDeclaredMemberCountOk() (*int32, bool)`

GetDeclaredMemberCountOk returns a tuple with the DeclaredMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredMemberCount

`func (o *RwdNetworksRow) SetDeclaredMemberCount(v int32)`

SetDeclaredMemberCount sets DeclaredMemberCount field to given value.

### HasDeclaredMemberCount

`func (o *RwdNetworksRow) HasDeclaredMemberCount() bool`

HasDeclaredMemberCount returns a boolean if a field has been set.

### SetDeclaredMemberCountNil

`func (o *RwdNetworksRow) SetDeclaredMemberCountNil(b bool)`

 SetDeclaredMemberCountNil sets the value for DeclaredMemberCount to be an explicit nil

### UnsetDeclaredMemberCount
`func (o *RwdNetworksRow) UnsetDeclaredMemberCount()`

UnsetDeclaredMemberCount ensures that no value is present for DeclaredMemberCount, not even an explicit nil
### GetDeclaredMemberNameKeys

`func (o *RwdNetworksRow) GetDeclaredMemberNameKeys() []string`

GetDeclaredMemberNameKeys returns the DeclaredMemberNameKeys field if non-nil, zero value otherwise.

### GetDeclaredMemberNameKeysOk

`func (o *RwdNetworksRow) GetDeclaredMemberNameKeysOk() (*[]string, bool)`

GetDeclaredMemberNameKeysOk returns a tuple with the DeclaredMemberNameKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredMemberNameKeys

`func (o *RwdNetworksRow) SetDeclaredMemberNameKeys(v []string)`

SetDeclaredMemberNameKeys sets DeclaredMemberNameKeys field to given value.

### HasDeclaredMemberNameKeys

`func (o *RwdNetworksRow) HasDeclaredMemberNameKeys() bool`

HasDeclaredMemberNameKeys returns a boolean if a field has been set.

### SetDeclaredMemberNameKeysNil

`func (o *RwdNetworksRow) SetDeclaredMemberNameKeysNil(b bool)`

 SetDeclaredMemberNameKeysNil sets the value for DeclaredMemberNameKeys to be an explicit nil

### UnsetDeclaredMemberNameKeys
`func (o *RwdNetworksRow) UnsetDeclaredMemberNameKeys()`

UnsetDeclaredMemberNameKeys ensures that no value is present for DeclaredMemberNameKeys, not even an explicit nil
### GetDeclaredMemberNames

`func (o *RwdNetworksRow) GetDeclaredMemberNames() []string`

GetDeclaredMemberNames returns the DeclaredMemberNames field if non-nil, zero value otherwise.

### GetDeclaredMemberNamesOk

`func (o *RwdNetworksRow) GetDeclaredMemberNamesOk() (*[]string, bool)`

GetDeclaredMemberNamesOk returns a tuple with the DeclaredMemberNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredMemberNames

`func (o *RwdNetworksRow) SetDeclaredMemberNames(v []string)`

SetDeclaredMemberNames sets DeclaredMemberNames field to given value.

### HasDeclaredMemberNames

`func (o *RwdNetworksRow) HasDeclaredMemberNames() bool`

HasDeclaredMemberNames returns a boolean if a field has been set.

### SetDeclaredMemberNamesNil

`func (o *RwdNetworksRow) SetDeclaredMemberNamesNil(b bool)`

 SetDeclaredMemberNamesNil sets the value for DeclaredMemberNames to be an explicit nil

### UnsetDeclaredMemberNames
`func (o *RwdNetworksRow) UnsetDeclaredMemberNames()`

UnsetDeclaredMemberNames ensures that no value is present for DeclaredMemberNames, not even an explicit nil
### GetListsAnyMember

`func (o *RwdNetworksRow) GetListsAnyMember() bool`

GetListsAnyMember returns the ListsAnyMember field if non-nil, zero value otherwise.

### GetListsAnyMemberOk

`func (o *RwdNetworksRow) GetListsAnyMemberOk() (*bool, bool)`

GetListsAnyMemberOk returns a tuple with the ListsAnyMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListsAnyMember

`func (o *RwdNetworksRow) SetListsAnyMember(v bool)`

SetListsAnyMember sets ListsAnyMember field to given value.

### HasListsAnyMember

`func (o *RwdNetworksRow) HasListsAnyMember() bool`

HasListsAnyMember returns a boolean if a field has been set.

### SetListsAnyMemberNil

`func (o *RwdNetworksRow) SetListsAnyMemberNil(b bool)`

 SetListsAnyMemberNil sets the value for ListsAnyMember to be an explicit nil

### UnsetListsAnyMember
`func (o *RwdNetworksRow) UnsetListsAnyMember()`

UnsetListsAnyMember ensures that no value is present for ListsAnyMember, not even an explicit nil
### GetDeliveredStudiesLifetime

`func (o *RwdNetworksRow) GetDeliveredStudiesLifetime() int32`

GetDeliveredStudiesLifetime returns the DeliveredStudiesLifetime field if non-nil, zero value otherwise.

### GetDeliveredStudiesLifetimeOk

`func (o *RwdNetworksRow) GetDeliveredStudiesLifetimeOk() (*int32, bool)`

GetDeliveredStudiesLifetimeOk returns a tuple with the DeliveredStudiesLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredStudiesLifetime

`func (o *RwdNetworksRow) SetDeliveredStudiesLifetime(v int32)`

SetDeliveredStudiesLifetime sets DeliveredStudiesLifetime field to given value.

### HasDeliveredStudiesLifetime

`func (o *RwdNetworksRow) HasDeliveredStudiesLifetime() bool`

HasDeliveredStudiesLifetime returns a boolean if a field has been set.

### SetDeliveredStudiesLifetimeNil

`func (o *RwdNetworksRow) SetDeliveredStudiesLifetimeNil(b bool)`

 SetDeliveredStudiesLifetimeNil sets the value for DeliveredStudiesLifetime to be an explicit nil

### UnsetDeliveredStudiesLifetime
`func (o *RwdNetworksRow) UnsetDeliveredStudiesLifetime()`

UnsetDeliveredStudiesLifetime ensures that no value is present for DeliveredStudiesLifetime, not even an explicit nil
### GetDeliveredStudiesLast5Years

`func (o *RwdNetworksRow) GetDeliveredStudiesLast5Years() int32`

GetDeliveredStudiesLast5Years returns the DeliveredStudiesLast5Years field if non-nil, zero value otherwise.

### GetDeliveredStudiesLast5YearsOk

`func (o *RwdNetworksRow) GetDeliveredStudiesLast5YearsOk() (*int32, bool)`

GetDeliveredStudiesLast5YearsOk returns a tuple with the DeliveredStudiesLast5Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredStudiesLast5Years

`func (o *RwdNetworksRow) SetDeliveredStudiesLast5Years(v int32)`

SetDeliveredStudiesLast5Years sets DeliveredStudiesLast5Years field to given value.

### HasDeliveredStudiesLast5Years

`func (o *RwdNetworksRow) HasDeliveredStudiesLast5Years() bool`

HasDeliveredStudiesLast5Years returns a boolean if a field has been set.

### SetDeliveredStudiesLast5YearsNil

`func (o *RwdNetworksRow) SetDeliveredStudiesLast5YearsNil(b bool)`

 SetDeliveredStudiesLast5YearsNil sets the value for DeliveredStudiesLast5Years to be an explicit nil

### UnsetDeliveredStudiesLast5Years
`func (o *RwdNetworksRow) UnsetDeliveredStudiesLast5Years()`

UnsetDeliveredStudiesLast5Years ensures that no value is present for DeliveredStudiesLast5Years, not even an explicit nil
### GetDeliveredStudiesOngoing

`func (o *RwdNetworksRow) GetDeliveredStudiesOngoing() int32`

GetDeliveredStudiesOngoing returns the DeliveredStudiesOngoing field if non-nil, zero value otherwise.

### GetDeliveredStudiesOngoingOk

`func (o *RwdNetworksRow) GetDeliveredStudiesOngoingOk() (*int32, bool)`

GetDeliveredStudiesOngoingOk returns a tuple with the DeliveredStudiesOngoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredStudiesOngoing

`func (o *RwdNetworksRow) SetDeliveredStudiesOngoing(v int32)`

SetDeliveredStudiesOngoing sets DeliveredStudiesOngoing field to given value.

### HasDeliveredStudiesOngoing

`func (o *RwdNetworksRow) HasDeliveredStudiesOngoing() bool`

HasDeliveredStudiesOngoing returns a boolean if a field has been set.

### SetDeliveredStudiesOngoingNil

`func (o *RwdNetworksRow) SetDeliveredStudiesOngoingNil(b bool)`

 SetDeliveredStudiesOngoingNil sets the value for DeliveredStudiesOngoing to be an explicit nil

### UnsetDeliveredStudiesOngoing
`func (o *RwdNetworksRow) UnsetDeliveredStudiesOngoing()`

UnsetDeliveredStudiesOngoing ensures that no value is present for DeliveredStudiesOngoing, not even an explicit nil
### GetDeliveredStudiesPlanned

`func (o *RwdNetworksRow) GetDeliveredStudiesPlanned() int32`

GetDeliveredStudiesPlanned returns the DeliveredStudiesPlanned field if non-nil, zero value otherwise.

### GetDeliveredStudiesPlannedOk

`func (o *RwdNetworksRow) GetDeliveredStudiesPlannedOk() (*int32, bool)`

GetDeliveredStudiesPlannedOk returns a tuple with the DeliveredStudiesPlanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredStudiesPlanned

`func (o *RwdNetworksRow) SetDeliveredStudiesPlanned(v int32)`

SetDeliveredStudiesPlanned sets DeliveredStudiesPlanned field to given value.

### HasDeliveredStudiesPlanned

`func (o *RwdNetworksRow) HasDeliveredStudiesPlanned() bool`

HasDeliveredStudiesPlanned returns a boolean if a field has been set.

### SetDeliveredStudiesPlannedNil

`func (o *RwdNetworksRow) SetDeliveredStudiesPlannedNil(b bool)`

 SetDeliveredStudiesPlannedNil sets the value for DeliveredStudiesPlanned to be an explicit nil

### UnsetDeliveredStudiesPlanned
`func (o *RwdNetworksRow) UnsetDeliveredStudiesPlanned()`

UnsetDeliveredStudiesPlanned ensures that no value is present for DeliveredStudiesPlanned, not even an explicit nil
### GetMostRecentStudyStartOn

`func (o *RwdNetworksRow) GetMostRecentStudyStartOn() string`

GetMostRecentStudyStartOn returns the MostRecentStudyStartOn field if non-nil, zero value otherwise.

### GetMostRecentStudyStartOnOk

`func (o *RwdNetworksRow) GetMostRecentStudyStartOnOk() (*string, bool)`

GetMostRecentStudyStartOnOk returns a tuple with the MostRecentStudyStartOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentStudyStartOn

`func (o *RwdNetworksRow) SetMostRecentStudyStartOn(v string)`

SetMostRecentStudyStartOn sets MostRecentStudyStartOn field to given value.

### HasMostRecentStudyStartOn

`func (o *RwdNetworksRow) HasMostRecentStudyStartOn() bool`

HasMostRecentStudyStartOn returns a boolean if a field has been set.

### SetMostRecentStudyStartOnNil

`func (o *RwdNetworksRow) SetMostRecentStudyStartOnNil(b bool)`

 SetMostRecentStudyStartOnNil sets the value for MostRecentStudyStartOn to be an explicit nil

### UnsetMostRecentStudyStartOn
`func (o *RwdNetworksRow) UnsetMostRecentStudyStartOn()`

UnsetMostRecentStudyStartOn ensures that no value is present for MostRecentStudyStartOn, not even an explicit nil
### GetRecordUpdatedAt

`func (o *RwdNetworksRow) GetRecordUpdatedAt() time.Time`

GetRecordUpdatedAt returns the RecordUpdatedAt field if non-nil, zero value otherwise.

### GetRecordUpdatedAtOk

`func (o *RwdNetworksRow) GetRecordUpdatedAtOk() (*time.Time, bool)`

GetRecordUpdatedAtOk returns a tuple with the RecordUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordUpdatedAt

`func (o *RwdNetworksRow) SetRecordUpdatedAt(v time.Time)`

SetRecordUpdatedAt sets RecordUpdatedAt field to given value.

### HasRecordUpdatedAt

`func (o *RwdNetworksRow) HasRecordUpdatedAt() bool`

HasRecordUpdatedAt returns a boolean if a field has been set.

### SetRecordUpdatedAtNil

`func (o *RwdNetworksRow) SetRecordUpdatedAtNil(b bool)`

 SetRecordUpdatedAtNil sets the value for RecordUpdatedAt to be an explicit nil

### UnsetRecordUpdatedAt
`func (o *RwdNetworksRow) UnsetRecordUpdatedAt()`

UnsetRecordUpdatedAt ensures that no value is present for RecordUpdatedAt, not even an explicit nil
### GetRecordUpdatedIsMassEdit

`func (o *RwdNetworksRow) GetRecordUpdatedIsMassEdit() bool`

GetRecordUpdatedIsMassEdit returns the RecordUpdatedIsMassEdit field if non-nil, zero value otherwise.

### GetRecordUpdatedIsMassEditOk

`func (o *RwdNetworksRow) GetRecordUpdatedIsMassEditOk() (*bool, bool)`

GetRecordUpdatedIsMassEditOk returns a tuple with the RecordUpdatedIsMassEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordUpdatedIsMassEdit

`func (o *RwdNetworksRow) SetRecordUpdatedIsMassEdit(v bool)`

SetRecordUpdatedIsMassEdit sets RecordUpdatedIsMassEdit field to given value.

### HasRecordUpdatedIsMassEdit

`func (o *RwdNetworksRow) HasRecordUpdatedIsMassEdit() bool`

HasRecordUpdatedIsMassEdit returns a boolean if a field has been set.

### SetRecordUpdatedIsMassEditNil

`func (o *RwdNetworksRow) SetRecordUpdatedIsMassEditNil(b bool)`

 SetRecordUpdatedIsMassEditNil sets the value for RecordUpdatedIsMassEdit to be an explicit nil

### UnsetRecordUpdatedIsMassEdit
`func (o *RwdNetworksRow) UnsetRecordUpdatedIsMassEdit()`

UnsetRecordUpdatedIsMassEdit ensures that no value is present for RecordUpdatedIsMassEdit, not even an explicit nil
### GetRecordUpdatedAtAttested

`func (o *RwdNetworksRow) GetRecordUpdatedAtAttested() time.Time`

GetRecordUpdatedAtAttested returns the RecordUpdatedAtAttested field if non-nil, zero value otherwise.

### GetRecordUpdatedAtAttestedOk

`func (o *RwdNetworksRow) GetRecordUpdatedAtAttestedOk() (*time.Time, bool)`

GetRecordUpdatedAtAttestedOk returns a tuple with the RecordUpdatedAtAttested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordUpdatedAtAttested

`func (o *RwdNetworksRow) SetRecordUpdatedAtAttested(v time.Time)`

SetRecordUpdatedAtAttested sets RecordUpdatedAtAttested field to given value.

### HasRecordUpdatedAtAttested

`func (o *RwdNetworksRow) HasRecordUpdatedAtAttested() bool`

HasRecordUpdatedAtAttested returns a boolean if a field has been set.

### SetRecordUpdatedAtAttestedNil

`func (o *RwdNetworksRow) SetRecordUpdatedAtAttestedNil(b bool)`

 SetRecordUpdatedAtAttestedNil sets the value for RecordUpdatedAtAttested to be an explicit nil

### UnsetRecordUpdatedAtAttested
`func (o *RwdNetworksRow) UnsetRecordUpdatedAtAttested()`

UnsetRecordUpdatedAtAttested ensures that no value is present for RecordUpdatedAtAttested, not even an explicit nil
### GetRecordUpdatedDateShare

`func (o *RwdNetworksRow) GetRecordUpdatedDateShare() float32`

GetRecordUpdatedDateShare returns the RecordUpdatedDateShare field if non-nil, zero value otherwise.

### GetRecordUpdatedDateShareOk

`func (o *RwdNetworksRow) GetRecordUpdatedDateShareOk() (*float32, bool)`

GetRecordUpdatedDateShareOk returns a tuple with the RecordUpdatedDateShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordUpdatedDateShare

`func (o *RwdNetworksRow) SetRecordUpdatedDateShare(v float32)`

SetRecordUpdatedDateShare sets RecordUpdatedDateShare field to given value.

### HasRecordUpdatedDateShare

`func (o *RwdNetworksRow) HasRecordUpdatedDateShare() bool`

HasRecordUpdatedDateShare returns a boolean if a field has been set.

### SetRecordUpdatedDateShareNil

`func (o *RwdNetworksRow) SetRecordUpdatedDateShareNil(b bool)`

 SetRecordUpdatedDateShareNil sets the value for RecordUpdatedDateShare to be an explicit nil

### UnsetRecordUpdatedDateShare
`func (o *RwdNetworksRow) UnsetRecordUpdatedDateShare()`

UnsetRecordUpdatedDateShare ensures that no value is present for RecordUpdatedDateShare, not even an explicit nil
### GetRecordFirstPublishedAt

`func (o *RwdNetworksRow) GetRecordFirstPublishedAt() time.Time`

GetRecordFirstPublishedAt returns the RecordFirstPublishedAt field if non-nil, zero value otherwise.

### GetRecordFirstPublishedAtOk

`func (o *RwdNetworksRow) GetRecordFirstPublishedAtOk() (*time.Time, bool)`

GetRecordFirstPublishedAtOk returns a tuple with the RecordFirstPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordFirstPublishedAt

`func (o *RwdNetworksRow) SetRecordFirstPublishedAt(v time.Time)`

SetRecordFirstPublishedAt sets RecordFirstPublishedAt field to given value.

### HasRecordFirstPublishedAt

`func (o *RwdNetworksRow) HasRecordFirstPublishedAt() bool`

HasRecordFirstPublishedAt returns a boolean if a field has been set.

### SetRecordFirstPublishedAtNil

`func (o *RwdNetworksRow) SetRecordFirstPublishedAtNil(b bool)`

 SetRecordFirstPublishedAtNil sets the value for RecordFirstPublishedAt to be an explicit nil

### UnsetRecordFirstPublishedAt
`func (o *RwdNetworksRow) UnsetRecordFirstPublishedAt()`

UnsetRecordFirstPublishedAt ensures that no value is present for RecordFirstPublishedAt, not even an explicit nil
### GetIsOutdated

`func (o *RwdNetworksRow) GetIsOutdated() bool`

GetIsOutdated returns the IsOutdated field if non-nil, zero value otherwise.

### GetIsOutdatedOk

`func (o *RwdNetworksRow) GetIsOutdatedOk() (*bool, bool)`

GetIsOutdatedOk returns a tuple with the IsOutdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOutdated

`func (o *RwdNetworksRow) SetIsOutdated(v bool)`

SetIsOutdated sets IsOutdated field to given value.

### HasIsOutdated

`func (o *RwdNetworksRow) HasIsOutdated() bool`

HasIsOutdated returns a boolean if a field has been set.

### SetIsOutdatedNil

`func (o *RwdNetworksRow) SetIsOutdatedNil(b bool)`

 SetIsOutdatedNil sets the value for IsOutdated to be an explicit nil

### UnsetIsOutdated
`func (o *RwdNetworksRow) UnsetIsOutdated()`

UnsetIsOutdated ensures that no value is present for IsOutdated, not even an explicit nil
### GetNetworkRegisterTotal

`func (o *RwdNetworksRow) GetNetworkRegisterTotal() int32`

GetNetworkRegisterTotal returns the NetworkRegisterTotal field if non-nil, zero value otherwise.

### GetNetworkRegisterTotalOk

`func (o *RwdNetworksRow) GetNetworkRegisterTotalOk() (*int32, bool)`

GetNetworkRegisterTotalOk returns a tuple with the NetworkRegisterTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRegisterTotal

`func (o *RwdNetworksRow) SetNetworkRegisterTotal(v int32)`

SetNetworkRegisterTotal sets NetworkRegisterTotal field to given value.

### HasNetworkRegisterTotal

`func (o *RwdNetworksRow) HasNetworkRegisterTotal() bool`

HasNetworkRegisterTotal returns a boolean if a field has been set.

### SetNetworkRegisterTotalNil

`func (o *RwdNetworksRow) SetNetworkRegisterTotalNil(b bool)`

 SetNetworkRegisterTotalNil sets the value for NetworkRegisterTotal to be an explicit nil

### UnsetNetworkRegisterTotal
`func (o *RwdNetworksRow) UnsetNetworkRegisterTotal()`

UnsetNetworkRegisterTotal ensures that no value is present for NetworkRegisterTotal, not even an explicit nil
### GetNetworksListingAnyMember

`func (o *RwdNetworksRow) GetNetworksListingAnyMember() int32`

GetNetworksListingAnyMember returns the NetworksListingAnyMember field if non-nil, zero value otherwise.

### GetNetworksListingAnyMemberOk

`func (o *RwdNetworksRow) GetNetworksListingAnyMemberOk() (*int32, bool)`

GetNetworksListingAnyMemberOk returns a tuple with the NetworksListingAnyMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworksListingAnyMember

`func (o *RwdNetworksRow) SetNetworksListingAnyMember(v int32)`

SetNetworksListingAnyMember sets NetworksListingAnyMember field to given value.

### HasNetworksListingAnyMember

`func (o *RwdNetworksRow) HasNetworksListingAnyMember() bool`

HasNetworksListingAnyMember returns a boolean if a field has been set.

### SetNetworksListingAnyMemberNil

`func (o *RwdNetworksRow) SetNetworksListingAnyMemberNil(b bool)`

 SetNetworksListingAnyMemberNil sets the value for NetworksListingAnyMember to be an explicit nil

### UnsetNetworksListingAnyMember
`func (o *RwdNetworksRow) UnsetNetworksListingAnyMember()`

UnsetNetworksListingAnyMember ensures that no value is present for NetworksListingAnyMember, not even an explicit nil
### GetMembershipEdgeTotal

`func (o *RwdNetworksRow) GetMembershipEdgeTotal() int32`

GetMembershipEdgeTotal returns the MembershipEdgeTotal field if non-nil, zero value otherwise.

### GetMembershipEdgeTotalOk

`func (o *RwdNetworksRow) GetMembershipEdgeTotalOk() (*int32, bool)`

GetMembershipEdgeTotalOk returns a tuple with the MembershipEdgeTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipEdgeTotal

`func (o *RwdNetworksRow) SetMembershipEdgeTotal(v int32)`

SetMembershipEdgeTotal sets MembershipEdgeTotal field to given value.

### HasMembershipEdgeTotal

`func (o *RwdNetworksRow) HasMembershipEdgeTotal() bool`

HasMembershipEdgeTotal returns a boolean if a field has been set.

### SetMembershipEdgeTotalNil

`func (o *RwdNetworksRow) SetMembershipEdgeTotalNil(b bool)`

 SetMembershipEdgeTotalNil sets the value for MembershipEdgeTotal to be an explicit nil

### UnsetMembershipEdgeTotal
`func (o *RwdNetworksRow) UnsetMembershipEdgeTotal()`

UnsetMembershipEdgeTotal ensures that no value is present for MembershipEdgeTotal, not even an explicit nil
### GetInstitutionsInAnyMemberList

`func (o *RwdNetworksRow) GetInstitutionsInAnyMemberList() int32`

GetInstitutionsInAnyMemberList returns the InstitutionsInAnyMemberList field if non-nil, zero value otherwise.

### GetInstitutionsInAnyMemberListOk

`func (o *RwdNetworksRow) GetInstitutionsInAnyMemberListOk() (*int32, bool)`

GetInstitutionsInAnyMemberListOk returns a tuple with the InstitutionsInAnyMemberList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionsInAnyMemberList

`func (o *RwdNetworksRow) SetInstitutionsInAnyMemberList(v int32)`

SetInstitutionsInAnyMemberList sets InstitutionsInAnyMemberList field to given value.

### HasInstitutionsInAnyMemberList

`func (o *RwdNetworksRow) HasInstitutionsInAnyMemberList() bool`

HasInstitutionsInAnyMemberList returns a boolean if a field has been set.

### SetInstitutionsInAnyMemberListNil

`func (o *RwdNetworksRow) SetInstitutionsInAnyMemberListNil(b bool)`

 SetInstitutionsInAnyMemberListNil sets the value for InstitutionsInAnyMemberList to be an explicit nil

### UnsetInstitutionsInAnyMemberList
`func (o *RwdNetworksRow) UnsetInstitutionsInAnyMemberList()`

UnsetInstitutionsInAnyMemberList ensures that no value is present for InstitutionsInAnyMemberList, not even an explicit nil
### GetInstitutionNameKeyTotal

`func (o *RwdNetworksRow) GetInstitutionNameKeyTotal() int32`

GetInstitutionNameKeyTotal returns the InstitutionNameKeyTotal field if non-nil, zero value otherwise.

### GetInstitutionNameKeyTotalOk

`func (o *RwdNetworksRow) GetInstitutionNameKeyTotalOk() (*int32, bool)`

GetInstitutionNameKeyTotalOk returns a tuple with the InstitutionNameKeyTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionNameKeyTotal

`func (o *RwdNetworksRow) SetInstitutionNameKeyTotal(v int32)`

SetInstitutionNameKeyTotal sets InstitutionNameKeyTotal field to given value.

### HasInstitutionNameKeyTotal

`func (o *RwdNetworksRow) HasInstitutionNameKeyTotal() bool`

HasInstitutionNameKeyTotal returns a boolean if a field has been set.

### SetInstitutionNameKeyTotalNil

`func (o *RwdNetworksRow) SetInstitutionNameKeyTotalNil(b bool)`

 SetInstitutionNameKeyTotalNil sets the value for InstitutionNameKeyTotal to be an explicit nil

### UnsetInstitutionNameKeyTotal
`func (o *RwdNetworksRow) UnsetInstitutionNameKeyTotal()`

UnsetInstitutionNameKeyTotal ensures that no value is present for InstitutionNameKeyTotal, not even an explicit nil
### GetNetworksDeclaringATherapeuticArea

`func (o *RwdNetworksRow) GetNetworksDeclaringATherapeuticArea() int32`

GetNetworksDeclaringATherapeuticArea returns the NetworksDeclaringATherapeuticArea field if non-nil, zero value otherwise.

### GetNetworksDeclaringATherapeuticAreaOk

`func (o *RwdNetworksRow) GetNetworksDeclaringATherapeuticAreaOk() (*int32, bool)`

GetNetworksDeclaringATherapeuticAreaOk returns a tuple with the NetworksDeclaringATherapeuticArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworksDeclaringATherapeuticArea

`func (o *RwdNetworksRow) SetNetworksDeclaringATherapeuticArea(v int32)`

SetNetworksDeclaringATherapeuticArea sets NetworksDeclaringATherapeuticArea field to given value.

### HasNetworksDeclaringATherapeuticArea

`func (o *RwdNetworksRow) HasNetworksDeclaringATherapeuticArea() bool`

HasNetworksDeclaringATherapeuticArea returns a boolean if a field has been set.

### SetNetworksDeclaringATherapeuticAreaNil

`func (o *RwdNetworksRow) SetNetworksDeclaringATherapeuticAreaNil(b bool)`

 SetNetworksDeclaringATherapeuticAreaNil sets the value for NetworksDeclaringATherapeuticArea to be an explicit nil

### UnsetNetworksDeclaringATherapeuticArea
`func (o *RwdNetworksRow) UnsetNetworksDeclaringATherapeuticArea()`

UnsetNetworksDeclaringATherapeuticArea ensures that no value is present for NetworksDeclaringATherapeuticArea, not even an explicit nil
### GetNetworksWithAnyDeliveredStudy

`func (o *RwdNetworksRow) GetNetworksWithAnyDeliveredStudy() int32`

GetNetworksWithAnyDeliveredStudy returns the NetworksWithAnyDeliveredStudy field if non-nil, zero value otherwise.

### GetNetworksWithAnyDeliveredStudyOk

`func (o *RwdNetworksRow) GetNetworksWithAnyDeliveredStudyOk() (*int32, bool)`

GetNetworksWithAnyDeliveredStudyOk returns a tuple with the NetworksWithAnyDeliveredStudy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworksWithAnyDeliveredStudy

`func (o *RwdNetworksRow) SetNetworksWithAnyDeliveredStudy(v int32)`

SetNetworksWithAnyDeliveredStudy sets NetworksWithAnyDeliveredStudy field to given value.

### HasNetworksWithAnyDeliveredStudy

`func (o *RwdNetworksRow) HasNetworksWithAnyDeliveredStudy() bool`

HasNetworksWithAnyDeliveredStudy returns a boolean if a field has been set.

### SetNetworksWithAnyDeliveredStudyNil

`func (o *RwdNetworksRow) SetNetworksWithAnyDeliveredStudyNil(b bool)`

 SetNetworksWithAnyDeliveredStudyNil sets the value for NetworksWithAnyDeliveredStudy to be an explicit nil

### UnsetNetworksWithAnyDeliveredStudy
`func (o *RwdNetworksRow) UnsetNetworksWithAnyDeliveredStudy()`

UnsetNetworksWithAnyDeliveredStudy ensures that no value is present for NetworksWithAnyDeliveredStudy, not even an explicit nil
### GetSelfRegisteredProfileTotal

`func (o *RwdNetworksRow) GetSelfRegisteredProfileTotal() int32`

GetSelfRegisteredProfileTotal returns the SelfRegisteredProfileTotal field if non-nil, zero value otherwise.

### GetSelfRegisteredProfileTotalOk

`func (o *RwdNetworksRow) GetSelfRegisteredProfileTotalOk() (*int32, bool)`

GetSelfRegisteredProfileTotalOk returns a tuple with the SelfRegisteredProfileTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfRegisteredProfileTotal

`func (o *RwdNetworksRow) SetSelfRegisteredProfileTotal(v int32)`

SetSelfRegisteredProfileTotal sets SelfRegisteredProfileTotal field to given value.

### HasSelfRegisteredProfileTotal

`func (o *RwdNetworksRow) HasSelfRegisteredProfileTotal() bool`

HasSelfRegisteredProfileTotal returns a boolean if a field has been set.

### SetSelfRegisteredProfileTotalNil

`func (o *RwdNetworksRow) SetSelfRegisteredProfileTotalNil(b bool)`

 SetSelfRegisteredProfileTotalNil sets the value for SelfRegisteredProfileTotal to be an explicit nil

### UnsetSelfRegisteredProfileTotal
`func (o *RwdNetworksRow) UnsetSelfRegisteredProfileTotal()`

UnsetSelfRegisteredProfileTotal ensures that no value is present for SelfRegisteredProfileTotal, not even an explicit nil
### GetCatalogueNameEntryTotal

`func (o *RwdNetworksRow) GetCatalogueNameEntryTotal() int32`

GetCatalogueNameEntryTotal returns the CatalogueNameEntryTotal field if non-nil, zero value otherwise.

### GetCatalogueNameEntryTotalOk

`func (o *RwdNetworksRow) GetCatalogueNameEntryTotalOk() (*int32, bool)`

GetCatalogueNameEntryTotalOk returns a tuple with the CatalogueNameEntryTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueNameEntryTotal

`func (o *RwdNetworksRow) SetCatalogueNameEntryTotal(v int32)`

SetCatalogueNameEntryTotal sets CatalogueNameEntryTotal field to given value.

### HasCatalogueNameEntryTotal

`func (o *RwdNetworksRow) HasCatalogueNameEntryTotal() bool`

HasCatalogueNameEntryTotal returns a boolean if a field has been set.

### SetCatalogueNameEntryTotalNil

`func (o *RwdNetworksRow) SetCatalogueNameEntryTotalNil(b bool)`

 SetCatalogueNameEntryTotalNil sets the value for CatalogueNameEntryTotal to be an explicit nil

### UnsetCatalogueNameEntryTotal
`func (o *RwdNetworksRow) UnsetCatalogueNameEntryTotal()`

UnsetCatalogueNameEntryTotal ensures that no value is present for CatalogueNameEntryTotal, not even an explicit nil
### GetSelfRegisteredProfileEnceppYesTotal

`func (o *RwdNetworksRow) GetSelfRegisteredProfileEnceppYesTotal() int32`

GetSelfRegisteredProfileEnceppYesTotal returns the SelfRegisteredProfileEnceppYesTotal field if non-nil, zero value otherwise.

### GetSelfRegisteredProfileEnceppYesTotalOk

`func (o *RwdNetworksRow) GetSelfRegisteredProfileEnceppYesTotalOk() (*int32, bool)`

GetSelfRegisteredProfileEnceppYesTotalOk returns a tuple with the SelfRegisteredProfileEnceppYesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfRegisteredProfileEnceppYesTotal

`func (o *RwdNetworksRow) SetSelfRegisteredProfileEnceppYesTotal(v int32)`

SetSelfRegisteredProfileEnceppYesTotal sets SelfRegisteredProfileEnceppYesTotal field to given value.

### HasSelfRegisteredProfileEnceppYesTotal

`func (o *RwdNetworksRow) HasSelfRegisteredProfileEnceppYesTotal() bool`

HasSelfRegisteredProfileEnceppYesTotal returns a boolean if a field has been set.

### SetSelfRegisteredProfileEnceppYesTotalNil

`func (o *RwdNetworksRow) SetSelfRegisteredProfileEnceppYesTotalNil(b bool)`

 SetSelfRegisteredProfileEnceppYesTotalNil sets the value for SelfRegisteredProfileEnceppYesTotal to be an explicit nil

### UnsetSelfRegisteredProfileEnceppYesTotal
`func (o *RwdNetworksRow) UnsetSelfRegisteredProfileEnceppYesTotal()`

UnsetSelfRegisteredProfileEnceppYesTotal ensures that no value is present for SelfRegisteredProfileEnceppYesTotal, not even an explicit nil
### GetCatalogueNameEntryEnceppYesTotal

`func (o *RwdNetworksRow) GetCatalogueNameEntryEnceppYesTotal() int32`

GetCatalogueNameEntryEnceppYesTotal returns the CatalogueNameEntryEnceppYesTotal field if non-nil, zero value otherwise.

### GetCatalogueNameEntryEnceppYesTotalOk

`func (o *RwdNetworksRow) GetCatalogueNameEntryEnceppYesTotalOk() (*int32, bool)`

GetCatalogueNameEntryEnceppYesTotalOk returns a tuple with the CatalogueNameEntryEnceppYesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueNameEntryEnceppYesTotal

`func (o *RwdNetworksRow) SetCatalogueNameEntryEnceppYesTotal(v int32)`

SetCatalogueNameEntryEnceppYesTotal sets CatalogueNameEntryEnceppYesTotal field to given value.

### HasCatalogueNameEntryEnceppYesTotal

`func (o *RwdNetworksRow) HasCatalogueNameEntryEnceppYesTotal() bool`

HasCatalogueNameEntryEnceppYesTotal returns a boolean if a field has been set.

### SetCatalogueNameEntryEnceppYesTotalNil

`func (o *RwdNetworksRow) SetCatalogueNameEntryEnceppYesTotalNil(b bool)`

 SetCatalogueNameEntryEnceppYesTotalNil sets the value for CatalogueNameEntryEnceppYesTotal to be an explicit nil

### UnsetCatalogueNameEntryEnceppYesTotal
`func (o *RwdNetworksRow) UnsetCatalogueNameEntryEnceppYesTotal()`

UnsetCatalogueNameEntryEnceppYesTotal ensures that no value is present for CatalogueNameEntryEnceppYesTotal, not even an explicit nil
### GetRegisterSourceCount

`func (o *RwdNetworksRow) GetRegisterSourceCount() int32`

GetRegisterSourceCount returns the RegisterSourceCount field if non-nil, zero value otherwise.

### GetRegisterSourceCountOk

`func (o *RwdNetworksRow) GetRegisterSourceCountOk() (*int32, bool)`

GetRegisterSourceCountOk returns a tuple with the RegisterSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterSourceCount

`func (o *RwdNetworksRow) SetRegisterSourceCount(v int32)`

SetRegisterSourceCount sets RegisterSourceCount field to given value.

### HasRegisterSourceCount

`func (o *RwdNetworksRow) HasRegisterSourceCount() bool`

HasRegisterSourceCount returns a boolean if a field has been set.

### SetRegisterSourceCountNil

`func (o *RwdNetworksRow) SetRegisterSourceCountNil(b bool)`

 SetRegisterSourceCountNil sets the value for RegisterSourceCount to be an explicit nil

### UnsetRegisterSourceCount
`func (o *RwdNetworksRow) UnsetRegisterSourceCount()`

UnsetRegisterSourceCount ensures that no value is present for RegisterSourceCount, not even an explicit nil
### GetStudyRegisterTotal

`func (o *RwdNetworksRow) GetStudyRegisterTotal() int32`

GetStudyRegisterTotal returns the StudyRegisterTotal field if non-nil, zero value otherwise.

### GetStudyRegisterTotalOk

`func (o *RwdNetworksRow) GetStudyRegisterTotalOk() (*int32, bool)`

GetStudyRegisterTotalOk returns a tuple with the StudyRegisterTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyRegisterTotal

`func (o *RwdNetworksRow) SetStudyRegisterTotal(v int32)`

SetStudyRegisterTotal sets StudyRegisterTotal field to given value.

### HasStudyRegisterTotal

`func (o *RwdNetworksRow) HasStudyRegisterTotal() bool`

HasStudyRegisterTotal returns a boolean if a field has been set.

### SetStudyRegisterTotalNil

`func (o *RwdNetworksRow) SetStudyRegisterTotalNil(b bool)`

 SetStudyRegisterTotalNil sets the value for StudyRegisterTotal to be an explicit nil

### UnsetStudyRegisterTotal
`func (o *RwdNetworksRow) UnsetStudyRegisterTotal()`

UnsetStudyRegisterTotal ensures that no value is present for StudyRegisterTotal, not even an explicit nil
### GetStudiesNamingARegisteredSource

`func (o *RwdNetworksRow) GetStudiesNamingARegisteredSource() int32`

GetStudiesNamingARegisteredSource returns the StudiesNamingARegisteredSource field if non-nil, zero value otherwise.

### GetStudiesNamingARegisteredSourceOk

`func (o *RwdNetworksRow) GetStudiesNamingARegisteredSourceOk() (*int32, bool)`

GetStudiesNamingARegisteredSourceOk returns a tuple with the StudiesNamingARegisteredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudiesNamingARegisteredSource

`func (o *RwdNetworksRow) SetStudiesNamingARegisteredSource(v int32)`

SetStudiesNamingARegisteredSource sets StudiesNamingARegisteredSource field to given value.

### HasStudiesNamingARegisteredSource

`func (o *RwdNetworksRow) HasStudiesNamingARegisteredSource() bool`

HasStudiesNamingARegisteredSource returns a boolean if a field has been set.

### SetStudiesNamingARegisteredSourceNil

`func (o *RwdNetworksRow) SetStudiesNamingARegisteredSourceNil(b bool)`

 SetStudiesNamingARegisteredSourceNil sets the value for StudiesNamingARegisteredSource to be an explicit nil

### UnsetStudiesNamingARegisteredSource
`func (o *RwdNetworksRow) UnsetStudiesNamingARegisteredSource()`

UnsetStudiesNamingARegisteredSource ensures that no value is present for StudiesNamingARegisteredSource, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdNetworksRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdNetworksRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdNetworksRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdNetworksRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdNetworksRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdNetworksRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *RwdNetworksRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdNetworksRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdNetworksRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdNetworksRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdNetworksRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdNetworksRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdNetworksRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdNetworksRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdNetworksRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdNetworksRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdNetworksRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdNetworksRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdNetworksRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdNetworksRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdNetworksRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdNetworksRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdNetworksRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdNetworksRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


