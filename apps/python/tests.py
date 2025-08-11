import unittest
from main import add_six

class TestAddSix(unittest.TestCase):

    def test_add_six(self):
        self.assertEqual(add_six(2), 8)
        self.assertEqual(add_six(0), 6)
        self.assertEqual(add_six(-10), -4)
        self.assertEqual(add_six(5.5), 11.5)

if __name__ == '__main__':
    unittest.main()
