// Package apis provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package apis

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	CookieAuthScopes = "cookieAuth.Scopes"
)

// SignUpValidationError defines model for SignUpValidationError.
type SignUpValidationError struct {
	BackIdentification  *[]string `json:"backIdentification,omitempty"`
	Birthday            *[]string `json:"birthday,omitempty"`
	Email               *[]string `json:"email,omitempty"`
	FirstName           *[]string `json:"firstName,omitempty"`
	FrontIdentification *[]string `json:"frontIdentification,omitempty"`
	LastName            *[]string `json:"lastName,omitempty"`
	Password            *[]string `json:"password,omitempty"`
}

// StoreTodoValidationError defines model for StoreTodoValidationError.
type StoreTodoValidationError struct {
	Content *[]string `json:"content,omitempty"`
	Title   *[]string `json:"title,omitempty"`
}

// Todo defines model for Todo.
type Todo struct {
	Content string `json:"content"`
	Id      int    `json:"id"`
	Title   string `json:"title"`
}

// CsrfResponse defines model for CsrfResponse.
type CsrfResponse struct {
	CsrfToken string `json:"csrf_token"`
}

// DeleteTodoResponse defines model for DeleteTodoResponse.
type DeleteTodoResponse struct {
	Code   int64 `json:"code"`
	Result bool  `json:"result"`
}

// FetchTodosResponse defines model for FetchTodosResponse.
type FetchTodosResponse struct {
	Todos []Todo `json:"todos"`
}

// InternalServerErrorResponse defines model for InternalServerErrorResponse.
type InternalServerErrorResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// NotFoundErrorResponse defines model for NotFoundErrorResponse.
type NotFoundErrorResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// ShowTodoResponse defines model for ShowTodoResponse.
type ShowTodoResponse struct {
	Todo Todo `json:"todo"`
}

// SignInBadRequestResponse defines model for SignInBadRequestResponse.
type SignInBadRequestResponse struct {
	Errors []string `json:"errors"`
}

// SignInOkResponse defines model for SignInOkResponse.
type SignInOkResponse = map[string]interface{}

// SignUpResponse defines model for SignUpResponse.
type SignUpResponse struct {
	Code   int64                 `json:"code"`
	Errors SignUpValidationError `json:"errors"`
}

// StoreTodoResponse defines model for StoreTodoResponse.
type StoreTodoResponse struct {
	Code   int64                    `json:"code"`
	Errors StoreTodoValidationError `json:"errors"`
}

// UnauthorizedErrorResponse defines model for UnauthorizedErrorResponse.
type UnauthorizedErrorResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// SignInInput defines model for SignInInput.
type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// StoreTodoInput defines model for StoreTodoInput.
type StoreTodoInput struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}

// PostAuthSignInJSONBody defines parameters for PostAuthSignIn.
type PostAuthSignInJSONBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// PostAuthSignUpMultipartBody defines parameters for PostAuthSignUp.
type PostAuthSignUpMultipartBody struct {
	BackIdentification  *openapi_types.File `json:"backIdentification,omitempty"`
	Birthday            *openapi_types.Date `json:"birthday,omitempty"`
	Email               string              `json:"email"`
	FirstName           string              `json:"firstName"`
	FrontIdentification *openapi_types.File `json:"frontIdentification,omitempty"`
	LastName            string              `json:"lastName"`
	Password            string              `json:"password"`
}

// PostAuthValidateSignUpMultipartBody defines parameters for PostAuthValidateSignUp.
type PostAuthValidateSignUpMultipartBody struct {
	BackIdentification  *openapi_types.File `json:"backIdentification,omitempty"`
	Birthday            *openapi_types.Date `json:"birthday,omitempty"`
	Email               string              `json:"email"`
	FirstName           string              `json:"firstName"`
	FrontIdentification *openapi_types.File `json:"frontIdentification,omitempty"`
	LastName            string              `json:"lastName"`
	Password            string              `json:"password"`
}

// PostTodosJSONBody defines parameters for PostTodos.
type PostTodosJSONBody struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}

// PatchTodoJSONBody defines parameters for PatchTodo.
type PatchTodoJSONBody struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}

// PostAuthSignInJSONRequestBody defines body for PostAuthSignIn for application/json ContentType.
type PostAuthSignInJSONRequestBody PostAuthSignInJSONBody

// PostAuthSignUpMultipartRequestBody defines body for PostAuthSignUp for multipart/form-data ContentType.
type PostAuthSignUpMultipartRequestBody PostAuthSignUpMultipartBody

// PostAuthValidateSignUpMultipartRequestBody defines body for PostAuthValidateSignUp for multipart/form-data ContentType.
type PostAuthValidateSignUpMultipartRequestBody PostAuthValidateSignUpMultipartBody

// PostTodosJSONRequestBody defines body for PostTodos for application/json ContentType.
type PostTodosJSONRequestBody PostTodosJSONBody

// PatchTodoJSONRequestBody defines body for PatchTodo for application/json ContentType.
type PatchTodoJSONRequestBody PatchTodoJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get Csrf
	// (GET /auth/csrf)
	GetAuthCsrf(ctx echo.Context) error
	// Sign In
	// (POST /auth/signIn)
	PostAuthSignIn(ctx echo.Context) error
	// SignUp
	// (POST /auth/signUp)
	PostAuthSignUp(ctx echo.Context) error
	// Validate SignUp
	// (POST /auth/validateSignUp)
	PostAuthValidateSignUp(ctx echo.Context) error
	// Fetch Todos
	// (GET /todos)
	GetTodos(ctx echo.Context) error
	// Create Todo
	// (POST /todos)
	PostTodos(ctx echo.Context) error
	// Delete Todo
	// (DELETE /todos/{id})
	DeleteTodo(ctx echo.Context, id string) error
	// Show Todo
	// (GET /todos/{id})
	GetTodo(ctx echo.Context, id string) error
	// Update Todo
	// (PATCH /todos/{id})
	PatchTodo(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAuthCsrf converts echo context to params.
func (w *ServerInterfaceWrapper) GetAuthCsrf(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAuthCsrf(ctx)
	return err
}

// PostAuthSignIn converts echo context to params.
func (w *ServerInterfaceWrapper) PostAuthSignIn(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostAuthSignIn(ctx)
	return err
}

// PostAuthSignUp converts echo context to params.
func (w *ServerInterfaceWrapper) PostAuthSignUp(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostAuthSignUp(ctx)
	return err
}

// PostAuthValidateSignUp converts echo context to params.
func (w *ServerInterfaceWrapper) PostAuthValidateSignUp(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostAuthValidateSignUp(ctx)
	return err
}

// GetTodos converts echo context to params.
func (w *ServerInterfaceWrapper) GetTodos(ctx echo.Context) error {
	var err error

	ctx.Set(CookieAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTodos(ctx)
	return err
}

// PostTodos converts echo context to params.
func (w *ServerInterfaceWrapper) PostTodos(ctx echo.Context) error {
	var err error

	ctx.Set(CookieAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTodos(ctx)
	return err
}

// DeleteTodo converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTodo(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(CookieAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTodo(ctx, id)
	return err
}

// GetTodo converts echo context to params.
func (w *ServerInterfaceWrapper) GetTodo(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(CookieAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTodo(ctx, id)
	return err
}

// PatchTodo converts echo context to params.
func (w *ServerInterfaceWrapper) PatchTodo(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(CookieAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchTodo(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/auth/csrf", wrapper.GetAuthCsrf)
	router.POST(baseURL+"/auth/signIn", wrapper.PostAuthSignIn)
	router.POST(baseURL+"/auth/signUp", wrapper.PostAuthSignUp)
	router.POST(baseURL+"/auth/validateSignUp", wrapper.PostAuthValidateSignUp)
	router.GET(baseURL+"/todos", wrapper.GetTodos)
	router.POST(baseURL+"/todos", wrapper.PostTodos)
	router.DELETE(baseURL+"/todos/:id", wrapper.DeleteTodo)
	router.GET(baseURL+"/todos/:id", wrapper.GetTodo)
	router.PATCH(baseURL+"/todos/:id", wrapper.PatchTodo)

}

type CsrfResponseJSONResponse struct {
	CsrfToken string `json:"csrf_token"`
}

type DeleteTodoResponseJSONResponse struct {
	Code   int64 `json:"code"`
	Result bool  `json:"result"`
}

type FetchTodosResponseJSONResponse struct {
	Todos []Todo `json:"todos"`
}

type InternalServerErrorResponseJSONResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type NotFoundErrorResponseJSONResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ShowTodoResponseJSONResponse struct {
	Todo Todo `json:"todo"`
}

type SignInBadRequestResponseJSONResponse struct {
	Errors []string `json:"errors"`
}

type SignInOkResponseResponseHeaders struct {
	SetCookie string
}
type SignInOkResponseJSONResponse struct {
	Body map[string]interface{}

	Headers SignInOkResponseResponseHeaders
}

type SignUpResponseJSONResponse struct {
	Code   int64                 `json:"code"`
	Errors SignUpValidationError `json:"errors"`
}

type StoreTodoResponseJSONResponse struct {
	Code   int64                    `json:"code"`
	Errors StoreTodoValidationError `json:"errors"`
}

type UnauthorizedErrorResponseJSONResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type GetAuthCsrfRequestObject struct {
}

type GetAuthCsrfResponseObject interface {
	VisitGetAuthCsrfResponse(w http.ResponseWriter) error
}

type GetAuthCsrf200JSONResponse struct{ CsrfResponseJSONResponse }

func (response GetAuthCsrf200JSONResponse) VisitGetAuthCsrfResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetAuthCsrf500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response GetAuthCsrf500JSONResponse) VisitGetAuthCsrfResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostAuthSignInRequestObject struct {
	Body *PostAuthSignInJSONRequestBody
}

type PostAuthSignInResponseObject interface {
	VisitPostAuthSignInResponse(w http.ResponseWriter) error
}

type PostAuthSignIn200JSONResponse struct{ SignInOkResponseJSONResponse }

func (response PostAuthSignIn200JSONResponse) VisitPostAuthSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Set-Cookie", fmt.Sprint(response.Headers.SetCookie))
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostAuthSignIn400JSONResponse struct {
	SignInBadRequestResponseJSONResponse
}

func (response PostAuthSignIn400JSONResponse) VisitPostAuthSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostAuthSignIn500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response PostAuthSignIn500JSONResponse) VisitPostAuthSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostAuthSignUpRequestObject struct {
	Body *multipart.Reader
}

type PostAuthSignUpResponseObject interface {
	VisitPostAuthSignUpResponse(w http.ResponseWriter) error
}

type PostAuthSignUp200JSONResponse struct{ SignUpResponseJSONResponse }

func (response PostAuthSignUp200JSONResponse) VisitPostAuthSignUpResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostAuthSignUp400JSONResponse struct {
	Code   int64                 `json:"code"`
	Errors SignUpValidationError `json:"errors"`
}

func (response PostAuthSignUp400JSONResponse) VisitPostAuthSignUpResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostAuthSignUp500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response PostAuthSignUp500JSONResponse) VisitPostAuthSignUpResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostAuthValidateSignUpRequestObject struct {
	Body *multipart.Reader
}

type PostAuthValidateSignUpResponseObject interface {
	VisitPostAuthValidateSignUpResponse(w http.ResponseWriter) error
}

type PostAuthValidateSignUp200JSONResponse struct{ SignUpResponseJSONResponse }

func (response PostAuthValidateSignUp200JSONResponse) VisitPostAuthValidateSignUpResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostAuthValidateSignUp400JSONResponse struct {
	Code   int64                 `json:"code"`
	Errors SignUpValidationError `json:"errors"`
}

func (response PostAuthValidateSignUp400JSONResponse) VisitPostAuthValidateSignUpResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostAuthValidateSignUp500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response PostAuthValidateSignUp500JSONResponse) VisitPostAuthValidateSignUpResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetTodosRequestObject struct {
}

type GetTodosResponseObject interface {
	VisitGetTodosResponse(w http.ResponseWriter) error
}

type GetTodos200JSONResponse struct{ FetchTodosResponseJSONResponse }

func (response GetTodos200JSONResponse) VisitGetTodosResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTodos401JSONResponse struct {
	UnauthorizedErrorResponseJSONResponse
}

func (response GetTodos401JSONResponse) VisitGetTodosResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type GetTodos500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response GetTodos500JSONResponse) VisitGetTodosResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostTodosRequestObject struct {
	Body *PostTodosJSONRequestBody
}

type PostTodosResponseObject interface {
	VisitPostTodosResponse(w http.ResponseWriter) error
}

type PostTodos200JSONResponse struct{ StoreTodoResponseJSONResponse }

func (response PostTodos200JSONResponse) VisitPostTodosResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTodos400JSONResponse struct {
	Code   int64                    `json:"code"`
	Errors StoreTodoValidationError `json:"errors"`
}

func (response PostTodos400JSONResponse) VisitPostTodosResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostTodos401JSONResponse struct {
	UnauthorizedErrorResponseJSONResponse
}

func (response PostTodos401JSONResponse) VisitPostTodosResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type PostTodos500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response PostTodos500JSONResponse) VisitPostTodosResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTodoRequestObject struct {
	Id string `json:"id"`
}

type DeleteTodoResponseObject interface {
	VisitDeleteTodoResponse(w http.ResponseWriter) error
}

type DeleteTodo200JSONResponse struct{ DeleteTodoResponseJSONResponse }

func (response DeleteTodo200JSONResponse) VisitDeleteTodoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTodo401JSONResponse struct {
	UnauthorizedErrorResponseJSONResponse
}

func (response DeleteTodo401JSONResponse) VisitDeleteTodoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTodo404JSONResponse struct {
	NotFoundErrorResponseJSONResponse
}

func (response DeleteTodo404JSONResponse) VisitDeleteTodoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTodo500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response DeleteTodo500JSONResponse) VisitDeleteTodoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetTodoRequestObject struct {
	Id string `json:"id"`
}

type GetTodoResponseObject interface {
	VisitGetTodoResponse(w http.ResponseWriter) error
}

type GetTodo200JSONResponse struct{ ShowTodoResponseJSONResponse }

func (response GetTodo200JSONResponse) VisitGetTodoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTodo401JSONResponse struct {
	UnauthorizedErrorResponseJSONResponse
}

func (response GetTodo401JSONResponse) VisitGetTodoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type GetTodo404JSONResponse struct {
	NotFoundErrorResponseJSONResponse
}

func (response GetTodo404JSONResponse) VisitGetTodoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetTodo500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response GetTodo500JSONResponse) VisitGetTodoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PatchTodoRequestObject struct {
	Id   string `json:"id"`
	Body *PatchTodoJSONRequestBody
}

type PatchTodoResponseObject interface {
	VisitPatchTodoResponse(w http.ResponseWriter) error
}

type PatchTodo200JSONResponse struct{ StoreTodoResponseJSONResponse }

func (response PatchTodo200JSONResponse) VisitPatchTodoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchTodo400JSONResponse struct {
	Code   int64                    `json:"code"`
	Errors StoreTodoValidationError `json:"errors"`
}

func (response PatchTodo400JSONResponse) VisitPatchTodoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PatchTodo401JSONResponse struct {
	UnauthorizedErrorResponseJSONResponse
}

func (response PatchTodo401JSONResponse) VisitPatchTodoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type PatchTodo404JSONResponse struct {
	NotFoundErrorResponseJSONResponse
}

func (response PatchTodo404JSONResponse) VisitPatchTodoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type PatchTodo500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response PatchTodo500JSONResponse) VisitPatchTodoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get Csrf
	// (GET /auth/csrf)
	GetAuthCsrf(ctx context.Context, request GetAuthCsrfRequestObject) (GetAuthCsrfResponseObject, error)
	// Sign In
	// (POST /auth/signIn)
	PostAuthSignIn(ctx context.Context, request PostAuthSignInRequestObject) (PostAuthSignInResponseObject, error)
	// SignUp
	// (POST /auth/signUp)
	PostAuthSignUp(ctx context.Context, request PostAuthSignUpRequestObject) (PostAuthSignUpResponseObject, error)
	// Validate SignUp
	// (POST /auth/validateSignUp)
	PostAuthValidateSignUp(ctx context.Context, request PostAuthValidateSignUpRequestObject) (PostAuthValidateSignUpResponseObject, error)
	// Fetch Todos
	// (GET /todos)
	GetTodos(ctx context.Context, request GetTodosRequestObject) (GetTodosResponseObject, error)
	// Create Todo
	// (POST /todos)
	PostTodos(ctx context.Context, request PostTodosRequestObject) (PostTodosResponseObject, error)
	// Delete Todo
	// (DELETE /todos/{id})
	DeleteTodo(ctx context.Context, request DeleteTodoRequestObject) (DeleteTodoResponseObject, error)
	// Show Todo
	// (GET /todos/{id})
	GetTodo(ctx context.Context, request GetTodoRequestObject) (GetTodoResponseObject, error)
	// Update Todo
	// (PATCH /todos/{id})
	PatchTodo(ctx context.Context, request PatchTodoRequestObject) (PatchTodoResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetAuthCsrf operation middleware
func (sh *strictHandler) GetAuthCsrf(ctx echo.Context) error {
	var request GetAuthCsrfRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetAuthCsrf(ctx.Request().Context(), request.(GetAuthCsrfRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetAuthCsrf")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetAuthCsrfResponseObject); ok {
		return validResponse.VisitGetAuthCsrfResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostAuthSignIn operation middleware
func (sh *strictHandler) PostAuthSignIn(ctx echo.Context) error {
	var request PostAuthSignInRequestObject

	var body PostAuthSignInJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostAuthSignIn(ctx.Request().Context(), request.(PostAuthSignInRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostAuthSignIn")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostAuthSignInResponseObject); ok {
		return validResponse.VisitPostAuthSignInResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostAuthSignUp operation middleware
func (sh *strictHandler) PostAuthSignUp(ctx echo.Context) error {
	var request PostAuthSignUpRequestObject

	if reader, err := ctx.Request().MultipartReader(); err != nil {
		return err
	} else {
		request.Body = reader
	}

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostAuthSignUp(ctx.Request().Context(), request.(PostAuthSignUpRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostAuthSignUp")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostAuthSignUpResponseObject); ok {
		return validResponse.VisitPostAuthSignUpResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostAuthValidateSignUp operation middleware
func (sh *strictHandler) PostAuthValidateSignUp(ctx echo.Context) error {
	var request PostAuthValidateSignUpRequestObject

	if reader, err := ctx.Request().MultipartReader(); err != nil {
		return err
	} else {
		request.Body = reader
	}

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostAuthValidateSignUp(ctx.Request().Context(), request.(PostAuthValidateSignUpRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostAuthValidateSignUp")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostAuthValidateSignUpResponseObject); ok {
		return validResponse.VisitPostAuthValidateSignUpResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTodos operation middleware
func (sh *strictHandler) GetTodos(ctx echo.Context) error {
	var request GetTodosRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTodos(ctx.Request().Context(), request.(GetTodosRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTodos")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTodosResponseObject); ok {
		return validResponse.VisitGetTodosResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostTodos operation middleware
func (sh *strictHandler) PostTodos(ctx echo.Context) error {
	var request PostTodosRequestObject

	var body PostTodosJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTodos(ctx.Request().Context(), request.(PostTodosRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTodos")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTodosResponseObject); ok {
		return validResponse.VisitPostTodosResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteTodo operation middleware
func (sh *strictHandler) DeleteTodo(ctx echo.Context, id string) error {
	var request DeleteTodoRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTodo(ctx.Request().Context(), request.(DeleteTodoRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTodo")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTodoResponseObject); ok {
		return validResponse.VisitDeleteTodoResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTodo operation middleware
func (sh *strictHandler) GetTodo(ctx echo.Context, id string) error {
	var request GetTodoRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTodo(ctx.Request().Context(), request.(GetTodoRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTodo")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTodoResponseObject); ok {
		return validResponse.VisitGetTodoResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchTodo operation middleware
func (sh *strictHandler) PatchTodo(ctx echo.Context, id string) error {
	var request PatchTodoRequestObject

	request.Id = id

	var body PatchTodoJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchTodo(ctx.Request().Context(), request.(PatchTodoRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchTodo")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchTodoResponseObject); ok {
		return validResponse.VisitPatchTodoResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xYUY/bNgz+K4a2R7dOt9tQ+K3t1iIY0BZN05dDMOhsJlHPkVyJbpEd/N8HSo4tn5XE",
	"Se8Oa9c3xxFpft9HUSJvWKY2pZIg0bD0hmn4VIHB5yoXYF/MxEpO5VSWFdLPTEkEaR95WRYi4yiUTD4a",
	"Jemdydaw4fRUalWCxsYLbLgo6AG3JbCUGdRCrlgds5Ib80XpPPBnHdtwhIacpZeND89iEe8s1NVHyJDV",
	"ZJKDybQoKSyWNuFHkQNQx/bFvAzh2VQFipJrTJZKbx7lHPkhSFc8u57mIFEsGxboLZlyZCm7EpLrLYuH",
	"iK+ExnXOt73lOUcILd5P3FJog6/5BsL/aiXxrPAKfsDteLW68DyX8fkizssomlY7EVFpeK9y9bV56ZkN",
	"sKLAAo4Ddcvi1tUYQBT6Do11Z0oljQvphdHLd82Lr0Fm9PJvVNcgj0Pw1o6JniKMdjETU39AAWj1uIvA",
	"VQ69VBUSf7/oMlVIhBVo5nirCl+8K6UK4HIIkHy268dgJPcvAbM1oTJ3AAvJDz0IhI19+FnDkqXsp6Sr",
	"wImzNgl91eagi5NrzbfDzLMux6CxSCILJXrnCTeVCFryYgb6M+g/tVb6YRXcgDF8NWKXNQru1o8BvQMX",
	"OXSRhdeD/1rhS1XJ/DsD/lphZHEFIM/W6ssdbVRKv3F5HMjbcXV/rb7YrO1DsEf6c56/c3eVO4ACxFN/",
	"dw6Pg0NbsXFwwo2kiz8A7s31WaDGfrt1HrM18Bwc9BngoxdKXQsIem3zs25vUQ+7ZTqRDmWci+wDL0Ru",
	"I7BbYN+uOkG23rXjPwl8F9z9YJ9LXuFaafEPfG/10oc2KJl13ITfdkPD7BrZG4ytLf0WYbxV2yuMN+k1",
	"ESeYhbuL8Q78JmO8ld99nFCom4v8HvEGCeLt86NCe4k/HkbbWJyDYV9oARjvmwN6fMsj/MbO230jWyGR",
	"szjYDzXB26P8jYtvuC9jZiCrtMDtjDbcLlo6kJ5VuKZft04zMEYoGdl/YybonVvPYiZtdjHX1XSEluIv",
	"2LoiIORSDZ0ilwZ5dh19qkBvo1LzDEUG0bO3U8NiZqrNhhrnlLEOl7uIx+wzaOO8PHk8Id5UCZKXgqXs",
	"18f0ilIY1xZYQjUnoc6Lfq3ACkI6WU2nOUvZK0CCRu0Wu9Uk/jKZ7DsP2nVJr5OsY/bbGKNDLYGvEUsv",
	"Fz4drwCjJlLkK0PpQAjZgowcWGNvHzYjlQngfauMBexuKSz2hlHb/YF786rEH1bV51A2uH7VMbsYbxi4",
	"lN437/ThyLJ1iPZ5OY72eXku7buZ2tm0e9fJU0jvm9071ZagfUx/djUZZgPG+zVmty4iaaKKXIYl+dB3",
	"+EOafdLsiIoOadSOYJpqu39GMnMX1nhYkN83pf50IgPTJEvmk+Om++/e98Br/8y9XNQ9oj2WPJKbSRRd",
	"0YIJ/0IDiWPP/z3cUsp35J6a5f158HmJPmjuRud62PLbEtaTKCBsu32SG5HXTuACEIZSu1nwQam7cfFZ",
	"Gykwbb4Tvi8mF8c9hEeGD66Wx3JwGwYrXDdOO1zfzlJlMFj832nS0hsujFzzDaCdtV3euGaFmoGuVbHd",
	"U9dOoa4gPjCGW9hmIgv0RfMyP1pteXMY/ai232CmeQKHarV1RZ9wqVbpgqVsjVimSVKojBdrZTB9Onk6",
	"YeS3sb+dRcRVBDIvlZDYZam9T9XxoG2296bhchdUvaj/DQAA//9b7FpvaSAAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
