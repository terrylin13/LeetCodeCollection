#!/usr/bin/env php
<?php
class ListNode
{
    public $val = 0;
    public $next = null;
    function __construct($val)
    {
        $this->val = $val;
    }
}

class Solution
{

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2)
    {
        $next1 = $l1;
        $next2 = $l2;
        $returnList = new ListNode(0);
        $tempNext = &$returnList;
        $carry = 0;
        while ($next1 != null || $next2 != null) {
            $val = $next1->val + $next2->val + $carry;
            if ($tempNext == null) {
                $tempNext = new ListNode($val);
            }
            $tempNext->val = $val;
            if ($carry > 0) {
                $carry = 0;
            }
            if ($val >= 10) {
                $carry = (int) ($val / 10);
                $tempNext->val = (int) ($val  % 10);
            }
            $next1 = $next1->next;
            $next2 = $next2->next;
            $tempNext = &$tempNext->next;
        }
        if ($carry > 0) {
            $tempNext = new ListNode($carry);
        }
        return $returnList;
    }

    public function printList(ListNode $list)
    {
        while ($list) {
            echo $list->val, "->";
            $list = $list->next;
        }
        echo "null\n";
    }
}
