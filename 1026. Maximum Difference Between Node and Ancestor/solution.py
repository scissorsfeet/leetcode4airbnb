import unittest
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x, left, right):
        self.val = x
        self.left = left
        self.right = right

class Solution(object):
    def maxAncestorDiff(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        def helper(r, parentStack, pRes):
            if r is None:
                return
            if len(parentStack) > 0:
                for v in parentStack:
                    res[0] = max(res[0], abs(v-r.val))

            if r.left is not None or r.right is not None:
                parentStack.append(r.val)
                if r.left is not None:
                    helper(r.left, parentStack, pRes)
                if r.right is not None:
                    helper(r.right, parentStack, pRes)
                parentStack.pop()

        res = [0]
        parentStack = []
        helper(root, parentStack, res)
        return res[0]


class Test(unittest.TestCase):
    def test_a(self):
        s = Solution()
        r1 = TreeNode(2, None, None)
        r2 = TreeNode(3, None, None)
        root = TreeNode(1, r1, r2)
        print(s.maxAncestorDiff(root))

if __name__ == "__main__":
    unittest.main()