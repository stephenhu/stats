package stats


func NbaStoreAll(s string) {

	if s == "" {
		logf("NbaStoreAll", "Cannot process empty season string.")
	} else {
		
		players := NbaGetPlayers(s)

		if players != nil {
			
			NbaStorePlayers(players)

			profiles := NbaGetProfiles(s, players)

			if profiles != nil {
				NbaStoreProfiles(profiles)
			}

		}

		teams := NbaGetTeams(s)
		
		if teams != nil {
			
			NbaStoreTeams(teams)

			ranks := NbaGetTeamRanks(s)

			if ranks != nil {
				NbaStoreTeamRanks(ranks)
			}

		}
		
		NbaStoreSeason(s)

	}

} // NbaStoreAll
