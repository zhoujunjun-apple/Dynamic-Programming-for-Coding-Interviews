gcount = 0

def wayNumRecursive(rowIdx: int, colIdx: int)->int:
    """the total number of ways which moves from arr[0][0] to arr[rowIdx][colIdx]"""
    global gcount
    gcount += 1

    if rowIdx < 0 or colIdx < 0:
        return 0
    if rowIdx == 0 or colIdx == 0:
        return 1
    return wayNumRecursive(rowIdx, colIdx-1) + wayNumRecursive(rowIdx-1, colIdx)


def wayNumDP(rows: int, cols: int)->int:
    import numpy as np
    
    dp = np.ones(shape=(rows, cols), dtype=np.int32)
    for ri in range(1, rows):
        for ci in range(1, cols):
            dp[ri][ci] = dp[ri, ci-1] + dp[ri-1, ci]
    return dp[rows-1][cols-1]


if __name__ == "__main__":
    recursiveWay = wayNumRecursive(rowIdx=10, colIdx=10)
    dpWay = wayNumDP(rows=11, cols=11)
    print("recursive: {}({} calls), dp: {}\n".format(recursiveWay, gcount, dpWay))