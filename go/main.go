package main

import (
	"github.com/gin-gonic/gin"
)

// userData contains user's entitlement of shares and their
// snapshot of the pool accumulator.
type userData struct {
	// number of shares of the user in a pool.
	shares int
	// user's most recently updated accumulator snapshot.
	accumulatorSnapshot int
}

// pool defines the data structure for holding data in the pool.
type pool struct {
	// liqudity is the total liquidity in the pool.
	// must add up to the sum of shares of all users in userSnapshots
	liquidity int
	// global per-pool ever-increasing accumulator.
	// This is a per-unit of liquidity value.
	globalAccumulator int

	// userAddress -> data
	userSnapshots map[string]userData
}

// Mocked data for demonstration purposes
var poolData map[string]pool = map[string]pool{
	"1": {
		liquidity:         300,
		globalAccumulator: 800,

		userSnapshots: map[string]userData{
			"userA": {
				shares:              300,
				accumulatorSnapshot: 0,
			},
		},
	},
	"2": {
		liquidity:         300,
		globalAccumulator: 800,

		userSnapshots: map[string]userData{
			"userA": {
				shares:              300,
				accumulatorSnapshot: 400,
			},
		},
	},
	"3": {
		liquidity:         400,
		globalAccumulator: 500,

		userSnapshots: map[string]userData{
			"userA": {
				shares:              200,
				accumulatorSnapshot: 300,
			},
			"userB": {
				shares:              200,
				accumulatorSnapshot: 0,
			},
		},
	},
}

// Endpoint for Claimable Incentive Rewards
func getClaimableRewards(c *gin.Context) {
	// TODO 1: uncomment and use parameters to read mock data.
	// Perform proper validation as you see fit.
	// poolID := c.Query("pool_id")
	// userAddress := c.Query("user_address")

	// TODO 2: uncoment placeholders and use the read mock data to calculate claimable rewards.
	// Calculate claimable rewards based on your logic (replace with actual calculation)
	// claimableRewards := calculateClaimableRewards(
	// 	...,
	// 	...,
	// 	...,
	// )

	// TODO 3: uncomment once claimableRewards is calculated.
	// c.JSON(200, gin.H{"claimable_rewards": fmt.Sprintf("%d", claimableRewards)})
}

// TODO 4 (HARD): copy your originan endpoint logic and add a few lines to make it mutative
// Mutative Endpoint to claim rewards.
func claimRewards(c *gin.Context) {

}

func calculateClaimableRewards(poolSnapshot int, userSnapshot int, userShares int) int {
	// TODO: Replace this with your actual claimable rewards calculation logic
	return 0
}

func main() {
	r := gin.Default()

	// Define routes
	r.GET("/claimable_rewards", getClaimableRewards)
	r.POST("/claim_rewards", claimRewards)

	// Run the server
	r.Run(":8080")
}
