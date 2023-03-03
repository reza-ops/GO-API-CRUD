package main

import (
	book_migration "example/server/db/migration/book"
)

func main() {
	book_migration.MigrateBooks()
}
