package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct { // ✅ fixed struct syntax
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ") // ✅ corrected variable name
	if err != nil {
		return err // ✅ typo fixed: "retun" → "return"
	}
	return os.WriteFile(s.FileName, fileData, 0644) // ✅ typo fixed
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, data) // ✅ typo fixed
}
