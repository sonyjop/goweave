package util

import (
	"reflect"
	"testing"
	//"github.com/sonyjop/zegrav/core"
)

func TestParseNodeURI_ftpWithUserInPath(t *testing.T) {
	str := "ftp://foo@myserver?password=secret&recursive=true"
	result, _ := ParseNodeURI(str)
	m := map[string]string{
		"password":  "secret",
		"recursive": "true",
	}
	if result.ComponentName != "ftp" || result.Path != "foo@myserver" ||
		!reflect.DeepEqual(result.Query, m) {
		t.Fatalf("Parsing of URI " + str + "  failed")
	}
}

func TestParseNodeURI_KafkaUri(t *testing.T) {
	str := "kafka://myserver?user=abc&password=secret"
	result, _ := ParseNodeURI(str)
	m := map[string]string{
		"user":     "abc",
		"password": "secret",
	}
	if result.ComponentName != "kafka" || result.Path != "myserver" ||
		!reflect.DeepEqual(result.Query, m) {
		t.Fatalf("Parsing of URI " + str + "  failed")
	}
}
func TestParseNodeURI_ComponentMissing(t *testing.T) {
	str := "://myserver?user=abc&password=secret"
	_, err := ParseNodeURI(str)

	if err == nil {
		t.Fatalf("Component Name Missing")
	}
}

func TestParseNodeURI_PathMissing(t *testing.T) {
	str := "rest://?user=abc&password=secret"
	_, err := ParseNodeURI(str)

	if err == nil {
		t.Fatalf("Path missing")
	}
}

/*
	func TestParseNodeURI_QueryMising(t *testing.T) {
		str := "http://myserver?"
		_, err := ParseNodeURI(str)

		if err == nil {
			t.Fatalf("Query missing but a separator is available in the uri")
		}
	}
*/
func TestParseNodeURI_QueryKeyorValueMissing(t *testing.T) {
	str := "kafka://myserver?userabc&password=secret"
	_, err := ParseNodeURI(str)

	if err == nil {
		t.Fatalf("Query parameter key/value both must be available")
	}
}
func TestParseNodeURI_ComponentStartingWithNumber(t *testing.T) {
	str := "3kafka://myserver?userabc&password=secret"
	_, err := ParseNodeURI(str)

	if err == nil {
		t.Fatalf("Component name starting with number")
	}
}
func TestParseNodeURI_ComponentHasSplChar(t *testing.T) {
	str := "kaf?ka://myserver?userabc&password=secret"
	_, err := ParseNodeURI(str)

	if err == nil {
		t.Fatalf("Component has a special character")
	}
}
func TestParseNodeURI_PathStartsWithNumber(t *testing.T) {
	str := "kafka://3myserver?userabc&password=secret"
	_, err := ParseNodeURI(str)

	if err == nil {
		t.Fatalf("path name starting with number")
	}
}
func TestParseNodeURI_PathHasSplChar(t *testing.T) {
	str := "kafka://myser?ver?userabc&password=secret"
	_, err := ParseNodeURI(str)

	if err == nil {
		t.Fatalf("path has a special character")
	}
}
func TestParseNodeURI_KeyStartsWithNumber(t *testing.T) {
	str := "kafka://3myserver?userabc&3password=secret"
	_, err := ParseNodeURI(str)

	if err == nil {
		t.Fatalf("key starting with number")
	}
}
func TestParseNodeURI_KeyHasSplChar(t *testing.T) {
	str := "kafka://myserver?use?rabc&password=secret"
	_, err := ParseNodeURI(str)

	if err == nil {
		t.Fatalf("key has a special character")
	}
}
