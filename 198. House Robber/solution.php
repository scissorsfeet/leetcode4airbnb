<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function rob($nums) {
        $n = count($nums);
        if($n==0) {
            return 0;
        } else if($n==1) {
            return $nums[0];
        }

        $dp = array($nums[0], max($nums[0], $nums[1]));
        for($i=2;$i<$n;$i++) {
            $cur = max($dp[0]+$nums[$i], $dp[1]);
            $dp[0] = $dp[1];
            $dp[1] = $cur;
        }

        return $dp[1];
    }
}

$s = new Solution();
var_dump($s->rob(array(1,2,3,1)));