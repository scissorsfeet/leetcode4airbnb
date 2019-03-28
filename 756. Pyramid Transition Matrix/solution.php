<?php
class Solution {

    /**
     * @param String $bottom
     * @param String[] $allowed
     * @return Boolean
     */
    function pyramidTransition($bottom, $allowed) {
        $m = array();
        foreach($allowed as $v) {
            $m[substr($v, 0, 2)][] = $v[2];
        }

        return $this->helper($bottom, "", $m);
    }

    private function helper($cur, $above, $m) {
        $curLen = strlen($cur);
        $aboveLen = strlen($above);
        if($curLen == 2 && $aboveLen == 1) {
            return true;
        }
        if($curLen == $aboveLen+1) {
            return $this->helper($above, "", $m);
        }

        $pos = $aboveLen;
        $base = substr($cur, $pos, 2);
        if(isset($m[$base])) {
            foreach($m[$base] as $v) {
                if($this->helper($cur, $above.$v, $m)) {
                    return true;
                }
            }
        }

        return false;
    }
}

$s = new Solution();
var_dump($s->pyramidTransition("XYZ", array("XYD", "YZE", "DEA", "FFF")));