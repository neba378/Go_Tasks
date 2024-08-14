# Task Management REST API

### Introduction
The Task Management REST API is designed to manage tasks within a system. It supports operations such as creating, updating, retrieving, and deleting tasks. It is connected to a MongoDB database online.

### System Requirements
- Go 1.22 or higher
- A terminal or command-line interface
- Postman (optional) to test the API

### Running the Application
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


## Task Management REST API - Testing Documentation

### Overview

This project includes comprehensive testing for the repositories, controllers, and use cases using the `testify` library in Go. The testing is designed to ensure that each layer of the application functions correctly and adheres to the principles of Clean Architecture.

### Testing Libraries Used
- **Testify:** A toolkit with tools such as `assert`, `mock`, and `suite` to facilitate writing tests.
  - **Assert:** Used for making assertions in tests, verifying that expected outcomes match actual results.
  - **Mock:** Used to mock dependencies and isolate the unit being tested.
  - **Suite:** Provides a way to group related tests together and run setup/teardown code before/after tests.

## Test Structure

### 1. **Repository Tests**
Repository tests are designed to ensure that the data persistence layer interacts correctly with the MongoDB database.

- **File Structure:**
  - `Repositories/task_repository_test.go`
  - `Repositories/user_repository_test.go`

- **Tests Include:**
  - **CRUD operations:** Ensuring tasks and users can be created, read, updated, and deleted.
  - **Database interactions:** Verifying that queries are correctly executed and results are as expected.

- **Example Test:**
```go
  func (suite *TaskRepositoryTestSuite) TestAdd() {
	task := domain.Task{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "Pending"}
	err := suite.repo.Add(task)
	suite.NoError(err)

	var result domain.Task
	err = suite.collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: task.ID}}).Decode(&result)
	suite.NoError(err)
	suite.Equal(task.ID, result.ID)
}
```

## Controller Tests

Controller tests ensure that the HTTP handlers in your application correctly process requests and return appropriate responses.

### File Structure
- `Delivery/controllers/task_controller_test.go`
- `Delivery/controllers/user_controller_test.go`

### Testing Methodology
1. **Request Handling:** Tests verify how requests are processed by the controllers.
2. **Response Validation:** Tests ensure that the correct status codes and response bodies are returned by the controllers.

### Example Tests

#### **Test Register User Success**

This test verifies that the user registration endpoint processes requests correctly and returns the expected response.

```go
func (suite *UserControllerTestSuite) TestRegisterUser_Success() {
    user := domain.User{
        Username: "new_user",
        Password: "password123",
    }

    // Marshal the user to JSON
    payload, _ := json.Marshal(user)

    // Mock the use case to expect the user registration and return no error
    suite.mockUsecase.On("RegisterUser", mock.MatchedBy(func(u domain.User) bool {
        return u.Username == user.Username && u.Password == user.Password
    })).Return(nil)

    // Create a new POST request with the user data
    req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(payload))
    suite.NoError(err)
    req.Header.Set("Content-Type", "application/json")

    // Create a response recorder to capture the response
    w := httptest.NewRecorder()

    // Serve the request
    suite.router.ServeHTTP(w, req)

    // Check that the status code is 201 Created
    assert.Equal(suite.T(), http.StatusCreated, w.Code)

    // Define the expected response body
    expectedBody := `{"message":"User registered successfully"}`

    // Check that the response body matches the expected JSON
    assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}
```

## Use Case Tests

Use case tests validate the business logic of the application. These tests ensure that the use case layer correctly processes data and interacts with the repository.

### File Structure
- `Usecases/task_usecases_test.go`
- `Usecases/user_usecases_test.go`

### Testing Methodology
1. **Business Logic:** Verify that the use cases implement the correct business logic.
2. **Interaction with Repositories:** Ensure that the use cases interact with the repository layer as expected, including correct handling of errors and results.

### Example Tests

#### **Test Register User**

This test verifies that a user can be successfully registered.

```go
func (suite *UserUsecaseSuite) TestRegisterUser() {
	user := domain.User{Username: "testuser", Password: "password"}

	// Check if username already exists
	suite.mockRepo.On("UsernameExists", user.Username).Return(false, nil)

	// Hash the password before passing it to the mock
	hashedPassword, err := infrastructures.HashPassword(user.Password)
	suite.Require().Nil(err)

	user.Password = hashedPassword // Update the user with the hashed password

	// Use a wildcard matcher to ignore the specific value of the password
	suite.mockRepo.On("Register", mock.MatchedBy(func(u domain.User) bool {
		return u.Username == user.Username && u.Password != "" // Check username and ensure password is not empty
	})).Return(nil)

	err = suite.usecase.Register(user)

	suite.Assert().Nil(err)
	suite.mockRepo.AssertExpectations(suite.T())
}
```
## Main function to run the test suite
```go
func TestTaskHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskHandlerTestSuite))
}
```


## Error Handling
The API handles errors such as:
- **Invalid Task IDs:** Returns a `400 Bad Request` or `404 Not Found` with a descriptive error message.
- **Non-Existent Task Operations:** When trying to update or delete a task that does not exist, the API will return a `404 Not Found` status with an appropriate message.
- **Invalid User Information:** Returns a `400 Bad Request` if the user data is invalid (e.g., missing required fields).
- **Non-Existent User Operations:** When trying to promote, activate, or deactivate a user that does not exist, the API will return a `404 Not Found` status with an appropriate message.
- **Authentication and Authorization:** When a user tries to access a protected route without proper authentication, the API will return a `401 Unauthorized` status with a descriptive message.

Errors are reported with descriptive messages to guide the user in correcting the issue.

## Code Structure
- **`Delivery/main.go`:** The entry point of the application. It starts the process by calling a function in `router/router.go`.
- **`controllers/task_controller.go`:** Contains the logic for handling HTTP requests and responses related to tasks.
- **`controllers/user_controller.go`:** Handles HTTP requests and responses related to user operations like registration, login, promotion, activation, and deactivation.
- **`models/task.go`:** Defines the Task data model.
- **`models/user.go`:** Defines the User data model.
- **`Delivery/router/router.go`:** Sets up and configures the API routes, linking the endpoints to specific functions in the controller.
- **`Infrastructure/auth_middleWare.go`:** Contains the middleware for handling authentication and authorization.
- **`Infrastructure/jwt_service.go`:** Provides the logic for JWT-based authentication.
- **`Infrastructure/password_service.go`:** Handles password hashing and validation.
- **`Repositories/task_repository.go`:** Handles the data persistence logic for tasks.
- **`Repositories/user_repository.go`:** Handles the data persistence logic for users.
- **`Usecases/task_usecases.go`:** Contains the business logic related to tasks.
- **`Usecases/user_usecases.go`:** Contains the business logic related to user operations.
- **`controllers/task_controller_test.go`:** Contains unit tests for the task controller.
- **`controllers/user_controller_test.go`:** Contains unit tests for the user controller.
- **`Usecases/task_usecases_test.go`:** Contains unit tests for the task use cases.
- **`Usecases/user_usecases_test.go`:** Contains unit tests for the user use cases.
- **`Repositories/task_repository_test.go`:** Contains unit tests for the task repository.
- **`Repositories/user_repository_test.go`:** Contains unit tests for the user repository.
- **`docs/api_documentation.md`:** Contains this documentation.
- **`Delivery/.env`:** Contains the MongoDB URI for database connection.

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
```sh
curl -X DELETE http://localhost:8080/tasks/{id}
```