<?php
/**
 * Created by PhpStorm.
 * User: leo
 * Date: 2019-03-27
 * Time: 12:09
 */

class Solution {

    /**
     * @param String[] $words
     * @return Integer[][]
     */
    public function palindromePairs($words) {
        $res = array();
        $wordMap = array();
        $wordsLen = count($words);
        $lenArr = array();
        for($i=0;$i<$wordsLen;$i++) {
            $wordMap[$words[$i]] = $i;
            $lenArr[strlen($words[$i])] = 1;
        }
        $lenArr = array_keys($lenArr);
        sort($lenArr);

        for($i=0;$i<$wordsLen;$i++) {
            $revWord = strrev($words[$i]);
            if(isset($wordMap[$revWord]) && $wordMap[$revWord] != $i) {
                $res[] = array($i, $wordMap[$revWord]);
            }
            $wordLen = strlen($revWord);
            $index = array_search($wordLen, $lenArr);
            for($j=0;$j<$index;$j++) {
                $d = $lenArr[$j];
                $sub = substr($revWord, $wordLen-$d);
                if($this->isPalindrome($revWord, 0,$wordLen-$d-1) && isset($wordMap[$sub])) {
                    $res[] = array($i, $wordMap[$sub]);
                }
                $sub = substr($revWord, 0, $d);
                if($this->isPalindrome($revWord, $d, $wordLen-1) && isset($wordMap[$sub])) {
                    $res[] = array($wordMap[$sub], $i);
                }
            }
        }

        return $res;
    }

    private function isPalindrome($arr, $left, $right) {
        $res = true;
        while($left < $right) {
            if($arr[$left++] != $arr[$right--]) {
                $res = false;
                break;
            }
        }
        return $res;
    }
}

$s = new Solution();
//print_r($s->palindromePairs(array("abcdd","cba")));
print_r($s->palindromePairs(array("abcd","dcba","lls","s","sssll")));
//print_r($s->palindromePairs(array("lls","s")));