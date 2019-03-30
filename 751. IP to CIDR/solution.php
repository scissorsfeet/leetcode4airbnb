<?php

class Solution {
    function ip2cidr($ip, $n) {
        $x = 0;
        $set = explode(".", $ip);
        for($i=0;$i<4;$i++) {
            $x = $x*256 + intval($set[$i]);
        }

        $res = array();
        while($n>0) {
            $step = $x & -$x;
            while($step > $n) {
                $step = $step >> 1;
            }

            $tmp = array();
            $tmp[] = $x>>24&255;
            $tmp[] = $x>>16&255;
            $tmp[] = $x>>8&255;
            $tmp[] = $x&255;
            $num = 32-log($step, 2);
            $res[] = implode(".", $tmp)."/".$num;

            $x += $step;
            $n -= $step;
        }

        return $res;
    }
}

$s = new Solution();
print_r($s->ip2cidr("255.0.0.7", 10));