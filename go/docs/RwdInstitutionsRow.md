# RwdInstitutionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstitutionId** | Pointer to **NullableString** | The register&#39;s own Institution ID. The entry key of the rwd_institutions object type. The register can list one name under two ids; counts are per id and are not additive across rows. | [optional] 
**InstitutionName** | Pointer to **NullableString** | The institution&#39;s name as the register spells it; an institution here can be registered under a person&#39;s name. | [optional] 
**InstitutionOrgId** | Pointer to **NullableString** | The legal entity (organization record id) the institution resolved to; null when it did not resolve. An unresolved institution keeps its row with null ids. | [optional] 
**InstitutionCompanyId** | Pointer to **NullableString** | The company the institution&#39;s legal entity belongs to; null when unresolved. | [optional] 
**InstitutionCompanyName** | Pointer to **NullableString** | The company&#39;s name. | [optional] 
**InstitutionGroupId** | Pointer to **NullableString** | The current ultimate parent company of the institution&#39;s company. | [optional] 
**InstitutionGroupName** | Pointer to **NullableString** | The group&#39;s name. | [optional] 
**OrgMatchRung** | Pointer to **NullableString** | How the institution name resolved to a legal entity (the resolution rung); null when unresolved. | [optional] 
**OrgMatchTier** | Pointer to **NullableString** | The confidence tier of that resolution; null when unresolved. | [optional] 
**ParentOrgId** | Pointer to **NullableString** | The legal entity the register&#39;s parent institution claim resolved to. Most parent claims are claims between register entries, not between organizations: read it with is_org_to_org. | [optional] 
**IsOrgToOrg** | Pointer to **NullableBool** | True when the parent claim is a claim between two resolved legal entities, false when it only links register entries. | [optional] 
**InstitutionNameKey** | Pointer to **NullableString** | The normalised name key the study, membership and holder edges join on. | [optional] 
**ParentName** | Pointer to **NullableString** | The parent institution&#39;s name as the register spells it. | [optional] 
**ParentNameKey** | Pointer to **NullableString** | The normalised key of the parent name. | [optional] 
**RegistrationState** | Pointer to **NullableString** | self_registered_profile when the entry carries at least one descriptive field, catalogue_name_entry when it is a name and an id only (a catalogue pick-list entry nobody registered a profile for). The declaration columns are null on a catalogue name entry. | [optional] 
**IsSelfRegisteredProfile** | Pointer to **NullableBool** | True when registration_state is self_registered_profile. | [optional] 
**DeclaredProfileFields** | Pointer to **NullableInt32** | How many of the declarable profile fields the entry answered. | [optional] 
**DeclarableProfileFields** | Pointer to **NullableInt32** | How many profile fields the register offers; the denominator of declared_profile_fields. | [optional] 
**TypeOfInstitution** | Pointer to **NullableString** | The type of institution the entry declares, as typed. | [optional] 
**InstitutionRole** | Pointer to **NullableString** | The role the entry declares for the institution, as typed. | [optional] 
**InstitutionCountries** | Pointer to **NullableString** | The country text the entry declares, as typed. | [optional] 
**InstitutionCountryNames** | Pointer to **[]string** | The declared countries split by the register&#39;s country list, one element per country. | [optional] 
**InstitutionCountryCount** | Pointer to **NullableInt32** | How many countries institution_country_names holds. | [optional] 
**UnmatchedCountryFragments** | Pointer to **[]string** | Country text items the country vocabulary of this build does not hold. | [optional] 
**SingleDeclaredCountryIso2** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 code when the entry declares exactly one country, null otherwise. | [optional] 
**InstitutionWebsite** | Pointer to **NullableString** | The institution&#39;s website as the register holds it. | [optional] 
**CatalogueUrl** | Pointer to **NullableString** | The address of this entry on the HMA-EMA catalogue site (the register&#39;s URL LINK). | [optional] 
**OpenToIndustryFundedResearch** | Pointer to **NullableBool** | Whether the institution declares it is open to industry-funded research; null on a catalogue name entry, where the export writes a default No that is not an answer. | [optional] 
**CollectsDataDirectlyFromPatients** | Pointer to **NullableBool** | Whether the institution declares it collects data directly from patients; null on a catalogue name entry. | [optional] 
**IsEnceppPartner** | Pointer to **NullableBool** | Whether the entry declares an ENCePP partnership; null on a catalogue name entry, where the export writes a default No that is not an answer. | [optional] 
**RecordUpdatedAt** | Pointer to **NullableTime** | When the registrant last edited this catalogue entry. A different clock from last_refreshed_on: it dates the web form, not the database. | [optional] 
**RecordUpdatedIsMassEdit** | Pointer to **NullableBool** | True when the record&#39;s last update date is a mass-edit date: one shared by an outsized part of the register, so it dates a migration and not the registrant. | [optional] 
**RecordUpdatedAtAttested** | Pointer to **NullableTime** | record_updated_at with the mass-edit dates removed (null for a mass-edit date or no timestamp): the column to sort a freshness question on. | [optional] 
**RecordUpdatedDateShare** | Pointer to **NullableFloat32** | The share of the register whose last update carries the same date, on every row so a mass-edit date is visible. | [optional] 
**RecordFirstPublishedAt** | Pointer to **NullableTime** | When the catalogue entry was first published. | [optional] 
**IsOutdated** | Pointer to **NullableBool** | EMA&#39;s own Outdated flag, carried through untouched; the register flags none of the data sources, so read false as EMA&#39;s answer and not as a currency finding. | [optional] 
**ConductingStudiesLifetime** | Pointer to **NullableInt32** | Distinct registered studies the institution&#39;s name conducted. Four separate activity numbers; never added together. | [optional] 
**ConductingStudiesLast5Years** | Pointer to **NullableInt32** | Conducting studies that started within five years of the snapshot. | [optional] 
**ConductingStudiesOngoing** | Pointer to **NullableInt32** | Conducting studies whose status is Ongoing. | [optional] 
**ConductingStudiesPlanned** | Pointer to **NullableInt32** | Conducting studies whose status is Planned. | [optional] 
**AdditionalRoleStudies** | Pointer to **NullableInt32** | Distinct registered studies the conductor is named on as an additional institution; a different claim from conducting, never added to it. | [optional] 
**MostRecentStudyStartOn** | Pointer to **NullableString** | The latest actual start date among the studies it conducted (or delivered, for a network). | [optional] 
**HeldSourceCount** | Pointer to **NullableInt32** | How many registered data sources name this institution as their data holder; zero is no source naming it, not a finding about the institution. | [optional] 
**HeldSourceIds** | Pointer to **[]string** | The register Data source IDs of those sources. | [optional] 
**HeldActivePopulation** | Pointer to **NullableInt64** | The sum of the active populations the held sources declare. | [optional] 
**HeldMaxActivePopulation** | Pointer to **NullableInt64** | The largest active population any held source declares. | [optional] 
**HeldSourcesDeclaringActivePopulation** | Pointer to **NullableInt32** | How many held sources declare an active population. | [optional] 
**HeldQualifiedSourceCount** | Pointer to **NullableInt32** | How many held sources declare a data source qualification. | [optional] 
**HeldDarwinPartnerSourceCount** | Pointer to **NullableInt32** | How many held sources are DARWIN EU data partners. | [optional] 
**NetworkMembershipCount** | Pointer to **NullableInt32** | How many networks&#39; own member lists name this institution; zero is the network&#39;s silence, not a property of the institution. | [optional] 
**NetworkIds** | Pointer to **[]string** | The register Network IDs of those networks. | [optional] 
**NetworkNames** | Pointer to **[]string** | The names of those networks. | [optional] 
**InstitutionRegisterTotal** | Pointer to **NullableInt32** | The number of institutions in the whole institution register at the snapshot. | [optional] 
**SelfRegisteredProfileTotal** | Pointer to **NullableInt32** | How many register entries are self-registered profiles. | [optional] 
**CatalogueNameEntryTotal** | Pointer to **NullableInt32** | How many register entries are catalogue name entries. | [optional] 
**InstitutionsDeclaringACountry** | Pointer to **NullableInt32** | How many register institutions declare a country. | [optional] 
**InstitutionsInAnyMemberList** | Pointer to **NullableInt32** | How many register institutions at least one network member list names. | [optional] 
**InstitutionNameKeyTotal** | Pointer to **NullableInt32** | How many distinct institution name keys the register holds. | [optional] 
**InstitutionsHoldingASource** | Pointer to **NullableInt32** | How many register institutions hold a registered data source. | [optional] 
**DeclaredHolderNameTotal** | Pointer to **NullableInt32** | How many distinct data holder names the data source register declares. | [optional] 
**DeclaredHolderNameMatched** | Pointer to **NullableInt32** | How many of those holder names match a register institution. | [optional] 
**RegisterSourceCount** | Pointer to **NullableInt32** | The number of data sources in the whole data source register at the snapshot: the denominator every source count is over. Not a count of real-world data sources in the world. | [optional] 
**StudyRegisterTotal** | Pointer to **NullableInt32** | The number of studies in the whole study register at the snapshot. | [optional] 
**StudiesNamingARegisteredSource** | Pointer to **NullableInt32** | How many registered studies name at least one data source that is itself registered; the denominator under study_count. | [optional] 
**StudiesNamingARegisteredConductor** | Pointer to **NullableInt32** | How many registered studies name a conducting institution that is itself registered. | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdInstitutionsRow

`func NewRwdInstitutionsRow() *RwdInstitutionsRow`

NewRwdInstitutionsRow instantiates a new RwdInstitutionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdInstitutionsRowWithDefaults

`func NewRwdInstitutionsRowWithDefaults() *RwdInstitutionsRow`

NewRwdInstitutionsRowWithDefaults instantiates a new RwdInstitutionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstitutionId

`func (o *RwdInstitutionsRow) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *RwdInstitutionsRow) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *RwdInstitutionsRow) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.

### HasInstitutionId

`func (o *RwdInstitutionsRow) HasInstitutionId() bool`

HasInstitutionId returns a boolean if a field has been set.

### SetInstitutionIdNil

`func (o *RwdInstitutionsRow) SetInstitutionIdNil(b bool)`

 SetInstitutionIdNil sets the value for InstitutionId to be an explicit nil

### UnsetInstitutionId
`func (o *RwdInstitutionsRow) UnsetInstitutionId()`

UnsetInstitutionId ensures that no value is present for InstitutionId, not even an explicit nil
### GetInstitutionName

`func (o *RwdInstitutionsRow) GetInstitutionName() string`

GetInstitutionName returns the InstitutionName field if non-nil, zero value otherwise.

### GetInstitutionNameOk

`func (o *RwdInstitutionsRow) GetInstitutionNameOk() (*string, bool)`

GetInstitutionNameOk returns a tuple with the InstitutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionName

`func (o *RwdInstitutionsRow) SetInstitutionName(v string)`

SetInstitutionName sets InstitutionName field to given value.

### HasInstitutionName

`func (o *RwdInstitutionsRow) HasInstitutionName() bool`

HasInstitutionName returns a boolean if a field has been set.

### SetInstitutionNameNil

`func (o *RwdInstitutionsRow) SetInstitutionNameNil(b bool)`

 SetInstitutionNameNil sets the value for InstitutionName to be an explicit nil

### UnsetInstitutionName
`func (o *RwdInstitutionsRow) UnsetInstitutionName()`

UnsetInstitutionName ensures that no value is present for InstitutionName, not even an explicit nil
### GetInstitutionOrgId

`func (o *RwdInstitutionsRow) GetInstitutionOrgId() string`

GetInstitutionOrgId returns the InstitutionOrgId field if non-nil, zero value otherwise.

### GetInstitutionOrgIdOk

`func (o *RwdInstitutionsRow) GetInstitutionOrgIdOk() (*string, bool)`

GetInstitutionOrgIdOk returns a tuple with the InstitutionOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionOrgId

`func (o *RwdInstitutionsRow) SetInstitutionOrgId(v string)`

SetInstitutionOrgId sets InstitutionOrgId field to given value.

### HasInstitutionOrgId

`func (o *RwdInstitutionsRow) HasInstitutionOrgId() bool`

HasInstitutionOrgId returns a boolean if a field has been set.

### SetInstitutionOrgIdNil

`func (o *RwdInstitutionsRow) SetInstitutionOrgIdNil(b bool)`

 SetInstitutionOrgIdNil sets the value for InstitutionOrgId to be an explicit nil

### UnsetInstitutionOrgId
`func (o *RwdInstitutionsRow) UnsetInstitutionOrgId()`

UnsetInstitutionOrgId ensures that no value is present for InstitutionOrgId, not even an explicit nil
### GetInstitutionCompanyId

`func (o *RwdInstitutionsRow) GetInstitutionCompanyId() string`

GetInstitutionCompanyId returns the InstitutionCompanyId field if non-nil, zero value otherwise.

### GetInstitutionCompanyIdOk

`func (o *RwdInstitutionsRow) GetInstitutionCompanyIdOk() (*string, bool)`

GetInstitutionCompanyIdOk returns a tuple with the InstitutionCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionCompanyId

`func (o *RwdInstitutionsRow) SetInstitutionCompanyId(v string)`

SetInstitutionCompanyId sets InstitutionCompanyId field to given value.

### HasInstitutionCompanyId

`func (o *RwdInstitutionsRow) HasInstitutionCompanyId() bool`

HasInstitutionCompanyId returns a boolean if a field has been set.

### SetInstitutionCompanyIdNil

`func (o *RwdInstitutionsRow) SetInstitutionCompanyIdNil(b bool)`

 SetInstitutionCompanyIdNil sets the value for InstitutionCompanyId to be an explicit nil

### UnsetInstitutionCompanyId
`func (o *RwdInstitutionsRow) UnsetInstitutionCompanyId()`

UnsetInstitutionCompanyId ensures that no value is present for InstitutionCompanyId, not even an explicit nil
### GetInstitutionCompanyName

`func (o *RwdInstitutionsRow) GetInstitutionCompanyName() string`

GetInstitutionCompanyName returns the InstitutionCompanyName field if non-nil, zero value otherwise.

### GetInstitutionCompanyNameOk

`func (o *RwdInstitutionsRow) GetInstitutionCompanyNameOk() (*string, bool)`

GetInstitutionCompanyNameOk returns a tuple with the InstitutionCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionCompanyName

`func (o *RwdInstitutionsRow) SetInstitutionCompanyName(v string)`

SetInstitutionCompanyName sets InstitutionCompanyName field to given value.

### HasInstitutionCompanyName

`func (o *RwdInstitutionsRow) HasInstitutionCompanyName() bool`

HasInstitutionCompanyName returns a boolean if a field has been set.

### SetInstitutionCompanyNameNil

`func (o *RwdInstitutionsRow) SetInstitutionCompanyNameNil(b bool)`

 SetInstitutionCompanyNameNil sets the value for InstitutionCompanyName to be an explicit nil

### UnsetInstitutionCompanyName
`func (o *RwdInstitutionsRow) UnsetInstitutionCompanyName()`

UnsetInstitutionCompanyName ensures that no value is present for InstitutionCompanyName, not even an explicit nil
### GetInstitutionGroupId

`func (o *RwdInstitutionsRow) GetInstitutionGroupId() string`

GetInstitutionGroupId returns the InstitutionGroupId field if non-nil, zero value otherwise.

### GetInstitutionGroupIdOk

`func (o *RwdInstitutionsRow) GetInstitutionGroupIdOk() (*string, bool)`

GetInstitutionGroupIdOk returns a tuple with the InstitutionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionGroupId

`func (o *RwdInstitutionsRow) SetInstitutionGroupId(v string)`

SetInstitutionGroupId sets InstitutionGroupId field to given value.

### HasInstitutionGroupId

`func (o *RwdInstitutionsRow) HasInstitutionGroupId() bool`

HasInstitutionGroupId returns a boolean if a field has been set.

### SetInstitutionGroupIdNil

`func (o *RwdInstitutionsRow) SetInstitutionGroupIdNil(b bool)`

 SetInstitutionGroupIdNil sets the value for InstitutionGroupId to be an explicit nil

### UnsetInstitutionGroupId
`func (o *RwdInstitutionsRow) UnsetInstitutionGroupId()`

UnsetInstitutionGroupId ensures that no value is present for InstitutionGroupId, not even an explicit nil
### GetInstitutionGroupName

`func (o *RwdInstitutionsRow) GetInstitutionGroupName() string`

GetInstitutionGroupName returns the InstitutionGroupName field if non-nil, zero value otherwise.

### GetInstitutionGroupNameOk

`func (o *RwdInstitutionsRow) GetInstitutionGroupNameOk() (*string, bool)`

GetInstitutionGroupNameOk returns a tuple with the InstitutionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionGroupName

`func (o *RwdInstitutionsRow) SetInstitutionGroupName(v string)`

SetInstitutionGroupName sets InstitutionGroupName field to given value.

### HasInstitutionGroupName

`func (o *RwdInstitutionsRow) HasInstitutionGroupName() bool`

HasInstitutionGroupName returns a boolean if a field has been set.

### SetInstitutionGroupNameNil

`func (o *RwdInstitutionsRow) SetInstitutionGroupNameNil(b bool)`

 SetInstitutionGroupNameNil sets the value for InstitutionGroupName to be an explicit nil

### UnsetInstitutionGroupName
`func (o *RwdInstitutionsRow) UnsetInstitutionGroupName()`

UnsetInstitutionGroupName ensures that no value is present for InstitutionGroupName, not even an explicit nil
### GetOrgMatchRung

`func (o *RwdInstitutionsRow) GetOrgMatchRung() string`

GetOrgMatchRung returns the OrgMatchRung field if non-nil, zero value otherwise.

### GetOrgMatchRungOk

`func (o *RwdInstitutionsRow) GetOrgMatchRungOk() (*string, bool)`

GetOrgMatchRungOk returns a tuple with the OrgMatchRung field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMatchRung

`func (o *RwdInstitutionsRow) SetOrgMatchRung(v string)`

SetOrgMatchRung sets OrgMatchRung field to given value.

### HasOrgMatchRung

`func (o *RwdInstitutionsRow) HasOrgMatchRung() bool`

HasOrgMatchRung returns a boolean if a field has been set.

### SetOrgMatchRungNil

`func (o *RwdInstitutionsRow) SetOrgMatchRungNil(b bool)`

 SetOrgMatchRungNil sets the value for OrgMatchRung to be an explicit nil

### UnsetOrgMatchRung
`func (o *RwdInstitutionsRow) UnsetOrgMatchRung()`

UnsetOrgMatchRung ensures that no value is present for OrgMatchRung, not even an explicit nil
### GetOrgMatchTier

`func (o *RwdInstitutionsRow) GetOrgMatchTier() string`

GetOrgMatchTier returns the OrgMatchTier field if non-nil, zero value otherwise.

### GetOrgMatchTierOk

`func (o *RwdInstitutionsRow) GetOrgMatchTierOk() (*string, bool)`

GetOrgMatchTierOk returns a tuple with the OrgMatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMatchTier

`func (o *RwdInstitutionsRow) SetOrgMatchTier(v string)`

SetOrgMatchTier sets OrgMatchTier field to given value.

### HasOrgMatchTier

`func (o *RwdInstitutionsRow) HasOrgMatchTier() bool`

HasOrgMatchTier returns a boolean if a field has been set.

### SetOrgMatchTierNil

`func (o *RwdInstitutionsRow) SetOrgMatchTierNil(b bool)`

 SetOrgMatchTierNil sets the value for OrgMatchTier to be an explicit nil

### UnsetOrgMatchTier
`func (o *RwdInstitutionsRow) UnsetOrgMatchTier()`

UnsetOrgMatchTier ensures that no value is present for OrgMatchTier, not even an explicit nil
### GetParentOrgId

`func (o *RwdInstitutionsRow) GetParentOrgId() string`

GetParentOrgId returns the ParentOrgId field if non-nil, zero value otherwise.

### GetParentOrgIdOk

`func (o *RwdInstitutionsRow) GetParentOrgIdOk() (*string, bool)`

GetParentOrgIdOk returns a tuple with the ParentOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrgId

`func (o *RwdInstitutionsRow) SetParentOrgId(v string)`

SetParentOrgId sets ParentOrgId field to given value.

### HasParentOrgId

`func (o *RwdInstitutionsRow) HasParentOrgId() bool`

HasParentOrgId returns a boolean if a field has been set.

### SetParentOrgIdNil

`func (o *RwdInstitutionsRow) SetParentOrgIdNil(b bool)`

 SetParentOrgIdNil sets the value for ParentOrgId to be an explicit nil

### UnsetParentOrgId
`func (o *RwdInstitutionsRow) UnsetParentOrgId()`

UnsetParentOrgId ensures that no value is present for ParentOrgId, not even an explicit nil
### GetIsOrgToOrg

`func (o *RwdInstitutionsRow) GetIsOrgToOrg() bool`

GetIsOrgToOrg returns the IsOrgToOrg field if non-nil, zero value otherwise.

### GetIsOrgToOrgOk

`func (o *RwdInstitutionsRow) GetIsOrgToOrgOk() (*bool, bool)`

GetIsOrgToOrgOk returns a tuple with the IsOrgToOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgToOrg

`func (o *RwdInstitutionsRow) SetIsOrgToOrg(v bool)`

SetIsOrgToOrg sets IsOrgToOrg field to given value.

### HasIsOrgToOrg

`func (o *RwdInstitutionsRow) HasIsOrgToOrg() bool`

HasIsOrgToOrg returns a boolean if a field has been set.

### SetIsOrgToOrgNil

`func (o *RwdInstitutionsRow) SetIsOrgToOrgNil(b bool)`

 SetIsOrgToOrgNil sets the value for IsOrgToOrg to be an explicit nil

### UnsetIsOrgToOrg
`func (o *RwdInstitutionsRow) UnsetIsOrgToOrg()`

UnsetIsOrgToOrg ensures that no value is present for IsOrgToOrg, not even an explicit nil
### GetInstitutionNameKey

`func (o *RwdInstitutionsRow) GetInstitutionNameKey() string`

GetInstitutionNameKey returns the InstitutionNameKey field if non-nil, zero value otherwise.

### GetInstitutionNameKeyOk

`func (o *RwdInstitutionsRow) GetInstitutionNameKeyOk() (*string, bool)`

GetInstitutionNameKeyOk returns a tuple with the InstitutionNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionNameKey

`func (o *RwdInstitutionsRow) SetInstitutionNameKey(v string)`

SetInstitutionNameKey sets InstitutionNameKey field to given value.

### HasInstitutionNameKey

`func (o *RwdInstitutionsRow) HasInstitutionNameKey() bool`

HasInstitutionNameKey returns a boolean if a field has been set.

### SetInstitutionNameKeyNil

`func (o *RwdInstitutionsRow) SetInstitutionNameKeyNil(b bool)`

 SetInstitutionNameKeyNil sets the value for InstitutionNameKey to be an explicit nil

### UnsetInstitutionNameKey
`func (o *RwdInstitutionsRow) UnsetInstitutionNameKey()`

UnsetInstitutionNameKey ensures that no value is present for InstitutionNameKey, not even an explicit nil
### GetParentName

`func (o *RwdInstitutionsRow) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *RwdInstitutionsRow) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *RwdInstitutionsRow) SetParentName(v string)`

SetParentName sets ParentName field to given value.

### HasParentName

`func (o *RwdInstitutionsRow) HasParentName() bool`

HasParentName returns a boolean if a field has been set.

### SetParentNameNil

`func (o *RwdInstitutionsRow) SetParentNameNil(b bool)`

 SetParentNameNil sets the value for ParentName to be an explicit nil

### UnsetParentName
`func (o *RwdInstitutionsRow) UnsetParentName()`

UnsetParentName ensures that no value is present for ParentName, not even an explicit nil
### GetParentNameKey

`func (o *RwdInstitutionsRow) GetParentNameKey() string`

GetParentNameKey returns the ParentNameKey field if non-nil, zero value otherwise.

### GetParentNameKeyOk

`func (o *RwdInstitutionsRow) GetParentNameKeyOk() (*string, bool)`

GetParentNameKeyOk returns a tuple with the ParentNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNameKey

`func (o *RwdInstitutionsRow) SetParentNameKey(v string)`

SetParentNameKey sets ParentNameKey field to given value.

### HasParentNameKey

`func (o *RwdInstitutionsRow) HasParentNameKey() bool`

HasParentNameKey returns a boolean if a field has been set.

### SetParentNameKeyNil

`func (o *RwdInstitutionsRow) SetParentNameKeyNil(b bool)`

 SetParentNameKeyNil sets the value for ParentNameKey to be an explicit nil

### UnsetParentNameKey
`func (o *RwdInstitutionsRow) UnsetParentNameKey()`

UnsetParentNameKey ensures that no value is present for ParentNameKey, not even an explicit nil
### GetRegistrationState

`func (o *RwdInstitutionsRow) GetRegistrationState() string`

GetRegistrationState returns the RegistrationState field if non-nil, zero value otherwise.

### GetRegistrationStateOk

`func (o *RwdInstitutionsRow) GetRegistrationStateOk() (*string, bool)`

GetRegistrationStateOk returns a tuple with the RegistrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationState

`func (o *RwdInstitutionsRow) SetRegistrationState(v string)`

SetRegistrationState sets RegistrationState field to given value.

### HasRegistrationState

`func (o *RwdInstitutionsRow) HasRegistrationState() bool`

HasRegistrationState returns a boolean if a field has been set.

### SetRegistrationStateNil

`func (o *RwdInstitutionsRow) SetRegistrationStateNil(b bool)`

 SetRegistrationStateNil sets the value for RegistrationState to be an explicit nil

### UnsetRegistrationState
`func (o *RwdInstitutionsRow) UnsetRegistrationState()`

UnsetRegistrationState ensures that no value is present for RegistrationState, not even an explicit nil
### GetIsSelfRegisteredProfile

`func (o *RwdInstitutionsRow) GetIsSelfRegisteredProfile() bool`

GetIsSelfRegisteredProfile returns the IsSelfRegisteredProfile field if non-nil, zero value otherwise.

### GetIsSelfRegisteredProfileOk

`func (o *RwdInstitutionsRow) GetIsSelfRegisteredProfileOk() (*bool, bool)`

GetIsSelfRegisteredProfileOk returns a tuple with the IsSelfRegisteredProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfRegisteredProfile

`func (o *RwdInstitutionsRow) SetIsSelfRegisteredProfile(v bool)`

SetIsSelfRegisteredProfile sets IsSelfRegisteredProfile field to given value.

### HasIsSelfRegisteredProfile

`func (o *RwdInstitutionsRow) HasIsSelfRegisteredProfile() bool`

HasIsSelfRegisteredProfile returns a boolean if a field has been set.

### SetIsSelfRegisteredProfileNil

`func (o *RwdInstitutionsRow) SetIsSelfRegisteredProfileNil(b bool)`

 SetIsSelfRegisteredProfileNil sets the value for IsSelfRegisteredProfile to be an explicit nil

### UnsetIsSelfRegisteredProfile
`func (o *RwdInstitutionsRow) UnsetIsSelfRegisteredProfile()`

UnsetIsSelfRegisteredProfile ensures that no value is present for IsSelfRegisteredProfile, not even an explicit nil
### GetDeclaredProfileFields

`func (o *RwdInstitutionsRow) GetDeclaredProfileFields() int32`

GetDeclaredProfileFields returns the DeclaredProfileFields field if non-nil, zero value otherwise.

### GetDeclaredProfileFieldsOk

`func (o *RwdInstitutionsRow) GetDeclaredProfileFieldsOk() (*int32, bool)`

GetDeclaredProfileFieldsOk returns a tuple with the DeclaredProfileFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredProfileFields

`func (o *RwdInstitutionsRow) SetDeclaredProfileFields(v int32)`

SetDeclaredProfileFields sets DeclaredProfileFields field to given value.

### HasDeclaredProfileFields

`func (o *RwdInstitutionsRow) HasDeclaredProfileFields() bool`

HasDeclaredProfileFields returns a boolean if a field has been set.

### SetDeclaredProfileFieldsNil

`func (o *RwdInstitutionsRow) SetDeclaredProfileFieldsNil(b bool)`

 SetDeclaredProfileFieldsNil sets the value for DeclaredProfileFields to be an explicit nil

### UnsetDeclaredProfileFields
`func (o *RwdInstitutionsRow) UnsetDeclaredProfileFields()`

UnsetDeclaredProfileFields ensures that no value is present for DeclaredProfileFields, not even an explicit nil
### GetDeclarableProfileFields

`func (o *RwdInstitutionsRow) GetDeclarableProfileFields() int32`

GetDeclarableProfileFields returns the DeclarableProfileFields field if non-nil, zero value otherwise.

### GetDeclarableProfileFieldsOk

`func (o *RwdInstitutionsRow) GetDeclarableProfileFieldsOk() (*int32, bool)`

GetDeclarableProfileFieldsOk returns a tuple with the DeclarableProfileFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarableProfileFields

`func (o *RwdInstitutionsRow) SetDeclarableProfileFields(v int32)`

SetDeclarableProfileFields sets DeclarableProfileFields field to given value.

### HasDeclarableProfileFields

`func (o *RwdInstitutionsRow) HasDeclarableProfileFields() bool`

HasDeclarableProfileFields returns a boolean if a field has been set.

### SetDeclarableProfileFieldsNil

`func (o *RwdInstitutionsRow) SetDeclarableProfileFieldsNil(b bool)`

 SetDeclarableProfileFieldsNil sets the value for DeclarableProfileFields to be an explicit nil

### UnsetDeclarableProfileFields
`func (o *RwdInstitutionsRow) UnsetDeclarableProfileFields()`

UnsetDeclarableProfileFields ensures that no value is present for DeclarableProfileFields, not even an explicit nil
### GetTypeOfInstitution

`func (o *RwdInstitutionsRow) GetTypeOfInstitution() string`

GetTypeOfInstitution returns the TypeOfInstitution field if non-nil, zero value otherwise.

### GetTypeOfInstitutionOk

`func (o *RwdInstitutionsRow) GetTypeOfInstitutionOk() (*string, bool)`

GetTypeOfInstitutionOk returns a tuple with the TypeOfInstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfInstitution

`func (o *RwdInstitutionsRow) SetTypeOfInstitution(v string)`

SetTypeOfInstitution sets TypeOfInstitution field to given value.

### HasTypeOfInstitution

`func (o *RwdInstitutionsRow) HasTypeOfInstitution() bool`

HasTypeOfInstitution returns a boolean if a field has been set.

### SetTypeOfInstitutionNil

`func (o *RwdInstitutionsRow) SetTypeOfInstitutionNil(b bool)`

 SetTypeOfInstitutionNil sets the value for TypeOfInstitution to be an explicit nil

### UnsetTypeOfInstitution
`func (o *RwdInstitutionsRow) UnsetTypeOfInstitution()`

UnsetTypeOfInstitution ensures that no value is present for TypeOfInstitution, not even an explicit nil
### GetInstitutionRole

`func (o *RwdInstitutionsRow) GetInstitutionRole() string`

GetInstitutionRole returns the InstitutionRole field if non-nil, zero value otherwise.

### GetInstitutionRoleOk

`func (o *RwdInstitutionsRow) GetInstitutionRoleOk() (*string, bool)`

GetInstitutionRoleOk returns a tuple with the InstitutionRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionRole

`func (o *RwdInstitutionsRow) SetInstitutionRole(v string)`

SetInstitutionRole sets InstitutionRole field to given value.

### HasInstitutionRole

`func (o *RwdInstitutionsRow) HasInstitutionRole() bool`

HasInstitutionRole returns a boolean if a field has been set.

### SetInstitutionRoleNil

`func (o *RwdInstitutionsRow) SetInstitutionRoleNil(b bool)`

 SetInstitutionRoleNil sets the value for InstitutionRole to be an explicit nil

### UnsetInstitutionRole
`func (o *RwdInstitutionsRow) UnsetInstitutionRole()`

UnsetInstitutionRole ensures that no value is present for InstitutionRole, not even an explicit nil
### GetInstitutionCountries

`func (o *RwdInstitutionsRow) GetInstitutionCountries() string`

GetInstitutionCountries returns the InstitutionCountries field if non-nil, zero value otherwise.

### GetInstitutionCountriesOk

`func (o *RwdInstitutionsRow) GetInstitutionCountriesOk() (*string, bool)`

GetInstitutionCountriesOk returns a tuple with the InstitutionCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionCountries

`func (o *RwdInstitutionsRow) SetInstitutionCountries(v string)`

SetInstitutionCountries sets InstitutionCountries field to given value.

### HasInstitutionCountries

`func (o *RwdInstitutionsRow) HasInstitutionCountries() bool`

HasInstitutionCountries returns a boolean if a field has been set.

### SetInstitutionCountriesNil

`func (o *RwdInstitutionsRow) SetInstitutionCountriesNil(b bool)`

 SetInstitutionCountriesNil sets the value for InstitutionCountries to be an explicit nil

### UnsetInstitutionCountries
`func (o *RwdInstitutionsRow) UnsetInstitutionCountries()`

UnsetInstitutionCountries ensures that no value is present for InstitutionCountries, not even an explicit nil
### GetInstitutionCountryNames

`func (o *RwdInstitutionsRow) GetInstitutionCountryNames() []string`

GetInstitutionCountryNames returns the InstitutionCountryNames field if non-nil, zero value otherwise.

### GetInstitutionCountryNamesOk

`func (o *RwdInstitutionsRow) GetInstitutionCountryNamesOk() (*[]string, bool)`

GetInstitutionCountryNamesOk returns a tuple with the InstitutionCountryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionCountryNames

`func (o *RwdInstitutionsRow) SetInstitutionCountryNames(v []string)`

SetInstitutionCountryNames sets InstitutionCountryNames field to given value.

### HasInstitutionCountryNames

`func (o *RwdInstitutionsRow) HasInstitutionCountryNames() bool`

HasInstitutionCountryNames returns a boolean if a field has been set.

### SetInstitutionCountryNamesNil

`func (o *RwdInstitutionsRow) SetInstitutionCountryNamesNil(b bool)`

 SetInstitutionCountryNamesNil sets the value for InstitutionCountryNames to be an explicit nil

### UnsetInstitutionCountryNames
`func (o *RwdInstitutionsRow) UnsetInstitutionCountryNames()`

UnsetInstitutionCountryNames ensures that no value is present for InstitutionCountryNames, not even an explicit nil
### GetInstitutionCountryCount

`func (o *RwdInstitutionsRow) GetInstitutionCountryCount() int32`

GetInstitutionCountryCount returns the InstitutionCountryCount field if non-nil, zero value otherwise.

### GetInstitutionCountryCountOk

`func (o *RwdInstitutionsRow) GetInstitutionCountryCountOk() (*int32, bool)`

GetInstitutionCountryCountOk returns a tuple with the InstitutionCountryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionCountryCount

`func (o *RwdInstitutionsRow) SetInstitutionCountryCount(v int32)`

SetInstitutionCountryCount sets InstitutionCountryCount field to given value.

### HasInstitutionCountryCount

`func (o *RwdInstitutionsRow) HasInstitutionCountryCount() bool`

HasInstitutionCountryCount returns a boolean if a field has been set.

### SetInstitutionCountryCountNil

`func (o *RwdInstitutionsRow) SetInstitutionCountryCountNil(b bool)`

 SetInstitutionCountryCountNil sets the value for InstitutionCountryCount to be an explicit nil

### UnsetInstitutionCountryCount
`func (o *RwdInstitutionsRow) UnsetInstitutionCountryCount()`

UnsetInstitutionCountryCount ensures that no value is present for InstitutionCountryCount, not even an explicit nil
### GetUnmatchedCountryFragments

`func (o *RwdInstitutionsRow) GetUnmatchedCountryFragments() []string`

GetUnmatchedCountryFragments returns the UnmatchedCountryFragments field if non-nil, zero value otherwise.

### GetUnmatchedCountryFragmentsOk

`func (o *RwdInstitutionsRow) GetUnmatchedCountryFragmentsOk() (*[]string, bool)`

GetUnmatchedCountryFragmentsOk returns a tuple with the UnmatchedCountryFragments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedCountryFragments

`func (o *RwdInstitutionsRow) SetUnmatchedCountryFragments(v []string)`

SetUnmatchedCountryFragments sets UnmatchedCountryFragments field to given value.

### HasUnmatchedCountryFragments

`func (o *RwdInstitutionsRow) HasUnmatchedCountryFragments() bool`

HasUnmatchedCountryFragments returns a boolean if a field has been set.

### SetUnmatchedCountryFragmentsNil

`func (o *RwdInstitutionsRow) SetUnmatchedCountryFragmentsNil(b bool)`

 SetUnmatchedCountryFragmentsNil sets the value for UnmatchedCountryFragments to be an explicit nil

### UnsetUnmatchedCountryFragments
`func (o *RwdInstitutionsRow) UnsetUnmatchedCountryFragments()`

UnsetUnmatchedCountryFragments ensures that no value is present for UnmatchedCountryFragments, not even an explicit nil
### GetSingleDeclaredCountryIso2

`func (o *RwdInstitutionsRow) GetSingleDeclaredCountryIso2() string`

GetSingleDeclaredCountryIso2 returns the SingleDeclaredCountryIso2 field if non-nil, zero value otherwise.

### GetSingleDeclaredCountryIso2Ok

`func (o *RwdInstitutionsRow) GetSingleDeclaredCountryIso2Ok() (*string, bool)`

GetSingleDeclaredCountryIso2Ok returns a tuple with the SingleDeclaredCountryIso2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleDeclaredCountryIso2

`func (o *RwdInstitutionsRow) SetSingleDeclaredCountryIso2(v string)`

SetSingleDeclaredCountryIso2 sets SingleDeclaredCountryIso2 field to given value.

### HasSingleDeclaredCountryIso2

`func (o *RwdInstitutionsRow) HasSingleDeclaredCountryIso2() bool`

HasSingleDeclaredCountryIso2 returns a boolean if a field has been set.

### SetSingleDeclaredCountryIso2Nil

`func (o *RwdInstitutionsRow) SetSingleDeclaredCountryIso2Nil(b bool)`

 SetSingleDeclaredCountryIso2Nil sets the value for SingleDeclaredCountryIso2 to be an explicit nil

### UnsetSingleDeclaredCountryIso2
`func (o *RwdInstitutionsRow) UnsetSingleDeclaredCountryIso2()`

UnsetSingleDeclaredCountryIso2 ensures that no value is present for SingleDeclaredCountryIso2, not even an explicit nil
### GetInstitutionWebsite

`func (o *RwdInstitutionsRow) GetInstitutionWebsite() string`

GetInstitutionWebsite returns the InstitutionWebsite field if non-nil, zero value otherwise.

### GetInstitutionWebsiteOk

`func (o *RwdInstitutionsRow) GetInstitutionWebsiteOk() (*string, bool)`

GetInstitutionWebsiteOk returns a tuple with the InstitutionWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionWebsite

`func (o *RwdInstitutionsRow) SetInstitutionWebsite(v string)`

SetInstitutionWebsite sets InstitutionWebsite field to given value.

### HasInstitutionWebsite

`func (o *RwdInstitutionsRow) HasInstitutionWebsite() bool`

HasInstitutionWebsite returns a boolean if a field has been set.

### SetInstitutionWebsiteNil

`func (o *RwdInstitutionsRow) SetInstitutionWebsiteNil(b bool)`

 SetInstitutionWebsiteNil sets the value for InstitutionWebsite to be an explicit nil

### UnsetInstitutionWebsite
`func (o *RwdInstitutionsRow) UnsetInstitutionWebsite()`

UnsetInstitutionWebsite ensures that no value is present for InstitutionWebsite, not even an explicit nil
### GetCatalogueUrl

`func (o *RwdInstitutionsRow) GetCatalogueUrl() string`

GetCatalogueUrl returns the CatalogueUrl field if non-nil, zero value otherwise.

### GetCatalogueUrlOk

`func (o *RwdInstitutionsRow) GetCatalogueUrlOk() (*string, bool)`

GetCatalogueUrlOk returns a tuple with the CatalogueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueUrl

`func (o *RwdInstitutionsRow) SetCatalogueUrl(v string)`

SetCatalogueUrl sets CatalogueUrl field to given value.

### HasCatalogueUrl

`func (o *RwdInstitutionsRow) HasCatalogueUrl() bool`

HasCatalogueUrl returns a boolean if a field has been set.

### SetCatalogueUrlNil

`func (o *RwdInstitutionsRow) SetCatalogueUrlNil(b bool)`

 SetCatalogueUrlNil sets the value for CatalogueUrl to be an explicit nil

### UnsetCatalogueUrl
`func (o *RwdInstitutionsRow) UnsetCatalogueUrl()`

UnsetCatalogueUrl ensures that no value is present for CatalogueUrl, not even an explicit nil
### GetOpenToIndustryFundedResearch

`func (o *RwdInstitutionsRow) GetOpenToIndustryFundedResearch() bool`

GetOpenToIndustryFundedResearch returns the OpenToIndustryFundedResearch field if non-nil, zero value otherwise.

### GetOpenToIndustryFundedResearchOk

`func (o *RwdInstitutionsRow) GetOpenToIndustryFundedResearchOk() (*bool, bool)`

GetOpenToIndustryFundedResearchOk returns a tuple with the OpenToIndustryFundedResearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenToIndustryFundedResearch

`func (o *RwdInstitutionsRow) SetOpenToIndustryFundedResearch(v bool)`

SetOpenToIndustryFundedResearch sets OpenToIndustryFundedResearch field to given value.

### HasOpenToIndustryFundedResearch

`func (o *RwdInstitutionsRow) HasOpenToIndustryFundedResearch() bool`

HasOpenToIndustryFundedResearch returns a boolean if a field has been set.

### SetOpenToIndustryFundedResearchNil

`func (o *RwdInstitutionsRow) SetOpenToIndustryFundedResearchNil(b bool)`

 SetOpenToIndustryFundedResearchNil sets the value for OpenToIndustryFundedResearch to be an explicit nil

### UnsetOpenToIndustryFundedResearch
`func (o *RwdInstitutionsRow) UnsetOpenToIndustryFundedResearch()`

UnsetOpenToIndustryFundedResearch ensures that no value is present for OpenToIndustryFundedResearch, not even an explicit nil
### GetCollectsDataDirectlyFromPatients

`func (o *RwdInstitutionsRow) GetCollectsDataDirectlyFromPatients() bool`

GetCollectsDataDirectlyFromPatients returns the CollectsDataDirectlyFromPatients field if non-nil, zero value otherwise.

### GetCollectsDataDirectlyFromPatientsOk

`func (o *RwdInstitutionsRow) GetCollectsDataDirectlyFromPatientsOk() (*bool, bool)`

GetCollectsDataDirectlyFromPatientsOk returns a tuple with the CollectsDataDirectlyFromPatients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectsDataDirectlyFromPatients

`func (o *RwdInstitutionsRow) SetCollectsDataDirectlyFromPatients(v bool)`

SetCollectsDataDirectlyFromPatients sets CollectsDataDirectlyFromPatients field to given value.

### HasCollectsDataDirectlyFromPatients

`func (o *RwdInstitutionsRow) HasCollectsDataDirectlyFromPatients() bool`

HasCollectsDataDirectlyFromPatients returns a boolean if a field has been set.

### SetCollectsDataDirectlyFromPatientsNil

`func (o *RwdInstitutionsRow) SetCollectsDataDirectlyFromPatientsNil(b bool)`

 SetCollectsDataDirectlyFromPatientsNil sets the value for CollectsDataDirectlyFromPatients to be an explicit nil

### UnsetCollectsDataDirectlyFromPatients
`func (o *RwdInstitutionsRow) UnsetCollectsDataDirectlyFromPatients()`

UnsetCollectsDataDirectlyFromPatients ensures that no value is present for CollectsDataDirectlyFromPatients, not even an explicit nil
### GetIsEnceppPartner

`func (o *RwdInstitutionsRow) GetIsEnceppPartner() bool`

GetIsEnceppPartner returns the IsEnceppPartner field if non-nil, zero value otherwise.

### GetIsEnceppPartnerOk

`func (o *RwdInstitutionsRow) GetIsEnceppPartnerOk() (*bool, bool)`

GetIsEnceppPartnerOk returns a tuple with the IsEnceppPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnceppPartner

`func (o *RwdInstitutionsRow) SetIsEnceppPartner(v bool)`

SetIsEnceppPartner sets IsEnceppPartner field to given value.

### HasIsEnceppPartner

`func (o *RwdInstitutionsRow) HasIsEnceppPartner() bool`

HasIsEnceppPartner returns a boolean if a field has been set.

### SetIsEnceppPartnerNil

`func (o *RwdInstitutionsRow) SetIsEnceppPartnerNil(b bool)`

 SetIsEnceppPartnerNil sets the value for IsEnceppPartner to be an explicit nil

### UnsetIsEnceppPartner
`func (o *RwdInstitutionsRow) UnsetIsEnceppPartner()`

UnsetIsEnceppPartner ensures that no value is present for IsEnceppPartner, not even an explicit nil
### GetRecordUpdatedAt

`func (o *RwdInstitutionsRow) GetRecordUpdatedAt() time.Time`

GetRecordUpdatedAt returns the RecordUpdatedAt field if non-nil, zero value otherwise.

### GetRecordUpdatedAtOk

`func (o *RwdInstitutionsRow) GetRecordUpdatedAtOk() (*time.Time, bool)`

GetRecordUpdatedAtOk returns a tuple with the RecordUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordUpdatedAt

`func (o *RwdInstitutionsRow) SetRecordUpdatedAt(v time.Time)`

SetRecordUpdatedAt sets RecordUpdatedAt field to given value.

### HasRecordUpdatedAt

`func (o *RwdInstitutionsRow) HasRecordUpdatedAt() bool`

HasRecordUpdatedAt returns a boolean if a field has been set.

### SetRecordUpdatedAtNil

`func (o *RwdInstitutionsRow) SetRecordUpdatedAtNil(b bool)`

 SetRecordUpdatedAtNil sets the value for RecordUpdatedAt to be an explicit nil

### UnsetRecordUpdatedAt
`func (o *RwdInstitutionsRow) UnsetRecordUpdatedAt()`

UnsetRecordUpdatedAt ensures that no value is present for RecordUpdatedAt, not even an explicit nil
### GetRecordUpdatedIsMassEdit

`func (o *RwdInstitutionsRow) GetRecordUpdatedIsMassEdit() bool`

GetRecordUpdatedIsMassEdit returns the RecordUpdatedIsMassEdit field if non-nil, zero value otherwise.

### GetRecordUpdatedIsMassEditOk

`func (o *RwdInstitutionsRow) GetRecordUpdatedIsMassEditOk() (*bool, bool)`

GetRecordUpdatedIsMassEditOk returns a tuple with the RecordUpdatedIsMassEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordUpdatedIsMassEdit

`func (o *RwdInstitutionsRow) SetRecordUpdatedIsMassEdit(v bool)`

SetRecordUpdatedIsMassEdit sets RecordUpdatedIsMassEdit field to given value.

### HasRecordUpdatedIsMassEdit

`func (o *RwdInstitutionsRow) HasRecordUpdatedIsMassEdit() bool`

HasRecordUpdatedIsMassEdit returns a boolean if a field has been set.

### SetRecordUpdatedIsMassEditNil

`func (o *RwdInstitutionsRow) SetRecordUpdatedIsMassEditNil(b bool)`

 SetRecordUpdatedIsMassEditNil sets the value for RecordUpdatedIsMassEdit to be an explicit nil

### UnsetRecordUpdatedIsMassEdit
`func (o *RwdInstitutionsRow) UnsetRecordUpdatedIsMassEdit()`

UnsetRecordUpdatedIsMassEdit ensures that no value is present for RecordUpdatedIsMassEdit, not even an explicit nil
### GetRecordUpdatedAtAttested

`func (o *RwdInstitutionsRow) GetRecordUpdatedAtAttested() time.Time`

GetRecordUpdatedAtAttested returns the RecordUpdatedAtAttested field if non-nil, zero value otherwise.

### GetRecordUpdatedAtAttestedOk

`func (o *RwdInstitutionsRow) GetRecordUpdatedAtAttestedOk() (*time.Time, bool)`

GetRecordUpdatedAtAttestedOk returns a tuple with the RecordUpdatedAtAttested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordUpdatedAtAttested

`func (o *RwdInstitutionsRow) SetRecordUpdatedAtAttested(v time.Time)`

SetRecordUpdatedAtAttested sets RecordUpdatedAtAttested field to given value.

### HasRecordUpdatedAtAttested

`func (o *RwdInstitutionsRow) HasRecordUpdatedAtAttested() bool`

HasRecordUpdatedAtAttested returns a boolean if a field has been set.

### SetRecordUpdatedAtAttestedNil

`func (o *RwdInstitutionsRow) SetRecordUpdatedAtAttestedNil(b bool)`

 SetRecordUpdatedAtAttestedNil sets the value for RecordUpdatedAtAttested to be an explicit nil

### UnsetRecordUpdatedAtAttested
`func (o *RwdInstitutionsRow) UnsetRecordUpdatedAtAttested()`

UnsetRecordUpdatedAtAttested ensures that no value is present for RecordUpdatedAtAttested, not even an explicit nil
### GetRecordUpdatedDateShare

`func (o *RwdInstitutionsRow) GetRecordUpdatedDateShare() float32`

GetRecordUpdatedDateShare returns the RecordUpdatedDateShare field if non-nil, zero value otherwise.

### GetRecordUpdatedDateShareOk

`func (o *RwdInstitutionsRow) GetRecordUpdatedDateShareOk() (*float32, bool)`

GetRecordUpdatedDateShareOk returns a tuple with the RecordUpdatedDateShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordUpdatedDateShare

`func (o *RwdInstitutionsRow) SetRecordUpdatedDateShare(v float32)`

SetRecordUpdatedDateShare sets RecordUpdatedDateShare field to given value.

### HasRecordUpdatedDateShare

`func (o *RwdInstitutionsRow) HasRecordUpdatedDateShare() bool`

HasRecordUpdatedDateShare returns a boolean if a field has been set.

### SetRecordUpdatedDateShareNil

`func (o *RwdInstitutionsRow) SetRecordUpdatedDateShareNil(b bool)`

 SetRecordUpdatedDateShareNil sets the value for RecordUpdatedDateShare to be an explicit nil

### UnsetRecordUpdatedDateShare
`func (o *RwdInstitutionsRow) UnsetRecordUpdatedDateShare()`

UnsetRecordUpdatedDateShare ensures that no value is present for RecordUpdatedDateShare, not even an explicit nil
### GetRecordFirstPublishedAt

`func (o *RwdInstitutionsRow) GetRecordFirstPublishedAt() time.Time`

GetRecordFirstPublishedAt returns the RecordFirstPublishedAt field if non-nil, zero value otherwise.

### GetRecordFirstPublishedAtOk

`func (o *RwdInstitutionsRow) GetRecordFirstPublishedAtOk() (*time.Time, bool)`

GetRecordFirstPublishedAtOk returns a tuple with the RecordFirstPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordFirstPublishedAt

`func (o *RwdInstitutionsRow) SetRecordFirstPublishedAt(v time.Time)`

SetRecordFirstPublishedAt sets RecordFirstPublishedAt field to given value.

### HasRecordFirstPublishedAt

`func (o *RwdInstitutionsRow) HasRecordFirstPublishedAt() bool`

HasRecordFirstPublishedAt returns a boolean if a field has been set.

### SetRecordFirstPublishedAtNil

`func (o *RwdInstitutionsRow) SetRecordFirstPublishedAtNil(b bool)`

 SetRecordFirstPublishedAtNil sets the value for RecordFirstPublishedAt to be an explicit nil

### UnsetRecordFirstPublishedAt
`func (o *RwdInstitutionsRow) UnsetRecordFirstPublishedAt()`

UnsetRecordFirstPublishedAt ensures that no value is present for RecordFirstPublishedAt, not even an explicit nil
### GetIsOutdated

`func (o *RwdInstitutionsRow) GetIsOutdated() bool`

GetIsOutdated returns the IsOutdated field if non-nil, zero value otherwise.

### GetIsOutdatedOk

`func (o *RwdInstitutionsRow) GetIsOutdatedOk() (*bool, bool)`

GetIsOutdatedOk returns a tuple with the IsOutdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOutdated

`func (o *RwdInstitutionsRow) SetIsOutdated(v bool)`

SetIsOutdated sets IsOutdated field to given value.

### HasIsOutdated

`func (o *RwdInstitutionsRow) HasIsOutdated() bool`

HasIsOutdated returns a boolean if a field has been set.

### SetIsOutdatedNil

`func (o *RwdInstitutionsRow) SetIsOutdatedNil(b bool)`

 SetIsOutdatedNil sets the value for IsOutdated to be an explicit nil

### UnsetIsOutdated
`func (o *RwdInstitutionsRow) UnsetIsOutdated()`

UnsetIsOutdated ensures that no value is present for IsOutdated, not even an explicit nil
### GetConductingStudiesLifetime

`func (o *RwdInstitutionsRow) GetConductingStudiesLifetime() int32`

GetConductingStudiesLifetime returns the ConductingStudiesLifetime field if non-nil, zero value otherwise.

### GetConductingStudiesLifetimeOk

`func (o *RwdInstitutionsRow) GetConductingStudiesLifetimeOk() (*int32, bool)`

GetConductingStudiesLifetimeOk returns a tuple with the ConductingStudiesLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductingStudiesLifetime

`func (o *RwdInstitutionsRow) SetConductingStudiesLifetime(v int32)`

SetConductingStudiesLifetime sets ConductingStudiesLifetime field to given value.

### HasConductingStudiesLifetime

`func (o *RwdInstitutionsRow) HasConductingStudiesLifetime() bool`

HasConductingStudiesLifetime returns a boolean if a field has been set.

### SetConductingStudiesLifetimeNil

`func (o *RwdInstitutionsRow) SetConductingStudiesLifetimeNil(b bool)`

 SetConductingStudiesLifetimeNil sets the value for ConductingStudiesLifetime to be an explicit nil

### UnsetConductingStudiesLifetime
`func (o *RwdInstitutionsRow) UnsetConductingStudiesLifetime()`

UnsetConductingStudiesLifetime ensures that no value is present for ConductingStudiesLifetime, not even an explicit nil
### GetConductingStudiesLast5Years

`func (o *RwdInstitutionsRow) GetConductingStudiesLast5Years() int32`

GetConductingStudiesLast5Years returns the ConductingStudiesLast5Years field if non-nil, zero value otherwise.

### GetConductingStudiesLast5YearsOk

`func (o *RwdInstitutionsRow) GetConductingStudiesLast5YearsOk() (*int32, bool)`

GetConductingStudiesLast5YearsOk returns a tuple with the ConductingStudiesLast5Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductingStudiesLast5Years

`func (o *RwdInstitutionsRow) SetConductingStudiesLast5Years(v int32)`

SetConductingStudiesLast5Years sets ConductingStudiesLast5Years field to given value.

### HasConductingStudiesLast5Years

`func (o *RwdInstitutionsRow) HasConductingStudiesLast5Years() bool`

HasConductingStudiesLast5Years returns a boolean if a field has been set.

### SetConductingStudiesLast5YearsNil

`func (o *RwdInstitutionsRow) SetConductingStudiesLast5YearsNil(b bool)`

 SetConductingStudiesLast5YearsNil sets the value for ConductingStudiesLast5Years to be an explicit nil

### UnsetConductingStudiesLast5Years
`func (o *RwdInstitutionsRow) UnsetConductingStudiesLast5Years()`

UnsetConductingStudiesLast5Years ensures that no value is present for ConductingStudiesLast5Years, not even an explicit nil
### GetConductingStudiesOngoing

`func (o *RwdInstitutionsRow) GetConductingStudiesOngoing() int32`

GetConductingStudiesOngoing returns the ConductingStudiesOngoing field if non-nil, zero value otherwise.

### GetConductingStudiesOngoingOk

`func (o *RwdInstitutionsRow) GetConductingStudiesOngoingOk() (*int32, bool)`

GetConductingStudiesOngoingOk returns a tuple with the ConductingStudiesOngoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductingStudiesOngoing

`func (o *RwdInstitutionsRow) SetConductingStudiesOngoing(v int32)`

SetConductingStudiesOngoing sets ConductingStudiesOngoing field to given value.

### HasConductingStudiesOngoing

`func (o *RwdInstitutionsRow) HasConductingStudiesOngoing() bool`

HasConductingStudiesOngoing returns a boolean if a field has been set.

### SetConductingStudiesOngoingNil

`func (o *RwdInstitutionsRow) SetConductingStudiesOngoingNil(b bool)`

 SetConductingStudiesOngoingNil sets the value for ConductingStudiesOngoing to be an explicit nil

### UnsetConductingStudiesOngoing
`func (o *RwdInstitutionsRow) UnsetConductingStudiesOngoing()`

UnsetConductingStudiesOngoing ensures that no value is present for ConductingStudiesOngoing, not even an explicit nil
### GetConductingStudiesPlanned

`func (o *RwdInstitutionsRow) GetConductingStudiesPlanned() int32`

GetConductingStudiesPlanned returns the ConductingStudiesPlanned field if non-nil, zero value otherwise.

### GetConductingStudiesPlannedOk

`func (o *RwdInstitutionsRow) GetConductingStudiesPlannedOk() (*int32, bool)`

GetConductingStudiesPlannedOk returns a tuple with the ConductingStudiesPlanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductingStudiesPlanned

`func (o *RwdInstitutionsRow) SetConductingStudiesPlanned(v int32)`

SetConductingStudiesPlanned sets ConductingStudiesPlanned field to given value.

### HasConductingStudiesPlanned

`func (o *RwdInstitutionsRow) HasConductingStudiesPlanned() bool`

HasConductingStudiesPlanned returns a boolean if a field has been set.

### SetConductingStudiesPlannedNil

`func (o *RwdInstitutionsRow) SetConductingStudiesPlannedNil(b bool)`

 SetConductingStudiesPlannedNil sets the value for ConductingStudiesPlanned to be an explicit nil

### UnsetConductingStudiesPlanned
`func (o *RwdInstitutionsRow) UnsetConductingStudiesPlanned()`

UnsetConductingStudiesPlanned ensures that no value is present for ConductingStudiesPlanned, not even an explicit nil
### GetAdditionalRoleStudies

`func (o *RwdInstitutionsRow) GetAdditionalRoleStudies() int32`

GetAdditionalRoleStudies returns the AdditionalRoleStudies field if non-nil, zero value otherwise.

### GetAdditionalRoleStudiesOk

`func (o *RwdInstitutionsRow) GetAdditionalRoleStudiesOk() (*int32, bool)`

GetAdditionalRoleStudiesOk returns a tuple with the AdditionalRoleStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalRoleStudies

`func (o *RwdInstitutionsRow) SetAdditionalRoleStudies(v int32)`

SetAdditionalRoleStudies sets AdditionalRoleStudies field to given value.

### HasAdditionalRoleStudies

`func (o *RwdInstitutionsRow) HasAdditionalRoleStudies() bool`

HasAdditionalRoleStudies returns a boolean if a field has been set.

### SetAdditionalRoleStudiesNil

`func (o *RwdInstitutionsRow) SetAdditionalRoleStudiesNil(b bool)`

 SetAdditionalRoleStudiesNil sets the value for AdditionalRoleStudies to be an explicit nil

### UnsetAdditionalRoleStudies
`func (o *RwdInstitutionsRow) UnsetAdditionalRoleStudies()`

UnsetAdditionalRoleStudies ensures that no value is present for AdditionalRoleStudies, not even an explicit nil
### GetMostRecentStudyStartOn

`func (o *RwdInstitutionsRow) GetMostRecentStudyStartOn() string`

GetMostRecentStudyStartOn returns the MostRecentStudyStartOn field if non-nil, zero value otherwise.

### GetMostRecentStudyStartOnOk

`func (o *RwdInstitutionsRow) GetMostRecentStudyStartOnOk() (*string, bool)`

GetMostRecentStudyStartOnOk returns a tuple with the MostRecentStudyStartOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentStudyStartOn

`func (o *RwdInstitutionsRow) SetMostRecentStudyStartOn(v string)`

SetMostRecentStudyStartOn sets MostRecentStudyStartOn field to given value.

### HasMostRecentStudyStartOn

`func (o *RwdInstitutionsRow) HasMostRecentStudyStartOn() bool`

HasMostRecentStudyStartOn returns a boolean if a field has been set.

### SetMostRecentStudyStartOnNil

`func (o *RwdInstitutionsRow) SetMostRecentStudyStartOnNil(b bool)`

 SetMostRecentStudyStartOnNil sets the value for MostRecentStudyStartOn to be an explicit nil

### UnsetMostRecentStudyStartOn
`func (o *RwdInstitutionsRow) UnsetMostRecentStudyStartOn()`

UnsetMostRecentStudyStartOn ensures that no value is present for MostRecentStudyStartOn, not even an explicit nil
### GetHeldSourceCount

`func (o *RwdInstitutionsRow) GetHeldSourceCount() int32`

GetHeldSourceCount returns the HeldSourceCount field if non-nil, zero value otherwise.

### GetHeldSourceCountOk

`func (o *RwdInstitutionsRow) GetHeldSourceCountOk() (*int32, bool)`

GetHeldSourceCountOk returns a tuple with the HeldSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldSourceCount

`func (o *RwdInstitutionsRow) SetHeldSourceCount(v int32)`

SetHeldSourceCount sets HeldSourceCount field to given value.

### HasHeldSourceCount

`func (o *RwdInstitutionsRow) HasHeldSourceCount() bool`

HasHeldSourceCount returns a boolean if a field has been set.

### SetHeldSourceCountNil

`func (o *RwdInstitutionsRow) SetHeldSourceCountNil(b bool)`

 SetHeldSourceCountNil sets the value for HeldSourceCount to be an explicit nil

### UnsetHeldSourceCount
`func (o *RwdInstitutionsRow) UnsetHeldSourceCount()`

UnsetHeldSourceCount ensures that no value is present for HeldSourceCount, not even an explicit nil
### GetHeldSourceIds

`func (o *RwdInstitutionsRow) GetHeldSourceIds() []string`

GetHeldSourceIds returns the HeldSourceIds field if non-nil, zero value otherwise.

### GetHeldSourceIdsOk

`func (o *RwdInstitutionsRow) GetHeldSourceIdsOk() (*[]string, bool)`

GetHeldSourceIdsOk returns a tuple with the HeldSourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldSourceIds

`func (o *RwdInstitutionsRow) SetHeldSourceIds(v []string)`

SetHeldSourceIds sets HeldSourceIds field to given value.

### HasHeldSourceIds

`func (o *RwdInstitutionsRow) HasHeldSourceIds() bool`

HasHeldSourceIds returns a boolean if a field has been set.

### SetHeldSourceIdsNil

`func (o *RwdInstitutionsRow) SetHeldSourceIdsNil(b bool)`

 SetHeldSourceIdsNil sets the value for HeldSourceIds to be an explicit nil

### UnsetHeldSourceIds
`func (o *RwdInstitutionsRow) UnsetHeldSourceIds()`

UnsetHeldSourceIds ensures that no value is present for HeldSourceIds, not even an explicit nil
### GetHeldActivePopulation

`func (o *RwdInstitutionsRow) GetHeldActivePopulation() int64`

GetHeldActivePopulation returns the HeldActivePopulation field if non-nil, zero value otherwise.

### GetHeldActivePopulationOk

`func (o *RwdInstitutionsRow) GetHeldActivePopulationOk() (*int64, bool)`

GetHeldActivePopulationOk returns a tuple with the HeldActivePopulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldActivePopulation

`func (o *RwdInstitutionsRow) SetHeldActivePopulation(v int64)`

SetHeldActivePopulation sets HeldActivePopulation field to given value.

### HasHeldActivePopulation

`func (o *RwdInstitutionsRow) HasHeldActivePopulation() bool`

HasHeldActivePopulation returns a boolean if a field has been set.

### SetHeldActivePopulationNil

`func (o *RwdInstitutionsRow) SetHeldActivePopulationNil(b bool)`

 SetHeldActivePopulationNil sets the value for HeldActivePopulation to be an explicit nil

### UnsetHeldActivePopulation
`func (o *RwdInstitutionsRow) UnsetHeldActivePopulation()`

UnsetHeldActivePopulation ensures that no value is present for HeldActivePopulation, not even an explicit nil
### GetHeldMaxActivePopulation

`func (o *RwdInstitutionsRow) GetHeldMaxActivePopulation() int64`

GetHeldMaxActivePopulation returns the HeldMaxActivePopulation field if non-nil, zero value otherwise.

### GetHeldMaxActivePopulationOk

`func (o *RwdInstitutionsRow) GetHeldMaxActivePopulationOk() (*int64, bool)`

GetHeldMaxActivePopulationOk returns a tuple with the HeldMaxActivePopulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldMaxActivePopulation

`func (o *RwdInstitutionsRow) SetHeldMaxActivePopulation(v int64)`

SetHeldMaxActivePopulation sets HeldMaxActivePopulation field to given value.

### HasHeldMaxActivePopulation

`func (o *RwdInstitutionsRow) HasHeldMaxActivePopulation() bool`

HasHeldMaxActivePopulation returns a boolean if a field has been set.

### SetHeldMaxActivePopulationNil

`func (o *RwdInstitutionsRow) SetHeldMaxActivePopulationNil(b bool)`

 SetHeldMaxActivePopulationNil sets the value for HeldMaxActivePopulation to be an explicit nil

### UnsetHeldMaxActivePopulation
`func (o *RwdInstitutionsRow) UnsetHeldMaxActivePopulation()`

UnsetHeldMaxActivePopulation ensures that no value is present for HeldMaxActivePopulation, not even an explicit nil
### GetHeldSourcesDeclaringActivePopulation

`func (o *RwdInstitutionsRow) GetHeldSourcesDeclaringActivePopulation() int32`

GetHeldSourcesDeclaringActivePopulation returns the HeldSourcesDeclaringActivePopulation field if non-nil, zero value otherwise.

### GetHeldSourcesDeclaringActivePopulationOk

`func (o *RwdInstitutionsRow) GetHeldSourcesDeclaringActivePopulationOk() (*int32, bool)`

GetHeldSourcesDeclaringActivePopulationOk returns a tuple with the HeldSourcesDeclaringActivePopulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldSourcesDeclaringActivePopulation

`func (o *RwdInstitutionsRow) SetHeldSourcesDeclaringActivePopulation(v int32)`

SetHeldSourcesDeclaringActivePopulation sets HeldSourcesDeclaringActivePopulation field to given value.

### HasHeldSourcesDeclaringActivePopulation

`func (o *RwdInstitutionsRow) HasHeldSourcesDeclaringActivePopulation() bool`

HasHeldSourcesDeclaringActivePopulation returns a boolean if a field has been set.

### SetHeldSourcesDeclaringActivePopulationNil

`func (o *RwdInstitutionsRow) SetHeldSourcesDeclaringActivePopulationNil(b bool)`

 SetHeldSourcesDeclaringActivePopulationNil sets the value for HeldSourcesDeclaringActivePopulation to be an explicit nil

### UnsetHeldSourcesDeclaringActivePopulation
`func (o *RwdInstitutionsRow) UnsetHeldSourcesDeclaringActivePopulation()`

UnsetHeldSourcesDeclaringActivePopulation ensures that no value is present for HeldSourcesDeclaringActivePopulation, not even an explicit nil
### GetHeldQualifiedSourceCount

`func (o *RwdInstitutionsRow) GetHeldQualifiedSourceCount() int32`

GetHeldQualifiedSourceCount returns the HeldQualifiedSourceCount field if non-nil, zero value otherwise.

### GetHeldQualifiedSourceCountOk

`func (o *RwdInstitutionsRow) GetHeldQualifiedSourceCountOk() (*int32, bool)`

GetHeldQualifiedSourceCountOk returns a tuple with the HeldQualifiedSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldQualifiedSourceCount

`func (o *RwdInstitutionsRow) SetHeldQualifiedSourceCount(v int32)`

SetHeldQualifiedSourceCount sets HeldQualifiedSourceCount field to given value.

### HasHeldQualifiedSourceCount

`func (o *RwdInstitutionsRow) HasHeldQualifiedSourceCount() bool`

HasHeldQualifiedSourceCount returns a boolean if a field has been set.

### SetHeldQualifiedSourceCountNil

`func (o *RwdInstitutionsRow) SetHeldQualifiedSourceCountNil(b bool)`

 SetHeldQualifiedSourceCountNil sets the value for HeldQualifiedSourceCount to be an explicit nil

### UnsetHeldQualifiedSourceCount
`func (o *RwdInstitutionsRow) UnsetHeldQualifiedSourceCount()`

UnsetHeldQualifiedSourceCount ensures that no value is present for HeldQualifiedSourceCount, not even an explicit nil
### GetHeldDarwinPartnerSourceCount

`func (o *RwdInstitutionsRow) GetHeldDarwinPartnerSourceCount() int32`

GetHeldDarwinPartnerSourceCount returns the HeldDarwinPartnerSourceCount field if non-nil, zero value otherwise.

### GetHeldDarwinPartnerSourceCountOk

`func (o *RwdInstitutionsRow) GetHeldDarwinPartnerSourceCountOk() (*int32, bool)`

GetHeldDarwinPartnerSourceCountOk returns a tuple with the HeldDarwinPartnerSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldDarwinPartnerSourceCount

`func (o *RwdInstitutionsRow) SetHeldDarwinPartnerSourceCount(v int32)`

SetHeldDarwinPartnerSourceCount sets HeldDarwinPartnerSourceCount field to given value.

### HasHeldDarwinPartnerSourceCount

`func (o *RwdInstitutionsRow) HasHeldDarwinPartnerSourceCount() bool`

HasHeldDarwinPartnerSourceCount returns a boolean if a field has been set.

### SetHeldDarwinPartnerSourceCountNil

`func (o *RwdInstitutionsRow) SetHeldDarwinPartnerSourceCountNil(b bool)`

 SetHeldDarwinPartnerSourceCountNil sets the value for HeldDarwinPartnerSourceCount to be an explicit nil

### UnsetHeldDarwinPartnerSourceCount
`func (o *RwdInstitutionsRow) UnsetHeldDarwinPartnerSourceCount()`

UnsetHeldDarwinPartnerSourceCount ensures that no value is present for HeldDarwinPartnerSourceCount, not even an explicit nil
### GetNetworkMembershipCount

`func (o *RwdInstitutionsRow) GetNetworkMembershipCount() int32`

GetNetworkMembershipCount returns the NetworkMembershipCount field if non-nil, zero value otherwise.

### GetNetworkMembershipCountOk

`func (o *RwdInstitutionsRow) GetNetworkMembershipCountOk() (*int32, bool)`

GetNetworkMembershipCountOk returns a tuple with the NetworkMembershipCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMembershipCount

`func (o *RwdInstitutionsRow) SetNetworkMembershipCount(v int32)`

SetNetworkMembershipCount sets NetworkMembershipCount field to given value.

### HasNetworkMembershipCount

`func (o *RwdInstitutionsRow) HasNetworkMembershipCount() bool`

HasNetworkMembershipCount returns a boolean if a field has been set.

### SetNetworkMembershipCountNil

`func (o *RwdInstitutionsRow) SetNetworkMembershipCountNil(b bool)`

 SetNetworkMembershipCountNil sets the value for NetworkMembershipCount to be an explicit nil

### UnsetNetworkMembershipCount
`func (o *RwdInstitutionsRow) UnsetNetworkMembershipCount()`

UnsetNetworkMembershipCount ensures that no value is present for NetworkMembershipCount, not even an explicit nil
### GetNetworkIds

`func (o *RwdInstitutionsRow) GetNetworkIds() []string`

GetNetworkIds returns the NetworkIds field if non-nil, zero value otherwise.

### GetNetworkIdsOk

`func (o *RwdInstitutionsRow) GetNetworkIdsOk() (*[]string, bool)`

GetNetworkIdsOk returns a tuple with the NetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIds

`func (o *RwdInstitutionsRow) SetNetworkIds(v []string)`

SetNetworkIds sets NetworkIds field to given value.

### HasNetworkIds

`func (o *RwdInstitutionsRow) HasNetworkIds() bool`

HasNetworkIds returns a boolean if a field has been set.

### SetNetworkIdsNil

`func (o *RwdInstitutionsRow) SetNetworkIdsNil(b bool)`

 SetNetworkIdsNil sets the value for NetworkIds to be an explicit nil

### UnsetNetworkIds
`func (o *RwdInstitutionsRow) UnsetNetworkIds()`

UnsetNetworkIds ensures that no value is present for NetworkIds, not even an explicit nil
### GetNetworkNames

`func (o *RwdInstitutionsRow) GetNetworkNames() []string`

GetNetworkNames returns the NetworkNames field if non-nil, zero value otherwise.

### GetNetworkNamesOk

`func (o *RwdInstitutionsRow) GetNetworkNamesOk() (*[]string, bool)`

GetNetworkNamesOk returns a tuple with the NetworkNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkNames

`func (o *RwdInstitutionsRow) SetNetworkNames(v []string)`

SetNetworkNames sets NetworkNames field to given value.

### HasNetworkNames

`func (o *RwdInstitutionsRow) HasNetworkNames() bool`

HasNetworkNames returns a boolean if a field has been set.

### SetNetworkNamesNil

`func (o *RwdInstitutionsRow) SetNetworkNamesNil(b bool)`

 SetNetworkNamesNil sets the value for NetworkNames to be an explicit nil

### UnsetNetworkNames
`func (o *RwdInstitutionsRow) UnsetNetworkNames()`

UnsetNetworkNames ensures that no value is present for NetworkNames, not even an explicit nil
### GetInstitutionRegisterTotal

`func (o *RwdInstitutionsRow) GetInstitutionRegisterTotal() int32`

GetInstitutionRegisterTotal returns the InstitutionRegisterTotal field if non-nil, zero value otherwise.

### GetInstitutionRegisterTotalOk

`func (o *RwdInstitutionsRow) GetInstitutionRegisterTotalOk() (*int32, bool)`

GetInstitutionRegisterTotalOk returns a tuple with the InstitutionRegisterTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionRegisterTotal

`func (o *RwdInstitutionsRow) SetInstitutionRegisterTotal(v int32)`

SetInstitutionRegisterTotal sets InstitutionRegisterTotal field to given value.

### HasInstitutionRegisterTotal

`func (o *RwdInstitutionsRow) HasInstitutionRegisterTotal() bool`

HasInstitutionRegisterTotal returns a boolean if a field has been set.

### SetInstitutionRegisterTotalNil

`func (o *RwdInstitutionsRow) SetInstitutionRegisterTotalNil(b bool)`

 SetInstitutionRegisterTotalNil sets the value for InstitutionRegisterTotal to be an explicit nil

### UnsetInstitutionRegisterTotal
`func (o *RwdInstitutionsRow) UnsetInstitutionRegisterTotal()`

UnsetInstitutionRegisterTotal ensures that no value is present for InstitutionRegisterTotal, not even an explicit nil
### GetSelfRegisteredProfileTotal

`func (o *RwdInstitutionsRow) GetSelfRegisteredProfileTotal() int32`

GetSelfRegisteredProfileTotal returns the SelfRegisteredProfileTotal field if non-nil, zero value otherwise.

### GetSelfRegisteredProfileTotalOk

`func (o *RwdInstitutionsRow) GetSelfRegisteredProfileTotalOk() (*int32, bool)`

GetSelfRegisteredProfileTotalOk returns a tuple with the SelfRegisteredProfileTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfRegisteredProfileTotal

`func (o *RwdInstitutionsRow) SetSelfRegisteredProfileTotal(v int32)`

SetSelfRegisteredProfileTotal sets SelfRegisteredProfileTotal field to given value.

### HasSelfRegisteredProfileTotal

`func (o *RwdInstitutionsRow) HasSelfRegisteredProfileTotal() bool`

HasSelfRegisteredProfileTotal returns a boolean if a field has been set.

### SetSelfRegisteredProfileTotalNil

`func (o *RwdInstitutionsRow) SetSelfRegisteredProfileTotalNil(b bool)`

 SetSelfRegisteredProfileTotalNil sets the value for SelfRegisteredProfileTotal to be an explicit nil

### UnsetSelfRegisteredProfileTotal
`func (o *RwdInstitutionsRow) UnsetSelfRegisteredProfileTotal()`

UnsetSelfRegisteredProfileTotal ensures that no value is present for SelfRegisteredProfileTotal, not even an explicit nil
### GetCatalogueNameEntryTotal

`func (o *RwdInstitutionsRow) GetCatalogueNameEntryTotal() int32`

GetCatalogueNameEntryTotal returns the CatalogueNameEntryTotal field if non-nil, zero value otherwise.

### GetCatalogueNameEntryTotalOk

`func (o *RwdInstitutionsRow) GetCatalogueNameEntryTotalOk() (*int32, bool)`

GetCatalogueNameEntryTotalOk returns a tuple with the CatalogueNameEntryTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueNameEntryTotal

`func (o *RwdInstitutionsRow) SetCatalogueNameEntryTotal(v int32)`

SetCatalogueNameEntryTotal sets CatalogueNameEntryTotal field to given value.

### HasCatalogueNameEntryTotal

`func (o *RwdInstitutionsRow) HasCatalogueNameEntryTotal() bool`

HasCatalogueNameEntryTotal returns a boolean if a field has been set.

### SetCatalogueNameEntryTotalNil

`func (o *RwdInstitutionsRow) SetCatalogueNameEntryTotalNil(b bool)`

 SetCatalogueNameEntryTotalNil sets the value for CatalogueNameEntryTotal to be an explicit nil

### UnsetCatalogueNameEntryTotal
`func (o *RwdInstitutionsRow) UnsetCatalogueNameEntryTotal()`

UnsetCatalogueNameEntryTotal ensures that no value is present for CatalogueNameEntryTotal, not even an explicit nil
### GetInstitutionsDeclaringACountry

`func (o *RwdInstitutionsRow) GetInstitutionsDeclaringACountry() int32`

GetInstitutionsDeclaringACountry returns the InstitutionsDeclaringACountry field if non-nil, zero value otherwise.

### GetInstitutionsDeclaringACountryOk

`func (o *RwdInstitutionsRow) GetInstitutionsDeclaringACountryOk() (*int32, bool)`

GetInstitutionsDeclaringACountryOk returns a tuple with the InstitutionsDeclaringACountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionsDeclaringACountry

`func (o *RwdInstitutionsRow) SetInstitutionsDeclaringACountry(v int32)`

SetInstitutionsDeclaringACountry sets InstitutionsDeclaringACountry field to given value.

### HasInstitutionsDeclaringACountry

`func (o *RwdInstitutionsRow) HasInstitutionsDeclaringACountry() bool`

HasInstitutionsDeclaringACountry returns a boolean if a field has been set.

### SetInstitutionsDeclaringACountryNil

`func (o *RwdInstitutionsRow) SetInstitutionsDeclaringACountryNil(b bool)`

 SetInstitutionsDeclaringACountryNil sets the value for InstitutionsDeclaringACountry to be an explicit nil

### UnsetInstitutionsDeclaringACountry
`func (o *RwdInstitutionsRow) UnsetInstitutionsDeclaringACountry()`

UnsetInstitutionsDeclaringACountry ensures that no value is present for InstitutionsDeclaringACountry, not even an explicit nil
### GetInstitutionsInAnyMemberList

`func (o *RwdInstitutionsRow) GetInstitutionsInAnyMemberList() int32`

GetInstitutionsInAnyMemberList returns the InstitutionsInAnyMemberList field if non-nil, zero value otherwise.

### GetInstitutionsInAnyMemberListOk

`func (o *RwdInstitutionsRow) GetInstitutionsInAnyMemberListOk() (*int32, bool)`

GetInstitutionsInAnyMemberListOk returns a tuple with the InstitutionsInAnyMemberList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionsInAnyMemberList

`func (o *RwdInstitutionsRow) SetInstitutionsInAnyMemberList(v int32)`

SetInstitutionsInAnyMemberList sets InstitutionsInAnyMemberList field to given value.

### HasInstitutionsInAnyMemberList

`func (o *RwdInstitutionsRow) HasInstitutionsInAnyMemberList() bool`

HasInstitutionsInAnyMemberList returns a boolean if a field has been set.

### SetInstitutionsInAnyMemberListNil

`func (o *RwdInstitutionsRow) SetInstitutionsInAnyMemberListNil(b bool)`

 SetInstitutionsInAnyMemberListNil sets the value for InstitutionsInAnyMemberList to be an explicit nil

### UnsetInstitutionsInAnyMemberList
`func (o *RwdInstitutionsRow) UnsetInstitutionsInAnyMemberList()`

UnsetInstitutionsInAnyMemberList ensures that no value is present for InstitutionsInAnyMemberList, not even an explicit nil
### GetInstitutionNameKeyTotal

`func (o *RwdInstitutionsRow) GetInstitutionNameKeyTotal() int32`

GetInstitutionNameKeyTotal returns the InstitutionNameKeyTotal field if non-nil, zero value otherwise.

### GetInstitutionNameKeyTotalOk

`func (o *RwdInstitutionsRow) GetInstitutionNameKeyTotalOk() (*int32, bool)`

GetInstitutionNameKeyTotalOk returns a tuple with the InstitutionNameKeyTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionNameKeyTotal

`func (o *RwdInstitutionsRow) SetInstitutionNameKeyTotal(v int32)`

SetInstitutionNameKeyTotal sets InstitutionNameKeyTotal field to given value.

### HasInstitutionNameKeyTotal

`func (o *RwdInstitutionsRow) HasInstitutionNameKeyTotal() bool`

HasInstitutionNameKeyTotal returns a boolean if a field has been set.

### SetInstitutionNameKeyTotalNil

`func (o *RwdInstitutionsRow) SetInstitutionNameKeyTotalNil(b bool)`

 SetInstitutionNameKeyTotalNil sets the value for InstitutionNameKeyTotal to be an explicit nil

### UnsetInstitutionNameKeyTotal
`func (o *RwdInstitutionsRow) UnsetInstitutionNameKeyTotal()`

UnsetInstitutionNameKeyTotal ensures that no value is present for InstitutionNameKeyTotal, not even an explicit nil
### GetInstitutionsHoldingASource

`func (o *RwdInstitutionsRow) GetInstitutionsHoldingASource() int32`

GetInstitutionsHoldingASource returns the InstitutionsHoldingASource field if non-nil, zero value otherwise.

### GetInstitutionsHoldingASourceOk

`func (o *RwdInstitutionsRow) GetInstitutionsHoldingASourceOk() (*int32, bool)`

GetInstitutionsHoldingASourceOk returns a tuple with the InstitutionsHoldingASource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionsHoldingASource

`func (o *RwdInstitutionsRow) SetInstitutionsHoldingASource(v int32)`

SetInstitutionsHoldingASource sets InstitutionsHoldingASource field to given value.

### HasInstitutionsHoldingASource

`func (o *RwdInstitutionsRow) HasInstitutionsHoldingASource() bool`

HasInstitutionsHoldingASource returns a boolean if a field has been set.

### SetInstitutionsHoldingASourceNil

`func (o *RwdInstitutionsRow) SetInstitutionsHoldingASourceNil(b bool)`

 SetInstitutionsHoldingASourceNil sets the value for InstitutionsHoldingASource to be an explicit nil

### UnsetInstitutionsHoldingASource
`func (o *RwdInstitutionsRow) UnsetInstitutionsHoldingASource()`

UnsetInstitutionsHoldingASource ensures that no value is present for InstitutionsHoldingASource, not even an explicit nil
### GetDeclaredHolderNameTotal

`func (o *RwdInstitutionsRow) GetDeclaredHolderNameTotal() int32`

GetDeclaredHolderNameTotal returns the DeclaredHolderNameTotal field if non-nil, zero value otherwise.

### GetDeclaredHolderNameTotalOk

`func (o *RwdInstitutionsRow) GetDeclaredHolderNameTotalOk() (*int32, bool)`

GetDeclaredHolderNameTotalOk returns a tuple with the DeclaredHolderNameTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredHolderNameTotal

`func (o *RwdInstitutionsRow) SetDeclaredHolderNameTotal(v int32)`

SetDeclaredHolderNameTotal sets DeclaredHolderNameTotal field to given value.

### HasDeclaredHolderNameTotal

`func (o *RwdInstitutionsRow) HasDeclaredHolderNameTotal() bool`

HasDeclaredHolderNameTotal returns a boolean if a field has been set.

### SetDeclaredHolderNameTotalNil

`func (o *RwdInstitutionsRow) SetDeclaredHolderNameTotalNil(b bool)`

 SetDeclaredHolderNameTotalNil sets the value for DeclaredHolderNameTotal to be an explicit nil

### UnsetDeclaredHolderNameTotal
`func (o *RwdInstitutionsRow) UnsetDeclaredHolderNameTotal()`

UnsetDeclaredHolderNameTotal ensures that no value is present for DeclaredHolderNameTotal, not even an explicit nil
### GetDeclaredHolderNameMatched

`func (o *RwdInstitutionsRow) GetDeclaredHolderNameMatched() int32`

GetDeclaredHolderNameMatched returns the DeclaredHolderNameMatched field if non-nil, zero value otherwise.

### GetDeclaredHolderNameMatchedOk

`func (o *RwdInstitutionsRow) GetDeclaredHolderNameMatchedOk() (*int32, bool)`

GetDeclaredHolderNameMatchedOk returns a tuple with the DeclaredHolderNameMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredHolderNameMatched

`func (o *RwdInstitutionsRow) SetDeclaredHolderNameMatched(v int32)`

SetDeclaredHolderNameMatched sets DeclaredHolderNameMatched field to given value.

### HasDeclaredHolderNameMatched

`func (o *RwdInstitutionsRow) HasDeclaredHolderNameMatched() bool`

HasDeclaredHolderNameMatched returns a boolean if a field has been set.

### SetDeclaredHolderNameMatchedNil

`func (o *RwdInstitutionsRow) SetDeclaredHolderNameMatchedNil(b bool)`

 SetDeclaredHolderNameMatchedNil sets the value for DeclaredHolderNameMatched to be an explicit nil

### UnsetDeclaredHolderNameMatched
`func (o *RwdInstitutionsRow) UnsetDeclaredHolderNameMatched()`

UnsetDeclaredHolderNameMatched ensures that no value is present for DeclaredHolderNameMatched, not even an explicit nil
### GetRegisterSourceCount

`func (o *RwdInstitutionsRow) GetRegisterSourceCount() int32`

GetRegisterSourceCount returns the RegisterSourceCount field if non-nil, zero value otherwise.

### GetRegisterSourceCountOk

`func (o *RwdInstitutionsRow) GetRegisterSourceCountOk() (*int32, bool)`

GetRegisterSourceCountOk returns a tuple with the RegisterSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterSourceCount

`func (o *RwdInstitutionsRow) SetRegisterSourceCount(v int32)`

SetRegisterSourceCount sets RegisterSourceCount field to given value.

### HasRegisterSourceCount

`func (o *RwdInstitutionsRow) HasRegisterSourceCount() bool`

HasRegisterSourceCount returns a boolean if a field has been set.

### SetRegisterSourceCountNil

`func (o *RwdInstitutionsRow) SetRegisterSourceCountNil(b bool)`

 SetRegisterSourceCountNil sets the value for RegisterSourceCount to be an explicit nil

### UnsetRegisterSourceCount
`func (o *RwdInstitutionsRow) UnsetRegisterSourceCount()`

UnsetRegisterSourceCount ensures that no value is present for RegisterSourceCount, not even an explicit nil
### GetStudyRegisterTotal

`func (o *RwdInstitutionsRow) GetStudyRegisterTotal() int32`

GetStudyRegisterTotal returns the StudyRegisterTotal field if non-nil, zero value otherwise.

### GetStudyRegisterTotalOk

`func (o *RwdInstitutionsRow) GetStudyRegisterTotalOk() (*int32, bool)`

GetStudyRegisterTotalOk returns a tuple with the StudyRegisterTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyRegisterTotal

`func (o *RwdInstitutionsRow) SetStudyRegisterTotal(v int32)`

SetStudyRegisterTotal sets StudyRegisterTotal field to given value.

### HasStudyRegisterTotal

`func (o *RwdInstitutionsRow) HasStudyRegisterTotal() bool`

HasStudyRegisterTotal returns a boolean if a field has been set.

### SetStudyRegisterTotalNil

`func (o *RwdInstitutionsRow) SetStudyRegisterTotalNil(b bool)`

 SetStudyRegisterTotalNil sets the value for StudyRegisterTotal to be an explicit nil

### UnsetStudyRegisterTotal
`func (o *RwdInstitutionsRow) UnsetStudyRegisterTotal()`

UnsetStudyRegisterTotal ensures that no value is present for StudyRegisterTotal, not even an explicit nil
### GetStudiesNamingARegisteredSource

`func (o *RwdInstitutionsRow) GetStudiesNamingARegisteredSource() int32`

GetStudiesNamingARegisteredSource returns the StudiesNamingARegisteredSource field if non-nil, zero value otherwise.

### GetStudiesNamingARegisteredSourceOk

`func (o *RwdInstitutionsRow) GetStudiesNamingARegisteredSourceOk() (*int32, bool)`

GetStudiesNamingARegisteredSourceOk returns a tuple with the StudiesNamingARegisteredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudiesNamingARegisteredSource

`func (o *RwdInstitutionsRow) SetStudiesNamingARegisteredSource(v int32)`

SetStudiesNamingARegisteredSource sets StudiesNamingARegisteredSource field to given value.

### HasStudiesNamingARegisteredSource

`func (o *RwdInstitutionsRow) HasStudiesNamingARegisteredSource() bool`

HasStudiesNamingARegisteredSource returns a boolean if a field has been set.

### SetStudiesNamingARegisteredSourceNil

`func (o *RwdInstitutionsRow) SetStudiesNamingARegisteredSourceNil(b bool)`

 SetStudiesNamingARegisteredSourceNil sets the value for StudiesNamingARegisteredSource to be an explicit nil

### UnsetStudiesNamingARegisteredSource
`func (o *RwdInstitutionsRow) UnsetStudiesNamingARegisteredSource()`

UnsetStudiesNamingARegisteredSource ensures that no value is present for StudiesNamingARegisteredSource, not even an explicit nil
### GetStudiesNamingARegisteredConductor

`func (o *RwdInstitutionsRow) GetStudiesNamingARegisteredConductor() int32`

GetStudiesNamingARegisteredConductor returns the StudiesNamingARegisteredConductor field if non-nil, zero value otherwise.

### GetStudiesNamingARegisteredConductorOk

`func (o *RwdInstitutionsRow) GetStudiesNamingARegisteredConductorOk() (*int32, bool)`

GetStudiesNamingARegisteredConductorOk returns a tuple with the StudiesNamingARegisteredConductor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudiesNamingARegisteredConductor

`func (o *RwdInstitutionsRow) SetStudiesNamingARegisteredConductor(v int32)`

SetStudiesNamingARegisteredConductor sets StudiesNamingARegisteredConductor field to given value.

### HasStudiesNamingARegisteredConductor

`func (o *RwdInstitutionsRow) HasStudiesNamingARegisteredConductor() bool`

HasStudiesNamingARegisteredConductor returns a boolean if a field has been set.

### SetStudiesNamingARegisteredConductorNil

`func (o *RwdInstitutionsRow) SetStudiesNamingARegisteredConductorNil(b bool)`

 SetStudiesNamingARegisteredConductorNil sets the value for StudiesNamingARegisteredConductor to be an explicit nil

### UnsetStudiesNamingARegisteredConductor
`func (o *RwdInstitutionsRow) UnsetStudiesNamingARegisteredConductor()`

UnsetStudiesNamingARegisteredConductor ensures that no value is present for StudiesNamingARegisteredConductor, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdInstitutionsRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdInstitutionsRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdInstitutionsRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdInstitutionsRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdInstitutionsRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdInstitutionsRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *RwdInstitutionsRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdInstitutionsRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdInstitutionsRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdInstitutionsRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdInstitutionsRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdInstitutionsRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdInstitutionsRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdInstitutionsRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdInstitutionsRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdInstitutionsRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdInstitutionsRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdInstitutionsRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdInstitutionsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdInstitutionsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdInstitutionsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdInstitutionsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdInstitutionsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdInstitutionsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


