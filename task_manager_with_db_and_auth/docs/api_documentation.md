# Task Management REST API
**Introduction**
The Task Management REST API is designed to manage tasks within a system. It supports operations such as creating, updating, retrieving, and deleting tasks. It also handles user registration and login, utilizing JWT for user authentication. There are two types of users: Normal users and Admins. Admins have additional privileges, such as activating/deactivating users and registering other admins. A middleware using JWT ensures that users have appropriate access based on their roles.

# Summary of Clean Architecture Usage
**Main Components**
1.  Entities (Domain Layer)

* `Domain/domain.go`: Defines the core domain models, including entities such as Task and User. This layer contains the fundamental business logic and entities, which are agnostic to any frameworks or external systems.

2. Use Cases (Application Layer)

* `Usecases/task_usecases.go`: Contains the application-specific business rules for managing tasks. Implements use cases such as creating, updating, and retrieving tasks.
* `Usecases/user_usecases.go`: Contains the application-specific business rules for managing users. Implements use cases such as user registration, login, and role management.

3. Interface Adapters (Adapter Layer)

* `Delivery/handler/task_handler.go`: Provides HTTP handlers for managing tasks. These handlers interact with the TaskUsecase to perform actions and return responses.
* `Delivery/handler/user_handler.go`: Provides HTTP handlers for user management. These handlers interact with the UserUsecase to perform actions and return responses.
* `Delivery/routers/router.go`: Configures the API routes and links them to the appropriate handlers. This layer adapts the routing configuration to the handlers.

4. Frameworks and Drivers (Infrastructure Layer)

`Infrastructure/auth_middleWare.go`: Implements middleware for JWT authentication and authorization. Ensures that routes are protected and only accessible by authorized users.
* `Infrastructure/jwt_service.go`: Contains Functions to generate and validate JWT tokens.
* `Infrastructure/password_service.go`: Contains Functions for hashing and comparing passwords to ensure secure storage of user credentials.
* `Repositories/task_repository.go`: Handles data access for tasks, including CRUD operations with the database.
* `Repositories/user_repository.go`: Handles data access for users and admins.

**System Requirements**

- Go 1.22 or higher
- A terminal or command-line interface
- Postman (optional) for testing
- MongoDB (local or cloud)
- Use `go get` to install required packages.
**Running the Application**

1. Install MongoDB locally or obtain a URI from [MongoDB Atlas](https://cloud.mongodb.com/) . Add this URI or `mongodb://localhost:27017` to the `.env` file located in the same directory as `main.go`.

2. Set your `JWT_SECRET` in the `.env` file:

```sh
MONGO_URI=your_mongodb_uri
JWT_SECRET=your_jwt_secret
```
3. Install the required Go packages:

```sh

go get github.com/gin-gonic/gin
go get github.com/joho/godotenv
go get github.com/dgrijalva/jwt-go
go get go.mongodb.org/mongo-driver
go get go.mongodb.org/mongo-driver/bson
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
go get golang.org/x/crypto/bcrypt
```
4. Navigate to the main.go file and run the application:

```sh
go run main.go
```

**Endpoints**

1. Get Tasks

- Description: Retrieves all tasks.
- Method: GET
- Endpoint: /tasks
- Response:
```json
[
  {
    "id": "1",
    "title": "Task Title 1",
    "description": "Task description 1",
    "status": "Pending",
    "due_date": "2024-08-06T00:00:00Z"
  },
  {
    "id": "2",
    "title": "Task Title 2",
    "description": "Task description 2",
    "status": "Completed",
    "due_date": "2024-08-07T00:00:00Z"
  }
]
```
2. Get Specific Task

- Description: Retrieves details of a specific task by its ID.
- Method: GET
- Endpoint: /tasks/:id
- Parameters:
- id: The ID of the task to retrieve.
- Response:
- Success:
```json
{
  "id": "1",
  "title": "Task Title",
  "description": "Task description",
  "status": "Pending",
  "due_date": "2024-08-06T00:00:00Z"
}
```
- Error:
```json

{
  "error": "task not found"
}
```
3. Create Task

- Description: Adds a new task.
- Method: POST
- Endpoint: /tasks
- Input:
```json

{
  "title": "New Task",
  "description": "Task description"
}
```
- Response:
- - Success:
```json

{
  "message": "Task successfully added",
  "task_id": "new_task_id"
}
```
- Error:
```json
{
  "error": "failed to add task"
}
```
4. Update Task

- Description: Updates an existing task by its ID.
- Method: PUT
- Endpoint: /tasks/:id
- Parameters:
- id: The ID of the task to update.
- Input:
```json
{
  "title": "Updated Task Title",
  "description": "Updated task description"
}
```
- Response:
- - Success:
```json

{
  "message": "Task successfully updated"
}
```
- Error:
```json

{
  "error": "failed to update task"
}
```
5. Delete Task

- Description: Deletes a task by its ID.
- Method: DELETE
- Endpoint: /tasks/:id
- Parameters:
- -  `id`: The ID of the task to delete.
- Response:
- - Success:
```json

{
  "message": "Task successfully deleted"
}
```
- Error:
```json

{
  "error": "failed to delete task"
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


**Error Handling**
The API handles errors such as:

- Invalid task IDs
- Attempting to update or delete a non-existent task
- JWT Token validation
- Login only allowed for existing users

Errors are reported with descriptive messages to guide the user in correcting the issue.

**Code Structure**

`main.go`: The entry point of the application. Starts the server and calls functions in `router/router.go`.
`controllers/controller.go`: Provides functions to handle requests.
`models/task.go`: Defines the Task data model.
`models/user.go`: Defines the User data model.
`data/task_service.go`: Contains functions related to task management.
`data/user_service.go`: Contains functions related to user management, including login and registration.
`middleware/auth_middleware.go`: Manages JWT authentication and authorization.
`router/router.go`: Sets up and configures the API routes and links them to controller functions.
`docs/api_documentation.md`: Contains this documentation.
`.env`: Contains the MongoDB URI and JWT secret key.