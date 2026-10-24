# InvestigatorTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestigatorPersonId** | Pointer to **NullableString** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] 
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**Roles** | Pointer to **[]string** | The roles the person holds on this trial as registered. | [optional] 
**IsPrincipalInvestigator** | Pointer to **NullableBool** | True when the person is named principal investigator on this trial. | [optional] 
**TrialTitle** | Pointer to **NullableString** | The trial&#39;s brief title. | [optional] 
**PhaseLabel** | Pointer to **NullableString** | The trial phase as a display label. | [optional] 
**OverallStatus** | Pointer to **NullableString** | The registry&#39;s overall recruitment status. | [optional] 
**StudyType** | Pointer to **NullableString** | The registry&#39;s study type. | [optional] 
**StartDate** | Pointer to **NullableString** | Trial start date as registered. | [optional] 
**StartYear** | Pointer to **NullableInt32** | Calendar year of the start date. | [optional] 
**CompletionDate** | Pointer to **NullableString** | Trial completion date as registered. | [optional] 
**Enrollment** | Pointer to **NullableInt32** | Enrollment as registered. | [optional] 
**NSites** | Pointer to **NullableInt64** | The trial&#39;s distinct geocoded sites, the count the people page shows (0 when none is geocoded). | [optional] 
**SingleSiteId** | Pointer to **NullableString** | The site key when the trial lists exactly one site (null otherwise); joins the sites dataset. | [optional] 
**SingleSiteName** | Pointer to **NullableString** | The single site&#39;s facility name as the registry lists it. | [optional] 
**SingleSiteCity** | Pointer to **NullableString** | The single site&#39;s city. | [optional] 
**SingleSiteState** | Pointer to **NullableString** | The single site&#39;s state or province. | [optional] 
**SingleSiteCountry** | Pointer to **NullableString** | The single site&#39;s country. | [optional] 
**SponsorName** | Pointer to **NullableString** | The trial&#39;s lead sponsor as filed. | [optional] 
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the lead sponsor resolved to (null when unresolved). | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the company the lead sponsor belongs to (null when unresolved). | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the sponsor&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the sponsor&#39;s group. | [optional] 

## Methods

### NewInvestigatorTrialsRow

`func NewInvestigatorTrialsRow() *InvestigatorTrialsRow`

NewInvestigatorTrialsRow instantiates a new InvestigatorTrialsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigatorTrialsRowWithDefaults

`func NewInvestigatorTrialsRowWithDefaults() *InvestigatorTrialsRow`

NewInvestigatorTrialsRowWithDefaults instantiates a new InvestigatorTrialsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestigatorPersonId

`func (o *InvestigatorTrialsRow) GetInvestigatorPersonId() string`

GetInvestigatorPersonId returns the InvestigatorPersonId field if non-nil, zero value otherwise.

### GetInvestigatorPersonIdOk

`func (o *InvestigatorTrialsRow) GetInvestigatorPersonIdOk() (*string, bool)`

GetInvestigatorPersonIdOk returns a tuple with the InvestigatorPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonId

`func (o *InvestigatorTrialsRow) SetInvestigatorPersonId(v string)`

SetInvestigatorPersonId sets InvestigatorPersonId field to given value.

### HasInvestigatorPersonId

`func (o *InvestigatorTrialsRow) HasInvestigatorPersonId() bool`

HasInvestigatorPersonId returns a boolean if a field has been set.

### SetInvestigatorPersonIdNil

`func (o *InvestigatorTrialsRow) SetInvestigatorPersonIdNil(b bool)`

 SetInvestigatorPersonIdNil sets the value for InvestigatorPersonId to be an explicit nil

### UnsetInvestigatorPersonId
`func (o *InvestigatorTrialsRow) UnsetInvestigatorPersonId()`

UnsetInvestigatorPersonId ensures that no value is present for InvestigatorPersonId, not even an explicit nil
### GetTrialId

`func (o *InvestigatorTrialsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *InvestigatorTrialsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *InvestigatorTrialsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *InvestigatorTrialsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *InvestigatorTrialsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *InvestigatorTrialsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetRoles

`func (o *InvestigatorTrialsRow) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *InvestigatorTrialsRow) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *InvestigatorTrialsRow) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *InvestigatorTrialsRow) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *InvestigatorTrialsRow) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *InvestigatorTrialsRow) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetIsPrincipalInvestigator

`func (o *InvestigatorTrialsRow) GetIsPrincipalInvestigator() bool`

GetIsPrincipalInvestigator returns the IsPrincipalInvestigator field if non-nil, zero value otherwise.

### GetIsPrincipalInvestigatorOk

`func (o *InvestigatorTrialsRow) GetIsPrincipalInvestigatorOk() (*bool, bool)`

GetIsPrincipalInvestigatorOk returns a tuple with the IsPrincipalInvestigator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrincipalInvestigator

`func (o *InvestigatorTrialsRow) SetIsPrincipalInvestigator(v bool)`

SetIsPrincipalInvestigator sets IsPrincipalInvestigator field to given value.

### HasIsPrincipalInvestigator

`func (o *InvestigatorTrialsRow) HasIsPrincipalInvestigator() bool`

HasIsPrincipalInvestigator returns a boolean if a field has been set.

### SetIsPrincipalInvestigatorNil

`func (o *InvestigatorTrialsRow) SetIsPrincipalInvestigatorNil(b bool)`

 SetIsPrincipalInvestigatorNil sets the value for IsPrincipalInvestigator to be an explicit nil

### UnsetIsPrincipalInvestigator
`func (o *InvestigatorTrialsRow) UnsetIsPrincipalInvestigator()`

UnsetIsPrincipalInvestigator ensures that no value is present for IsPrincipalInvestigator, not even an explicit nil
### GetTrialTitle

`func (o *InvestigatorTrialsRow) GetTrialTitle() string`

GetTrialTitle returns the TrialTitle field if non-nil, zero value otherwise.

### GetTrialTitleOk

`func (o *InvestigatorTrialsRow) GetTrialTitleOk() (*string, bool)`

GetTrialTitleOk returns a tuple with the TrialTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialTitle

`func (o *InvestigatorTrialsRow) SetTrialTitle(v string)`

SetTrialTitle sets TrialTitle field to given value.

### HasTrialTitle

`func (o *InvestigatorTrialsRow) HasTrialTitle() bool`

HasTrialTitle returns a boolean if a field has been set.

### SetTrialTitleNil

`func (o *InvestigatorTrialsRow) SetTrialTitleNil(b bool)`

 SetTrialTitleNil sets the value for TrialTitle to be an explicit nil

### UnsetTrialTitle
`func (o *InvestigatorTrialsRow) UnsetTrialTitle()`

UnsetTrialTitle ensures that no value is present for TrialTitle, not even an explicit nil
### GetPhaseLabel

`func (o *InvestigatorTrialsRow) GetPhaseLabel() string`

GetPhaseLabel returns the PhaseLabel field if non-nil, zero value otherwise.

### GetPhaseLabelOk

`func (o *InvestigatorTrialsRow) GetPhaseLabelOk() (*string, bool)`

GetPhaseLabelOk returns a tuple with the PhaseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseLabel

`func (o *InvestigatorTrialsRow) SetPhaseLabel(v string)`

SetPhaseLabel sets PhaseLabel field to given value.

### HasPhaseLabel

`func (o *InvestigatorTrialsRow) HasPhaseLabel() bool`

HasPhaseLabel returns a boolean if a field has been set.

### SetPhaseLabelNil

`func (o *InvestigatorTrialsRow) SetPhaseLabelNil(b bool)`

 SetPhaseLabelNil sets the value for PhaseLabel to be an explicit nil

### UnsetPhaseLabel
`func (o *InvestigatorTrialsRow) UnsetPhaseLabel()`

UnsetPhaseLabel ensures that no value is present for PhaseLabel, not even an explicit nil
### GetOverallStatus

`func (o *InvestigatorTrialsRow) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *InvestigatorTrialsRow) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *InvestigatorTrialsRow) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *InvestigatorTrialsRow) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### SetOverallStatusNil

`func (o *InvestigatorTrialsRow) SetOverallStatusNil(b bool)`

 SetOverallStatusNil sets the value for OverallStatus to be an explicit nil

### UnsetOverallStatus
`func (o *InvestigatorTrialsRow) UnsetOverallStatus()`

UnsetOverallStatus ensures that no value is present for OverallStatus, not even an explicit nil
### GetStudyType

`func (o *InvestigatorTrialsRow) GetStudyType() string`

GetStudyType returns the StudyType field if non-nil, zero value otherwise.

### GetStudyTypeOk

`func (o *InvestigatorTrialsRow) GetStudyTypeOk() (*string, bool)`

GetStudyTypeOk returns a tuple with the StudyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyType

`func (o *InvestigatorTrialsRow) SetStudyType(v string)`

SetStudyType sets StudyType field to given value.

### HasStudyType

`func (o *InvestigatorTrialsRow) HasStudyType() bool`

HasStudyType returns a boolean if a field has been set.

### SetStudyTypeNil

`func (o *InvestigatorTrialsRow) SetStudyTypeNil(b bool)`

 SetStudyTypeNil sets the value for StudyType to be an explicit nil

### UnsetStudyType
`func (o *InvestigatorTrialsRow) UnsetStudyType()`

UnsetStudyType ensures that no value is present for StudyType, not even an explicit nil
### GetStartDate

`func (o *InvestigatorTrialsRow) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InvestigatorTrialsRow) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InvestigatorTrialsRow) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InvestigatorTrialsRow) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *InvestigatorTrialsRow) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *InvestigatorTrialsRow) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetStartYear

`func (o *InvestigatorTrialsRow) GetStartYear() int32`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *InvestigatorTrialsRow) GetStartYearOk() (*int32, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *InvestigatorTrialsRow) SetStartYear(v int32)`

SetStartYear sets StartYear field to given value.

### HasStartYear

`func (o *InvestigatorTrialsRow) HasStartYear() bool`

HasStartYear returns a boolean if a field has been set.

### SetStartYearNil

`func (o *InvestigatorTrialsRow) SetStartYearNil(b bool)`

 SetStartYearNil sets the value for StartYear to be an explicit nil

### UnsetStartYear
`func (o *InvestigatorTrialsRow) UnsetStartYear()`

UnsetStartYear ensures that no value is present for StartYear, not even an explicit nil
### GetCompletionDate

`func (o *InvestigatorTrialsRow) GetCompletionDate() string`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *InvestigatorTrialsRow) GetCompletionDateOk() (*string, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *InvestigatorTrialsRow) SetCompletionDate(v string)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *InvestigatorTrialsRow) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### SetCompletionDateNil

`func (o *InvestigatorTrialsRow) SetCompletionDateNil(b bool)`

 SetCompletionDateNil sets the value for CompletionDate to be an explicit nil

### UnsetCompletionDate
`func (o *InvestigatorTrialsRow) UnsetCompletionDate()`

UnsetCompletionDate ensures that no value is present for CompletionDate, not even an explicit nil
### GetEnrollment

`func (o *InvestigatorTrialsRow) GetEnrollment() int32`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *InvestigatorTrialsRow) GetEnrollmentOk() (*int32, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *InvestigatorTrialsRow) SetEnrollment(v int32)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *InvestigatorTrialsRow) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### SetEnrollmentNil

`func (o *InvestigatorTrialsRow) SetEnrollmentNil(b bool)`

 SetEnrollmentNil sets the value for Enrollment to be an explicit nil

### UnsetEnrollment
`func (o *InvestigatorTrialsRow) UnsetEnrollment()`

UnsetEnrollment ensures that no value is present for Enrollment, not even an explicit nil
### GetNSites

`func (o *InvestigatorTrialsRow) GetNSites() int64`

GetNSites returns the NSites field if non-nil, zero value otherwise.

### GetNSitesOk

`func (o *InvestigatorTrialsRow) GetNSitesOk() (*int64, bool)`

GetNSitesOk returns a tuple with the NSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSites

`func (o *InvestigatorTrialsRow) SetNSites(v int64)`

SetNSites sets NSites field to given value.

### HasNSites

`func (o *InvestigatorTrialsRow) HasNSites() bool`

HasNSites returns a boolean if a field has been set.

### SetNSitesNil

`func (o *InvestigatorTrialsRow) SetNSitesNil(b bool)`

 SetNSitesNil sets the value for NSites to be an explicit nil

### UnsetNSites
`func (o *InvestigatorTrialsRow) UnsetNSites()`

UnsetNSites ensures that no value is present for NSites, not even an explicit nil
### GetSingleSiteId

`func (o *InvestigatorTrialsRow) GetSingleSiteId() string`

GetSingleSiteId returns the SingleSiteId field if non-nil, zero value otherwise.

### GetSingleSiteIdOk

`func (o *InvestigatorTrialsRow) GetSingleSiteIdOk() (*string, bool)`

GetSingleSiteIdOk returns a tuple with the SingleSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSiteId

`func (o *InvestigatorTrialsRow) SetSingleSiteId(v string)`

SetSingleSiteId sets SingleSiteId field to given value.

### HasSingleSiteId

`func (o *InvestigatorTrialsRow) HasSingleSiteId() bool`

HasSingleSiteId returns a boolean if a field has been set.

### SetSingleSiteIdNil

`func (o *InvestigatorTrialsRow) SetSingleSiteIdNil(b bool)`

 SetSingleSiteIdNil sets the value for SingleSiteId to be an explicit nil

### UnsetSingleSiteId
`func (o *InvestigatorTrialsRow) UnsetSingleSiteId()`

UnsetSingleSiteId ensures that no value is present for SingleSiteId, not even an explicit nil
### GetSingleSiteName

`func (o *InvestigatorTrialsRow) GetSingleSiteName() string`

GetSingleSiteName returns the SingleSiteName field if non-nil, zero value otherwise.

### GetSingleSiteNameOk

`func (o *InvestigatorTrialsRow) GetSingleSiteNameOk() (*string, bool)`

GetSingleSiteNameOk returns a tuple with the SingleSiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSiteName

`func (o *InvestigatorTrialsRow) SetSingleSiteName(v string)`

SetSingleSiteName sets SingleSiteName field to given value.

### HasSingleSiteName

`func (o *InvestigatorTrialsRow) HasSingleSiteName() bool`

HasSingleSiteName returns a boolean if a field has been set.

### SetSingleSiteNameNil

`func (o *InvestigatorTrialsRow) SetSingleSiteNameNil(b bool)`

 SetSingleSiteNameNil sets the value for SingleSiteName to be an explicit nil

### UnsetSingleSiteName
`func (o *InvestigatorTrialsRow) UnsetSingleSiteName()`

UnsetSingleSiteName ensures that no value is present for SingleSiteName, not even an explicit nil
### GetSingleSiteCity

`func (o *InvestigatorTrialsRow) GetSingleSiteCity() string`

GetSingleSiteCity returns the SingleSiteCity field if non-nil, zero value otherwise.

### GetSingleSiteCityOk

`func (o *InvestigatorTrialsRow) GetSingleSiteCityOk() (*string, bool)`

GetSingleSiteCityOk returns a tuple with the SingleSiteCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSiteCity

`func (o *InvestigatorTrialsRow) SetSingleSiteCity(v string)`

SetSingleSiteCity sets SingleSiteCity field to given value.

### HasSingleSiteCity

`func (o *InvestigatorTrialsRow) HasSingleSiteCity() bool`

HasSingleSiteCity returns a boolean if a field has been set.

### SetSingleSiteCityNil

`func (o *InvestigatorTrialsRow) SetSingleSiteCityNil(b bool)`

 SetSingleSiteCityNil sets the value for SingleSiteCity to be an explicit nil

### UnsetSingleSiteCity
`func (o *InvestigatorTrialsRow) UnsetSingleSiteCity()`

UnsetSingleSiteCity ensures that no value is present for SingleSiteCity, not even an explicit nil
### GetSingleSiteState

`func (o *InvestigatorTrialsRow) GetSingleSiteState() string`

GetSingleSiteState returns the SingleSiteState field if non-nil, zero value otherwise.

### GetSingleSiteStateOk

`func (o *InvestigatorTrialsRow) GetSingleSiteStateOk() (*string, bool)`

GetSingleSiteStateOk returns a tuple with the SingleSiteState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSiteState

`func (o *InvestigatorTrialsRow) SetSingleSiteState(v string)`

SetSingleSiteState sets SingleSiteState field to given value.

### HasSingleSiteState

`func (o *InvestigatorTrialsRow) HasSingleSiteState() bool`

HasSingleSiteState returns a boolean if a field has been set.

### SetSingleSiteStateNil

`func (o *InvestigatorTrialsRow) SetSingleSiteStateNil(b bool)`

 SetSingleSiteStateNil sets the value for SingleSiteState to be an explicit nil

### UnsetSingleSiteState
`func (o *InvestigatorTrialsRow) UnsetSingleSiteState()`

UnsetSingleSiteState ensures that no value is present for SingleSiteState, not even an explicit nil
### GetSingleSiteCountry

`func (o *InvestigatorTrialsRow) GetSingleSiteCountry() string`

GetSingleSiteCountry returns the SingleSiteCountry field if non-nil, zero value otherwise.

### GetSingleSiteCountryOk

`func (o *InvestigatorTrialsRow) GetSingleSiteCountryOk() (*string, bool)`

GetSingleSiteCountryOk returns a tuple with the SingleSiteCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSiteCountry

`func (o *InvestigatorTrialsRow) SetSingleSiteCountry(v string)`

SetSingleSiteCountry sets SingleSiteCountry field to given value.

### HasSingleSiteCountry

`func (o *InvestigatorTrialsRow) HasSingleSiteCountry() bool`

HasSingleSiteCountry returns a boolean if a field has been set.

### SetSingleSiteCountryNil

`func (o *InvestigatorTrialsRow) SetSingleSiteCountryNil(b bool)`

 SetSingleSiteCountryNil sets the value for SingleSiteCountry to be an explicit nil

### UnsetSingleSiteCountry
`func (o *InvestigatorTrialsRow) UnsetSingleSiteCountry()`

UnsetSingleSiteCountry ensures that no value is present for SingleSiteCountry, not even an explicit nil
### GetSponsorName

`func (o *InvestigatorTrialsRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *InvestigatorTrialsRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *InvestigatorTrialsRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *InvestigatorTrialsRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *InvestigatorTrialsRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *InvestigatorTrialsRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorOrgId

`func (o *InvestigatorTrialsRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *InvestigatorTrialsRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *InvestigatorTrialsRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *InvestigatorTrialsRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *InvestigatorTrialsRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *InvestigatorTrialsRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorCompanyId

`func (o *InvestigatorTrialsRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *InvestigatorTrialsRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *InvestigatorTrialsRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *InvestigatorTrialsRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *InvestigatorTrialsRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *InvestigatorTrialsRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *InvestigatorTrialsRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *InvestigatorTrialsRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *InvestigatorTrialsRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *InvestigatorTrialsRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *InvestigatorTrialsRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *InvestigatorTrialsRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *InvestigatorTrialsRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *InvestigatorTrialsRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *InvestigatorTrialsRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *InvestigatorTrialsRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *InvestigatorTrialsRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *InvestigatorTrialsRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *InvestigatorTrialsRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *InvestigatorTrialsRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *InvestigatorTrialsRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *InvestigatorTrialsRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *InvestigatorTrialsRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *InvestigatorTrialsRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


