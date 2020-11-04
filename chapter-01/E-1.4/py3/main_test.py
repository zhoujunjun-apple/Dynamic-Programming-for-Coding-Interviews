import unittest

from main import selectSort, bubbleSort


class TestMain(unittest.TestCase):
    def test_selectSort(self):
        tt = [
            {
                'input': [0, ],
                'output': [0, ],
            },
            {
                'input': [2, 1, 0],
                'output': [0, 1, 2],
            },
            {
                'input': [1, 4, 3, 2, 5, 9, 6],
                'output': [1, 2, 3, 4, 5, 6, 9],
            },
            {
                'input': [],
                'output': [],
            },
        ]
        for ts in tt:
            inpdata = ts['input']
            selectSort(arr=inpdata)
            self.assertEqual(inpdata, ts['output'])
    
    def test_bubbleSort(self):
        tt = [
            {
                'input': [0, ],
                'output': [0, ],
            },
            {
                'input': [2, 1, 0],
                'output': [0, 1, 2],
            },
            {
                'input': [1, 4, 3, 2, 5, 9, 6],
                'output': [1, 2, 3, 4, 5, 6, 9],
            },
            {
                'input': [],
                'output': [],
            },
        ]
        for ts in tt:
            inpdata = ts['input']
            bubbleSort(arr=inpdata)
            self.assertEqual(inpdata, ts['output'])


if __name__ == "__main__":
    unittest.main()