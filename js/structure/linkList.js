/**
 * @file linkList 
 */
'use strict';

function Node(value){
  this.value = value;
  this.next = null;
}

function LinkList() {
  this.head = new Node('head');
  this.find = find;
  this.insert = insert;
  this.remove = remove;
  this.print_r = print_r;
}

function find(item) {
  let currNode = this.head;
  while (currNode !== null && currNode.value !== item) {
    currNode = currNode.next;
  }
  console.log(currNode);
  return currNode;
}

function insert(value,item='') {
  let newNode = new Node(value);
  let current;
  if(item){
    current = this.find(item);
    console.log(item,'value:',current.value)
    newNode.next = current.next;
    current.next = newNode
  }else {
    current = this.head;
    while (current.next !== null) {
      current = current.next;
    }
    current.next = newNode
  }
}

function remove(item){
  let currNode = this.head;
  while(currNode.next !== null && currNode.next.value !== item){
    currNode = currNode.next;
  }
  if(currNode.next !== null ){
    let rmNode = currNode.next;
    currNode.next = rmNode.next;
    rmNode.next = null;
  }
}

function print_r(){
  let currNode = this.head;
  while(currNode !== null){
    console.log(currNode.value);
    currNode = currNode.next;
  }
}

//Test code
const cities = new LinkList();
cities.insert('Conway');
cities.insert('Russellville');
cities.insert('Alma');
cities.insert('testNode','Russellville');
cities.print_r();
console.log('-------------');
// cities.remove('Alma');
cities.find('Russellville');

