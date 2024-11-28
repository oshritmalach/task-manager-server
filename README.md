
# Task-manager-server

## Overview
The task-manager-server is a simple RESTapi service for managing tasks. It supports creating, updating, deleting, and retrieving tasks. The service is implemented in Go.

## Features
- Add new tasks with required fields: `Title`, `Description`, and `Status`.
- Update existing task.
- Delete tasks.
- Get a task by ID or get list of tasks.

---

## Getting Started

### Docker Instructions
#### 1. Use Docker Compose

```bash
docker-compose up --build
```

#### 2. Build and Run with Docker
Build the Docker image and run the container:
```bash
docker build -t task-manager-server .
docker run -d -p 8083:8083 task-manager-server
```


### Clone the Repository
   ```bash
   git clone https://github.com/oshritmalach/task-manager-server.git
   cd task-manager-server
   ```

### 2. Run Locally (Optional)
Install dependencies and start the server:
```bash
  go mod tidy
  go run main.go
```

#### Run the tests with:
   ```bash
    go test ./repository ./api/handler -v
   ```
### Using the Makefile

   ```bash
   make
   ```

## Usage

### API Endpoints
| Method | Endpoint               | Description         | Example Request Body                           |
|--------|-------------------------|---------------------|-----------------------------------------------|
| GET    | `/tasks`               | Get all tasks       | -                                             |
| POST   | `/task`                | Create a new task   | `{ "title": "Task 1", "description": "Test description", "status": "open" }` |
| GET    | `/task/{id}`           | Get a task by ID    | -                                             |
| POST   | `/task/{id}`           | Update a task by ID | `{ "title": "Updated Title" }`                |
| DELETE | `/task/{id}`           | Delete a task by ID | -                                             |

### Example Request and Response

#### Create a Task
**Request:**
```bash
curl -X POST http://localhost:8083/task \
-H "Content-Type: application/json" \
-d '{"title": "My Task", "description": "This is a test task", "status": "open"}'
```

**Response:**
```
Status Code: 201 Created
```
#### Update a Task
**Request:**
```bash
curl -X POST http://localhost:8083/task/1 \
-H "Content-Type: application/json" \
-d '{"status": "in_progress"}'
```

**Response:**
```
Status Code: 200 Ok
```
```json
{
   "title": "2 updated",
   "description": "This is a test task",
   "status": "in_progress",
   "createdAt": "2024-11-26T12:00:00Z"
}
```

#### Delete a Task
**Request:**
```bash
curl -X DELETE http://localhost:8083/task/1
```

**Response:**
```
Status Code: 204 No Content
```
#### Get a Task
```bash
curl -X GET http://localhost:8083/task/1
```
**Response:**
```
Status Code: 200 Ok
```
```json
{
   "title": "Item",
   "description": "3333",
   "status": "pending",
   "created_at": "2024-11-26T13:16:40.974907+02:00"
}
```
#### Get all Tasks
```bash
curl -X GET http://localhost:8083/tasks
```
**Response:**
```
Status Code: 200 Ok
```
```json
{
   "1": {
      "title": "Item",
      "description": "3333",
      "status": "pending",
      "created_at": "2024-11-26T13:16:40.974907+02:00"
   }
}
```

### Error Handling

All failed requests return a response in a unified JSON structure containing the following fields:

- **`code`**: The HTTP status code indicating the type of error (e.g., 400, 404, etc.).
- **`message`**: A descriptive message explaining the reason for the error.

#### Example Error Response:
```json
{
    "code": 404,
    "message": "task with ID 1 does not exist"
}
```



