# Task Management REST API

## Introduction
The Task Management REST API is designed to manage tasks within a system. It supports operations such as creating, updating, retrieving, and deleting tasks. It is connected to a MongoDB database online.

## System Requirements
- Go 1.22 or higher
- A terminal or command-line interface
- Postman (optional) to test the API

## Running the Application
- Before starting, ensure that MongoDB is installed on your PC or get a URI from the official website: [MongoDB Cloud](https://cloud.mongodb.com/). Then, add that URI or `mongodb://localhost:27017` to the `.env` file in the same directory as `main.go`.

- Follow these steps:

1. Navigate to the `main.go` file in your terminal or command prompt.
2. Run the application using the following command:
    ```sh
    go run main.go
    ```
   This will start the server, and you can interact with the API via HTTP requests.

## Endpoints

### 1. **Get Tasks**
   - **Description:** Retrieves all tasks.
   - **Method:** GET
   - **Endpoint:** `/tasks`
   - **Response:**
     - **Success:** 
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         [
           {
             "id": "1",
             "title": "Complete Documentation",
             "description": "Finish writing the API documentation",
             "due_date": "created time",
             "status": "pending"
           },
           {
             "id": "2",
             "title": "post Documentation",
             "description": "Finish writing the API documentation and post",
             "due_date": "created time",
             "status": "pending"
           }
         ]
         ```
     - **Error:**
       - **Status Code:** `500 Internal Server Error`
       - **Message:** "error while retrieving the tasks!"

### 2. **Create Task**
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
       - **Status Code:** `201 Created`
       - **Example:**
         ```json
         {
           "message": "Task created"
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Message:** "Invalid input data."
       - **Example:**
         ```json
         {
           "error": "Invalid input data."
         }
         ```

### 3. **Update Task**
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
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "Task updated"
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Message:** "Invalid task ID or input data."
       - **Example:**
         ```json
         {
           "message": "Invalid task ID or input data."
         }
         ```
       - **Status Code:** `404 Not Found`
       - **Message:** "Task not found"
       - **Example:**
         ```json
         {
           "message": "Task not found"
         }
         ```

### 4. **Delete Task**
   - **Description:** Deletes a task by its ID.
   - **Method:** DELETE
   - **Endpoint:** `/tasks/{id}`
   - **Response:**
     - **Success:** 
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "Task removed"
         }
         ```
     - **Error:**
       - **Status Code:** `404 Not Found`
       - **Message:** "Task not found with the provided ID."
       - **Example:**
         ```json
         {
           "message": "Task not found"
         }
         ```

## Error Handling
The API handles errors such as:
- **Invalid Task IDs:** Returns a `400 Bad Request` or `404 Not Found` with a descriptive error message.
- **Non-Existent Task Operations:** When trying to update or delete a task that does not exist, the API will return a `404 Not Found` status with an appropriate message.
- **Input Validation:** When creating or updating a task, if the input data is invalid (e.g., missing title or description), the API will return a `400 Bad Request` status with details on the required fields.

Errors are reported with descriptive messages to guide the user in correcting the issue.

## Code Structure
- **`main.go`:** The entry point of the application. It starts the process by calling a function in `router/router.go`.
- **`controllers/task_controller.go`:** Contains the logic for handling HTTP requests and responses.
- **`models/task.go`:** Defines the Task data model.
- **`data/task_service.go`:** Contains the service layer functions that interact with the repository.
- **`router/router.go`:** Sets up and configures the API routes, linking the endpoints to specific functions in the controller.
- **`docs/api_documentation.md`:** Contains this documentation.
- **`.env`:** Contains the MongoDB URI for database connection.

## Example Usage with cURL

### Get All Tasks
```sh
curl -X GET http://localhost:8080/tasks
```

### Create a New Task
```sh
curl -X POST http://localhost:8080/tasks -d '{"title":"New Task","description":"Task description"}' -H "Content-Type: application/json"
```
### Update Task

```sh
curl -X PUT http://localhost:8080/tasks/{id} -d '{"title":"Updated Task Title","description":"Updated Task Description"}' -H "Content-Type: application/json"
```
### Delete Task
```
curl -X DELETE http://localhost:8080/tasks/{id}
```