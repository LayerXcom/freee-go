# \BanksApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBank**](BanksApi.md#GetBank) | **Get** /api/1/banks/{id} | 連携サービスの取得
[**GetBanks**](BanksApi.md#GetBanks) | **Get** /api/1/banks | 連携サービス一覧の取得



## GetBank

> BankResponse GetBank(ctx, id)

連携サービスの取得

 <h2 id=\"\">概要</h2>  <p>連携しているサービスを取得する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li>type <ul> <li>bank_account : 銀行口座</li>  <li>credit_card : クレジットカード</li>  <li>wallet : その他の決済口座</li> </ul> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 連携サービスID | 

### Return type

[**BankResponse**](bankResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBanks

> InlineResponse20013 GetBanks(ctx, optional)

連携サービス一覧の取得

 <h2 id=\"\">概要</h2>  <p>連携しているサービス一覧を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li>type <ul> <li>bank_account : 銀行口座</li>  <li>credit_card : クレジットカード</li>  <li>wallet : その他の決済口座</li> </ul> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBanksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBanksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 500) | 
 **type_** | **optional.String**| サービス種別 | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

