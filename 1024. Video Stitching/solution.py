class Solution(object):
    def videoStitching(self, clips, T):
        """
        :type clips: List[List[int]]
        :type T: int
        :rtype: int
        """

        clipsLen = len(clips)
        if clipsLen == 0:
            return -1
        clips = sorted(clips, key=lambda x: (x[0],x[1]))
        if clips[0][0] != 0:
            return -1
        stack = []
        for i in range(0,clipsLen):
            start = clips[i][0]
            end = clips[i][1]
            if start >= end:
                continue
            if len(stack) == 0:
                stack.append([start, end])
            else:
                top = stack[len(stack)-1]
                if start > top[1]:
                    return -1
                if end > top[1]:
                    if start <= top[0]:
                        stack.pop()
                        stack.append([top[0],end])
                    elif start <= top[1]:
                        stack.append([top[1], end])
            if stack[len(stack)-1][1] >= T:
                return len(stack)

        return -1


if __name__ == '__main__':
    clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]]
    s = Solution()
    print(s.videoStitching(clips,10))

    clips = [[0,1],[1,2]]
    s = Solution()
    print(s.videoStitching(clips,5))

    clips = [[5,7],[1,8],[0,0],[2,3],[4,5],[0,6],[5,10],[7,10]]
    s = Solution()
    print(s.videoStitching(clips,5))

    clips = [[17,18],[25,26],[16,21],[3,3],[19,23],[1,5],[0,2],[9,20],[5,17],[8,10]]
    s = Solution()
    print(s.videoStitching(clips,15))
