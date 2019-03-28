<?php
class Solution {

    /**
     * @param String[][] $board
     * @param String[] $words
     * @return String[]
     */
    function findWords($board, $words) {
        $rowLen = count($board);
        if ($rowLen == 0) {
            return array();
        }
        $wordsLen = count($words);
        $colLen = count($board[0]);
        if ($wordsLen == 0 || $colLen == 0) {
            return array();
        }

        $t = new Trie();
        foreach($words as $s) {
            $t->insert($s);
        }

        $res = array();
        $visited = array();
        for($i=0;$i<$rowLen;$i++) {
            for($j=0;$j<$colLen;$j++) {
                if(isset($t->root->childs[$board[$i][$j]])) {
                    $this->search($board, $i, $j, $t->root->childs[$board[$i][$j]], $visited, $res, $rowLen, $colLen);
                }
            }
        }

        return $res;
    }

    private function search($board, $i, $j, $t, &$visited, &$res, $rowLen, $colLen) {
        if($t->value != "") {
            $res[] = $t->value;
            $t->value = "";
        }
        $dirs = array(array(0,-1),array(0,1),array(-1,0),array(1,0));
        $visited[$i][$j] = 1;
        foreach($dirs as $dir) {
            $x = $i+$dir[0];
            $y = $j+$dir[1];
            if($x<0||$x>=$rowLen||$y<0||$y>=$colLen||!isset($t->childs[$board[$x][$y]])||$visited[$x][$y]) {
                continue;
            }
            $this->search($board, $x, $y, $t->childs[$board[$x][$y]], $visited, $res, $rowLen, $colLen);
        }
        $visited[$i][$j] = 0;
    }
}

class TrieNode {
    public $childs = array();
    public $value = "";
}

class Trie {
    public $root;
    public function __construct() {
        $this->root = new TrieNode();
    }

    public function insert($s) {
        $p = $this->root;
        $sLen = strlen($s);
        for($i=0;$i<$sLen;$i++) {
            $b = $s[$i];
            if(!isset($p->childs[$b])) {
                $p->childs[$b] = new TrieNode();
            }
            $p = $p->childs[$b];
        }
        $p->value = $s;
    }
}

$s = new Solution();
$res=$s->findWords(
    array(
      array('o','a','a','n'),
      array('e','t','a','e'),
      array('i','h','k','r'),
      array('i','f','l','v'),
    ),
    array("oath","pea","eat","rain")
);
var_dump($res);