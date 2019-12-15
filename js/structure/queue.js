/**
 * @file Queue 
 */
'use strict';
function Queue() {
  let items = [];

  this.enqueue = (element) => {
    items.push(element);
  };

  this.dequeue = () => {
    return items.shift();
  };

  this.front = () => {
    return items[0];
  };

  this.isEmpty = () => {
    return items.length == 0;
  };

  this.size = () => {
    return items.length;
  }

  this.print = () => {
    console.log(items.toString());
  }

}

//Test Code
let queue = new Queue();
console.log(queue.isEmpty());
queue.enqueue('test');
queue.enqueue('Terry')
queue.enqueue(1);
console.log(queue.size());
console.log(queue.dequeue());
queue.print();