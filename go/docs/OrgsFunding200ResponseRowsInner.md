# OrgsFunding200ResponseRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventDate** | **string** |  | 
**RoundType** | **string** |  | 
**AmountUsd** | **NullableFloat32** |  | 
**AmountIsDisclosed** | **bool** |  | 
**Source** | **string** |  | 
**FundingEventId** | Pointer to **string** |  | [optional] 
**RoundSubtype** | Pointer to **string** |  | [optional] 
**IsDilutive** | Pointer to **bool** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CapitalTreatment** | Pointer to **string** |  | [optional] 
**CumulativeRaised** | Pointer to **float32** |  | [optional] 

## Methods

### NewOrgsFunding200ResponseRowsInner

`func NewOrgsFunding200ResponseRowsInner(eventDate string, roundType string, amountUsd NullableFloat32, amountIsDisclosed bool, source string, ) *OrgsFunding200ResponseRowsInner`

NewOrgsFunding200ResponseRowsInner instantiates a new OrgsFunding200ResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsFunding200ResponseRowsInnerWithDefaults

`func NewOrgsFunding200ResponseRowsInnerWithDefaults() *OrgsFunding200ResponseRowsInner`

NewOrgsFunding200ResponseRowsInnerWithDefaults instantiates a new OrgsFunding200ResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventDate

`func (o *OrgsFunding200ResponseRowsInner) GetEventDate() string`

GetEventDate returns the EventDate field if non-nil, zero value otherwise.

### GetEventDateOk

`func (o *OrgsFunding200ResponseRowsInner) GetEventDateOk() (*string, bool)`

GetEventDateOk returns a tuple with the EventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDate

`func (o *OrgsFunding200ResponseRowsInner) SetEventDate(v string)`

SetEventDate sets EventDate field to given value.


### GetRoundType

`func (o *OrgsFunding200ResponseRowsInner) GetRoundType() string`

GetRoundType returns the RoundType field if non-nil, zero value otherwise.

### GetRoundTypeOk

`func (o *OrgsFunding200ResponseRowsInner) GetRoundTypeOk() (*string, bool)`

GetRoundTypeOk returns a tuple with the RoundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundType

`func (o *OrgsFunding200ResponseRowsInner) SetRoundType(v string)`

SetRoundType sets RoundType field to given value.


### GetAmountUsd

`func (o *OrgsFunding200ResponseRowsInner) GetAmountUsd() float32`

GetAmountUsd returns the AmountUsd field if non-nil, zero value otherwise.

### GetAmountUsdOk

`func (o *OrgsFunding200ResponseRowsInner) GetAmountUsdOk() (*float32, bool)`

GetAmountUsdOk returns a tuple with the AmountUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountUsd

`func (o *OrgsFunding200ResponseRowsInner) SetAmountUsd(v float32)`

SetAmountUsd sets AmountUsd field to given value.


### SetAmountUsdNil

`func (o *OrgsFunding200ResponseRowsInner) SetAmountUsdNil(b bool)`

 SetAmountUsdNil sets the value for AmountUsd to be an explicit nil

### UnsetAmountUsd
`func (o *OrgsFunding200ResponseRowsInner) UnsetAmountUsd()`

UnsetAmountUsd ensures that no value is present for AmountUsd, not even an explicit nil
### GetAmountIsDisclosed

`func (o *OrgsFunding200ResponseRowsInner) GetAmountIsDisclosed() bool`

GetAmountIsDisclosed returns the AmountIsDisclosed field if non-nil, zero value otherwise.

### GetAmountIsDisclosedOk

`func (o *OrgsFunding200ResponseRowsInner) GetAmountIsDisclosedOk() (*bool, bool)`

GetAmountIsDisclosedOk returns a tuple with the AmountIsDisclosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountIsDisclosed

`func (o *OrgsFunding200ResponseRowsInner) SetAmountIsDisclosed(v bool)`

SetAmountIsDisclosed sets AmountIsDisclosed field to given value.


### GetSource

`func (o *OrgsFunding200ResponseRowsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OrgsFunding200ResponseRowsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OrgsFunding200ResponseRowsInner) SetSource(v string)`

SetSource sets Source field to given value.


### GetFundingEventId

`func (o *OrgsFunding200ResponseRowsInner) GetFundingEventId() string`

GetFundingEventId returns the FundingEventId field if non-nil, zero value otherwise.

### GetFundingEventIdOk

`func (o *OrgsFunding200ResponseRowsInner) GetFundingEventIdOk() (*string, bool)`

GetFundingEventIdOk returns a tuple with the FundingEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingEventId

`func (o *OrgsFunding200ResponseRowsInner) SetFundingEventId(v string)`

SetFundingEventId sets FundingEventId field to given value.

### HasFundingEventId

`func (o *OrgsFunding200ResponseRowsInner) HasFundingEventId() bool`

HasFundingEventId returns a boolean if a field has been set.

### GetRoundSubtype

`func (o *OrgsFunding200ResponseRowsInner) GetRoundSubtype() string`

GetRoundSubtype returns the RoundSubtype field if non-nil, zero value otherwise.

### GetRoundSubtypeOk

`func (o *OrgsFunding200ResponseRowsInner) GetRoundSubtypeOk() (*string, bool)`

GetRoundSubtypeOk returns a tuple with the RoundSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundSubtype

`func (o *OrgsFunding200ResponseRowsInner) SetRoundSubtype(v string)`

SetRoundSubtype sets RoundSubtype field to given value.

### HasRoundSubtype

`func (o *OrgsFunding200ResponseRowsInner) HasRoundSubtype() bool`

HasRoundSubtype returns a boolean if a field has been set.

### GetIsDilutive

`func (o *OrgsFunding200ResponseRowsInner) GetIsDilutive() bool`

GetIsDilutive returns the IsDilutive field if non-nil, zero value otherwise.

### GetIsDilutiveOk

`func (o *OrgsFunding200ResponseRowsInner) GetIsDilutiveOk() (*bool, bool)`

GetIsDilutiveOk returns a tuple with the IsDilutive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDilutive

`func (o *OrgsFunding200ResponseRowsInner) SetIsDilutive(v bool)`

SetIsDilutive sets IsDilutive field to given value.

### HasIsDilutive

`func (o *OrgsFunding200ResponseRowsInner) HasIsDilutive() bool`

HasIsDilutive returns a boolean if a field has been set.

### GetCurrency

`func (o *OrgsFunding200ResponseRowsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrgsFunding200ResponseRowsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrgsFunding200ResponseRowsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrgsFunding200ResponseRowsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCapitalTreatment

`func (o *OrgsFunding200ResponseRowsInner) GetCapitalTreatment() string`

GetCapitalTreatment returns the CapitalTreatment field if non-nil, zero value otherwise.

### GetCapitalTreatmentOk

`func (o *OrgsFunding200ResponseRowsInner) GetCapitalTreatmentOk() (*string, bool)`

GetCapitalTreatmentOk returns a tuple with the CapitalTreatment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalTreatment

`func (o *OrgsFunding200ResponseRowsInner) SetCapitalTreatment(v string)`

SetCapitalTreatment sets CapitalTreatment field to given value.

### HasCapitalTreatment

`func (o *OrgsFunding200ResponseRowsInner) HasCapitalTreatment() bool`

HasCapitalTreatment returns a boolean if a field has been set.

### GetCumulativeRaised

`func (o *OrgsFunding200ResponseRowsInner) GetCumulativeRaised() float32`

GetCumulativeRaised returns the CumulativeRaised field if non-nil, zero value otherwise.

### GetCumulativeRaisedOk

`func (o *OrgsFunding200ResponseRowsInner) GetCumulativeRaisedOk() (*float32, bool)`

GetCumulativeRaisedOk returns a tuple with the CumulativeRaised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeRaised

`func (o *OrgsFunding200ResponseRowsInner) SetCumulativeRaised(v float32)`

SetCumulativeRaised sets CumulativeRaised field to given value.

### HasCumulativeRaised

`func (o *OrgsFunding200ResponseRowsInner) HasCumulativeRaised() bool`

HasCumulativeRaised returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


