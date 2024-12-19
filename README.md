# Goals

The objective with this project is to asses household investment strategies in their statistical context. That means instead of simply looking at the mean return for a given investment like an online calculator, or many of the axioms you hear about would, you asses all investment outcomes as a distribution. I plan to achieve this with Monte Carlo Simulation to collect a sample of outcomes for any 'query'. This allows us to consider what I see as the most important statistics for a lifelong investment strategy: the Probability of Portfolio Failure $P(balance ≤ 0)$@life_expectancy (hereby refered to as PPF). This metric can then be used to optmize what I see as a rational investment strategy based on a users abilities and needs, setting aside the notion of risk tollerance.

The reason for approaching the problem this way this is two fold:

1. I have a suspicion that many people are undervaluing the risk that comes with high volatility investmnets, espiecally for young investors. I've heard it reccomended that you shouldn't even consider bonds in your 20s and 30s, but I question if the volatility, and therefore potientally increased PPF, that introduces is worth the inarguably higher mean value. Again, for the simple rational investor minimizing PPF at expected life expectancy should be their only priority, even if that means accepting some fraction of the mean value of you could achieve with a higher risk strategy
2. I don't think investment advice is individualized enough, even if you hired a personal financial advisor they aren't running these simulations, instead just using the known reccomendations, but with this system, you could brute force a strategy for anyone in any finanical situation with a much finer granularity, and actual evidence to back it up. This also comes with the ability to reevaluate constantly based on ever changing political and economic consequences.

# Implementation

## Data Aggregation

The dificulty with simulating something over a long term as a person's life is the quantity of data you need to generate to get a stable result. Given that the values I'm dealing with don't need to be incredibly particularly precise, I will use an approximation of any values that would require me to store this data points. My DistributionLearner, inspired by [this python library](https://github.com/cxxr/LiveStats/blob/master/livestats/livestats.py), should give me the best balance of time and space complexity, with $O(1)$ updates and minimal storage needs. I've also implemented a Heuristic to check the stability of the learned distribtuion. By looking at the range of a window of recent mean values (subject to change) and stopping when that range is below an acceptable value. At any time this DistributionLearner can be summarized into a set of statistics that should do an adequate job of representing the data for any analysis I would want to do.

Each investment within each account would need to be tracked with one of these DistributionLearners, which could then be summarized and checked for stability incrementally. The statistics from these summaries could then be saved to a csv file that the frontend could periodically access. This means that it could automatically update independantley of the simulation, making the process easier to follow visually (you could see how close to stable each year's data is as it runs). if the stability check is met at summary time it could then cancel execution.

## Modeling Real World Constructs

### Household

A household is what all constructs discussed below will eventually feed into or out of. ...

### Household Factor

Household factors are some kind of historically or anecdotally informed influence on the simulation from inside of a household, this includes things like income, variable budgets, children, social security, etc. ...

### Economic Factors

Economic factors are some kind of historically informed influence on the simulation from outside of a household, this includes things like stock/bond/housing market performance, and inflation. An economic factor includes the raw data that its based on, a name, and a way to sample random data from it.

### Accounts

The Account is really central to this project, since it's kind of the lowest level construct. An Account is implmeented like an abstract class, and can be used to represent anything from a high yield savings account or a Roth IRA, to a mortgage. The abstract nature of an account means that a new one can be added with as little boilerplate as possible. The current version can be implemented with just an AllowedContribution function, that informs the inner abstract account's deposit function. A similar strategy will be used later in withdrawals. This strategy also let me make a deposit function `account.Deposit(economicFactor string, amount float64)` which works for any account.

Central to an account is a map of string to Investment called Investments. The key in that map is the name of the economic factor the investment is in, and the investment itself has data like the balance, year's stats, and a refrence to the economic factor that is shared among all accounts. This mean's an account like the HYSA can have an investment in 'cash', a RothIRA can have one in stocks and bonds, and a mortgage can have on in the housing market. The allowed economic factors for an account are defined on creation.

# Current Todo List:

- [x] Move historic data to it's own package and refactor EconomicFactors
- [ ] Flesh out SimulationResults structure
- [ ] Implement SimulationResults in SimpleSimulation
- [ ] Make SimpleSimulation anynchronus with periodic polling of current SimulationResults
- [ ] Display live stability check as the simulation runs
- [ ] Find a easier frontend UI library
- [ ] Write actual tests for existing code
- [ ] Remove dead code
- [ ] Flesh out Household / HouseholdFactors
- [ ] Consider go routines
- [ ] Replace hardcoded historic data with modular config of some kind? start with SQLLite
- [ ] Implement more accounts
- [ ] Implement taxes
- [ ] Implement social security
- [ ] write a function to export SimulationResults as csv
- [ ] Figure out what a "strategy" looks like
- [ ] Implement a simulation on a strategy
- [ ] ???
