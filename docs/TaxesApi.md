# \TaxesApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTaxCode**](TaxesApi.md#GetTaxCode) | **Get** /api/1/taxes/codes/{code} | 税区分コードの取得
[**GetTaxCodes**](TaxesApi.md#GetTaxCodes) | **Get** /api/1/taxes/codes | 税区分コード一覧の取得
[**GetTaxesCompanies**](TaxesApi.md#GetTaxesCompanies) | **Get** /api/1/taxes/companies/{company_id} | 税区分コード詳細一覧の取得



## GetTaxCode

> TaxResponse GetTaxCode(ctx, code)

税区分コードの取得

 <h2 id=\"\">概要</h2>  <p>税区分コードを取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **int32**| 税区分コード | 

### Return type

[**TaxResponse**](taxResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxCodes

> InlineResponse2009 GetTaxCodes(ctx, )

税区分コード一覧の取得

 <h2 id=\"\">概要</h2>  <p>税区分コード一覧を取得する</p>

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxesCompanies

> InlineResponse20010 GetTaxesCompanies(ctx, companyId)

税区分コード詳細一覧の取得

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

