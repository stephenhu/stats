package stats

import (
	"testing"
)


func TestStoreGameDay(t *testing.T) {

	t.Skip("deprecate in favor of data.nba.net")

	StoreGameDay("20200130")
	
} // TestStoreGameDay


func TestStoreSeason(t *testing.T) {
	
	t.Skip("deprecate in favor of data.nba.net")

	StoreSeason("1819")

} // TestStoreSeason


func TestStoreFromDay(t *testing.T) {
	
	t.Skip("deprecate in favor of data.nba.net")

	StoreFromDay("20200128")

} // TestStoreFromDay
