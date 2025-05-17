package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func Migrate() {
	paths := getSchemaFilesPaths()
	fileBytes := readSchemaFiles(paths)

	fmt.Println()

	for index, bytes := range fileBytes {
		sql := getStringFromBytes(bytes)

		_, err := Connection.Exec(context.Background(), sql)

		if err != nil {
			fmt.Println("Failed to execute schema for file", paths[index], ":", err)
			continue
		}

		fmt.Println("Successfully migrated", paths[index])
	}

	fmt.Println()

}

func getSchemaFilesPaths() []string {
	files, err := filepath.Glob("db/schemas/*.sql")

	if err != nil {
		log.Fatal("Failed to get schema files: ", err)
	}

	sort.Strings(files)

	return files
}

func readSchemaFiles(paths []string) [][]byte {
	var result [][]byte

	for _, path := range paths {
		bytes, err := os.ReadFile(path)

		if err != nil {
			fmt.Println("Error reading file at ", path, ": ", err)
			continue
		}

		result = append(result, bytes)
	}

	return result
}

func getStringFromBytes(bytes []byte) string {
	return string(bytes)
}
