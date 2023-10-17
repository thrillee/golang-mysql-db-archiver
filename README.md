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
```


The script will perform the following steps for each specified table:

- Export table data to a CSV file in the specified export directory.
- Rename the original table by appending the current date in YYYYMMDD format.
- Create a new table with the same structure as the original table.


## Example

```env
DB_USERNAME=your_mysql_user
DB_PASSWORD=your_mysql_password
DB_NAME=your_database_name
DB_HOST=localhost
DB_PORT=3306
DB_EXPORT_DIR=/var/lib/mysql-files/
DB_EXPORT_TABLES=table1,table2,table3
```
Make sure to customize the environment variables to match your specific setup and requirements.

## Disclaimer

This script performs operations on your database. Use it with caution, and ensure you have appropriate backups and permissions before running it in a production environment.

## Author

Oluwatobi Bello

For questions or support, contact bellotobiloba01@gmail.com.
