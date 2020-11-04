def selectSort(arr: list):
    putMaxToEnd(arr=arr, idx=len(arr)-1)

def putMaxToEnd(arr: list, idx: int):
    """
    find the max in arr[0, idx] and swap it with arr[idx]
    """
    if idx <= 0:
        return
    
    maxVal = arr[0]
    maxIdx = 0
    for i in range(0, idx+1):
        now = arr[i]
        if now > maxVal:
            maxVal = now
            maxIdx = i
    arr[maxIdx] = arr[idx]
    arr[idx] = maxVal
    putMaxToEnd(arr=arr, idx=idx-1)

def bubbleSort(arr: list):
    swapMaxToEnd(arr=arr, idx=len(arr)-1)

def swapMaxToEnd(arr: list, idx: int):
    """
    compare the adjacent values in arr[0, idx] and swap them if left value is not less than right value 
    """
    if idx <= 0:
        return
    
    for i in range(0, idx):
        left = arr[i]
        right = arr[i+1]
        if left > right:
            temp = left
            arr[i] = right
            arr[i+1] = temp
    swapMaxToEnd(arr=arr, idx=idx-1)