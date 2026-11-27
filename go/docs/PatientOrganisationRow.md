# PatientOrganisationRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetKind** | Pointer to **NullableString** | Which kind of page the row is listed on: rare_disease (a rare disease page, entered on orpha_code) or condition (a condition page, entered on condition_id). | [optional] 
**TargetKey** | Pointer to **NullableString** | The orpha_code of a rare_disease row or the condition_id of a condition row: the key the row is entered on. Always present, so it can stand in a sort or a cursor. | [optional] 
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease the row is listed under, the digits in the web page address (/rare-diseases/558); null on a condition row. The entry key of the rare_diseases object type. | [optional] 
**ConditionId** | Pointer to **NullableString** | The MeSH descriptor id of the condition the row is listed under (D followed by six to nine digits); null on a rare disease row. The entry key of the conditions object type. | [optional] 
**PatientOrganisationId** | Pointer to **NullableString** | The organisation&#39;s id: a spine organisation id (edgar:, ror:, sponsor: and the like) when the organisation folded onto an organisation we already hold, else pag:&lt;registrable domain of its website&gt;. | [optional] 
**OrgId** | Pointer to **NullableString** | The legal-entity id of the organisation when it folded onto one we already hold, the join to legal_entities; null for a pag: organisation, which has no spine record. | [optional] 
**Name** | Pointer to **NullableString** | The organisation&#39;s name as its own website states it. | [optional] 
**WebsiteUrl** | Pointer to **NullableString** | The organisation&#39;s website as it states it. | [optional] 
**Country** | Pointer to **NullableString** | The country the organisation declares (ISO 3166-1 alpha-2), else its primary location&#39;s country. | [optional] 
**Region** | Pointer to **NullableString** | The region of its primary location. | [optional] 
**City** | Pointer to **NullableString** | The city of its primary location. | [optional] 
**LocationCount** | Pointer to **NullableInt32** | How many locations the organisation lists. | [optional] 
**Languages** | Pointer to **[]string** | The languages the organisation says it works in, as it states them. | [optional] 
**Phones** | Pointer to **[]string** | The organisation&#39;s own phone numbers, as stated. Organisation-level channels only: no person&#39;s line. | [optional] 
**Emails** | Pointer to **[]string** | The organisation&#39;s own email addresses, as stated, minus any address whose local part looks like a person&#39;s name. Organisation-level channels only. | [optional] 
**ContactUrls** | Pointer to **[]string** | The organisation&#39;s own contact page addresses. | [optional] 
**SocialUrls** | Pointer to **[]string** | The organisation&#39;s own social media addresses. | [optional] 
**ServiceCodes** | Pointer to **[]string** | The services the organisation&#39;s own pages say it offers, from a closed list: support_groups, helpline, information_and_education, patient_registry, research_funding, advocacy_and_policy, financial_assistance, care_navigation, clinical_trial_matching, other. | [optional] 
**OtherServiceLabels** | Pointer to **[]string** | For a service coded other, the organisation&#39;s own words for it. Empty when it states none. | [optional] 
**LinkKind** | Pointer to **NullableString** | How the organisation reaches the disease or condition the row is listed under: declared (its own page names it), mapped (a condition row: it declares a rare disease this condition is mapped to; the mapped_* columns say which), broader (a rare disease row: it declared a broader parent disease) or narrower (a rare disease row: it declared a narrower child disease). A broader or narrower row never says the organisation declares THIS disease. | [optional] 
**LinkRank** | Pointer to **NullableInt32** | The page&#39;s group order: 1 declared, 2 mapped, 3 broader, 4 narrower. | [optional] 
**DeclaredScope** | Pointer to **NullableString** | How much of its work the organisation says the disease is: sole_focus, one_of_several or umbrella; null when it states none. | [optional] 
**EdgeVerdict** | Pointer to **NullableString** | Whether the organisation&#39;s pages were judged to SERVE people with the disease it declared: serves, unclear (asked and not callable) or unjudged (not yet asked). Only unclear and unjudged rows carry a &#39;not yet confirmed&#39; label on the page; an edge judged to be a mention only is not a row. A judgement of ours about the organisation&#39;s pages, not the organisation&#39;s own statement. On a broader, narrower or mapped row it describes the declared disease. | [optional] 
**DeclaredName** | Pointer to **NullableString** | The disease name the organisation actually declared. On a broader, narrower or mapped row it is NOT this page&#39;s disease. | [optional] 
**MappedOrphaCode** | Pointer to **NullableString** | On a mapped row, the ORPHA code of the declared rare disease the condition is mapped to; null otherwise. | [optional] 
**MappedDiseaseName** | Pointer to **NullableString** | On a mapped row, the preferred name of that rare disease; null otherwise. | [optional] 
**MappingRelation** | Pointer to **NullableString** | On a mapped row, how the rare disease relates to the condition (exact, broader or narrower); null otherwise. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The page of the organisation&#39;s own website the declaration came from. | [optional] 
**RetrievedOn** | Pointer to **NullableString** | The date that page was retrieved. | [optional] 
**Credit** | Pointer to **NullableString** | The source credit line for the row: Source: &lt;organisation name&gt;, &lt;page address&gt;, retrieved &lt;date&gt;. Mandatory wherever the row is shown. | [optional] 

## Methods

### NewPatientOrganisationRow

`func NewPatientOrganisationRow() *PatientOrganisationRow`

NewPatientOrganisationRow instantiates a new PatientOrganisationRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatientOrganisationRowWithDefaults

`func NewPatientOrganisationRowWithDefaults() *PatientOrganisationRow`

NewPatientOrganisationRowWithDefaults instantiates a new PatientOrganisationRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetKind

`func (o *PatientOrganisationRow) GetTargetKind() string`

GetTargetKind returns the TargetKind field if non-nil, zero value otherwise.

### GetTargetKindOk

`func (o *PatientOrganisationRow) GetTargetKindOk() (*string, bool)`

GetTargetKindOk returns a tuple with the TargetKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKind

`func (o *PatientOrganisationRow) SetTargetKind(v string)`

SetTargetKind sets TargetKind field to given value.

### HasTargetKind

`func (o *PatientOrganisationRow) HasTargetKind() bool`

HasTargetKind returns a boolean if a field has been set.

### SetTargetKindNil

`func (o *PatientOrganisationRow) SetTargetKindNil(b bool)`

 SetTargetKindNil sets the value for TargetKind to be an explicit nil

### UnsetTargetKind
`func (o *PatientOrganisationRow) UnsetTargetKind()`

UnsetTargetKind ensures that no value is present for TargetKind, not even an explicit nil
### GetTargetKey

`func (o *PatientOrganisationRow) GetTargetKey() string`

GetTargetKey returns the TargetKey field if non-nil, zero value otherwise.

### GetTargetKeyOk

`func (o *PatientOrganisationRow) GetTargetKeyOk() (*string, bool)`

GetTargetKeyOk returns a tuple with the TargetKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKey

`func (o *PatientOrganisationRow) SetTargetKey(v string)`

SetTargetKey sets TargetKey field to given value.

### HasTargetKey

`func (o *PatientOrganisationRow) HasTargetKey() bool`

HasTargetKey returns a boolean if a field has been set.

### SetTargetKeyNil

`func (o *PatientOrganisationRow) SetTargetKeyNil(b bool)`

 SetTargetKeyNil sets the value for TargetKey to be an explicit nil

### UnsetTargetKey
`func (o *PatientOrganisationRow) UnsetTargetKey()`

UnsetTargetKey ensures that no value is present for TargetKey, not even an explicit nil
### GetOrphaCode

`func (o *PatientOrganisationRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *PatientOrganisationRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *PatientOrganisationRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *PatientOrganisationRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *PatientOrganisationRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *PatientOrganisationRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetConditionId

`func (o *PatientOrganisationRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *PatientOrganisationRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *PatientOrganisationRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *PatientOrganisationRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *PatientOrganisationRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *PatientOrganisationRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetPatientOrganisationId

`func (o *PatientOrganisationRow) GetPatientOrganisationId() string`

GetPatientOrganisationId returns the PatientOrganisationId field if non-nil, zero value otherwise.

### GetPatientOrganisationIdOk

`func (o *PatientOrganisationRow) GetPatientOrganisationIdOk() (*string, bool)`

GetPatientOrganisationIdOk returns a tuple with the PatientOrganisationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatientOrganisationId

`func (o *PatientOrganisationRow) SetPatientOrganisationId(v string)`

SetPatientOrganisationId sets PatientOrganisationId field to given value.

### HasPatientOrganisationId

`func (o *PatientOrganisationRow) HasPatientOrganisationId() bool`

HasPatientOrganisationId returns a boolean if a field has been set.

### SetPatientOrganisationIdNil

`func (o *PatientOrganisationRow) SetPatientOrganisationIdNil(b bool)`

 SetPatientOrganisationIdNil sets the value for PatientOrganisationId to be an explicit nil

### UnsetPatientOrganisationId
`func (o *PatientOrganisationRow) UnsetPatientOrganisationId()`

UnsetPatientOrganisationId ensures that no value is present for PatientOrganisationId, not even an explicit nil
### GetOrgId

`func (o *PatientOrganisationRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PatientOrganisationRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PatientOrganisationRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PatientOrganisationRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *PatientOrganisationRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *PatientOrganisationRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetName

`func (o *PatientOrganisationRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatientOrganisationRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatientOrganisationRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatientOrganisationRow) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatientOrganisationRow) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatientOrganisationRow) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetWebsiteUrl

`func (o *PatientOrganisationRow) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *PatientOrganisationRow) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *PatientOrganisationRow) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *PatientOrganisationRow) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *PatientOrganisationRow) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *PatientOrganisationRow) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetCountry

`func (o *PatientOrganisationRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PatientOrganisationRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PatientOrganisationRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PatientOrganisationRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PatientOrganisationRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PatientOrganisationRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetRegion

`func (o *PatientOrganisationRow) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PatientOrganisationRow) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PatientOrganisationRow) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PatientOrganisationRow) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *PatientOrganisationRow) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *PatientOrganisationRow) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetCity

`func (o *PatientOrganisationRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PatientOrganisationRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PatientOrganisationRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PatientOrganisationRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *PatientOrganisationRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *PatientOrganisationRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetLocationCount

`func (o *PatientOrganisationRow) GetLocationCount() int32`

GetLocationCount returns the LocationCount field if non-nil, zero value otherwise.

### GetLocationCountOk

`func (o *PatientOrganisationRow) GetLocationCountOk() (*int32, bool)`

GetLocationCountOk returns a tuple with the LocationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationCount

`func (o *PatientOrganisationRow) SetLocationCount(v int32)`

SetLocationCount sets LocationCount field to given value.

### HasLocationCount

`func (o *PatientOrganisationRow) HasLocationCount() bool`

HasLocationCount returns a boolean if a field has been set.

### SetLocationCountNil

`func (o *PatientOrganisationRow) SetLocationCountNil(b bool)`

 SetLocationCountNil sets the value for LocationCount to be an explicit nil

### UnsetLocationCount
`func (o *PatientOrganisationRow) UnsetLocationCount()`

UnsetLocationCount ensures that no value is present for LocationCount, not even an explicit nil
### GetLanguages

`func (o *PatientOrganisationRow) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *PatientOrganisationRow) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *PatientOrganisationRow) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *PatientOrganisationRow) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *PatientOrganisationRow) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *PatientOrganisationRow) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetPhones

`func (o *PatientOrganisationRow) GetPhones() []string`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *PatientOrganisationRow) GetPhonesOk() (*[]string, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *PatientOrganisationRow) SetPhones(v []string)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *PatientOrganisationRow) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### SetPhonesNil

`func (o *PatientOrganisationRow) SetPhonesNil(b bool)`

 SetPhonesNil sets the value for Phones to be an explicit nil

### UnsetPhones
`func (o *PatientOrganisationRow) UnsetPhones()`

UnsetPhones ensures that no value is present for Phones, not even an explicit nil
### GetEmails

`func (o *PatientOrganisationRow) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *PatientOrganisationRow) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *PatientOrganisationRow) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *PatientOrganisationRow) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### SetEmailsNil

`func (o *PatientOrganisationRow) SetEmailsNil(b bool)`

 SetEmailsNil sets the value for Emails to be an explicit nil

### UnsetEmails
`func (o *PatientOrganisationRow) UnsetEmails()`

UnsetEmails ensures that no value is present for Emails, not even an explicit nil
### GetContactUrls

`func (o *PatientOrganisationRow) GetContactUrls() []string`

GetContactUrls returns the ContactUrls field if non-nil, zero value otherwise.

### GetContactUrlsOk

`func (o *PatientOrganisationRow) GetContactUrlsOk() (*[]string, bool)`

GetContactUrlsOk returns a tuple with the ContactUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactUrls

`func (o *PatientOrganisationRow) SetContactUrls(v []string)`

SetContactUrls sets ContactUrls field to given value.

### HasContactUrls

`func (o *PatientOrganisationRow) HasContactUrls() bool`

HasContactUrls returns a boolean if a field has been set.

### SetContactUrlsNil

`func (o *PatientOrganisationRow) SetContactUrlsNil(b bool)`

 SetContactUrlsNil sets the value for ContactUrls to be an explicit nil

### UnsetContactUrls
`func (o *PatientOrganisationRow) UnsetContactUrls()`

UnsetContactUrls ensures that no value is present for ContactUrls, not even an explicit nil
### GetSocialUrls

`func (o *PatientOrganisationRow) GetSocialUrls() []string`

GetSocialUrls returns the SocialUrls field if non-nil, zero value otherwise.

### GetSocialUrlsOk

`func (o *PatientOrganisationRow) GetSocialUrlsOk() (*[]string, bool)`

GetSocialUrlsOk returns a tuple with the SocialUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialUrls

`func (o *PatientOrganisationRow) SetSocialUrls(v []string)`

SetSocialUrls sets SocialUrls field to given value.

### HasSocialUrls

`func (o *PatientOrganisationRow) HasSocialUrls() bool`

HasSocialUrls returns a boolean if a field has been set.

### SetSocialUrlsNil

`func (o *PatientOrganisationRow) SetSocialUrlsNil(b bool)`

 SetSocialUrlsNil sets the value for SocialUrls to be an explicit nil

### UnsetSocialUrls
`func (o *PatientOrganisationRow) UnsetSocialUrls()`

UnsetSocialUrls ensures that no value is present for SocialUrls, not even an explicit nil
### GetServiceCodes

`func (o *PatientOrganisationRow) GetServiceCodes() []string`

GetServiceCodes returns the ServiceCodes field if non-nil, zero value otherwise.

### GetServiceCodesOk

`func (o *PatientOrganisationRow) GetServiceCodesOk() (*[]string, bool)`

GetServiceCodesOk returns a tuple with the ServiceCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCodes

`func (o *PatientOrganisationRow) SetServiceCodes(v []string)`

SetServiceCodes sets ServiceCodes field to given value.

### HasServiceCodes

`func (o *PatientOrganisationRow) HasServiceCodes() bool`

HasServiceCodes returns a boolean if a field has been set.

### SetServiceCodesNil

`func (o *PatientOrganisationRow) SetServiceCodesNil(b bool)`

 SetServiceCodesNil sets the value for ServiceCodes to be an explicit nil

### UnsetServiceCodes
`func (o *PatientOrganisationRow) UnsetServiceCodes()`

UnsetServiceCodes ensures that no value is present for ServiceCodes, not even an explicit nil
### GetOtherServiceLabels

`func (o *PatientOrganisationRow) GetOtherServiceLabels() []string`

GetOtherServiceLabels returns the OtherServiceLabels field if non-nil, zero value otherwise.

### GetOtherServiceLabelsOk

`func (o *PatientOrganisationRow) GetOtherServiceLabelsOk() (*[]string, bool)`

GetOtherServiceLabelsOk returns a tuple with the OtherServiceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherServiceLabels

`func (o *PatientOrganisationRow) SetOtherServiceLabels(v []string)`

SetOtherServiceLabels sets OtherServiceLabels field to given value.

### HasOtherServiceLabels

`func (o *PatientOrganisationRow) HasOtherServiceLabels() bool`

HasOtherServiceLabels returns a boolean if a field has been set.

### SetOtherServiceLabelsNil

`func (o *PatientOrganisationRow) SetOtherServiceLabelsNil(b bool)`

 SetOtherServiceLabelsNil sets the value for OtherServiceLabels to be an explicit nil

### UnsetOtherServiceLabels
`func (o *PatientOrganisationRow) UnsetOtherServiceLabels()`

UnsetOtherServiceLabels ensures that no value is present for OtherServiceLabels, not even an explicit nil
### GetLinkKind

`func (o *PatientOrganisationRow) GetLinkKind() string`

GetLinkKind returns the LinkKind field if non-nil, zero value otherwise.

### GetLinkKindOk

`func (o *PatientOrganisationRow) GetLinkKindOk() (*string, bool)`

GetLinkKindOk returns a tuple with the LinkKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkKind

`func (o *PatientOrganisationRow) SetLinkKind(v string)`

SetLinkKind sets LinkKind field to given value.

### HasLinkKind

`func (o *PatientOrganisationRow) HasLinkKind() bool`

HasLinkKind returns a boolean if a field has been set.

### SetLinkKindNil

`func (o *PatientOrganisationRow) SetLinkKindNil(b bool)`

 SetLinkKindNil sets the value for LinkKind to be an explicit nil

### UnsetLinkKind
`func (o *PatientOrganisationRow) UnsetLinkKind()`

UnsetLinkKind ensures that no value is present for LinkKind, not even an explicit nil
### GetLinkRank

`func (o *PatientOrganisationRow) GetLinkRank() int32`

GetLinkRank returns the LinkRank field if non-nil, zero value otherwise.

### GetLinkRankOk

`func (o *PatientOrganisationRow) GetLinkRankOk() (*int32, bool)`

GetLinkRankOk returns a tuple with the LinkRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkRank

`func (o *PatientOrganisationRow) SetLinkRank(v int32)`

SetLinkRank sets LinkRank field to given value.

### HasLinkRank

`func (o *PatientOrganisationRow) HasLinkRank() bool`

HasLinkRank returns a boolean if a field has been set.

### SetLinkRankNil

`func (o *PatientOrganisationRow) SetLinkRankNil(b bool)`

 SetLinkRankNil sets the value for LinkRank to be an explicit nil

### UnsetLinkRank
`func (o *PatientOrganisationRow) UnsetLinkRank()`

UnsetLinkRank ensures that no value is present for LinkRank, not even an explicit nil
### GetDeclaredScope

`func (o *PatientOrganisationRow) GetDeclaredScope() string`

GetDeclaredScope returns the DeclaredScope field if non-nil, zero value otherwise.

### GetDeclaredScopeOk

`func (o *PatientOrganisationRow) GetDeclaredScopeOk() (*string, bool)`

GetDeclaredScopeOk returns a tuple with the DeclaredScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredScope

`func (o *PatientOrganisationRow) SetDeclaredScope(v string)`

SetDeclaredScope sets DeclaredScope field to given value.

### HasDeclaredScope

`func (o *PatientOrganisationRow) HasDeclaredScope() bool`

HasDeclaredScope returns a boolean if a field has been set.

### SetDeclaredScopeNil

`func (o *PatientOrganisationRow) SetDeclaredScopeNil(b bool)`

 SetDeclaredScopeNil sets the value for DeclaredScope to be an explicit nil

### UnsetDeclaredScope
`func (o *PatientOrganisationRow) UnsetDeclaredScope()`

UnsetDeclaredScope ensures that no value is present for DeclaredScope, not even an explicit nil
### GetEdgeVerdict

`func (o *PatientOrganisationRow) GetEdgeVerdict() string`

GetEdgeVerdict returns the EdgeVerdict field if non-nil, zero value otherwise.

### GetEdgeVerdictOk

`func (o *PatientOrganisationRow) GetEdgeVerdictOk() (*string, bool)`

GetEdgeVerdictOk returns a tuple with the EdgeVerdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeVerdict

`func (o *PatientOrganisationRow) SetEdgeVerdict(v string)`

SetEdgeVerdict sets EdgeVerdict field to given value.

### HasEdgeVerdict

`func (o *PatientOrganisationRow) HasEdgeVerdict() bool`

HasEdgeVerdict returns a boolean if a field has been set.

### SetEdgeVerdictNil

`func (o *PatientOrganisationRow) SetEdgeVerdictNil(b bool)`

 SetEdgeVerdictNil sets the value for EdgeVerdict to be an explicit nil

### UnsetEdgeVerdict
`func (o *PatientOrganisationRow) UnsetEdgeVerdict()`

UnsetEdgeVerdict ensures that no value is present for EdgeVerdict, not even an explicit nil
### GetDeclaredName

`func (o *PatientOrganisationRow) GetDeclaredName() string`

GetDeclaredName returns the DeclaredName field if non-nil, zero value otherwise.

### GetDeclaredNameOk

`func (o *PatientOrganisationRow) GetDeclaredNameOk() (*string, bool)`

GetDeclaredNameOk returns a tuple with the DeclaredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredName

`func (o *PatientOrganisationRow) SetDeclaredName(v string)`

SetDeclaredName sets DeclaredName field to given value.

### HasDeclaredName

`func (o *PatientOrganisationRow) HasDeclaredName() bool`

HasDeclaredName returns a boolean if a field has been set.

### SetDeclaredNameNil

`func (o *PatientOrganisationRow) SetDeclaredNameNil(b bool)`

 SetDeclaredNameNil sets the value for DeclaredName to be an explicit nil

### UnsetDeclaredName
`func (o *PatientOrganisationRow) UnsetDeclaredName()`

UnsetDeclaredName ensures that no value is present for DeclaredName, not even an explicit nil
### GetMappedOrphaCode

`func (o *PatientOrganisationRow) GetMappedOrphaCode() string`

GetMappedOrphaCode returns the MappedOrphaCode field if non-nil, zero value otherwise.

### GetMappedOrphaCodeOk

`func (o *PatientOrganisationRow) GetMappedOrphaCodeOk() (*string, bool)`

GetMappedOrphaCodeOk returns a tuple with the MappedOrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedOrphaCode

`func (o *PatientOrganisationRow) SetMappedOrphaCode(v string)`

SetMappedOrphaCode sets MappedOrphaCode field to given value.

### HasMappedOrphaCode

`func (o *PatientOrganisationRow) HasMappedOrphaCode() bool`

HasMappedOrphaCode returns a boolean if a field has been set.

### SetMappedOrphaCodeNil

`func (o *PatientOrganisationRow) SetMappedOrphaCodeNil(b bool)`

 SetMappedOrphaCodeNil sets the value for MappedOrphaCode to be an explicit nil

### UnsetMappedOrphaCode
`func (o *PatientOrganisationRow) UnsetMappedOrphaCode()`

UnsetMappedOrphaCode ensures that no value is present for MappedOrphaCode, not even an explicit nil
### GetMappedDiseaseName

`func (o *PatientOrganisationRow) GetMappedDiseaseName() string`

GetMappedDiseaseName returns the MappedDiseaseName field if non-nil, zero value otherwise.

### GetMappedDiseaseNameOk

`func (o *PatientOrganisationRow) GetMappedDiseaseNameOk() (*string, bool)`

GetMappedDiseaseNameOk returns a tuple with the MappedDiseaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedDiseaseName

`func (o *PatientOrganisationRow) SetMappedDiseaseName(v string)`

SetMappedDiseaseName sets MappedDiseaseName field to given value.

### HasMappedDiseaseName

`func (o *PatientOrganisationRow) HasMappedDiseaseName() bool`

HasMappedDiseaseName returns a boolean if a field has been set.

### SetMappedDiseaseNameNil

`func (o *PatientOrganisationRow) SetMappedDiseaseNameNil(b bool)`

 SetMappedDiseaseNameNil sets the value for MappedDiseaseName to be an explicit nil

### UnsetMappedDiseaseName
`func (o *PatientOrganisationRow) UnsetMappedDiseaseName()`

UnsetMappedDiseaseName ensures that no value is present for MappedDiseaseName, not even an explicit nil
### GetMappingRelation

`func (o *PatientOrganisationRow) GetMappingRelation() string`

GetMappingRelation returns the MappingRelation field if non-nil, zero value otherwise.

### GetMappingRelationOk

`func (o *PatientOrganisationRow) GetMappingRelationOk() (*string, bool)`

GetMappingRelationOk returns a tuple with the MappingRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingRelation

`func (o *PatientOrganisationRow) SetMappingRelation(v string)`

SetMappingRelation sets MappingRelation field to given value.

### HasMappingRelation

`func (o *PatientOrganisationRow) HasMappingRelation() bool`

HasMappingRelation returns a boolean if a field has been set.

### SetMappingRelationNil

`func (o *PatientOrganisationRow) SetMappingRelationNil(b bool)`

 SetMappingRelationNil sets the value for MappingRelation to be an explicit nil

### UnsetMappingRelation
`func (o *PatientOrganisationRow) UnsetMappingRelation()`

UnsetMappingRelation ensures that no value is present for MappingRelation, not even an explicit nil
### GetSourceUrl

`func (o *PatientOrganisationRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *PatientOrganisationRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *PatientOrganisationRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *PatientOrganisationRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *PatientOrganisationRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *PatientOrganisationRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetRetrievedOn

`func (o *PatientOrganisationRow) GetRetrievedOn() string`

GetRetrievedOn returns the RetrievedOn field if non-nil, zero value otherwise.

### GetRetrievedOnOk

`func (o *PatientOrganisationRow) GetRetrievedOnOk() (*string, bool)`

GetRetrievedOnOk returns a tuple with the RetrievedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievedOn

`func (o *PatientOrganisationRow) SetRetrievedOn(v string)`

SetRetrievedOn sets RetrievedOn field to given value.

### HasRetrievedOn

`func (o *PatientOrganisationRow) HasRetrievedOn() bool`

HasRetrievedOn returns a boolean if a field has been set.

### SetRetrievedOnNil

`func (o *PatientOrganisationRow) SetRetrievedOnNil(b bool)`

 SetRetrievedOnNil sets the value for RetrievedOn to be an explicit nil

### UnsetRetrievedOn
`func (o *PatientOrganisationRow) UnsetRetrievedOn()`

UnsetRetrievedOn ensures that no value is present for RetrievedOn, not even an explicit nil
### GetCredit

`func (o *PatientOrganisationRow) GetCredit() string`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *PatientOrganisationRow) GetCreditOk() (*string, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *PatientOrganisationRow) SetCredit(v string)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *PatientOrganisationRow) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### SetCreditNil

`func (o *PatientOrganisationRow) SetCreditNil(b bool)`

 SetCreditNil sets the value for Credit to be an explicit nil

### UnsetCredit
`func (o *PatientOrganisationRow) UnsetCredit()`

UnsetCredit ensures that no value is present for Credit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


