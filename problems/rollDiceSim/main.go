package main

const DIV = 1000000007

func dieSimulator(n int, rollMax []int) int {
    dp := make([][6][16]int, n)
    for i := range dp{
        for j := range dp[i] {
            for k := range dp[i][j]{
                dp[i][j][k] = -1
            }
        }
    }
    

    return roll(n, 0, 0, 0, rollMax, dp)  
}

func roll(n int, cur int, lastRoll int, consecutive int, rollMax []int, memo [][6][16]int) int {
    if n == cur{
        return 1
    }

    if memo[cur][lastRoll][consecutive] != -1 {
        return memo[cur][lastRoll][consecutive]
    }
    
    total := 0
    for i := 0; i < 6; i++{
        if i == lastRoll && rollMax[i] <= consecutive{
            continue
        } else if i == lastRoll && rollMax[i] > consecutive {
            total += roll(n, cur+1, i, consecutive+1, rollMax, memo )
        } else {
            total += roll(n, cur+1, i, 1, rollMax, memo )
        }
    }
    
    memo[cur][lastRoll][consecutive] = total % DIV
    return memo[cur][lastRoll][consecutive]
}