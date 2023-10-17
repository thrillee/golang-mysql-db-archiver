# Database Table Export and Management

This Go script is designed to automate the process of exporting data from specified database tables, renaming the tables, and creating clones of the tables. It's useful for archiving data and maintaining table structures for historical records.

## Prerequisites

Before running the script, ensure you have the following prerequisites in place:

- Go environment set up.
- MySQL server with the required access permissions.
- Necessary Go packages installed: `github.com/go-sql-driver/mysql` and `github.com/joho/godotenv`.

## Configuration

The script uses environment variables to configure the database connection and export settings. You can set these environment variables in a `.env` file or directly in the script.

- **DB_USERNAME**: MySQL username for database access.
- **DB_PASSWORD**: MySQL password for the specified user.
- **DB_NAME**: The name of the database to connect to.
- **DB_HOST**: The host where the MySQL database server is running.
- **DB_PORT**: The port number to connect to the MySQL server.
- **DB_EXPORT_DIR**: The directory where exported CSV files will be stored.
- **DB_EXPORT_TABLES**: A comma-separated list of tables to be processed by the script.

## Usage

To run the script, execute the following command:

```bash
go run main.go

