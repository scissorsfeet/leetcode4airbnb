import unittest

class Solution(object):
    def maxSumTwoNoOverlap(self, A, L, M):
        """
        :type A: List[int]
        :type L: int
        :type M: int
        :rtype: int
        """
        n = len(A)
        return max(self._helper(A, L, M, n), self._helper(A, L, M, n, True))

    def _helper(self, A, L, M, n, reverse=False):
        if reverse:
            A.reverse()
        s = min(L, M)
        l = max(L, M)
        res = rsum = lsum = tmp = 0
        for i in range(n):
            if i < l:
                res += A[i]
                lsum = res
            else:
                lsum -= A[i - l]
                lsum += A[i]
                if i - l >= 0:
                    tmp += A[i-l]
                    if i - l - s >= 0:
                        tmp -= A[i-l-s]
                    if tmp > rsum:
                        rsum = tmp
                    if lsum + rsum > res:
                        res = lsum + rsum
        return res

class Test(unittest.TestCase):
    def test_a(self):
        s = Solution()
        print(s.maxSumTwoNoOverlap([3, 8, 1, 3, 2, 1, 8, 9, 0], 3, 2))
        # assert(s.maxSumTwoNoOverlap([3,8,1,3,2,1,8,9,0], 3, 2) == 23)

    # def test_b(self):
    #     s = Solution()
    #     assert(s.maxSumTwoNoOverlap([0,6,5,2,2,5,1,9,4], 1, 2) == 20)
    #
    # def test_c(self):
    #     s = Solution()
    #     assert(s.maxSumTwoNoOverlap([2,1,5,6,0,9,5,0,3,8], 4, 3) == 31)

if __name__ == "__main__":
    # unittest.main()
    s = Solution()
    print(s.maxSumTwoNoOverlap([3,8,1,3,2,1,8,9,0], 3, 2))