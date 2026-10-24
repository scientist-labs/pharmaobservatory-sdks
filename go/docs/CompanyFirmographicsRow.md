# CompanyFirmographicsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | Company level: the company the row is about; the entry key for the companies object type. | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**FoundedYear** | Pointer to **NullableInt32** | The earliest of the ROR and Wikidata founding years, else our provider profile, else the GLEIF entity creation year last (GLEIF dates re-registrations and reused shells wrongly). | [optional] 
**FoundedYearSource** | Pointer to **NullableString** | The source that supplied founded year. | [optional] 
**Status** | Pointer to **NullableString** | public when an organization record of the company carries an SEC filer id (not a Form D-only filer), else private. | [optional] 
**HqCity** | Pointer to **NullableString** | Headquarters city: GLEIF, then SEC filing business address, then ROR, then Wikidata. | [optional] 
**HqCitySource** | Pointer to **NullableString** | The source that supplied hq city. | [optional] 
**HqCountry** | Pointer to **NullableString** | Headquarters country: GLEIF, then ROR, then Wikidata. | [optional] 
**HqCountrySource** | Pointer to **NullableString** | The source that supplied hq country. | [optional] 
**PrimaryIndustry** | Pointer to **NullableString** | Primary industry: SEC standard industrial classification description, then ROR organization type, then Wikidata industry. | [optional] 
**PrimaryIndustrySource** | Pointer to **NullableString** | The source that supplied primary industry. | [optional] 
**Website** | Pointer to **NullableString** | Website: ROR, then Wikidata, then our provider profile. | [optional] 
**WebsiteSource** | Pointer to **NullableString** | The source that supplied website. | [optional] 
**Employees** | Pointer to **NullableInt64** | Latest employee headcount from Wikidata; a zero or negative value is treated as absent. | [optional] 
**EmployeesSource** | Pointer to **NullableString** | The source that supplied employees. | [optional] 
**EmployeesAsOf** | Pointer to **NullableString** | The date Wikidata states for the headcount. | [optional] 
**Verticals** | Pointer to **[]string** | Service areas from our provider profile (the pipe-separated list, split). | [optional] 
**VerticalsSource** | Pointer to **NullableString** | The source that supplied verticals. | [optional] 
**LegalName** | Pointer to **NullableString** | Legal name: GLEIF, then Wikidata, then our provider profile&#39;s official name. | [optional] 
**LegalNameSource** | Pointer to **NullableString** | The source that supplied legal name. | [optional] 
**CorporateParent** | Pointer to **NullableString** | Corporate parent name: GLEIF group name, then Wikidata parent, then our provider profile&#39;s parent. This is a stated source value, not the company&#39;s ladder group. | [optional] 
**CorporateParentSource** | Pointer to **NullableString** | The source that supplied corporate parent. | [optional] 
**LegalForm** | Pointer to **NullableString** | Legal form: GLEIF, then Wikidata. | [optional] 
**LegalFormSource** | Pointer to **NullableString** | The source that supplied legal form. | [optional] 
**WikidataQid** | Pointer to **NullableString** | The Wikidata item id of the organization record whose profile supplies the five fields below. | [optional] 
**WikipediaUrl** | Pointer to **NullableString** | Wikipedia page: ROR, then Wikidata. | [optional] 
**Linkedin** | Pointer to **NullableString** | LinkedIn organization id from the Wikidata profile. | [optional] 
**XHandle** | Pointer to **NullableString** | X (Twitter) username from the Wikidata profile. | [optional] 
**Facebook** | Pointer to **NullableString** | Facebook id from the Wikidata profile. | [optional] 
**ShortDescription** | Pointer to **NullableString** | The Wikidata short description. | [optional] 

## Methods

### NewCompanyFirmographicsRow

`func NewCompanyFirmographicsRow() *CompanyFirmographicsRow`

NewCompanyFirmographicsRow instantiates a new CompanyFirmographicsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyFirmographicsRowWithDefaults

`func NewCompanyFirmographicsRowWithDefaults() *CompanyFirmographicsRow`

NewCompanyFirmographicsRowWithDefaults instantiates a new CompanyFirmographicsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyFirmographicsRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyFirmographicsRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyFirmographicsRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyFirmographicsRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyFirmographicsRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyFirmographicsRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyFirmographicsRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyFirmographicsRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyFirmographicsRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyFirmographicsRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyFirmographicsRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyFirmographicsRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyFirmographicsRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyFirmographicsRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyFirmographicsRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyFirmographicsRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyFirmographicsRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyFirmographicsRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyFirmographicsRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyFirmographicsRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyFirmographicsRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyFirmographicsRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyFirmographicsRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyFirmographicsRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetFoundedYear

`func (o *CompanyFirmographicsRow) GetFoundedYear() int32`

GetFoundedYear returns the FoundedYear field if non-nil, zero value otherwise.

### GetFoundedYearOk

`func (o *CompanyFirmographicsRow) GetFoundedYearOk() (*int32, bool)`

GetFoundedYearOk returns a tuple with the FoundedYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundedYear

`func (o *CompanyFirmographicsRow) SetFoundedYear(v int32)`

SetFoundedYear sets FoundedYear field to given value.

### HasFoundedYear

`func (o *CompanyFirmographicsRow) HasFoundedYear() bool`

HasFoundedYear returns a boolean if a field has been set.

### SetFoundedYearNil

`func (o *CompanyFirmographicsRow) SetFoundedYearNil(b bool)`

 SetFoundedYearNil sets the value for FoundedYear to be an explicit nil

### UnsetFoundedYear
`func (o *CompanyFirmographicsRow) UnsetFoundedYear()`

UnsetFoundedYear ensures that no value is present for FoundedYear, not even an explicit nil
### GetFoundedYearSource

`func (o *CompanyFirmographicsRow) GetFoundedYearSource() string`

GetFoundedYearSource returns the FoundedYearSource field if non-nil, zero value otherwise.

### GetFoundedYearSourceOk

`func (o *CompanyFirmographicsRow) GetFoundedYearSourceOk() (*string, bool)`

GetFoundedYearSourceOk returns a tuple with the FoundedYearSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundedYearSource

`func (o *CompanyFirmographicsRow) SetFoundedYearSource(v string)`

SetFoundedYearSource sets FoundedYearSource field to given value.

### HasFoundedYearSource

`func (o *CompanyFirmographicsRow) HasFoundedYearSource() bool`

HasFoundedYearSource returns a boolean if a field has been set.

### SetFoundedYearSourceNil

`func (o *CompanyFirmographicsRow) SetFoundedYearSourceNil(b bool)`

 SetFoundedYearSourceNil sets the value for FoundedYearSource to be an explicit nil

### UnsetFoundedYearSource
`func (o *CompanyFirmographicsRow) UnsetFoundedYearSource()`

UnsetFoundedYearSource ensures that no value is present for FoundedYearSource, not even an explicit nil
### GetStatus

`func (o *CompanyFirmographicsRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompanyFirmographicsRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompanyFirmographicsRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CompanyFirmographicsRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CompanyFirmographicsRow) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CompanyFirmographicsRow) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetHqCity

`func (o *CompanyFirmographicsRow) GetHqCity() string`

GetHqCity returns the HqCity field if non-nil, zero value otherwise.

### GetHqCityOk

`func (o *CompanyFirmographicsRow) GetHqCityOk() (*string, bool)`

GetHqCityOk returns a tuple with the HqCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqCity

`func (o *CompanyFirmographicsRow) SetHqCity(v string)`

SetHqCity sets HqCity field to given value.

### HasHqCity

`func (o *CompanyFirmographicsRow) HasHqCity() bool`

HasHqCity returns a boolean if a field has been set.

### SetHqCityNil

`func (o *CompanyFirmographicsRow) SetHqCityNil(b bool)`

 SetHqCityNil sets the value for HqCity to be an explicit nil

### UnsetHqCity
`func (o *CompanyFirmographicsRow) UnsetHqCity()`

UnsetHqCity ensures that no value is present for HqCity, not even an explicit nil
### GetHqCitySource

`func (o *CompanyFirmographicsRow) GetHqCitySource() string`

GetHqCitySource returns the HqCitySource field if non-nil, zero value otherwise.

### GetHqCitySourceOk

`func (o *CompanyFirmographicsRow) GetHqCitySourceOk() (*string, bool)`

GetHqCitySourceOk returns a tuple with the HqCitySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqCitySource

`func (o *CompanyFirmographicsRow) SetHqCitySource(v string)`

SetHqCitySource sets HqCitySource field to given value.

### HasHqCitySource

`func (o *CompanyFirmographicsRow) HasHqCitySource() bool`

HasHqCitySource returns a boolean if a field has been set.

### SetHqCitySourceNil

`func (o *CompanyFirmographicsRow) SetHqCitySourceNil(b bool)`

 SetHqCitySourceNil sets the value for HqCitySource to be an explicit nil

### UnsetHqCitySource
`func (o *CompanyFirmographicsRow) UnsetHqCitySource()`

UnsetHqCitySource ensures that no value is present for HqCitySource, not even an explicit nil
### GetHqCountry

`func (o *CompanyFirmographicsRow) GetHqCountry() string`

GetHqCountry returns the HqCountry field if non-nil, zero value otherwise.

### GetHqCountryOk

`func (o *CompanyFirmographicsRow) GetHqCountryOk() (*string, bool)`

GetHqCountryOk returns a tuple with the HqCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqCountry

`func (o *CompanyFirmographicsRow) SetHqCountry(v string)`

SetHqCountry sets HqCountry field to given value.

### HasHqCountry

`func (o *CompanyFirmographicsRow) HasHqCountry() bool`

HasHqCountry returns a boolean if a field has been set.

### SetHqCountryNil

`func (o *CompanyFirmographicsRow) SetHqCountryNil(b bool)`

 SetHqCountryNil sets the value for HqCountry to be an explicit nil

### UnsetHqCountry
`func (o *CompanyFirmographicsRow) UnsetHqCountry()`

UnsetHqCountry ensures that no value is present for HqCountry, not even an explicit nil
### GetHqCountrySource

`func (o *CompanyFirmographicsRow) GetHqCountrySource() string`

GetHqCountrySource returns the HqCountrySource field if non-nil, zero value otherwise.

### GetHqCountrySourceOk

`func (o *CompanyFirmographicsRow) GetHqCountrySourceOk() (*string, bool)`

GetHqCountrySourceOk returns a tuple with the HqCountrySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqCountrySource

`func (o *CompanyFirmographicsRow) SetHqCountrySource(v string)`

SetHqCountrySource sets HqCountrySource field to given value.

### HasHqCountrySource

`func (o *CompanyFirmographicsRow) HasHqCountrySource() bool`

HasHqCountrySource returns a boolean if a field has been set.

### SetHqCountrySourceNil

`func (o *CompanyFirmographicsRow) SetHqCountrySourceNil(b bool)`

 SetHqCountrySourceNil sets the value for HqCountrySource to be an explicit nil

### UnsetHqCountrySource
`func (o *CompanyFirmographicsRow) UnsetHqCountrySource()`

UnsetHqCountrySource ensures that no value is present for HqCountrySource, not even an explicit nil
### GetPrimaryIndustry

`func (o *CompanyFirmographicsRow) GetPrimaryIndustry() string`

GetPrimaryIndustry returns the PrimaryIndustry field if non-nil, zero value otherwise.

### GetPrimaryIndustryOk

`func (o *CompanyFirmographicsRow) GetPrimaryIndustryOk() (*string, bool)`

GetPrimaryIndustryOk returns a tuple with the PrimaryIndustry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIndustry

`func (o *CompanyFirmographicsRow) SetPrimaryIndustry(v string)`

SetPrimaryIndustry sets PrimaryIndustry field to given value.

### HasPrimaryIndustry

`func (o *CompanyFirmographicsRow) HasPrimaryIndustry() bool`

HasPrimaryIndustry returns a boolean if a field has been set.

### SetPrimaryIndustryNil

`func (o *CompanyFirmographicsRow) SetPrimaryIndustryNil(b bool)`

 SetPrimaryIndustryNil sets the value for PrimaryIndustry to be an explicit nil

### UnsetPrimaryIndustry
`func (o *CompanyFirmographicsRow) UnsetPrimaryIndustry()`

UnsetPrimaryIndustry ensures that no value is present for PrimaryIndustry, not even an explicit nil
### GetPrimaryIndustrySource

`func (o *CompanyFirmographicsRow) GetPrimaryIndustrySource() string`

GetPrimaryIndustrySource returns the PrimaryIndustrySource field if non-nil, zero value otherwise.

### GetPrimaryIndustrySourceOk

`func (o *CompanyFirmographicsRow) GetPrimaryIndustrySourceOk() (*string, bool)`

GetPrimaryIndustrySourceOk returns a tuple with the PrimaryIndustrySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIndustrySource

`func (o *CompanyFirmographicsRow) SetPrimaryIndustrySource(v string)`

SetPrimaryIndustrySource sets PrimaryIndustrySource field to given value.

### HasPrimaryIndustrySource

`func (o *CompanyFirmographicsRow) HasPrimaryIndustrySource() bool`

HasPrimaryIndustrySource returns a boolean if a field has been set.

### SetPrimaryIndustrySourceNil

`func (o *CompanyFirmographicsRow) SetPrimaryIndustrySourceNil(b bool)`

 SetPrimaryIndustrySourceNil sets the value for PrimaryIndustrySource to be an explicit nil

### UnsetPrimaryIndustrySource
`func (o *CompanyFirmographicsRow) UnsetPrimaryIndustrySource()`

UnsetPrimaryIndustrySource ensures that no value is present for PrimaryIndustrySource, not even an explicit nil
### GetWebsite

`func (o *CompanyFirmographicsRow) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CompanyFirmographicsRow) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CompanyFirmographicsRow) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *CompanyFirmographicsRow) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *CompanyFirmographicsRow) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *CompanyFirmographicsRow) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetWebsiteSource

`func (o *CompanyFirmographicsRow) GetWebsiteSource() string`

GetWebsiteSource returns the WebsiteSource field if non-nil, zero value otherwise.

### GetWebsiteSourceOk

`func (o *CompanyFirmographicsRow) GetWebsiteSourceOk() (*string, bool)`

GetWebsiteSourceOk returns a tuple with the WebsiteSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteSource

`func (o *CompanyFirmographicsRow) SetWebsiteSource(v string)`

SetWebsiteSource sets WebsiteSource field to given value.

### HasWebsiteSource

`func (o *CompanyFirmographicsRow) HasWebsiteSource() bool`

HasWebsiteSource returns a boolean if a field has been set.

### SetWebsiteSourceNil

`func (o *CompanyFirmographicsRow) SetWebsiteSourceNil(b bool)`

 SetWebsiteSourceNil sets the value for WebsiteSource to be an explicit nil

### UnsetWebsiteSource
`func (o *CompanyFirmographicsRow) UnsetWebsiteSource()`

UnsetWebsiteSource ensures that no value is present for WebsiteSource, not even an explicit nil
### GetEmployees

`func (o *CompanyFirmographicsRow) GetEmployees() int64`

GetEmployees returns the Employees field if non-nil, zero value otherwise.

### GetEmployeesOk

`func (o *CompanyFirmographicsRow) GetEmployeesOk() (*int64, bool)`

GetEmployeesOk returns a tuple with the Employees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployees

`func (o *CompanyFirmographicsRow) SetEmployees(v int64)`

SetEmployees sets Employees field to given value.

### HasEmployees

`func (o *CompanyFirmographicsRow) HasEmployees() bool`

HasEmployees returns a boolean if a field has been set.

### SetEmployeesNil

`func (o *CompanyFirmographicsRow) SetEmployeesNil(b bool)`

 SetEmployeesNil sets the value for Employees to be an explicit nil

### UnsetEmployees
`func (o *CompanyFirmographicsRow) UnsetEmployees()`

UnsetEmployees ensures that no value is present for Employees, not even an explicit nil
### GetEmployeesSource

`func (o *CompanyFirmographicsRow) GetEmployeesSource() string`

GetEmployeesSource returns the EmployeesSource field if non-nil, zero value otherwise.

### GetEmployeesSourceOk

`func (o *CompanyFirmographicsRow) GetEmployeesSourceOk() (*string, bool)`

GetEmployeesSourceOk returns a tuple with the EmployeesSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeesSource

`func (o *CompanyFirmographicsRow) SetEmployeesSource(v string)`

SetEmployeesSource sets EmployeesSource field to given value.

### HasEmployeesSource

`func (o *CompanyFirmographicsRow) HasEmployeesSource() bool`

HasEmployeesSource returns a boolean if a field has been set.

### SetEmployeesSourceNil

`func (o *CompanyFirmographicsRow) SetEmployeesSourceNil(b bool)`

 SetEmployeesSourceNil sets the value for EmployeesSource to be an explicit nil

### UnsetEmployeesSource
`func (o *CompanyFirmographicsRow) UnsetEmployeesSource()`

UnsetEmployeesSource ensures that no value is present for EmployeesSource, not even an explicit nil
### GetEmployeesAsOf

`func (o *CompanyFirmographicsRow) GetEmployeesAsOf() string`

GetEmployeesAsOf returns the EmployeesAsOf field if non-nil, zero value otherwise.

### GetEmployeesAsOfOk

`func (o *CompanyFirmographicsRow) GetEmployeesAsOfOk() (*string, bool)`

GetEmployeesAsOfOk returns a tuple with the EmployeesAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeesAsOf

`func (o *CompanyFirmographicsRow) SetEmployeesAsOf(v string)`

SetEmployeesAsOf sets EmployeesAsOf field to given value.

### HasEmployeesAsOf

`func (o *CompanyFirmographicsRow) HasEmployeesAsOf() bool`

HasEmployeesAsOf returns a boolean if a field has been set.

### SetEmployeesAsOfNil

`func (o *CompanyFirmographicsRow) SetEmployeesAsOfNil(b bool)`

 SetEmployeesAsOfNil sets the value for EmployeesAsOf to be an explicit nil

### UnsetEmployeesAsOf
`func (o *CompanyFirmographicsRow) UnsetEmployeesAsOf()`

UnsetEmployeesAsOf ensures that no value is present for EmployeesAsOf, not even an explicit nil
### GetVerticals

`func (o *CompanyFirmographicsRow) GetVerticals() []string`

GetVerticals returns the Verticals field if non-nil, zero value otherwise.

### GetVerticalsOk

`func (o *CompanyFirmographicsRow) GetVerticalsOk() (*[]string, bool)`

GetVerticalsOk returns a tuple with the Verticals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticals

`func (o *CompanyFirmographicsRow) SetVerticals(v []string)`

SetVerticals sets Verticals field to given value.

### HasVerticals

`func (o *CompanyFirmographicsRow) HasVerticals() bool`

HasVerticals returns a boolean if a field has been set.

### SetVerticalsNil

`func (o *CompanyFirmographicsRow) SetVerticalsNil(b bool)`

 SetVerticalsNil sets the value for Verticals to be an explicit nil

### UnsetVerticals
`func (o *CompanyFirmographicsRow) UnsetVerticals()`

UnsetVerticals ensures that no value is present for Verticals, not even an explicit nil
### GetVerticalsSource

`func (o *CompanyFirmographicsRow) GetVerticalsSource() string`

GetVerticalsSource returns the VerticalsSource field if non-nil, zero value otherwise.

### GetVerticalsSourceOk

`func (o *CompanyFirmographicsRow) GetVerticalsSourceOk() (*string, bool)`

GetVerticalsSourceOk returns a tuple with the VerticalsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalsSource

`func (o *CompanyFirmographicsRow) SetVerticalsSource(v string)`

SetVerticalsSource sets VerticalsSource field to given value.

### HasVerticalsSource

`func (o *CompanyFirmographicsRow) HasVerticalsSource() bool`

HasVerticalsSource returns a boolean if a field has been set.

### SetVerticalsSourceNil

`func (o *CompanyFirmographicsRow) SetVerticalsSourceNil(b bool)`

 SetVerticalsSourceNil sets the value for VerticalsSource to be an explicit nil

### UnsetVerticalsSource
`func (o *CompanyFirmographicsRow) UnsetVerticalsSource()`

UnsetVerticalsSource ensures that no value is present for VerticalsSource, not even an explicit nil
### GetLegalName

`func (o *CompanyFirmographicsRow) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *CompanyFirmographicsRow) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *CompanyFirmographicsRow) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *CompanyFirmographicsRow) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### SetLegalNameNil

`func (o *CompanyFirmographicsRow) SetLegalNameNil(b bool)`

 SetLegalNameNil sets the value for LegalName to be an explicit nil

### UnsetLegalName
`func (o *CompanyFirmographicsRow) UnsetLegalName()`

UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
### GetLegalNameSource

`func (o *CompanyFirmographicsRow) GetLegalNameSource() string`

GetLegalNameSource returns the LegalNameSource field if non-nil, zero value otherwise.

### GetLegalNameSourceOk

`func (o *CompanyFirmographicsRow) GetLegalNameSourceOk() (*string, bool)`

GetLegalNameSourceOk returns a tuple with the LegalNameSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalNameSource

`func (o *CompanyFirmographicsRow) SetLegalNameSource(v string)`

SetLegalNameSource sets LegalNameSource field to given value.

### HasLegalNameSource

`func (o *CompanyFirmographicsRow) HasLegalNameSource() bool`

HasLegalNameSource returns a boolean if a field has been set.

### SetLegalNameSourceNil

`func (o *CompanyFirmographicsRow) SetLegalNameSourceNil(b bool)`

 SetLegalNameSourceNil sets the value for LegalNameSource to be an explicit nil

### UnsetLegalNameSource
`func (o *CompanyFirmographicsRow) UnsetLegalNameSource()`

UnsetLegalNameSource ensures that no value is present for LegalNameSource, not even an explicit nil
### GetCorporateParent

`func (o *CompanyFirmographicsRow) GetCorporateParent() string`

GetCorporateParent returns the CorporateParent field if non-nil, zero value otherwise.

### GetCorporateParentOk

`func (o *CompanyFirmographicsRow) GetCorporateParentOk() (*string, bool)`

GetCorporateParentOk returns a tuple with the CorporateParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateParent

`func (o *CompanyFirmographicsRow) SetCorporateParent(v string)`

SetCorporateParent sets CorporateParent field to given value.

### HasCorporateParent

`func (o *CompanyFirmographicsRow) HasCorporateParent() bool`

HasCorporateParent returns a boolean if a field has been set.

### SetCorporateParentNil

`func (o *CompanyFirmographicsRow) SetCorporateParentNil(b bool)`

 SetCorporateParentNil sets the value for CorporateParent to be an explicit nil

### UnsetCorporateParent
`func (o *CompanyFirmographicsRow) UnsetCorporateParent()`

UnsetCorporateParent ensures that no value is present for CorporateParent, not even an explicit nil
### GetCorporateParentSource

`func (o *CompanyFirmographicsRow) GetCorporateParentSource() string`

GetCorporateParentSource returns the CorporateParentSource field if non-nil, zero value otherwise.

### GetCorporateParentSourceOk

`func (o *CompanyFirmographicsRow) GetCorporateParentSourceOk() (*string, bool)`

GetCorporateParentSourceOk returns a tuple with the CorporateParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateParentSource

`func (o *CompanyFirmographicsRow) SetCorporateParentSource(v string)`

SetCorporateParentSource sets CorporateParentSource field to given value.

### HasCorporateParentSource

`func (o *CompanyFirmographicsRow) HasCorporateParentSource() bool`

HasCorporateParentSource returns a boolean if a field has been set.

### SetCorporateParentSourceNil

`func (o *CompanyFirmographicsRow) SetCorporateParentSourceNil(b bool)`

 SetCorporateParentSourceNil sets the value for CorporateParentSource to be an explicit nil

### UnsetCorporateParentSource
`func (o *CompanyFirmographicsRow) UnsetCorporateParentSource()`

UnsetCorporateParentSource ensures that no value is present for CorporateParentSource, not even an explicit nil
### GetLegalForm

`func (o *CompanyFirmographicsRow) GetLegalForm() string`

GetLegalForm returns the LegalForm field if non-nil, zero value otherwise.

### GetLegalFormOk

`func (o *CompanyFirmographicsRow) GetLegalFormOk() (*string, bool)`

GetLegalFormOk returns a tuple with the LegalForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalForm

`func (o *CompanyFirmographicsRow) SetLegalForm(v string)`

SetLegalForm sets LegalForm field to given value.

### HasLegalForm

`func (o *CompanyFirmographicsRow) HasLegalForm() bool`

HasLegalForm returns a boolean if a field has been set.

### SetLegalFormNil

`func (o *CompanyFirmographicsRow) SetLegalFormNil(b bool)`

 SetLegalFormNil sets the value for LegalForm to be an explicit nil

### UnsetLegalForm
`func (o *CompanyFirmographicsRow) UnsetLegalForm()`

UnsetLegalForm ensures that no value is present for LegalForm, not even an explicit nil
### GetLegalFormSource

`func (o *CompanyFirmographicsRow) GetLegalFormSource() string`

GetLegalFormSource returns the LegalFormSource field if non-nil, zero value otherwise.

### GetLegalFormSourceOk

`func (o *CompanyFirmographicsRow) GetLegalFormSourceOk() (*string, bool)`

GetLegalFormSourceOk returns a tuple with the LegalFormSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalFormSource

`func (o *CompanyFirmographicsRow) SetLegalFormSource(v string)`

SetLegalFormSource sets LegalFormSource field to given value.

### HasLegalFormSource

`func (o *CompanyFirmographicsRow) HasLegalFormSource() bool`

HasLegalFormSource returns a boolean if a field has been set.

### SetLegalFormSourceNil

`func (o *CompanyFirmographicsRow) SetLegalFormSourceNil(b bool)`

 SetLegalFormSourceNil sets the value for LegalFormSource to be an explicit nil

### UnsetLegalFormSource
`func (o *CompanyFirmographicsRow) UnsetLegalFormSource()`

UnsetLegalFormSource ensures that no value is present for LegalFormSource, not even an explicit nil
### GetWikidataQid

`func (o *CompanyFirmographicsRow) GetWikidataQid() string`

GetWikidataQid returns the WikidataQid field if non-nil, zero value otherwise.

### GetWikidataQidOk

`func (o *CompanyFirmographicsRow) GetWikidataQidOk() (*string, bool)`

GetWikidataQidOk returns a tuple with the WikidataQid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikidataQid

`func (o *CompanyFirmographicsRow) SetWikidataQid(v string)`

SetWikidataQid sets WikidataQid field to given value.

### HasWikidataQid

`func (o *CompanyFirmographicsRow) HasWikidataQid() bool`

HasWikidataQid returns a boolean if a field has been set.

### SetWikidataQidNil

`func (o *CompanyFirmographicsRow) SetWikidataQidNil(b bool)`

 SetWikidataQidNil sets the value for WikidataQid to be an explicit nil

### UnsetWikidataQid
`func (o *CompanyFirmographicsRow) UnsetWikidataQid()`

UnsetWikidataQid ensures that no value is present for WikidataQid, not even an explicit nil
### GetWikipediaUrl

`func (o *CompanyFirmographicsRow) GetWikipediaUrl() string`

GetWikipediaUrl returns the WikipediaUrl field if non-nil, zero value otherwise.

### GetWikipediaUrlOk

`func (o *CompanyFirmographicsRow) GetWikipediaUrlOk() (*string, bool)`

GetWikipediaUrlOk returns a tuple with the WikipediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikipediaUrl

`func (o *CompanyFirmographicsRow) SetWikipediaUrl(v string)`

SetWikipediaUrl sets WikipediaUrl field to given value.

### HasWikipediaUrl

`func (o *CompanyFirmographicsRow) HasWikipediaUrl() bool`

HasWikipediaUrl returns a boolean if a field has been set.

### SetWikipediaUrlNil

`func (o *CompanyFirmographicsRow) SetWikipediaUrlNil(b bool)`

 SetWikipediaUrlNil sets the value for WikipediaUrl to be an explicit nil

### UnsetWikipediaUrl
`func (o *CompanyFirmographicsRow) UnsetWikipediaUrl()`

UnsetWikipediaUrl ensures that no value is present for WikipediaUrl, not even an explicit nil
### GetLinkedin

`func (o *CompanyFirmographicsRow) GetLinkedin() string`

GetLinkedin returns the Linkedin field if non-nil, zero value otherwise.

### GetLinkedinOk

`func (o *CompanyFirmographicsRow) GetLinkedinOk() (*string, bool)`

GetLinkedinOk returns a tuple with the Linkedin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedin

`func (o *CompanyFirmographicsRow) SetLinkedin(v string)`

SetLinkedin sets Linkedin field to given value.

### HasLinkedin

`func (o *CompanyFirmographicsRow) HasLinkedin() bool`

HasLinkedin returns a boolean if a field has been set.

### SetLinkedinNil

`func (o *CompanyFirmographicsRow) SetLinkedinNil(b bool)`

 SetLinkedinNil sets the value for Linkedin to be an explicit nil

### UnsetLinkedin
`func (o *CompanyFirmographicsRow) UnsetLinkedin()`

UnsetLinkedin ensures that no value is present for Linkedin, not even an explicit nil
### GetXHandle

`func (o *CompanyFirmographicsRow) GetXHandle() string`

GetXHandle returns the XHandle field if non-nil, zero value otherwise.

### GetXHandleOk

`func (o *CompanyFirmographicsRow) GetXHandleOk() (*string, bool)`

GetXHandleOk returns a tuple with the XHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXHandle

`func (o *CompanyFirmographicsRow) SetXHandle(v string)`

SetXHandle sets XHandle field to given value.

### HasXHandle

`func (o *CompanyFirmographicsRow) HasXHandle() bool`

HasXHandle returns a boolean if a field has been set.

### SetXHandleNil

`func (o *CompanyFirmographicsRow) SetXHandleNil(b bool)`

 SetXHandleNil sets the value for XHandle to be an explicit nil

### UnsetXHandle
`func (o *CompanyFirmographicsRow) UnsetXHandle()`

UnsetXHandle ensures that no value is present for XHandle, not even an explicit nil
### GetFacebook

`func (o *CompanyFirmographicsRow) GetFacebook() string`

GetFacebook returns the Facebook field if non-nil, zero value otherwise.

### GetFacebookOk

`func (o *CompanyFirmographicsRow) GetFacebookOk() (*string, bool)`

GetFacebookOk returns a tuple with the Facebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebook

`func (o *CompanyFirmographicsRow) SetFacebook(v string)`

SetFacebook sets Facebook field to given value.

### HasFacebook

`func (o *CompanyFirmographicsRow) HasFacebook() bool`

HasFacebook returns a boolean if a field has been set.

### SetFacebookNil

`func (o *CompanyFirmographicsRow) SetFacebookNil(b bool)`

 SetFacebookNil sets the value for Facebook to be an explicit nil

### UnsetFacebook
`func (o *CompanyFirmographicsRow) UnsetFacebook()`

UnsetFacebook ensures that no value is present for Facebook, not even an explicit nil
### GetShortDescription

`func (o *CompanyFirmographicsRow) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *CompanyFirmographicsRow) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *CompanyFirmographicsRow) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *CompanyFirmographicsRow) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### SetShortDescriptionNil

`func (o *CompanyFirmographicsRow) SetShortDescriptionNil(b bool)`

 SetShortDescriptionNil sets the value for ShortDescription to be an explicit nil

### UnsetShortDescription
`func (o *CompanyFirmographicsRow) UnsetShortDescription()`

UnsetShortDescription ensures that no value is present for ShortDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


