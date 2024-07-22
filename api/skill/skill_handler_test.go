package skill

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestGetSkillHandler(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   string
		mockStorage    *mockSkillStorage
	}{
		{
			name:           "get skill success",
			url:            "/skills/python",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status": "success", "data": {"key": "python", "name": "Python", "description": "Python is a programming language that lets you work quickly and integrate systems more effectively.", "logo": "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1200px-Python-logo-notext.svg.png", "tags": ["programming", "scripting", "web", "data science"]}}`,
			mockStorage: &mockSkillStorage{
				skill: &Skill{
					Key:         "python",
					Name:        "Python",
					Description: "Python is a programming language that lets you work quickly and integrate systems more effectively.",
					Logo:        "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1200px-Python-logo-notext.svg.png",
					Tags:        []string{"programming", "scripting", "web", "data science"},
				},
			},
		},
		{
			name:           "not exist skill",
			url:            "/skills/python3",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"status": "error", "message": "Skill not found" }`,
			mockStorage: &mockSkillStorage{
				errGet: sql.ErrNoRows,
			},
		},
		{
			name:           "database connection error",
			url:            "/skills/python4",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"status": "error", "message": "not be able to get skill" }`,
			mockStorage: &mockSkillStorage{
				errGet: sql.ErrConnDone,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			res := httptest.NewRecorder()
			c, r := gin.CreateTestContext(res)
			c.Request = httptest.NewRequest(http.MethodGet, tt.url, nil)

			h := NewSkillHandler(tt.mockStorage, nil)
			r.GET("/skills/:key", h.GetSkill) // Call to a handler method
			r.ServeHTTP(res, c.Request)

			// Assert response
			if status := res.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// Parse and compare JSON
			var actual, expectedJSON map[string]interface{}
			err := json.Unmarshal(res.Body.Bytes(), &actual)
			if err != nil {
				t.Fatalf("could not unmarshal response body: %v", err)
			}

			err = json.Unmarshal([]byte(tt.expectedBody), &expectedJSON)
			if err != nil {
				t.Fatalf("could not unmarshal expected JSON: %v", err)
			}

			// Assert response body
			if !reflect.DeepEqual(expectedJSON, actual) {
				t.Errorf("handler returned unexpected body: got %v want %v", actual, expectedJSON)
			}
		})
	}
}

func TestGetSkillsHandler(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   string
		mockStorage    *mockSkillStorage
	}{
		{
			name:           "get skills success",
			url:            "/skills",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status": "success", "data": [{"key": "python", "name": "Python", "description": "Python is a programming language that lets you work quickly and integrate systems more effectively.", "logo": "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1200px-Python-logo-notext.svg.png", "tags": ["programming", "scripting", "web", "data science"]}, {"key": "go", "name": "Go", "description": "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.", "logo": "https://blog.golang.org/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg", "tags": ["programming", "web", "cloud", "concurrency"]}]}`,
			mockStorage: &mockSkillStorage{
				skills: []Skill{
					{
						Key:         "python",
						Name:        "Python",
						Description: "Python is a programming language that lets you work quickly and integrate systems more effectively.",
						Logo:        "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1200px-Python-logo-notext.svg.png",
						Tags:        []string{"programming", "scripting", "web", "data science"},
					},
					{
						Key:         "go",
						Name:        "Go",
						Description: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.",
						Logo:        "https://blog.golang.org/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg",
						Tags:        []string{"programming", "web", "cloud", "concurrency"},
					},
				},
			},
		},
		{
			name:           "empty skills success",
			url:            "/skills",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status": "success", "data": []}`,
			mockStorage: &mockSkillStorage{
				skills: []Skill{},
			},
		},
		{
			name:           "database connection error",
			url:            "/skills",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"status": "error", "message": "not be able to get skills"}`,
			mockStorage: &mockSkillStorage{
				errGet: sql.ErrConnDone,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			res := httptest.NewRecorder()
			c, r := gin.CreateTestContext(res)
			c.Request = httptest.NewRequest(http.MethodGet, tt.url, nil)

			h := NewSkillHandler(tt.mockStorage, nil)
			r.GET("/skills", h.GetSkills) // Call to a handler method
			r.ServeHTTP(res, c.Request)

			// Assert response
			if status := res.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// Parse and compare JSON
			var actual, expectedJSON map[string]interface{}
			err := json.Unmarshal(res.Body.Bytes(), &actual)
			if err != nil {
				t.Fatalf("could not unmarshal response body: %v", err)
			}

			err = json.Unmarshal([]byte(tt.expectedBody), &expectedJSON)
			if err != nil {
				t.Fatalf("could not unmarshal expected JSON: %v", err)
			}

			// Assert response body
			if !reflect.DeepEqual(expectedJSON, actual) {
				t.Errorf("handler returned unexpected body: got %v want %v", actual, expectedJSON)
			}
		})
	}
}

func TestCreateSkillHandler(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		payload        string
		expectedStatus int
		expectedBody   string
		mockStorage    *mockSkillStorage
		mockSkillQueue *mockSkillQueue
	}{
		{
			name:           "create skill success",
			url:            "/skills",
			payload:        `{"key": "python", "name": "Python", "description": "Python is a programming language that lets you work quickly and integrate systems more effectively.", "logo": "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1200px-Python-logo-notext.svg.png", "tags": ["programming", "scripting", "web", "data science"]}`,
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"status": "success", "message": "creating skill already in progress"}`,
			mockStorage:    &mockSkillStorage{},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "invalid request",
			url:            "/skills",
			payload:        `{"key": "python", "name": "Python", "description": "Python is a programming language that lets you work quickly and integrate systems more effectively.", "logo": "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1200px-Python-logo-notext.svg.png"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"status": "error", "message": "invalid request"}`,
			mockStorage:    &mockSkillStorage{},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "skill already exists",
			url:            "/skills",
			payload:        `{"key": "python", "name": "Python", "description": "Python is a programming language that lets you work quickly and integrate systems more effectively.", "logo": "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1200px-Python-logo-notext.svg.png", "tags": ["programming", "scripting", "web", "data science"]}`,
			expectedStatus: http.StatusConflict,
			expectedBody:   `{"status": "error", "message": "skill already exists"}`,
			mockStorage: &mockSkillStorage{
				skill: &Skill{
					Key: "python",
				},
			},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "database connection error",
			url:            "/skills",
			payload:        `{"key": "python", "name": "Python", "description": "Python is a programming language that lets you work quickly and integrate systems more effectively.", "logo": "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1200px-Python-logo-notext.svg.png", "tags": ["programming", "scripting", "web", "data science"]}`,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"status": "error", "message": "not be able to get skill"}`,
			mockStorage: &mockSkillStorage{
				errGet: sql.ErrConnDone,
			},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "publish skill error",
			url:            "/skills",
			payload:        `{"key": "python", "name": "Python", "description": "Python is a programming language that lets you work quickly and integrate systems more effectively.", "logo": "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1200px-Python-logo-notext.svg.png", "tags": ["programming", "scripting", "web", "data science"]}`,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"status": "error", "message": "not be able to create skill"}`,
			mockStorage:    &mockSkillStorage{},
			mockSkillQueue: &mockSkillQueue{errPublish: errors.New("publish error")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			res := httptest.NewRecorder()
			c, r := gin.CreateTestContext(res)
			c.Request = httptest.NewRequest(http.MethodPost, tt.url, nil)
			c.Request.Body = ioutil.NopCloser(strings.NewReader(tt.payload))

			h := NewSkillHandler(tt.mockStorage, tt.mockSkillQueue)
			r.POST("/skills", h.CreateSkill) // Call to a handler method
			r.ServeHTTP(res, c.Request)

			// Assert response
			if status := res.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// Parse and compare JSON
			var actual, expectedJSON map[string]interface{}
			err := json.Unmarshal(res.Body.Bytes(), &actual)
			if err != nil {
				t.Fatalf("could not unmarshal response body: %v", err)
			}

			err = json.Unmarshal([]byte(tt.expectedBody), &expectedJSON)
			if err != nil {
				t.Fatalf("could not unmarshal expected JSON: %v", err)
			}

			// Assert response body
			if !reflect.DeepEqual(expectedJSON, actual) {
				t.Errorf("handler returned unexpected body: got %v want %v", actual, expectedJSON)
			}
		})
	}
}

func TestUpdateSkillHandler(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		payload        string
		expectedStatus int
		expectedBody   string
		mockStorage    *mockSkillStorage
		mockSkillQueue *mockSkillQueue
	}{
		{
			name:           "update skill success",
			url:            "/skills/python",
			payload:        `{"name": "Python", "description": "Python is a programming language that lets you work quickly and integrate systems more effectively.", "logo": "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1200px-Python-logo-notext.svg.png", "tags": ["programming", "scripting", "web", "data science"]}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status": "success", "message": "updating skill already in progress"}`,
			mockStorage: &mockSkillStorage{
				skill: &Skill{
					Key: "python",
				},
			},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "invalid request",
			url:            "/skills/python",
			payload:        `{"name": "Python", "description": "Python is a programming language that lets you work quickly and integrate systems more effectively.", "logo": "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1200px-Python-logo-notext.svg.png"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"status": "error", "message": "invalid request"}`,
			mockStorage:    &mockSkillStorage{},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "not exist skill",
			url:            "/skills/python3",
			payload:        `{"name": "Python", "description": "Python is a programming language that lets you work quickly and integrate systems more effectively.", "logo": "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1200px-Python-logo-notext.svg.png", "tags": ["programming", "scripting", "web", "data science"]}`,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"status": "error", "message": "skill not found"}`,
			mockStorage: &mockSkillStorage{
				errGet: sql.ErrNoRows,
			},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "database connection error",
			url:            "/skills/python4",
			payload:        `{"name": "Python", "description": "Python is a programming language that lets you work quickly and integrate systems more effectively.", "logo": "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1200px-Python-logo-notext.svg.png", "tags": ["programming", "scripting", "web", "data science"]}`,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"status": "error", "message": "not be able to get skill"}`,
			mockStorage: &mockSkillStorage{
				errGet: sql.ErrConnDone,
			},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "publish skill error",
			url:            "/skills/python5",
			payload:        `{"name": "Python", "description": "Python is a programming language that lets you work quickly and integrate systems more effectively.", "logo": "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1200px-Python-logo-notext.svg.png", "tags": ["programming", "scripting", "web", "data science"]}`,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"status": "error", "message": "not be able to update skill"}`,
			mockStorage: &mockSkillStorage{
				skill: &Skill{
					Key: "python5",
				},
			},
			mockSkillQueue: &mockSkillQueue{errPublish: errors.New("publish error")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			res := httptest.NewRecorder()
			c, r := gin.CreateTestContext(res)
			c.Request = httptest.NewRequest(http.MethodPut, tt.url, nil)
			c.Request.Body = ioutil.NopCloser(strings.NewReader(tt.payload))

			h := NewSkillHandler(tt.mockStorage, tt.mockSkillQueue)
			r.PUT("/skills/:key", h.UpdateSkill) // Call to a handler method
			r.ServeHTTP(res, c.Request)

			// Assert response
			if status := res.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// Parse and compare JSON
			var actual, expectedJSON map[string]interface{}
			err := json.Unmarshal(res.Body.Bytes(), &actual)
			if err != nil {
				t.Fatalf("could not unmarshal response body: %v", err)
			}

			err = json.Unmarshal([]byte(tt.expectedBody), &expectedJSON)
			if err != nil {
				t.Fatalf("could not unmarshal expected JSON: %v", err)
			}

			// Assert response body
			if !reflect.DeepEqual(expectedJSON, actual) {
				t.Errorf("handler returned unexpected body: got %v want %v", actual, expectedJSON)
			}
		})
	}
}

func TestDeleteSkillHandler(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   string
		mockStorage    *mockSkillStorage
		mockSkillQueue *mockSkillQueue
	}{
		{
			name:           "delete skill success",
			url:            "/skills/python",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status": "success", "message": "deleting skill already in progress"}`,
			mockStorage: &mockSkillStorage{
				skill: &Skill{
					Key: "python",
				},
			},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "not exist skill",
			url:            "/skills/python3",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"status": "error", "message": "skill not found"}`,
			mockStorage: &mockSkillStorage{
				errGet: sql.ErrNoRows,
			},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "database connection error",
			url:            "/skills/python4",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"status": "error", "message": "not be able to get skill"}`,
			mockStorage: &mockSkillStorage{
				errGet: sql.ErrConnDone,
			},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "publish skill error",
			url:            "/skills/python5",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"status": "error", "message": "not be able to delete skill"}`,
			mockStorage: &mockSkillStorage{
				skill: &Skill{
					Key: "python5",
				},
			},
			mockSkillQueue: &mockSkillQueue{errPublish: errors.New("publish error")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			res := httptest.NewRecorder()
			c, r := gin.CreateTestContext(res)
			c.Request = httptest.NewRequest(http.MethodDelete, tt.url, nil)

			h := NewSkillHandler(tt.mockStorage, tt.mockSkillQueue)
			r.DELETE("/skills/:key", h.DeleteSkill) // Call to a handler method
			r.ServeHTTP(res, c.Request)

			// Assert response
			if status := res.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// Parse and compare JSON
			var actual, expectedJSON map[string]interface{}
			err := json.Unmarshal(res.Body.Bytes(), &actual)
			if err != nil {
				t.Fatalf("could not unmarshal response body: %v", err)
			}

			err = json.Unmarshal([]byte(tt.expectedBody), &expectedJSON)
			if err != nil {
				t.Fatalf("could not unmarshal expected JSON: %v", err)
			}

			// Assert response body
			if !reflect.DeepEqual(expectedJSON, actual) {
				t.Errorf("handler returned unexpected body: got %v want %v", actual, expectedJSON)
			}
		})
	}
}

func TestUpdateSkillNameHandler(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		payload        string
		expectedStatus int
		expectedBody   string
		mockStorage    *mockSkillStorage
		mockSkillQueue *mockSkillQueue
	}{
		{
			name:           "update skill name success",
			url:            "/skills/python/name",
			payload:        `{"name": "Python"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status": "success", "message": "updating skill name already in progress"}`,
			mockStorage: &mockSkillStorage{
				skill: &Skill{
					Key: "python",
				},
			},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "invalid request",
			url:            "/skills/python/name",
			payload:        `{"name": ""}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"status": "error", "message": "invalid request"}`,
			mockStorage:    &mockSkillStorage{},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "not exist skill",
			url:            "/skills/python3/name",
			payload:        `{"name": "Python"}`,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"status": "error", "message": "skill not found"}`,
			mockStorage: &mockSkillStorage{
				errGet: sql.ErrNoRows,
			},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "database connection error",
			url:            "/skills/python4/name",
			payload:        `{"name": "Python"}`,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"status": "error", "message": "not be able to get skill"}`,
			mockStorage: &mockSkillStorage{
				errGet: sql.ErrConnDone,
			},
			mockSkillQueue: &mockSkillQueue{},
		},
		{
			name:           "publish skill error",
			url:            "/skills/python5/name",
			payload:        `{"name": "Python"}`,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"status": "error", "message": "not be able to update skill name"}`,
			mockStorage: &mockSkillStorage{
				skill: &Skill{
					Key: "python5",
				},
			},
			mockSkillQueue: &mockSkillQueue{errPublish: errors.New("publish error")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			res := httptest.NewRecorder()
			c, r := gin.CreateTestContext(res)
			c.Request = httptest.NewRequest(http.MethodPut, tt.url, nil)
			c.Request.Body = ioutil.NopCloser(strings.NewReader(tt.payload))

			h := NewSkillHandler(tt.mockStorage, tt.mockSkillQueue)
			r.PUT("/skills/:key/name", h.UpdateName) // Call to a handler method
			r.ServeHTTP(res, c.Request)

			// Assert response
			if status := res.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// Parse and compare JSON
			var actual, expectedJSON map[string]interface{}
			err := json.Unmarshal(res.Body.Bytes(), &actual)
			if err != nil {
				t.Fatalf("could not unmarshal response body: %v", err)
			}

			err = json.Unmarshal([]byte(tt.expectedBody), &expectedJSON)
			if err != nil {
				t.Fatalf("could not unmarshal expected JSON: %v", err)
			}

			// Assert response body
			if !reflect.DeepEqual(expectedJSON, actual) {
				t.Errorf("handler returned unexpected body: got %v want %v", actual, expectedJSON)
			}
		})
	}
}
