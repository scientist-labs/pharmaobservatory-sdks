# SuppliersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | Pointer to **NullableString** | The supplier id: the Scientist.com provider id as text, the id in the supplier page address (/suppliers/&lt;id&gt;). A supplier is not a company. | [optional] 
**Name** | Pointer to **NullableString** | The supplier&#39;s name as listed on the marketplace. | [optional] 
**CompanyType** | Pointer to **NullableString** | The supplier&#39;s company type as the marketplace spells it (small-business, academic-core and so on). An open set; null when the supplier lists none. | [optional] 
**MarketplaceTier** | Pointer to **NullableString** | The supplier&#39;s marketplace tier word; null when the marketplace assigns none. | [optional] 
**ServiceAreas** | Pointer to **[]string** | The service areas the supplier lists, one array element each (empty when none). Filter with &#39;Biology&#39; &#x3D; ANY(service_areas). | [optional] 
**HqCountry** | Pointer to **NullableString** | ISO 3166-1 alpha-2 country code of the supplier&#39;s headquarters point (null when unknown). | [optional] 
**HqRegion** | Pointer to **NullableString** | The broad world region of hq_country, from the country_regions reference table; other when the country is unknown or in none of the six regions. | [optional] 
**HqLatitude** | Pointer to **NullableFloat32** | Latitude of the supplier&#39;s headquarters point: the Headquarters site, else the first site by category then site id. | [optional] 
**HqLongitude** | Pointer to **NullableFloat32** | Longitude of the supplier&#39;s headquarters point. | [optional] 
**NSites** | Pointer to **NullableInt64** | How many located sites the supplier has. | [optional] 
**NCountries** | Pointer to **NullableInt64** | How many distinct countries the supplier&#39;s sites are in. | [optional] 
**IsCro** | Pointer to **NullableBool** | True when the supplier&#39;s organization is a contract research organization (CRO) under the one organization-level definition. | [optional] 
**IsTrialSponsor** | Pointer to **NullableBool** | True when the supplier name matches a registered trial sponsor. A name match, not a verified identity. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Registered trials where an organization matched to the supplier by name is lead sponsor or collaborator. | [optional] 
**NLeadTrials** | Pointer to **NullableInt64** | The subset of n_trials where the matched organization is the lead sponsor. | [optional] 
**Website** | Pointer to **NullableString** | The website the supplier listed; a third-party address, not vouched for. | [optional] 
**ProfileUrl** | Pointer to **NullableString** | The supplier&#39;s page on the Scientist.com marketplace; the marketplace asks a visitor to sign in. | [optional] 
**LogoPath** | Pointer to **NullableString** | The path of the supplier&#39;s logo on our own origin (/supplier-logos/&lt;sha&gt;.png); null when it has none. | [optional] 
**HasDescription** | Pointer to **NullableBool** | True when the supplier wrote a profile description. The description text is not published. | [optional] 
**IsVeteranOwned** | Pointer to **NullableBool** | SAM.gov self-certified veteran-owned designation. | [optional] 
**IsWomanOwned** | Pointer to **NullableBool** | SAM.gov self-certified woman-owned designation. | [optional] 
**IsMinorityOwned** | Pointer to **NullableBool** | SAM.gov self-certified minority-owned designation (8(a) or small disadvantaged business). | [optional] 
**IsHubzone** | Pointer to **NullableBool** | SAM.gov self-certified HUBZone (historically underutilized business zone) designation. | [optional] 
**SupplierOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the supplier resolved to. Null when the supplier has no organization record of its own. | [optional] 
**SupplierCompanyId** | Pointer to **NullableString** | Company level: the company the supplier&#39;s organization belongs to (null when none). | [optional] 
**SupplierCompanyName** | Pointer to **NullableString** | Display name of the supplier&#39;s company. | [optional] 
**SupplierGroupId** | Pointer to **NullableString** | Group level: the supplier&#39;s company&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SupplierGroupName** | Pointer to **NullableString** | Display name of the supplier&#39;s group. | [optional] 
**OrgLinkMethod** | Pointer to **NullableString** | How the supplier was linked to its organization record, so a caller can apply a confidence floor. The supplier pages on organization pages show only provider_ror_exact, provider_edgar_exact and supplier_curated. | [optional] 
**PhsAnimalWelfareAssuranceId** | Pointer to **NullableString** | The NIH Office of Laboratory Animal Welfare (OLAW) Public Health Service Animal Welfare Assurance id the supplier page badges (#5286). Set only when the supplier is linked to its organization with confidence (provider_ror_exact, provider_edgar_exact or supplier_curated) and the OLAW institution matched that organization at a high or medium tier, the same rule as the page. Null means no badge, never &#39;not assured&#39;: OLAW&#39;s list is matched to the organizations we hold, and most institutions do not resolve. Full detail, with the raw status code, is the animal_welfare_assurances dataset. | [optional] 
**HasPublicPage** | Pointer to **NullableBool** | True when the supplier has a public page on the web (every row of this dataset does). | [optional] 

## Methods

### NewSuppliersRow

`func NewSuppliersRow() *SuppliersRow`

NewSuppliersRow instantiates a new SuppliersRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppliersRowWithDefaults

`func NewSuppliersRowWithDefaults() *SuppliersRow`

NewSuppliersRowWithDefaults instantiates a new SuppliersRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *SuppliersRow) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *SuppliersRow) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *SuppliersRow) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *SuppliersRow) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *SuppliersRow) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *SuppliersRow) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetName

`func (o *SuppliersRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SuppliersRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SuppliersRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SuppliersRow) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SuppliersRow) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SuppliersRow) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCompanyType

`func (o *SuppliersRow) GetCompanyType() string`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *SuppliersRow) GetCompanyTypeOk() (*string, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *SuppliersRow) SetCompanyType(v string)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *SuppliersRow) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### SetCompanyTypeNil

`func (o *SuppliersRow) SetCompanyTypeNil(b bool)`

 SetCompanyTypeNil sets the value for CompanyType to be an explicit nil

### UnsetCompanyType
`func (o *SuppliersRow) UnsetCompanyType()`

UnsetCompanyType ensures that no value is present for CompanyType, not even an explicit nil
### GetMarketplaceTier

`func (o *SuppliersRow) GetMarketplaceTier() string`

GetMarketplaceTier returns the MarketplaceTier field if non-nil, zero value otherwise.

### GetMarketplaceTierOk

`func (o *SuppliersRow) GetMarketplaceTierOk() (*string, bool)`

GetMarketplaceTierOk returns a tuple with the MarketplaceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceTier

`func (o *SuppliersRow) SetMarketplaceTier(v string)`

SetMarketplaceTier sets MarketplaceTier field to given value.

### HasMarketplaceTier

`func (o *SuppliersRow) HasMarketplaceTier() bool`

HasMarketplaceTier returns a boolean if a field has been set.

### SetMarketplaceTierNil

`func (o *SuppliersRow) SetMarketplaceTierNil(b bool)`

 SetMarketplaceTierNil sets the value for MarketplaceTier to be an explicit nil

### UnsetMarketplaceTier
`func (o *SuppliersRow) UnsetMarketplaceTier()`

UnsetMarketplaceTier ensures that no value is present for MarketplaceTier, not even an explicit nil
### GetServiceAreas

`func (o *SuppliersRow) GetServiceAreas() []string`

GetServiceAreas returns the ServiceAreas field if non-nil, zero value otherwise.

### GetServiceAreasOk

`func (o *SuppliersRow) GetServiceAreasOk() (*[]string, bool)`

GetServiceAreasOk returns a tuple with the ServiceAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAreas

`func (o *SuppliersRow) SetServiceAreas(v []string)`

SetServiceAreas sets ServiceAreas field to given value.

### HasServiceAreas

`func (o *SuppliersRow) HasServiceAreas() bool`

HasServiceAreas returns a boolean if a field has been set.

### SetServiceAreasNil

`func (o *SuppliersRow) SetServiceAreasNil(b bool)`

 SetServiceAreasNil sets the value for ServiceAreas to be an explicit nil

### UnsetServiceAreas
`func (o *SuppliersRow) UnsetServiceAreas()`

UnsetServiceAreas ensures that no value is present for ServiceAreas, not even an explicit nil
### GetHqCountry

`func (o *SuppliersRow) GetHqCountry() string`

GetHqCountry returns the HqCountry field if non-nil, zero value otherwise.

### GetHqCountryOk

`func (o *SuppliersRow) GetHqCountryOk() (*string, bool)`

GetHqCountryOk returns a tuple with the HqCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqCountry

`func (o *SuppliersRow) SetHqCountry(v string)`

SetHqCountry sets HqCountry field to given value.

### HasHqCountry

`func (o *SuppliersRow) HasHqCountry() bool`

HasHqCountry returns a boolean if a field has been set.

### SetHqCountryNil

`func (o *SuppliersRow) SetHqCountryNil(b bool)`

 SetHqCountryNil sets the value for HqCountry to be an explicit nil

### UnsetHqCountry
`func (o *SuppliersRow) UnsetHqCountry()`

UnsetHqCountry ensures that no value is present for HqCountry, not even an explicit nil
### GetHqRegion

`func (o *SuppliersRow) GetHqRegion() string`

GetHqRegion returns the HqRegion field if non-nil, zero value otherwise.

### GetHqRegionOk

`func (o *SuppliersRow) GetHqRegionOk() (*string, bool)`

GetHqRegionOk returns a tuple with the HqRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqRegion

`func (o *SuppliersRow) SetHqRegion(v string)`

SetHqRegion sets HqRegion field to given value.

### HasHqRegion

`func (o *SuppliersRow) HasHqRegion() bool`

HasHqRegion returns a boolean if a field has been set.

### SetHqRegionNil

`func (o *SuppliersRow) SetHqRegionNil(b bool)`

 SetHqRegionNil sets the value for HqRegion to be an explicit nil

### UnsetHqRegion
`func (o *SuppliersRow) UnsetHqRegion()`

UnsetHqRegion ensures that no value is present for HqRegion, not even an explicit nil
### GetHqLatitude

`func (o *SuppliersRow) GetHqLatitude() float32`

GetHqLatitude returns the HqLatitude field if non-nil, zero value otherwise.

### GetHqLatitudeOk

`func (o *SuppliersRow) GetHqLatitudeOk() (*float32, bool)`

GetHqLatitudeOk returns a tuple with the HqLatitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqLatitude

`func (o *SuppliersRow) SetHqLatitude(v float32)`

SetHqLatitude sets HqLatitude field to given value.

### HasHqLatitude

`func (o *SuppliersRow) HasHqLatitude() bool`

HasHqLatitude returns a boolean if a field has been set.

### SetHqLatitudeNil

`func (o *SuppliersRow) SetHqLatitudeNil(b bool)`

 SetHqLatitudeNil sets the value for HqLatitude to be an explicit nil

### UnsetHqLatitude
`func (o *SuppliersRow) UnsetHqLatitude()`

UnsetHqLatitude ensures that no value is present for HqLatitude, not even an explicit nil
### GetHqLongitude

`func (o *SuppliersRow) GetHqLongitude() float32`

GetHqLongitude returns the HqLongitude field if non-nil, zero value otherwise.

### GetHqLongitudeOk

`func (o *SuppliersRow) GetHqLongitudeOk() (*float32, bool)`

GetHqLongitudeOk returns a tuple with the HqLongitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqLongitude

`func (o *SuppliersRow) SetHqLongitude(v float32)`

SetHqLongitude sets HqLongitude field to given value.

### HasHqLongitude

`func (o *SuppliersRow) HasHqLongitude() bool`

HasHqLongitude returns a boolean if a field has been set.

### SetHqLongitudeNil

`func (o *SuppliersRow) SetHqLongitudeNil(b bool)`

 SetHqLongitudeNil sets the value for HqLongitude to be an explicit nil

### UnsetHqLongitude
`func (o *SuppliersRow) UnsetHqLongitude()`

UnsetHqLongitude ensures that no value is present for HqLongitude, not even an explicit nil
### GetNSites

`func (o *SuppliersRow) GetNSites() int64`

GetNSites returns the NSites field if non-nil, zero value otherwise.

### GetNSitesOk

`func (o *SuppliersRow) GetNSitesOk() (*int64, bool)`

GetNSitesOk returns a tuple with the NSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSites

`func (o *SuppliersRow) SetNSites(v int64)`

SetNSites sets NSites field to given value.

### HasNSites

`func (o *SuppliersRow) HasNSites() bool`

HasNSites returns a boolean if a field has been set.

### SetNSitesNil

`func (o *SuppliersRow) SetNSitesNil(b bool)`

 SetNSitesNil sets the value for NSites to be an explicit nil

### UnsetNSites
`func (o *SuppliersRow) UnsetNSites()`

UnsetNSites ensures that no value is present for NSites, not even an explicit nil
### GetNCountries

`func (o *SuppliersRow) GetNCountries() int64`

GetNCountries returns the NCountries field if non-nil, zero value otherwise.

### GetNCountriesOk

`func (o *SuppliersRow) GetNCountriesOk() (*int64, bool)`

GetNCountriesOk returns a tuple with the NCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCountries

`func (o *SuppliersRow) SetNCountries(v int64)`

SetNCountries sets NCountries field to given value.

### HasNCountries

`func (o *SuppliersRow) HasNCountries() bool`

HasNCountries returns a boolean if a field has been set.

### SetNCountriesNil

`func (o *SuppliersRow) SetNCountriesNil(b bool)`

 SetNCountriesNil sets the value for NCountries to be an explicit nil

### UnsetNCountries
`func (o *SuppliersRow) UnsetNCountries()`

UnsetNCountries ensures that no value is present for NCountries, not even an explicit nil
### GetIsCro

`func (o *SuppliersRow) GetIsCro() bool`

GetIsCro returns the IsCro field if non-nil, zero value otherwise.

### GetIsCroOk

`func (o *SuppliersRow) GetIsCroOk() (*bool, bool)`

GetIsCroOk returns a tuple with the IsCro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCro

`func (o *SuppliersRow) SetIsCro(v bool)`

SetIsCro sets IsCro field to given value.

### HasIsCro

`func (o *SuppliersRow) HasIsCro() bool`

HasIsCro returns a boolean if a field has been set.

### SetIsCroNil

`func (o *SuppliersRow) SetIsCroNil(b bool)`

 SetIsCroNil sets the value for IsCro to be an explicit nil

### UnsetIsCro
`func (o *SuppliersRow) UnsetIsCro()`

UnsetIsCro ensures that no value is present for IsCro, not even an explicit nil
### GetIsTrialSponsor

`func (o *SuppliersRow) GetIsTrialSponsor() bool`

GetIsTrialSponsor returns the IsTrialSponsor field if non-nil, zero value otherwise.

### GetIsTrialSponsorOk

`func (o *SuppliersRow) GetIsTrialSponsorOk() (*bool, bool)`

GetIsTrialSponsorOk returns a tuple with the IsTrialSponsor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrialSponsor

`func (o *SuppliersRow) SetIsTrialSponsor(v bool)`

SetIsTrialSponsor sets IsTrialSponsor field to given value.

### HasIsTrialSponsor

`func (o *SuppliersRow) HasIsTrialSponsor() bool`

HasIsTrialSponsor returns a boolean if a field has been set.

### SetIsTrialSponsorNil

`func (o *SuppliersRow) SetIsTrialSponsorNil(b bool)`

 SetIsTrialSponsorNil sets the value for IsTrialSponsor to be an explicit nil

### UnsetIsTrialSponsor
`func (o *SuppliersRow) UnsetIsTrialSponsor()`

UnsetIsTrialSponsor ensures that no value is present for IsTrialSponsor, not even an explicit nil
### GetNTrials

`func (o *SuppliersRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *SuppliersRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *SuppliersRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *SuppliersRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *SuppliersRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *SuppliersRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNLeadTrials

`func (o *SuppliersRow) GetNLeadTrials() int64`

GetNLeadTrials returns the NLeadTrials field if non-nil, zero value otherwise.

### GetNLeadTrialsOk

`func (o *SuppliersRow) GetNLeadTrialsOk() (*int64, bool)`

GetNLeadTrialsOk returns a tuple with the NLeadTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNLeadTrials

`func (o *SuppliersRow) SetNLeadTrials(v int64)`

SetNLeadTrials sets NLeadTrials field to given value.

### HasNLeadTrials

`func (o *SuppliersRow) HasNLeadTrials() bool`

HasNLeadTrials returns a boolean if a field has been set.

### SetNLeadTrialsNil

`func (o *SuppliersRow) SetNLeadTrialsNil(b bool)`

 SetNLeadTrialsNil sets the value for NLeadTrials to be an explicit nil

### UnsetNLeadTrials
`func (o *SuppliersRow) UnsetNLeadTrials()`

UnsetNLeadTrials ensures that no value is present for NLeadTrials, not even an explicit nil
### GetWebsite

`func (o *SuppliersRow) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *SuppliersRow) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *SuppliersRow) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *SuppliersRow) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *SuppliersRow) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *SuppliersRow) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetProfileUrl

`func (o *SuppliersRow) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *SuppliersRow) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *SuppliersRow) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *SuppliersRow) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### SetProfileUrlNil

`func (o *SuppliersRow) SetProfileUrlNil(b bool)`

 SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil

### UnsetProfileUrl
`func (o *SuppliersRow) UnsetProfileUrl()`

UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
### GetLogoPath

`func (o *SuppliersRow) GetLogoPath() string`

GetLogoPath returns the LogoPath field if non-nil, zero value otherwise.

### GetLogoPathOk

`func (o *SuppliersRow) GetLogoPathOk() (*string, bool)`

GetLogoPathOk returns a tuple with the LogoPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoPath

`func (o *SuppliersRow) SetLogoPath(v string)`

SetLogoPath sets LogoPath field to given value.

### HasLogoPath

`func (o *SuppliersRow) HasLogoPath() bool`

HasLogoPath returns a boolean if a field has been set.

### SetLogoPathNil

`func (o *SuppliersRow) SetLogoPathNil(b bool)`

 SetLogoPathNil sets the value for LogoPath to be an explicit nil

### UnsetLogoPath
`func (o *SuppliersRow) UnsetLogoPath()`

UnsetLogoPath ensures that no value is present for LogoPath, not even an explicit nil
### GetHasDescription

`func (o *SuppliersRow) GetHasDescription() bool`

GetHasDescription returns the HasDescription field if non-nil, zero value otherwise.

### GetHasDescriptionOk

`func (o *SuppliersRow) GetHasDescriptionOk() (*bool, bool)`

GetHasDescriptionOk returns a tuple with the HasDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDescription

`func (o *SuppliersRow) SetHasDescription(v bool)`

SetHasDescription sets HasDescription field to given value.

### HasHasDescription

`func (o *SuppliersRow) HasHasDescription() bool`

HasHasDescription returns a boolean if a field has been set.

### SetHasDescriptionNil

`func (o *SuppliersRow) SetHasDescriptionNil(b bool)`

 SetHasDescriptionNil sets the value for HasDescription to be an explicit nil

### UnsetHasDescription
`func (o *SuppliersRow) UnsetHasDescription()`

UnsetHasDescription ensures that no value is present for HasDescription, not even an explicit nil
### GetIsVeteranOwned

`func (o *SuppliersRow) GetIsVeteranOwned() bool`

GetIsVeteranOwned returns the IsVeteranOwned field if non-nil, zero value otherwise.

### GetIsVeteranOwnedOk

`func (o *SuppliersRow) GetIsVeteranOwnedOk() (*bool, bool)`

GetIsVeteranOwnedOk returns a tuple with the IsVeteranOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVeteranOwned

`func (o *SuppliersRow) SetIsVeteranOwned(v bool)`

SetIsVeteranOwned sets IsVeteranOwned field to given value.

### HasIsVeteranOwned

`func (o *SuppliersRow) HasIsVeteranOwned() bool`

HasIsVeteranOwned returns a boolean if a field has been set.

### SetIsVeteranOwnedNil

`func (o *SuppliersRow) SetIsVeteranOwnedNil(b bool)`

 SetIsVeteranOwnedNil sets the value for IsVeteranOwned to be an explicit nil

### UnsetIsVeteranOwned
`func (o *SuppliersRow) UnsetIsVeteranOwned()`

UnsetIsVeteranOwned ensures that no value is present for IsVeteranOwned, not even an explicit nil
### GetIsWomanOwned

`func (o *SuppliersRow) GetIsWomanOwned() bool`

GetIsWomanOwned returns the IsWomanOwned field if non-nil, zero value otherwise.

### GetIsWomanOwnedOk

`func (o *SuppliersRow) GetIsWomanOwnedOk() (*bool, bool)`

GetIsWomanOwnedOk returns a tuple with the IsWomanOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWomanOwned

`func (o *SuppliersRow) SetIsWomanOwned(v bool)`

SetIsWomanOwned sets IsWomanOwned field to given value.

### HasIsWomanOwned

`func (o *SuppliersRow) HasIsWomanOwned() bool`

HasIsWomanOwned returns a boolean if a field has been set.

### SetIsWomanOwnedNil

`func (o *SuppliersRow) SetIsWomanOwnedNil(b bool)`

 SetIsWomanOwnedNil sets the value for IsWomanOwned to be an explicit nil

### UnsetIsWomanOwned
`func (o *SuppliersRow) UnsetIsWomanOwned()`

UnsetIsWomanOwned ensures that no value is present for IsWomanOwned, not even an explicit nil
### GetIsMinorityOwned

`func (o *SuppliersRow) GetIsMinorityOwned() bool`

GetIsMinorityOwned returns the IsMinorityOwned field if non-nil, zero value otherwise.

### GetIsMinorityOwnedOk

`func (o *SuppliersRow) GetIsMinorityOwnedOk() (*bool, bool)`

GetIsMinorityOwnedOk returns a tuple with the IsMinorityOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMinorityOwned

`func (o *SuppliersRow) SetIsMinorityOwned(v bool)`

SetIsMinorityOwned sets IsMinorityOwned field to given value.

### HasIsMinorityOwned

`func (o *SuppliersRow) HasIsMinorityOwned() bool`

HasIsMinorityOwned returns a boolean if a field has been set.

### SetIsMinorityOwnedNil

`func (o *SuppliersRow) SetIsMinorityOwnedNil(b bool)`

 SetIsMinorityOwnedNil sets the value for IsMinorityOwned to be an explicit nil

### UnsetIsMinorityOwned
`func (o *SuppliersRow) UnsetIsMinorityOwned()`

UnsetIsMinorityOwned ensures that no value is present for IsMinorityOwned, not even an explicit nil
### GetIsHubzone

`func (o *SuppliersRow) GetIsHubzone() bool`

GetIsHubzone returns the IsHubzone field if non-nil, zero value otherwise.

### GetIsHubzoneOk

`func (o *SuppliersRow) GetIsHubzoneOk() (*bool, bool)`

GetIsHubzoneOk returns a tuple with the IsHubzone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHubzone

`func (o *SuppliersRow) SetIsHubzone(v bool)`

SetIsHubzone sets IsHubzone field to given value.

### HasIsHubzone

`func (o *SuppliersRow) HasIsHubzone() bool`

HasIsHubzone returns a boolean if a field has been set.

### SetIsHubzoneNil

`func (o *SuppliersRow) SetIsHubzoneNil(b bool)`

 SetIsHubzoneNil sets the value for IsHubzone to be an explicit nil

### UnsetIsHubzone
`func (o *SuppliersRow) UnsetIsHubzone()`

UnsetIsHubzone ensures that no value is present for IsHubzone, not even an explicit nil
### GetSupplierOrgId

`func (o *SuppliersRow) GetSupplierOrgId() string`

GetSupplierOrgId returns the SupplierOrgId field if non-nil, zero value otherwise.

### GetSupplierOrgIdOk

`func (o *SuppliersRow) GetSupplierOrgIdOk() (*string, bool)`

GetSupplierOrgIdOk returns a tuple with the SupplierOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierOrgId

`func (o *SuppliersRow) SetSupplierOrgId(v string)`

SetSupplierOrgId sets SupplierOrgId field to given value.

### HasSupplierOrgId

`func (o *SuppliersRow) HasSupplierOrgId() bool`

HasSupplierOrgId returns a boolean if a field has been set.

### SetSupplierOrgIdNil

`func (o *SuppliersRow) SetSupplierOrgIdNil(b bool)`

 SetSupplierOrgIdNil sets the value for SupplierOrgId to be an explicit nil

### UnsetSupplierOrgId
`func (o *SuppliersRow) UnsetSupplierOrgId()`

UnsetSupplierOrgId ensures that no value is present for SupplierOrgId, not even an explicit nil
### GetSupplierCompanyId

`func (o *SuppliersRow) GetSupplierCompanyId() string`

GetSupplierCompanyId returns the SupplierCompanyId field if non-nil, zero value otherwise.

### GetSupplierCompanyIdOk

`func (o *SuppliersRow) GetSupplierCompanyIdOk() (*string, bool)`

GetSupplierCompanyIdOk returns a tuple with the SupplierCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierCompanyId

`func (o *SuppliersRow) SetSupplierCompanyId(v string)`

SetSupplierCompanyId sets SupplierCompanyId field to given value.

### HasSupplierCompanyId

`func (o *SuppliersRow) HasSupplierCompanyId() bool`

HasSupplierCompanyId returns a boolean if a field has been set.

### SetSupplierCompanyIdNil

`func (o *SuppliersRow) SetSupplierCompanyIdNil(b bool)`

 SetSupplierCompanyIdNil sets the value for SupplierCompanyId to be an explicit nil

### UnsetSupplierCompanyId
`func (o *SuppliersRow) UnsetSupplierCompanyId()`

UnsetSupplierCompanyId ensures that no value is present for SupplierCompanyId, not even an explicit nil
### GetSupplierCompanyName

`func (o *SuppliersRow) GetSupplierCompanyName() string`

GetSupplierCompanyName returns the SupplierCompanyName field if non-nil, zero value otherwise.

### GetSupplierCompanyNameOk

`func (o *SuppliersRow) GetSupplierCompanyNameOk() (*string, bool)`

GetSupplierCompanyNameOk returns a tuple with the SupplierCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierCompanyName

`func (o *SuppliersRow) SetSupplierCompanyName(v string)`

SetSupplierCompanyName sets SupplierCompanyName field to given value.

### HasSupplierCompanyName

`func (o *SuppliersRow) HasSupplierCompanyName() bool`

HasSupplierCompanyName returns a boolean if a field has been set.

### SetSupplierCompanyNameNil

`func (o *SuppliersRow) SetSupplierCompanyNameNil(b bool)`

 SetSupplierCompanyNameNil sets the value for SupplierCompanyName to be an explicit nil

### UnsetSupplierCompanyName
`func (o *SuppliersRow) UnsetSupplierCompanyName()`

UnsetSupplierCompanyName ensures that no value is present for SupplierCompanyName, not even an explicit nil
### GetSupplierGroupId

`func (o *SuppliersRow) GetSupplierGroupId() string`

GetSupplierGroupId returns the SupplierGroupId field if non-nil, zero value otherwise.

### GetSupplierGroupIdOk

`func (o *SuppliersRow) GetSupplierGroupIdOk() (*string, bool)`

GetSupplierGroupIdOk returns a tuple with the SupplierGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierGroupId

`func (o *SuppliersRow) SetSupplierGroupId(v string)`

SetSupplierGroupId sets SupplierGroupId field to given value.

### HasSupplierGroupId

`func (o *SuppliersRow) HasSupplierGroupId() bool`

HasSupplierGroupId returns a boolean if a field has been set.

### SetSupplierGroupIdNil

`func (o *SuppliersRow) SetSupplierGroupIdNil(b bool)`

 SetSupplierGroupIdNil sets the value for SupplierGroupId to be an explicit nil

### UnsetSupplierGroupId
`func (o *SuppliersRow) UnsetSupplierGroupId()`

UnsetSupplierGroupId ensures that no value is present for SupplierGroupId, not even an explicit nil
### GetSupplierGroupName

`func (o *SuppliersRow) GetSupplierGroupName() string`

GetSupplierGroupName returns the SupplierGroupName field if non-nil, zero value otherwise.

### GetSupplierGroupNameOk

`func (o *SuppliersRow) GetSupplierGroupNameOk() (*string, bool)`

GetSupplierGroupNameOk returns a tuple with the SupplierGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierGroupName

`func (o *SuppliersRow) SetSupplierGroupName(v string)`

SetSupplierGroupName sets SupplierGroupName field to given value.

### HasSupplierGroupName

`func (o *SuppliersRow) HasSupplierGroupName() bool`

HasSupplierGroupName returns a boolean if a field has been set.

### SetSupplierGroupNameNil

`func (o *SuppliersRow) SetSupplierGroupNameNil(b bool)`

 SetSupplierGroupNameNil sets the value for SupplierGroupName to be an explicit nil

### UnsetSupplierGroupName
`func (o *SuppliersRow) UnsetSupplierGroupName()`

UnsetSupplierGroupName ensures that no value is present for SupplierGroupName, not even an explicit nil
### GetOrgLinkMethod

`func (o *SuppliersRow) GetOrgLinkMethod() string`

GetOrgLinkMethod returns the OrgLinkMethod field if non-nil, zero value otherwise.

### GetOrgLinkMethodOk

`func (o *SuppliersRow) GetOrgLinkMethodOk() (*string, bool)`

GetOrgLinkMethodOk returns a tuple with the OrgLinkMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgLinkMethod

`func (o *SuppliersRow) SetOrgLinkMethod(v string)`

SetOrgLinkMethod sets OrgLinkMethod field to given value.

### HasOrgLinkMethod

`func (o *SuppliersRow) HasOrgLinkMethod() bool`

HasOrgLinkMethod returns a boolean if a field has been set.

### SetOrgLinkMethodNil

`func (o *SuppliersRow) SetOrgLinkMethodNil(b bool)`

 SetOrgLinkMethodNil sets the value for OrgLinkMethod to be an explicit nil

### UnsetOrgLinkMethod
`func (o *SuppliersRow) UnsetOrgLinkMethod()`

UnsetOrgLinkMethod ensures that no value is present for OrgLinkMethod, not even an explicit nil
### GetPhsAnimalWelfareAssuranceId

`func (o *SuppliersRow) GetPhsAnimalWelfareAssuranceId() string`

GetPhsAnimalWelfareAssuranceId returns the PhsAnimalWelfareAssuranceId field if non-nil, zero value otherwise.

### GetPhsAnimalWelfareAssuranceIdOk

`func (o *SuppliersRow) GetPhsAnimalWelfareAssuranceIdOk() (*string, bool)`

GetPhsAnimalWelfareAssuranceIdOk returns a tuple with the PhsAnimalWelfareAssuranceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhsAnimalWelfareAssuranceId

`func (o *SuppliersRow) SetPhsAnimalWelfareAssuranceId(v string)`

SetPhsAnimalWelfareAssuranceId sets PhsAnimalWelfareAssuranceId field to given value.

### HasPhsAnimalWelfareAssuranceId

`func (o *SuppliersRow) HasPhsAnimalWelfareAssuranceId() bool`

HasPhsAnimalWelfareAssuranceId returns a boolean if a field has been set.

### SetPhsAnimalWelfareAssuranceIdNil

`func (o *SuppliersRow) SetPhsAnimalWelfareAssuranceIdNil(b bool)`

 SetPhsAnimalWelfareAssuranceIdNil sets the value for PhsAnimalWelfareAssuranceId to be an explicit nil

### UnsetPhsAnimalWelfareAssuranceId
`func (o *SuppliersRow) UnsetPhsAnimalWelfareAssuranceId()`

UnsetPhsAnimalWelfareAssuranceId ensures that no value is present for PhsAnimalWelfareAssuranceId, not even an explicit nil
### GetHasPublicPage

`func (o *SuppliersRow) GetHasPublicPage() bool`

GetHasPublicPage returns the HasPublicPage field if non-nil, zero value otherwise.

### GetHasPublicPageOk

`func (o *SuppliersRow) GetHasPublicPageOk() (*bool, bool)`

GetHasPublicPageOk returns a tuple with the HasPublicPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPublicPage

`func (o *SuppliersRow) SetHasPublicPage(v bool)`

SetHasPublicPage sets HasPublicPage field to given value.

### HasHasPublicPage

`func (o *SuppliersRow) HasHasPublicPage() bool`

HasHasPublicPage returns a boolean if a field has been set.

### SetHasPublicPageNil

`func (o *SuppliersRow) SetHasPublicPageNil(b bool)`

 SetHasPublicPageNil sets the value for HasPublicPage to be an explicit nil

### UnsetHasPublicPage
`func (o *SuppliersRow) UnsetHasPublicPage()`

UnsetHasPublicPage ensures that no value is present for HasPublicPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


