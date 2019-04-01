class Solution(object):
    def combinationSum2(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        def dfs(candidates, target, start, path, res):
            if target == 0:
                return res.append(path+[])

            for i in range(start, len(candidates)):
                if target - candidates[i] >= 0:
                    if i > start and candidates[i] == candidates[i-1]:
                        continue
                    path.append(candidates[i])
                    dfs(candidates, target-candidates[i],i+1,path,res)
                    path.pop()

        res = []
        candidates.sort()
        dfs(candidates,target,0,[],res)
        return res

if __name__=="__main__":
    s = Solution()
    # print(s.combinationSum2([2, 3, 6, 7], 7))
    # print(s.combinationSum2([1, 98], 99))
    print(s.combinationSum2([10,1,2,7,6,1,5], 8))