# Task Management REST API

## Introduction
The Task Management REST API is built following the Clean Architecture principles to manage tasks within a system. It supports operations such as creating, updating, retrieving, and deleting tasks. The API is connected to a MongoDB database.

## System Requirements
- Go 1.22 or higher
- A terminal or command-line interface
- Postman (optional) to test the API

## Running the Application
- Ensure that MongoDB is installed on your PC or get a URI from the official website: [MongoDB Cloud](https://cloud.mongodb.com/). Add that URI or `mongodb://localhost:27017` to the `.env` file located in the same directory as `main.go`.

- Follow these steps to run the application:

1. Navigate to the `main.go` file in your terminal or command prompt.
2. Start the server using the following command:
    ```sh
    go run main.go
    ```
   The server will start, and you can interact with the API via HTTP requests.

## Clean Architecture Overview
The application is structured based on the Clean Architecture, ensuring separation of concerns and easy maintenance:

### **1. Delivery Layer (`Delivery/`)**
   - Handles incoming HTTP requests and returns responses.
   - **`main.go`:** Sets up the HTTP server, initializes dependencies, and defines the routing configuration.
   - **`controllers/`:** Handles HTTP requests and invokes the appropriate use case methods.
   - **`routers/routers.go`:** Sets up routes and initializes the Gin router.

### **2. Domain Layer (`Domain/`)**
   - Contains the core business entities and logic.
   - **`domain.go`:** Defines entities like `Task` and `User`.

### **3. Use Case Layer (`Usecases/`)**
   - Contains application-specific business rules.
   - **`task_usecases.go`:** Implements use cases for tasks, including creating, updating, retrieving, and deleting tasks.
   - **`user_usecases.go`:** Implements use cases for user operations like registration and login.

### **4. Repository Layer (`Repositories/`)**
   - Abstracts data access logic, separating it from business logic.
   - **`task_repository.go`:** Provides an interface and implementation for task data access.
   - **`user_repository.go`:** Provides an interface and implementation for user data access.

### **5. Infrastructure Layer (`Infrastructure/`)**
   - Implements external dependencies and services.
   - **`auth_middleWare.go`:** Middleware for handling authentication and authorization using JWT tokens.
   - **`jwt_service.go`:** Provides functions to generate and validate JWT tokens.
   - **`password_service.go`:** Provides functions for hashing and comparing passwords to securely store user credentials.

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
             "title": "Post Documentation",
             "description": "Finish writing the API documentation and post",
             "due_date": "created time",
             "status": "pending"
           }
         ]
         ```
     - **Error:**
       - **Status Code:** `500 Internal Server Error`
       - **Message:** "Error while retrieving the tasks!"

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

## User Related Endpoints

### 1. **Register User**
   - **Description:** Registers a new user.
   - **Method:** POST
   - **Endpoint:** `/register`
   - **Input:** JSON object with user details.
     ```json
     {
       "username": "new_user",
       "password": "password123"
     }
     ```
   - **Response:**
     - **Success:** 
       - **Status Code:** `201 Created`
       - **Example:**
         ```json
         {
           "message": "User registered"
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

### 2. **User Login**
   - **Description:** Authenticates a user and provides a JWT token.
   - **Method:** POST
   - **Endpoint:** `/login`
   - **Input:** JSON object with login credentials.
     ```json
     {
       "username": "existing_user",
       "password": "password123"
     }
     ```
   - **Response:**
     - **Success:** 
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "token": "jwt_token_here"
         }
         ```
     - **Error:**
       - **Status Code:** `401 Unauthorized`
       - **Message:** "Invalid username or password."
       - **Example:**
         ```json
         {
           "error": "Invalid username or password."
         }
         ```

### 3. **Promote User to Admin**
   - **Description:** Promotes an existing user to an admin role.
   - **Method:** PUT
   - **Endpoint:** `/admin/promote/{id}`
   - **Input:** User ID in the URL path.
   - **Response:**
     - **Success:** 
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "User promoted to admin"
         }
         ```
     - **Error:**
       - **Status Code:** `404 Not Found`
       - **Message:** "User not found."
       - **Example:**
         ```json
         {
           "error": "User not found"
         }
         ```

### 4. **Activate User**
   - **Description:** Activates a user account.
   - **Method:** PUT
   - **Endpoint:** `/admin/activate/{id}`
   - **Input:** User ID in the URL path.
   - **Response:**
     - **Success:** 
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "User account activated"
         }
         ```
     - **Error:**
       - **Status Code:** `404 Not Found`
       - **Message:** "User not found."
       - **Example:**
         ```json
         {
           "error": "User not found"
         }
         ```

### 5. **Deactivate User**
   - **Description:** Deactivates a user account.
   - **Method:** PUT
   - **Endpoint:** `/admin/deactivate/{id}`
   - **Input:** User ID in the URL path.
   - **Response:**
     - **Success:** 
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "User account deactivated"
         }
         ```
     - **Error:**
       - **Status Code:** `404 Not Found`
       - **Message:** "User not found."
       - **Example:**
         ```json
         {
           "error": "User not found"
         }
         ```


## Error Handling
The API handles errors such as:
- **Invalid Task IDs:** Returns a `400 Bad Request` or `404 Not Found` with a descriptive error message.
- **Non-Existent Task Operations:** When trying to update or delete a task that does not exist, the API will return a `404 Not Found` status with an appropriate message.
- **Input Validation:** When creating or updating a task, if the input data is invalid (e.g., missing title or description), the API will return a `400 Bad Request` status with details on the required fields.

Errors are reported with descriptive messages to guide the user in correcting the issue.

## Code Structure Overview
This project follows the Clean Architecture approach:

- **`Delivery/`:** Handles the request-response cycle.
- **`Domain/`:** Contains the core business entities and logic.
- **`Usecases/`:** Implements the application-specific business rules.
- **`Repositories/`:** Abstracts the data access logic.
- **`Infrastructure/`:** Manages external dependencies and services.

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