# BiomarkerPublishingCompaniesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteKey** | Pointer to **NullableString** | The biomarker route key of the gene the papers name (the gene symbol); the entry key for the biomarkers object type. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol. | [optional] 
**PublisherCompanyId** | Pointer to **NullableString** | Company level: the publisher&#39;s company. | [optional] 
**PublisherCompanyName** | Pointer to **NullableString** | Display name of the publisher&#39;s company. | [optional] 
**PublisherGroupId** | Pointer to **NullableString** | Group level: the publisher&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**PublisherGroupName** | Pointer to **NullableString** | Display name of the publisher&#39;s group. | [optional] 
**NPublications** | Pointer to **NullableInt32** | Papers that name the gene from the company&#39;s organizations (a consortium paper is excluded). | [optional] 
**NMemberOrgs** | Pointer to **NullableInt32** | The company&#39;s organizations the papers reach. | [optional] 
**BestMatchTier** | Pointer to **NullableString** | Strongest affiliation match tier. | [optional] 
**FirstPubYear** | Pointer to **NullableInt32** | Earliest publication year. | [optional] 
**LastPubYear** | Pointer to **NullableInt32** | Latest publication year. | [optional] 
**SupportingPmids** | Pointer to **[]string** | A sample of the supporting PubMed ids. | [optional] 
**SupportingPmidsTruncated** | Pointer to **NullableBool** | True when the sample is shorter than the paper count. | [optional] 
**NPapersExcludedAsConsortium** | Pointer to **NullableInt32** | Papers left out because they name too many organizations. | [optional] 

## Methods

### NewBiomarkerPublishingCompaniesRow

`func NewBiomarkerPublishingCompaniesRow() *BiomarkerPublishingCompaniesRow`

NewBiomarkerPublishingCompaniesRow instantiates a new BiomarkerPublishingCompaniesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerPublishingCompaniesRowWithDefaults

`func NewBiomarkerPublishingCompaniesRowWithDefaults() *BiomarkerPublishingCompaniesRow`

NewBiomarkerPublishingCompaniesRowWithDefaults instantiates a new BiomarkerPublishingCompaniesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteKey

`func (o *BiomarkerPublishingCompaniesRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerPublishingCompaniesRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerPublishingCompaniesRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerPublishingCompaniesRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerPublishingCompaniesRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerPublishingCompaniesRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetGeneSymbol

`func (o *BiomarkerPublishingCompaniesRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *BiomarkerPublishingCompaniesRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *BiomarkerPublishingCompaniesRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *BiomarkerPublishingCompaniesRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *BiomarkerPublishingCompaniesRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *BiomarkerPublishingCompaniesRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetPublisherCompanyId

`func (o *BiomarkerPublishingCompaniesRow) GetPublisherCompanyId() string`

GetPublisherCompanyId returns the PublisherCompanyId field if non-nil, zero value otherwise.

### GetPublisherCompanyIdOk

`func (o *BiomarkerPublishingCompaniesRow) GetPublisherCompanyIdOk() (*string, bool)`

GetPublisherCompanyIdOk returns a tuple with the PublisherCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherCompanyId

`func (o *BiomarkerPublishingCompaniesRow) SetPublisherCompanyId(v string)`

SetPublisherCompanyId sets PublisherCompanyId field to given value.

### HasPublisherCompanyId

`func (o *BiomarkerPublishingCompaniesRow) HasPublisherCompanyId() bool`

HasPublisherCompanyId returns a boolean if a field has been set.

### SetPublisherCompanyIdNil

`func (o *BiomarkerPublishingCompaniesRow) SetPublisherCompanyIdNil(b bool)`

 SetPublisherCompanyIdNil sets the value for PublisherCompanyId to be an explicit nil

### UnsetPublisherCompanyId
`func (o *BiomarkerPublishingCompaniesRow) UnsetPublisherCompanyId()`

UnsetPublisherCompanyId ensures that no value is present for PublisherCompanyId, not even an explicit nil
### GetPublisherCompanyName

`func (o *BiomarkerPublishingCompaniesRow) GetPublisherCompanyName() string`

GetPublisherCompanyName returns the PublisherCompanyName field if non-nil, zero value otherwise.

### GetPublisherCompanyNameOk

`func (o *BiomarkerPublishingCompaniesRow) GetPublisherCompanyNameOk() (*string, bool)`

GetPublisherCompanyNameOk returns a tuple with the PublisherCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherCompanyName

`func (o *BiomarkerPublishingCompaniesRow) SetPublisherCompanyName(v string)`

SetPublisherCompanyName sets PublisherCompanyName field to given value.

### HasPublisherCompanyName

`func (o *BiomarkerPublishingCompaniesRow) HasPublisherCompanyName() bool`

HasPublisherCompanyName returns a boolean if a field has been set.

### SetPublisherCompanyNameNil

`func (o *BiomarkerPublishingCompaniesRow) SetPublisherCompanyNameNil(b bool)`

 SetPublisherCompanyNameNil sets the value for PublisherCompanyName to be an explicit nil

### UnsetPublisherCompanyName
`func (o *BiomarkerPublishingCompaniesRow) UnsetPublisherCompanyName()`

UnsetPublisherCompanyName ensures that no value is present for PublisherCompanyName, not even an explicit nil
### GetPublisherGroupId

`func (o *BiomarkerPublishingCompaniesRow) GetPublisherGroupId() string`

GetPublisherGroupId returns the PublisherGroupId field if non-nil, zero value otherwise.

### GetPublisherGroupIdOk

`func (o *BiomarkerPublishingCompaniesRow) GetPublisherGroupIdOk() (*string, bool)`

GetPublisherGroupIdOk returns a tuple with the PublisherGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherGroupId

`func (o *BiomarkerPublishingCompaniesRow) SetPublisherGroupId(v string)`

SetPublisherGroupId sets PublisherGroupId field to given value.

### HasPublisherGroupId

`func (o *BiomarkerPublishingCompaniesRow) HasPublisherGroupId() bool`

HasPublisherGroupId returns a boolean if a field has been set.

### SetPublisherGroupIdNil

`func (o *BiomarkerPublishingCompaniesRow) SetPublisherGroupIdNil(b bool)`

 SetPublisherGroupIdNil sets the value for PublisherGroupId to be an explicit nil

### UnsetPublisherGroupId
`func (o *BiomarkerPublishingCompaniesRow) UnsetPublisherGroupId()`

UnsetPublisherGroupId ensures that no value is present for PublisherGroupId, not even an explicit nil
### GetPublisherGroupName

`func (o *BiomarkerPublishingCompaniesRow) GetPublisherGroupName() string`

GetPublisherGroupName returns the PublisherGroupName field if non-nil, zero value otherwise.

### GetPublisherGroupNameOk

`func (o *BiomarkerPublishingCompaniesRow) GetPublisherGroupNameOk() (*string, bool)`

GetPublisherGroupNameOk returns a tuple with the PublisherGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherGroupName

`func (o *BiomarkerPublishingCompaniesRow) SetPublisherGroupName(v string)`

SetPublisherGroupName sets PublisherGroupName field to given value.

### HasPublisherGroupName

`func (o *BiomarkerPublishingCompaniesRow) HasPublisherGroupName() bool`

HasPublisherGroupName returns a boolean if a field has been set.

### SetPublisherGroupNameNil

`func (o *BiomarkerPublishingCompaniesRow) SetPublisherGroupNameNil(b bool)`

 SetPublisherGroupNameNil sets the value for PublisherGroupName to be an explicit nil

### UnsetPublisherGroupName
`func (o *BiomarkerPublishingCompaniesRow) UnsetPublisherGroupName()`

UnsetPublisherGroupName ensures that no value is present for PublisherGroupName, not even an explicit nil
### GetNPublications

`func (o *BiomarkerPublishingCompaniesRow) GetNPublications() int32`

GetNPublications returns the NPublications field if non-nil, zero value otherwise.

### GetNPublicationsOk

`func (o *BiomarkerPublishingCompaniesRow) GetNPublicationsOk() (*int32, bool)`

GetNPublicationsOk returns a tuple with the NPublications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPublications

`func (o *BiomarkerPublishingCompaniesRow) SetNPublications(v int32)`

SetNPublications sets NPublications field to given value.

### HasNPublications

`func (o *BiomarkerPublishingCompaniesRow) HasNPublications() bool`

HasNPublications returns a boolean if a field has been set.

### SetNPublicationsNil

`func (o *BiomarkerPublishingCompaniesRow) SetNPublicationsNil(b bool)`

 SetNPublicationsNil sets the value for NPublications to be an explicit nil

### UnsetNPublications
`func (o *BiomarkerPublishingCompaniesRow) UnsetNPublications()`

UnsetNPublications ensures that no value is present for NPublications, not even an explicit nil
### GetNMemberOrgs

`func (o *BiomarkerPublishingCompaniesRow) GetNMemberOrgs() int32`

GetNMemberOrgs returns the NMemberOrgs field if non-nil, zero value otherwise.

### GetNMemberOrgsOk

`func (o *BiomarkerPublishingCompaniesRow) GetNMemberOrgsOk() (*int32, bool)`

GetNMemberOrgsOk returns a tuple with the NMemberOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNMemberOrgs

`func (o *BiomarkerPublishingCompaniesRow) SetNMemberOrgs(v int32)`

SetNMemberOrgs sets NMemberOrgs field to given value.

### HasNMemberOrgs

`func (o *BiomarkerPublishingCompaniesRow) HasNMemberOrgs() bool`

HasNMemberOrgs returns a boolean if a field has been set.

### SetNMemberOrgsNil

`func (o *BiomarkerPublishingCompaniesRow) SetNMemberOrgsNil(b bool)`

 SetNMemberOrgsNil sets the value for NMemberOrgs to be an explicit nil

### UnsetNMemberOrgs
`func (o *BiomarkerPublishingCompaniesRow) UnsetNMemberOrgs()`

UnsetNMemberOrgs ensures that no value is present for NMemberOrgs, not even an explicit nil
### GetBestMatchTier

`func (o *BiomarkerPublishingCompaniesRow) GetBestMatchTier() string`

GetBestMatchTier returns the BestMatchTier field if non-nil, zero value otherwise.

### GetBestMatchTierOk

`func (o *BiomarkerPublishingCompaniesRow) GetBestMatchTierOk() (*string, bool)`

GetBestMatchTierOk returns a tuple with the BestMatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestMatchTier

`func (o *BiomarkerPublishingCompaniesRow) SetBestMatchTier(v string)`

SetBestMatchTier sets BestMatchTier field to given value.

### HasBestMatchTier

`func (o *BiomarkerPublishingCompaniesRow) HasBestMatchTier() bool`

HasBestMatchTier returns a boolean if a field has been set.

### SetBestMatchTierNil

`func (o *BiomarkerPublishingCompaniesRow) SetBestMatchTierNil(b bool)`

 SetBestMatchTierNil sets the value for BestMatchTier to be an explicit nil

### UnsetBestMatchTier
`func (o *BiomarkerPublishingCompaniesRow) UnsetBestMatchTier()`

UnsetBestMatchTier ensures that no value is present for BestMatchTier, not even an explicit nil
### GetFirstPubYear

`func (o *BiomarkerPublishingCompaniesRow) GetFirstPubYear() int32`

GetFirstPubYear returns the FirstPubYear field if non-nil, zero value otherwise.

### GetFirstPubYearOk

`func (o *BiomarkerPublishingCompaniesRow) GetFirstPubYearOk() (*int32, bool)`

GetFirstPubYearOk returns a tuple with the FirstPubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPubYear

`func (o *BiomarkerPublishingCompaniesRow) SetFirstPubYear(v int32)`

SetFirstPubYear sets FirstPubYear field to given value.

### HasFirstPubYear

`func (o *BiomarkerPublishingCompaniesRow) HasFirstPubYear() bool`

HasFirstPubYear returns a boolean if a field has been set.

### SetFirstPubYearNil

`func (o *BiomarkerPublishingCompaniesRow) SetFirstPubYearNil(b bool)`

 SetFirstPubYearNil sets the value for FirstPubYear to be an explicit nil

### UnsetFirstPubYear
`func (o *BiomarkerPublishingCompaniesRow) UnsetFirstPubYear()`

UnsetFirstPubYear ensures that no value is present for FirstPubYear, not even an explicit nil
### GetLastPubYear

`func (o *BiomarkerPublishingCompaniesRow) GetLastPubYear() int32`

GetLastPubYear returns the LastPubYear field if non-nil, zero value otherwise.

### GetLastPubYearOk

`func (o *BiomarkerPublishingCompaniesRow) GetLastPubYearOk() (*int32, bool)`

GetLastPubYearOk returns a tuple with the LastPubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPubYear

`func (o *BiomarkerPublishingCompaniesRow) SetLastPubYear(v int32)`

SetLastPubYear sets LastPubYear field to given value.

### HasLastPubYear

`func (o *BiomarkerPublishingCompaniesRow) HasLastPubYear() bool`

HasLastPubYear returns a boolean if a field has been set.

### SetLastPubYearNil

`func (o *BiomarkerPublishingCompaniesRow) SetLastPubYearNil(b bool)`

 SetLastPubYearNil sets the value for LastPubYear to be an explicit nil

### UnsetLastPubYear
`func (o *BiomarkerPublishingCompaniesRow) UnsetLastPubYear()`

UnsetLastPubYear ensures that no value is present for LastPubYear, not even an explicit nil
### GetSupportingPmids

`func (o *BiomarkerPublishingCompaniesRow) GetSupportingPmids() []string`

GetSupportingPmids returns the SupportingPmids field if non-nil, zero value otherwise.

### GetSupportingPmidsOk

`func (o *BiomarkerPublishingCompaniesRow) GetSupportingPmidsOk() (*[]string, bool)`

GetSupportingPmidsOk returns a tuple with the SupportingPmids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportingPmids

`func (o *BiomarkerPublishingCompaniesRow) SetSupportingPmids(v []string)`

SetSupportingPmids sets SupportingPmids field to given value.

### HasSupportingPmids

`func (o *BiomarkerPublishingCompaniesRow) HasSupportingPmids() bool`

HasSupportingPmids returns a boolean if a field has been set.

### SetSupportingPmidsNil

`func (o *BiomarkerPublishingCompaniesRow) SetSupportingPmidsNil(b bool)`

 SetSupportingPmidsNil sets the value for SupportingPmids to be an explicit nil

### UnsetSupportingPmids
`func (o *BiomarkerPublishingCompaniesRow) UnsetSupportingPmids()`

UnsetSupportingPmids ensures that no value is present for SupportingPmids, not even an explicit nil
### GetSupportingPmidsTruncated

`func (o *BiomarkerPublishingCompaniesRow) GetSupportingPmidsTruncated() bool`

GetSupportingPmidsTruncated returns the SupportingPmidsTruncated field if non-nil, zero value otherwise.

### GetSupportingPmidsTruncatedOk

`func (o *BiomarkerPublishingCompaniesRow) GetSupportingPmidsTruncatedOk() (*bool, bool)`

GetSupportingPmidsTruncatedOk returns a tuple with the SupportingPmidsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportingPmidsTruncated

`func (o *BiomarkerPublishingCompaniesRow) SetSupportingPmidsTruncated(v bool)`

SetSupportingPmidsTruncated sets SupportingPmidsTruncated field to given value.

### HasSupportingPmidsTruncated

`func (o *BiomarkerPublishingCompaniesRow) HasSupportingPmidsTruncated() bool`

HasSupportingPmidsTruncated returns a boolean if a field has been set.

### SetSupportingPmidsTruncatedNil

`func (o *BiomarkerPublishingCompaniesRow) SetSupportingPmidsTruncatedNil(b bool)`

 SetSupportingPmidsTruncatedNil sets the value for SupportingPmidsTruncated to be an explicit nil

### UnsetSupportingPmidsTruncated
`func (o *BiomarkerPublishingCompaniesRow) UnsetSupportingPmidsTruncated()`

UnsetSupportingPmidsTruncated ensures that no value is present for SupportingPmidsTruncated, not even an explicit nil
### GetNPapersExcludedAsConsortium

`func (o *BiomarkerPublishingCompaniesRow) GetNPapersExcludedAsConsortium() int32`

GetNPapersExcludedAsConsortium returns the NPapersExcludedAsConsortium field if non-nil, zero value otherwise.

### GetNPapersExcludedAsConsortiumOk

`func (o *BiomarkerPublishingCompaniesRow) GetNPapersExcludedAsConsortiumOk() (*int32, bool)`

GetNPapersExcludedAsConsortiumOk returns a tuple with the NPapersExcludedAsConsortium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPapersExcludedAsConsortium

`func (o *BiomarkerPublishingCompaniesRow) SetNPapersExcludedAsConsortium(v int32)`

SetNPapersExcludedAsConsortium sets NPapersExcludedAsConsortium field to given value.

### HasNPapersExcludedAsConsortium

`func (o *BiomarkerPublishingCompaniesRow) HasNPapersExcludedAsConsortium() bool`

HasNPapersExcludedAsConsortium returns a boolean if a field has been set.

### SetNPapersExcludedAsConsortiumNil

`func (o *BiomarkerPublishingCompaniesRow) SetNPapersExcludedAsConsortiumNil(b bool)`

 SetNPapersExcludedAsConsortiumNil sets the value for NPapersExcludedAsConsortium to be an explicit nil

### UnsetNPapersExcludedAsConsortium
`func (o *BiomarkerPublishingCompaniesRow) UnsetNPapersExcludedAsConsortium()`

UnsetNPapersExcludedAsConsortium ensures that no value is present for NPapersExcludedAsConsortium, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


