import unittest

class Solution(object):
    def allCellsDistOrder(self, R, C, r0, c0):
        """
        :type R: int
        :type C: int
        :type r0: int
        :type c0: int
        :rtype: List[List[int]]
        """

        queue = [[r0, c0]]
        visited = dict()
        res = []
        while len(queue) > 0:
            top = queue.pop(0)
            key = "%d:%d" % (top[0], top[1])
            if visited.get(key):
                continue
            res.append(top)
            visited[key] = True

            if top[1] - 1 >= 0 and not visited.get("%d:%d" % (top[0], top[1] - 1)):
                queue.append([top[0], top[1] - 1])
            if top[0] - 1 >= 0 and not visited.get("%d:%d" % (top[0] - 1, top[1])):
                queue.append([top[0] - 1, top[1]])
            if top[0] + 1 < R and not visited.get("%d:%d" % (top[0] + 1, top[1])):
                queue.append([top[0] + 1, top[1]])
            if top[1] + 1 < C and not visited.get("%d:%d" % (top[0], top[1] + 1)):
                queue.append([top[0], top[1] + 1])

        return res

class Test(unittest.TestCase):
    def test_a(self):
        s = Solution()
        print(s.allCellsDistOrder(1,2,0,0))

    def test_b(self):
        s = Solution()
        print(s.allCellsDistOrder(2,2,0,1))

    def test_c(self):
        s = Solution()
        print(s.allCellsDistOrder(2,3,1,2))

if __name__ == "__main__":
    unittest.main()