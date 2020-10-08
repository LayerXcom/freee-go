# \CompaniesApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanies**](CompaniesApi.md#GetCompanies) | **Get** /api/1/companies | 事業所一覧の取得
[**GetCompany**](CompaniesApi.md#GetCompany) | **Get** /api/1/companies/{id} | 事業所の詳細情報の取得
[**UpdateCompany**](CompaniesApi.md#UpdateCompany) | **Put** /api/1/companies/{id} | 事業所情報の更新



## GetCompanies

> CompanyIndexResponse GetCompanies(ctx, )

事業所一覧の取得

 <h2 id=\"\">概要</h2>  <p>ユーザーが所属する事業所の一覧を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li>role <ul> <li>admin : 管理者</li>  <li>simple_accounting : 一般</li>  <li>self_only : 取引登録のみ</li>  <li>read_only : 閲覧のみ</li> </ul> </li> </ul>

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**CompanyIndexResponse**](companyIndexResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompany

> CompanyResponse GetCompany(ctx, id, optional)

事業所の詳細情報の取得

 <h2 id=\"\">概要</h2>  <p>ユーザーが所属する事業所の詳細を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li>role <ul> <li>admin : 管理者</li>  <li>simple_accounting : 一般</li>  <li>self_only : 取引登録のみ</li>  <li>read_only : 閲覧のみ</li> </ul> </li> </ul>  <h2 id=\"_3\">

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 事業所ID | 
 **optional** | ***GetCompanyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCompanyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **details** | **optional.Bool**| 取得情報に勘定科目・税区分コード・税区分・品目・取引先・部門・メモタグ・口座の一覧を含める | 
 **accountItems** | **optional.Bool**| 取得情報に勘定科目一覧を含める | 
 **taxes** | **optional.Bool**| 取得情報に税区分コード・税区分一覧を含める | 
 **items** | **optional.Bool**| 取得情報に品目一覧を含める | 
 **partners** | **optional.Bool**| 取得情報に取引先一覧を含める | 
 **sections** | **optional.Bool**| 取得情報に部門一覧を含める | 
 **tags** | **optional.Bool**| 取得情報にメモタグ一覧を含める | 
 **walletables** | **optional.Bool**| 取得情報に口座一覧を含める | 

### Return type

[**CompanyResponse**](companyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompany

> CompanyUpdateResponse UpdateCompany(ctx, id, optional)

事業所情報の更新

 <h2 id=\"\">概要</h2>  <p>ユーザーが所属する事業所の情報を更新する</p>  <p>※同時に複数のリクエストを受け付けない</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 事業所ID | 
 **optional** | ***UpdateCompanyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCompanyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyParams** | [**optional.Interface of CompanyParams**](CompanyParams.md)|  | 

### Return type

[**CompanyUpdateResponse**](companyUpdateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

