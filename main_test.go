package main

import (
	"testing"

	"github.com/Sumit0x00/SHA1-Password-Cracker/pass"
)

func TestHash_Superman(t *testing.T) {
	actual := pass.CrackSHA1Hash("18c28604dd31094a8d69dae60f1bcd347f1afc5a", false)

	if actual == "superman" {
		t.Logf("success : expected %v , Got %v ", "superman", actual)
	} else {
		t.Errorf("failed : expectedt %v , Got %v ", "superman", actual)
	}
}

func TestHash_q1w2e3r4t5(t *testing.T) {
	actual := pass.CrackSHA1Hash("5d70c3d101efd9cc0a69f4df2ddf33b21e641f6a", false)

	if actual == "q1w2e3r4t5" {
		t.Logf("success : expected %v , Got %v ", "q1w2e3r4t5", actual)
	} else {
		t.Errorf("failed : expectedt %v , Got %v ", "q1w2e3r4t5", actual)
	}
}

func TestHash_bubbles1(t *testing.T) {
	actual := pass.CrackSHA1Hash("b80abc2feeb1e37c66477b0824ac046f9e2e84a0", false)

	if actual == "bubbles1" {
		t.Logf("success : expected %v , Got %v ", "bubbles1", actual)
	} else {
		t.Errorf("failed : expectedt %v , Got %v ", "bubbles1", actual)
	}
}

func TestHash_NOT_FOUND(t *testing.T) {
	actual := pass.CrackSHA1Hash("03810a46a2c1a0eae58d9332f01c32bdcec9a01a", false)
	if actual == "PASSWORD NOT IN DATABASE" {
		t.Logf("success: expected %v, got %v", "PASSWORD NOT IN DATABASE", actual)
	} else {
		t.Errorf("failed: expected %v, got %v", "PASSWORD NOT IN DATABASE", actual)
	}
}

func TestSalt_superman(t *testing.T) {
	actual := pass.CrackSHA1Hash("53d8b3dc9d39f0184144674e310185e41a87ffd5", true)
	if actual == "superman" {
		t.Logf("success: expected %v, got %v", "superman", actual)
	} else {
		t.Errorf("failed: expected %v, got %v", "superman", actual)
	}
}

func TestSalt_q1w2e3r4t5(t *testing.T) {
	actual := pass.CrackSHA1Hash("da5a4e8cf89539e66097acd2f8af128acae2f8ae", true)
	if actual == "q1w2e3r4t5" {
		t.Logf("success: expected %v, got %v", "q1w2e3r4t5", actual)
	} else {
		t.Errorf("failed: expected %v, got %v", "q1w2e3r4t5", actual)
	}
}

func TestSalt_bubbles1(t *testing.T) {
	actual := pass.CrackSHA1Hash("ea3f62d498e3b98557f9f9cd0d905028b3b019e1", true)
	if actual == "bubbles1" {
		t.Logf("success: expected %v, got %v", "bubbles1", actual)
	} else {
		t.Errorf("failed: expected %v, got %v", "bubbles1", actual)
	}
}
