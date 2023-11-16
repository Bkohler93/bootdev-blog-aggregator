# bootdev-blog-aggregator

Blog aggregator written in Go!

## Database Migrations

This project uses [goose](https://github.com/pressly/goose) to handle database migrations!

To run migrations use a command like this, which tells goose which database you are using, the database connection string, and whether you want to do an up or down migration:

```bash
goose postgres postgres://brettkohler:@localhost:5432/bloggo up
```

## SQL to Go: sqlc

This project uses [sqlc](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html) to handle generating Go code directly from SQL queries.
