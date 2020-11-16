from main import brute_force, max_asc_substring_dp
from main import is_valid_parenthese, max_valid_parenthese

def TestBruteForce():
    ts = [
        {
            'input': '124241',
            'output': 6,
        },
        {
            'input': '0',
            'output': 1,
        },
        {
            'input': '1234522218',
            'output': 10,
        },
        {
            'input': '12345810',
            'output': 4,
        },
    ]

    for tt in ts:
        inp = tt['input']
        oup = tt['output']

        got = brute_force(digits=inp)
        if got != oup:
            print('faild. input={}, expected={}, got={}'.format(inp, oup, got))
    

def TestMaxAscSubstringDP():
    ts = [
        {
            'input': '987654321',
            'output': 1,
        },
        {
            'input': '1234123412345',
            'output': 5,
        },
        {
            'input': '123145689001',
            'output': 6,
        },
    ]

    for t in ts:
        inp = t['input']
        oup = t['output']

        got = max_asc_substring_dp(s=inp)
        if got != oup:
            print('faild. s={}, expected={}, got={}'.format(inp, oup, got))


def TestIsValidParenthese():
    ts = [
        {
            'input': '((',
            'output': False,
        },
        {
            'input': '())',
            'output': False,
        },
        {
            'input': '()',
            'output': True,
        },
        {
            'input': '(()()(()()))',
            'output': True,
        },
    ]

    for t in ts:
        inp = t['input']
        oup = t['output']

        got = is_valid_parenthese(s=inp)
        if got != oup:
            print('input={}, expected={}, got={}'.format(inp, oup, got))


def TestMaxValidParenthese():
    ts = [
        {
            'input': '(()(((()()',
            'output': 4,
        },
        {
            'input': ')())()()',
            'output': 4,
        },
        {
            'input': '((',
            'output': 0,
        },
        {
            'input': '))',
            'output': 0,
        },
        {
            'input': '(()',
            'output': 2,
        },
        {
            'input': '))()(',
            'output': 2,
        },
        {
            'input': '(()())',
            'output': 6,
        },
        {
            'input': ')())()())((())))',
            'output': 6,
        },
    ]

    for t in ts:
        inp = t['input']
        oup = t['output']

        got = max_valid_parenthese(s=inp)
        if got != oup:
            print('input={}, expected={}, got={}'.format(inp, oup, got))

if __name__ == "__main__":
    # TestBruteForce()
    # TestMaxAscSubstringDP()
    # TestIsValidParenthese()
    TestMaxValidParenthese()