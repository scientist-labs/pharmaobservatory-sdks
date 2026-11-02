# CompanyCosponsorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | Company level: the company the row is about; the entry key for the companies object type. | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**RowLevel** | Pointer to **NullableString** | root: the co-sponsor&#39;s ultimate parent company, with the distinct trials any company of that root shares with this company. member: a company under that root that itself co-sponsored, with its own count. Filter on one level; the two do not add. | [optional] 
**CosponsorCompanyId** | Pointer to **NullableString** | Company level: the co-sponsoring company. At row_level root this is the root company; at member it is the member company itself. | [optional] 
**CosponsorCompanyName** | Pointer to **NullableString** | Display name of the co-sponsoring company. | [optional] 
**CosponsorGroupId** | Pointer to **NullableString** | Group level: the co-sponsor&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**CosponsorGroupName** | Pointer to **NullableString** | Display name of the co-sponsor&#39;s group. | [optional] 
**CosponsorRootId** | Pointer to **NullableString** | The ultimate parent company of the co-sponsor, the ladder company&#39;s &#x60;coalesce(ultimate_parent_company_id, company_id)&#x60;. Equal to cosponsor_company_id at row_level root, and the root a member row sits under. | [optional] 
**CosponsorRank** | Pointer to **NullableInt32** | The rank of the ROOT among this company&#39;s co-sponsor roots by distinct shared trials (1 is the most; ties broken by root id). The same on a root row and on each of its member rows. | [optional] 
**MemberRank** | Pointer to **NullableInt32** | 0 on a root row; 1 and up on the member rows of that root, by the member&#39;s own shared trials. Order by cosponsor_rank then member_rank to list each root followed by its members. | [optional] 
**NMembers** | Pointer to **NullableInt32** | Member companies of this root that co-sponsored with this company, before the member cut. The same on a root row and its member rows. | [optional] 
**NCandidates** | Pointer to **NullableInt32** | Distinct co-sponsor roots for this company before the top-25 cut; the same on every row of a company. | [optional] 
**NSharedTrials** | Pointer to **NullableInt32** | Distinct trials that name both. At row_level root, the distinct trials of any company under the root (never a sum over members); at member, that company&#39;s own trials. | [optional] 
**NTrialsCompanyLead** | Pointer to **NullableInt32** | Of the shared trials, those where an organization record of this company holds the ClinicalTrials.gov lead sponsor role. At most n_shared_trials; it may overlap n_trials_cosponsor_lead. | [optional] 
**NTrialsCosponsorLead** | Pointer to **NullableInt32** | Of the shared trials, those where the co-sponsor&#39;s organization record holds the lead sponsor role. At most n_shared_trials. | [optional] 
**FirstSharedStartYear** | Pointer to **NullableInt32** | The earliest registered start year over the shared trials; null when none has a plausible year. | [optional] 
**LastSharedStartYear** | Pointer to **NullableInt32** | The latest registered start year over the shared trials; null when none has a plausible year. | [optional] 
**NSharedConditions** | Pointer to **NullableInt32** | Distinct specific MeSH disease areas over the shared trials. | [optional] 
**TopSharedConditions** | Pointer to **[]string** | Up to 5 MeSH disease names by shared-trial count then name. An empty list when none. | [optional] 

## Methods

### NewCompanyCosponsorsRow

`func NewCompanyCosponsorsRow() *CompanyCosponsorsRow`

NewCompanyCosponsorsRow instantiates a new CompanyCosponsorsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyCosponsorsRowWithDefaults

`func NewCompanyCosponsorsRowWithDefaults() *CompanyCosponsorsRow`

NewCompanyCosponsorsRowWithDefaults instantiates a new CompanyCosponsorsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyCosponsorsRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyCosponsorsRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyCosponsorsRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyCosponsorsRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyCosponsorsRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyCosponsorsRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyCosponsorsRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyCosponsorsRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyCosponsorsRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyCosponsorsRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyCosponsorsRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyCosponsorsRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyCosponsorsRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyCosponsorsRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyCosponsorsRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyCosponsorsRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyCosponsorsRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyCosponsorsRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyCosponsorsRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyCosponsorsRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyCosponsorsRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyCosponsorsRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyCosponsorsRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyCosponsorsRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetRowLevel

`func (o *CompanyCosponsorsRow) GetRowLevel() string`

GetRowLevel returns the RowLevel field if non-nil, zero value otherwise.

### GetRowLevelOk

`func (o *CompanyCosponsorsRow) GetRowLevelOk() (*string, bool)`

GetRowLevelOk returns a tuple with the RowLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowLevel

`func (o *CompanyCosponsorsRow) SetRowLevel(v string)`

SetRowLevel sets RowLevel field to given value.

### HasRowLevel

`func (o *CompanyCosponsorsRow) HasRowLevel() bool`

HasRowLevel returns a boolean if a field has been set.

### SetRowLevelNil

`func (o *CompanyCosponsorsRow) SetRowLevelNil(b bool)`

 SetRowLevelNil sets the value for RowLevel to be an explicit nil

### UnsetRowLevel
`func (o *CompanyCosponsorsRow) UnsetRowLevel()`

UnsetRowLevel ensures that no value is present for RowLevel, not even an explicit nil
### GetCosponsorCompanyId

`func (o *CompanyCosponsorsRow) GetCosponsorCompanyId() string`

GetCosponsorCompanyId returns the CosponsorCompanyId field if non-nil, zero value otherwise.

### GetCosponsorCompanyIdOk

`func (o *CompanyCosponsorsRow) GetCosponsorCompanyIdOk() (*string, bool)`

GetCosponsorCompanyIdOk returns a tuple with the CosponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosponsorCompanyId

`func (o *CompanyCosponsorsRow) SetCosponsorCompanyId(v string)`

SetCosponsorCompanyId sets CosponsorCompanyId field to given value.

### HasCosponsorCompanyId

`func (o *CompanyCosponsorsRow) HasCosponsorCompanyId() bool`

HasCosponsorCompanyId returns a boolean if a field has been set.

### SetCosponsorCompanyIdNil

`func (o *CompanyCosponsorsRow) SetCosponsorCompanyIdNil(b bool)`

 SetCosponsorCompanyIdNil sets the value for CosponsorCompanyId to be an explicit nil

### UnsetCosponsorCompanyId
`func (o *CompanyCosponsorsRow) UnsetCosponsorCompanyId()`

UnsetCosponsorCompanyId ensures that no value is present for CosponsorCompanyId, not even an explicit nil
### GetCosponsorCompanyName

`func (o *CompanyCosponsorsRow) GetCosponsorCompanyName() string`

GetCosponsorCompanyName returns the CosponsorCompanyName field if non-nil, zero value otherwise.

### GetCosponsorCompanyNameOk

`func (o *CompanyCosponsorsRow) GetCosponsorCompanyNameOk() (*string, bool)`

GetCosponsorCompanyNameOk returns a tuple with the CosponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosponsorCompanyName

`func (o *CompanyCosponsorsRow) SetCosponsorCompanyName(v string)`

SetCosponsorCompanyName sets CosponsorCompanyName field to given value.

### HasCosponsorCompanyName

`func (o *CompanyCosponsorsRow) HasCosponsorCompanyName() bool`

HasCosponsorCompanyName returns a boolean if a field has been set.

### SetCosponsorCompanyNameNil

`func (o *CompanyCosponsorsRow) SetCosponsorCompanyNameNil(b bool)`

 SetCosponsorCompanyNameNil sets the value for CosponsorCompanyName to be an explicit nil

### UnsetCosponsorCompanyName
`func (o *CompanyCosponsorsRow) UnsetCosponsorCompanyName()`

UnsetCosponsorCompanyName ensures that no value is present for CosponsorCompanyName, not even an explicit nil
### GetCosponsorGroupId

`func (o *CompanyCosponsorsRow) GetCosponsorGroupId() string`

GetCosponsorGroupId returns the CosponsorGroupId field if non-nil, zero value otherwise.

### GetCosponsorGroupIdOk

`func (o *CompanyCosponsorsRow) GetCosponsorGroupIdOk() (*string, bool)`

GetCosponsorGroupIdOk returns a tuple with the CosponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosponsorGroupId

`func (o *CompanyCosponsorsRow) SetCosponsorGroupId(v string)`

SetCosponsorGroupId sets CosponsorGroupId field to given value.

### HasCosponsorGroupId

`func (o *CompanyCosponsorsRow) HasCosponsorGroupId() bool`

HasCosponsorGroupId returns a boolean if a field has been set.

### SetCosponsorGroupIdNil

`func (o *CompanyCosponsorsRow) SetCosponsorGroupIdNil(b bool)`

 SetCosponsorGroupIdNil sets the value for CosponsorGroupId to be an explicit nil

### UnsetCosponsorGroupId
`func (o *CompanyCosponsorsRow) UnsetCosponsorGroupId()`

UnsetCosponsorGroupId ensures that no value is present for CosponsorGroupId, not even an explicit nil
### GetCosponsorGroupName

`func (o *CompanyCosponsorsRow) GetCosponsorGroupName() string`

GetCosponsorGroupName returns the CosponsorGroupName field if non-nil, zero value otherwise.

### GetCosponsorGroupNameOk

`func (o *CompanyCosponsorsRow) GetCosponsorGroupNameOk() (*string, bool)`

GetCosponsorGroupNameOk returns a tuple with the CosponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosponsorGroupName

`func (o *CompanyCosponsorsRow) SetCosponsorGroupName(v string)`

SetCosponsorGroupName sets CosponsorGroupName field to given value.

### HasCosponsorGroupName

`func (o *CompanyCosponsorsRow) HasCosponsorGroupName() bool`

HasCosponsorGroupName returns a boolean if a field has been set.

### SetCosponsorGroupNameNil

`func (o *CompanyCosponsorsRow) SetCosponsorGroupNameNil(b bool)`

 SetCosponsorGroupNameNil sets the value for CosponsorGroupName to be an explicit nil

### UnsetCosponsorGroupName
`func (o *CompanyCosponsorsRow) UnsetCosponsorGroupName()`

UnsetCosponsorGroupName ensures that no value is present for CosponsorGroupName, not even an explicit nil
### GetCosponsorRootId

`func (o *CompanyCosponsorsRow) GetCosponsorRootId() string`

GetCosponsorRootId returns the CosponsorRootId field if non-nil, zero value otherwise.

### GetCosponsorRootIdOk

`func (o *CompanyCosponsorsRow) GetCosponsorRootIdOk() (*string, bool)`

GetCosponsorRootIdOk returns a tuple with the CosponsorRootId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosponsorRootId

`func (o *CompanyCosponsorsRow) SetCosponsorRootId(v string)`

SetCosponsorRootId sets CosponsorRootId field to given value.

### HasCosponsorRootId

`func (o *CompanyCosponsorsRow) HasCosponsorRootId() bool`

HasCosponsorRootId returns a boolean if a field has been set.

### SetCosponsorRootIdNil

`func (o *CompanyCosponsorsRow) SetCosponsorRootIdNil(b bool)`

 SetCosponsorRootIdNil sets the value for CosponsorRootId to be an explicit nil

### UnsetCosponsorRootId
`func (o *CompanyCosponsorsRow) UnsetCosponsorRootId()`

UnsetCosponsorRootId ensures that no value is present for CosponsorRootId, not even an explicit nil
### GetCosponsorRank

`func (o *CompanyCosponsorsRow) GetCosponsorRank() int32`

GetCosponsorRank returns the CosponsorRank field if non-nil, zero value otherwise.

### GetCosponsorRankOk

`func (o *CompanyCosponsorsRow) GetCosponsorRankOk() (*int32, bool)`

GetCosponsorRankOk returns a tuple with the CosponsorRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosponsorRank

`func (o *CompanyCosponsorsRow) SetCosponsorRank(v int32)`

SetCosponsorRank sets CosponsorRank field to given value.

### HasCosponsorRank

`func (o *CompanyCosponsorsRow) HasCosponsorRank() bool`

HasCosponsorRank returns a boolean if a field has been set.

### SetCosponsorRankNil

`func (o *CompanyCosponsorsRow) SetCosponsorRankNil(b bool)`

 SetCosponsorRankNil sets the value for CosponsorRank to be an explicit nil

### UnsetCosponsorRank
`func (o *CompanyCosponsorsRow) UnsetCosponsorRank()`

UnsetCosponsorRank ensures that no value is present for CosponsorRank, not even an explicit nil
### GetMemberRank

`func (o *CompanyCosponsorsRow) GetMemberRank() int32`

GetMemberRank returns the MemberRank field if non-nil, zero value otherwise.

### GetMemberRankOk

`func (o *CompanyCosponsorsRow) GetMemberRankOk() (*int32, bool)`

GetMemberRankOk returns a tuple with the MemberRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberRank

`func (o *CompanyCosponsorsRow) SetMemberRank(v int32)`

SetMemberRank sets MemberRank field to given value.

### HasMemberRank

`func (o *CompanyCosponsorsRow) HasMemberRank() bool`

HasMemberRank returns a boolean if a field has been set.

### SetMemberRankNil

`func (o *CompanyCosponsorsRow) SetMemberRankNil(b bool)`

 SetMemberRankNil sets the value for MemberRank to be an explicit nil

### UnsetMemberRank
`func (o *CompanyCosponsorsRow) UnsetMemberRank()`

UnsetMemberRank ensures that no value is present for MemberRank, not even an explicit nil
### GetNMembers

`func (o *CompanyCosponsorsRow) GetNMembers() int32`

GetNMembers returns the NMembers field if non-nil, zero value otherwise.

### GetNMembersOk

`func (o *CompanyCosponsorsRow) GetNMembersOk() (*int32, bool)`

GetNMembersOk returns a tuple with the NMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNMembers

`func (o *CompanyCosponsorsRow) SetNMembers(v int32)`

SetNMembers sets NMembers field to given value.

### HasNMembers

`func (o *CompanyCosponsorsRow) HasNMembers() bool`

HasNMembers returns a boolean if a field has been set.

### SetNMembersNil

`func (o *CompanyCosponsorsRow) SetNMembersNil(b bool)`

 SetNMembersNil sets the value for NMembers to be an explicit nil

### UnsetNMembers
`func (o *CompanyCosponsorsRow) UnsetNMembers()`

UnsetNMembers ensures that no value is present for NMembers, not even an explicit nil
### GetNCandidates

`func (o *CompanyCosponsorsRow) GetNCandidates() int32`

GetNCandidates returns the NCandidates field if non-nil, zero value otherwise.

### GetNCandidatesOk

`func (o *CompanyCosponsorsRow) GetNCandidatesOk() (*int32, bool)`

GetNCandidatesOk returns a tuple with the NCandidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCandidates

`func (o *CompanyCosponsorsRow) SetNCandidates(v int32)`

SetNCandidates sets NCandidates field to given value.

### HasNCandidates

`func (o *CompanyCosponsorsRow) HasNCandidates() bool`

HasNCandidates returns a boolean if a field has been set.

### SetNCandidatesNil

`func (o *CompanyCosponsorsRow) SetNCandidatesNil(b bool)`

 SetNCandidatesNil sets the value for NCandidates to be an explicit nil

### UnsetNCandidates
`func (o *CompanyCosponsorsRow) UnsetNCandidates()`

UnsetNCandidates ensures that no value is present for NCandidates, not even an explicit nil
### GetNSharedTrials

`func (o *CompanyCosponsorsRow) GetNSharedTrials() int32`

GetNSharedTrials returns the NSharedTrials field if non-nil, zero value otherwise.

### GetNSharedTrialsOk

`func (o *CompanyCosponsorsRow) GetNSharedTrialsOk() (*int32, bool)`

GetNSharedTrialsOk returns a tuple with the NSharedTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSharedTrials

`func (o *CompanyCosponsorsRow) SetNSharedTrials(v int32)`

SetNSharedTrials sets NSharedTrials field to given value.

### HasNSharedTrials

`func (o *CompanyCosponsorsRow) HasNSharedTrials() bool`

HasNSharedTrials returns a boolean if a field has been set.

### SetNSharedTrialsNil

`func (o *CompanyCosponsorsRow) SetNSharedTrialsNil(b bool)`

 SetNSharedTrialsNil sets the value for NSharedTrials to be an explicit nil

### UnsetNSharedTrials
`func (o *CompanyCosponsorsRow) UnsetNSharedTrials()`

UnsetNSharedTrials ensures that no value is present for NSharedTrials, not even an explicit nil
### GetNTrialsCompanyLead

`func (o *CompanyCosponsorsRow) GetNTrialsCompanyLead() int32`

GetNTrialsCompanyLead returns the NTrialsCompanyLead field if non-nil, zero value otherwise.

### GetNTrialsCompanyLeadOk

`func (o *CompanyCosponsorsRow) GetNTrialsCompanyLeadOk() (*int32, bool)`

GetNTrialsCompanyLeadOk returns a tuple with the NTrialsCompanyLead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsCompanyLead

`func (o *CompanyCosponsorsRow) SetNTrialsCompanyLead(v int32)`

SetNTrialsCompanyLead sets NTrialsCompanyLead field to given value.

### HasNTrialsCompanyLead

`func (o *CompanyCosponsorsRow) HasNTrialsCompanyLead() bool`

HasNTrialsCompanyLead returns a boolean if a field has been set.

### SetNTrialsCompanyLeadNil

`func (o *CompanyCosponsorsRow) SetNTrialsCompanyLeadNil(b bool)`

 SetNTrialsCompanyLeadNil sets the value for NTrialsCompanyLead to be an explicit nil

### UnsetNTrialsCompanyLead
`func (o *CompanyCosponsorsRow) UnsetNTrialsCompanyLead()`

UnsetNTrialsCompanyLead ensures that no value is present for NTrialsCompanyLead, not even an explicit nil
### GetNTrialsCosponsorLead

`func (o *CompanyCosponsorsRow) GetNTrialsCosponsorLead() int32`

GetNTrialsCosponsorLead returns the NTrialsCosponsorLead field if non-nil, zero value otherwise.

### GetNTrialsCosponsorLeadOk

`func (o *CompanyCosponsorsRow) GetNTrialsCosponsorLeadOk() (*int32, bool)`

GetNTrialsCosponsorLeadOk returns a tuple with the NTrialsCosponsorLead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsCosponsorLead

`func (o *CompanyCosponsorsRow) SetNTrialsCosponsorLead(v int32)`

SetNTrialsCosponsorLead sets NTrialsCosponsorLead field to given value.

### HasNTrialsCosponsorLead

`func (o *CompanyCosponsorsRow) HasNTrialsCosponsorLead() bool`

HasNTrialsCosponsorLead returns a boolean if a field has been set.

### SetNTrialsCosponsorLeadNil

`func (o *CompanyCosponsorsRow) SetNTrialsCosponsorLeadNil(b bool)`

 SetNTrialsCosponsorLeadNil sets the value for NTrialsCosponsorLead to be an explicit nil

### UnsetNTrialsCosponsorLead
`func (o *CompanyCosponsorsRow) UnsetNTrialsCosponsorLead()`

UnsetNTrialsCosponsorLead ensures that no value is present for NTrialsCosponsorLead, not even an explicit nil
### GetFirstSharedStartYear

`func (o *CompanyCosponsorsRow) GetFirstSharedStartYear() int32`

GetFirstSharedStartYear returns the FirstSharedStartYear field if non-nil, zero value otherwise.

### GetFirstSharedStartYearOk

`func (o *CompanyCosponsorsRow) GetFirstSharedStartYearOk() (*int32, bool)`

GetFirstSharedStartYearOk returns a tuple with the FirstSharedStartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSharedStartYear

`func (o *CompanyCosponsorsRow) SetFirstSharedStartYear(v int32)`

SetFirstSharedStartYear sets FirstSharedStartYear field to given value.

### HasFirstSharedStartYear

`func (o *CompanyCosponsorsRow) HasFirstSharedStartYear() bool`

HasFirstSharedStartYear returns a boolean if a field has been set.

### SetFirstSharedStartYearNil

`func (o *CompanyCosponsorsRow) SetFirstSharedStartYearNil(b bool)`

 SetFirstSharedStartYearNil sets the value for FirstSharedStartYear to be an explicit nil

### UnsetFirstSharedStartYear
`func (o *CompanyCosponsorsRow) UnsetFirstSharedStartYear()`

UnsetFirstSharedStartYear ensures that no value is present for FirstSharedStartYear, not even an explicit nil
### GetLastSharedStartYear

`func (o *CompanyCosponsorsRow) GetLastSharedStartYear() int32`

GetLastSharedStartYear returns the LastSharedStartYear field if non-nil, zero value otherwise.

### GetLastSharedStartYearOk

`func (o *CompanyCosponsorsRow) GetLastSharedStartYearOk() (*int32, bool)`

GetLastSharedStartYearOk returns a tuple with the LastSharedStartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSharedStartYear

`func (o *CompanyCosponsorsRow) SetLastSharedStartYear(v int32)`

SetLastSharedStartYear sets LastSharedStartYear field to given value.

### HasLastSharedStartYear

`func (o *CompanyCosponsorsRow) HasLastSharedStartYear() bool`

HasLastSharedStartYear returns a boolean if a field has been set.

### SetLastSharedStartYearNil

`func (o *CompanyCosponsorsRow) SetLastSharedStartYearNil(b bool)`

 SetLastSharedStartYearNil sets the value for LastSharedStartYear to be an explicit nil

### UnsetLastSharedStartYear
`func (o *CompanyCosponsorsRow) UnsetLastSharedStartYear()`

UnsetLastSharedStartYear ensures that no value is present for LastSharedStartYear, not even an explicit nil
### GetNSharedConditions

`func (o *CompanyCosponsorsRow) GetNSharedConditions() int32`

GetNSharedConditions returns the NSharedConditions field if non-nil, zero value otherwise.

### GetNSharedConditionsOk

`func (o *CompanyCosponsorsRow) GetNSharedConditionsOk() (*int32, bool)`

GetNSharedConditionsOk returns a tuple with the NSharedConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSharedConditions

`func (o *CompanyCosponsorsRow) SetNSharedConditions(v int32)`

SetNSharedConditions sets NSharedConditions field to given value.

### HasNSharedConditions

`func (o *CompanyCosponsorsRow) HasNSharedConditions() bool`

HasNSharedConditions returns a boolean if a field has been set.

### SetNSharedConditionsNil

`func (o *CompanyCosponsorsRow) SetNSharedConditionsNil(b bool)`

 SetNSharedConditionsNil sets the value for NSharedConditions to be an explicit nil

### UnsetNSharedConditions
`func (o *CompanyCosponsorsRow) UnsetNSharedConditions()`

UnsetNSharedConditions ensures that no value is present for NSharedConditions, not even an explicit nil
### GetTopSharedConditions

`func (o *CompanyCosponsorsRow) GetTopSharedConditions() []string`

GetTopSharedConditions returns the TopSharedConditions field if non-nil, zero value otherwise.

### GetTopSharedConditionsOk

`func (o *CompanyCosponsorsRow) GetTopSharedConditionsOk() (*[]string, bool)`

GetTopSharedConditionsOk returns a tuple with the TopSharedConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSharedConditions

`func (o *CompanyCosponsorsRow) SetTopSharedConditions(v []string)`

SetTopSharedConditions sets TopSharedConditions field to given value.

### HasTopSharedConditions

`func (o *CompanyCosponsorsRow) HasTopSharedConditions() bool`

HasTopSharedConditions returns a boolean if a field has been set.

### SetTopSharedConditionsNil

`func (o *CompanyCosponsorsRow) SetTopSharedConditionsNil(b bool)`

 SetTopSharedConditionsNil sets the value for TopSharedConditions to be an explicit nil

### UnsetTopSharedConditions
`func (o *CompanyCosponsorsRow) UnsetTopSharedConditions()`

UnsetTopSharedConditions ensures that no value is present for TopSharedConditions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


