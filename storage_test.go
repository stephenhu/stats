package stats

import (
	"testing"
)


func TestStoreGameDay(t *testing.T) {
	StoreGameDay("20200130")
} // TestStoreGameDay


func TestStoreSeason(t *testing.T) {
	StoreSeason("1819")
} // TestStoreSeason


func TestStoreFromDay(t *testing.T) {
	StoreFromDay("20200128")
} // TestStoreFromDay
