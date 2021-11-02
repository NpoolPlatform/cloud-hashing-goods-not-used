# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/cloud-hashing-goods.proto](#npool/cloud-hashing-goods.proto)
    - [AuthorizeAppGoodRequest](#cloud.hashing.goods.v1.AuthorizeAppGoodRequest)
    - [AuthorizeAppGoodResponse](#cloud.hashing.goods.v1.AuthorizeAppGoodResponse)
    - [AuthorizeAppGoodTargetAreaRequest](#cloud.hashing.goods.v1.AuthorizeAppGoodTargetAreaRequest)
    - [AuthorizeAppTargetAreaRequest](#cloud.hashing.goods.v1.AuthorizeAppTargetAreaRequest)
    - [CreateDeviceInfoRequest](#cloud.hashing.goods.v1.CreateDeviceInfoRequest)
    - [CreateDeviceInfoResponse](#cloud.hashing.goods.v1.CreateDeviceInfoResponse)
    - [CreateGoodCommentRequest](#cloud.hashing.goods.v1.CreateGoodCommentRequest)
    - [CreateGoodCommentResponse](#cloud.hashing.goods.v1.CreateGoodCommentResponse)
    - [CreateGoodRequest](#cloud.hashing.goods.v1.CreateGoodRequest)
    - [CreateGoodResponse](#cloud.hashing.goods.v1.CreateGoodResponse)
    - [CreateTargetAreaRequest](#cloud.hashing.goods.v1.CreateTargetAreaRequest)
    - [CreateTargetAreaResponse](#cloud.hashing.goods.v1.CreateTargetAreaResponse)
    - [CreateVendorLocationRequest](#cloud.hashing.goods.v1.CreateVendorLocationRequest)
    - [CreateVendorLocationResponse](#cloud.hashing.goods.v1.CreateVendorLocationResponse)
    - [DeleteDeviceInfoRequest](#cloud.hashing.goods.v1.DeleteDeviceInfoRequest)
    - [DeleteDeviceInfoResponse](#cloud.hashing.goods.v1.DeleteDeviceInfoResponse)
    - [DeleteTargetAreaByContinentCountryRequest](#cloud.hashing.goods.v1.DeleteTargetAreaByContinentCountryRequest)
    - [DeleteTargetAreaByContinentCountryResponse](#cloud.hashing.goods.v1.DeleteTargetAreaByContinentCountryResponse)
    - [DeleteTargetAreaRequest](#cloud.hashing.goods.v1.DeleteTargetAreaRequest)
    - [DeleteTargetAreaResponse](#cloud.hashing.goods.v1.DeleteTargetAreaResponse)
    - [DeleteVendorLocationRequest](#cloud.hashing.goods.v1.DeleteVendorLocationRequest)
    - [DeleteVendorLocationResponse](#cloud.hashing.goods.v1.DeleteVendorLocationResponse)
    - [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo)
    - [GetAllGoodsRequest](#cloud.hashing.goods.v1.GetAllGoodsRequest)
    - [GetAllGoodsResponse](#cloud.hashing.goods.v1.GetAllGoodsResponse)
    - [GetDeviceInfoRequest](#cloud.hashing.goods.v1.GetDeviceInfoRequest)
    - [GetDeviceInfoResponse](#cloud.hashing.goods.v1.GetDeviceInfoResponse)
    - [GetDeviceInfosRequest](#cloud.hashing.goods.v1.GetDeviceInfosRequest)
    - [GetDeviceInfosResponse](#cloud.hashing.goods.v1.GetDeviceInfosResponse)
    - [GetGoodCommentsRequest](#cloud.hashing.goods.v1.GetGoodCommentsRequest)
    - [GetGoodCommentsResponse](#cloud.hashing.goods.v1.GetGoodCommentsResponse)
    - [GetGoodsRequest](#cloud.hashing.goods.v1.GetGoodsRequest)
    - [GetGoodsResponse](#cloud.hashing.goods.v1.GetGoodsResponse)
    - [GetTargetAreaRequest](#cloud.hashing.goods.v1.GetTargetAreaRequest)
    - [GetTargetAreaResponse](#cloud.hashing.goods.v1.GetTargetAreaResponse)
    - [GetTargetAreasRequest](#cloud.hashing.goods.v1.GetTargetAreasRequest)
    - [GetTargetAreasResponse](#cloud.hashing.goods.v1.GetTargetAreasResponse)
    - [GetVendorLocationRequest](#cloud.hashing.goods.v1.GetVendorLocationRequest)
    - [GetVendorLocationResponse](#cloud.hashing.goods.v1.GetVendorLocationResponse)
    - [GetVendorLocationsRequest](#cloud.hashing.goods.v1.GetVendorLocationsRequest)
    - [GetVendorLocationsResponse](#cloud.hashing.goods.v1.GetVendorLocationsResponse)
    - [GoodAuthInfo](#cloud.hashing.goods.v1.GoodAuthInfo)
    - [GoodComment](#cloud.hashing.goods.v1.GoodComment)
    - [GoodInfo](#cloud.hashing.goods.v1.GoodInfo)
    - [GoodPrice](#cloud.hashing.goods.v1.GoodPrice)
    - [PageInfo](#cloud.hashing.goods.v1.PageInfo)
    - [SetAppGoodPriceRequest](#cloud.hashing.goods.v1.SetAppGoodPriceRequest)
    - [SetAppGoodPriceResponse](#cloud.hashing.goods.v1.SetAppGoodPriceResponse)
    - [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo)
    - [UnauthorizeAppGoodRequest](#cloud.hashing.goods.v1.UnauthorizeAppGoodRequest)
    - [UnauthorizeAppGoodResponse](#cloud.hashing.goods.v1.UnauthorizeAppGoodResponse)
    - [UnauthorizeAppGoodTargetAreaRequest](#cloud.hashing.goods.v1.UnauthorizeAppGoodTargetAreaRequest)
    - [UnauthorizeAppTargetAreaRequest](#cloud.hashing.goods.v1.UnauthorizeAppTargetAreaRequest)
    - [UpdateDeviceInfoRequest](#cloud.hashing.goods.v1.UpdateDeviceInfoRequest)
    - [UpdateDeviceInfoResponse](#cloud.hashing.goods.v1.UpdateDeviceInfoResponse)
    - [UpdateGoodCommentRequest](#cloud.hashing.goods.v1.UpdateGoodCommentRequest)
    - [UpdateGoodCommentResponse](#cloud.hashing.goods.v1.UpdateGoodCommentResponse)
    - [UpdateGoodRequest](#cloud.hashing.goods.v1.UpdateGoodRequest)
    - [UpdateGoodResponse](#cloud.hashing.goods.v1.UpdateGoodResponse)
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



<a name="cloud.hashing.goods.v1.AuthorizeAppGoodRequest"></a>

### AuthorizeAppGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| AuthorizeAllTargetArea | [bool](#bool) |  |  |






<a name="cloud.hashing.goods.v1.AuthorizeAppGoodResponse"></a>

### AuthorizeAppGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodAuthInfo](#cloud.hashing.goods.v1.GoodAuthInfo) |  |  |






<a name="cloud.hashing.goods.v1.AuthorizeAppGoodTargetAreaRequest"></a>

### AuthorizeAppGoodTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| TargetAreaID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.AuthorizeAppTargetAreaRequest"></a>

### AuthorizeAppTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| TargetAreaID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.CreateDeviceInfoRequest"></a>

### CreateDeviceInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateDeviceInfoResponse"></a>

### CreateDeviceInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateGoodCommentRequest"></a>

### CreateGoodCommentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Comment | [GoodComment](#cloud.hashing.goods.v1.GoodComment) |  |  |






<a name="cloud.hashing.goods.v1.CreateGoodCommentResponse"></a>

### CreateGoodCommentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Comment | [GoodComment](#cloud.hashing.goods.v1.GoodComment) |  |  |






<a name="cloud.hashing.goods.v1.CreateGoodRequest"></a>

### CreateGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateGoodResponse"></a>

### CreateGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) |  |  |






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






<a name="cloud.hashing.goods.v1.DeleteDeviceInfoRequest"></a>

### DeleteDeviceInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.DeleteDeviceInfoResponse"></a>

### DeleteDeviceInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) |  |  |






<a name="cloud.hashing.goods.v1.DeleteTargetAreaByContinentCountryRequest"></a>

### DeleteTargetAreaByContinentCountryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Continent | [string](#string) |  |  |
| Country | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.DeleteTargetAreaByContinentCountryResponse"></a>

### DeleteTargetAreaByContinentCountryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.DeleteTargetAreaRequest"></a>

### DeleteTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.DeleteTargetAreaResponse"></a>

### DeleteTargetAreaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.DeleteVendorLocationRequest"></a>

### DeleteVendorLocationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.DeleteVendorLocationResponse"></a>

### DeleteVendorLocationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) |  |  |






<a name="cloud.hashing.goods.v1.DeviceInfo"></a>

### DeviceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Type | [string](#string) |  |  |
| Manufacturer | [string](#string) |  |  |
| PowerComsuption | [int32](#int32) |  |  |
| ShipmentAt | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.GetAllGoodsRequest"></a>

### GetAllGoodsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetAllGoodsResponse"></a>

### GetAllGoodsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GoodAuthInfo](#cloud.hashing.goods.v1.GoodAuthInfo) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.GetDeviceInfoRequest"></a>

### GetDeviceInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetDeviceInfoResponse"></a>

### GetDeviceInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetDeviceInfosRequest"></a>

### GetDeviceInfosRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetDeviceInfosResponse"></a>

### GetDeviceInfosResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodCommentsRequest"></a>

### GetGoodCommentsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodCommentsResponse"></a>

### GetGoodCommentsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Comments | [GoodComment](#cloud.hashing.goods.v1.GoodComment) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodsRequest"></a>

### GetGoodsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodsResponse"></a>

### GetGoodsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.GetTargetAreaRequest"></a>

### GetTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetTargetAreaResponse"></a>

### GetTargetAreaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetTargetAreasRequest"></a>

### GetTargetAreasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetTargetAreasResponse"></a>

### GetTargetAreasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.GetVendorLocationRequest"></a>

### GetVendorLocationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetVendorLocationResponse"></a>

### GetVendorLocationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetVendorLocationsRequest"></a>

### GetVendorLocationsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetVendorLocationsResponse"></a>

### GetVendorLocationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.GoodAuthInfo"></a>

### GoodAuthInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) |  |  |
| AppAuthorized | [bool](#bool) |  |  |






<a name="cloud.hashing.goods.v1.GoodComment"></a>

### GoodComment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| ReplyTo | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| Content | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GoodInfo"></a>

### GoodInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| DeviceInfoID | [string](#string) |  |  |
| GasPrice | [int64](#int64) |  |  |
| SeparateGasFee | [bool](#bool) |  |  |
| UnitPower | [int32](#int32) |  |  |
| DurationDays | [int32](#int32) |  |  |
| CoinInfoID | [string](#string) |  |  |
| Actuals | [bool](#bool) |  |  |
| DeliveryAt | [int32](#int32) |  |  |
| InheritFromGoodID | [string](#string) |  |  |
| VendorLocationID | [string](#string) |  |  |
| Price | [int64](#int64) |  |  |
| BenefitType | [string](#string) |  |  |
| Classic | [bool](#bool) |  |  |
| SupportCoinTypeIDs | [string](#string) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.GoodPrice"></a>

### GoodPrice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Price | [float](#float) |  |  |
| GasPrice | [float](#float) |  |  |






<a name="cloud.hashing.goods.v1.PageInfo"></a>

### PageInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageIndex | [int32](#int32) |  |  |
| PageSize | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.SetAppGoodPriceRequest"></a>

### SetAppGoodPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |
| GoodPrice | [GoodPrice](#cloud.hashing.goods.v1.GoodPrice) |  |  |






<a name="cloud.hashing.goods.v1.SetAppGoodPriceResponse"></a>

### SetAppGoodPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodPrice | [GoodPrice](#cloud.hashing.goods.v1.GoodPrice) |  |  |
| GoodInfo | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.TargetAreaInfo"></a>

### TargetAreaInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Continent | [string](#string) |  |  |
| Country | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.UnauthorizeAppGoodRequest"></a>

### UnauthorizeAppGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.UnauthorizeAppGoodResponse"></a>

### UnauthorizeAppGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodAuthInfo](#cloud.hashing.goods.v1.GoodAuthInfo) |  |  |






<a name="cloud.hashing.goods.v1.UnauthorizeAppGoodTargetAreaRequest"></a>

### UnauthorizeAppGoodTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| TargetAreaID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.UnauthorizeAppTargetAreaRequest"></a>

### UnauthorizeAppTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| TargetAreaID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.UpdateDeviceInfoRequest"></a>

### UpdateDeviceInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateDeviceInfoResponse"></a>

### UpdateDeviceInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateGoodCommentRequest"></a>

### UpdateGoodCommentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Comment | [GoodComment](#cloud.hashing.goods.v1.GoodComment) |  |  |






<a name="cloud.hashing.goods.v1.UpdateGoodCommentResponse"></a>

### UpdateGoodCommentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Comment | [GoodComment](#cloud.hashing.goods.v1.GoodComment) |  |  |






<a name="cloud.hashing.goods.v1.UpdateGoodRequest"></a>

### UpdateGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateGoodResponse"></a>

### UpdateGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) |  |  |






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
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#cloud.hashing.goods.v1.VersionResponse) |  |
| CreateVendorLocation | [CreateVendorLocationRequest](#cloud.hashing.goods.v1.CreateVendorLocationRequest) | [CreateVendorLocationResponse](#cloud.hashing.goods.v1.CreateVendorLocationResponse) |  |
| UpdateVendorLocation | [UpdateVendorLocationRequest](#cloud.hashing.goods.v1.UpdateVendorLocationRequest) | [UpdateVendorLocationResponse](#cloud.hashing.goods.v1.UpdateVendorLocationResponse) |  |
| GetVendorLocation | [GetVendorLocationRequest](#cloud.hashing.goods.v1.GetVendorLocationRequest) | [GetVendorLocationResponse](#cloud.hashing.goods.v1.GetVendorLocationResponse) |  |
| DeleteVendorLocation | [DeleteVendorLocationRequest](#cloud.hashing.goods.v1.DeleteVendorLocationRequest) | [DeleteVendorLocationResponse](#cloud.hashing.goods.v1.DeleteVendorLocationResponse) |  |
| GetVendorLocations | [GetVendorLocationsRequest](#cloud.hashing.goods.v1.GetVendorLocationsRequest) | [GetVendorLocationsResponse](#cloud.hashing.goods.v1.GetVendorLocationsResponse) |  |
| CreateTargetArea | [CreateTargetAreaRequest](#cloud.hashing.goods.v1.CreateTargetAreaRequest) | [CreateTargetAreaResponse](#cloud.hashing.goods.v1.CreateTargetAreaResponse) |  |
| UpdateTargetArea | [UpdateTargetAreaRequest](#cloud.hashing.goods.v1.UpdateTargetAreaRequest) | [UpdateTargetAreaResponse](#cloud.hashing.goods.v1.UpdateTargetAreaResponse) |  |
| GetTargetArea | [GetTargetAreaRequest](#cloud.hashing.goods.v1.GetTargetAreaRequest) | [GetTargetAreaResponse](#cloud.hashing.goods.v1.GetTargetAreaResponse) |  |
| DeleteTargetArea | [DeleteTargetAreaRequest](#cloud.hashing.goods.v1.DeleteTargetAreaRequest) | [DeleteTargetAreaResponse](#cloud.hashing.goods.v1.DeleteTargetAreaResponse) |  |
| DeleteTargetAreaByContinentCountry | [DeleteTargetAreaByContinentCountryRequest](#cloud.hashing.goods.v1.DeleteTargetAreaByContinentCountryRequest) | [DeleteTargetAreaByContinentCountryResponse](#cloud.hashing.goods.v1.DeleteTargetAreaByContinentCountryResponse) |  |
| GetTargetAreas | [GetTargetAreasRequest](#cloud.hashing.goods.v1.GetTargetAreasRequest) | [GetTargetAreasResponse](#cloud.hashing.goods.v1.GetTargetAreasResponse) |  |
| CreateDeviceInfo | [CreateDeviceInfoRequest](#cloud.hashing.goods.v1.CreateDeviceInfoRequest) | [CreateDeviceInfoResponse](#cloud.hashing.goods.v1.CreateDeviceInfoResponse) |  |
| UpdateDeviceInfo | [UpdateDeviceInfoRequest](#cloud.hashing.goods.v1.UpdateDeviceInfoRequest) | [UpdateDeviceInfoResponse](#cloud.hashing.goods.v1.UpdateDeviceInfoResponse) |  |
| GetDeviceInfo | [GetDeviceInfoRequest](#cloud.hashing.goods.v1.GetDeviceInfoRequest) | [GetDeviceInfoResponse](#cloud.hashing.goods.v1.GetDeviceInfoResponse) |  |
| DeleteDeviceInfo | [DeleteDeviceInfoRequest](#cloud.hashing.goods.v1.DeleteDeviceInfoRequest) | [DeleteDeviceInfoResponse](#cloud.hashing.goods.v1.DeleteDeviceInfoResponse) |  |
| GetDeviceInfos | [GetDeviceInfosRequest](#cloud.hashing.goods.v1.GetDeviceInfosRequest) | [GetDeviceInfosResponse](#cloud.hashing.goods.v1.GetDeviceInfosResponse) |  |
| CreateGood | [CreateGoodRequest](#cloud.hashing.goods.v1.CreateGoodRequest) | [CreateGoodResponse](#cloud.hashing.goods.v1.CreateGoodResponse) |  |
| UpdateGood | [UpdateGoodRequest](#cloud.hashing.goods.v1.UpdateGoodRequest) | [UpdateGoodResponse](#cloud.hashing.goods.v1.UpdateGoodResponse) |  |
| GetAllGoods | [GetAllGoodsRequest](#cloud.hashing.goods.v1.GetAllGoodsRequest) | [GetAllGoodsResponse](#cloud.hashing.goods.v1.GetAllGoodsResponse) | Can only accessed by APP administrator and platform administrator |
| GetGoods | [GetGoodsRequest](#cloud.hashing.goods.v1.GetGoodsRequest) | [GetGoodsResponse](#cloud.hashing.goods.v1.GetGoodsResponse) | Accessed by APP user |
| SetAppGoodPrice | [SetAppGoodPriceRequest](#cloud.hashing.goods.v1.SetAppGoodPriceRequest) | [SetAppGoodPriceResponse](#cloud.hashing.goods.v1.SetAppGoodPriceResponse) |  |
| AuthorizeAppGood | [AuthorizeAppGoodRequest](#cloud.hashing.goods.v1.AuthorizeAppGoodRequest) | [AuthorizeAppGoodResponse](#cloud.hashing.goods.v1.AuthorizeAppGoodResponse) |  |
| UnauthorizeAppGood | [UnauthorizeAppGoodRequest](#cloud.hashing.goods.v1.UnauthorizeAppGoodRequest) | [UnauthorizeAppGoodResponse](#cloud.hashing.goods.v1.UnauthorizeAppGoodResponse) |  |
| AuthorizeAppTargetArea | [AuthorizeAppTargetAreaRequest](#cloud.hashing.goods.v1.AuthorizeAppTargetAreaRequest) | [.google.protobuf.Empty](#google.protobuf.Empty) |  |
| UnauthorizeAppTargetArea | [UnauthorizeAppTargetAreaRequest](#cloud.hashing.goods.v1.UnauthorizeAppTargetAreaRequest) | [.google.protobuf.Empty](#google.protobuf.Empty) |  |
| AuthorizeAppGoodTargetArea | [AuthorizeAppGoodTargetAreaRequest](#cloud.hashing.goods.v1.AuthorizeAppGoodTargetAreaRequest) | [.google.protobuf.Empty](#google.protobuf.Empty) |  |
| UnauthorizeAppGoodTargetArea | [UnauthorizeAppGoodTargetAreaRequest](#cloud.hashing.goods.v1.UnauthorizeAppGoodTargetAreaRequest) | [.google.protobuf.Empty](#google.protobuf.Empty) |  |
| CreateGoodComment | [CreateGoodCommentRequest](#cloud.hashing.goods.v1.CreateGoodCommentRequest) | [CreateGoodCommentResponse](#cloud.hashing.goods.v1.CreateGoodCommentResponse) |  |
| UpdateGoodComment | [UpdateGoodCommentRequest](#cloud.hashing.goods.v1.UpdateGoodCommentRequest) | [UpdateGoodCommentResponse](#cloud.hashing.goods.v1.UpdateGoodCommentResponse) |  |
| GetGoodComments | [GetGoodCommentsRequest](#cloud.hashing.goods.v1.GetGoodCommentsRequest) | [GetGoodCommentsResponse](#cloud.hashing.goods.v1.GetGoodCommentsResponse) |  |

 



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

