# ActiveWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** |  | 
**PositionCount** | **int32** |  | 

## Methods

### NewActiveWallet

`func NewActiveWallet(owner string, positionCount int32, ) *ActiveWallet`

NewActiveWallet instantiates a new ActiveWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveWalletWithDefaults

`func NewActiveWalletWithDefaults() *ActiveWallet`

NewActiveWalletWithDefaults instantiates a new ActiveWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *ActiveWallet) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ActiveWallet) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ActiveWallet) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetPositionCount

`func (o *ActiveWallet) GetPositionCount() int32`

GetPositionCount returns the PositionCount field if non-nil, zero value otherwise.

### GetPositionCountOk

`func (o *ActiveWallet) GetPositionCountOk() (*int32, bool)`

GetPositionCountOk returns a tuple with the PositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionCount

`func (o *ActiveWallet) SetPositionCount(v int32)`

SetPositionCount sets PositionCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


