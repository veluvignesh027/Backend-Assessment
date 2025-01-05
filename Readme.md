**Project Name:**
    Backend Assessment
**Overview:**
    This project uses Docker and Go to manage and run a web application which loads csv data into postgres DB


**Makefile Commands:**

The following commands are available in the Makefile:

docker-run
This command builds and starts the Docker containers defined in the docker-compose.yml file. It attempts to use Docker Compose V2 first, and if that fails, it falls back to Docker Compose V1.

Usage:_ make docker-run_

docker-down
This command stops and removes the Docker containers defined in the docker-compose.yml file. Similar to docker-run, it tries to use Docker Compose V2 first and falls back to V1 if necessary.

Usage: _make docker-down_

run
This command runs the Go application. It executes the go run command on all Go files in the current directory and its subdirectories.

Usage: _make run_

clean
This command cleans up the project by removing the compiled binary file named main. It is useful for ensuring that you start with a clean state.

Usage: _make clean_

**Prerequisites:**

Docker: Ensure that Docker is installed and running on your machine.
Docker Compose: Make sure Docker Compose is installed. This Makefile is compatible with both Docker Compose V1 and V2.
Go: Ensure that Go is installed on your machine if you plan to run the Go application directly.
Make: Ensure make is installed


**Getting Started:**

    1. create a .env file in the home directory and set the ENV variables
        PORT=9010
        APP_ENV=local
        BLUEPRINT_DB_HOST=<db ip>
        BLUEPRINT_DB_PORT=5432
        BLUEPRINT_DB_DATABASE=<db name>
        BLUEPRINT_DB_USERNAME=<db user>
        BLUEPRINT_DB_PASSWORD=password1234
        BLUEPRINT_DB_SCHEMA=public

    2. Install all the Prerequisites
    3. make docker-run
    4. make run


**REST API ENDPOINT:**

    1. "/resync"
        Method: HTTP GET
        Objective: Loads the .csv files into database

    2. "/customer/analysis/:startdate/:enddate/:query"
        Method: HTTP GET
        Objective:
            startdate: starting date for the query
            enddate: ending date for the query
            query: Query options 
                1. totalcustomers
                2. totalorders
                3. averageorders
