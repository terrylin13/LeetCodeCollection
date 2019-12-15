/**
 * @file Reverse Linked List
 * @link https://leetcode.com/problems/reverse-linked-list/
 */
/**
 * Brute force
 * @param {ListNode} head
 * @return {ListNode}
 */
var reverseList = function(head) {
  var result = null;
  var stack = [];
  
  var current = head;
  //when head is null
  if(!current) return head

  while (current) {
      stack.push(current);
      current = current.next;
  }
  
  // Set head to end of list.
  result = stack.pop() || [];
  current = result;
  
  while (current) {
      current.next = stack.pop();
      current = current.next;
  }

  return result;
};


/**
 * Common practice
 * @param {ListNode} head
 * @return {ListNode}
 */
var reverseList = function(head) {
  var pre = null;
  var next = null;
  var current = head;
  //when head is null
  if(!current) return head
  while(current){
    next = current.next
    current.next = pre
    pre = current
    current = next
  }
  return  pre
}