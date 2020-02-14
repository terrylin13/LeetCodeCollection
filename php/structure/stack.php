<?php
class Stack 
{
    protected $_stack;

    public function __construct()
    {
        $this->_stack = new SplStack();
    }

    public function push($data)
    {
        return $this->_stack->push($data);
    }

    public function pop()
    {
        return $this->_stack->pop();
    }

    public function __invoke()
    {
        return $this->_stack;
    }
}