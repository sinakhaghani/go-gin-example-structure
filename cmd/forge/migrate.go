package forge

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func MakeMigration(name string) {
	dir := "database/migrations"
	versionFile := filepath.Join(dir, ".migration_version")

	// ساخت مسیر اگر وجود ندارد
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	// خواندن نسخه فعلی
	version := 1
	if data, err := os.ReadFile(versionFile); err == nil {
		v, err := strconv.Atoi(strings.TrimSpace(string(data)))
		if err == nil {
			version = v + 1
		}
	}

	// ذخیره نسخه جدید
	os.WriteFile(versionFile, []byte(strconv.Itoa(version)), 0644)

	// ساخت نام فایل
	paddedVersion := fmt.Sprintf("%06d", version)
	snakeName := toSnakeCase(name)

	upFile := fmt.Sprintf("%s/%s_%s.up.sql", dir, paddedVersion, snakeName)
	downFile := fmt.Sprintf("%s/%s_%s.down.sql", dir, paddedVersion, snakeName)

	createEmptyFile(upFile)
	createEmptyFile(downFile)

	fmt.Println("Created migration:")
	fmt.Println("  ↑", upFile)
	fmt.Println("  ↓", downFile)
}

func createEmptyFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("Failed to create migration file:", err)
		return
	}
	defer f.Close()
}

func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_', r+32)
		} else {
			if r >= 'A' && r <= 'Z' {
				result = append(result, r+32)
			} else {
				result = append(result, r)
			}
		}
	}
	return string(result)
}
