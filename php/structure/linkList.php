<?php
class Node 
{
    public $value;
    public $next;
    public function __construct($value,$next = null)
    {
        $this->value = $value;
        $this->next  = $next;
    }
}

class LinkList 
{
    private $_link;

    public function __construct()
    {
        $this->_link = new Node('head');
    }

    public function insert($value,$item = null)
    {
        $node = new Node($value);
        $currNode = null;
        if($item){
            $currNode = $this->find($item);
            if(!$currNode) return false;
            if($currNode->next != null) $node->next = $currNode->next;
        } else {
            $currNode = $this->_link;
            while($currNode && $currNode->next){
                $currNode = $currNode->next;
            }
        }
        $currNode->next = $node;
        return true;
    }

    public function find($value)
    {
        $currNode = $this->_link;
        while($currNode && $currNode->value != $value)
        {
            $currNode = $currNode->next;
        }
        return $currNode ?? false;
    }

    public function display()
    {
        $currNode = $this->_link;
        $returnString = '';
        while($currNode){
            $returnString.=$currNode->value.'->';
            $currNode = $currNode->next;
        }
        $returnString.='tail';
        return $returnString;
    }
}