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
class Solution
{

    /**
     * @param TreeNode $root
     * @param Integer $sum
     * @return Boolean
     */
    function hasPathSum($root, $sum)
    {
        if ($root == null) {
            return false;
        }
        $queue = [];
        array_push($queue, $root);
        while (!empty($queue)) {
            $node = array_shift($queue);
            if (!$node->left && !$node->right) {
                if ($node->val == $sum) return true;
            }
            if ($node->left) {
                $node->left->val += $node->val;
                array_push($queue, $node->left);
            }
            if ($node->right) {
                $node->right->val += $node->val;
                array_push($queue, $node->right);
            }
        }
        return false;
    }
}
