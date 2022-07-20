# ListProtoConfigsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | **string** |  | 
**Granularity** | **int32** |  | 
**TriggerDcaSpread** | **int32** |  | 
**BaseWithdrawalSpread** | **int32** |  | 

## Methods

### NewListProtoConfigsInner

`func NewListProtoConfigsInner(pubkey string, granularity int32, triggerDcaSpread int32, baseWithdrawalSpread int32, ) *ListProtoConfigsInner`

NewListProtoConfigsInner instantiates a new ListProtoConfigsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProtoConfigsInnerWithDefaults

`func NewListProtoConfigsInnerWithDefaults() *ListProtoConfigsInner`

NewListProtoConfigsInnerWithDefaults instantiates a new ListProtoConfigsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *ListProtoConfigsInner) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *ListProtoConfigsInner) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *ListProtoConfigsInner) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.


### GetGranularity

`func (o *ListProtoConfigsInner) GetGranularity() int32`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *ListProtoConfigsInner) GetGranularityOk() (*int32, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *ListProtoConfigsInner) SetGranularity(v int32)`

SetGranularity sets Granularity field to given value.


### GetTriggerDcaSpread

`func (o *ListProtoConfigsInner) GetTriggerDcaSpread() int32`

GetTriggerDcaSpread returns the TriggerDcaSpread field if non-nil, zero value otherwise.

### GetTriggerDcaSpreadOk

`func (o *ListProtoConfigsInner) GetTriggerDcaSpreadOk() (*int32, bool)`

GetTriggerDcaSpreadOk returns a tuple with the TriggerDcaSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerDcaSpread

`func (o *ListProtoConfigsInner) SetTriggerDcaSpread(v int32)`

SetTriggerDcaSpread sets TriggerDcaSpread field to given value.


### GetBaseWithdrawalSpread

`func (o *ListProtoConfigsInner) GetBaseWithdrawalSpread() int32`

GetBaseWithdrawalSpread returns the BaseWithdrawalSpread field if non-nil, zero value otherwise.

### GetBaseWithdrawalSpreadOk

`func (o *ListProtoConfigsInner) GetBaseWithdrawalSpreadOk() (*int32, bool)`

GetBaseWithdrawalSpreadOk returns a tuple with the BaseWithdrawalSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseWithdrawalSpread

`func (o *ListProtoConfigsInner) SetBaseWithdrawalSpread(v int32)`

SetBaseWithdrawalSpread sets BaseWithdrawalSpread field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


