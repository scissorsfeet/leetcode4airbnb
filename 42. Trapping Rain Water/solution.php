<?php

class Solution {

    /**
     * @param Integer[] $height
     * @return Integer
     */
    function trap($height) {
        $l = 0;
        $r = count($height)-1;
        $res = 0;
        while($l < $r) {
            $mx = min($height[$l], $height[$r]);

            while($l < $r && $height[$l] <= $mx) {
                $res += $mx - $height[$l++];
            }
            while($l < $r && $height[$r] <= $mx) {
                $res += $mx - $height[$r--];
            }
        }

        return $res;
    }
}

$s = new Solution();
var_dump($s->trap(array(0,1,0,2,1,0,1,3,2,1,2,1)));
var_dump($s->trap(array(2,1)));
var_dump($s->trap(array(2,0,1)));