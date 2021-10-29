# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/cloud-hashing-goods.proto](#npool/cloud-hashing-goods.proto)
    - [CreateTargetAreaRequest](#cloud.hashing.goods.v1.CreateTargetAreaRequest)
    - [CreateTargetAreaResponse](#cloud.hashing.goods.v1.CreateTargetAreaResponse)
    - [CreateVendorLocationRequest](#cloud.hashing.goods.v1.CreateVendorLocationRequest)
    - [CreateVendorLocationResponse](#cloud.hashing.goods.v1.CreateVendorLocationResponse)
    - [GetTargetAreasResponse](#cloud.hashing.goods.v1.GetTargetAreasResponse)
    - [GetVendorLocationsResponse](#cloud.hashing.goods.v1.GetVendorLocationsResponse)
    - [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo)
    - [UpdateTargetAreaRequest](#cloud.hashing.goods.v1.UpdateTargetAreaRequest)
    - [UpdateTargetAreaResponse](#cloud.hashing.goods.v1.UpdateTargetAreaResponse)
    - [UpdateVendorLocationRequest](#cloud.hashing.goods.v1.UpdateVendorLocationRequest)
    - [UpdateVendorLocationResponse](#cloud.hashing.goods.v1.UpdateVendorLocationResponse)
    - [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo)
    - [VersionResponse](#cloud.hashing.goods.v1.VersionResponse)
  
    - [CloudHashingGoods](#cloud.hashing.goods.v1.CloudHashingGoods)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/cloud-hashing-goods.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/cloud-hashing-goods.proto



<a name="cloud.hashing.goods.v1.CreateTargetAreaRequest"></a>

### CreateTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateTargetAreaResponse"></a>

### CreateTargetAreaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateVendorLocationRequest"></a>

### CreateVendorLocationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateVendorLocationResponse"></a>

### CreateVendorLocationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetTargetAreasResponse"></a>

### GetTargetAreasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) | repeated |  |






<a name="cloud.hashing.goods.v1.GetVendorLocationsResponse"></a>

### GetVendorLocationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) | repeated |  |






<a name="cloud.hashing.goods.v1.TargetAreaInfo"></a>

### TargetAreaInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Continent | [string](#string) |  |  |
| Country | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.UpdateTargetAreaRequest"></a>

### UpdateTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateTargetAreaResponse"></a>

### UpdateTargetAreaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateVendorLocationRequest"></a>

### UpdateVendorLocationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateVendorLocationResponse"></a>

### UpdateVendorLocationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) |  |  |






<a name="cloud.hashing.goods.v1.VendorLocationInfo"></a>

### VendorLocationInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Country | [string](#string) |  |  |
| Province | [string](#string) |  |  |
| City | [string](#string) |  |  |
| Address | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.VersionResponse"></a>

### VersionResponse
Request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="cloud.hashing.goods.v1.CloudHashingGoods"></a>

### CloudHashingGoods
Cloud Hashing Goods

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#cloud.hashing.goods.v1.VersionResponse) | Method Version |
| CreateVendorLocation | [CreateVendorLocationRequest](#cloud.hashing.goods.v1.CreateVendorLocationRequest) | [CreateVendorLocationResponse](#cloud.hashing.goods.v1.CreateVendorLocationResponse) |  |
| UpdateVendorLocation | [UpdateVendorLocationRequest](#cloud.hashing.goods.v1.UpdateVendorLocationRequest) | [UpdateVendorLocationResponse](#cloud.hashing.goods.v1.UpdateVendorLocationResponse) |  |
| GetVendorLocations | [.google.protobuf.Empty](#google.protobuf.Empty) | [GetVendorLocationsResponse](#cloud.hashing.goods.v1.GetVendorLocationsResponse) |  |
| CreateTargetArea | [CreateTargetAreaRequest](#cloud.hashing.goods.v1.CreateTargetAreaRequest) | [CreateTargetAreaResponse](#cloud.hashing.goods.v1.CreateTargetAreaResponse) |  |
| UpdateTargetArea | [UpdateTargetAreaRequest](#cloud.hashing.goods.v1.UpdateTargetAreaRequest) | [UpdateTargetAreaResponse](#cloud.hashing.goods.v1.UpdateTargetAreaResponse) |  |
| GetTargetAreas | [.google.protobuf.Empty](#google.protobuf.Empty) | [GetTargetAreasResponse](#cloud.hashing.goods.v1.GetTargetAreasResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

