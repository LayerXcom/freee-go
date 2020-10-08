# \WalletablesApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWalletable**](WalletablesApi.md#CreateWalletable) | **Post** /api/1/walletables | 口座の作成
[**DestroyWalletable**](WalletablesApi.md#DestroyWalletable) | **Delete** /api/1/walletables/{type}/{id} | 口座の削除
[**GetWalletable**](WalletablesApi.md#GetWalletable) | **Get** /api/1/walletables/{type}/{id} | 口座情報の取得
[**GetWalletables**](WalletablesApi.md#GetWalletables) | **Get** /api/1/walletables | 口座一覧の取得
[**UpdateWalletable**](WalletablesApi.md#UpdateWalletable) | **Put** /api/1/walletables/{type}/{id} | 口座の更新



## CreateWalletable

> WalletableCreateResponse CreateWalletable(ctx, optional)

口座の作成

 <h2 id=\"\">概要</h2>  <p>指定した事業所に口座を作成する</p>  <h2 id=\"\">注意点</h2> <ul><li>同期に対応した口座はこのAPIでは作成できません</li></ul>  <h2 id=\"_2\">定義</h2>  <ul> <li> <p>type</p>  <ul> <li>bank_account : 銀行口座</li>  <li>credit_card : クレジットカード</li>  <li>wallet : その他の決済口座</li> </ul> </li>  <li> <p>name : 口座名</p> </li>  <li> <p>bank_id : サービスID</p> </li>  <li> <p>group_name : 決算書表示名（小カテゴリー）　例：売掛金, 受取手形, 未収入金（法人のみ）, 買掛金, 支払手形, 未払金, 預り金, 前受金</p> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateWalletableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWalletableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletableCreateParams** | [**optional.Interface of WalletableCreateParams**](WalletableCreateParams.md)| 口座の作成 | 

### Return type

[**WalletableCreateResponse**](walletableCreateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyWalletable

> DestroyWalletable(ctx, id, type_, companyId)

口座の削除

 <h2 id=\"\">概要</h2>  <p>指定した事業所の口座を削除する</p>  <h2 id=\"\">注意点</h2> <ul> <li>削除を実行するには、当該口座に関連する仕訳データを事前に削除する必要があります。</li> <li>当該口座に仕訳が残っていないか確認するには、レポートの「仕訳帳」等を参照し、必要に応じて、「取引」や「口座振替」も削除します。</li>  </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 口座ID | 
**type_** | **string**| 口座種別（bank_account : 銀行口座, credit_card : クレジットカード, wallet : その他の決済口座） | 
**companyId** | **int32**| 事業所ID | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletable

> InlineResponse20012 GetWalletable(ctx, id, type_, companyId)

口座情報の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の口座情報を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li>type <ul> <li>bank_account : 銀行口座</li>  <li>credit_card : クレジットカード</li>  <li>wallet : その他の決済口座</li> </ul> </li>  <li>walletable_balance : 登録残高</li>  <li>last_balance : 同期残高</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 口座ID | 
**type_** | **string**| 口座種別（bank_account : 銀行口座, credit_card : クレジットカード, wallet : その他の決済口座） | 
**companyId** | **int32**| 事業所ID | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletables

> InlineResponse20011 GetWalletables(ctx, companyId, optional)

口座一覧の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の口座一覧を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li>type <ul> <li>bank_account : 銀行口座</li>  <li>credit_card : クレジットカード</li>  <li>wallet : その他の決済口座</li> </ul> </li>  <li>walletable_balance : 登録残高</li>  <li>last_balance : 同期残高</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetWalletablesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWalletablesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withBalance** | **optional.Bool**| 残高情報を含める | 
 **type_** | **optional.String**| 口座種別（bank_account : 銀行口座, credit_card : クレジットカード, wallet : その他の決済口座） | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWalletable

> InlineResponse20012 UpdateWalletable(ctx, id, type_, optional)

口座の更新

 <h2 id=\"\">概要</h2>  <p>口座を更新する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**type_** | **string**| 口座種別（bank_account : 銀行口座, credit_card : クレジットカード, wallet : その他の決済口座） | 
 **optional** | ***UpdateWalletableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWalletableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **walletableUpdateParams** | [**optional.Interface of WalletableUpdateParams**](WalletableUpdateParams.md)| 口座の作成 | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

