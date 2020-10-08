# \PaymentsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDealPayment**](PaymentsApi.md#CreateDealPayment) | **Post** /api/1/deals/{id}/payments | 取引（収入／支出）の支払行作成
[**DestroyDealPayment**](PaymentsApi.md#DestroyDealPayment) | **Delete** /api/1/deals/{id}/payments/{payment_id} | 取引（収入／支出）の支払行削除
[**UpdateDealPayment**](PaymentsApi.md#UpdateDealPayment) | **Put** /api/1/deals/{id}/payments/{payment_id} | 取引（収入／支出）の支払行更新



## CreateDealPayment

> DealResponse CreateDealPayment(ctx, id, paymentParams)

取引（収入／支出）の支払行作成

<h2 id=\"\">概要</h2> <p>指定した事業所の取引（収入／支出）の支払行を作成する</p> <h2 id=\"_2\">定義</h2> <ul> <li> <p>issue_date : 発生日</p> </li> <li> <p>due_date : 支払期日</p> </li> <li> <p>amount : 金額</p> </li> <li> <p>due_amount : 支払残額</p> </li> <li> <p>type</p> <ul> <li>income : 収入</li> <li>expense : 支出</li> </ul> </li> <li> <p>details : 取引の明細行</p> </li> <li> <p>renews : 取引の+更新行</p> </li> <li> <p>payments : 取引の支払行</p> </li> <li> <p>from_walletable_type</p> <ul> <li>bank_account : 銀行口座</li> <li>credit_card : クレジットカード</li> <li>wallet : 現金</li> <li>private_account_item : プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）</li> </ul> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 取引ID | 
**paymentParams** | [**PaymentParams**](PaymentParams.md)| 取引（収入／支出）の支払行作成 | 

### Return type

[**DealResponse**](dealResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyDealPayment

> DestroyDealPayment(ctx, id, paymentId, companyId)

取引（収入／支出）の支払行削除

<h2 id=\"\">概要</h2> <p>指定した事業所の取引（収入／支出）の支払行を削除する</p> <h2 id=\"_2\">定義</h2> <ul> <li> <p>issue_date : 発生日</p> </li> <li> <p>due_date : 支払期日</p> </li> <li> <p>amount : 金額</p> </li> <li> <p>due_amount : 支払残額</p> </li> <li> <p>type</p> <ul> <li>income : 収入</li> <li>expense : 支出</li> </ul> </li> <li> <p>details : 取引の明細行</p> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 取引ID | 
**paymentId** | **int32**| 決済ID | 
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


## UpdateDealPayment

> DealResponse UpdateDealPayment(ctx, id, paymentId, paymentParams)

取引（収入／支出）の支払行更新

<h2 id=\"\">概要</h2> <p>指定した事業所の取引（収入／支出）の支払行を更新する</p> <h2 id=\"_2\">定義</h2> <ul> <li> <p>issue_date : 発生日</p> </li> <li> <p>due_date : 支払期日</p> </li> <li> <p>amount : 金額</p> </li> <li> <p>due_amount : 支払残額</p> </li> <li> <p>type</p> <ul> <li>income : 収入</li> <li>expense : 支出</li> </ul> </li> <li> <p>details : 取引の明細行</p> </li> <li> <p>renews : 取引の+更新行</p> </li> <li> <p>payments : 取引の支払行</p> </li> <li> <p>from_walletable_type</p> <ul> <li>bank_account : 銀行口座</li> <li>credit_card : クレジットカード</li> <li>wallet : 現金</li> <li>private_account_item : プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）</li> </ul> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 取引ID | 
**paymentId** | **int32**| 決済ID | 
**paymentParams** | [**PaymentParams**](PaymentParams.md)| 取引（収入／支出）の支払行更新 | 

### Return type

[**DealResponse**](dealResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

