/**
 * @param {ListNode} head
 * @return {boolean}
 */
var hasCycle = function(head) {
  if (!head) return false
  let walker = head
  let runner = head
  while (walker.next && (runner.next && runner.next.next)) {
    if (walker.next == runner.next.next) return true
    walker = walker.next
    runner = runner.next.next
  }
  return false
};