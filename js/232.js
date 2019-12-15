/**
 * Initialize your data structure here.
 */
var MyQueue = function() {
    this.inStack = []
    this.outStack = []
    this.count = 0;
};

/**
 * Push element x to the back of queue. 
 * @param {number} x
 * @return {void}
 */
MyQueue.prototype.push = function(x) {
    this.inStack.push(x)
    this.count++
};

/**
 * Removes the element from in front of queue and returns that element.
 * @return {number}
 */
MyQueue.prototype.pop = function() {
    if (this.outStack.length === 0 )
        while(this.inStack.length > 0)
            this.outStack.push(this.inStack.pop())
    this.count--
    return this.outStack.pop()
};

/**
 * Get the front element.
 */
MyQueue.prototype.peek = function() {
    if (this.outStack.length === 0 )
        while(this.inStack.length > 0)
            this.outStack.push(this.inStack.pop())
    return this.outStack[this.count - 1]
};

/**
 * Returns whether the queue is empty.
 * @return {boolean}
 */
MyQueue.prototype.empty = function() {
    return this.count === 0 ? true : false;
};

/** 
 * Your MyQueue object will be instantiated and called as such:
 * var obj = new MyQueue()
 * obj.push(x)
 * var param_2 = obj.pop()
 * var param_3 = obj.peek()
 * var param_4 = obj.empty()
 */