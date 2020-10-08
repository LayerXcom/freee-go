# \ManualJournalsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManualJournal**](ManualJournalsApi.md#CreateManualJournal) | **Post** /api/1/manual_journals | 振替伝票の作成
[**DestroyManualJournal**](ManualJournalsApi.md#DestroyManualJournal) | **Delete** /api/1/manual_journals/{id} | 振替伝票の削除
[**GetManualJournal**](ManualJournalsApi.md#GetManualJournal) | **Get** /api/1/manual_journals/{id} | 振替伝票の取得
[**GetManualJournals**](ManualJournalsApi.md#GetManualJournals) | **Get** /api/1/manual_journals | 振替伝票一覧の取得
[**UpdateManualJournal**](ManualJournalsApi.md#UpdateManualJournal) | **Put** /api/1/manual_journals/{id} | 振替伝票の更新



## CreateManualJournal

> ManualJournalResponse CreateManualJournal(ctx, optional)

振替伝票の作成

 <h2 id=\"\">概要</h2>  <p>指定した事業所の振替伝票を作成する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li> <p>issue_date : 発生日</p> </li>  <li> <p>adjustment : 決算整理仕訳フラグ（true: 決算整理仕訳, false: 日常仕訳）</p> </li>  <li> <p>txn_number : 仕訳番号</p> </li>  <li> <p>details : 振替伝票の貸借行</p> </li>  <li> <p>entry_side : 貸借区分</p>  <ul> <li>credit : 貸方</li>  <li>debit : 借方</li> </ul> </li>  <li> <p>amount : 金額</p> </li> </ul>  <h2 id=\"_3\">注意点</h2>  <ul> <li>振替伝票は売掛・買掛レポートには反映されません。債権・債務データの登録は取引(Deals)をお使いください。</li> <li>事業所の仕訳番号形式が有効な場合のみ、レスポンスで仕訳番号(txn_number)を返します。</li> <li>貸借合わせて100行まで仕訳行を登録できます。</li> <li>セグメントタグ情報は法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。</li> <li>partner_codeを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。またpartner_codeとpartner_idは同時に指定することはできません。</li></ul>  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateManualJournalOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateManualJournalOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manualJournalCreateParams** | [**optional.Interface of ManualJournalCreateParams**](ManualJournalCreateParams.md)| 振替伝票の作成 | 

### Return type

[**ManualJournalResponse**](manualJournalResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyManualJournal

> DestroyManualJournal(ctx, id, companyId)

振替伝票の削除

 <h2 id=\"\">概要</h2>  <p>指定した事業所の振替伝票を削除する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
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


## GetManualJournal

> ManualJournalResponse GetManualJournal(ctx, companyId, id)

振替伝票の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の振替伝票を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li> <p>issue_date : 発生日</p> </li> <li> <p>adjustment : 決算整理仕訳フラグ（true: 決算整理仕訳, false: 日常仕訳）</p> </li> <li> <p>txn_number : 仕訳番号</p> </li> <li> <p>details : 振替伝票の貸借行</p> </li> <li> <p>entry_side : 貸借区分</p> <ul> <li>credit : 貸方</li> <li>debit : 借方</li> </ul> </li> <li> <p>amount : 金額</p> </li> </ul>  <h2 id=\"_3\">注意点</h2>  <ul> <li>振替伝票は売掛・買掛レポートには反映されません。債権・債務データの登録は取引(Deals)をお使いください。</li> <li>事業所の仕訳番号形式が有効な場合のみ、レスポンスで仕訳番号(txn_number)を返します。</li> <li>セグメントタグ情報は法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
**id** | **int32**|  | 

### Return type

[**ManualJournalResponse**](manualJournalResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualJournals

> InlineResponse2005 GetManualJournals(ctx, companyId, optional)

振替伝票一覧の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の振替伝票一覧を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li> <p>issue_date : 発生日</p> </li>  <li> <p>adjustment : 決算整理仕訳フラグ（true: 決算整理仕訳, false: 日常仕訳）</p> </li>  <li> <p>txn_number : 仕訳番号</p> </li>  <li> <p>details : 振替伝票の貸借行</p> </li>  <li> <p>entry_side : 貸借区分</p>  <ul> <li>credit : 貸方</li>  <li>debit : 借方</li> </ul> </li>  <li> <p>amount : 金額</p> </li> </ul>  <h2 id=\"_3\">注意点</h2>  <ul> <li>振替伝票は売掛・買掛レポートには反映されません。債権・債務データの登録は取引(Deals)をお使いください。</li> <li>事業所の仕訳番号形式が有効な場合のみ、レスポンスで仕訳番号(txn_number)を返します。</li> <li>セグメントタグ情報は法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。</li> <li>partner_codeを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。またpartner_codeとpartner_idは同時に指定することはできません。</li></ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetManualJournalsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetManualJournalsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIssueDate** | **optional.String**| 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endIssueDate** | **optional.String**| 発生日で絞込：終了日(yyyy-mm-dd) | 
 **entrySide** | **optional.String**| 貸借で絞込 (貸方: credit, 借方: debit) | 
 **accountItemId** | **optional.Int32**| 勘定科目IDで絞込 | 
 **minAmount** | **optional.Int32**| 金額で絞込：下限 | 
 **maxAmount** | **optional.Int32**| 金額で絞込：上限 | 
 **partnerId** | **optional.Int32**| 取引先IDで絞込（0を指定すると、取引先が未選択の貸借行を絞り込めます） | 
 **partnerCode** | **optional.String**| 取引先コードで絞込 | 
 **itemId** | **optional.Int32**| 品目IDで絞込（0を指定すると、品目が未選択の貸借行を絞り込めます） | 
 **sectionId** | **optional.Int32**| 部門IDで絞込（0を指定すると、部門が未選択の貸借行を絞り込めます） | 
 **segment1TagId** | **optional.Int32**| セグメント１IDで絞込（0を指定すると、セグメント１が未選択の貸借行を絞り込めます） | 
 **segment2TagId** | **optional.Int32**| セグメント２IDで絞込（0を指定すると、セグメント２が未選択の貸借行を絞り込めます） | 
 **segment3TagId** | **optional.Int32**| セグメント３IDで絞込（0を指定すると、セグメント３が未選択の貸借行を絞り込めます） | 
 **commentStatus** | **optional.String**| コメント状態で絞込（自分宛のコメント: posted_with_mention, 自分宛のコメント-未解決: raised_with_mention, 自分宛のコメント-解決済: resolved_with_mention, コメントあり: posted, 未解決: raised, 解決済: resolved, コメントなし: none） | 
 **commentImportant** | **optional.Bool**| 重要コメント付きの振替伝票を絞込 | 
 **adjustment** | **optional.String**| 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 
 **txnNumber** | **optional.String**| 仕訳番号で絞込（事業所の仕訳番号形式が有効な場合のみ） | 
 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 500)  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManualJournal

> ManualJournalResponse UpdateManualJournal(ctx, id, optional)

振替伝票の更新

 <h2 id=\"\">概要</h2>  <p>指定した事業所の振替伝票を更新する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li> <p>issue_date : 発生日</p> </li>  <li> <p>adjustment : 決算整理仕訳フラグ（true: 決算整理仕訳, false: 日常仕訳）</p> </li>  <li> <p>txn_number : 仕訳番号</p> </li>  <li> <p>details : 振替伝票の貸借行</p> </li>  <li> <p>entry_side : 貸借区分</p>  <ul> <li>credit : 貸方</li>  <li>debit : 借方</li> </ul> </li>  <li> <p>amount : 金額</p> </li> </ul>  <h2 id=\"_3\">注意点</h2>  <ul> <li>振替伝票は売掛・買掛レポートには反映されません。債権・債務データの登録は取引(Deals)をお使いください。</li>  <li>事業所の仕訳番号形式が有効な場合のみ、レスポンスで仕訳番号(txn_number)を返します。</li> <li>貸借合わせて100行まで仕訳行を登録できます。</li>  <li>detailsに含まれない既存の貸借行は削除されます。更新後も残したい行は、必ず貸借行IDを指定してdetailsに含めてください。</li>  <li>detailsに含まれる貸借行IDの指定がある行は、更新行として扱われ更新されます。</li>  <li>detailsに含まれる貸借行IDの指定がない行は、新規行として扱われ追加されます。</li> <li>セグメントタグ情報は法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。</li> <li>partner_codeを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。またpartner_codeとpartner_idは同時に指定することはできません。</li></ul>  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***UpdateManualJournalOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateManualJournalOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **manualJournalUpdateParams** | [**optional.Interface of ManualJournalUpdateParams**](ManualJournalUpdateParams.md)| 振替伝票の更新 | 

### Return type

[**ManualJournalResponse**](manualJournalResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

