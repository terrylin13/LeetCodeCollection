<?php

use PHPUnit\Framework\TestCase;
require(__DIR__.'/../structure/stack.php');

class StackTest extends TestCase
{
    public function testInsert()
    {
        $stack = new Stack();       
        $this->assertTrue($stack->push(1));
        $this->assertTrue($stack->push(2));
        $this->assertTrue($stack->push(3));
        $this->assertTrue($stack->push(4));
        return $stack;
    }

    /**
     * @depends testInsert
     */
    public function testPop($stack)
    {
        $this->assertEquals($stack->pop(),4);
    }
}