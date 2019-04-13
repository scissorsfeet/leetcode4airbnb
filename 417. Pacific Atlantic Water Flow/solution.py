class Solution(object):
    def pacificAtlantic(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: List[List[int]]
        """
        m = len(matrix)
        if m == 0:
            return []
        n = len(matrix[0])
        if n == 0:
            return []

        p = [[False for _ in range(n)] for _ in range(m)]
        a = [[False for _ in range(n)] for _ in range(m)]

        qp = []
        qa = []
        for i in range(m):
            qp.append((i,0,matrix[i][0],0)) #left
            qa.append((i,n-1,matrix[i][n-1],0)) #right
        for i in range(n):
            qp.append((0,i,matrix[0][i],0))   #top
            qa.append((m-1,i,matrix[m-1][i],0)) #bottom

        def bfs(q, v):
            while len(q) > 0:
                t = q.pop(0)
                x = t[0]
                y = t[1]
                selfHeight = t[2]
                parentHeight = t[3]
                if x < 0 or y < 0 or x >= m or y >= n:
                    continue
                if v[x][y] or parentHeight > selfHeight:
                    continue
                v[x][y] = True
                if x+1 < m:
                    q.append((x+1,y,matrix[x+1][y],selfHeight))
                if x-1 >= 0:
                    q.append((x-1,y,matrix[x-1][y],selfHeight))
                if y+1 < n:
                    q.append((x,y+1,matrix[x][y+1],selfHeight))
                if y-1 >= 0:
                    q.append((x,y-1,matrix[x][y-1],selfHeight))

        bfs(qp, p)
        bfs(qa, a)

        res = []
        for i in range(m):
            for j in range(n):
                if p[i][j] and a[i][j]:
                    res.append([i,j])
        return res

import unittest

class Test(unittest.TestCase):
    def test_a(self):
        s = Solution()
        matrix = [
            [1,2,2,3,5],
            [3,2,3,4,4],
            [2,4,5,3,1],
            [6,7,1,4,5],
            [5,1,1,2,4],
        ]
        self.assertEqual(s.pacificAtlantic(matrix), [[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]])

if __name__ == "__main__":
    unittest.main()
