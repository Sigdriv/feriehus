package handler

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
)

// resolveFileReferences recursively iterates through all fields in the config
// and replaces string values starting with "file::" with the content of the file
func (srv *Handler) resolveFileReferences(config interface{}) {
	srv.resolveFileReferencesRecursive(reflect.ValueOf(config))
}

func (srv *Handler) resolveFileReferencesRecursive(v reflect.Value) {
	// Handle pointers
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}

	// Only process structs
	if v.Kind() != reflect.Struct {
		return
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Skip unexported fields
		if !field.CanSet() {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			// Process string fields with "file::" prefix
			value := field.String()
			if len(value) > 6 && value[:6] == "file::" {
				filePath := value[6:]
				correctPath := fmt.Sprintf("./cfg/%s", strings.Replace(filePath, "./", "", 1))

				fileData, err := os.ReadFile(correctPath)
				if err != nil {
					logrus.Fatalf("Failed to read file for field %s: %v", fieldType.Name, err)
				}

				// Trim whitespace/newlines from file content
				field.SetString(strings.TrimSpace(string(fileData)))
				logrus.Infof("Loaded %s from file: %s", fieldType.Name, correctPath)
			}

		case reflect.Struct:
			// Recursively process nested structs
			srv.resolveFileReferencesRecursive(field)

		case reflect.Ptr:
			// Recursively process pointer fields
			if !field.IsNil() {
				srv.resolveFileReferencesRecursive(field)
			}
		}
	}
}
