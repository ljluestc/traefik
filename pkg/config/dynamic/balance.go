package dynamic

// BalancerStrategy defines the strategy for a load balancer.
type BalancerStrategy string

const (
	// BalancerStrategyRoundRobin is the round-robin load balancing strategy.
	BalancerStrategyRoundRobin BalancerStrategy = "roundRobin"
	
	// BalancerStrategyLeastConn is the least connection load balancing strategy.
	BalancerStrategyLeastConn BalancerStrategy = "leastConn"
	
	// BalancerStrategyRandom is the random load balancing strategy.
	BalancerStrategyRandom BalancerStrategy = "random"
	
	// BalancerStrategyIPHash is the IP hash load balancing strategy.
	BalancerStrategyIPHash BalancerStrategy = "ipHash"
)
