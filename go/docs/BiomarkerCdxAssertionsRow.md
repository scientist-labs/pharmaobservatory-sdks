# BiomarkerCdxAssertionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAssertionId** | Pointer to **NullableString** | The assertion id. | [optional] 
**RouteKey** | Pointer to **NullableString** | The biomarker&#39;s route key, the entry key for the biomarkers object type. | [optional] 
**BiomarkerId** | Pointer to **NullableString** | The biomarker key. | [optional] 
**BiomarkerToken** | Pointer to **NullableString** | The biomarker as the FDA list names it. | [optional] 
**DiagnosticName** | Pointer to **NullableString** | The companion diagnostic&#39;s name. | [optional] 
**Indication** | Pointer to **NullableString** | The indication. | [optional] 
**SampleType** | Pointer to **NullableString** | The sample type. | [optional] 
**SubmissionCell** | Pointer to **NullableString** | The submission number cell as listed. | [optional] 
**LatestStatedDecisionDate** | Pointer to **NullableString** | Latest decision date the list states. | [optional] 
**SourceUrl** | Pointer to **NullableString** | Link to the FDA list. | [optional] 
**ManufacturerName** | Pointer to **NullableString** | The manufacturer as filed. | [optional] 
**ManufacturerOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the manufacturer resolved to (null when unresolved). | [optional] 
**ManufacturerCompanyId** | Pointer to **NullableString** | Company level: the company the manufacturer belongs to (null when unresolved). | [optional] 
**ManufacturerCompanyName** | Pointer to **NullableString** | Display name of the manufacturer&#39;s company. | [optional] 
**ManufacturerGroupId** | Pointer to **NullableString** | Group level: the manufacturer&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**ManufacturerGroupName** | Pointer to **NullableString** | Display name of the manufacturer&#39;s group. | [optional] 
**ManufacturerMatchTier** | Pointer to **NullableString** | How firmly the manufacturer name matched an organization. | [optional] 

## Methods

### NewBiomarkerCdxAssertionsRow

`func NewBiomarkerCdxAssertionsRow() *BiomarkerCdxAssertionsRow`

NewBiomarkerCdxAssertionsRow instantiates a new BiomarkerCdxAssertionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerCdxAssertionsRowWithDefaults

`func NewBiomarkerCdxAssertionsRowWithDefaults() *BiomarkerCdxAssertionsRow`

NewBiomarkerCdxAssertionsRowWithDefaults instantiates a new BiomarkerCdxAssertionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAssertionId

`func (o *BiomarkerCdxAssertionsRow) GetSourceAssertionId() string`

GetSourceAssertionId returns the SourceAssertionId field if non-nil, zero value otherwise.

### GetSourceAssertionIdOk

`func (o *BiomarkerCdxAssertionsRow) GetSourceAssertionIdOk() (*string, bool)`

GetSourceAssertionIdOk returns a tuple with the SourceAssertionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAssertionId

`func (o *BiomarkerCdxAssertionsRow) SetSourceAssertionId(v string)`

SetSourceAssertionId sets SourceAssertionId field to given value.

### HasSourceAssertionId

`func (o *BiomarkerCdxAssertionsRow) HasSourceAssertionId() bool`

HasSourceAssertionId returns a boolean if a field has been set.

### SetSourceAssertionIdNil

`func (o *BiomarkerCdxAssertionsRow) SetSourceAssertionIdNil(b bool)`

 SetSourceAssertionIdNil sets the value for SourceAssertionId to be an explicit nil

### UnsetSourceAssertionId
`func (o *BiomarkerCdxAssertionsRow) UnsetSourceAssertionId()`

UnsetSourceAssertionId ensures that no value is present for SourceAssertionId, not even an explicit nil
### GetRouteKey

`func (o *BiomarkerCdxAssertionsRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerCdxAssertionsRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerCdxAssertionsRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerCdxAssertionsRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerCdxAssertionsRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerCdxAssertionsRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetBiomarkerId

`func (o *BiomarkerCdxAssertionsRow) GetBiomarkerId() string`

GetBiomarkerId returns the BiomarkerId field if non-nil, zero value otherwise.

### GetBiomarkerIdOk

`func (o *BiomarkerCdxAssertionsRow) GetBiomarkerIdOk() (*string, bool)`

GetBiomarkerIdOk returns a tuple with the BiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerId

`func (o *BiomarkerCdxAssertionsRow) SetBiomarkerId(v string)`

SetBiomarkerId sets BiomarkerId field to given value.

### HasBiomarkerId

`func (o *BiomarkerCdxAssertionsRow) HasBiomarkerId() bool`

HasBiomarkerId returns a boolean if a field has been set.

### SetBiomarkerIdNil

`func (o *BiomarkerCdxAssertionsRow) SetBiomarkerIdNil(b bool)`

 SetBiomarkerIdNil sets the value for BiomarkerId to be an explicit nil

### UnsetBiomarkerId
`func (o *BiomarkerCdxAssertionsRow) UnsetBiomarkerId()`

UnsetBiomarkerId ensures that no value is present for BiomarkerId, not even an explicit nil
### GetBiomarkerToken

`func (o *BiomarkerCdxAssertionsRow) GetBiomarkerToken() string`

GetBiomarkerToken returns the BiomarkerToken field if non-nil, zero value otherwise.

### GetBiomarkerTokenOk

`func (o *BiomarkerCdxAssertionsRow) GetBiomarkerTokenOk() (*string, bool)`

GetBiomarkerTokenOk returns a tuple with the BiomarkerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerToken

`func (o *BiomarkerCdxAssertionsRow) SetBiomarkerToken(v string)`

SetBiomarkerToken sets BiomarkerToken field to given value.

### HasBiomarkerToken

`func (o *BiomarkerCdxAssertionsRow) HasBiomarkerToken() bool`

HasBiomarkerToken returns a boolean if a field has been set.

### SetBiomarkerTokenNil

`func (o *BiomarkerCdxAssertionsRow) SetBiomarkerTokenNil(b bool)`

 SetBiomarkerTokenNil sets the value for BiomarkerToken to be an explicit nil

### UnsetBiomarkerToken
`func (o *BiomarkerCdxAssertionsRow) UnsetBiomarkerToken()`

UnsetBiomarkerToken ensures that no value is present for BiomarkerToken, not even an explicit nil
### GetDiagnosticName

`func (o *BiomarkerCdxAssertionsRow) GetDiagnosticName() string`

GetDiagnosticName returns the DiagnosticName field if non-nil, zero value otherwise.

### GetDiagnosticNameOk

`func (o *BiomarkerCdxAssertionsRow) GetDiagnosticNameOk() (*string, bool)`

GetDiagnosticNameOk returns a tuple with the DiagnosticName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticName

`func (o *BiomarkerCdxAssertionsRow) SetDiagnosticName(v string)`

SetDiagnosticName sets DiagnosticName field to given value.

### HasDiagnosticName

`func (o *BiomarkerCdxAssertionsRow) HasDiagnosticName() bool`

HasDiagnosticName returns a boolean if a field has been set.

### SetDiagnosticNameNil

`func (o *BiomarkerCdxAssertionsRow) SetDiagnosticNameNil(b bool)`

 SetDiagnosticNameNil sets the value for DiagnosticName to be an explicit nil

### UnsetDiagnosticName
`func (o *BiomarkerCdxAssertionsRow) UnsetDiagnosticName()`

UnsetDiagnosticName ensures that no value is present for DiagnosticName, not even an explicit nil
### GetIndication

`func (o *BiomarkerCdxAssertionsRow) GetIndication() string`

GetIndication returns the Indication field if non-nil, zero value otherwise.

### GetIndicationOk

`func (o *BiomarkerCdxAssertionsRow) GetIndicationOk() (*string, bool)`

GetIndicationOk returns a tuple with the Indication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndication

`func (o *BiomarkerCdxAssertionsRow) SetIndication(v string)`

SetIndication sets Indication field to given value.

### HasIndication

`func (o *BiomarkerCdxAssertionsRow) HasIndication() bool`

HasIndication returns a boolean if a field has been set.

### SetIndicationNil

`func (o *BiomarkerCdxAssertionsRow) SetIndicationNil(b bool)`

 SetIndicationNil sets the value for Indication to be an explicit nil

### UnsetIndication
`func (o *BiomarkerCdxAssertionsRow) UnsetIndication()`

UnsetIndication ensures that no value is present for Indication, not even an explicit nil
### GetSampleType

`func (o *BiomarkerCdxAssertionsRow) GetSampleType() string`

GetSampleType returns the SampleType field if non-nil, zero value otherwise.

### GetSampleTypeOk

`func (o *BiomarkerCdxAssertionsRow) GetSampleTypeOk() (*string, bool)`

GetSampleTypeOk returns a tuple with the SampleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleType

`func (o *BiomarkerCdxAssertionsRow) SetSampleType(v string)`

SetSampleType sets SampleType field to given value.

### HasSampleType

`func (o *BiomarkerCdxAssertionsRow) HasSampleType() bool`

HasSampleType returns a boolean if a field has been set.

### SetSampleTypeNil

`func (o *BiomarkerCdxAssertionsRow) SetSampleTypeNil(b bool)`

 SetSampleTypeNil sets the value for SampleType to be an explicit nil

### UnsetSampleType
`func (o *BiomarkerCdxAssertionsRow) UnsetSampleType()`

UnsetSampleType ensures that no value is present for SampleType, not even an explicit nil
### GetSubmissionCell

`func (o *BiomarkerCdxAssertionsRow) GetSubmissionCell() string`

GetSubmissionCell returns the SubmissionCell field if non-nil, zero value otherwise.

### GetSubmissionCellOk

`func (o *BiomarkerCdxAssertionsRow) GetSubmissionCellOk() (*string, bool)`

GetSubmissionCellOk returns a tuple with the SubmissionCell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionCell

`func (o *BiomarkerCdxAssertionsRow) SetSubmissionCell(v string)`

SetSubmissionCell sets SubmissionCell field to given value.

### HasSubmissionCell

`func (o *BiomarkerCdxAssertionsRow) HasSubmissionCell() bool`

HasSubmissionCell returns a boolean if a field has been set.

### SetSubmissionCellNil

`func (o *BiomarkerCdxAssertionsRow) SetSubmissionCellNil(b bool)`

 SetSubmissionCellNil sets the value for SubmissionCell to be an explicit nil

### UnsetSubmissionCell
`func (o *BiomarkerCdxAssertionsRow) UnsetSubmissionCell()`

UnsetSubmissionCell ensures that no value is present for SubmissionCell, not even an explicit nil
### GetLatestStatedDecisionDate

`func (o *BiomarkerCdxAssertionsRow) GetLatestStatedDecisionDate() string`

GetLatestStatedDecisionDate returns the LatestStatedDecisionDate field if non-nil, zero value otherwise.

### GetLatestStatedDecisionDateOk

`func (o *BiomarkerCdxAssertionsRow) GetLatestStatedDecisionDateOk() (*string, bool)`

GetLatestStatedDecisionDateOk returns a tuple with the LatestStatedDecisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestStatedDecisionDate

`func (o *BiomarkerCdxAssertionsRow) SetLatestStatedDecisionDate(v string)`

SetLatestStatedDecisionDate sets LatestStatedDecisionDate field to given value.

### HasLatestStatedDecisionDate

`func (o *BiomarkerCdxAssertionsRow) HasLatestStatedDecisionDate() bool`

HasLatestStatedDecisionDate returns a boolean if a field has been set.

### SetLatestStatedDecisionDateNil

`func (o *BiomarkerCdxAssertionsRow) SetLatestStatedDecisionDateNil(b bool)`

 SetLatestStatedDecisionDateNil sets the value for LatestStatedDecisionDate to be an explicit nil

### UnsetLatestStatedDecisionDate
`func (o *BiomarkerCdxAssertionsRow) UnsetLatestStatedDecisionDate()`

UnsetLatestStatedDecisionDate ensures that no value is present for LatestStatedDecisionDate, not even an explicit nil
### GetSourceUrl

`func (o *BiomarkerCdxAssertionsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *BiomarkerCdxAssertionsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *BiomarkerCdxAssertionsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *BiomarkerCdxAssertionsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *BiomarkerCdxAssertionsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *BiomarkerCdxAssertionsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetManufacturerName

`func (o *BiomarkerCdxAssertionsRow) GetManufacturerName() string`

GetManufacturerName returns the ManufacturerName field if non-nil, zero value otherwise.

### GetManufacturerNameOk

`func (o *BiomarkerCdxAssertionsRow) GetManufacturerNameOk() (*string, bool)`

GetManufacturerNameOk returns a tuple with the ManufacturerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerName

`func (o *BiomarkerCdxAssertionsRow) SetManufacturerName(v string)`

SetManufacturerName sets ManufacturerName field to given value.

### HasManufacturerName

`func (o *BiomarkerCdxAssertionsRow) HasManufacturerName() bool`

HasManufacturerName returns a boolean if a field has been set.

### SetManufacturerNameNil

`func (o *BiomarkerCdxAssertionsRow) SetManufacturerNameNil(b bool)`

 SetManufacturerNameNil sets the value for ManufacturerName to be an explicit nil

### UnsetManufacturerName
`func (o *BiomarkerCdxAssertionsRow) UnsetManufacturerName()`

UnsetManufacturerName ensures that no value is present for ManufacturerName, not even an explicit nil
### GetManufacturerOrgId

`func (o *BiomarkerCdxAssertionsRow) GetManufacturerOrgId() string`

GetManufacturerOrgId returns the ManufacturerOrgId field if non-nil, zero value otherwise.

### GetManufacturerOrgIdOk

`func (o *BiomarkerCdxAssertionsRow) GetManufacturerOrgIdOk() (*string, bool)`

GetManufacturerOrgIdOk returns a tuple with the ManufacturerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerOrgId

`func (o *BiomarkerCdxAssertionsRow) SetManufacturerOrgId(v string)`

SetManufacturerOrgId sets ManufacturerOrgId field to given value.

### HasManufacturerOrgId

`func (o *BiomarkerCdxAssertionsRow) HasManufacturerOrgId() bool`

HasManufacturerOrgId returns a boolean if a field has been set.

### SetManufacturerOrgIdNil

`func (o *BiomarkerCdxAssertionsRow) SetManufacturerOrgIdNil(b bool)`

 SetManufacturerOrgIdNil sets the value for ManufacturerOrgId to be an explicit nil

### UnsetManufacturerOrgId
`func (o *BiomarkerCdxAssertionsRow) UnsetManufacturerOrgId()`

UnsetManufacturerOrgId ensures that no value is present for ManufacturerOrgId, not even an explicit nil
### GetManufacturerCompanyId

`func (o *BiomarkerCdxAssertionsRow) GetManufacturerCompanyId() string`

GetManufacturerCompanyId returns the ManufacturerCompanyId field if non-nil, zero value otherwise.

### GetManufacturerCompanyIdOk

`func (o *BiomarkerCdxAssertionsRow) GetManufacturerCompanyIdOk() (*string, bool)`

GetManufacturerCompanyIdOk returns a tuple with the ManufacturerCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerCompanyId

`func (o *BiomarkerCdxAssertionsRow) SetManufacturerCompanyId(v string)`

SetManufacturerCompanyId sets ManufacturerCompanyId field to given value.

### HasManufacturerCompanyId

`func (o *BiomarkerCdxAssertionsRow) HasManufacturerCompanyId() bool`

HasManufacturerCompanyId returns a boolean if a field has been set.

### SetManufacturerCompanyIdNil

`func (o *BiomarkerCdxAssertionsRow) SetManufacturerCompanyIdNil(b bool)`

 SetManufacturerCompanyIdNil sets the value for ManufacturerCompanyId to be an explicit nil

### UnsetManufacturerCompanyId
`func (o *BiomarkerCdxAssertionsRow) UnsetManufacturerCompanyId()`

UnsetManufacturerCompanyId ensures that no value is present for ManufacturerCompanyId, not even an explicit nil
### GetManufacturerCompanyName

`func (o *BiomarkerCdxAssertionsRow) GetManufacturerCompanyName() string`

GetManufacturerCompanyName returns the ManufacturerCompanyName field if non-nil, zero value otherwise.

### GetManufacturerCompanyNameOk

`func (o *BiomarkerCdxAssertionsRow) GetManufacturerCompanyNameOk() (*string, bool)`

GetManufacturerCompanyNameOk returns a tuple with the ManufacturerCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerCompanyName

`func (o *BiomarkerCdxAssertionsRow) SetManufacturerCompanyName(v string)`

SetManufacturerCompanyName sets ManufacturerCompanyName field to given value.

### HasManufacturerCompanyName

`func (o *BiomarkerCdxAssertionsRow) HasManufacturerCompanyName() bool`

HasManufacturerCompanyName returns a boolean if a field has been set.

### SetManufacturerCompanyNameNil

`func (o *BiomarkerCdxAssertionsRow) SetManufacturerCompanyNameNil(b bool)`

 SetManufacturerCompanyNameNil sets the value for ManufacturerCompanyName to be an explicit nil

### UnsetManufacturerCompanyName
`func (o *BiomarkerCdxAssertionsRow) UnsetManufacturerCompanyName()`

UnsetManufacturerCompanyName ensures that no value is present for ManufacturerCompanyName, not even an explicit nil
### GetManufacturerGroupId

`func (o *BiomarkerCdxAssertionsRow) GetManufacturerGroupId() string`

GetManufacturerGroupId returns the ManufacturerGroupId field if non-nil, zero value otherwise.

### GetManufacturerGroupIdOk

`func (o *BiomarkerCdxAssertionsRow) GetManufacturerGroupIdOk() (*string, bool)`

GetManufacturerGroupIdOk returns a tuple with the ManufacturerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerGroupId

`func (o *BiomarkerCdxAssertionsRow) SetManufacturerGroupId(v string)`

SetManufacturerGroupId sets ManufacturerGroupId field to given value.

### HasManufacturerGroupId

`func (o *BiomarkerCdxAssertionsRow) HasManufacturerGroupId() bool`

HasManufacturerGroupId returns a boolean if a field has been set.

### SetManufacturerGroupIdNil

`func (o *BiomarkerCdxAssertionsRow) SetManufacturerGroupIdNil(b bool)`

 SetManufacturerGroupIdNil sets the value for ManufacturerGroupId to be an explicit nil

### UnsetManufacturerGroupId
`func (o *BiomarkerCdxAssertionsRow) UnsetManufacturerGroupId()`

UnsetManufacturerGroupId ensures that no value is present for ManufacturerGroupId, not even an explicit nil
### GetManufacturerGroupName

`func (o *BiomarkerCdxAssertionsRow) GetManufacturerGroupName() string`

GetManufacturerGroupName returns the ManufacturerGroupName field if non-nil, zero value otherwise.

### GetManufacturerGroupNameOk

`func (o *BiomarkerCdxAssertionsRow) GetManufacturerGroupNameOk() (*string, bool)`

GetManufacturerGroupNameOk returns a tuple with the ManufacturerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerGroupName

`func (o *BiomarkerCdxAssertionsRow) SetManufacturerGroupName(v string)`

SetManufacturerGroupName sets ManufacturerGroupName field to given value.

### HasManufacturerGroupName

`func (o *BiomarkerCdxAssertionsRow) HasManufacturerGroupName() bool`

HasManufacturerGroupName returns a boolean if a field has been set.

### SetManufacturerGroupNameNil

`func (o *BiomarkerCdxAssertionsRow) SetManufacturerGroupNameNil(b bool)`

 SetManufacturerGroupNameNil sets the value for ManufacturerGroupName to be an explicit nil

### UnsetManufacturerGroupName
`func (o *BiomarkerCdxAssertionsRow) UnsetManufacturerGroupName()`

UnsetManufacturerGroupName ensures that no value is present for ManufacturerGroupName, not even an explicit nil
### GetManufacturerMatchTier

`func (o *BiomarkerCdxAssertionsRow) GetManufacturerMatchTier() string`

GetManufacturerMatchTier returns the ManufacturerMatchTier field if non-nil, zero value otherwise.

### GetManufacturerMatchTierOk

`func (o *BiomarkerCdxAssertionsRow) GetManufacturerMatchTierOk() (*string, bool)`

GetManufacturerMatchTierOk returns a tuple with the ManufacturerMatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerMatchTier

`func (o *BiomarkerCdxAssertionsRow) SetManufacturerMatchTier(v string)`

SetManufacturerMatchTier sets ManufacturerMatchTier field to given value.

### HasManufacturerMatchTier

`func (o *BiomarkerCdxAssertionsRow) HasManufacturerMatchTier() bool`

HasManufacturerMatchTier returns a boolean if a field has been set.

### SetManufacturerMatchTierNil

`func (o *BiomarkerCdxAssertionsRow) SetManufacturerMatchTierNil(b bool)`

 SetManufacturerMatchTierNil sets the value for ManufacturerMatchTier to be an explicit nil

### UnsetManufacturerMatchTier
`func (o *BiomarkerCdxAssertionsRow) UnsetManufacturerMatchTier()`

UnsetManufacturerMatchTier ensures that no value is present for ManufacturerMatchTier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


