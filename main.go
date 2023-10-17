package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var db *sql.DB

func getDBConn() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")

	/* root:password@tcp(localhost:3306)/testdb */
	db_url := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s", db_username, db_password, db_host, db_port, db_name)

	db, err := sql.Open("mysql", db_url)
	errCheck(err)
	return db
}

func handleIndexRename(tableName string, batchName string, indexName string, columns []string) {
	log.Printf(">>>>>>>>>>>>>>>Index Modification for table: %s<<<<<<<<<<<<<<<", tableName)
	dropIndexQuery := fmt.Sprintf("ALTER TABLE %s DROP INDEX %s", tableName, indexName)
	log.Printf("Drop Index Query: %s", dropIndexQuery)

	_, dropIndexErr := db.Exec(dropIndexQuery)
	errCheck(dropIndexErr)

	newIndexName := fmt.Sprintf("%s%s", indexName, batchName)
	columnsStr := strings.Join(columns, ",")
	addIndexQuery := fmt.Sprintf("ALTER TABLE %s ADD INDEX %s (%s)", tableName, newIndexName, columnsStr)
	log.Printf("Add Index Query: %s", addIndexQuery)

	_, addIdxErr := db.Exec(addIndexQuery)
	errCheck(addIdxErr)

	log.Printf(">>>>>>>>>>>>>>>Index Modification completed for table: %s<<<<<<<<<<<<<<<", tableName)
	log.Println("")

}

func renameIndex(tableName string, batchName string) {
	showIndexQuery := fmt.Sprintf("show index from %s", tableName)
	indexResults, err := db.Query(showIndexQuery)
	errCheck(err)

	tableIndexes := make(map[string][]string)

	for indexResults.Next() {
		var (
			tableName    string
			nonUnique    int
			keyName      string
			seqInIndex   int
			columnName   string
			collation    string
			cardinality  int
			subPart      *int
			packed       *string
			isNull       string
			indexType    string
			comment      string
			indexComment string
			visible      string
			expression   *string
		)

		err := indexResults.Scan(&tableName, &nonUnique, &keyName, &seqInIndex, &columnName, &collation,
			&cardinality, &subPart, &packed, &isNull, &indexType, &comment,
			&indexComment, &visible, &expression,
		)
		errCheck(err)

		tableIndexes[keyName] = append(tableIndexes[keyName], columnName)

	}

	for key, columns := range tableIndexes {
		if key == "PRIMARY"{
			continue
		}
		handleIndexRename(tableName, batchName, key, columns)
	}

}

func makeExport(tableName string, dateSuffix string) {
	/* /var/lib/mysql-files/ */
	db_export_dir := os.Getenv("DB_EXPORT_DIR")
	export_dir := fmt.Sprintf("%s%s-%s.csv", db_export_dir, tableName, dateSuffix)
	log.Printf("Export Dir: %s", export_dir)

	export_query_str := fmt.Sprintf(
		"SELECT * INTO OUTFILE '%s' FIELDS TERMINATED BY ',' LINES TERMINATED BY '\\n' FROM %s", export_dir, tableName)
	log.Printf("Query: %s", export_query_str)
	_, err := db.Exec(export_query_str)
	errCheck(err)

	oldTableName := fmt.Sprintf("%s%s", tableName, dateSuffix)

	rename_table_query_str := fmt.Sprintf("ALTER TABLE %s RENAME %s", tableName, oldTableName)
	log.Printf("Rename query: %s", rename_table_query_str)
	_, rename_err := db.Exec(rename_table_query_str)
	errCheck(rename_err)

	recreate_table_query_str := fmt.Sprintf("CREATE TABLE %s LIKE %s", tableName, oldTableName)
	log.Printf("Recreate query: %s", recreate_table_query_str)
	_, recreate_err := db.Exec(recreate_table_query_str)
	errCheck(recreate_err)

	renameIndex(tableName, dateSuffix)
}

func manageTable() {
	db_export_tables := os.Getenv("DB_EXPORT_TABLES")
	tables := strings.Split(db_export_tables, ",")

	log.Printf(">>>>>>>Archiving Data<<<<<<<")

	now := time.Now()
	dateSuffix := now.Format("20060102")

	for _, table := range tables {
		log.Printf("Table: %s", table)

		makeExport(table, dateSuffix)
	}

	log.Printf(">>>>>>>Archiving Data Completed<<<<<<<")

}

func main() {
	db = getDBConn()
	defer db.Close()
	manageTable()
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
