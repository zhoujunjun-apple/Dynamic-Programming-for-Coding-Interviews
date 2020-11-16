def brute_force(digits: str)->int:
    """暴力求解"""
    digits_num = [int(i) for i in digits]
    ret = 1

    for left in range(len(digits)):
        for right in range(left+ret, len(digits)):
            # check if digits[left, right] is match the pattern
            left_right = right - left  # the digits num in [left, right]
            half = int(left_right / 2)  # the half length of left_right
            mid = left + half 

            left_val = sum(digits_num[left:mid])
            if left_right % 2 == 0:
                # digits[left, right] has odd digits
                right_val = sum(digits_num[mid: right])
            else:
                # digits[left, right] has even digits
                right_val = sum(digits_num[mid+1: right])
            
            if left_val == right_val:
                if left_right > ret:
                    ret = left_right
    return ret

def max_asc_substring_dp(s: str):
    """最长递增字串DP算法"""
    import numpy as np

    dp = np.ones(shape=len(s), dtype=np.int64)
    for i in range(1, len(s)):
        if s[i] > s[i-1]:
            dp[i] = dp[i-1] + 1
    
    return max(dp)

def is_valid_parenthese(s: str)->bool:
    """判断字符串是否是有效的括号序列"""
    if len(s) % 2 != 0:
        return False
    
    stack = []
    for c in s:
        if c == '(':
            stack.append(c)
        elif c == ')':
            if stack:
                stack = stack[:-1]
            else:
                return False
        else:
            raise Exception('invalid char {}'.format(c))
    return len(stack) == 0


def max_valid_parenthese(s: str)->int:
    """找出最长有效括号字串长度(错误解法)"""

    max_len = 0
    cur_len = 0

    stack = []
    for c in s:
        if c == '(':
            stack.append(c)
        else:
            if stack:
                top = stack[-1]
                if top == '(':
                    cur_len += 2
                    stack = stack[:-1]
                else:
                    max_len = max(max_len, cur_len)
                    cur_len = 0
            else:
                max_len = max(max_len, cur_len)
                cur_len = 0
    return max(max_len, cur_len)


if __name__ == "__main__":
    pass