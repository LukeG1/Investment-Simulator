# Goals

The objective with this project is to asses household investment strategies in their statistical context. That means instead of simply looking at the mean return for a given investment like an online calculator, or many of the axioms you hear about would, you asses all investment outcomes as a distribution. I plan to achieve this with Monte Carlo Simulation to collect a sample of outcomes for any 'query'. This allows us to consider what I see as the most important statistics for a lifelong investment strategy: the Probability of Portfolio Failure $P(balance ≤ 0)$@life_expectancy (hereby refered to as PPF). This metric can then be used to optmize what I see as a rational investment strategy based on a users abilities and needs, setting aside the notion of risk tollerance.

The reason for approaching the problem this way this is two fold:

1. I have a suspicion that many people are undervaluing the risk that comes with high volatility investmnets, espiecally for young investors. I've heard it reccomended that you shouldn't even consider bonds in your 20s and 30s, but I question if the volatility, and therefore potientally increased PPF, that introduces is worth the inarguably higher mean value. Again, for the simple rational investor minimizing PPF at expected life expectancy should be their only priority, even if that means accepting some fraction of the mean value of you could achieve with a higher risk strategy
2. I don't think investment advice is individualized enough, even if you hired a personal financial advisor they aren't running these simulations, instead just using the known reccomendations, but with this system, you could brute force a strategy for anyone in any finanical situation with a much finer granularity, and actual evidence to back it up. This also comes with the ability to reevaluate constantly based on ever changing political and economic consequences.

# Implementation

## Data Aggregation

The dificulty with simulating something over a long term as a person's life is the quantity of data you need to generate to get a stable result. To minimize this data I run simulations in one year batches, aggregating the data every time, this way I only need to ever store one complete year's worth of data. In addition to this batching, I've also implemented a heuristic to check if the data has 'stabalized' before a simulation count limit (restricted by memory) based on the range of the total mean in recent window. (note it is unclear if this heuristic is realistic to keep long term)

I'm also considering switching this around a little, noting how effective the heuristc is, what if I made a global threadsafe data aggregator that did full simulations at a time, but when a year is stabalized, which the early ones will much faster, I aggregate them and just stop saving data, it introduces some risk of overflow if even the early years are highly variable, but it's worth considering, I would need to estimate how much storage is actually needed in a worst casenario, but it would make development much simpler

lastly there is my favorite but not most percise, the bucket approach, here you would have a couple of maps to allow you to store an array for each account's investment, for each year, say you had 500k values in that array, each could represent a bucket of say 100 dollars, so you could represent ammounts up to 50 million, if a sim had a Roth IRA's bonds investment at 50 grand, you would increment the index @ $\lfloor \frac{new value}{bucket size} \rfloor$ this obviously has the risk of missing outliers, limited precision, and a gaurnteed minimum of memory use, I estimate minimum 5gb (could possibly take advantage of sparseness or a hash table to improve) but it would be pretty easy to implement and lets you do the most intresting math on the data after becuase it really has the littlest aggregation, while it limits the usefulness of the project for smaller applications, over the course of a lifetime precision beyond 100 dollars feels silly, plus a simpler representation could be automatically used for simpler calculations
