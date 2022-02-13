package database

var database string

// Automatically initialized when package is imported
func init() {
	database = "MySQL"
}

func GetDatabase() string {
	return database
}
