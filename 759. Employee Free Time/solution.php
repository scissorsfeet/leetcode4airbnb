<?php
/**
 * Created by PhpStorm.
 * User: yangyueqian
 * Date: 2019-03-27
 * Time: 16:02
 */

class Solution {
    public function employeeFreeTime($schedule) {
        $intervals = array();
        foreach($schedule as $arr) {
            foreach($arr as $item) {
                $intervals[] = $item;
            }
        }
        usort($intervals, array($this, "sortByStart"));
        $n = count($intervals);
        $t = $intervals[0];
        $freeTime = array();
        for($i=1;$i<$n;$i++) {
            if($intervals[$i][0] > $t[1]) {
                $freeTime[] = array($t[1], $intervals[$i][0]);
                $t = $intervals[$i];
            } else {
                $t = array($t[0], max($t[1], $intervals[$i][1]));
            }
        }
        return $freeTime;
    }

    private function sortByStart($a, $b) {
        if($a[0] < $b[1]) {
            return -1;
        }
        return 1;
    }
}

$s = new Solution();
//print_r($s->employeeFreeTime(array(array(array(1,2),array(5,6)),array(array(1,3)),array(array(4,10)))));
print_r($s->employeeFreeTime(array(array(array(1,3),array(6,7)), array(array(2,4)), array(array(2,5), array(9,12)))));