package stats


func NbaStoreAll(s string) {

	if s == "" {
		logf("NbaStoreAll", "Cannot process empty season string.")
	} else {
		
		players := NbaGetPlayers(s)

		if players != nil {
			NbaStorePlayers(players)
		}

		teams := NbaGetTeams(s)
		
		if teams != nil {
			NbaStoreTeams(teams)
		}
		
		NbaStoreSeason(s)

	}

} // NbaStoreAll
