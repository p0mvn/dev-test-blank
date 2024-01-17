### Rewards based on accumulator task

One of the common problems for blockchains is distributing rewards for fractional ownership of assets in an efficient way. For example, consider providing liquidity into a standard liquidity pool. Upon depositing assets into the pool, a user gets back an amount of liquidity provision (LP) shares that is proportional to their ownership in the pool.

As a protocol, Osmosis wants to incentivize users to provide this liquidity through incentive rewards. At regular intervals (i.e. every day), the protocol mints OSMO tokens that get pipelined into pools based on governance-approved proportions.

The problem now becomes determining how to distribute rewards efficiently to each user. Imagine that there are thousands or even millions of users. A naive implementation would iterate over all users and determine the appropriate allocation to each at the time of distribution.

Instead, we define a framework that allows users to claim rewards on-demand in constant time. Read the "Sources" section at the bottom for more context about this framework. To achieve constant time claiming, the chain defines a per-pool accumulator that represents incentive reward growth per unit of liquidity in the pool. Its value is ever-increasing. At the time of reward distribution, we allocate a per-unit-of-liquidity distribution amount to the pool accumulator.

For each user, whenever they LP into the pool, the system takes a snapshot of the pool accumulator and associates it with the user. Additionally, if the user adds, removes liquidity, or claims rewards, their accumulator snapshot is updated. The system also persists the number of shares users have after the update.

For example, assume that a pool has 900 units of liquidity from various users and 100 units from user A. At the start, each user's individual snapshot is equal to 0. Pool receives 10000 OSMO. Then, the global accumulator becomes 10000 / 1000 = 10 OSMO / unit of liquidity.

Now, user A decides to claim their rewards. They take the current value of the pool accumulator (10) and subtract their snapshot from it. Finally, they multiply the difference by the number of shares that they have (10 - 0 = 10 * 100 = 1000 OSMO). To reflect that they claimed the amount and to avoid being able to claim next time, the user's pool accumulator snapshot is updated to 10.

In order to incentivize users to LP into an Osmosis pool the FE needs to ability to display APR for that and all other available pools [4]. Additionally, it needs to show the users what their claimable rewards are.

Requirement:
- Get claimable rewards function: 
   * Inputs: Pool ID, user address
   * Outputs: Incentive rewards denominated in USDC

Note that each pool has a unique ID.

For your main task, define the function with the above requirements. It has to query the chain to keep the relevant data in sync with the latest state. You have the flexibility to expose chain queries as needed. Define the chain query APIs needed to pull the relevant data. For simplicity, feel free to mock and read these outputs from JSON files in your project.

Expose HTTP endpoints for the front end and other clients to query.

You will be evaluated on design choices, algorithm selection, design patterns, testability, readability, and security of your code.

Most importantly, we hope that you have a great experience working through this problem. Good luck!

## Follow-Up

Let's now assume that this is a mutative operation to claim rewards. How would you modify your function to ensure that it works correctly per protocol requirements?

### Sources

- https://uploads-ssl.webflow.com/5ad71ffeb79acc67c8bcdaba/5ad8d1193a40977462982470_scalable-reward-distribution-paper.pdf
- https://drops.dagstuhl.de/storage/01oasics/oasics-vol071-tokenomics2019/OASIcs.Tokenomics.2019.10/OASIcs.Tokenomics.2019.10.pdf
