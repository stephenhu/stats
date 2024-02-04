# last 10

based on the given data structures, the last 10 games per team is sort of hard
to calculate, the algorithm works like this:

1. go to the latest day that has boxscores
1. find all games, track per team up to 10 games
1. calculate the wins and losses, they cannot be more than 10