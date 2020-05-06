<?php
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class Solution {

    /**
     * @param TreeNode $root
     * @return Boolean
     */
    function isValidBST($root) {
        return $this->helper($root,null,null);
    }

    function helper($node,$lower,$upper){
        if($node == null) return true;
        $val = $node->val;

        if ($lower !== null && $val <= $lower) return false;
        if ($upper !== null && $val >= $upper) return false;

        if (!$this->helper($node->left, $lower, $val)) return false;
        if (!$this->helper($node->right, $val, $upper)) return false;

        return true;
    }
}