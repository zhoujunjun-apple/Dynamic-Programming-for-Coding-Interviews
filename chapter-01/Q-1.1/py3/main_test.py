import unittest

from main import factorialFunc


class TestMain(unittest.TestCase):
    def test_factorial(self):
        tt = [
            {
                'input': 0,
                'output': 0,
            },
            {
                'input': 1,
                'output': 1,
            },
            {
                'input': 2,
                'output': 2,
            },
            {
                'input': 3,
                'output': 6,
            },
            {
                'input': 4,
                'output': 24,
            },
            {
                'input': 5,
                'output': 120,
            },
        ]
        for titem in tt:
            self.assertEqual(factorialFunc(titem['input']), titem['output'])


if __name__ == "__main__":
    unittest.main()