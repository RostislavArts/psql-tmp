# psql-temp
This is a very simple program for running .sql files in temporary PostgreSQL database so you don't have to setup any users, passwords, etc.

It was made for learning SQL and testing queries in safe and easy-to-use environment

## Building
To build this project you must first install the following:
- [PostgreSQL](https://github.com/postgres/postgres) itself
- [Go Compiler](https://go.dev/)
- [pq_tmp](https://github.com/eradman/ephemeralpg/) utility

To build use:
- **make build** to just get binary compiled by Go in your current directory
- **make install** to compile the binary and move it to /usr/bin (so you can use it by just typing psql-tmp from any directory) *(requires sudo privileges)*

## Usage
To use it just enter 'psql-tmp /path/to/file.sql' and it will execute your query
