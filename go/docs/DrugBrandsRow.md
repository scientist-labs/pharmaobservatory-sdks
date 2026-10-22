# DrugBrandsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**BrandKey** | Pointer to **NullableString** | The brand key: an upper-case alphanumeric trade-name token (the brand&#39;s leading name tokens), the same key the brands pages use. Not a display name; brand_name is the display form. | [optional] 
**BrandName** | Pointer to **NullableString** | The brand&#39;s display name (the brand key in display case). | [optional] 
**MoleculeName** | Pointer to **NullableString** | The display name of the drug the brand is a brand of. | [optional] 
**ApprovedAnywhere** | Pointer to **NullableBool** | True when the drug is approved in at least one jurisdiction. | [optional] 
**IsMarketed** | Pointer to **NullableBool** | True when the drug, or a record folded under it, has a package in the US National Drug Code (NDC) directory. | [optional] 
**AtcClass** | Pointer to **NullableString** | The top-level Anatomical Therapeutic Chemical (ATC) class of the drug. | [optional] 
**IsBiologic** | Pointer to **NullableBool** | True when the drug is a biologic. | [optional] 
**FirstApprovalYear** | Pointer to **NullableInt32** | The earliest approval year ChEMBL records for the drug. | [optional] 
**HolderName** | Pointer to **NullableString** | The FDA approval holder&#39;s name as filed, present only when the brand has exactly one holder name. | [optional] 
**HolderOrgId** | Pointer to **NullableString** | Legal-entity level: the approval holder&#39;s organization record (null when the brand has no holder, more than one, or an unresolved one). | [optional] 
**HolderCompanyId** | Pointer to **NullableString** | Company level: the approval holder&#39;s company (null when unresolved). | [optional] 
**HolderCompanyName** | Pointer to **NullableString** | Display name of the approval holder&#39;s company. | [optional] 
**HolderGroupId** | Pointer to **NullableString** | Group level: the approval holder&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**HolderGroupName** | Pointer to **NullableString** | Display name of the approval holder&#39;s group. | [optional] 
**HolderMatchScope** | Pointer to **NullableString** | How firmly the holder name resolved to the organization (spine_unique, anchored and establishment_registry are folded onto an organization the spine already held; anchored_scoped, registry_scoped, alias_scoped and none are newly minted organizations). Null when the brand has no single holder organization. | [optional] 
**NHolderNames** | Pointer to **NullableInt32** | Distinct approval-holder names FDA publishes for the brand (0 for a brand with no FDA applicant, an RxNorm-only brand). | [optional] 
**NHolderOrgs** | Pointer to **NullableInt32** | Distinct organizations those names resolve to; above 1 with a null holder_org_id means more than one company holds the brand. | [optional] 
**HolderAsOf** | Pointer to **NullableString** | The ingest snapshot date the holder was read from. | [optional] 
**HeadlineIndication** | Pointer to **NullableString** | The brand&#39;s headline indication: the one most often named in the filings that disclose its revenue. | [optional] 
**LatestRevenueUsdMillions** | Pointer to **NullableFloat64** | The brand&#39;s latest disclosed worldwide revenue in USD millions (null when none is disclosed). | [optional] 
**LatestRevenueYear** | Pointer to **NullableInt32** | The fiscal year of latest_revenue_usd_millions. | [optional] 

## Methods

### NewDrugBrandsRow

`func NewDrugBrandsRow() *DrugBrandsRow`

NewDrugBrandsRow instantiates a new DrugBrandsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugBrandsRowWithDefaults

`func NewDrugBrandsRowWithDefaults() *DrugBrandsRow`

NewDrugBrandsRowWithDefaults instantiates a new DrugBrandsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugBrandsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugBrandsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugBrandsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugBrandsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugBrandsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugBrandsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetBrandKey

`func (o *DrugBrandsRow) GetBrandKey() string`

GetBrandKey returns the BrandKey field if non-nil, zero value otherwise.

### GetBrandKeyOk

`func (o *DrugBrandsRow) GetBrandKeyOk() (*string, bool)`

GetBrandKeyOk returns a tuple with the BrandKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandKey

`func (o *DrugBrandsRow) SetBrandKey(v string)`

SetBrandKey sets BrandKey field to given value.

### HasBrandKey

`func (o *DrugBrandsRow) HasBrandKey() bool`

HasBrandKey returns a boolean if a field has been set.

### SetBrandKeyNil

`func (o *DrugBrandsRow) SetBrandKeyNil(b bool)`

 SetBrandKeyNil sets the value for BrandKey to be an explicit nil

### UnsetBrandKey
`func (o *DrugBrandsRow) UnsetBrandKey()`

UnsetBrandKey ensures that no value is present for BrandKey, not even an explicit nil
### GetBrandName

`func (o *DrugBrandsRow) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *DrugBrandsRow) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *DrugBrandsRow) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.

### HasBrandName

`func (o *DrugBrandsRow) HasBrandName() bool`

HasBrandName returns a boolean if a field has been set.

### SetBrandNameNil

`func (o *DrugBrandsRow) SetBrandNameNil(b bool)`

 SetBrandNameNil sets the value for BrandName to be an explicit nil

### UnsetBrandName
`func (o *DrugBrandsRow) UnsetBrandName()`

UnsetBrandName ensures that no value is present for BrandName, not even an explicit nil
### GetMoleculeName

`func (o *DrugBrandsRow) GetMoleculeName() string`

GetMoleculeName returns the MoleculeName field if non-nil, zero value otherwise.

### GetMoleculeNameOk

`func (o *DrugBrandsRow) GetMoleculeNameOk() (*string, bool)`

GetMoleculeNameOk returns a tuple with the MoleculeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeName

`func (o *DrugBrandsRow) SetMoleculeName(v string)`

SetMoleculeName sets MoleculeName field to given value.

### HasMoleculeName

`func (o *DrugBrandsRow) HasMoleculeName() bool`

HasMoleculeName returns a boolean if a field has been set.

### SetMoleculeNameNil

`func (o *DrugBrandsRow) SetMoleculeNameNil(b bool)`

 SetMoleculeNameNil sets the value for MoleculeName to be an explicit nil

### UnsetMoleculeName
`func (o *DrugBrandsRow) UnsetMoleculeName()`

UnsetMoleculeName ensures that no value is present for MoleculeName, not even an explicit nil
### GetApprovedAnywhere

`func (o *DrugBrandsRow) GetApprovedAnywhere() bool`

GetApprovedAnywhere returns the ApprovedAnywhere field if non-nil, zero value otherwise.

### GetApprovedAnywhereOk

`func (o *DrugBrandsRow) GetApprovedAnywhereOk() (*bool, bool)`

GetApprovedAnywhereOk returns a tuple with the ApprovedAnywhere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAnywhere

`func (o *DrugBrandsRow) SetApprovedAnywhere(v bool)`

SetApprovedAnywhere sets ApprovedAnywhere field to given value.

### HasApprovedAnywhere

`func (o *DrugBrandsRow) HasApprovedAnywhere() bool`

HasApprovedAnywhere returns a boolean if a field has been set.

### SetApprovedAnywhereNil

`func (o *DrugBrandsRow) SetApprovedAnywhereNil(b bool)`

 SetApprovedAnywhereNil sets the value for ApprovedAnywhere to be an explicit nil

### UnsetApprovedAnywhere
`func (o *DrugBrandsRow) UnsetApprovedAnywhere()`

UnsetApprovedAnywhere ensures that no value is present for ApprovedAnywhere, not even an explicit nil
### GetIsMarketed

`func (o *DrugBrandsRow) GetIsMarketed() bool`

GetIsMarketed returns the IsMarketed field if non-nil, zero value otherwise.

### GetIsMarketedOk

`func (o *DrugBrandsRow) GetIsMarketedOk() (*bool, bool)`

GetIsMarketedOk returns a tuple with the IsMarketed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarketed

`func (o *DrugBrandsRow) SetIsMarketed(v bool)`

SetIsMarketed sets IsMarketed field to given value.

### HasIsMarketed

`func (o *DrugBrandsRow) HasIsMarketed() bool`

HasIsMarketed returns a boolean if a field has been set.

### SetIsMarketedNil

`func (o *DrugBrandsRow) SetIsMarketedNil(b bool)`

 SetIsMarketedNil sets the value for IsMarketed to be an explicit nil

### UnsetIsMarketed
`func (o *DrugBrandsRow) UnsetIsMarketed()`

UnsetIsMarketed ensures that no value is present for IsMarketed, not even an explicit nil
### GetAtcClass

`func (o *DrugBrandsRow) GetAtcClass() string`

GetAtcClass returns the AtcClass field if non-nil, zero value otherwise.

### GetAtcClassOk

`func (o *DrugBrandsRow) GetAtcClassOk() (*string, bool)`

GetAtcClassOk returns a tuple with the AtcClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcClass

`func (o *DrugBrandsRow) SetAtcClass(v string)`

SetAtcClass sets AtcClass field to given value.

### HasAtcClass

`func (o *DrugBrandsRow) HasAtcClass() bool`

HasAtcClass returns a boolean if a field has been set.

### SetAtcClassNil

`func (o *DrugBrandsRow) SetAtcClassNil(b bool)`

 SetAtcClassNil sets the value for AtcClass to be an explicit nil

### UnsetAtcClass
`func (o *DrugBrandsRow) UnsetAtcClass()`

UnsetAtcClass ensures that no value is present for AtcClass, not even an explicit nil
### GetIsBiologic

`func (o *DrugBrandsRow) GetIsBiologic() bool`

GetIsBiologic returns the IsBiologic field if non-nil, zero value otherwise.

### GetIsBiologicOk

`func (o *DrugBrandsRow) GetIsBiologicOk() (*bool, bool)`

GetIsBiologicOk returns a tuple with the IsBiologic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBiologic

`func (o *DrugBrandsRow) SetIsBiologic(v bool)`

SetIsBiologic sets IsBiologic field to given value.

### HasIsBiologic

`func (o *DrugBrandsRow) HasIsBiologic() bool`

HasIsBiologic returns a boolean if a field has been set.

### SetIsBiologicNil

`func (o *DrugBrandsRow) SetIsBiologicNil(b bool)`

 SetIsBiologicNil sets the value for IsBiologic to be an explicit nil

### UnsetIsBiologic
`func (o *DrugBrandsRow) UnsetIsBiologic()`

UnsetIsBiologic ensures that no value is present for IsBiologic, not even an explicit nil
### GetFirstApprovalYear

`func (o *DrugBrandsRow) GetFirstApprovalYear() int32`

GetFirstApprovalYear returns the FirstApprovalYear field if non-nil, zero value otherwise.

### GetFirstApprovalYearOk

`func (o *DrugBrandsRow) GetFirstApprovalYearOk() (*int32, bool)`

GetFirstApprovalYearOk returns a tuple with the FirstApprovalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstApprovalYear

`func (o *DrugBrandsRow) SetFirstApprovalYear(v int32)`

SetFirstApprovalYear sets FirstApprovalYear field to given value.

### HasFirstApprovalYear

`func (o *DrugBrandsRow) HasFirstApprovalYear() bool`

HasFirstApprovalYear returns a boolean if a field has been set.

### SetFirstApprovalYearNil

`func (o *DrugBrandsRow) SetFirstApprovalYearNil(b bool)`

 SetFirstApprovalYearNil sets the value for FirstApprovalYear to be an explicit nil

### UnsetFirstApprovalYear
`func (o *DrugBrandsRow) UnsetFirstApprovalYear()`

UnsetFirstApprovalYear ensures that no value is present for FirstApprovalYear, not even an explicit nil
### GetHolderName

`func (o *DrugBrandsRow) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *DrugBrandsRow) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *DrugBrandsRow) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.

### HasHolderName

`func (o *DrugBrandsRow) HasHolderName() bool`

HasHolderName returns a boolean if a field has been set.

### SetHolderNameNil

`func (o *DrugBrandsRow) SetHolderNameNil(b bool)`

 SetHolderNameNil sets the value for HolderName to be an explicit nil

### UnsetHolderName
`func (o *DrugBrandsRow) UnsetHolderName()`

UnsetHolderName ensures that no value is present for HolderName, not even an explicit nil
### GetHolderOrgId

`func (o *DrugBrandsRow) GetHolderOrgId() string`

GetHolderOrgId returns the HolderOrgId field if non-nil, zero value otherwise.

### GetHolderOrgIdOk

`func (o *DrugBrandsRow) GetHolderOrgIdOk() (*string, bool)`

GetHolderOrgIdOk returns a tuple with the HolderOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderOrgId

`func (o *DrugBrandsRow) SetHolderOrgId(v string)`

SetHolderOrgId sets HolderOrgId field to given value.

### HasHolderOrgId

`func (o *DrugBrandsRow) HasHolderOrgId() bool`

HasHolderOrgId returns a boolean if a field has been set.

### SetHolderOrgIdNil

`func (o *DrugBrandsRow) SetHolderOrgIdNil(b bool)`

 SetHolderOrgIdNil sets the value for HolderOrgId to be an explicit nil

### UnsetHolderOrgId
`func (o *DrugBrandsRow) UnsetHolderOrgId()`

UnsetHolderOrgId ensures that no value is present for HolderOrgId, not even an explicit nil
### GetHolderCompanyId

`func (o *DrugBrandsRow) GetHolderCompanyId() string`

GetHolderCompanyId returns the HolderCompanyId field if non-nil, zero value otherwise.

### GetHolderCompanyIdOk

`func (o *DrugBrandsRow) GetHolderCompanyIdOk() (*string, bool)`

GetHolderCompanyIdOk returns a tuple with the HolderCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderCompanyId

`func (o *DrugBrandsRow) SetHolderCompanyId(v string)`

SetHolderCompanyId sets HolderCompanyId field to given value.

### HasHolderCompanyId

`func (o *DrugBrandsRow) HasHolderCompanyId() bool`

HasHolderCompanyId returns a boolean if a field has been set.

### SetHolderCompanyIdNil

`func (o *DrugBrandsRow) SetHolderCompanyIdNil(b bool)`

 SetHolderCompanyIdNil sets the value for HolderCompanyId to be an explicit nil

### UnsetHolderCompanyId
`func (o *DrugBrandsRow) UnsetHolderCompanyId()`

UnsetHolderCompanyId ensures that no value is present for HolderCompanyId, not even an explicit nil
### GetHolderCompanyName

`func (o *DrugBrandsRow) GetHolderCompanyName() string`

GetHolderCompanyName returns the HolderCompanyName field if non-nil, zero value otherwise.

### GetHolderCompanyNameOk

`func (o *DrugBrandsRow) GetHolderCompanyNameOk() (*string, bool)`

GetHolderCompanyNameOk returns a tuple with the HolderCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderCompanyName

`func (o *DrugBrandsRow) SetHolderCompanyName(v string)`

SetHolderCompanyName sets HolderCompanyName field to given value.

### HasHolderCompanyName

`func (o *DrugBrandsRow) HasHolderCompanyName() bool`

HasHolderCompanyName returns a boolean if a field has been set.

### SetHolderCompanyNameNil

`func (o *DrugBrandsRow) SetHolderCompanyNameNil(b bool)`

 SetHolderCompanyNameNil sets the value for HolderCompanyName to be an explicit nil

### UnsetHolderCompanyName
`func (o *DrugBrandsRow) UnsetHolderCompanyName()`

UnsetHolderCompanyName ensures that no value is present for HolderCompanyName, not even an explicit nil
### GetHolderGroupId

`func (o *DrugBrandsRow) GetHolderGroupId() string`

GetHolderGroupId returns the HolderGroupId field if non-nil, zero value otherwise.

### GetHolderGroupIdOk

`func (o *DrugBrandsRow) GetHolderGroupIdOk() (*string, bool)`

GetHolderGroupIdOk returns a tuple with the HolderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderGroupId

`func (o *DrugBrandsRow) SetHolderGroupId(v string)`

SetHolderGroupId sets HolderGroupId field to given value.

### HasHolderGroupId

`func (o *DrugBrandsRow) HasHolderGroupId() bool`

HasHolderGroupId returns a boolean if a field has been set.

### SetHolderGroupIdNil

`func (o *DrugBrandsRow) SetHolderGroupIdNil(b bool)`

 SetHolderGroupIdNil sets the value for HolderGroupId to be an explicit nil

### UnsetHolderGroupId
`func (o *DrugBrandsRow) UnsetHolderGroupId()`

UnsetHolderGroupId ensures that no value is present for HolderGroupId, not even an explicit nil
### GetHolderGroupName

`func (o *DrugBrandsRow) GetHolderGroupName() string`

GetHolderGroupName returns the HolderGroupName field if non-nil, zero value otherwise.

### GetHolderGroupNameOk

`func (o *DrugBrandsRow) GetHolderGroupNameOk() (*string, bool)`

GetHolderGroupNameOk returns a tuple with the HolderGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderGroupName

`func (o *DrugBrandsRow) SetHolderGroupName(v string)`

SetHolderGroupName sets HolderGroupName field to given value.

### HasHolderGroupName

`func (o *DrugBrandsRow) HasHolderGroupName() bool`

HasHolderGroupName returns a boolean if a field has been set.

### SetHolderGroupNameNil

`func (o *DrugBrandsRow) SetHolderGroupNameNil(b bool)`

 SetHolderGroupNameNil sets the value for HolderGroupName to be an explicit nil

### UnsetHolderGroupName
`func (o *DrugBrandsRow) UnsetHolderGroupName()`

UnsetHolderGroupName ensures that no value is present for HolderGroupName, not even an explicit nil
### GetHolderMatchScope

`func (o *DrugBrandsRow) GetHolderMatchScope() string`

GetHolderMatchScope returns the HolderMatchScope field if non-nil, zero value otherwise.

### GetHolderMatchScopeOk

`func (o *DrugBrandsRow) GetHolderMatchScopeOk() (*string, bool)`

GetHolderMatchScopeOk returns a tuple with the HolderMatchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderMatchScope

`func (o *DrugBrandsRow) SetHolderMatchScope(v string)`

SetHolderMatchScope sets HolderMatchScope field to given value.

### HasHolderMatchScope

`func (o *DrugBrandsRow) HasHolderMatchScope() bool`

HasHolderMatchScope returns a boolean if a field has been set.

### SetHolderMatchScopeNil

`func (o *DrugBrandsRow) SetHolderMatchScopeNil(b bool)`

 SetHolderMatchScopeNil sets the value for HolderMatchScope to be an explicit nil

### UnsetHolderMatchScope
`func (o *DrugBrandsRow) UnsetHolderMatchScope()`

UnsetHolderMatchScope ensures that no value is present for HolderMatchScope, not even an explicit nil
### GetNHolderNames

`func (o *DrugBrandsRow) GetNHolderNames() int32`

GetNHolderNames returns the NHolderNames field if non-nil, zero value otherwise.

### GetNHolderNamesOk

`func (o *DrugBrandsRow) GetNHolderNamesOk() (*int32, bool)`

GetNHolderNamesOk returns a tuple with the NHolderNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNHolderNames

`func (o *DrugBrandsRow) SetNHolderNames(v int32)`

SetNHolderNames sets NHolderNames field to given value.

### HasNHolderNames

`func (o *DrugBrandsRow) HasNHolderNames() bool`

HasNHolderNames returns a boolean if a field has been set.

### SetNHolderNamesNil

`func (o *DrugBrandsRow) SetNHolderNamesNil(b bool)`

 SetNHolderNamesNil sets the value for NHolderNames to be an explicit nil

### UnsetNHolderNames
`func (o *DrugBrandsRow) UnsetNHolderNames()`

UnsetNHolderNames ensures that no value is present for NHolderNames, not even an explicit nil
### GetNHolderOrgs

`func (o *DrugBrandsRow) GetNHolderOrgs() int32`

GetNHolderOrgs returns the NHolderOrgs field if non-nil, zero value otherwise.

### GetNHolderOrgsOk

`func (o *DrugBrandsRow) GetNHolderOrgsOk() (*int32, bool)`

GetNHolderOrgsOk returns a tuple with the NHolderOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNHolderOrgs

`func (o *DrugBrandsRow) SetNHolderOrgs(v int32)`

SetNHolderOrgs sets NHolderOrgs field to given value.

### HasNHolderOrgs

`func (o *DrugBrandsRow) HasNHolderOrgs() bool`

HasNHolderOrgs returns a boolean if a field has been set.

### SetNHolderOrgsNil

`func (o *DrugBrandsRow) SetNHolderOrgsNil(b bool)`

 SetNHolderOrgsNil sets the value for NHolderOrgs to be an explicit nil

### UnsetNHolderOrgs
`func (o *DrugBrandsRow) UnsetNHolderOrgs()`

UnsetNHolderOrgs ensures that no value is present for NHolderOrgs, not even an explicit nil
### GetHolderAsOf

`func (o *DrugBrandsRow) GetHolderAsOf() string`

GetHolderAsOf returns the HolderAsOf field if non-nil, zero value otherwise.

### GetHolderAsOfOk

`func (o *DrugBrandsRow) GetHolderAsOfOk() (*string, bool)`

GetHolderAsOfOk returns a tuple with the HolderAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderAsOf

`func (o *DrugBrandsRow) SetHolderAsOf(v string)`

SetHolderAsOf sets HolderAsOf field to given value.

### HasHolderAsOf

`func (o *DrugBrandsRow) HasHolderAsOf() bool`

HasHolderAsOf returns a boolean if a field has been set.

### SetHolderAsOfNil

`func (o *DrugBrandsRow) SetHolderAsOfNil(b bool)`

 SetHolderAsOfNil sets the value for HolderAsOf to be an explicit nil

### UnsetHolderAsOf
`func (o *DrugBrandsRow) UnsetHolderAsOf()`

UnsetHolderAsOf ensures that no value is present for HolderAsOf, not even an explicit nil
### GetHeadlineIndication

`func (o *DrugBrandsRow) GetHeadlineIndication() string`

GetHeadlineIndication returns the HeadlineIndication field if non-nil, zero value otherwise.

### GetHeadlineIndicationOk

`func (o *DrugBrandsRow) GetHeadlineIndicationOk() (*string, bool)`

GetHeadlineIndicationOk returns a tuple with the HeadlineIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadlineIndication

`func (o *DrugBrandsRow) SetHeadlineIndication(v string)`

SetHeadlineIndication sets HeadlineIndication field to given value.

### HasHeadlineIndication

`func (o *DrugBrandsRow) HasHeadlineIndication() bool`

HasHeadlineIndication returns a boolean if a field has been set.

### SetHeadlineIndicationNil

`func (o *DrugBrandsRow) SetHeadlineIndicationNil(b bool)`

 SetHeadlineIndicationNil sets the value for HeadlineIndication to be an explicit nil

### UnsetHeadlineIndication
`func (o *DrugBrandsRow) UnsetHeadlineIndication()`

UnsetHeadlineIndication ensures that no value is present for HeadlineIndication, not even an explicit nil
### GetLatestRevenueUsdMillions

`func (o *DrugBrandsRow) GetLatestRevenueUsdMillions() float64`

GetLatestRevenueUsdMillions returns the LatestRevenueUsdMillions field if non-nil, zero value otherwise.

### GetLatestRevenueUsdMillionsOk

`func (o *DrugBrandsRow) GetLatestRevenueUsdMillionsOk() (*float64, bool)`

GetLatestRevenueUsdMillionsOk returns a tuple with the LatestRevenueUsdMillions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevenueUsdMillions

`func (o *DrugBrandsRow) SetLatestRevenueUsdMillions(v float64)`

SetLatestRevenueUsdMillions sets LatestRevenueUsdMillions field to given value.

### HasLatestRevenueUsdMillions

`func (o *DrugBrandsRow) HasLatestRevenueUsdMillions() bool`

HasLatestRevenueUsdMillions returns a boolean if a field has been set.

### SetLatestRevenueUsdMillionsNil

`func (o *DrugBrandsRow) SetLatestRevenueUsdMillionsNil(b bool)`

 SetLatestRevenueUsdMillionsNil sets the value for LatestRevenueUsdMillions to be an explicit nil

### UnsetLatestRevenueUsdMillions
`func (o *DrugBrandsRow) UnsetLatestRevenueUsdMillions()`

UnsetLatestRevenueUsdMillions ensures that no value is present for LatestRevenueUsdMillions, not even an explicit nil
### GetLatestRevenueYear

`func (o *DrugBrandsRow) GetLatestRevenueYear() int32`

GetLatestRevenueYear returns the LatestRevenueYear field if non-nil, zero value otherwise.

### GetLatestRevenueYearOk

`func (o *DrugBrandsRow) GetLatestRevenueYearOk() (*int32, bool)`

GetLatestRevenueYearOk returns a tuple with the LatestRevenueYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevenueYear

`func (o *DrugBrandsRow) SetLatestRevenueYear(v int32)`

SetLatestRevenueYear sets LatestRevenueYear field to given value.

### HasLatestRevenueYear

`func (o *DrugBrandsRow) HasLatestRevenueYear() bool`

HasLatestRevenueYear returns a boolean if a field has been set.

### SetLatestRevenueYearNil

`func (o *DrugBrandsRow) SetLatestRevenueYearNil(b bool)`

 SetLatestRevenueYearNil sets the value for LatestRevenueYear to be an explicit nil

### UnsetLatestRevenueYear
`func (o *DrugBrandsRow) UnsetLatestRevenueYear()`

UnsetLatestRevenueYear ensures that no value is present for LatestRevenueYear, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


