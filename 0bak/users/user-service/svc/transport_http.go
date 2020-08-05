// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 8907ffca23
// Version Date: Wed 27 Nov 2019 21:28:21 UTC

package svc

// This file provides server-side bindings for the HTTP transport.
// It utilizes the transport/http.Server.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"

	"context"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	// This service
	pb "github.com/sabnak227/jwt-demo/bak/users"
)

const contentType = "application/json; charset=utf-8"

var (
	_ = fmt.Sprint
	_ = bytes.Compare
	_ = strconv.Atoi
	_ = httptransport.NewServer
	_ = ioutil.NopCloser
	_ = pb.NewUserClient
	_ = io.Copy
	_ = errors.Wrap
)

// MakeHTTPHandler returns a handler that makes a set of endpoints available
// on predefined paths.
func MakeHTTPHandler(endpoints Endpoints, options ...httptransport.ServerOption) http.Handler {
	serverOptions := []httptransport.ServerOption{
		httptransport.ServerBefore(headersToContext),
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerAfter(httptransport.SetContentType(contentType)),
	}
	serverOptions = append(serverOptions, options...)
	m := mux.NewRouter()

	m.Methods("GET").Path("/users/").Handler(httptransport.NewServer(
		endpoints.ListUserEndpoint,
		DecodeHTTPListUserZeroRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))
	m.Methods("GET").Path("/users").Handler(httptransport.NewServer(
		endpoints.ListUserEndpoint,
		DecodeHTTPListUserOneRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))

	m.Methods("GET").Path("/users/{id}").Handler(httptransport.NewServer(
		endpoints.GetUserEndpoint,
		DecodeHTTPGetUserZeroRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))

	m.Methods("POST").Path("/users").Handler(httptransport.NewServer(
		endpoints.CreateUserEndpoint,
		DecodeHTTPCreateUserZeroRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))
	m.Methods("HEAD").Path("/users").Handler(httptransport.NewServer(
		endpoints.CreateUserEndpoint,
		DecodeHTTPCreateUserOneRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))

	m.Methods("PUT").Path("/users/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateUserEndpoint,
		DecodeHTTPUpdateUserZeroRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))
	m.Methods("HEAD").Path("/users/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateUserEndpoint,
		DecodeHTTPUpdateUserOneRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))

	m.Methods("DELETE").Path("/users/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteUserEndpoint,
		DecodeHTTPDeleteUserZeroRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))
	m.Methods("HEAD").Path("/users/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteUserEndpoint,
		DecodeHTTPDeleteUserOneRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))

	m.Methods("POST").Path("/users/{id}/auth").Handler(httptransport.NewServer(
		endpoints.AuthUserEndpoint,
		DecodeHTTPAuthUserZeroRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))
	m.Methods("HEAD").Path("/users/{id}/auth").Handler(httptransport.NewServer(
		endpoints.AuthUserEndpoint,
		DecodeHTTPAuthUserOneRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))
	return m
}

// ErrorEncoder writes the error to the ResponseWriter, by default a content
// type of application/json, a body of json with key "error" and the value
// error.Error(), and a status code of 500. If the error implements Headerer,
// the provided headers will be applied to the response. If the error
// implements json.Marshaler, and the marshaling succeeds, the JSON encoded
// form of the error will be used. If the error implements StatusCoder, the
// provided StatusCode will be used instead of 500.
func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	body, _ := json.Marshal(errorWrapper{Error: err.Error()})
	if marshaler, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
			body = jsonBody
		}
	}
	w.Header().Set("Content-Type", contentType)
	if headerer, ok := err.(httptransport.Headerer); ok {
		for k := range headerer.Headers() {
			w.Header().Set(k, headerer.Headers().Get(k))
		}
	}
	code := http.StatusInternalServerError
	if sc, ok := err.(httptransport.StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	w.Write(body)
}

type errorWrapper struct {
	Error string `json:"error"`
}

// httpError satisfies the Headerer and StatusCoder interfaces in
// package github.com/go-kit/kit/transport/http.
type httpError struct {
	error
	statusCode int
	headers    map[string][]string
}

func (h httpError) StatusCode() int {
	return h.statusCode
}

func (h httpError) Headers() http.Header {
	return h.headers
}

// Server Decode

// DecodeHTTPListUserZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded listuser request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPListUserZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.ListUserRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if LastIdListUserStrArr, ok := queryParams["last_id"]; ok {
		LastIdListUserStr := LastIdListUserStrArr[0]
		LastIdListUser, err := strconv.ParseInt(LastIdListUserStr, 10, 32)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Error while extracting LastIdListUser from query, queryParams: %v", queryParams))
		}
		req.LastId = int32(LastIdListUser)
	}

	if Size_ListUserStrArr, ok := queryParams["size"]; ok {
		Size_ListUserStr := Size_ListUserStrArr[0]
		Size_ListUser, err := strconv.ParseInt(Size_ListUserStr, 10, 32)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Error while extracting Size_ListUser from query, queryParams: %v", queryParams))
		}
		req.Size_ = int32(Size_ListUser)
	}

	return &req, err
}

// DecodeHTTPListUserOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded listuser request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPListUserOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.ListUserRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if LastIdListUserStrArr, ok := queryParams["last_id"]; ok {
		LastIdListUserStr := LastIdListUserStrArr[0]
		LastIdListUser, err := strconv.ParseInt(LastIdListUserStr, 10, 32)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Error while extracting LastIdListUser from query, queryParams: %v", queryParams))
		}
		req.LastId = int32(LastIdListUser)
	}

	if Size_ListUserStrArr, ok := queryParams["size"]; ok {
		Size_ListUserStr := Size_ListUserStrArr[0]
		Size_ListUser, err := strconv.ParseInt(Size_ListUserStr, 10, 32)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Error while extracting Size_ListUser from query, queryParams: %v", queryParams))
		}
		req.Size_ = int32(Size_ListUser)
	}

	return &req, err
}

// DecodeHTTPGetUserZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded getuser request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPGetUserZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.GetUserRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	return &req, err
}

// DecodeHTTPCreateUserZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded createuser request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPCreateUserZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.CreateUserRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	return &req, err
}

// DecodeHTTPCreateUserOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded createuser request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPCreateUserOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.CreateUserRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if FirstNameCreateUserStrArr, ok := queryParams["first_name"]; ok {
		FirstNameCreateUserStr := FirstNameCreateUserStrArr[0]
		FirstNameCreateUser := FirstNameCreateUserStr
		req.FirstName = FirstNameCreateUser
	}

	if LastNameCreateUserStrArr, ok := queryParams["last_name"]; ok {
		LastNameCreateUserStr := LastNameCreateUserStrArr[0]
		LastNameCreateUser := LastNameCreateUserStr
		req.LastName = LastNameCreateUser
	}

	if EmailCreateUserStrArr, ok := queryParams["email"]; ok {
		EmailCreateUserStr := EmailCreateUserStrArr[0]
		EmailCreateUser := EmailCreateUserStr
		req.Email = EmailCreateUser
	}

	if PasswordCreateUserStrArr, ok := queryParams["password"]; ok {
		PasswordCreateUserStr := PasswordCreateUserStrArr[0]
		PasswordCreateUser := PasswordCreateUserStr
		req.Password = PasswordCreateUser
	}

	if Address1CreateUserStrArr, ok := queryParams["address1"]; ok {
		Address1CreateUserStr := Address1CreateUserStrArr[0]
		Address1CreateUser := Address1CreateUserStr
		req.Address1 = Address1CreateUser
	}

	if Address2CreateUserStrArr, ok := queryParams["address2"]; ok {
		Address2CreateUserStr := Address2CreateUserStrArr[0]
		Address2CreateUser := Address2CreateUserStr
		req.Address2 = Address2CreateUser
	}

	if CityCreateUserStrArr, ok := queryParams["city"]; ok {
		CityCreateUserStr := CityCreateUserStrArr[0]
		CityCreateUser := CityCreateUserStr
		req.City = CityCreateUser
	}

	if StateCreateUserStrArr, ok := queryParams["state"]; ok {
		StateCreateUserStr := StateCreateUserStrArr[0]
		StateCreateUser := StateCreateUserStr
		req.State = StateCreateUser
	}

	if CountryCreateUserStrArr, ok := queryParams["country"]; ok {
		CountryCreateUserStr := CountryCreateUserStrArr[0]
		CountryCreateUser := CountryCreateUserStr
		req.Country = CountryCreateUser
	}

	if PhoneCreateUserStrArr, ok := queryParams["phone"]; ok {
		PhoneCreateUserStr := PhoneCreateUserStrArr[0]
		PhoneCreateUser := PhoneCreateUserStr
		req.Phone = PhoneCreateUser
	}

	return &req, err
}

// DecodeHTTPUpdateUserZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded updateuser request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPUpdateUserZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.UpdateUserRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	return &req, err
}

// DecodeHTTPUpdateUserOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded updateuser request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPUpdateUserOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.UpdateUserRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if FirstNameUpdateUserStrArr, ok := queryParams["first_name"]; ok {
		FirstNameUpdateUserStr := FirstNameUpdateUserStrArr[0]
		FirstNameUpdateUser := FirstNameUpdateUserStr
		req.FirstName = FirstNameUpdateUser
	}

	if LastNameUpdateUserStrArr, ok := queryParams["last_name"]; ok {
		LastNameUpdateUserStr := LastNameUpdateUserStrArr[0]
		LastNameUpdateUser := LastNameUpdateUserStr
		req.LastName = LastNameUpdateUser
	}

	if EmailUpdateUserStrArr, ok := queryParams["email"]; ok {
		EmailUpdateUserStr := EmailUpdateUserStrArr[0]
		EmailUpdateUser := EmailUpdateUserStr
		req.Email = EmailUpdateUser
	}

	if PasswordUpdateUserStrArr, ok := queryParams["password"]; ok {
		PasswordUpdateUserStr := PasswordUpdateUserStrArr[0]
		PasswordUpdateUser := PasswordUpdateUserStr
		req.Password = PasswordUpdateUser
	}

	if Address1UpdateUserStrArr, ok := queryParams["address1"]; ok {
		Address1UpdateUserStr := Address1UpdateUserStrArr[0]
		Address1UpdateUser := Address1UpdateUserStr
		req.Address1 = Address1UpdateUser
	}

	if Address2UpdateUserStrArr, ok := queryParams["address2"]; ok {
		Address2UpdateUserStr := Address2UpdateUserStrArr[0]
		Address2UpdateUser := Address2UpdateUserStr
		req.Address2 = Address2UpdateUser
	}

	if CityUpdateUserStrArr, ok := queryParams["city"]; ok {
		CityUpdateUserStr := CityUpdateUserStrArr[0]
		CityUpdateUser := CityUpdateUserStr
		req.City = CityUpdateUser
	}

	if StateUpdateUserStrArr, ok := queryParams["state"]; ok {
		StateUpdateUserStr := StateUpdateUserStrArr[0]
		StateUpdateUser := StateUpdateUserStr
		req.State = StateUpdateUser
	}

	if CountryUpdateUserStrArr, ok := queryParams["country"]; ok {
		CountryUpdateUserStr := CountryUpdateUserStrArr[0]
		CountryUpdateUser := CountryUpdateUserStr
		req.Country = CountryUpdateUser
	}

	if PhoneUpdateUserStrArr, ok := queryParams["phone"]; ok {
		PhoneUpdateUserStr := PhoneUpdateUserStrArr[0]
		PhoneUpdateUser := PhoneUpdateUserStr
		req.Phone = PhoneUpdateUser
	}

	return &req, err
}

// DecodeHTTPDeleteUserZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded deleteuser request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPDeleteUserZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.DeleteUserRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	return &req, err
}

// DecodeHTTPDeleteUserOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded deleteuser request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPDeleteUserOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.DeleteUserRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	return &req, err
}

// DecodeHTTPAuthUserZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded authuser request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPAuthUserZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.AuthUserRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	return &req, err
}

// DecodeHTTPAuthUserOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded authuser request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPAuthUserOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.AuthUserRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if EmailAuthUserStrArr, ok := queryParams["email"]; ok {
		EmailAuthUserStr := EmailAuthUserStrArr[0]
		EmailAuthUser := EmailAuthUserStr
		req.Email = EmailAuthUser
	}

	if PasswordAuthUserStrArr, ok := queryParams["password"]; ok {
		PasswordAuthUserStr := PasswordAuthUserStrArr[0]
		PasswordAuthUser := PasswordAuthUserStr
		req.Password = PasswordAuthUser
	}

	return &req, err
}

// EncodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeHTTPGenericResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	marshaller := jsonpb.Marshaler{
		EmitDefaults: false,
		OrigName:     true,
	}

	return marshaller.Marshal(w, response.(proto.Message))
}

// Helper functions

func headersToContext(ctx context.Context, r *http.Request) context.Context {
	for k, _ := range r.Header {
		// The key is added both in http format (k) which has had
		// http.CanonicalHeaderKey called on it in transport as well as the
		// strings.ToLower which is the grpc metadata format of the key so
		// that it can be accessed in either format
		ctx = context.WithValue(ctx, k, r.Header.Get(k))
		ctx = context.WithValue(ctx, strings.ToLower(k), r.Header.Get(k))
	}

	// Tune specific change.
	// also add the request url
	ctx = context.WithValue(ctx, "request-url", r.URL.Path)
	ctx = context.WithValue(ctx, "transport", "HTTPJSON")

	return ctx
}
