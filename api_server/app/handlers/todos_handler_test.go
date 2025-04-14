package handlers

import (
	models "app/models/generated"
	apis "app/openapi"
	"app/test/factories"
	"net/http"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/oapi-codegen/testutil"
)

type testTodosHandlerSuite struct {
	WithDBSuite
}

func (s *testTodosHandlerSuite) SetupTest() {
	s.SetDBCon()

	s.initializeHandlers()

	s.SetCsrfHeaderValues()
}

func (s *testTodosHandlerSuite) TearDownTest() {
	s.CloseDB()
}

func (s *testTodosHandlerSuite) TestPostTodos_StatusOk() {
	s.SignIn()

	reqBody := apis.StoreTodoInput{
		Title: "test_title",
		Content: "test_content",
	}
	result := testutil.NewRequest().Post("/todos").WithHeader("Cookie", token+"; "+csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res apis.PostTodos200JSONResponse
	result.UnmarshalBodyToObject(&res)
	
	assert.Equal(s.T(), int64(http.StatusOK), res.Code)

	// NOTE: TODOリストが作成されていることを確認
	todo, err := models.Todos(
		qm.Where("title = ?", "test_title"),
	).One(ctx, DBCon)
	if err != nil {
		s.T().Fatalf("failed to create todo %v", err)
	}
	assert.Equal(s.T(), null.String{String: "test_content", Valid: true}, todo.Content)
}

func (s *testTodosHandlerSuite) TestPostTodos_BadRequest() {
	s.SignIn()

	reqBody := apis.StoreTodoInput{
		Title: "",
		Content: "test_content",
	}
	result := testutil.NewRequest().Post("/todos").WithHeader("Cookie", token+"; "+csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res apis.PostTodos200JSONResponse
	result.UnmarshalBodyToObject(&res)
	titleValidationErrors := *res.Errors.Title
	assert.Equal(s.T(), []string{"タイトルは必須入力です。"}, titleValidationErrors)
	
	assert.Equal(s.T(), int64(http.StatusOK), res.Code)

	// NOTE: TODOリストが作成されていないことを確認
	isExistTodo, _ := models.Todos(
		qm.Where("title = ?", ""),
	).Exists(ctx, DBCon)
	assert.False(s.T(), isExistTodo)
}

func (s *testTodosHandlerSuite) TestPostTodos_StatusUnauthorized() {
	reqBody := apis.StoreTodoInput{
		Title: "test_title",
		Content: "test_content",
	}
	result := testutil.NewRequest().Post("/todos").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())

	// NOTE: TODOリストが作成されていないことを確認
	isExistTodo, _ := models.Todos(
		qm.Where("title = ?", "test_title"),
	).Exists(ctx, DBCon)
	assert.False(s.T(), isExistTodo)
}

func (s *testTodosHandlerSuite) TestGetTodos_StatusOk() {
	s.SignIn()

	var todosSlice models.TodoSlice
	todosSlice = append(todosSlice, &models.Todo{
		Title:   "test title 1",
		Content: null.String{String: "test content 1", Valid: true},
		UserID:  int64(user.ID),
	})
	todosSlice = append(todosSlice, &models.Todo{
		Title:   "test title 2",
		Content: null.String{String: "test content 2", Valid: true},
		UserID:  int64(user.ID),
	})
	_, err := todosSlice.InsertAll(ctx, DBCon, boil.Infer())
	if err != nil {
		s.T().Fatalf("failed to create TestFetchTodosList Data: %v", err)
	}
	
	result := testutil.NewRequest().Get("/todos").WithHeader("Cookie", token+"; "+csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res apis.GetTodos200JSONResponse
	result.UnmarshalBodyToObject(&res)

	assert.Equal(s.T(), 2, len(res.Todos))
	assert.Equal(s.T(), "test title 1", res.Todos[0].Title)
	assert.Equal(s.T(), "test content 1", res.Todos[0].Content)
}

func (s *testTodosHandlerSuite) TestGetTodos_StatusUnauthorized() {
	result := testutil.NewRequest().Get("/todos").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
}

func (s *testTodosHandlerSuite) TestGetTodo_StatusOk() {
	s.SignIn()

	todoParam := map[string]interface{}{"UserID": int64(user.ID), "Title": "test title 1", "Content": null.String{String: "test content 1", Valid: true}}
	todo := factories.TodoFactory.MustCreateWithOption(todoParam).(*models.Todo)
	if err := todo.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test todo %v", err)
	}
	todo.Reload(ctx, DBCon)

	result := testutil.NewRequest().Get("/todos/"+strconv.Itoa(int(todo.ID))).WithHeader("Cookie", token+"; "+csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res apis.GetTodo200JSONResponse
	result.UnmarshalBodyToObject(&res)

	assert.Equal(s.T(), "test title 1", res.Todo.Title)
	assert.Equal(s.T(), "test content 1", res.Todo.Content)
}

func (s *testTodosHandlerSuite) TestGetTodo_StatusUnauthorized() {
	result := testutil.NewRequest().Get("/todos/1").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
}

func (s *testTodosHandlerSuite) TestGetTodo_StatusNotFound() {
	s.SignIn()

	todoParam := map[string]interface{}{"UserID": int64(user.ID), "Title": "test title 1", "Content": null.String{String: "test content 1", Valid: true}}
	todo := factories.TodoFactory.MustCreateWithOption(todoParam).(*models.Todo)
	if err := todo.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test todo %v", err)
	}
	todo.Reload(ctx, DBCon)
	
	result := testutil.NewRequest().Get("/todos/"+strconv.Itoa(int(todo.ID + 1))).WithHeader("Cookie", token+"; "+csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusNotFound, result.Code())
}

func (s *testTodosHandlerSuite) TestPatchTodo_StatusOk() {
	s.SignIn()

	todoParam := map[string]interface{}{"UserID": int64(user.ID), "Title": "test title 1", "Content": null.String{String: "test content 1", Valid: true}}
	todo := factories.TodoFactory.MustCreateWithOption(todoParam).(*models.Todo)
	if err := todo.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test todo %v", err)
	}
	todo.Reload(ctx, DBCon)

	reqBody := apis.StoreTodoInput{
		Title: "test updated title 1",
		Content: "test updated content 1",
	}
	result := testutil.NewRequest().Patch("/todos/"+strconv.Itoa(int(todo.ID))).WithHeader("Cookie", token+"; "+csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res apis.PatchTodo200JSONResponse
	result.UnmarshalBodyToObject(&res)
	
	assert.Equal(s.T(), int64(http.StatusOK), res.Code)

	// NOTE: TODOリストが更新されていることを確認
	if err := todo.Reload(ctx, DBCon); err != nil {
		s.T().Fatalf("failed to reload test todos %v", err)
	}
	assert.Equal(s.T(), "test updated title 1", todo.Title)
	assert.Equal(s.T(), null.String{String: "test updated content 1", Valid: true}, todo.Content)
}

func (s *testTodosHandlerSuite) TestPatchTodo_BadRequest() {
	s.SignIn()

	todoParam := map[string]interface{}{"UserID": int64(user.ID), "Title": "test title 1", "Content": null.String{String: "test content 1", Valid: true}}
	todo := factories.TodoFactory.MustCreateWithOption(todoParam).(*models.Todo)
	if err := todo.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test todo %v", err)
	}
	todo.Reload(ctx, DBCon)
	
	reqBody := apis.StoreTodoInput{
		Title: "",
		Content: "test updated content 1",
	}
	result := testutil.NewRequest().Patch("/todos/"+strconv.Itoa(int(todo.ID))).WithHeader("Cookie", token+"; "+csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res apis.PatchTodo200JSONResponse
	result.UnmarshalBodyToObject(&res)
	titleValidationErrors := *res.Errors.Title
	assert.Equal(s.T(), []string{"タイトルは必須入力です。"}, titleValidationErrors)
	
	assert.Equal(s.T(), int64(http.StatusOK), res.Code)

	// NOTE: TODOリストが更新されていないことを確認
	if err := todo.Reload(ctx, DBCon); err != nil {
		s.T().Fatalf("failed to reload test todos %v", err)
	}
	assert.Equal(s.T(), "test title 1", todo.Title)
	assert.Equal(s.T(), null.String{String: "test content 1", Valid: true}, todo.Content)
}

func (s *testTodosHandlerSuite) TestPatchTodo_StatusUnauthorized() {
	reqBody := apis.StoreTodoInput{
		Title: "test_title",
		Content: "test_content",
	}
	result := testutil.NewRequest().Patch("/todos/1").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
}

func (s *testTodosHandlerSuite) TestPatchTodo_StatusNotFound() {
	s.SignIn()

	todoParam := map[string]interface{}{"UserID": int64(user.ID), "Title": "test title 1", "Content": null.String{String: "test content 1", Valid: true}}
	todo := factories.TodoFactory.MustCreateWithOption(todoParam).(*models.Todo)
	if err := todo.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test todo %v", err)
	}
	todo.Reload(ctx, DBCon)
	
	reqBody := apis.StoreTodoInput{
		Title: "test updated title 1",
		Content: "test updated content 1",
	}
	result := testutil.NewRequest().Patch("/todos/"+strconv.Itoa(int(todo.ID + 1))).WithHeader("Cookie", token+"; "+csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusNotFound, result.Code())

	// NOTE: TODOリストが更新されていないことを確認
	if err := todo.Reload(ctx, DBCon); err != nil {
		s.T().Fatalf("failed to reload test todos %v", err)
	}
	assert.Equal(s.T(), "test title 1", todo.Title)
	assert.Equal(s.T(), null.String{String: "test content 1", Valid: true}, todo.Content)
}

func (s *testTodosHandlerSuite) TestDeleteTodo_StatusOk() {
	s.SignIn()

	todoParam := map[string]interface{}{"UserID": int64(user.ID), "Title": "test title 1", "Content": null.String{String: "test content 1", Valid: true}}
	todo := factories.TodoFactory.MustCreateWithOption(todoParam).(*models.Todo)
	if err := todo.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test todo %v", err)
	}
	todo.Reload(ctx, DBCon)

	result := testutil.NewRequest().Delete("/todos/"+strconv.Itoa(int(todo.ID))).WithHeader("Cookie", token+"; "+csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res apis.DeleteTodo200JSONResponse
	result.UnmarshalBodyToObject(&res)
	
	assert.Equal(s.T(), int64(http.StatusOK), res.Code)
	assert.Equal(s.T(), true, res.Result)

	// NOTE: TODOリストが削除されていることを確認
	err := todo.Reload(ctx, DBCon)
	assert.NotNil(s.T(), err)
}

func (s *testTodosHandlerSuite) TestDeleteTodo_StatusUnauthorized() {
	result := testutil.NewRequest().Delete("/todos/1").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
}

func (s *testTodosHandlerSuite) TestDeleteTodo_StatusNotFound() {
	s.SignIn()

	todoParam := map[string]interface{}{"UserID": int64(user.ID), "Title": "test title 1", "Content": null.String{String: "test content 1", Valid: true}}
	todo := factories.TodoFactory.MustCreateWithOption(todoParam).(*models.Todo)
	if err := todo.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test todo %v", err)
	}
	todo.Reload(ctx, DBCon)
	
	result := testutil.NewRequest().Delete("/todos/"+strconv.Itoa(int(todo.ID + 1))).WithHeader("Cookie", token+"; "+csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusNotFound, result.Code())

	// NOTE: TODOリストが削除されていないことを確認
	err := todo.Reload(ctx, DBCon)
	assert.Nil(s.T(), err)
}

func TestTodosHandler(t *testing.T) {
	// テストスイートを実施
	suite.Run(t, new(testTodosHandlerSuite))
}
