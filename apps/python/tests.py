import unittest

def add_six(i):
    return i + 6

class TestAddSix(unittest.TestCase):

    def test_add_six(self):
        self.assertEqual(add_six(2), 8)

if __name__ == '__main__':
    unittest.main()
