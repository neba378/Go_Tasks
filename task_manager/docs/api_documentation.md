**Task Management REST API**

**Introduction**
The Task Management REST API is designed to manage tasks within a system. It supports operations such as creating, updating, retrieving, and deleting tasks.

**System Requirements**
- Go 1.22 or higher
- A terminal or command-line interface

**Running the Application**
1. Navigate to the `main.go` file in your terminal or command prompt.
2. Run the application using the following command:
    ```sh
    go run main.go
    ```

   This will start the server and you can interact with the API via HTTP requests.

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

**Error Handling**
The API handles errors such as:
- Invalid task IDs
- Attempting to update or delete a non-existent task

Errors are reported with descriptive messages to guide the user in correcting the issue.

**Code Structure**
- `main.go`: The entry point of the application. It stats the process by calling a function in the `router/router.go`.
- `controllers/task_controller.go`: It provides functions to the router.
- `models/task.go`: Defines the Task data model.
- `data/task_service.go`: provides functions to the controller.
- `router/router.go`: Sets up and configures the API routes => relates the endpoints with the specific functions from controller.
- `docs/api_documentation.md`: Contains this documentation.

