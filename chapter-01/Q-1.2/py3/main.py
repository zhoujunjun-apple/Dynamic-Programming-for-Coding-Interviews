def add_previous(arr: list):
    def add_recursive(arr: list, idx: int):
        """add in range [0, idx] from arr list and put sum result in arr[idx] """
        if idx <= 0:
            # the 0-th position don't need to calculate
            return
        add_recursive(arr=arr, idx=idx-1)
        # sum(arr[0, idx-1]) has saved at arr[idx-1]
        arr[idx] += arr[idx-1]  # update arr[idx] use arr[idx-1]
        
    add_recursive(arr=arr, idx=len(arr)-1)
