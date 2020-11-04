import unittest

from main import add_previous


class TestMain(unittest.TestCase):
    def test_add_previous(self):
        tt = [
            {
                'input': [0, ],
                'output': [0, ],
            },
            {
                'input': [0, 1], 
                'output': [0, 1],
            },
            {
                'input': [0, 1, 2],
                'output': [0, 1, 3],
            },
            {
                'input': [1, 3, 2, 4, 3, 5, 4, 6],
                'output': [1, 4, 6, 10, 13, 18, 22, 28],
            },
            {
                'input': [0, 0, 1, -1, 2, -2, 3, -3],
                'output': [0, 0, 1, 0, 2, 0, 3, 0],
            },
            {
                'input': [1, 2, 3, 4, -4, -3, -2, -1],
                'output': [1, 3, 6, 10, 6, 3, 1, 0],
            },
        ]
        for titem in tt:
            inpdata = titem['input']
            add_previous(arr=inpdata)
            self.assertEqual(inpdata, titem['output'])


if __name__ == "__main__":
    unittest.main()