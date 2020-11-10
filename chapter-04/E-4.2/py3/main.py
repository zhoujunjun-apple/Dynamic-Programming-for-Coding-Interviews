import sys

gCount = 0

def findMiniCost(costs: list, start: int, end: int)->int:
    """find the minimal cost in range [start, end]"""
    global gCount
    gCount += 1
    if start >= end:
        print(f"{start}->{end}.(no need to move)")
        return 0
    if start + 1 == end:
        ret = costs[start][end]
        print(f"{start}->{end}({ret})")
        return ret
    
    minIdx = start
    minVal = costs[start][end]
    for idx in range(start+1, end):
        leftVal = findMiniCost(costs=costs, start=start, end=idx)
        rightVal = findMiniCost(costs=costs, start=idx, end=end)
        idxVal = leftVal + rightVal
        if idxVal < minVal:
            minVal = idxVal
            minIdx = idx
    print(f"{start}->{minIdx}({leftVal});{minIdx}->{end}({rightVal})")
    return minVal


def DPMiniCost(costs: list)->int:
    from copy import deepcopy

    dp = deepcopy(costs)
    colNum = len(costs[0])
    
    for colIdx in range(1, colNum):
        # column by column from left to right
        for rowDelta in range(1, colIdx+1):
            # for each column, from midden to head
            rowIdx = colIdx - rowDelta
            baseCost = costs[rowIdx][colIdx]
            for k in range(1, colIdx-rowIdx):
                kVal = dp[rowIdx][colIdx-k] + dp[colIdx-k][colIdx]
                if kVal < baseCost:
                    baseCost = kVal
            dp[rowIdx][colIdx] = baseCost
    return dp[0][colNum-1]



if __name__ == "__main__":
    tickets = [
        [0, 10, 75, 94],
        [-1, 0, 35, 50],
        [-1, -1, 0, 80],
        [-1, -1, -1, 0],
    ]
    # ret = findMiniCost(
    #     costs=tickets,
    #     start=0,
    #     end=3,
    # )
    ret = DPMiniCost(costs=tickets)
    print(f'minimal cost: {ret}, call times: {gCount}\n')