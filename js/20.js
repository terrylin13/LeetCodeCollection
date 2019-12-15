/**
 * @file Valid Parentheses
 * @link https://leetcode.com/problems/valid-parentheses/
 */

/**
 * @param {string} 
 * @return {boolean}
 */
var isValid = function (s) {
    let len = s.length;
    let first = s[0];
    
    if (len == 0) return true;
    if ([')','}',']'].indexOf(first) !== -1) return false;
    
    let stack = [];
    stack.push(first);

    const VALID_MAP = {
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for (let i = 1; i< len; i++) {
        let length = stack.length;
        let top = length > 0 ? stack[length - 1] : null;
        let now = s[i];
        if (VALID_MAP[now] && VALID_MAP[now] === top) stack.pop();
        else if (VALID_MAP[now]) return false;
        else stack.push(now)
    }
    return stack.length === 0;
};
  