package main

import (
	"reflect"
	"testing"
	"time"
)

var testFilePath string = "REDACTED";

func TestDeviceFile(t *testing.T) {
	start := time.Now()
	test1 := toolDetection(testFilePath, 512)
	reference := map[string]int{"FreeBSD GEOM ELI": 1, "Windows BitLocker": 0, "Linux LUKS": 0}
	execTime := time.Since(start)
	if !reflect.DeepEqual(test1, reference) {
		t.Errorf("test failed in %s, expected %v, got %v", execTime, reference, test1)
	} else {
		t.Logf("test passed in %s, expected %v, got %v", execTime, test1, test1)
	}
}

func TestDeviceFile5000(t *testing.T) {
	var totalExecTime time.Duration
	start := time.Now()
	for iter := 0; iter < 5000; iter++ {
		_ = toolDetection(testFilePath, 512)
		_ = map[string]int{"FreeBSD GEOM ELI": 1, "Windows BitLocker": 0, "Linux LUKS": 0}

	}

	totalExecTime = time.Since(start)
	t.Logf("Total execution time: %s, average exec time: %s", totalExecTime, totalExecTime/5000)
}
