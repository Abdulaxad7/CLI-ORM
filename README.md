[![Build Status](http://209.38.236.180/buildStatus/icon?job=pi-challenge)](http://209.38.236.180/job/pi-challenge/)

# CLI-Orm

A CLI-based ORM tool for interacting with MySQL databases using Go. This project allows developers to perform essential database operations such as creating databases, managing tables, inserting values, and executing updates—all through a command-line interface.

---

## Features

- Create and drop MySQL databases.
- Show existing databases and tables through queries.
- Create and manage tables (columns and data types).
- Insert, update, and delete table records.
- Query database content and schemas.
- Lightweight and intuitive CLI-based workflow.
- Built using Go's GORM library for MySQL connection.

---

## Project Structure

The project is organized as follows:

---

## Installation

### Preconditions
1. Go **1.22+** is required for building and running the project.
2. A MySQL server instance must be available and configured.

### Clone the Repository

```bash
git clone https://github.com/your-username/Cli-Orm.git
cd Cli-Orm
```

### Install Dependencies

Run the following command to install necessary dependencies:

```bash
go mod download
```

---

## Usage

### Configuration
Update the configuration with your MySQL credentials (if required). This likely involves modifying an environment file or `login.go`.

### Run the CLI
To start interacting with the CLI-ORM project, run the following command:

```bash
go run src/msql.go
```

### Available Commands

The project provides the following functionalities:

1. **Databases**
   - `CreateDb(dbName)` – Create a new database.
   - `DropDb(dbName)` – Drop (delete) an existing database.
   - `ShowDb()` – List all databases.

2. **Tables**
   - `CreateTable(dbName, tableName, columns)` – Create a new table in the database.
   - `DropTable(dbName, tableName)` – Drop (delete) a table from the database.
   - `ShowTables(dbName)` – Show all tables in the selected database.

3. **Data Manipulation**
   - `InsertToTable(tableName, values)` – Insert rows into a table.
   - `Update(dbName, tableName, columnName, valueBefore, valueAfter)` – Update specific rows in a table.
   - `Delete(tableName, column, value)` – Delete rows matching a condition.
   - `ShowValues(tableName)` – Display all data from a table.

---

## Key Files

### Main Modules
- **`src/msql.go`** – Entry-point for MySQL operations in the CLI interface.
- **`src/login.go`** – Authentication and login management for MySQL.
- **`config/mq/msql-queries.go`** – Contains SQL query logic implemented using GORM.

### Database Operations
Located in `src/msql/dbs/` and includes:
- **`CreateDb.go`** – Logic for creating a database.
- **`DropDb.go`** – Logic for deleting a database.
- **`ShowDbMsql.go`** – Displaying all available databases in the connected server.
- **`ShowVlMsql.go`** – Query operations to fetch table data.
- **`insertTb.go`** – Insertion of records into the database tables.
- **`UpdateDb.go`** – Update operations for rows in MySQL tables.

---

## Requirements
- **Go 1.22+**
- **MySQL Server**: Ensure your server is running and accessible.
- Dependencies are managed using `go mod`.

---


## License

This project is licensed under MIT License.
