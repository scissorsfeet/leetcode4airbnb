# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def nextLargerNodes(self, head):
        """
        :type head: ListNode
        :rtype: List[int]
        """
        l = []
        while head != None:
            l.append(head.val)
            head = head.next

        s = []
        ans = [i for i in range(0,len(l))]
        for i in range(0,len(l)):
            num = l[i]
            while len(s)>0 and num > s[len(s)-1][1]:
                tmp = s.pop()
                ans[tmp[0]] = num
            s.append([i, num])
        while len(s) > 0 :
            tmp = s.pop()
            ans[tmp[0]] = 0

        return ans
