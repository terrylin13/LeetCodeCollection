<?php

use PHPUnit\Framework\TestCase;
require(__DIR__.'/../structure/linkList.php');

class LinkListTest extends TestCase
{   
    public function testInsert()
    {
        $link = new LinkList();       
        $this->assertTrue($link->insert(2));
        $this->assertTrue($link->insert(5));
        return $link;
    }

    /**
     * @depends testInsert
     */
    public function testInsertByFindItem($link)
    {
        $this->assertTrue($link->insert(3,2));
        return $link;
    }

    /**
     * @depends testInsertByFindItem
     */
    public function testDisplay($link)
    {
        $string = $link->display();
        $this->assertStringStartsWith('head', $string);
        echo "\n","testDisplay: ",$string,"\n";
        
    }

}