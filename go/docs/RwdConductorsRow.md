# RwdConductorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConductorId** | Pointer to **NullableString** | The conductor slug: the register&#39;s normalised institution name key with spaces as hyphens, the id in the web page address (/rwd-conductors/&lt;id&gt;). The entry key of the rwd_conductors object type; injective because no name key carries a hyphen. | [optional] 
**ConductorNameKey** | Pointer to **NullableString** | The normalised institution name key the conductor rollup is keyed on; every register institution whose name folds to it is one conductor. | [optional] 
**ConductorName** | Pointer to **NullableString** | The conductor&#39;s display name: one register spelling of the key. | [optional] 
**ConductorNames** | Pointer to **[]string** | Every register spelling of the key, as the study lists name it. | [optional] 
**RegisterInstitutionIds** | Pointer to **[]string** | The register Institution IDs of every register row that folds to this conductor (more than one when the register lists the name twice). | [optional] 
**RegisterRowCount** | Pointer to **NullableInt32** | How many register institution rows fold to this conductor; above one means the register lists the name more than once and the counts are per name, not per row. | [optional] 
**DeclaredOfficeCountries** | Pointer to **[]string** | The office country text the register institutions declare, one element per distinct declaration; the office is not where the conductor delivered studies (see rwd_conductor_footprint). | [optional] 
**AnyEnceppPartner** | Pointer to **NullableBool** | True when any register row of the conductor declares itself an ENCePP (European Network of Centres for Pharmacoepidemiology and Pharmacovigilance) partner. | [optional] 
**LifetimeStudies** | Pointer to **NullableInt32** | Distinct registered studies the conductor is named as conducting institution on. Not additive with additional_role_studies. | [optional] 
**RecentStudies** | Pointer to **NullableInt32** | Conducting studies that started within five years of the snapshot. | [optional] 
**OngoingStudies** | Pointer to **NullableInt32** | Conducting studies whose status is Ongoing. | [optional] 
**AdditionalRoleStudies** | Pointer to **NullableInt32** | Distinct registered studies the conductor is named on as an additional institution; a different claim from conducting, never added to it. | [optional] 
**MandatedStudies** | Pointer to **NullableInt32** | Conducting studies a regulator required (mandate_stratum mandated). | [optional] 
**MandatedShare** | Pointer to **NullableFloat32** | mandated_studies over lifetime_studies; null when the conductor has no conducting study. | [optional] 
**FinalReportPairs** | Pointer to **NullableInt32** | Conducting studies with both a planned and an actual final report date. | [optional] 
**BackfilledPairs** | Pointer to **NullableInt32** | Of those pairs, the studies whose planned and actual date are equal to the day (a likely backfill). | [optional] 
**ComparablePairs** | Pointer to **NullableInt32** | Conducting studies with a final report slip (both dates present and different); the denominator of the delivery figures. | [optional] 
**BackfilledShare** | Pointer to **NullableFloat32** | backfilled_pairs over final_report_pairs; null when there is no pair. | [optional] 
**NFloor** | Pointer to **NullableInt32** | The floor on comparable_pairs below which the median, quartiles and late count are null (suppressed, not zero). | [optional] 
**MeetsNFloor** | Pointer to **NullableBool** | True when comparable_pairs reaches n_floor, so the slip figures are shown. | [optional] 
**MedianSlipDays** | Pointer to **NullableFloat64** | Median final report slip in days (positive is late) over the conductor&#39;s comparable pairs; null below the floor. | [optional] 
**P25SlipDays** | Pointer to **NullableFloat64** | First quartile of the final report slip in days; null below the floor. | [optional] 
**P75SlipDays** | Pointer to **NullableFloat64** | Third quartile of the final report slip in days; null below the floor. | [optional] 
**LateOver30d** | Pointer to **NullableInt32** | Comparable studies whose final report slipped more than 30 days; null below the floor. | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdConductorsRow

`func NewRwdConductorsRow() *RwdConductorsRow`

NewRwdConductorsRow instantiates a new RwdConductorsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdConductorsRowWithDefaults

`func NewRwdConductorsRowWithDefaults() *RwdConductorsRow`

NewRwdConductorsRowWithDefaults instantiates a new RwdConductorsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConductorId

`func (o *RwdConductorsRow) GetConductorId() string`

GetConductorId returns the ConductorId field if non-nil, zero value otherwise.

### GetConductorIdOk

`func (o *RwdConductorsRow) GetConductorIdOk() (*string, bool)`

GetConductorIdOk returns a tuple with the ConductorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductorId

`func (o *RwdConductorsRow) SetConductorId(v string)`

SetConductorId sets ConductorId field to given value.

### HasConductorId

`func (o *RwdConductorsRow) HasConductorId() bool`

HasConductorId returns a boolean if a field has been set.

### SetConductorIdNil

`func (o *RwdConductorsRow) SetConductorIdNil(b bool)`

 SetConductorIdNil sets the value for ConductorId to be an explicit nil

### UnsetConductorId
`func (o *RwdConductorsRow) UnsetConductorId()`

UnsetConductorId ensures that no value is present for ConductorId, not even an explicit nil
### GetConductorNameKey

`func (o *RwdConductorsRow) GetConductorNameKey() string`

GetConductorNameKey returns the ConductorNameKey field if non-nil, zero value otherwise.

### GetConductorNameKeyOk

`func (o *RwdConductorsRow) GetConductorNameKeyOk() (*string, bool)`

GetConductorNameKeyOk returns a tuple with the ConductorNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductorNameKey

`func (o *RwdConductorsRow) SetConductorNameKey(v string)`

SetConductorNameKey sets ConductorNameKey field to given value.

### HasConductorNameKey

`func (o *RwdConductorsRow) HasConductorNameKey() bool`

HasConductorNameKey returns a boolean if a field has been set.

### SetConductorNameKeyNil

`func (o *RwdConductorsRow) SetConductorNameKeyNil(b bool)`

 SetConductorNameKeyNil sets the value for ConductorNameKey to be an explicit nil

### UnsetConductorNameKey
`func (o *RwdConductorsRow) UnsetConductorNameKey()`

UnsetConductorNameKey ensures that no value is present for ConductorNameKey, not even an explicit nil
### GetConductorName

`func (o *RwdConductorsRow) GetConductorName() string`

GetConductorName returns the ConductorName field if non-nil, zero value otherwise.

### GetConductorNameOk

`func (o *RwdConductorsRow) GetConductorNameOk() (*string, bool)`

GetConductorNameOk returns a tuple with the ConductorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductorName

`func (o *RwdConductorsRow) SetConductorName(v string)`

SetConductorName sets ConductorName field to given value.

### HasConductorName

`func (o *RwdConductorsRow) HasConductorName() bool`

HasConductorName returns a boolean if a field has been set.

### SetConductorNameNil

`func (o *RwdConductorsRow) SetConductorNameNil(b bool)`

 SetConductorNameNil sets the value for ConductorName to be an explicit nil

### UnsetConductorName
`func (o *RwdConductorsRow) UnsetConductorName()`

UnsetConductorName ensures that no value is present for ConductorName, not even an explicit nil
### GetConductorNames

`func (o *RwdConductorsRow) GetConductorNames() []string`

GetConductorNames returns the ConductorNames field if non-nil, zero value otherwise.

### GetConductorNamesOk

`func (o *RwdConductorsRow) GetConductorNamesOk() (*[]string, bool)`

GetConductorNamesOk returns a tuple with the ConductorNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductorNames

`func (o *RwdConductorsRow) SetConductorNames(v []string)`

SetConductorNames sets ConductorNames field to given value.

### HasConductorNames

`func (o *RwdConductorsRow) HasConductorNames() bool`

HasConductorNames returns a boolean if a field has been set.

### SetConductorNamesNil

`func (o *RwdConductorsRow) SetConductorNamesNil(b bool)`

 SetConductorNamesNil sets the value for ConductorNames to be an explicit nil

### UnsetConductorNames
`func (o *RwdConductorsRow) UnsetConductorNames()`

UnsetConductorNames ensures that no value is present for ConductorNames, not even an explicit nil
### GetRegisterInstitutionIds

`func (o *RwdConductorsRow) GetRegisterInstitutionIds() []string`

GetRegisterInstitutionIds returns the RegisterInstitutionIds field if non-nil, zero value otherwise.

### GetRegisterInstitutionIdsOk

`func (o *RwdConductorsRow) GetRegisterInstitutionIdsOk() (*[]string, bool)`

GetRegisterInstitutionIdsOk returns a tuple with the RegisterInstitutionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterInstitutionIds

`func (o *RwdConductorsRow) SetRegisterInstitutionIds(v []string)`

SetRegisterInstitutionIds sets RegisterInstitutionIds field to given value.

### HasRegisterInstitutionIds

`func (o *RwdConductorsRow) HasRegisterInstitutionIds() bool`

HasRegisterInstitutionIds returns a boolean if a field has been set.

### SetRegisterInstitutionIdsNil

`func (o *RwdConductorsRow) SetRegisterInstitutionIdsNil(b bool)`

 SetRegisterInstitutionIdsNil sets the value for RegisterInstitutionIds to be an explicit nil

### UnsetRegisterInstitutionIds
`func (o *RwdConductorsRow) UnsetRegisterInstitutionIds()`

UnsetRegisterInstitutionIds ensures that no value is present for RegisterInstitutionIds, not even an explicit nil
### GetRegisterRowCount

`func (o *RwdConductorsRow) GetRegisterRowCount() int32`

GetRegisterRowCount returns the RegisterRowCount field if non-nil, zero value otherwise.

### GetRegisterRowCountOk

`func (o *RwdConductorsRow) GetRegisterRowCountOk() (*int32, bool)`

GetRegisterRowCountOk returns a tuple with the RegisterRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterRowCount

`func (o *RwdConductorsRow) SetRegisterRowCount(v int32)`

SetRegisterRowCount sets RegisterRowCount field to given value.

### HasRegisterRowCount

`func (o *RwdConductorsRow) HasRegisterRowCount() bool`

HasRegisterRowCount returns a boolean if a field has been set.

### SetRegisterRowCountNil

`func (o *RwdConductorsRow) SetRegisterRowCountNil(b bool)`

 SetRegisterRowCountNil sets the value for RegisterRowCount to be an explicit nil

### UnsetRegisterRowCount
`func (o *RwdConductorsRow) UnsetRegisterRowCount()`

UnsetRegisterRowCount ensures that no value is present for RegisterRowCount, not even an explicit nil
### GetDeclaredOfficeCountries

`func (o *RwdConductorsRow) GetDeclaredOfficeCountries() []string`

GetDeclaredOfficeCountries returns the DeclaredOfficeCountries field if non-nil, zero value otherwise.

### GetDeclaredOfficeCountriesOk

`func (o *RwdConductorsRow) GetDeclaredOfficeCountriesOk() (*[]string, bool)`

GetDeclaredOfficeCountriesOk returns a tuple with the DeclaredOfficeCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredOfficeCountries

`func (o *RwdConductorsRow) SetDeclaredOfficeCountries(v []string)`

SetDeclaredOfficeCountries sets DeclaredOfficeCountries field to given value.

### HasDeclaredOfficeCountries

`func (o *RwdConductorsRow) HasDeclaredOfficeCountries() bool`

HasDeclaredOfficeCountries returns a boolean if a field has been set.

### SetDeclaredOfficeCountriesNil

`func (o *RwdConductorsRow) SetDeclaredOfficeCountriesNil(b bool)`

 SetDeclaredOfficeCountriesNil sets the value for DeclaredOfficeCountries to be an explicit nil

### UnsetDeclaredOfficeCountries
`func (o *RwdConductorsRow) UnsetDeclaredOfficeCountries()`

UnsetDeclaredOfficeCountries ensures that no value is present for DeclaredOfficeCountries, not even an explicit nil
### GetAnyEnceppPartner

`func (o *RwdConductorsRow) GetAnyEnceppPartner() bool`

GetAnyEnceppPartner returns the AnyEnceppPartner field if non-nil, zero value otherwise.

### GetAnyEnceppPartnerOk

`func (o *RwdConductorsRow) GetAnyEnceppPartnerOk() (*bool, bool)`

GetAnyEnceppPartnerOk returns a tuple with the AnyEnceppPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyEnceppPartner

`func (o *RwdConductorsRow) SetAnyEnceppPartner(v bool)`

SetAnyEnceppPartner sets AnyEnceppPartner field to given value.

### HasAnyEnceppPartner

`func (o *RwdConductorsRow) HasAnyEnceppPartner() bool`

HasAnyEnceppPartner returns a boolean if a field has been set.

### SetAnyEnceppPartnerNil

`func (o *RwdConductorsRow) SetAnyEnceppPartnerNil(b bool)`

 SetAnyEnceppPartnerNil sets the value for AnyEnceppPartner to be an explicit nil

### UnsetAnyEnceppPartner
`func (o *RwdConductorsRow) UnsetAnyEnceppPartner()`

UnsetAnyEnceppPartner ensures that no value is present for AnyEnceppPartner, not even an explicit nil
### GetLifetimeStudies

`func (o *RwdConductorsRow) GetLifetimeStudies() int32`

GetLifetimeStudies returns the LifetimeStudies field if non-nil, zero value otherwise.

### GetLifetimeStudiesOk

`func (o *RwdConductorsRow) GetLifetimeStudiesOk() (*int32, bool)`

GetLifetimeStudiesOk returns a tuple with the LifetimeStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeStudies

`func (o *RwdConductorsRow) SetLifetimeStudies(v int32)`

SetLifetimeStudies sets LifetimeStudies field to given value.

### HasLifetimeStudies

`func (o *RwdConductorsRow) HasLifetimeStudies() bool`

HasLifetimeStudies returns a boolean if a field has been set.

### SetLifetimeStudiesNil

`func (o *RwdConductorsRow) SetLifetimeStudiesNil(b bool)`

 SetLifetimeStudiesNil sets the value for LifetimeStudies to be an explicit nil

### UnsetLifetimeStudies
`func (o *RwdConductorsRow) UnsetLifetimeStudies()`

UnsetLifetimeStudies ensures that no value is present for LifetimeStudies, not even an explicit nil
### GetRecentStudies

`func (o *RwdConductorsRow) GetRecentStudies() int32`

GetRecentStudies returns the RecentStudies field if non-nil, zero value otherwise.

### GetRecentStudiesOk

`func (o *RwdConductorsRow) GetRecentStudiesOk() (*int32, bool)`

GetRecentStudiesOk returns a tuple with the RecentStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentStudies

`func (o *RwdConductorsRow) SetRecentStudies(v int32)`

SetRecentStudies sets RecentStudies field to given value.

### HasRecentStudies

`func (o *RwdConductorsRow) HasRecentStudies() bool`

HasRecentStudies returns a boolean if a field has been set.

### SetRecentStudiesNil

`func (o *RwdConductorsRow) SetRecentStudiesNil(b bool)`

 SetRecentStudiesNil sets the value for RecentStudies to be an explicit nil

### UnsetRecentStudies
`func (o *RwdConductorsRow) UnsetRecentStudies()`

UnsetRecentStudies ensures that no value is present for RecentStudies, not even an explicit nil
### GetOngoingStudies

`func (o *RwdConductorsRow) GetOngoingStudies() int32`

GetOngoingStudies returns the OngoingStudies field if non-nil, zero value otherwise.

### GetOngoingStudiesOk

`func (o *RwdConductorsRow) GetOngoingStudiesOk() (*int32, bool)`

GetOngoingStudiesOk returns a tuple with the OngoingStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoingStudies

`func (o *RwdConductorsRow) SetOngoingStudies(v int32)`

SetOngoingStudies sets OngoingStudies field to given value.

### HasOngoingStudies

`func (o *RwdConductorsRow) HasOngoingStudies() bool`

HasOngoingStudies returns a boolean if a field has been set.

### SetOngoingStudiesNil

`func (o *RwdConductorsRow) SetOngoingStudiesNil(b bool)`

 SetOngoingStudiesNil sets the value for OngoingStudies to be an explicit nil

### UnsetOngoingStudies
`func (o *RwdConductorsRow) UnsetOngoingStudies()`

UnsetOngoingStudies ensures that no value is present for OngoingStudies, not even an explicit nil
### GetAdditionalRoleStudies

`func (o *RwdConductorsRow) GetAdditionalRoleStudies() int32`

GetAdditionalRoleStudies returns the AdditionalRoleStudies field if non-nil, zero value otherwise.

### GetAdditionalRoleStudiesOk

`func (o *RwdConductorsRow) GetAdditionalRoleStudiesOk() (*int32, bool)`

GetAdditionalRoleStudiesOk returns a tuple with the AdditionalRoleStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalRoleStudies

`func (o *RwdConductorsRow) SetAdditionalRoleStudies(v int32)`

SetAdditionalRoleStudies sets AdditionalRoleStudies field to given value.

### HasAdditionalRoleStudies

`func (o *RwdConductorsRow) HasAdditionalRoleStudies() bool`

HasAdditionalRoleStudies returns a boolean if a field has been set.

### SetAdditionalRoleStudiesNil

`func (o *RwdConductorsRow) SetAdditionalRoleStudiesNil(b bool)`

 SetAdditionalRoleStudiesNil sets the value for AdditionalRoleStudies to be an explicit nil

### UnsetAdditionalRoleStudies
`func (o *RwdConductorsRow) UnsetAdditionalRoleStudies()`

UnsetAdditionalRoleStudies ensures that no value is present for AdditionalRoleStudies, not even an explicit nil
### GetMandatedStudies

`func (o *RwdConductorsRow) GetMandatedStudies() int32`

GetMandatedStudies returns the MandatedStudies field if non-nil, zero value otherwise.

### GetMandatedStudiesOk

`func (o *RwdConductorsRow) GetMandatedStudiesOk() (*int32, bool)`

GetMandatedStudiesOk returns a tuple with the MandatedStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatedStudies

`func (o *RwdConductorsRow) SetMandatedStudies(v int32)`

SetMandatedStudies sets MandatedStudies field to given value.

### HasMandatedStudies

`func (o *RwdConductorsRow) HasMandatedStudies() bool`

HasMandatedStudies returns a boolean if a field has been set.

### SetMandatedStudiesNil

`func (o *RwdConductorsRow) SetMandatedStudiesNil(b bool)`

 SetMandatedStudiesNil sets the value for MandatedStudies to be an explicit nil

### UnsetMandatedStudies
`func (o *RwdConductorsRow) UnsetMandatedStudies()`

UnsetMandatedStudies ensures that no value is present for MandatedStudies, not even an explicit nil
### GetMandatedShare

`func (o *RwdConductorsRow) GetMandatedShare() float32`

GetMandatedShare returns the MandatedShare field if non-nil, zero value otherwise.

### GetMandatedShareOk

`func (o *RwdConductorsRow) GetMandatedShareOk() (*float32, bool)`

GetMandatedShareOk returns a tuple with the MandatedShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatedShare

`func (o *RwdConductorsRow) SetMandatedShare(v float32)`

SetMandatedShare sets MandatedShare field to given value.

### HasMandatedShare

`func (o *RwdConductorsRow) HasMandatedShare() bool`

HasMandatedShare returns a boolean if a field has been set.

### SetMandatedShareNil

`func (o *RwdConductorsRow) SetMandatedShareNil(b bool)`

 SetMandatedShareNil sets the value for MandatedShare to be an explicit nil

### UnsetMandatedShare
`func (o *RwdConductorsRow) UnsetMandatedShare()`

UnsetMandatedShare ensures that no value is present for MandatedShare, not even an explicit nil
### GetFinalReportPairs

`func (o *RwdConductorsRow) GetFinalReportPairs() int32`

GetFinalReportPairs returns the FinalReportPairs field if non-nil, zero value otherwise.

### GetFinalReportPairsOk

`func (o *RwdConductorsRow) GetFinalReportPairsOk() (*int32, bool)`

GetFinalReportPairsOk returns a tuple with the FinalReportPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalReportPairs

`func (o *RwdConductorsRow) SetFinalReportPairs(v int32)`

SetFinalReportPairs sets FinalReportPairs field to given value.

### HasFinalReportPairs

`func (o *RwdConductorsRow) HasFinalReportPairs() bool`

HasFinalReportPairs returns a boolean if a field has been set.

### SetFinalReportPairsNil

`func (o *RwdConductorsRow) SetFinalReportPairsNil(b bool)`

 SetFinalReportPairsNil sets the value for FinalReportPairs to be an explicit nil

### UnsetFinalReportPairs
`func (o *RwdConductorsRow) UnsetFinalReportPairs()`

UnsetFinalReportPairs ensures that no value is present for FinalReportPairs, not even an explicit nil
### GetBackfilledPairs

`func (o *RwdConductorsRow) GetBackfilledPairs() int32`

GetBackfilledPairs returns the BackfilledPairs field if non-nil, zero value otherwise.

### GetBackfilledPairsOk

`func (o *RwdConductorsRow) GetBackfilledPairsOk() (*int32, bool)`

GetBackfilledPairsOk returns a tuple with the BackfilledPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfilledPairs

`func (o *RwdConductorsRow) SetBackfilledPairs(v int32)`

SetBackfilledPairs sets BackfilledPairs field to given value.

### HasBackfilledPairs

`func (o *RwdConductorsRow) HasBackfilledPairs() bool`

HasBackfilledPairs returns a boolean if a field has been set.

### SetBackfilledPairsNil

`func (o *RwdConductorsRow) SetBackfilledPairsNil(b bool)`

 SetBackfilledPairsNil sets the value for BackfilledPairs to be an explicit nil

### UnsetBackfilledPairs
`func (o *RwdConductorsRow) UnsetBackfilledPairs()`

UnsetBackfilledPairs ensures that no value is present for BackfilledPairs, not even an explicit nil
### GetComparablePairs

`func (o *RwdConductorsRow) GetComparablePairs() int32`

GetComparablePairs returns the ComparablePairs field if non-nil, zero value otherwise.

### GetComparablePairsOk

`func (o *RwdConductorsRow) GetComparablePairsOk() (*int32, bool)`

GetComparablePairsOk returns a tuple with the ComparablePairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparablePairs

`func (o *RwdConductorsRow) SetComparablePairs(v int32)`

SetComparablePairs sets ComparablePairs field to given value.

### HasComparablePairs

`func (o *RwdConductorsRow) HasComparablePairs() bool`

HasComparablePairs returns a boolean if a field has been set.

### SetComparablePairsNil

`func (o *RwdConductorsRow) SetComparablePairsNil(b bool)`

 SetComparablePairsNil sets the value for ComparablePairs to be an explicit nil

### UnsetComparablePairs
`func (o *RwdConductorsRow) UnsetComparablePairs()`

UnsetComparablePairs ensures that no value is present for ComparablePairs, not even an explicit nil
### GetBackfilledShare

`func (o *RwdConductorsRow) GetBackfilledShare() float32`

GetBackfilledShare returns the BackfilledShare field if non-nil, zero value otherwise.

### GetBackfilledShareOk

`func (o *RwdConductorsRow) GetBackfilledShareOk() (*float32, bool)`

GetBackfilledShareOk returns a tuple with the BackfilledShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfilledShare

`func (o *RwdConductorsRow) SetBackfilledShare(v float32)`

SetBackfilledShare sets BackfilledShare field to given value.

### HasBackfilledShare

`func (o *RwdConductorsRow) HasBackfilledShare() bool`

HasBackfilledShare returns a boolean if a field has been set.

### SetBackfilledShareNil

`func (o *RwdConductorsRow) SetBackfilledShareNil(b bool)`

 SetBackfilledShareNil sets the value for BackfilledShare to be an explicit nil

### UnsetBackfilledShare
`func (o *RwdConductorsRow) UnsetBackfilledShare()`

UnsetBackfilledShare ensures that no value is present for BackfilledShare, not even an explicit nil
### GetNFloor

`func (o *RwdConductorsRow) GetNFloor() int32`

GetNFloor returns the NFloor field if non-nil, zero value otherwise.

### GetNFloorOk

`func (o *RwdConductorsRow) GetNFloorOk() (*int32, bool)`

GetNFloorOk returns a tuple with the NFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFloor

`func (o *RwdConductorsRow) SetNFloor(v int32)`

SetNFloor sets NFloor field to given value.

### HasNFloor

`func (o *RwdConductorsRow) HasNFloor() bool`

HasNFloor returns a boolean if a field has been set.

### SetNFloorNil

`func (o *RwdConductorsRow) SetNFloorNil(b bool)`

 SetNFloorNil sets the value for NFloor to be an explicit nil

### UnsetNFloor
`func (o *RwdConductorsRow) UnsetNFloor()`

UnsetNFloor ensures that no value is present for NFloor, not even an explicit nil
### GetMeetsNFloor

`func (o *RwdConductorsRow) GetMeetsNFloor() bool`

GetMeetsNFloor returns the MeetsNFloor field if non-nil, zero value otherwise.

### GetMeetsNFloorOk

`func (o *RwdConductorsRow) GetMeetsNFloorOk() (*bool, bool)`

GetMeetsNFloorOk returns a tuple with the MeetsNFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetsNFloor

`func (o *RwdConductorsRow) SetMeetsNFloor(v bool)`

SetMeetsNFloor sets MeetsNFloor field to given value.

### HasMeetsNFloor

`func (o *RwdConductorsRow) HasMeetsNFloor() bool`

HasMeetsNFloor returns a boolean if a field has been set.

### SetMeetsNFloorNil

`func (o *RwdConductorsRow) SetMeetsNFloorNil(b bool)`

 SetMeetsNFloorNil sets the value for MeetsNFloor to be an explicit nil

### UnsetMeetsNFloor
`func (o *RwdConductorsRow) UnsetMeetsNFloor()`

UnsetMeetsNFloor ensures that no value is present for MeetsNFloor, not even an explicit nil
### GetMedianSlipDays

`func (o *RwdConductorsRow) GetMedianSlipDays() float64`

GetMedianSlipDays returns the MedianSlipDays field if non-nil, zero value otherwise.

### GetMedianSlipDaysOk

`func (o *RwdConductorsRow) GetMedianSlipDaysOk() (*float64, bool)`

GetMedianSlipDaysOk returns a tuple with the MedianSlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianSlipDays

`func (o *RwdConductorsRow) SetMedianSlipDays(v float64)`

SetMedianSlipDays sets MedianSlipDays field to given value.

### HasMedianSlipDays

`func (o *RwdConductorsRow) HasMedianSlipDays() bool`

HasMedianSlipDays returns a boolean if a field has been set.

### SetMedianSlipDaysNil

`func (o *RwdConductorsRow) SetMedianSlipDaysNil(b bool)`

 SetMedianSlipDaysNil sets the value for MedianSlipDays to be an explicit nil

### UnsetMedianSlipDays
`func (o *RwdConductorsRow) UnsetMedianSlipDays()`

UnsetMedianSlipDays ensures that no value is present for MedianSlipDays, not even an explicit nil
### GetP25SlipDays

`func (o *RwdConductorsRow) GetP25SlipDays() float64`

GetP25SlipDays returns the P25SlipDays field if non-nil, zero value otherwise.

### GetP25SlipDaysOk

`func (o *RwdConductorsRow) GetP25SlipDaysOk() (*float64, bool)`

GetP25SlipDaysOk returns a tuple with the P25SlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP25SlipDays

`func (o *RwdConductorsRow) SetP25SlipDays(v float64)`

SetP25SlipDays sets P25SlipDays field to given value.

### HasP25SlipDays

`func (o *RwdConductorsRow) HasP25SlipDays() bool`

HasP25SlipDays returns a boolean if a field has been set.

### SetP25SlipDaysNil

`func (o *RwdConductorsRow) SetP25SlipDaysNil(b bool)`

 SetP25SlipDaysNil sets the value for P25SlipDays to be an explicit nil

### UnsetP25SlipDays
`func (o *RwdConductorsRow) UnsetP25SlipDays()`

UnsetP25SlipDays ensures that no value is present for P25SlipDays, not even an explicit nil
### GetP75SlipDays

`func (o *RwdConductorsRow) GetP75SlipDays() float64`

GetP75SlipDays returns the P75SlipDays field if non-nil, zero value otherwise.

### GetP75SlipDaysOk

`func (o *RwdConductorsRow) GetP75SlipDaysOk() (*float64, bool)`

GetP75SlipDaysOk returns a tuple with the P75SlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP75SlipDays

`func (o *RwdConductorsRow) SetP75SlipDays(v float64)`

SetP75SlipDays sets P75SlipDays field to given value.

### HasP75SlipDays

`func (o *RwdConductorsRow) HasP75SlipDays() bool`

HasP75SlipDays returns a boolean if a field has been set.

### SetP75SlipDaysNil

`func (o *RwdConductorsRow) SetP75SlipDaysNil(b bool)`

 SetP75SlipDaysNil sets the value for P75SlipDays to be an explicit nil

### UnsetP75SlipDays
`func (o *RwdConductorsRow) UnsetP75SlipDays()`

UnsetP75SlipDays ensures that no value is present for P75SlipDays, not even an explicit nil
### GetLateOver30d

`func (o *RwdConductorsRow) GetLateOver30d() int32`

GetLateOver30d returns the LateOver30d field if non-nil, zero value otherwise.

### GetLateOver30dOk

`func (o *RwdConductorsRow) GetLateOver30dOk() (*int32, bool)`

GetLateOver30dOk returns a tuple with the LateOver30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateOver30d

`func (o *RwdConductorsRow) SetLateOver30d(v int32)`

SetLateOver30d sets LateOver30d field to given value.

### HasLateOver30d

`func (o *RwdConductorsRow) HasLateOver30d() bool`

HasLateOver30d returns a boolean if a field has been set.

### SetLateOver30dNil

`func (o *RwdConductorsRow) SetLateOver30dNil(b bool)`

 SetLateOver30dNil sets the value for LateOver30d to be an explicit nil

### UnsetLateOver30d
`func (o *RwdConductorsRow) UnsetLateOver30d()`

UnsetLateOver30d ensures that no value is present for LateOver30d, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdConductorsRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdConductorsRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdConductorsRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdConductorsRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdConductorsRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdConductorsRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *RwdConductorsRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdConductorsRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdConductorsRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdConductorsRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdConductorsRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdConductorsRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdConductorsRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdConductorsRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdConductorsRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdConductorsRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdConductorsRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdConductorsRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdConductorsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdConductorsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdConductorsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdConductorsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdConductorsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdConductorsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


