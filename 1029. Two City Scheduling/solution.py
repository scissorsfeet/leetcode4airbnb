import unittest

class Solution(object):
    def twoCitySchedCost(self, costs):
        """
        :type costs: List[List[int]]
        :rtype: int
        """
        n = len(costs)
        cost = 0
        for x in costs:
            cost += x[0]

        costs = sorted(costs, key=lambda x: x[0]-x[1], reverse=True)

        for i in range(n/2):
            cost -= costs[i][0]-costs[i][1]

        return cost

class Test(unittest.TestCase):
    def test_a(self):
        s = Solution()
        assert(s.twoCitySchedCost([[10, 20], [30, 200], [400, 50], [30, 20]]) == 110)

if __name__ == "__main__":
    unittest.main()
