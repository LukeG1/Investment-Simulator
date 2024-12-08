# Goals

The objective with this codebase is to asses household investment strategies in their statistical context, that means instead simply using a mean return for a given investment type like an online calculator would. Here I will use Monte Carlo Simulation to give a distribution of results for any 'query', this opens up the opportunity to look not at average returns, but instead $P(balance ≤ 0)$ over an entire lifetime of investing, and see how likley a person is to run out of money in their retirement, which can then be minimized against when choosing a strategy.

The reason for approaching the problem like this is two fold:

1. I have a suspicion that many people are undervaluing the risk that comes with high volatility, I have heard reccomended that you shouldn't even consider bonds in your 20s and 30s, but I question if that volatility that introduces is woth the inarguably higher mean value. Again, for the vast majority of people minimizing $P(balance ≤ 0)$ at expected life expectancy should be their highest priority, even if that means accepting some fraction of the mean value of your portfolio at that time
2. I don't think investment advice is individualized enough, even if you hired a personal financial advisor they aren't running these simulations, they are using the known reccomendations, but with this system, you could brute force a strategy for anyone in any finanical situation that gives them the best chance

# Implementation
