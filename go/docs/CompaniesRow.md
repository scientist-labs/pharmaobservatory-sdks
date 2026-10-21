# CompaniesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | The company key (company:&lt;key&gt;). | [optional] 
**CompanyName** | Pointer to **NullableString** | The company&#39;s display name: its own legal entity&#39;s name where one is known. | [optional] 
**NameBasis** | Pointer to **NullableString** | Which source supplied the name: gleif_legal_name, edgar_registrant, ror or sponsor_display_name. | [optional] 
**Country** | Pointer to **NullableString** | Country of the company&#39;s own legal entity, as the source states it. | [optional] 
**OrgType** | Pointer to **NullableString** | The organization type of the company&#39;s representative record: Academic, Government, Hospital, Industry, Nonprofit or Other (from the Research Organization Registry (ROR) record or the filing), or the lowercase tags investor and manufacturer that mark a minted adviser or manufacturer record. | [optional] 
**OrgKinds** | Pointer to **[]string** | The roles the company&#39;s organizations play (Manufacturer, Sponsor and so on). | [optional] 
**EdgarCik** | Pointer to **NullableString** | Central index key (CIK) in EDGAR, the US Securities and Exchange Commission (SEC) filing system, of the representative record, when it files. | [optional] 
**EdgarTicker** | Pointer to **NullableString** | Exchange ticker of the representative record, when listed. | [optional] 
**RepresentativeOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the company&#39;s page shows. | [optional] 
**ParentCompanyId** | Pointer to **NullableString** | The direct owner company (part_of), null for a root. | [optional] 
**ParentCompanyName** | Pointer to **NullableString** | Display name of the direct owner company. | [optional] 
**PartOfSince** | Pointer to **NullableString** | Date the ownership link began, when every hop is dated. | [optional] 
**PartOfSincePrecision** | Pointer to **NullableString** | Precision of part_of_since (day, month or year). | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the current ultimate parent company; a root company&#39;s group is itself. Current ownership only; dated links are in company_members. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**RollupDepth** | Pointer to **NullableInt32** | Hops from this company to its ultimate parent (0 for a root). | [optional] 
**NLegalEntities** | Pointer to **NullableInt64** | Registered legal entities that are members of the company. | [optional] 
**NMemberCompanies** | Pointer to **NullableInt64** | Companies folded into this one. | [optional] 
**NumLeadTrials** | Pointer to **NullableInt64** | Trials the company&#39;s organizations lead. | [optional] 
**NumCollaboratorTrials** | Pointer to **NullableInt64** | Trials the company&#39;s organizations collaborate on. | [optional] 
**NumTrials** | Pointer to **NullableInt64** | Trials the company&#39;s organizations lead or collaborate on. | [optional] 
**TotalLeadEnrollment** | Pointer to **NullableInt64** | Planned enrollment across led trials. | [optional] 
**NumRecentLeadTrials** | Pointer to **NullableInt64** | Led trials started in the recent window. | [optional] 
**ActivityScore** | Pointer to **NullableFloat32** | Trial activity score (interventional lead trials and enrollment, weighted toward recent). | [optional] 
**IsAdmitted** | Pointer to **NullableBool** | True when the company is admitted to the company roster (lead-trial activity, more than one legal entity or a recorded manufacturer role); false rows carry null firmographic, trial-count and activity columns. | [optional] 
**HasPublicPage** | Pointer to **NullableBool** | True when the company has a public organization page. | [optional] 

## Methods

### NewCompaniesRow

`func NewCompaniesRow() *CompaniesRow`

NewCompaniesRow instantiates a new CompaniesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompaniesRowWithDefaults

`func NewCompaniesRowWithDefaults() *CompaniesRow`

NewCompaniesRowWithDefaults instantiates a new CompaniesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompaniesRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompaniesRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompaniesRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompaniesRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompaniesRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompaniesRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompaniesRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompaniesRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompaniesRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompaniesRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompaniesRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompaniesRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetNameBasis

`func (o *CompaniesRow) GetNameBasis() string`

GetNameBasis returns the NameBasis field if non-nil, zero value otherwise.

### GetNameBasisOk

`func (o *CompaniesRow) GetNameBasisOk() (*string, bool)`

GetNameBasisOk returns a tuple with the NameBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameBasis

`func (o *CompaniesRow) SetNameBasis(v string)`

SetNameBasis sets NameBasis field to given value.

### HasNameBasis

`func (o *CompaniesRow) HasNameBasis() bool`

HasNameBasis returns a boolean if a field has been set.

### SetNameBasisNil

`func (o *CompaniesRow) SetNameBasisNil(b bool)`

 SetNameBasisNil sets the value for NameBasis to be an explicit nil

### UnsetNameBasis
`func (o *CompaniesRow) UnsetNameBasis()`

UnsetNameBasis ensures that no value is present for NameBasis, not even an explicit nil
### GetCountry

`func (o *CompaniesRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CompaniesRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CompaniesRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CompaniesRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *CompaniesRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *CompaniesRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetOrgType

`func (o *CompaniesRow) GetOrgType() string`

GetOrgType returns the OrgType field if non-nil, zero value otherwise.

### GetOrgTypeOk

`func (o *CompaniesRow) GetOrgTypeOk() (*string, bool)`

GetOrgTypeOk returns a tuple with the OrgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgType

`func (o *CompaniesRow) SetOrgType(v string)`

SetOrgType sets OrgType field to given value.

### HasOrgType

`func (o *CompaniesRow) HasOrgType() bool`

HasOrgType returns a boolean if a field has been set.

### SetOrgTypeNil

`func (o *CompaniesRow) SetOrgTypeNil(b bool)`

 SetOrgTypeNil sets the value for OrgType to be an explicit nil

### UnsetOrgType
`func (o *CompaniesRow) UnsetOrgType()`

UnsetOrgType ensures that no value is present for OrgType, not even an explicit nil
### GetOrgKinds

`func (o *CompaniesRow) GetOrgKinds() []string`

GetOrgKinds returns the OrgKinds field if non-nil, zero value otherwise.

### GetOrgKindsOk

`func (o *CompaniesRow) GetOrgKindsOk() (*[]string, bool)`

GetOrgKindsOk returns a tuple with the OrgKinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgKinds

`func (o *CompaniesRow) SetOrgKinds(v []string)`

SetOrgKinds sets OrgKinds field to given value.

### HasOrgKinds

`func (o *CompaniesRow) HasOrgKinds() bool`

HasOrgKinds returns a boolean if a field has been set.

### SetOrgKindsNil

`func (o *CompaniesRow) SetOrgKindsNil(b bool)`

 SetOrgKindsNil sets the value for OrgKinds to be an explicit nil

### UnsetOrgKinds
`func (o *CompaniesRow) UnsetOrgKinds()`

UnsetOrgKinds ensures that no value is present for OrgKinds, not even an explicit nil
### GetEdgarCik

`func (o *CompaniesRow) GetEdgarCik() string`

GetEdgarCik returns the EdgarCik field if non-nil, zero value otherwise.

### GetEdgarCikOk

`func (o *CompaniesRow) GetEdgarCikOk() (*string, bool)`

GetEdgarCikOk returns a tuple with the EdgarCik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgarCik

`func (o *CompaniesRow) SetEdgarCik(v string)`

SetEdgarCik sets EdgarCik field to given value.

### HasEdgarCik

`func (o *CompaniesRow) HasEdgarCik() bool`

HasEdgarCik returns a boolean if a field has been set.

### SetEdgarCikNil

`func (o *CompaniesRow) SetEdgarCikNil(b bool)`

 SetEdgarCikNil sets the value for EdgarCik to be an explicit nil

### UnsetEdgarCik
`func (o *CompaniesRow) UnsetEdgarCik()`

UnsetEdgarCik ensures that no value is present for EdgarCik, not even an explicit nil
### GetEdgarTicker

`func (o *CompaniesRow) GetEdgarTicker() string`

GetEdgarTicker returns the EdgarTicker field if non-nil, zero value otherwise.

### GetEdgarTickerOk

`func (o *CompaniesRow) GetEdgarTickerOk() (*string, bool)`

GetEdgarTickerOk returns a tuple with the EdgarTicker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgarTicker

`func (o *CompaniesRow) SetEdgarTicker(v string)`

SetEdgarTicker sets EdgarTicker field to given value.

### HasEdgarTicker

`func (o *CompaniesRow) HasEdgarTicker() bool`

HasEdgarTicker returns a boolean if a field has been set.

### SetEdgarTickerNil

`func (o *CompaniesRow) SetEdgarTickerNil(b bool)`

 SetEdgarTickerNil sets the value for EdgarTicker to be an explicit nil

### UnsetEdgarTicker
`func (o *CompaniesRow) UnsetEdgarTicker()`

UnsetEdgarTicker ensures that no value is present for EdgarTicker, not even an explicit nil
### GetRepresentativeOrgId

`func (o *CompaniesRow) GetRepresentativeOrgId() string`

GetRepresentativeOrgId returns the RepresentativeOrgId field if non-nil, zero value otherwise.

### GetRepresentativeOrgIdOk

`func (o *CompaniesRow) GetRepresentativeOrgIdOk() (*string, bool)`

GetRepresentativeOrgIdOk returns a tuple with the RepresentativeOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentativeOrgId

`func (o *CompaniesRow) SetRepresentativeOrgId(v string)`

SetRepresentativeOrgId sets RepresentativeOrgId field to given value.

### HasRepresentativeOrgId

`func (o *CompaniesRow) HasRepresentativeOrgId() bool`

HasRepresentativeOrgId returns a boolean if a field has been set.

### SetRepresentativeOrgIdNil

`func (o *CompaniesRow) SetRepresentativeOrgIdNil(b bool)`

 SetRepresentativeOrgIdNil sets the value for RepresentativeOrgId to be an explicit nil

### UnsetRepresentativeOrgId
`func (o *CompaniesRow) UnsetRepresentativeOrgId()`

UnsetRepresentativeOrgId ensures that no value is present for RepresentativeOrgId, not even an explicit nil
### GetParentCompanyId

`func (o *CompaniesRow) GetParentCompanyId() string`

GetParentCompanyId returns the ParentCompanyId field if non-nil, zero value otherwise.

### GetParentCompanyIdOk

`func (o *CompaniesRow) GetParentCompanyIdOk() (*string, bool)`

GetParentCompanyIdOk returns a tuple with the ParentCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCompanyId

`func (o *CompaniesRow) SetParentCompanyId(v string)`

SetParentCompanyId sets ParentCompanyId field to given value.

### HasParentCompanyId

`func (o *CompaniesRow) HasParentCompanyId() bool`

HasParentCompanyId returns a boolean if a field has been set.

### SetParentCompanyIdNil

`func (o *CompaniesRow) SetParentCompanyIdNil(b bool)`

 SetParentCompanyIdNil sets the value for ParentCompanyId to be an explicit nil

### UnsetParentCompanyId
`func (o *CompaniesRow) UnsetParentCompanyId()`

UnsetParentCompanyId ensures that no value is present for ParentCompanyId, not even an explicit nil
### GetParentCompanyName

`func (o *CompaniesRow) GetParentCompanyName() string`

GetParentCompanyName returns the ParentCompanyName field if non-nil, zero value otherwise.

### GetParentCompanyNameOk

`func (o *CompaniesRow) GetParentCompanyNameOk() (*string, bool)`

GetParentCompanyNameOk returns a tuple with the ParentCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCompanyName

`func (o *CompaniesRow) SetParentCompanyName(v string)`

SetParentCompanyName sets ParentCompanyName field to given value.

### HasParentCompanyName

`func (o *CompaniesRow) HasParentCompanyName() bool`

HasParentCompanyName returns a boolean if a field has been set.

### SetParentCompanyNameNil

`func (o *CompaniesRow) SetParentCompanyNameNil(b bool)`

 SetParentCompanyNameNil sets the value for ParentCompanyName to be an explicit nil

### UnsetParentCompanyName
`func (o *CompaniesRow) UnsetParentCompanyName()`

UnsetParentCompanyName ensures that no value is present for ParentCompanyName, not even an explicit nil
### GetPartOfSince

`func (o *CompaniesRow) GetPartOfSince() string`

GetPartOfSince returns the PartOfSince field if non-nil, zero value otherwise.

### GetPartOfSinceOk

`func (o *CompaniesRow) GetPartOfSinceOk() (*string, bool)`

GetPartOfSinceOk returns a tuple with the PartOfSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOfSince

`func (o *CompaniesRow) SetPartOfSince(v string)`

SetPartOfSince sets PartOfSince field to given value.

### HasPartOfSince

`func (o *CompaniesRow) HasPartOfSince() bool`

HasPartOfSince returns a boolean if a field has been set.

### SetPartOfSinceNil

`func (o *CompaniesRow) SetPartOfSinceNil(b bool)`

 SetPartOfSinceNil sets the value for PartOfSince to be an explicit nil

### UnsetPartOfSince
`func (o *CompaniesRow) UnsetPartOfSince()`

UnsetPartOfSince ensures that no value is present for PartOfSince, not even an explicit nil
### GetPartOfSincePrecision

`func (o *CompaniesRow) GetPartOfSincePrecision() string`

GetPartOfSincePrecision returns the PartOfSincePrecision field if non-nil, zero value otherwise.

### GetPartOfSincePrecisionOk

`func (o *CompaniesRow) GetPartOfSincePrecisionOk() (*string, bool)`

GetPartOfSincePrecisionOk returns a tuple with the PartOfSincePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOfSincePrecision

`func (o *CompaniesRow) SetPartOfSincePrecision(v string)`

SetPartOfSincePrecision sets PartOfSincePrecision field to given value.

### HasPartOfSincePrecision

`func (o *CompaniesRow) HasPartOfSincePrecision() bool`

HasPartOfSincePrecision returns a boolean if a field has been set.

### SetPartOfSincePrecisionNil

`func (o *CompaniesRow) SetPartOfSincePrecisionNil(b bool)`

 SetPartOfSincePrecisionNil sets the value for PartOfSincePrecision to be an explicit nil

### UnsetPartOfSincePrecision
`func (o *CompaniesRow) UnsetPartOfSincePrecision()`

UnsetPartOfSincePrecision ensures that no value is present for PartOfSincePrecision, not even an explicit nil
### GetGroupId

`func (o *CompaniesRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompaniesRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompaniesRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompaniesRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompaniesRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompaniesRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompaniesRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompaniesRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompaniesRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompaniesRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompaniesRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompaniesRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetRollupDepth

`func (o *CompaniesRow) GetRollupDepth() int32`

GetRollupDepth returns the RollupDepth field if non-nil, zero value otherwise.

### GetRollupDepthOk

`func (o *CompaniesRow) GetRollupDepthOk() (*int32, bool)`

GetRollupDepthOk returns a tuple with the RollupDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupDepth

`func (o *CompaniesRow) SetRollupDepth(v int32)`

SetRollupDepth sets RollupDepth field to given value.

### HasRollupDepth

`func (o *CompaniesRow) HasRollupDepth() bool`

HasRollupDepth returns a boolean if a field has been set.

### SetRollupDepthNil

`func (o *CompaniesRow) SetRollupDepthNil(b bool)`

 SetRollupDepthNil sets the value for RollupDepth to be an explicit nil

### UnsetRollupDepth
`func (o *CompaniesRow) UnsetRollupDepth()`

UnsetRollupDepth ensures that no value is present for RollupDepth, not even an explicit nil
### GetNLegalEntities

`func (o *CompaniesRow) GetNLegalEntities() int64`

GetNLegalEntities returns the NLegalEntities field if non-nil, zero value otherwise.

### GetNLegalEntitiesOk

`func (o *CompaniesRow) GetNLegalEntitiesOk() (*int64, bool)`

GetNLegalEntitiesOk returns a tuple with the NLegalEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNLegalEntities

`func (o *CompaniesRow) SetNLegalEntities(v int64)`

SetNLegalEntities sets NLegalEntities field to given value.

### HasNLegalEntities

`func (o *CompaniesRow) HasNLegalEntities() bool`

HasNLegalEntities returns a boolean if a field has been set.

### SetNLegalEntitiesNil

`func (o *CompaniesRow) SetNLegalEntitiesNil(b bool)`

 SetNLegalEntitiesNil sets the value for NLegalEntities to be an explicit nil

### UnsetNLegalEntities
`func (o *CompaniesRow) UnsetNLegalEntities()`

UnsetNLegalEntities ensures that no value is present for NLegalEntities, not even an explicit nil
### GetNMemberCompanies

`func (o *CompaniesRow) GetNMemberCompanies() int64`

GetNMemberCompanies returns the NMemberCompanies field if non-nil, zero value otherwise.

### GetNMemberCompaniesOk

`func (o *CompaniesRow) GetNMemberCompaniesOk() (*int64, bool)`

GetNMemberCompaniesOk returns a tuple with the NMemberCompanies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNMemberCompanies

`func (o *CompaniesRow) SetNMemberCompanies(v int64)`

SetNMemberCompanies sets NMemberCompanies field to given value.

### HasNMemberCompanies

`func (o *CompaniesRow) HasNMemberCompanies() bool`

HasNMemberCompanies returns a boolean if a field has been set.

### SetNMemberCompaniesNil

`func (o *CompaniesRow) SetNMemberCompaniesNil(b bool)`

 SetNMemberCompaniesNil sets the value for NMemberCompanies to be an explicit nil

### UnsetNMemberCompanies
`func (o *CompaniesRow) UnsetNMemberCompanies()`

UnsetNMemberCompanies ensures that no value is present for NMemberCompanies, not even an explicit nil
### GetNumLeadTrials

`func (o *CompaniesRow) GetNumLeadTrials() int64`

GetNumLeadTrials returns the NumLeadTrials field if non-nil, zero value otherwise.

### GetNumLeadTrialsOk

`func (o *CompaniesRow) GetNumLeadTrialsOk() (*int64, bool)`

GetNumLeadTrialsOk returns a tuple with the NumLeadTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLeadTrials

`func (o *CompaniesRow) SetNumLeadTrials(v int64)`

SetNumLeadTrials sets NumLeadTrials field to given value.

### HasNumLeadTrials

`func (o *CompaniesRow) HasNumLeadTrials() bool`

HasNumLeadTrials returns a boolean if a field has been set.

### SetNumLeadTrialsNil

`func (o *CompaniesRow) SetNumLeadTrialsNil(b bool)`

 SetNumLeadTrialsNil sets the value for NumLeadTrials to be an explicit nil

### UnsetNumLeadTrials
`func (o *CompaniesRow) UnsetNumLeadTrials()`

UnsetNumLeadTrials ensures that no value is present for NumLeadTrials, not even an explicit nil
### GetNumCollaboratorTrials

`func (o *CompaniesRow) GetNumCollaboratorTrials() int64`

GetNumCollaboratorTrials returns the NumCollaboratorTrials field if non-nil, zero value otherwise.

### GetNumCollaboratorTrialsOk

`func (o *CompaniesRow) GetNumCollaboratorTrialsOk() (*int64, bool)`

GetNumCollaboratorTrialsOk returns a tuple with the NumCollaboratorTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCollaboratorTrials

`func (o *CompaniesRow) SetNumCollaboratorTrials(v int64)`

SetNumCollaboratorTrials sets NumCollaboratorTrials field to given value.

### HasNumCollaboratorTrials

`func (o *CompaniesRow) HasNumCollaboratorTrials() bool`

HasNumCollaboratorTrials returns a boolean if a field has been set.

### SetNumCollaboratorTrialsNil

`func (o *CompaniesRow) SetNumCollaboratorTrialsNil(b bool)`

 SetNumCollaboratorTrialsNil sets the value for NumCollaboratorTrials to be an explicit nil

### UnsetNumCollaboratorTrials
`func (o *CompaniesRow) UnsetNumCollaboratorTrials()`

UnsetNumCollaboratorTrials ensures that no value is present for NumCollaboratorTrials, not even an explicit nil
### GetNumTrials

`func (o *CompaniesRow) GetNumTrials() int64`

GetNumTrials returns the NumTrials field if non-nil, zero value otherwise.

### GetNumTrialsOk

`func (o *CompaniesRow) GetNumTrialsOk() (*int64, bool)`

GetNumTrialsOk returns a tuple with the NumTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTrials

`func (o *CompaniesRow) SetNumTrials(v int64)`

SetNumTrials sets NumTrials field to given value.

### HasNumTrials

`func (o *CompaniesRow) HasNumTrials() bool`

HasNumTrials returns a boolean if a field has been set.

### SetNumTrialsNil

`func (o *CompaniesRow) SetNumTrialsNil(b bool)`

 SetNumTrialsNil sets the value for NumTrials to be an explicit nil

### UnsetNumTrials
`func (o *CompaniesRow) UnsetNumTrials()`

UnsetNumTrials ensures that no value is present for NumTrials, not even an explicit nil
### GetTotalLeadEnrollment

`func (o *CompaniesRow) GetTotalLeadEnrollment() int64`

GetTotalLeadEnrollment returns the TotalLeadEnrollment field if non-nil, zero value otherwise.

### GetTotalLeadEnrollmentOk

`func (o *CompaniesRow) GetTotalLeadEnrollmentOk() (*int64, bool)`

GetTotalLeadEnrollmentOk returns a tuple with the TotalLeadEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLeadEnrollment

`func (o *CompaniesRow) SetTotalLeadEnrollment(v int64)`

SetTotalLeadEnrollment sets TotalLeadEnrollment field to given value.

### HasTotalLeadEnrollment

`func (o *CompaniesRow) HasTotalLeadEnrollment() bool`

HasTotalLeadEnrollment returns a boolean if a field has been set.

### SetTotalLeadEnrollmentNil

`func (o *CompaniesRow) SetTotalLeadEnrollmentNil(b bool)`

 SetTotalLeadEnrollmentNil sets the value for TotalLeadEnrollment to be an explicit nil

### UnsetTotalLeadEnrollment
`func (o *CompaniesRow) UnsetTotalLeadEnrollment()`

UnsetTotalLeadEnrollment ensures that no value is present for TotalLeadEnrollment, not even an explicit nil
### GetNumRecentLeadTrials

`func (o *CompaniesRow) GetNumRecentLeadTrials() int64`

GetNumRecentLeadTrials returns the NumRecentLeadTrials field if non-nil, zero value otherwise.

### GetNumRecentLeadTrialsOk

`func (o *CompaniesRow) GetNumRecentLeadTrialsOk() (*int64, bool)`

GetNumRecentLeadTrialsOk returns a tuple with the NumRecentLeadTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRecentLeadTrials

`func (o *CompaniesRow) SetNumRecentLeadTrials(v int64)`

SetNumRecentLeadTrials sets NumRecentLeadTrials field to given value.

### HasNumRecentLeadTrials

`func (o *CompaniesRow) HasNumRecentLeadTrials() bool`

HasNumRecentLeadTrials returns a boolean if a field has been set.

### SetNumRecentLeadTrialsNil

`func (o *CompaniesRow) SetNumRecentLeadTrialsNil(b bool)`

 SetNumRecentLeadTrialsNil sets the value for NumRecentLeadTrials to be an explicit nil

### UnsetNumRecentLeadTrials
`func (o *CompaniesRow) UnsetNumRecentLeadTrials()`

UnsetNumRecentLeadTrials ensures that no value is present for NumRecentLeadTrials, not even an explicit nil
### GetActivityScore

`func (o *CompaniesRow) GetActivityScore() float32`

GetActivityScore returns the ActivityScore field if non-nil, zero value otherwise.

### GetActivityScoreOk

`func (o *CompaniesRow) GetActivityScoreOk() (*float32, bool)`

GetActivityScoreOk returns a tuple with the ActivityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityScore

`func (o *CompaniesRow) SetActivityScore(v float32)`

SetActivityScore sets ActivityScore field to given value.

### HasActivityScore

`func (o *CompaniesRow) HasActivityScore() bool`

HasActivityScore returns a boolean if a field has been set.

### SetActivityScoreNil

`func (o *CompaniesRow) SetActivityScoreNil(b bool)`

 SetActivityScoreNil sets the value for ActivityScore to be an explicit nil

### UnsetActivityScore
`func (o *CompaniesRow) UnsetActivityScore()`

UnsetActivityScore ensures that no value is present for ActivityScore, not even an explicit nil
### GetIsAdmitted

`func (o *CompaniesRow) GetIsAdmitted() bool`

GetIsAdmitted returns the IsAdmitted field if non-nil, zero value otherwise.

### GetIsAdmittedOk

`func (o *CompaniesRow) GetIsAdmittedOk() (*bool, bool)`

GetIsAdmittedOk returns a tuple with the IsAdmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmitted

`func (o *CompaniesRow) SetIsAdmitted(v bool)`

SetIsAdmitted sets IsAdmitted field to given value.

### HasIsAdmitted

`func (o *CompaniesRow) HasIsAdmitted() bool`

HasIsAdmitted returns a boolean if a field has been set.

### SetIsAdmittedNil

`func (o *CompaniesRow) SetIsAdmittedNil(b bool)`

 SetIsAdmittedNil sets the value for IsAdmitted to be an explicit nil

### UnsetIsAdmitted
`func (o *CompaniesRow) UnsetIsAdmitted()`

UnsetIsAdmitted ensures that no value is present for IsAdmitted, not even an explicit nil
### GetHasPublicPage

`func (o *CompaniesRow) GetHasPublicPage() bool`

GetHasPublicPage returns the HasPublicPage field if non-nil, zero value otherwise.

### GetHasPublicPageOk

`func (o *CompaniesRow) GetHasPublicPageOk() (*bool, bool)`

GetHasPublicPageOk returns a tuple with the HasPublicPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPublicPage

`func (o *CompaniesRow) SetHasPublicPage(v bool)`

SetHasPublicPage sets HasPublicPage field to given value.

### HasHasPublicPage

`func (o *CompaniesRow) HasHasPublicPage() bool`

HasHasPublicPage returns a boolean if a field has been set.

### SetHasPublicPageNil

`func (o *CompaniesRow) SetHasPublicPageNil(b bool)`

 SetHasPublicPageNil sets the value for HasPublicPage to be an explicit nil

### UnsetHasPublicPage
`func (o *CompaniesRow) UnsetHasPublicPage()`

UnsetHasPublicPage ensures that no value is present for HasPublicPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


