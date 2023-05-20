# aqua-farm

This project is a Go backend application with a PostgreSQL database that can be run using Docker containers.

## Prerequisites

Before running this project, you will need the following:

- Go installed on your system
- Docker installed on your system

## Running the Project

To run the project, use the following commands in your terminal:

1. Run the PostgreSQL database container:
    ```bash
    make db
    ```

2. Initialize the database with the required schema and data:


    ```bash
    make init-db
    ```
3. Run the Go backend application:

    ```bash
    make backend
    ```

## APIS

1. GET all farm & POND
    ``` bash
        GET localhost:9090/farm
    ```
    sample response: 
    ``` json
        {
            "count": 4,
            "unique_user_agent": 1,
            "Data": [
                {
                    "id": 1,
                    "name": "farm1",
                    "Pond": [
                        {
                            "id": 1,
                            "farmid": 1,
                            "name": "pond1"
                        },
                        {
                            "id": 2,
                            "farmid": 1,
                            "name": "pond2"
                        }
                    ]
                },
                {
                    "id": 2,
                    "name": "farm2",
                    "Pond": [
                        {
                            "id": 3,
                            "farmid": 2,
                            "name": "pond3"
                        },
                        {
                            "id": 4,
                            "farmid": 2,
                            "name": "pond4"
                        }
                    ]
                },
                {
                    "id": 3,
                    "name": "farm4",
                    "Pond": []
                }
            ]
        }
    ```


2. POST add Farm 
    ``` bash
        POST localhost:9090/farm
    ```
    Request:
    ``` json
        {
            "name":"farm4" 
        }
    ```

3. Update Farm 
    ``` bash
        PUT localhost:9090/farm
    ```
    Request body:
    ``` json
        {
            "id":3,
            "name":"farm5" 
        }
    ```

4. Delete Farm 
    ``` bash
        DELETE localhost:9090/farm
    ```
    Request body:
    ``` json
        {
            "id":3,
        }
    ```

5. Get All Pond 
    ``` bash
        GET localhost:9090/pond
    ```
    sample response:
    ``` json
        {
            "count": 1,
            "unique_user_agent": 1,
            "Data": [
                {
                    "id": 1,
                    "farmid": 1,
                    "name": "pond1"
                },
                {
                    "id": 2,
                    "farmid": 1,
                    "name": "pond2"
                },
                {
                    "id": 3,
                    "farmid": 2,
                    "name": "pond3"
                },
                {
                    "id": 4,
                    "farmid": 2,
                    "name": "pond4"
                }
            ]
        }
    ```

4. Add Pond 
    ``` bash
        POST localhost:9090/pond
    ```
    Request body:
    ``` json
        {
            "farm_id":3,
            "name":"pond4"
        }
    ```

5. Update Pond 
    ``` bash
        PUT localhost:9090/pond
    ```
    Request body:
    ``` json
        {
            "id":2,
            "farm_id":3,
            "name":"pond_update"
        }
    ```

5. Delete Pond 
    ``` bash
        DELETE localhost:9090/pond
    ```
    Request body:
    ``` json
        {
            "id":2,
        }
    ```

## Makefile Targets
    `backend`: Runs the Go backend application.
    `db`: Starts a PostgreSQL database container using Docker.
    `init-db`: Initializes the database with the required schema .
    `fill-db`: Initializes the database with sample data.

