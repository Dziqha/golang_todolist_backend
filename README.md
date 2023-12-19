# Golang ToDo List Backend

This is a simple ToDo List backend built with Golang and the Echo Framework.

## Prerequisites

Ensure you have the following installed on your machine before proceeding:

- [Golang](https://golang.org/)
- [MySQL](https://www.mysql.com/)
- Text Editor (e.g., Visual Studio Code, Sublime Text)

## Getting Started

1. **Database Setup:**

   - Create a new MySQL database.
   - Import the `todolist.sql` file into your newly created database.

2. **Environment Setup:**

   - Duplicate the provided `.env.example` file and rename it to `.env`.
   - Update the `.env` file with your database credentials and any other configuration parameters.

3. **Install Dependencies:**

   Open your terminal, navigate to the project directory, and run the following command to install the project dependencies:

   ```bash
   go mod download
   ```

4. Run the Application

    Start the application by running the following command:

    ```bash
    go run api.go
    ```

## Documentation

The ToDo List backend provides endpoints for managing tasks. You can use API calls to interact with the application and manage your ToDo items.

Here are some example API endpoints:

### Auth

`POST /auth/register` for create a new account

`POST /auth/login` for login to the account

`PUT /auth/update` for update the account

### ToDo

`GET /todo/view` to retrieve all tasks

`POST /todo/create` for create a new task

`POST /todo/update` for update the task

`DELETE /todo/delete` to delete a task

Please refer to the API documentation or explore the source code for more details on available endpoints and their usage.

## Contributing

Feel free to contribute to the project by submitting issues or pull requests. Your feedback and contributions are highly appreciated!