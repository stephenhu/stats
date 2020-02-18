package stats

import (
  "testing"
)


func TestGetEraUnsupported(t *testing.T) {
	GetEra("bad.era")
} // TestGetEraUnsupported


func TestGetEra(t *testing.T) {
	GetEra(PROFILE_SIMPLE_ERA)
} // TestGetEra
