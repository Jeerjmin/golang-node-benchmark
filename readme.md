# Node.js (Express) vs Golang (Gin) Performance Comparison

This project is designed to compare the performance of a Node.js (Express) application with a Golang (Gin) application.

## Test Description

The `test-task.lua` file performs the following tasks:
- `GET /users`
- `POST /users`
- `DELETE /users`

## How to Run the Test

1. **Start the web applications and database:**

   ```bash
   docker-compose up
    ```

2. **Run the benchmark against the Node.js application:**
    ```bash
    wrk -t2 -c50 -d30s -s test-task.lua http://localhost:3000
    ```

3. **Run the benchmark against the Golang application:**
    ```bash
    wrk -t2 -c50 -d30s -s test-task.lua http://localhost:8080
    ```
