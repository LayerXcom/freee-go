# \PartnersApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePartner**](PartnersApi.md#CreatePartner) | **Post** /api/1/partners | 取引先の作成
[**DestroyPartner**](PartnersApi.md#DestroyPartner) | **Delete** /api/1/partners/{id} | 取引先の削除
[**GetPartner**](PartnersApi.md#GetPartner) | **Get** /api/1/partners/{id} | 取引先の取得
[**GetPartners**](PartnersApi.md#GetPartners) | **Get** /api/1/partners | 取引先一覧の取得
[**UpdatePartner**](PartnersApi.md#UpdatePartner) | **Put** /api/1/partners/{id} | 取引先の更新
[**UpdatePartnerByCode**](PartnersApi.md#UpdatePartnerByCode) | **Put** /api/1/partners/code/{code} | 取引先の更新



## CreatePartner

> PartnerResponse CreatePartner(ctx, partnerCreateParams)

取引先の作成

 <h2 id=\"\">概要</h2>  <p>指定した事業所の取引先を作成する</p> <ul> <li>codeを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。</li> <li>取引先コードの利用を有効にしている場合は、codeの指定は必須です。</li> <li>振込元口座ID（payer_walletable_id）, 振込手数料負担（transfer_fee_handling_side）, 支払期日設定（payment_term_attributes, 請求の入金期日設定（invoice_payment_term_attributes）は法人向けのプロフェッショナルプラン以上で利用可能です。</li></ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerCreateParams** | [**PartnerCreateParams**](PartnerCreateParams.md)| 取引先の作成 | 

### Return type

[**PartnerResponse**](partnerResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyPartner

> DestroyPartner(ctx, id, companyId)

取引先の削除

 <h2 id=\"\">概要</h2>  <p>指定した事業所の取引先を削除する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 取引先ID | 
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


## GetPartner

> PartnerResponse GetPartner(ctx, id, companyId)

取引先の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の取引先を取得する</p> <ul> <li>振込元口座ID（payer_walletable_id）, 振込手数料負担（transfer_fee_handling_side）, 支払期日設定（payment_term_attributes, 請求の入金期日設定（invoice_payment_term_attributes）は法人向けのプロフェッショナルプラン以上で利用可能です。</li></ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 取引先ID | 
**companyId** | **int32**| 事業所ID | 

### Return type

[**PartnerResponse**](partnerResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartners

> PartnersResponse GetPartners(ctx, companyId, optional)

取引先一覧の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の取引先一覧を取得する</p> <ul> <li>振込元口座ID（payer_walletable_id）, 振込手数料負担（transfer_fee_handling_side）は法人向けのプロフェッショナルプラン以上で利用可能です。</li></ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetPartnersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPartnersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) | 
 **keyword** | **optional.String**| 検索キーワード：取引先名・正式名称・カナ名称に対するあいまい検索で一致、またはショートカットキー1・2のいずれかに完全一致 | 

### Return type

[**PartnersResponse**](partnersResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartner

> PartnerResponse UpdatePartner(ctx, id, partnerUpdateParams)

取引先の更新

 <h2 id=\"\">概要</h2>  <p>指定した取引先の情報を更新する</p> <ul> <li>codeを指定、更新することはできません。</li> <li>振込元口座ID（payer_walletable_id）, 振込手数料負担（transfer_fee_handling_side）, 支払期日設定（payment_term_attributes, 請求の入金期日設定（invoice_payment_term_attributes）は法人向けのプロフェッショナルプラン以上で利用可能です。</li></ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 取引先ID | 
**partnerUpdateParams** | [**PartnerUpdateParams**](PartnerUpdateParams.md)| 取引先の更新 | 

### Return type

[**PartnerResponse**](partnerResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartnerByCode

> PartnerResponse UpdatePartnerByCode(ctx, code, partnerUpdateParams)

取引先の更新

 <h2 id=\"\">概要</h2>  <p>取引先コードをキーに、指定した取引先の情報を更新する</p> <ul> <li>このAPIを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。</li> <li>コードを日本語に設定している場合は、URLエンコードしてURLに含めるようにしてください。</li> <li>振込元口座ID（payer_walletable_id）, 振込手数料負担（transfer_fee_handling_side）, 支払期日設定（payment_term_attributes, 請求の入金期日設定（invoice_payment_term_attributes）は法人向けのプロフェッショナルプラン以上で利用可能です。</li></ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string**| 取引先コード | 
**partnerUpdateParams** | [**PartnerUpdateParams**](PartnerUpdateParams.md)| 取引先の更新 | 

### Return type

[**PartnerResponse**](partnerResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

