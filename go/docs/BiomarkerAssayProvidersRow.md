# BiomarkerAssayProvidersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteKey** | Pointer to **NullableString** | The route key of the gene the biomarker resolved to, the entry key for the biomarkers object type. Null when the extraction did not resolve a gene. | [optional] 
**BiomarkerKey** | Pointer to **NullableString** | The grouping key of the biomarker. The biomarker&#39;s key when its gene resolved, otherwise raw: followed by the lower-cased name as written, so two spellings of one unresolved biomarker are two keys. | [optional] 
**BiomarkerId** | Pointer to **NullableString** | The biomarker&#39;s own key; null for a pair whose gene did not resolve. | [optional] 
**BiomarkerDisplayName** | Pointer to **NullableString** | The biomarker&#39;s name, as resolved or else as extracted. | [optional] 
**GeneResolutionStatus** | Pointer to **NullableString** | resolved: the gene resolved. unresolved: a gene symbol was extracted and matched nothing. not_applicable: no gene symbol was extracted. | [optional] 
**ProviderKey** | Pointer to **NullableString** | The provider the rows are grouped on. org: followed by the organization key when the company resolves to an organization, otherwise name: followed by the normalised company name. | [optional] 
**ProviderRole** | Pointer to **NullableString** | What the papers name the provider as. assay_or_test_maker makes the assay or test; testing_service runs it as a service; both when the pair has papers of each kind. | [optional] 
**ProviderName** | Pointer to **NullableString** | The provider company name as it appears in the papers. | [optional] 
**ProviderOrgId** | Pointer to **NullableString** | The provider&#39;s legal entity key, where the marketplace match resolves one. | [optional] 
**ProviderOrgMatchTier** | Pointer to **NullableString** | How the provider was resolved to an organization. rx_provider_link: through its marketplace supplier. unresolved: no organization, the provider stays a name. The remaining values are reserved for resolution rungs that have not shipped. | [optional] 
**ProviderParentOrgId** | Pointer to **NullableString** | The ultimate parent organization of the provider&#39;s organization; null when the provider has no organization. | [optional] 
**ProviderCompanyId** | Pointer to **NullableString** | The provider&#39;s company key. | [optional] 
**ProviderCompanyName** | Pointer to **NullableString** | The provider&#39;s company name. | [optional] 
**ProviderGroupId** | Pointer to **NullableString** | The provider&#39;s current ultimate parent group key. | [optional] 
**ProviderGroupName** | Pointer to **NullableString** | The provider&#39;s group name. | [optional] 
**RxMatched** | Pointer to **NullableBool** | Whether the provider name matches a supplier in the Scientist.com marketplace. | [optional] 
**RxProviderId** | Pointer to **NullableString** | The matched marketplace supplier&#39;s id; join to the suppliers dataset. | [optional] 
**RxProviderName** | Pointer to **NullableString** | The matched marketplace supplier&#39;s name. | [optional] 
**RxProviderSlug** | Pointer to **NullableString** | The matched supplier&#39;s marketplace profile slug. | [optional] 
**HasClinicalValidation** | Pointer to **NullableBool** | Whether any paper for the pair mentions clinical validation of the assay. | [optional] 
**HasAnalyticalValidation** | Pointer to **NullableBool** | Whether any paper for the pair mentions analytical validation of the assay. | [optional] 
**NPapers** | Pointer to **NullableInt64** | Papers naming the biomarker and the provider, dated or not. | [optional] 
**NDatedPapers** | Pointer to **NullableInt64** | Papers with a recoverable publication date. | [optional] 
**NUndatedPapers** | Pointer to **NullableInt64** | Papers with no recoverable publication date. n_dated_papers plus n_undated_papers is n_papers. | [optional] 
**MostRecentPublishedAt** | Pointer to **NullableTime** | The latest publication date among the dated papers, in UTC. Null when no paper is dated, which is not a statement that none was published recently. | [optional] 
**MostRecentPmid** | Pointer to **NullableString** | The PubMed id of the most recent dated paper. | [optional] 
**MostRecentTitle** | Pointer to **NullableString** | The title of that paper. | [optional] 
**MeetsCuratedThreshold** | Pointer to **NullableBool** | True when the provider matches a marketplace supplier or at least two papers name the pair. | [optional] 
**AssayProviderClass** | Pointer to **NullableString** | The company-level class. diagnostic_provider makes or supplies the diagnostic assay; reagent_supplier is the vendor whose kit appears in a methods section; context_dependent, pharma_or_other and not_a_company are the other decisions; unclassified means no decision yet. | [optional] 
**AssayProviderClassSource** | Pointer to **NullableString** | Where the class decision came from; null when unclassified. | [optional] 
**AssayProviderClassDecidedAt** | Pointer to **NullableTime** | When the class was decided, in UTC; null when unclassified. | [optional] 

## Methods

### NewBiomarkerAssayProvidersRow

`func NewBiomarkerAssayProvidersRow() *BiomarkerAssayProvidersRow`

NewBiomarkerAssayProvidersRow instantiates a new BiomarkerAssayProvidersRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerAssayProvidersRowWithDefaults

`func NewBiomarkerAssayProvidersRowWithDefaults() *BiomarkerAssayProvidersRow`

NewBiomarkerAssayProvidersRowWithDefaults instantiates a new BiomarkerAssayProvidersRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteKey

`func (o *BiomarkerAssayProvidersRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerAssayProvidersRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerAssayProvidersRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerAssayProvidersRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerAssayProvidersRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerAssayProvidersRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetBiomarkerKey

`func (o *BiomarkerAssayProvidersRow) GetBiomarkerKey() string`

GetBiomarkerKey returns the BiomarkerKey field if non-nil, zero value otherwise.

### GetBiomarkerKeyOk

`func (o *BiomarkerAssayProvidersRow) GetBiomarkerKeyOk() (*string, bool)`

GetBiomarkerKeyOk returns a tuple with the BiomarkerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerKey

`func (o *BiomarkerAssayProvidersRow) SetBiomarkerKey(v string)`

SetBiomarkerKey sets BiomarkerKey field to given value.

### HasBiomarkerKey

`func (o *BiomarkerAssayProvidersRow) HasBiomarkerKey() bool`

HasBiomarkerKey returns a boolean if a field has been set.

### SetBiomarkerKeyNil

`func (o *BiomarkerAssayProvidersRow) SetBiomarkerKeyNil(b bool)`

 SetBiomarkerKeyNil sets the value for BiomarkerKey to be an explicit nil

### UnsetBiomarkerKey
`func (o *BiomarkerAssayProvidersRow) UnsetBiomarkerKey()`

UnsetBiomarkerKey ensures that no value is present for BiomarkerKey, not even an explicit nil
### GetBiomarkerId

`func (o *BiomarkerAssayProvidersRow) GetBiomarkerId() string`

GetBiomarkerId returns the BiomarkerId field if non-nil, zero value otherwise.

### GetBiomarkerIdOk

`func (o *BiomarkerAssayProvidersRow) GetBiomarkerIdOk() (*string, bool)`

GetBiomarkerIdOk returns a tuple with the BiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerId

`func (o *BiomarkerAssayProvidersRow) SetBiomarkerId(v string)`

SetBiomarkerId sets BiomarkerId field to given value.

### HasBiomarkerId

`func (o *BiomarkerAssayProvidersRow) HasBiomarkerId() bool`

HasBiomarkerId returns a boolean if a field has been set.

### SetBiomarkerIdNil

`func (o *BiomarkerAssayProvidersRow) SetBiomarkerIdNil(b bool)`

 SetBiomarkerIdNil sets the value for BiomarkerId to be an explicit nil

### UnsetBiomarkerId
`func (o *BiomarkerAssayProvidersRow) UnsetBiomarkerId()`

UnsetBiomarkerId ensures that no value is present for BiomarkerId, not even an explicit nil
### GetBiomarkerDisplayName

`func (o *BiomarkerAssayProvidersRow) GetBiomarkerDisplayName() string`

GetBiomarkerDisplayName returns the BiomarkerDisplayName field if non-nil, zero value otherwise.

### GetBiomarkerDisplayNameOk

`func (o *BiomarkerAssayProvidersRow) GetBiomarkerDisplayNameOk() (*string, bool)`

GetBiomarkerDisplayNameOk returns a tuple with the BiomarkerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerDisplayName

`func (o *BiomarkerAssayProvidersRow) SetBiomarkerDisplayName(v string)`

SetBiomarkerDisplayName sets BiomarkerDisplayName field to given value.

### HasBiomarkerDisplayName

`func (o *BiomarkerAssayProvidersRow) HasBiomarkerDisplayName() bool`

HasBiomarkerDisplayName returns a boolean if a field has been set.

### SetBiomarkerDisplayNameNil

`func (o *BiomarkerAssayProvidersRow) SetBiomarkerDisplayNameNil(b bool)`

 SetBiomarkerDisplayNameNil sets the value for BiomarkerDisplayName to be an explicit nil

### UnsetBiomarkerDisplayName
`func (o *BiomarkerAssayProvidersRow) UnsetBiomarkerDisplayName()`

UnsetBiomarkerDisplayName ensures that no value is present for BiomarkerDisplayName, not even an explicit nil
### GetGeneResolutionStatus

`func (o *BiomarkerAssayProvidersRow) GetGeneResolutionStatus() string`

GetGeneResolutionStatus returns the GeneResolutionStatus field if non-nil, zero value otherwise.

### GetGeneResolutionStatusOk

`func (o *BiomarkerAssayProvidersRow) GetGeneResolutionStatusOk() (*string, bool)`

GetGeneResolutionStatusOk returns a tuple with the GeneResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneResolutionStatus

`func (o *BiomarkerAssayProvidersRow) SetGeneResolutionStatus(v string)`

SetGeneResolutionStatus sets GeneResolutionStatus field to given value.

### HasGeneResolutionStatus

`func (o *BiomarkerAssayProvidersRow) HasGeneResolutionStatus() bool`

HasGeneResolutionStatus returns a boolean if a field has been set.

### SetGeneResolutionStatusNil

`func (o *BiomarkerAssayProvidersRow) SetGeneResolutionStatusNil(b bool)`

 SetGeneResolutionStatusNil sets the value for GeneResolutionStatus to be an explicit nil

### UnsetGeneResolutionStatus
`func (o *BiomarkerAssayProvidersRow) UnsetGeneResolutionStatus()`

UnsetGeneResolutionStatus ensures that no value is present for GeneResolutionStatus, not even an explicit nil
### GetProviderKey

`func (o *BiomarkerAssayProvidersRow) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *BiomarkerAssayProvidersRow) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *BiomarkerAssayProvidersRow) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *BiomarkerAssayProvidersRow) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### SetProviderKeyNil

`func (o *BiomarkerAssayProvidersRow) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *BiomarkerAssayProvidersRow) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
### GetProviderRole

`func (o *BiomarkerAssayProvidersRow) GetProviderRole() string`

GetProviderRole returns the ProviderRole field if non-nil, zero value otherwise.

### GetProviderRoleOk

`func (o *BiomarkerAssayProvidersRow) GetProviderRoleOk() (*string, bool)`

GetProviderRoleOk returns a tuple with the ProviderRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderRole

`func (o *BiomarkerAssayProvidersRow) SetProviderRole(v string)`

SetProviderRole sets ProviderRole field to given value.

### HasProviderRole

`func (o *BiomarkerAssayProvidersRow) HasProviderRole() bool`

HasProviderRole returns a boolean if a field has been set.

### SetProviderRoleNil

`func (o *BiomarkerAssayProvidersRow) SetProviderRoleNil(b bool)`

 SetProviderRoleNil sets the value for ProviderRole to be an explicit nil

### UnsetProviderRole
`func (o *BiomarkerAssayProvidersRow) UnsetProviderRole()`

UnsetProviderRole ensures that no value is present for ProviderRole, not even an explicit nil
### GetProviderName

`func (o *BiomarkerAssayProvidersRow) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *BiomarkerAssayProvidersRow) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *BiomarkerAssayProvidersRow) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *BiomarkerAssayProvidersRow) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *BiomarkerAssayProvidersRow) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *BiomarkerAssayProvidersRow) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetProviderOrgId

`func (o *BiomarkerAssayProvidersRow) GetProviderOrgId() string`

GetProviderOrgId returns the ProviderOrgId field if non-nil, zero value otherwise.

### GetProviderOrgIdOk

`func (o *BiomarkerAssayProvidersRow) GetProviderOrgIdOk() (*string, bool)`

GetProviderOrgIdOk returns a tuple with the ProviderOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderOrgId

`func (o *BiomarkerAssayProvidersRow) SetProviderOrgId(v string)`

SetProviderOrgId sets ProviderOrgId field to given value.

### HasProviderOrgId

`func (o *BiomarkerAssayProvidersRow) HasProviderOrgId() bool`

HasProviderOrgId returns a boolean if a field has been set.

### SetProviderOrgIdNil

`func (o *BiomarkerAssayProvidersRow) SetProviderOrgIdNil(b bool)`

 SetProviderOrgIdNil sets the value for ProviderOrgId to be an explicit nil

### UnsetProviderOrgId
`func (o *BiomarkerAssayProvidersRow) UnsetProviderOrgId()`

UnsetProviderOrgId ensures that no value is present for ProviderOrgId, not even an explicit nil
### GetProviderOrgMatchTier

`func (o *BiomarkerAssayProvidersRow) GetProviderOrgMatchTier() string`

GetProviderOrgMatchTier returns the ProviderOrgMatchTier field if non-nil, zero value otherwise.

### GetProviderOrgMatchTierOk

`func (o *BiomarkerAssayProvidersRow) GetProviderOrgMatchTierOk() (*string, bool)`

GetProviderOrgMatchTierOk returns a tuple with the ProviderOrgMatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderOrgMatchTier

`func (o *BiomarkerAssayProvidersRow) SetProviderOrgMatchTier(v string)`

SetProviderOrgMatchTier sets ProviderOrgMatchTier field to given value.

### HasProviderOrgMatchTier

`func (o *BiomarkerAssayProvidersRow) HasProviderOrgMatchTier() bool`

HasProviderOrgMatchTier returns a boolean if a field has been set.

### SetProviderOrgMatchTierNil

`func (o *BiomarkerAssayProvidersRow) SetProviderOrgMatchTierNil(b bool)`

 SetProviderOrgMatchTierNil sets the value for ProviderOrgMatchTier to be an explicit nil

### UnsetProviderOrgMatchTier
`func (o *BiomarkerAssayProvidersRow) UnsetProviderOrgMatchTier()`

UnsetProviderOrgMatchTier ensures that no value is present for ProviderOrgMatchTier, not even an explicit nil
### GetProviderParentOrgId

`func (o *BiomarkerAssayProvidersRow) GetProviderParentOrgId() string`

GetProviderParentOrgId returns the ProviderParentOrgId field if non-nil, zero value otherwise.

### GetProviderParentOrgIdOk

`func (o *BiomarkerAssayProvidersRow) GetProviderParentOrgIdOk() (*string, bool)`

GetProviderParentOrgIdOk returns a tuple with the ProviderParentOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderParentOrgId

`func (o *BiomarkerAssayProvidersRow) SetProviderParentOrgId(v string)`

SetProviderParentOrgId sets ProviderParentOrgId field to given value.

### HasProviderParentOrgId

`func (o *BiomarkerAssayProvidersRow) HasProviderParentOrgId() bool`

HasProviderParentOrgId returns a boolean if a field has been set.

### SetProviderParentOrgIdNil

`func (o *BiomarkerAssayProvidersRow) SetProviderParentOrgIdNil(b bool)`

 SetProviderParentOrgIdNil sets the value for ProviderParentOrgId to be an explicit nil

### UnsetProviderParentOrgId
`func (o *BiomarkerAssayProvidersRow) UnsetProviderParentOrgId()`

UnsetProviderParentOrgId ensures that no value is present for ProviderParentOrgId, not even an explicit nil
### GetProviderCompanyId

`func (o *BiomarkerAssayProvidersRow) GetProviderCompanyId() string`

GetProviderCompanyId returns the ProviderCompanyId field if non-nil, zero value otherwise.

### GetProviderCompanyIdOk

`func (o *BiomarkerAssayProvidersRow) GetProviderCompanyIdOk() (*string, bool)`

GetProviderCompanyIdOk returns a tuple with the ProviderCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCompanyId

`func (o *BiomarkerAssayProvidersRow) SetProviderCompanyId(v string)`

SetProviderCompanyId sets ProviderCompanyId field to given value.

### HasProviderCompanyId

`func (o *BiomarkerAssayProvidersRow) HasProviderCompanyId() bool`

HasProviderCompanyId returns a boolean if a field has been set.

### SetProviderCompanyIdNil

`func (o *BiomarkerAssayProvidersRow) SetProviderCompanyIdNil(b bool)`

 SetProviderCompanyIdNil sets the value for ProviderCompanyId to be an explicit nil

### UnsetProviderCompanyId
`func (o *BiomarkerAssayProvidersRow) UnsetProviderCompanyId()`

UnsetProviderCompanyId ensures that no value is present for ProviderCompanyId, not even an explicit nil
### GetProviderCompanyName

`func (o *BiomarkerAssayProvidersRow) GetProviderCompanyName() string`

GetProviderCompanyName returns the ProviderCompanyName field if non-nil, zero value otherwise.

### GetProviderCompanyNameOk

`func (o *BiomarkerAssayProvidersRow) GetProviderCompanyNameOk() (*string, bool)`

GetProviderCompanyNameOk returns a tuple with the ProviderCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCompanyName

`func (o *BiomarkerAssayProvidersRow) SetProviderCompanyName(v string)`

SetProviderCompanyName sets ProviderCompanyName field to given value.

### HasProviderCompanyName

`func (o *BiomarkerAssayProvidersRow) HasProviderCompanyName() bool`

HasProviderCompanyName returns a boolean if a field has been set.

### SetProviderCompanyNameNil

`func (o *BiomarkerAssayProvidersRow) SetProviderCompanyNameNil(b bool)`

 SetProviderCompanyNameNil sets the value for ProviderCompanyName to be an explicit nil

### UnsetProviderCompanyName
`func (o *BiomarkerAssayProvidersRow) UnsetProviderCompanyName()`

UnsetProviderCompanyName ensures that no value is present for ProviderCompanyName, not even an explicit nil
### GetProviderGroupId

`func (o *BiomarkerAssayProvidersRow) GetProviderGroupId() string`

GetProviderGroupId returns the ProviderGroupId field if non-nil, zero value otherwise.

### GetProviderGroupIdOk

`func (o *BiomarkerAssayProvidersRow) GetProviderGroupIdOk() (*string, bool)`

GetProviderGroupIdOk returns a tuple with the ProviderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderGroupId

`func (o *BiomarkerAssayProvidersRow) SetProviderGroupId(v string)`

SetProviderGroupId sets ProviderGroupId field to given value.

### HasProviderGroupId

`func (o *BiomarkerAssayProvidersRow) HasProviderGroupId() bool`

HasProviderGroupId returns a boolean if a field has been set.

### SetProviderGroupIdNil

`func (o *BiomarkerAssayProvidersRow) SetProviderGroupIdNil(b bool)`

 SetProviderGroupIdNil sets the value for ProviderGroupId to be an explicit nil

### UnsetProviderGroupId
`func (o *BiomarkerAssayProvidersRow) UnsetProviderGroupId()`

UnsetProviderGroupId ensures that no value is present for ProviderGroupId, not even an explicit nil
### GetProviderGroupName

`func (o *BiomarkerAssayProvidersRow) GetProviderGroupName() string`

GetProviderGroupName returns the ProviderGroupName field if non-nil, zero value otherwise.

### GetProviderGroupNameOk

`func (o *BiomarkerAssayProvidersRow) GetProviderGroupNameOk() (*string, bool)`

GetProviderGroupNameOk returns a tuple with the ProviderGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderGroupName

`func (o *BiomarkerAssayProvidersRow) SetProviderGroupName(v string)`

SetProviderGroupName sets ProviderGroupName field to given value.

### HasProviderGroupName

`func (o *BiomarkerAssayProvidersRow) HasProviderGroupName() bool`

HasProviderGroupName returns a boolean if a field has been set.

### SetProviderGroupNameNil

`func (o *BiomarkerAssayProvidersRow) SetProviderGroupNameNil(b bool)`

 SetProviderGroupNameNil sets the value for ProviderGroupName to be an explicit nil

### UnsetProviderGroupName
`func (o *BiomarkerAssayProvidersRow) UnsetProviderGroupName()`

UnsetProviderGroupName ensures that no value is present for ProviderGroupName, not even an explicit nil
### GetRxMatched

`func (o *BiomarkerAssayProvidersRow) GetRxMatched() bool`

GetRxMatched returns the RxMatched field if non-nil, zero value otherwise.

### GetRxMatchedOk

`func (o *BiomarkerAssayProvidersRow) GetRxMatchedOk() (*bool, bool)`

GetRxMatchedOk returns a tuple with the RxMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMatched

`func (o *BiomarkerAssayProvidersRow) SetRxMatched(v bool)`

SetRxMatched sets RxMatched field to given value.

### HasRxMatched

`func (o *BiomarkerAssayProvidersRow) HasRxMatched() bool`

HasRxMatched returns a boolean if a field has been set.

### SetRxMatchedNil

`func (o *BiomarkerAssayProvidersRow) SetRxMatchedNil(b bool)`

 SetRxMatchedNil sets the value for RxMatched to be an explicit nil

### UnsetRxMatched
`func (o *BiomarkerAssayProvidersRow) UnsetRxMatched()`

UnsetRxMatched ensures that no value is present for RxMatched, not even an explicit nil
### GetRxProviderId

`func (o *BiomarkerAssayProvidersRow) GetRxProviderId() string`

GetRxProviderId returns the RxProviderId field if non-nil, zero value otherwise.

### GetRxProviderIdOk

`func (o *BiomarkerAssayProvidersRow) GetRxProviderIdOk() (*string, bool)`

GetRxProviderIdOk returns a tuple with the RxProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxProviderId

`func (o *BiomarkerAssayProvidersRow) SetRxProviderId(v string)`

SetRxProviderId sets RxProviderId field to given value.

### HasRxProviderId

`func (o *BiomarkerAssayProvidersRow) HasRxProviderId() bool`

HasRxProviderId returns a boolean if a field has been set.

### SetRxProviderIdNil

`func (o *BiomarkerAssayProvidersRow) SetRxProviderIdNil(b bool)`

 SetRxProviderIdNil sets the value for RxProviderId to be an explicit nil

### UnsetRxProviderId
`func (o *BiomarkerAssayProvidersRow) UnsetRxProviderId()`

UnsetRxProviderId ensures that no value is present for RxProviderId, not even an explicit nil
### GetRxProviderName

`func (o *BiomarkerAssayProvidersRow) GetRxProviderName() string`

GetRxProviderName returns the RxProviderName field if non-nil, zero value otherwise.

### GetRxProviderNameOk

`func (o *BiomarkerAssayProvidersRow) GetRxProviderNameOk() (*string, bool)`

GetRxProviderNameOk returns a tuple with the RxProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxProviderName

`func (o *BiomarkerAssayProvidersRow) SetRxProviderName(v string)`

SetRxProviderName sets RxProviderName field to given value.

### HasRxProviderName

`func (o *BiomarkerAssayProvidersRow) HasRxProviderName() bool`

HasRxProviderName returns a boolean if a field has been set.

### SetRxProviderNameNil

`func (o *BiomarkerAssayProvidersRow) SetRxProviderNameNil(b bool)`

 SetRxProviderNameNil sets the value for RxProviderName to be an explicit nil

### UnsetRxProviderName
`func (o *BiomarkerAssayProvidersRow) UnsetRxProviderName()`

UnsetRxProviderName ensures that no value is present for RxProviderName, not even an explicit nil
### GetRxProviderSlug

`func (o *BiomarkerAssayProvidersRow) GetRxProviderSlug() string`

GetRxProviderSlug returns the RxProviderSlug field if non-nil, zero value otherwise.

### GetRxProviderSlugOk

`func (o *BiomarkerAssayProvidersRow) GetRxProviderSlugOk() (*string, bool)`

GetRxProviderSlugOk returns a tuple with the RxProviderSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxProviderSlug

`func (o *BiomarkerAssayProvidersRow) SetRxProviderSlug(v string)`

SetRxProviderSlug sets RxProviderSlug field to given value.

### HasRxProviderSlug

`func (o *BiomarkerAssayProvidersRow) HasRxProviderSlug() bool`

HasRxProviderSlug returns a boolean if a field has been set.

### SetRxProviderSlugNil

`func (o *BiomarkerAssayProvidersRow) SetRxProviderSlugNil(b bool)`

 SetRxProviderSlugNil sets the value for RxProviderSlug to be an explicit nil

### UnsetRxProviderSlug
`func (o *BiomarkerAssayProvidersRow) UnsetRxProviderSlug()`

UnsetRxProviderSlug ensures that no value is present for RxProviderSlug, not even an explicit nil
### GetHasClinicalValidation

`func (o *BiomarkerAssayProvidersRow) GetHasClinicalValidation() bool`

GetHasClinicalValidation returns the HasClinicalValidation field if non-nil, zero value otherwise.

### GetHasClinicalValidationOk

`func (o *BiomarkerAssayProvidersRow) GetHasClinicalValidationOk() (*bool, bool)`

GetHasClinicalValidationOk returns a tuple with the HasClinicalValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasClinicalValidation

`func (o *BiomarkerAssayProvidersRow) SetHasClinicalValidation(v bool)`

SetHasClinicalValidation sets HasClinicalValidation field to given value.

### HasHasClinicalValidation

`func (o *BiomarkerAssayProvidersRow) HasHasClinicalValidation() bool`

HasHasClinicalValidation returns a boolean if a field has been set.

### SetHasClinicalValidationNil

`func (o *BiomarkerAssayProvidersRow) SetHasClinicalValidationNil(b bool)`

 SetHasClinicalValidationNil sets the value for HasClinicalValidation to be an explicit nil

### UnsetHasClinicalValidation
`func (o *BiomarkerAssayProvidersRow) UnsetHasClinicalValidation()`

UnsetHasClinicalValidation ensures that no value is present for HasClinicalValidation, not even an explicit nil
### GetHasAnalyticalValidation

`func (o *BiomarkerAssayProvidersRow) GetHasAnalyticalValidation() bool`

GetHasAnalyticalValidation returns the HasAnalyticalValidation field if non-nil, zero value otherwise.

### GetHasAnalyticalValidationOk

`func (o *BiomarkerAssayProvidersRow) GetHasAnalyticalValidationOk() (*bool, bool)`

GetHasAnalyticalValidationOk returns a tuple with the HasAnalyticalValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAnalyticalValidation

`func (o *BiomarkerAssayProvidersRow) SetHasAnalyticalValidation(v bool)`

SetHasAnalyticalValidation sets HasAnalyticalValidation field to given value.

### HasHasAnalyticalValidation

`func (o *BiomarkerAssayProvidersRow) HasHasAnalyticalValidation() bool`

HasHasAnalyticalValidation returns a boolean if a field has been set.

### SetHasAnalyticalValidationNil

`func (o *BiomarkerAssayProvidersRow) SetHasAnalyticalValidationNil(b bool)`

 SetHasAnalyticalValidationNil sets the value for HasAnalyticalValidation to be an explicit nil

### UnsetHasAnalyticalValidation
`func (o *BiomarkerAssayProvidersRow) UnsetHasAnalyticalValidation()`

UnsetHasAnalyticalValidation ensures that no value is present for HasAnalyticalValidation, not even an explicit nil
### GetNPapers

`func (o *BiomarkerAssayProvidersRow) GetNPapers() int64`

GetNPapers returns the NPapers field if non-nil, zero value otherwise.

### GetNPapersOk

`func (o *BiomarkerAssayProvidersRow) GetNPapersOk() (*int64, bool)`

GetNPapersOk returns a tuple with the NPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPapers

`func (o *BiomarkerAssayProvidersRow) SetNPapers(v int64)`

SetNPapers sets NPapers field to given value.

### HasNPapers

`func (o *BiomarkerAssayProvidersRow) HasNPapers() bool`

HasNPapers returns a boolean if a field has been set.

### SetNPapersNil

`func (o *BiomarkerAssayProvidersRow) SetNPapersNil(b bool)`

 SetNPapersNil sets the value for NPapers to be an explicit nil

### UnsetNPapers
`func (o *BiomarkerAssayProvidersRow) UnsetNPapers()`

UnsetNPapers ensures that no value is present for NPapers, not even an explicit nil
### GetNDatedPapers

`func (o *BiomarkerAssayProvidersRow) GetNDatedPapers() int64`

GetNDatedPapers returns the NDatedPapers field if non-nil, zero value otherwise.

### GetNDatedPapersOk

`func (o *BiomarkerAssayProvidersRow) GetNDatedPapersOk() (*int64, bool)`

GetNDatedPapersOk returns a tuple with the NDatedPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDatedPapers

`func (o *BiomarkerAssayProvidersRow) SetNDatedPapers(v int64)`

SetNDatedPapers sets NDatedPapers field to given value.

### HasNDatedPapers

`func (o *BiomarkerAssayProvidersRow) HasNDatedPapers() bool`

HasNDatedPapers returns a boolean if a field has been set.

### SetNDatedPapersNil

`func (o *BiomarkerAssayProvidersRow) SetNDatedPapersNil(b bool)`

 SetNDatedPapersNil sets the value for NDatedPapers to be an explicit nil

### UnsetNDatedPapers
`func (o *BiomarkerAssayProvidersRow) UnsetNDatedPapers()`

UnsetNDatedPapers ensures that no value is present for NDatedPapers, not even an explicit nil
### GetNUndatedPapers

`func (o *BiomarkerAssayProvidersRow) GetNUndatedPapers() int64`

GetNUndatedPapers returns the NUndatedPapers field if non-nil, zero value otherwise.

### GetNUndatedPapersOk

`func (o *BiomarkerAssayProvidersRow) GetNUndatedPapersOk() (*int64, bool)`

GetNUndatedPapersOk returns a tuple with the NUndatedPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNUndatedPapers

`func (o *BiomarkerAssayProvidersRow) SetNUndatedPapers(v int64)`

SetNUndatedPapers sets NUndatedPapers field to given value.

### HasNUndatedPapers

`func (o *BiomarkerAssayProvidersRow) HasNUndatedPapers() bool`

HasNUndatedPapers returns a boolean if a field has been set.

### SetNUndatedPapersNil

`func (o *BiomarkerAssayProvidersRow) SetNUndatedPapersNil(b bool)`

 SetNUndatedPapersNil sets the value for NUndatedPapers to be an explicit nil

### UnsetNUndatedPapers
`func (o *BiomarkerAssayProvidersRow) UnsetNUndatedPapers()`

UnsetNUndatedPapers ensures that no value is present for NUndatedPapers, not even an explicit nil
### GetMostRecentPublishedAt

`func (o *BiomarkerAssayProvidersRow) GetMostRecentPublishedAt() time.Time`

GetMostRecentPublishedAt returns the MostRecentPublishedAt field if non-nil, zero value otherwise.

### GetMostRecentPublishedAtOk

`func (o *BiomarkerAssayProvidersRow) GetMostRecentPublishedAtOk() (*time.Time, bool)`

GetMostRecentPublishedAtOk returns a tuple with the MostRecentPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentPublishedAt

`func (o *BiomarkerAssayProvidersRow) SetMostRecentPublishedAt(v time.Time)`

SetMostRecentPublishedAt sets MostRecentPublishedAt field to given value.

### HasMostRecentPublishedAt

`func (o *BiomarkerAssayProvidersRow) HasMostRecentPublishedAt() bool`

HasMostRecentPublishedAt returns a boolean if a field has been set.

### SetMostRecentPublishedAtNil

`func (o *BiomarkerAssayProvidersRow) SetMostRecentPublishedAtNil(b bool)`

 SetMostRecentPublishedAtNil sets the value for MostRecentPublishedAt to be an explicit nil

### UnsetMostRecentPublishedAt
`func (o *BiomarkerAssayProvidersRow) UnsetMostRecentPublishedAt()`

UnsetMostRecentPublishedAt ensures that no value is present for MostRecentPublishedAt, not even an explicit nil
### GetMostRecentPmid

`func (o *BiomarkerAssayProvidersRow) GetMostRecentPmid() string`

GetMostRecentPmid returns the MostRecentPmid field if non-nil, zero value otherwise.

### GetMostRecentPmidOk

`func (o *BiomarkerAssayProvidersRow) GetMostRecentPmidOk() (*string, bool)`

GetMostRecentPmidOk returns a tuple with the MostRecentPmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentPmid

`func (o *BiomarkerAssayProvidersRow) SetMostRecentPmid(v string)`

SetMostRecentPmid sets MostRecentPmid field to given value.

### HasMostRecentPmid

`func (o *BiomarkerAssayProvidersRow) HasMostRecentPmid() bool`

HasMostRecentPmid returns a boolean if a field has been set.

### SetMostRecentPmidNil

`func (o *BiomarkerAssayProvidersRow) SetMostRecentPmidNil(b bool)`

 SetMostRecentPmidNil sets the value for MostRecentPmid to be an explicit nil

### UnsetMostRecentPmid
`func (o *BiomarkerAssayProvidersRow) UnsetMostRecentPmid()`

UnsetMostRecentPmid ensures that no value is present for MostRecentPmid, not even an explicit nil
### GetMostRecentTitle

`func (o *BiomarkerAssayProvidersRow) GetMostRecentTitle() string`

GetMostRecentTitle returns the MostRecentTitle field if non-nil, zero value otherwise.

### GetMostRecentTitleOk

`func (o *BiomarkerAssayProvidersRow) GetMostRecentTitleOk() (*string, bool)`

GetMostRecentTitleOk returns a tuple with the MostRecentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentTitle

`func (o *BiomarkerAssayProvidersRow) SetMostRecentTitle(v string)`

SetMostRecentTitle sets MostRecentTitle field to given value.

### HasMostRecentTitle

`func (o *BiomarkerAssayProvidersRow) HasMostRecentTitle() bool`

HasMostRecentTitle returns a boolean if a field has been set.

### SetMostRecentTitleNil

`func (o *BiomarkerAssayProvidersRow) SetMostRecentTitleNil(b bool)`

 SetMostRecentTitleNil sets the value for MostRecentTitle to be an explicit nil

### UnsetMostRecentTitle
`func (o *BiomarkerAssayProvidersRow) UnsetMostRecentTitle()`

UnsetMostRecentTitle ensures that no value is present for MostRecentTitle, not even an explicit nil
### GetMeetsCuratedThreshold

`func (o *BiomarkerAssayProvidersRow) GetMeetsCuratedThreshold() bool`

GetMeetsCuratedThreshold returns the MeetsCuratedThreshold field if non-nil, zero value otherwise.

### GetMeetsCuratedThresholdOk

`func (o *BiomarkerAssayProvidersRow) GetMeetsCuratedThresholdOk() (*bool, bool)`

GetMeetsCuratedThresholdOk returns a tuple with the MeetsCuratedThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetsCuratedThreshold

`func (o *BiomarkerAssayProvidersRow) SetMeetsCuratedThreshold(v bool)`

SetMeetsCuratedThreshold sets MeetsCuratedThreshold field to given value.

### HasMeetsCuratedThreshold

`func (o *BiomarkerAssayProvidersRow) HasMeetsCuratedThreshold() bool`

HasMeetsCuratedThreshold returns a boolean if a field has been set.

### SetMeetsCuratedThresholdNil

`func (o *BiomarkerAssayProvidersRow) SetMeetsCuratedThresholdNil(b bool)`

 SetMeetsCuratedThresholdNil sets the value for MeetsCuratedThreshold to be an explicit nil

### UnsetMeetsCuratedThreshold
`func (o *BiomarkerAssayProvidersRow) UnsetMeetsCuratedThreshold()`

UnsetMeetsCuratedThreshold ensures that no value is present for MeetsCuratedThreshold, not even an explicit nil
### GetAssayProviderClass

`func (o *BiomarkerAssayProvidersRow) GetAssayProviderClass() string`

GetAssayProviderClass returns the AssayProviderClass field if non-nil, zero value otherwise.

### GetAssayProviderClassOk

`func (o *BiomarkerAssayProvidersRow) GetAssayProviderClassOk() (*string, bool)`

GetAssayProviderClassOk returns a tuple with the AssayProviderClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayProviderClass

`func (o *BiomarkerAssayProvidersRow) SetAssayProviderClass(v string)`

SetAssayProviderClass sets AssayProviderClass field to given value.

### HasAssayProviderClass

`func (o *BiomarkerAssayProvidersRow) HasAssayProviderClass() bool`

HasAssayProviderClass returns a boolean if a field has been set.

### SetAssayProviderClassNil

`func (o *BiomarkerAssayProvidersRow) SetAssayProviderClassNil(b bool)`

 SetAssayProviderClassNil sets the value for AssayProviderClass to be an explicit nil

### UnsetAssayProviderClass
`func (o *BiomarkerAssayProvidersRow) UnsetAssayProviderClass()`

UnsetAssayProviderClass ensures that no value is present for AssayProviderClass, not even an explicit nil
### GetAssayProviderClassSource

`func (o *BiomarkerAssayProvidersRow) GetAssayProviderClassSource() string`

GetAssayProviderClassSource returns the AssayProviderClassSource field if non-nil, zero value otherwise.

### GetAssayProviderClassSourceOk

`func (o *BiomarkerAssayProvidersRow) GetAssayProviderClassSourceOk() (*string, bool)`

GetAssayProviderClassSourceOk returns a tuple with the AssayProviderClassSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayProviderClassSource

`func (o *BiomarkerAssayProvidersRow) SetAssayProviderClassSource(v string)`

SetAssayProviderClassSource sets AssayProviderClassSource field to given value.

### HasAssayProviderClassSource

`func (o *BiomarkerAssayProvidersRow) HasAssayProviderClassSource() bool`

HasAssayProviderClassSource returns a boolean if a field has been set.

### SetAssayProviderClassSourceNil

`func (o *BiomarkerAssayProvidersRow) SetAssayProviderClassSourceNil(b bool)`

 SetAssayProviderClassSourceNil sets the value for AssayProviderClassSource to be an explicit nil

### UnsetAssayProviderClassSource
`func (o *BiomarkerAssayProvidersRow) UnsetAssayProviderClassSource()`

UnsetAssayProviderClassSource ensures that no value is present for AssayProviderClassSource, not even an explicit nil
### GetAssayProviderClassDecidedAt

`func (o *BiomarkerAssayProvidersRow) GetAssayProviderClassDecidedAt() time.Time`

GetAssayProviderClassDecidedAt returns the AssayProviderClassDecidedAt field if non-nil, zero value otherwise.

### GetAssayProviderClassDecidedAtOk

`func (o *BiomarkerAssayProvidersRow) GetAssayProviderClassDecidedAtOk() (*time.Time, bool)`

GetAssayProviderClassDecidedAtOk returns a tuple with the AssayProviderClassDecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayProviderClassDecidedAt

`func (o *BiomarkerAssayProvidersRow) SetAssayProviderClassDecidedAt(v time.Time)`

SetAssayProviderClassDecidedAt sets AssayProviderClassDecidedAt field to given value.

### HasAssayProviderClassDecidedAt

`func (o *BiomarkerAssayProvidersRow) HasAssayProviderClassDecidedAt() bool`

HasAssayProviderClassDecidedAt returns a boolean if a field has been set.

### SetAssayProviderClassDecidedAtNil

`func (o *BiomarkerAssayProvidersRow) SetAssayProviderClassDecidedAtNil(b bool)`

 SetAssayProviderClassDecidedAtNil sets the value for AssayProviderClassDecidedAt to be an explicit nil

### UnsetAssayProviderClassDecidedAt
`func (o *BiomarkerAssayProvidersRow) UnsetAssayProviderClassDecidedAt()`

UnsetAssayProviderClassDecidedAt ensures that no value is present for AssayProviderClassDecidedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


