# Goals

The objective with this codebase is to asses household investment strategies in their statistical context, that means instead simply using a mean return for a given investment type like an online calculator would. Here I will use Monte Carlo Simulation to give a distribution of results for any 'query', this opens up the opportunity to look not at average returns, but instead $P(balance ≤ 0)$ (hereby refered to as PPF) over an entire lifetime of investing, and see how likley a person is to run out of money in their retirement, which can then be minimized against when choosing a strategy.

The reason for approaching the problem like this is two fold:

1. I have a suspicion that many people are undervaluing the risk that comes with high volatility, I have heard reccomended that you shouldn't even consider bonds in your 20s and 30s, but I question if that volatility that introduces is woth the inarguably higher mean value. Again, for the vast majority of people minimizing PPF at expected life expectancy should be their highest priority, even if that means accepting some fraction of the mean value of your portfolio at that time
2. I don't think investment advice is individualized enough, even if you hired a personal financial advisor they aren't running these simulations, they are using the known reccomendations, but with this system, you could brute force a strategy for anyone in any finanical situation that gives them the best chance

# Implementation

## Data Aggregation

The dificulty with simulating something over such a long term as a person's life is the quantity of data, you need a huge number of samples to get an accurate idea of the values 80 some years later with any ammount of detirminisim. To acomplish this I run simulations one year at a time. Once a year is done I take a snapshot consisting of some basic summary stats, including that years PPF, then wait move on to the next year, this also opens the door for a heuristic of checking if the data has 'stabalized' within a year, and stopping simulation prematurely to save resources, I would still need to figure out how to gracefully move on to the next year that may need more or less data, but the concept is there. Low variance years in the begining would clearly require less data to get an accurate snapshot.
