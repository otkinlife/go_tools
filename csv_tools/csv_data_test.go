package csv_tools

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestNewCSVData(t *testing.T) {
	csv := NewCSVData()

	if csv == nil {
		t.Fatal("NewCSVData should return a non-nil pointer")
	}

	if csv.split != "," {
		t.Errorf("Default split should be ',', got %s", csv.split)
	}

	if len(csv.headers) != 0 {
		t.Errorf("Headers map should be empty, got %d items", len(csv.headers))
	}

	if len(csv.headersLine) != 0 {
		t.Errorf("HeadersLine should be empty, got %d items", len(csv.headersLine))
	}

	if len(csv.data) != 0 {
		t.Errorf("Data should be empty, got %d items", len(csv.data))
	}

	if csv.err != nil {
		t.Errorf("Error should be nil, got %v", csv.err)
	}
}

func TestSetSplit(t *testing.T) {
	csv := NewCSVData()

	// Test with valid split
	result := csv.SetSplit(";")
	if csv.split != ";" {
		t.Errorf("Split should be ';', got %s", csv.split)
	}
	if result != csv {
		t.Error("SetSplit should return the CSVData instance")
	}

	// Test with empty split (should default to comma)
	csv.SetSplit("")
	if csv.split != "," {
		t.Errorf("Split should default to ',', got %s", csv.split)
	}
}

func TestSetAndGetHeaders(t *testing.T) {
	csv := NewCSVData()
	headers := []string{"Name", "Age", "Email"}

	csv.SetHeaders(headers)

	// Check headersLine
	if !reflect.DeepEqual(csv.headersLine, headers) {
		t.Errorf("HeadersLine should be %v, got %v", headers, csv.headersLine)
	}

	// Check headers map
	expectedMap := map[string]int{"Name": 0, "Age": 1, "Email": 2}
	if !reflect.DeepEqual(csv.headers, expectedMap) {
		t.Errorf("Headers map should be %v, got %v", expectedMap, csv.headers)
	}

	// Test GetHeaders
	returnedHeaders := csv.GetHeaders()
	if !reflect.DeepEqual(returnedHeaders, expectedMap) {
		t.Errorf("GetHeaders should return %v, got %v", expectedMap, returnedHeaders)
	}
}

func TestLoadFromLines(t *testing.T) {
	csv := NewCSVData()

	// Test with valid data
	lines := [][]string{
		{"Name", "Age", "Email"},
		{"John", "30", "john@example.com"},
		{"Jane", "25", "jane@example.com"},
	}

	result := csv.LoadFromLines(lines)

	if result != csv {
		t.Error("LoadFromLines should return the CSVData instance")
	}

	if !reflect.DeepEqual(csv.headersLine, lines[0]) {
		t.Errorf("Headers should be %v, got %v", lines[0], csv.headersLine)
	}

	if !reflect.DeepEqual(csv.data, lines[1:]) {
		t.Errorf("Data should be %v, got %v", lines[1:], csv.data)
	}

	// Test with empty data
	csv = NewCSVData()
	emptyLines := [][]string{}
	csv.LoadFromLines(emptyLines)

	if csv.err != io.EOF {
		t.Errorf("Error should be io.EOF, got %v", csv.err)
	}
}

func TestLoadFromReader(t *testing.T) {
	// Test with valid data
	csvContent := "Name,Age,Email\nJohn,30,john@example.com\nJane,25,jane@example.com"
	reader := strings.NewReader(csvContent)

	csv := NewCSVData()
	result := csv.LoadFromReader(reader)

	if result != csv {
		t.Error("LoadFromReader should return the CSVData instance")
	}

	expectedHeaders := []string{"Name", "Age", "Email"}
	if !reflect.DeepEqual(csv.headersLine, expectedHeaders) {
		t.Errorf("Headers should be %v, got %v", expectedHeaders, csv.headersLine)
	}

	expectedData := [][]string{
		{"John", "30", "john@example.com"},
		{"Jane", "25", "jane@example.com"},
	}
	if !reflect.DeepEqual(csv.data, expectedData) {
		t.Errorf("Data should be %v, got %v", expectedData, csv.data)
	}

	// Test with custom delimiter
	csvContent = "Name;Age;Email\nJohn;30;john@example.com\nJane;25;jane@example.com"
	reader = strings.NewReader(csvContent)

	csv = NewCSVData().SetSplit(";")
	csv.LoadFromReader(reader)

	if !reflect.DeepEqual(csv.headersLine, expectedHeaders) {
		t.Errorf("Headers should be %v, got %v", expectedHeaders, csv.headersLine)
	}

	if !reflect.DeepEqual(csv.data, expectedData) {
		t.Errorf("Data should be %v, got %v", expectedData, csv.data)
	}

	// Test with invalid reader
	invalidReader := &errorReader{}
	csv = NewCSVData()
	csv.LoadFromReader(invalidReader)

	if csv.err == nil {
		t.Error("Error should not be nil when reading from invalid reader")
	}

	// Test with empty data
	emptyReader := strings.NewReader("")
	csv = NewCSVData()
	csv.LoadFromReader(emptyReader)

	if len(csv.data) != 0 {
		t.Errorf("Data should be empty, got %d items", len(csv.data))
	}
}

func TestGetHeaderIndex(t *testing.T) {
	csv := NewCSVData()
	headers := []string{"Name", "Age", "Email"}
	csv.SetHeaders(headers)

	// Test existing header
	index := csv.GetHeaderIndex("Age")
	if index != 1 {
		t.Errorf("Index for 'Age' should be 1, got %d", index)
	}

	// Test non-existing header
	index = csv.GetHeaderIndex("Address")
	if index != -1 {
		t.Errorf("Index for non-existing header should be -1, got %d", index)
	}
}

func TestGetLineValue(t *testing.T) {
	csv := NewCSVData()
	lines := [][]string{
		{"Name", "Age", "Email"},
		{"John", "30", "john@example.com"},
		{"Jane", "25", "jane@example.com"},
	}
	csv.LoadFromLines(lines)

	// Test valid key and line
	value := csv.GetLineValue("Email", 0)
	if value != "john@example.com" {
		t.Errorf("Value should be 'john@example.com', got '%s'", value)
	}

	// Test non-existing key
	value = csv.GetLineValue("Address", 0)
	if value != "" {
		t.Errorf("Value for non-existing key should be empty, got '%s'", value)
	}

	// Test out of bounds line index
	value = csv.GetLineValue("Name", 10)
	if value != "" {
		t.Errorf("Value for out of bounds line should be empty, got '%s'", value)
	}

	// Test negative line index
	value = csv.GetLineValue("Name", -1)
	if value != "" {
		t.Errorf("Value for negative line index should be empty, got '%s'", value)
	}
}

func TestGetData(t *testing.T) {
	csv := NewCSVData()
	expectedData := [][]string{
		{"John", "30", "john@example.com"},
		{"Jane", "25", "jane@example.com"},
	}
	csv.SetData(expectedData)

	data := csv.GetData()
	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("GetData should return %v, got %v", expectedData, data)
	}
}

func TestSaveToFile(t *testing.T) {
	// Create temporary file path
	tempFile := os.TempDir() + "/test_csv_data.csv"
	defer os.Remove(tempFile)

	// Create CSV data
	csv := NewCSVData()
	headers := []string{"Name", "Age", "Email"}
	data := [][]string{
		{"John", "30", "john@example.com"},
		{"Jane", "25", "jane@example.com"},
	}
	csv.SetHeaders(headers)
	csv.SetData(data)

	// Save to file
	csv.SaveToFile(tempFile)

	// Check if file exists
	if _, err := os.Stat(tempFile); os.IsNotExist(err) {
		t.Errorf("File %s should exist", tempFile)
	}

	// Read file content
	content, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	expectedContent := "Name,Age,Email\nJohn,30,john@example.com\nJane,25,jane@example.com\n"
	if string(content) != expectedContent {
		t.Errorf("File content should be '%s', got '%s'", expectedContent, string(content))
	}

	// Test with error
	csv.err = io.EOF
	csv.SaveToFile(tempFile + "/invalid/path")
	// Should not panic or create file
}

func TestGetError(t *testing.T) {
	csv := NewCSVData()
	expectedErr := io.EOF
	csv.err = expectedErr

	err := csv.GetError()
	if err != expectedErr {
		t.Errorf("GetError should return %v, got %v", expectedErr, err)
	}
}

// Helper error reader for testing
type errorReader struct{}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, io.ErrUnexpectedEOF
}
