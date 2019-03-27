<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target) {
        $m = array();
        foreach($nums as $i=>$num) {
            $m[$num] = $i;
        }

        $res = array();
        foreach($nums as $i=>$num) {
            $find = $target-$num;
            if(isset($m[$find]) && $i!=$m[$find]) {
                $res = array($i, $m[$find]);
                break;
            }
        }

        return $res;
    }
}

$s = new Solution();
//print_r($s->twoSum(array(2,7,11,15), 9));
print_r($s->twoSum(array(3,2,4), 6));