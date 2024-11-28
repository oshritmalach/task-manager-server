package handler

import (
	"Checkmarx/models"
	"Checkmarx/repository"
	"Checkmarx/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func startTestServer() *mux.Router {

	repo := repository.NewTaskManager()
	service := service.NewTaskService(repo)
	handler := TaskHandler{service: service}
	r := mux.NewRouter()

	r.HandleFunc("/task", handler.AddTask).Methods("POST")
	r.HandleFunc("/tasks", handler.GetAllTasks).Methods("GET")
	r.HandleFunc("/task/{id}", handler.GetTask).Methods("GET")
	r.HandleFunc("/task/{id}", handler.UpdateTask).Methods("POST")
	r.HandleFunc("/task/{id}", handler.DeleteTask).Methods("DELETE")
	return r
}

func TestAddTaskHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	requestBody := `{
        "title": "New Task",
        "description": "Description of the task",
        "status": "pending"
    }`

	req, err := http.NewRequest("POST", server.URL+"/task", strings.NewReader(requestBody))
	assert.NoError(t, err, "failed to create request")
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "failed to make request")
	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode, "expected status code %d, got %d", http.StatusCreated, resp.StatusCode)

}
func TestGetAllTasksHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL+"/tasks", nil)
	assert.NoError(t, err, "failed to create request")

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "failed to make request")
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "expected status code %d, got %d", http.StatusOK, resp.StatusCode)

}
func TestGetTaskIfNotExistsHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL+"/task/1", nil)
	assert.NoError(t, err, "failed to create request")

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "failed to make request")

	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)

}
func TestAddAndGetTaskHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	requestBody := `{"title": "New Task", "description": "Task description", "status": "pending"}`
	req, err := http.NewRequest("POST", server.URL+"/task", strings.NewReader(requestBody))
	assert.NoError(t, err, "failed to create request")

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "failed to make request")

	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode, "expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	req, err = http.NewRequest("GET", server.URL+"/task/1", nil)
	assert.NoError(t, err, "failed to create request")

	resp, err = http.DefaultClient.Do(req)
	assert.NoError(t, err, "failed to make request")

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode, "expected status code %d, got %d", http.StatusOK, resp.StatusCode)

	var task models.Task
	err = json.NewDecoder(resp.Body).Decode(&task)

	assert.NoError(t, err, "failed to decode response")
	assert.Equal(t, "New Task", task.Title, "expected task title 'New Task', got '%s'", task.Title)
	assert.Equal(t, "Task description", task.Description, "expected task description 'Task description', got '%s'", task.Description)
	assert.Equal(t, "pending", task.Status, "expected task status 'pending', got '%s'", task.Status)

}
func TestDeleteIfNotExistsHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	req, err := http.NewRequest("DELETE", server.URL+"/task/1", nil)
	assert.NoError(t, err, "failed to create request")

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "failed to make request")
	defer resp.Body.Close()
	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)

}
func TestAddAndDeleteTaskHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	requestBody := `{"title": "New Task", "description": "Task description", "status": "pending"}`
	req, err := http.NewRequest("POST", server.URL+"/task", strings.NewReader(requestBody))
	assert.NoError(t, err, "failed to create request")

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "failed to make request")

	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode, "expected status code %d, got %d", http.StatusCreated, resp.StatusCode)

	req, err = http.NewRequest("DELETE", server.URL+"/task/1", nil)
	assert.NoError(t, err, "failed to create request")

	resp, err = http.DefaultClient.Do(req)
	assert.NoError(t, err, "failed to make request")

	defer resp.Body.Close()
	assert.Equal(t, http.StatusNoContent, resp.StatusCode, "expected status code %d, got %d", http.StatusNoContent, resp.StatusCode)

}
func TestUpdateIfNotExistsHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	requestBody := `{"title": "New Task", "description": "Task description", "status": "pending"}`
	req, err := http.NewRequest("POST", server.URL+"/task/1", strings.NewReader(requestBody))
	assert.NoError(t, err, "failed to create request")

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "failed to make request")
	defer resp.Body.Close()
	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)

}
func TestAddAndUpdateTaskHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	requestBody := `{"title": "New Task", "description": "Task description", "status": "pending"}`
	req, err := http.NewRequest("POST", server.URL+"/task", strings.NewReader(requestBody))
	assert.NoError(t, err, "failed to create request")

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "failed to make request")

	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode, "expected status code %d, got %d", http.StatusCreated, resp.StatusCode)

	UpdatedRequestBody := `{"title": "New Task", "description": "Task description", "status": "pending"}`
	req, err = http.NewRequest("POST", server.URL+"/task/1", strings.NewReader(UpdatedRequestBody))
	assert.NoError(t, err, "failed to create request")

	resp, err = http.DefaultClient.Do(req)
	assert.NoError(t, err, "failed to make request")
	defer resp.Body.Close()

	var updatedTask models.Task
	err = json.NewDecoder(resp.Body).Decode(&updatedTask)
	assert.NoError(t, err, "failed to decode response")

	assert.Equal(t, "New Task", updatedTask.Title, "Task title does not match")
	assert.Equal(t, "Task description", updatedTask.Description, "Task description does not match")
	assert.Equal(t, "pending", updatedTask.Status, "Task status does not match")

}
