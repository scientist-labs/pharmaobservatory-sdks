# DrugRegulatoryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**FdaApproved** | Pointer to **NullableBool** | True when the drug has a record in Drugs@FDA. | [optional] 
**InNdc** | Pointer to **NullableBool** | True when the drug has a package in the National Drug Code (NDC) directory. | [optional] 
**HasGeneric** | Pointer to **NullableBool** | True when a generic is on record. | [optional] 
**BiosimilarAvailable** | Pointer to **NullableBool** | True when the Purple Book lists a biosimilar. | [optional] 
**InterchangeableAvailable** | Pointer to **NullableBool** | True when the Purple Book lists an interchangeable biosimilar. | [optional] 
**FirstLicensureDate** | Pointer to **NullableString** | The earliest biologic licensure date in the Purple Book. | [optional] 
**PteExpiry** | Pointer to **NullableString** | The extended expiry of a granted patent term extension. | [optional] 
**PatentProtected** | Pointer to **NullableBool** | True when an Orange Book patent is still active. | [optional] 
**LatestPatentExpiry** | Pointer to **NullableString** | The latest expiry among the Orange Book patents. | [optional] 
**ApplicantName** | Pointer to **NullableString** | The applicant (company) as filed in the Orange or Purple Book. | [optional] 

## Methods

### NewDrugRegulatoryRow

`func NewDrugRegulatoryRow() *DrugRegulatoryRow`

NewDrugRegulatoryRow instantiates a new DrugRegulatoryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugRegulatoryRowWithDefaults

`func NewDrugRegulatoryRowWithDefaults() *DrugRegulatoryRow`

NewDrugRegulatoryRowWithDefaults instantiates a new DrugRegulatoryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugRegulatoryRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugRegulatoryRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugRegulatoryRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugRegulatoryRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugRegulatoryRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugRegulatoryRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetFdaApproved

`func (o *DrugRegulatoryRow) GetFdaApproved() bool`

GetFdaApproved returns the FdaApproved field if non-nil, zero value otherwise.

### GetFdaApprovedOk

`func (o *DrugRegulatoryRow) GetFdaApprovedOk() (*bool, bool)`

GetFdaApprovedOk returns a tuple with the FdaApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdaApproved

`func (o *DrugRegulatoryRow) SetFdaApproved(v bool)`

SetFdaApproved sets FdaApproved field to given value.

### HasFdaApproved

`func (o *DrugRegulatoryRow) HasFdaApproved() bool`

HasFdaApproved returns a boolean if a field has been set.

### SetFdaApprovedNil

`func (o *DrugRegulatoryRow) SetFdaApprovedNil(b bool)`

 SetFdaApprovedNil sets the value for FdaApproved to be an explicit nil

### UnsetFdaApproved
`func (o *DrugRegulatoryRow) UnsetFdaApproved()`

UnsetFdaApproved ensures that no value is present for FdaApproved, not even an explicit nil
### GetInNdc

`func (o *DrugRegulatoryRow) GetInNdc() bool`

GetInNdc returns the InNdc field if non-nil, zero value otherwise.

### GetInNdcOk

`func (o *DrugRegulatoryRow) GetInNdcOk() (*bool, bool)`

GetInNdcOk returns a tuple with the InNdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInNdc

`func (o *DrugRegulatoryRow) SetInNdc(v bool)`

SetInNdc sets InNdc field to given value.

### HasInNdc

`func (o *DrugRegulatoryRow) HasInNdc() bool`

HasInNdc returns a boolean if a field has been set.

### SetInNdcNil

`func (o *DrugRegulatoryRow) SetInNdcNil(b bool)`

 SetInNdcNil sets the value for InNdc to be an explicit nil

### UnsetInNdc
`func (o *DrugRegulatoryRow) UnsetInNdc()`

UnsetInNdc ensures that no value is present for InNdc, not even an explicit nil
### GetHasGeneric

`func (o *DrugRegulatoryRow) GetHasGeneric() bool`

GetHasGeneric returns the HasGeneric field if non-nil, zero value otherwise.

### GetHasGenericOk

`func (o *DrugRegulatoryRow) GetHasGenericOk() (*bool, bool)`

GetHasGenericOk returns a tuple with the HasGeneric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGeneric

`func (o *DrugRegulatoryRow) SetHasGeneric(v bool)`

SetHasGeneric sets HasGeneric field to given value.

### HasHasGeneric

`func (o *DrugRegulatoryRow) HasHasGeneric() bool`

HasHasGeneric returns a boolean if a field has been set.

### SetHasGenericNil

`func (o *DrugRegulatoryRow) SetHasGenericNil(b bool)`

 SetHasGenericNil sets the value for HasGeneric to be an explicit nil

### UnsetHasGeneric
`func (o *DrugRegulatoryRow) UnsetHasGeneric()`

UnsetHasGeneric ensures that no value is present for HasGeneric, not even an explicit nil
### GetBiosimilarAvailable

`func (o *DrugRegulatoryRow) GetBiosimilarAvailable() bool`

GetBiosimilarAvailable returns the BiosimilarAvailable field if non-nil, zero value otherwise.

### GetBiosimilarAvailableOk

`func (o *DrugRegulatoryRow) GetBiosimilarAvailableOk() (*bool, bool)`

GetBiosimilarAvailableOk returns a tuple with the BiosimilarAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosimilarAvailable

`func (o *DrugRegulatoryRow) SetBiosimilarAvailable(v bool)`

SetBiosimilarAvailable sets BiosimilarAvailable field to given value.

### HasBiosimilarAvailable

`func (o *DrugRegulatoryRow) HasBiosimilarAvailable() bool`

HasBiosimilarAvailable returns a boolean if a field has been set.

### SetBiosimilarAvailableNil

`func (o *DrugRegulatoryRow) SetBiosimilarAvailableNil(b bool)`

 SetBiosimilarAvailableNil sets the value for BiosimilarAvailable to be an explicit nil

### UnsetBiosimilarAvailable
`func (o *DrugRegulatoryRow) UnsetBiosimilarAvailable()`

UnsetBiosimilarAvailable ensures that no value is present for BiosimilarAvailable, not even an explicit nil
### GetInterchangeableAvailable

`func (o *DrugRegulatoryRow) GetInterchangeableAvailable() bool`

GetInterchangeableAvailable returns the InterchangeableAvailable field if non-nil, zero value otherwise.

### GetInterchangeableAvailableOk

`func (o *DrugRegulatoryRow) GetInterchangeableAvailableOk() (*bool, bool)`

GetInterchangeableAvailableOk returns a tuple with the InterchangeableAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterchangeableAvailable

`func (o *DrugRegulatoryRow) SetInterchangeableAvailable(v bool)`

SetInterchangeableAvailable sets InterchangeableAvailable field to given value.

### HasInterchangeableAvailable

`func (o *DrugRegulatoryRow) HasInterchangeableAvailable() bool`

HasInterchangeableAvailable returns a boolean if a field has been set.

### SetInterchangeableAvailableNil

`func (o *DrugRegulatoryRow) SetInterchangeableAvailableNil(b bool)`

 SetInterchangeableAvailableNil sets the value for InterchangeableAvailable to be an explicit nil

### UnsetInterchangeableAvailable
`func (o *DrugRegulatoryRow) UnsetInterchangeableAvailable()`

UnsetInterchangeableAvailable ensures that no value is present for InterchangeableAvailable, not even an explicit nil
### GetFirstLicensureDate

`func (o *DrugRegulatoryRow) GetFirstLicensureDate() string`

GetFirstLicensureDate returns the FirstLicensureDate field if non-nil, zero value otherwise.

### GetFirstLicensureDateOk

`func (o *DrugRegulatoryRow) GetFirstLicensureDateOk() (*string, bool)`

GetFirstLicensureDateOk returns a tuple with the FirstLicensureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLicensureDate

`func (o *DrugRegulatoryRow) SetFirstLicensureDate(v string)`

SetFirstLicensureDate sets FirstLicensureDate field to given value.

### HasFirstLicensureDate

`func (o *DrugRegulatoryRow) HasFirstLicensureDate() bool`

HasFirstLicensureDate returns a boolean if a field has been set.

### SetFirstLicensureDateNil

`func (o *DrugRegulatoryRow) SetFirstLicensureDateNil(b bool)`

 SetFirstLicensureDateNil sets the value for FirstLicensureDate to be an explicit nil

### UnsetFirstLicensureDate
`func (o *DrugRegulatoryRow) UnsetFirstLicensureDate()`

UnsetFirstLicensureDate ensures that no value is present for FirstLicensureDate, not even an explicit nil
### GetPteExpiry

`func (o *DrugRegulatoryRow) GetPteExpiry() string`

GetPteExpiry returns the PteExpiry field if non-nil, zero value otherwise.

### GetPteExpiryOk

`func (o *DrugRegulatoryRow) GetPteExpiryOk() (*string, bool)`

GetPteExpiryOk returns a tuple with the PteExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPteExpiry

`func (o *DrugRegulatoryRow) SetPteExpiry(v string)`

SetPteExpiry sets PteExpiry field to given value.

### HasPteExpiry

`func (o *DrugRegulatoryRow) HasPteExpiry() bool`

HasPteExpiry returns a boolean if a field has been set.

### SetPteExpiryNil

`func (o *DrugRegulatoryRow) SetPteExpiryNil(b bool)`

 SetPteExpiryNil sets the value for PteExpiry to be an explicit nil

### UnsetPteExpiry
`func (o *DrugRegulatoryRow) UnsetPteExpiry()`

UnsetPteExpiry ensures that no value is present for PteExpiry, not even an explicit nil
### GetPatentProtected

`func (o *DrugRegulatoryRow) GetPatentProtected() bool`

GetPatentProtected returns the PatentProtected field if non-nil, zero value otherwise.

### GetPatentProtectedOk

`func (o *DrugRegulatoryRow) GetPatentProtectedOk() (*bool, bool)`

GetPatentProtectedOk returns a tuple with the PatentProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatentProtected

`func (o *DrugRegulatoryRow) SetPatentProtected(v bool)`

SetPatentProtected sets PatentProtected field to given value.

### HasPatentProtected

`func (o *DrugRegulatoryRow) HasPatentProtected() bool`

HasPatentProtected returns a boolean if a field has been set.

### SetPatentProtectedNil

`func (o *DrugRegulatoryRow) SetPatentProtectedNil(b bool)`

 SetPatentProtectedNil sets the value for PatentProtected to be an explicit nil

### UnsetPatentProtected
`func (o *DrugRegulatoryRow) UnsetPatentProtected()`

UnsetPatentProtected ensures that no value is present for PatentProtected, not even an explicit nil
### GetLatestPatentExpiry

`func (o *DrugRegulatoryRow) GetLatestPatentExpiry() string`

GetLatestPatentExpiry returns the LatestPatentExpiry field if non-nil, zero value otherwise.

### GetLatestPatentExpiryOk

`func (o *DrugRegulatoryRow) GetLatestPatentExpiryOk() (*string, bool)`

GetLatestPatentExpiryOk returns a tuple with the LatestPatentExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestPatentExpiry

`func (o *DrugRegulatoryRow) SetLatestPatentExpiry(v string)`

SetLatestPatentExpiry sets LatestPatentExpiry field to given value.

### HasLatestPatentExpiry

`func (o *DrugRegulatoryRow) HasLatestPatentExpiry() bool`

HasLatestPatentExpiry returns a boolean if a field has been set.

### SetLatestPatentExpiryNil

`func (o *DrugRegulatoryRow) SetLatestPatentExpiryNil(b bool)`

 SetLatestPatentExpiryNil sets the value for LatestPatentExpiry to be an explicit nil

### UnsetLatestPatentExpiry
`func (o *DrugRegulatoryRow) UnsetLatestPatentExpiry()`

UnsetLatestPatentExpiry ensures that no value is present for LatestPatentExpiry, not even an explicit nil
### GetApplicantName

`func (o *DrugRegulatoryRow) GetApplicantName() string`

GetApplicantName returns the ApplicantName field if non-nil, zero value otherwise.

### GetApplicantNameOk

`func (o *DrugRegulatoryRow) GetApplicantNameOk() (*string, bool)`

GetApplicantNameOk returns a tuple with the ApplicantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantName

`func (o *DrugRegulatoryRow) SetApplicantName(v string)`

SetApplicantName sets ApplicantName field to given value.

### HasApplicantName

`func (o *DrugRegulatoryRow) HasApplicantName() bool`

HasApplicantName returns a boolean if a field has been set.

### SetApplicantNameNil

`func (o *DrugRegulatoryRow) SetApplicantNameNil(b bool)`

 SetApplicantNameNil sets the value for ApplicantName to be an explicit nil

### UnsetApplicantName
`func (o *DrugRegulatoryRow) UnsetApplicantName()`

UnsetApplicantName ensures that no value is present for ApplicantName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


