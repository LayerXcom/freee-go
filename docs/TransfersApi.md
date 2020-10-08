# \TransfersApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransfer**](TransfersApi.md#CreateTransfer) | **Post** /api/1/transfers | 取引（振替）の作成
[**DestroyTransfer**](TransfersApi.md#DestroyTransfer) | **Delete** /api/1/transfers/{id} | 取引（振替）の削除する
[**GetTransfer**](TransfersApi.md#GetTransfer) | **Get** /api/1/transfers/{id} | 取引（振替）の取得
[**GetTransfers**](TransfersApi.md#GetTransfers) | **Get** /api/1/transfers | 取引（振替）一覧の取得
[**UpdateTransfer**](TransfersApi.md#UpdateTransfer) | **Put** /api/1/transfers/{id} | 取引（振替）の更新



## CreateTransfer

> TransferResponse CreateTransfer(ctx, optional)

取引（振替）の作成

 <h2 id=\"\">概要</h2>  <p>指定した事業所の取引（振替）を作成する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li> <p>amount : 振替金額</p> </li>  <li> <p>from_walletable_type, to_walletable_type</p>  <ul> <li>bank_account : 銀行口座</li>  <li>credit_card : クレジットカード</li>  <li>wallet : その他の決済口座</li> </ul> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTransferOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferParams** | [**optional.Interface of TransferParams**](TransferParams.md)| 取引（振替）の作成 | 

### Return type

[**TransferResponse**](transferResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTransfer

> DestroyTransfer(ctx, id, companyId)

取引（振替）の削除する

 <h2 id=\"\">概要</h2>  <p>指定した事業所の取引（振替）を削除する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 取引(振替)ID | 
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


## GetTransfer

> TransferResponse GetTransfer(ctx, id, companyId)

取引（振替）の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の取引（振替）を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li> <p>amount : 振替金額</p> </li>  <li> <p>from_walletable_type, to_walletable_type</p>  <ul> <li>bank_account : 銀行口座</li>  <li>credit_card : クレジットカード</li>  <li>wallet : その他の決済口座</li> </ul> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 取引(振替)ID | 
**companyId** | **int32**| 事業所ID | 

### Return type

[**TransferResponse**](transferResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransfers

> InlineResponse20014 GetTransfers(ctx, companyId, optional)

取引（振替）一覧の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の取引（振替）一覧を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li> <p>amount : 振替金額</p> </li>  <li> <p>from_walletable_type, to_walletable_type</p>  <ul> <li>bank_account : 銀行口座</li>  <li>credit_card : クレジットカード</li>  <li>wallet : その他の決済口座</li> </ul> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetTransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTransfersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **optional.String**| 振替日で絞込：開始日 (yyyy-mm-dd) | 
 **endDate** | **optional.String**| 振替日で絞込：終了日 (yyyy-mm-dd) | 
 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 100)  | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransfer

> TransferResponse UpdateTransfer(ctx, id, transferParams)

取引（振替）の更新

 <h2 id=\"\">概要</h2>  <p>指定した事業所の取引（振替）を更新する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li> <p>amount : 振替金額</p> </li>  <li> <p>from_walletable_type, to_walletable_type</p>  <ul> <li>bank_account : 銀行口座</li>  <li>credit_card : クレジットカード</li>  <li>wallet : その他の決済口座</li> </ul> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 取引(振替)ID | 
**transferParams** | [**TransferParams**](TransferParams.md)| 取引（振替）の更新 | 

### Return type

[**TransferResponse**](transferResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

