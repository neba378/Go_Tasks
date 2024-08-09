# Task Management REST API

## Introduction
The Task Management REST API is designed to manage tasks within a system. It supports operations such as creating, updating, retrieving, and deleting tasks.

## System Requirements
- Go 1.22 or higher
- A terminal or command-line interface

## Running the Application
1. Navigate to the `main.go` file in your terminal or command prompt.
2. Run the application using the following command:
    ```sh
    go run main.go
    ```

   This will start the server, and you can interact with the API via HTTP requests.

## Endpoints

### 1. Get Tasks
- **Description:** Retrieves all tasks.
- **Method:** GET
- **Endpoint:** `/tasks`
- **Response:**
  - **Success:**
    ```json
    [
      {
        "id": "1",
        "title": "Task 1",
        "description": "First task description"
      },
      {
        "id": "2",
        "title": "Task 2",
        "description": "Second task description"
      }
    ]
    ```
  - **Error:**
    ```json
    {
      "error": "Failed to retrieve tasks"
    }
    ```

### 2. Get Specific Task
- **Description:** Retrieves a specific task by its ID.
- **Method:** GET
- **Endpoint:** `/tasks/{id}`
- **Response:**
  - **Success:**
    ```json
    {
      "id": "1",
      "title": "Task 1",
      "description": "First task description"
    }
    ```
  - **Error:**
    ```json
    {
      "message": "No such task found!"
    }
    ```

### 3. Create Task
- **Description:** Adds a new task.
- **Method:** POST
- **Endpoint:** `/tasks`
- **Input:** JSON object with task details.
  ```json
  {
    "title": "New Task",
    "description": "Task description"
  }
   ```
  - **Response:**
  - **Success:**
    ```json
    {
      "message": "Task created"
    }
    ```
  - **Error:**
    ```json
    {
      "error": "Failed to create task"
    }
    ```

### 4. Update Task
- **Description:** Updates an existing task by its ID.
- **Method:** PUT
- **Endpoint:** `/tasks/{id}`
- **Input:** JSON object with updated task details.
  
  ```json
  {
    "title": "Updated Task Title",
    "description": "Updated Task Description"
  }
   ```
  - **Response:**
  - **Success:**
    ```json
    {
      "message": "Task updated"
    }
    ```
  - **Error:**
    ```json
    {
      "message": "Task not found"
    }
    ```


### 5. Delete Task
- **Description:** Deletes a task by its ID.
- **Method:** DELETE
- **Endpoint:** `/tasks/{id}`
- **Response:**
  - **Success:**
    ```json
    {
      "message": "Task removed"
    }
    ```
  - **Error:**
    ```json
    {
      "message": "Task not found"
    }
    ```

### 6. Get Specific Task
- **Description:** Retrieves a specific task by its ID.
- **Method:** GET
- **Endpoint:** `/tasks/{id}`
- **Response:**
  - **Success:**
    ```json
    {
      "id": "1",
      "title": "Task 1",
      "description": "First task description"
    }
    ```
  - **Error:**
    ```json
    {
      "message": "No such task found!"
    }
    ```

## Error Handling
The API handles errors such as:
- Invalid task IDs
- Attempting to update or delete a non-existent task

Errors are reported with descriptive messages to guide the user in correcting the issue.

## Code Structure
- `main.go`: The entry point of the application. It starts the process by calling a function in `router/router.go`.
- `controllers/task_controller.go`: Provides functions to the router.
- `models/task.go`: Defines the Task data model.
- `data/task_service.go`: Provides functions to the controller.
- `router/router.go`: Sets up and configures the API routes, relating the endpoints with specific functions from the controller.
- `docs/api_documentation.md`: Contains this documentation.
