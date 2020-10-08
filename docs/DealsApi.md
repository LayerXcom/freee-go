# \DealsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeal**](DealsApi.md#CreateDeal) | **Post** /api/1/deals | 取引（収入／支出）の作成
[**DestroyDeal**](DealsApi.md#DestroyDeal) | **Delete** /api/1/deals/{id} | 取引（収入／支出）の削除
[**GetDeal**](DealsApi.md#GetDeal) | **Get** /api/1/deals/{id} | 取引（収入／支出）の取得
[**GetDeals**](DealsApi.md#GetDeals) | **Get** /api/1/deals | 取引（収入／支出）一覧の取得
[**UpdateDeal**](DealsApi.md#UpdateDeal) | **Put** /api/1/deals/{id} | 取引（収入／支出）の更新



## CreateDeal

> DealCreateResponse CreateDeal(ctx, optional)

取引（収入／支出）の作成

<h2 id=\"\">概要</h2> <p>指定した事業所の取引（収入／支出）を作成する</p> <h2 id=\"_2\">定義</h2> <ul> <li> <p>issue_date : 発生日</p> </li> <li> <p>due_date : 支払期日</p> </li> <li> <p>amount : 金額</p> </li> <li> <p>due_amount : 支払残額</p> </li> <li> <p>type</p> <ul> <li>income : 収入</li> <li>expense : 支出</li> </ul> </li> <li> <p>ref_number : 管理番号</p> </li> <li> <p>details : 取引の明細行(最大40行)</p> </li> <li> <p>payments : 取引の支払行</p> </li> <li> <p>receipt_ids : 証憑ファイルID</p> </li> <li> <p>from_walletable_type</p> <ul> <li>bank_account : 銀行口座</li> <li>credit_card : クレジットカード</li> <li>wallet : 現金</li> <li>private_account_item : プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）</li> </ul> </li> </ul> <h2 id=\"_3\">注意点</h2> <ul>     <li><p>本APIでは+更新行(renews)の操作ができません。+更新行の作成APIをご利用ください。</p></li>     <li><p>セグメントタグ情報は法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。</p></li>     <li>         <p>partner_codeを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。またpartner_codeとpartner_idは同時に指定することはできません。</p></li>     <li>         <p>本APIでは取引の明細行(details)は、最大40行までになります。</p>     </li> </ul> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDealOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDealOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dealCreateParams** | [**optional.Interface of DealCreateParams**](DealCreateParams.md)| 取引（収入／支出）の作成 | 

### Return type

[**DealCreateResponse**](dealCreateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyDeal

> DestroyDeal(ctx, id, companyId)

取引（収入／支出）の削除

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 取引ID | 
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


## GetDeal

> DealResponse GetDeal(ctx, companyId, id, optional)

取引（収入／支出）の取得

<h2 id=\"\">概要</h2> <p>指定した事業所の取引（収入／支出）を取得する</p> <h2 id=\"_2\">定義</h2> <ul> <li> <p>issue_date : 発生日</p> </li> <li> <p>due_date : 支払期日</p> </li> <li> <p>amount : 金額</p> </li> <li> <p>due_amount : 支払残額</p> </li> <li> <p>type</p> <ul> <li>income : 収入</li> <li>expense : 支出</li> </ul> </li> <li> <p>details : 取引の明細行</p> </li> <li> <p>accruals : 取引の債権債務行</p> </li> <li> <p>renews : 取引の+更新行</p> </li> <li> <p>payments : 取引の支払行</p> </li> <li> <p>from_walletable_type</p> <ul> <li>bank_account : 銀行口座</li> <li>credit_card : クレジットカード</li> <li>wallet : 現金</li> <li>private_account_item : プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）</li> </ul> </li> <li> <p>registered_from</p> <ul> <li>all : すべての取引</li> <li>me : 自身が登録した取引</li> </ul> </li> </ul> <h2 id=\"_3\">注意点</h2> <ul> <li>セグメントタグ情報は法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
**id** | **int32**|  | 
 **optional** | ***GetDealOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDealOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accruals** | **optional.String**| 取引の債権債務行の表示（without: 表示しない(デフォルト), with: 表示する） | 

### Return type

[**DealResponse**](dealResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeals

> InlineResponse2004 GetDeals(ctx, companyId, optional)

取引（収入／支出）一覧の取得

<h2 id=\"\">概要</h2> <p>指定した事業所の取引一覧（収入／支出）を取得する</p> <h2 id=\"_2\">定義</h2> <ul> <li> <p>issue_date : 発生日</p> </li> <li> <p>due_date : 支払期日</p> </li> <li> <p>amount : 金額</p> </li> <li> <p>due_amount : 支払残額</p> </li> <li> <p>type</p> <ul> <li>income : 収入</li> <li>expense : 支出</li> </ul> </li> <li> <p>details : 取引の明細行</p> </li> <li> <p>accruals : 取引の債権債務行</p> </li> <li> <p>renews : 取引の+更新行</p> </li> <li> <p>payments : 取引の支払行</p> </li> <li> <p>from_walletable_type</p> <ul> <li>bank_account : 銀行口座</li> <li>credit_card : クレジットカード</li> <li>wallet : 現金</li> <li>private_account_item : プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）</li> </ul> </li> <li> <p>registered_from</p> <ul> <li>all : すべての取引</li> <li>me : 自身が登録した取引</li> </ul> </li> </ul> <h2 id=\"_3\">注意点</h2> <ul> <li>セグメントタグ情報は法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。</li> <li>partner_codeを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。またpartner_codeとpartner_idは同時に指定することはできません。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetDealsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDealsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partnerId** | **optional.Int32**| 取引先IDで絞込 | 
 **accountItemId** | **optional.Int32**| 勘定科目IDで絞込 | 
 **partnerCode** | **optional.String**| 取引先コードで絞込 | 
 **status** | **optional.String**| 決済状況で絞込 (未決済: unsettled, 完了: settled) | 
 **type_** | **optional.String**| 収支区分 (収入: income, 支出: expense) | 
 **startIssueDate** | **optional.String**| 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endIssueDate** | **optional.String**| 発生日で絞込：終了日(yyyy-mm-dd) | 
 **startDueDate** | **optional.String**| 支払期日で絞込：開始日(yyyy-mm-dd) | 
 **endDueDate** | **optional.String**| 支払期日で絞込：終了日(yyyy-mm-dd) | 
 **startRenewDate** | **optional.String**| +更新日で絞込：開始日(yyyy-mm-dd) | 
 **endRenewDate** | **optional.String**| +更新日で絞込：終了日(yyyy-mm-dd) | 
 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 20, 最大: 100)  | 
 **registeredFrom** | **optional.String**| 取引登録元アプリで絞込（me: 本APIを叩くアプリ自身から登録した取引のみ） | 
 **accruals** | **optional.String**| 取引の債権債務行の表示（without: 表示しない(デフォルト), with: 表示する） | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeal

> DealResponse UpdateDeal(ctx, id, optional)

取引（収入／支出）の更新

<h2 id=\"\">概要</h2> <p>指定した事業所の取引（収入／支出）を更新する</p> <h2 id=\"_2\">定義</h2> <ul> <li> <p>issue_date : 発生日</p> </li> <li> <p>due_date : 支払期日</p> </li> <li> <p>amount : 金額</p> </li> <li> <p>due_amount : 支払残額</p> </li> <li> <p>type</p> <ul> <li>income : 収入</li> <li>expense : 支出</li> </ul> </li> <li> <p>details : 取引の明細行(最大40行)</p> </li> <li> <p>renews : 取引の+更新行</p> </li> <li> <p>payments : 取引の支払行</p> </li> <li> <p>from_walletable_type</p> <ul> <li>bank_account : 銀行口座</li> <li>credit_card : クレジットカード</li> <li>wallet : 現金</li> <li>private_account_item : プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）</li> </ul> </li> <li> <p>receipt_ids : 証憑ファイルID</p> </li> </ul> <h2 id=\"_3\">注意点</h2> <ul>     <li><p>本APIでは支払行(payments)の操作ができません。支払行の作成・更新・削除APIをご利用ください。</p></li>     <li><p>本APIでは+更新行(renews)の操作ができません。+更新行の作成・更新・削除APIをご利用ください。</p></li>     <li><p>本APIでは収入／支出の切替えができません。既存の取引を削除後、再度作成してください。</p></li>     <li><p>本APIで取引を更新すると、消費税の計算方法は必ず内税方式が選択されます。</p></li>     <li><p>セグメントタグ情報は法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。</p></li>     <li><p>partner_codeを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。またpartner_codeとpartner_idは同時に指定することはできません。</p></li>     <li>         <p>本APIでは取引の明細行(details)は、最大40行までになります。</p>     </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 取引ID | 
 **optional** | ***UpdateDealOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDealOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dealUpdateParams** | [**optional.Interface of DealUpdateParams**](DealUpdateParams.md)| 取引（収入／支出）の更新 | 

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

