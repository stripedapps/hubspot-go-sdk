/*
Notes

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// BatchAPIService BatchAPI service
type BatchAPIService service

type ApiPostCrmV3ObjectsNotesBatchArchiveArchiveRequest struct {
	ctx context.Context
	ApiService *BatchAPIService
	batchInputSimplePublicObjectId *BatchInputSimplePublicObjectId
}

func (r ApiPostCrmV3ObjectsNotesBatchArchiveArchiveRequest) BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId BatchInputSimplePublicObjectId) ApiPostCrmV3ObjectsNotesBatchArchiveArchiveRequest {
	r.batchInputSimplePublicObjectId = &batchInputSimplePublicObjectId
	return r
}

func (r ApiPostCrmV3ObjectsNotesBatchArchiveArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsNotesBatchArchiveArchiveExecute(r)
}

/*
PostCrmV3ObjectsNotesBatchArchiveArchive Archive a batch of notes by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsNotesBatchArchiveArchiveRequest
*/
func (a *BatchAPIService) PostCrmV3ObjectsNotesBatchArchiveArchive(ctx context.Context) ApiPostCrmV3ObjectsNotesBatchArchiveArchiveRequest {
	return ApiPostCrmV3ObjectsNotesBatchArchiveArchiveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *BatchAPIService) PostCrmV3ObjectsNotesBatchArchiveArchiveExecute(r ApiPostCrmV3ObjectsNotesBatchArchiveArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchAPIService.PostCrmV3ObjectsNotesBatchArchiveArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/notes/batch/archive"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputSimplePublicObjectId == nil {
		return nil, reportError("batchInputSimplePublicObjectId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputSimplePublicObjectId
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPostCrmV3ObjectsNotesBatchCreateCreateRequest struct {
	ctx context.Context
	ApiService *BatchAPIService
	batchInputSimplePublicObjectInputForCreate *BatchInputSimplePublicObjectInputForCreate
}

func (r ApiPostCrmV3ObjectsNotesBatchCreateCreateRequest) BatchInputSimplePublicObjectInputForCreate(batchInputSimplePublicObjectInputForCreate BatchInputSimplePublicObjectInputForCreate) ApiPostCrmV3ObjectsNotesBatchCreateCreateRequest {
	r.batchInputSimplePublicObjectInputForCreate = &batchInputSimplePublicObjectInputForCreate
	return r
}

func (r ApiPostCrmV3ObjectsNotesBatchCreateCreateRequest) Execute() (*BatchResponseSimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsNotesBatchCreateCreateExecute(r)
}

/*
PostCrmV3ObjectsNotesBatchCreateCreate Create a batch of notes

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsNotesBatchCreateCreateRequest
*/
func (a *BatchAPIService) PostCrmV3ObjectsNotesBatchCreateCreate(ctx context.Context) ApiPostCrmV3ObjectsNotesBatchCreateCreateRequest {
	return ApiPostCrmV3ObjectsNotesBatchCreateCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BatchResponseSimplePublicObject
func (a *BatchAPIService) PostCrmV3ObjectsNotesBatchCreateCreateExecute(r ApiPostCrmV3ObjectsNotesBatchCreateCreateRequest) (*BatchResponseSimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchAPIService.PostCrmV3ObjectsNotesBatchCreateCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/notes/batch/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputSimplePublicObjectInputForCreate == nil {
		return localVarReturnValue, nil, reportError("batchInputSimplePublicObjectInputForCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputSimplePublicObjectInputForCreate
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostCrmV3ObjectsNotesBatchReadReadRequest struct {
	ctx context.Context
	ApiService *BatchAPIService
	batchReadInputSimplePublicObjectId *BatchReadInputSimplePublicObjectId
	archived *bool
}

func (r ApiPostCrmV3ObjectsNotesBatchReadReadRequest) BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId BatchReadInputSimplePublicObjectId) ApiPostCrmV3ObjectsNotesBatchReadReadRequest {
	r.batchReadInputSimplePublicObjectId = &batchReadInputSimplePublicObjectId
	return r
}

// Whether to return only results that have been archived.
func (r ApiPostCrmV3ObjectsNotesBatchReadReadRequest) Archived(archived bool) ApiPostCrmV3ObjectsNotesBatchReadReadRequest {
	r.archived = &archived
	return r
}

func (r ApiPostCrmV3ObjectsNotesBatchReadReadRequest) Execute() (*BatchResponseSimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsNotesBatchReadReadExecute(r)
}

/*
PostCrmV3ObjectsNotesBatchReadRead Read a batch of notes by internal ID, or unique property values

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsNotesBatchReadReadRequest
*/
func (a *BatchAPIService) PostCrmV3ObjectsNotesBatchReadRead(ctx context.Context) ApiPostCrmV3ObjectsNotesBatchReadReadRequest {
	return ApiPostCrmV3ObjectsNotesBatchReadReadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BatchResponseSimplePublicObject
func (a *BatchAPIService) PostCrmV3ObjectsNotesBatchReadReadExecute(r ApiPostCrmV3ObjectsNotesBatchReadReadRequest) (*BatchResponseSimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchAPIService.PostCrmV3ObjectsNotesBatchReadRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/notes/batch/read"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchReadInputSimplePublicObjectId == nil {
		return localVarReturnValue, nil, reportError("batchReadInputSimplePublicObjectId is required and must be specified")
	}

	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "form", "")
	} else {
		var defaultValue bool = false
		r.archived = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchReadInputSimplePublicObjectId
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostCrmV3ObjectsNotesBatchUpdateUpdateRequest struct {
	ctx context.Context
	ApiService *BatchAPIService
	batchInputSimplePublicObjectBatchInput *BatchInputSimplePublicObjectBatchInput
}

func (r ApiPostCrmV3ObjectsNotesBatchUpdateUpdateRequest) BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput BatchInputSimplePublicObjectBatchInput) ApiPostCrmV3ObjectsNotesBatchUpdateUpdateRequest {
	r.batchInputSimplePublicObjectBatchInput = &batchInputSimplePublicObjectBatchInput
	return r
}

func (r ApiPostCrmV3ObjectsNotesBatchUpdateUpdateRequest) Execute() (*BatchResponseSimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsNotesBatchUpdateUpdateExecute(r)
}

/*
PostCrmV3ObjectsNotesBatchUpdateUpdate Update a batch of notes by internal ID, or unique property values

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsNotesBatchUpdateUpdateRequest
*/
func (a *BatchAPIService) PostCrmV3ObjectsNotesBatchUpdateUpdate(ctx context.Context) ApiPostCrmV3ObjectsNotesBatchUpdateUpdateRequest {
	return ApiPostCrmV3ObjectsNotesBatchUpdateUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BatchResponseSimplePublicObject
func (a *BatchAPIService) PostCrmV3ObjectsNotesBatchUpdateUpdateExecute(r ApiPostCrmV3ObjectsNotesBatchUpdateUpdateRequest) (*BatchResponseSimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchAPIService.PostCrmV3ObjectsNotesBatchUpdateUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/notes/batch/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputSimplePublicObjectBatchInput == nil {
		return localVarReturnValue, nil, reportError("batchInputSimplePublicObjectBatchInput is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputSimplePublicObjectBatchInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostCrmV3ObjectsNotesBatchUpsertUpsertRequest struct {
	ctx context.Context
	ApiService *BatchAPIService
	batchInputSimplePublicObjectBatchInputUpsert *BatchInputSimplePublicObjectBatchInputUpsert
}

func (r ApiPostCrmV3ObjectsNotesBatchUpsertUpsertRequest) BatchInputSimplePublicObjectBatchInputUpsert(batchInputSimplePublicObjectBatchInputUpsert BatchInputSimplePublicObjectBatchInputUpsert) ApiPostCrmV3ObjectsNotesBatchUpsertUpsertRequest {
	r.batchInputSimplePublicObjectBatchInputUpsert = &batchInputSimplePublicObjectBatchInputUpsert
	return r
}

func (r ApiPostCrmV3ObjectsNotesBatchUpsertUpsertRequest) Execute() (*BatchResponseSimplePublicUpsertObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsNotesBatchUpsertUpsertExecute(r)
}

/*
PostCrmV3ObjectsNotesBatchUpsertUpsert Create or update a batch of notes by unique property values

Create or update records identified by a unique property value as specified by the `idProperty` query param. `idProperty` query param refers to a property whose values are unique for the object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsNotesBatchUpsertUpsertRequest
*/
func (a *BatchAPIService) PostCrmV3ObjectsNotesBatchUpsertUpsert(ctx context.Context) ApiPostCrmV3ObjectsNotesBatchUpsertUpsertRequest {
	return ApiPostCrmV3ObjectsNotesBatchUpsertUpsertRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BatchResponseSimplePublicUpsertObject
func (a *BatchAPIService) PostCrmV3ObjectsNotesBatchUpsertUpsertExecute(r ApiPostCrmV3ObjectsNotesBatchUpsertUpsertRequest) (*BatchResponseSimplePublicUpsertObject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BatchResponseSimplePublicUpsertObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchAPIService.PostCrmV3ObjectsNotesBatchUpsertUpsert")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/notes/batch/upsert"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputSimplePublicObjectBatchInputUpsert == nil {
		return localVarReturnValue, nil, reportError("batchInputSimplePublicObjectBatchInputUpsert is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputSimplePublicObjectBatchInputUpsert
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
