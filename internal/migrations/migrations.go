package migrations

import (
	"database/sql"
	"log"
	"os"
	"strings"
)

func Up(db *sql.DB, upMigrationsPath string) {
	files, e := os.ReadDir(upMigrationsPath)
	if e != nil {
		log.Fatal("package migrations.Up: cannot read upMigrationsPath ", e)
	}

	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, ".sql") {
			migration, e := os.ReadFile(fileName)
			if e != nil {
				log.Fatalf("package migrations.Up: cannot read migration file with name %s", fileName)
			}

			_, e = db.Exec(string(migration))
			if e != nil {
				log.Fatal("package migrations.Up: cannot exec migration statement ", e)
			}
		}
	}
}

// For feature usage
func Down(db *sql.DB, downMigrationsPath string) error {
	return nil
}
