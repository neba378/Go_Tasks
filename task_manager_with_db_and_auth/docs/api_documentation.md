# Task Management REST API

**Introduction**
The Task Management REST API is designed to manage tasks within a system. It supports operations such as creating, updating, retrieving, and deleting tasks. It is connected to the MongoDB database online. 
It also supports user registration and login. It uses JWT to authenticate the user. There are two kind of users in the process: User(Normal) and Admin. The admin has some rights that the normal users do not have like accessing `/protected` path.
There is a middleware that uses JWT to authenticate the user. It gives access based on the user role [admin,user].

**System Requirements**
- Go 1.22 or higher
- A terminal or command-line interface
- Postman (optional) for testing
- MongoDB (local or cloud)
- use `go get` and the package path to install any package.

**Running the Application**
1. Install MongoDB locally or get a URI from [MongoDB Atlas](https://cloud.mongodb.com/). Add this URI or `mongodb://localhost:27017` to the `.env` file located in the same directory as `main.go`.

2. Set your `JWT_SECRET` in the `.env` file.
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
**Running the Application**
- Before starting anything make sure to install mongodb on your pc or get a uri from the official website `https://cloud.mongodb.com/`. Then add that uri or `mongodb://localhost:27017` to the `.env` file in the same path as `main.go`.
Additionally for this version you need to have your `JWT_SECRETE` in `.env` file.
- Then follow these steps:
1. Navigate to the `main.go` file in your terminal or command prompt.
2. Run the application using the following command:
    ```sh
    go run main.go
    ```

   This will start the server and you can interact with the API via HTTP requests.

   optionally you can use `Air` read [here](https://github.com/air-verse/air) for more.

**Endpoints**

1. **Get Tasks**
   - **Description:** Retrieves all tasks.
   - **Method:** GET
   - **Endpoint:** `/tasks`
   - **Response:** List of tasks with their details.

2. **Create Task**
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
   - **Response:** Confirmation of task creation.

3. **Update Task**
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
   - **Response:** Confirmation of task update.

4. **Delete Task**
   - **Description:** Deletes a task by its ID.
   - **Method:** DELETE
   - **Endpoint:** `/tasks/{id}`
   - **Response:** Confirmation of task deletion.

5. **Resister**
   - **Description:** Registers users based on role they have given in the param.
   - **Method:** POST
   - **Endpoint:** `/register/{role}`
   - **Response:** Confirmation of user registration.
   - **Input**: JSON object with user details.
```sh
   {
   "username": "newuser",
   "password": "password"
   }
   ```

5. **Login**
   - **Description:** LogIn users based on role the username and password they provided.
   - **Method:** POST
   - **Endpoint:** `/login/`
   - **Response:** Confirmation of user login with the token created for an hour.
   - example response:
   ```sh
         {
      "token": "your_jwt_token"
      }
   ```

5. **Secrete page access**
   - **Description:** gives access to the user for a secrete page based on user role.
   - **Method:** GET
   - **Endpoint:** `/protected/secrete`
   - **Response:** message to inform they have accessed correctly.

**Error Handling**
The API handles errors such as:
- Invalid task IDs
- Attempting to update or delete a non-existent task
- It automatically gives new and unique IDs for new tasks
- JWT Token validation
- Login only allowed for existing users

Errors are reported with descriptive messages to guide the user in correcting the issue.

**Code Structure**
- `main.go`: The entry point of the application. It stats the process by calling a function in the `router/router.go`.
- `controllers/task_controller.go`: It provides functions to the router.
- `models/task.go`: Defines the Task data model.
- `models/user.go`: Defines the User data model.
- `data/task_service.go`: provides Task related functions to the controller.
- `data/user_service.go`: provides user related functions to the controller. Like the login(checkUser) and registration.
- `middleware/auth_middleware.go`: Provides validity for accessing some protected features (endpoints).
- `router/router.go`: Sets up and configures the API routes => relates the endpoints with the specific functions from controller.
- `docs/api_documentation.md`: Contains this documentation.
- `.env` : contains the db uri and JWT_SECRET key
