class Solution(object):
    def containsNearbyDuplicate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: bool
        """
        m = {}
        numsLen = len(nums)
        for i in range(0,numsLen):
            num = nums[i]
            if m.get(num) == None:
                m[num] = i
            else:
                d = i-m[num]
                if d <= k :
                    return True
                m[num] = i

        return False

if __name__ == "__main__":
    s = Solution()
    print(s.containsNearbyDuplicate([1,2,3,1], 3))
    print(s.containsNearbyDuplicate([1,0,1,1], 1))
    print(s.containsNearbyDuplicate([1,2,3,1,2,3], 2))