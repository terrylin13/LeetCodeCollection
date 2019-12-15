'use strict';
const BaseArray = require('./ArrayList.class')
class ArrayList extends BaseArray { 
  mergaSort() {
    return this._mergeSortRec(this.array)
  }
  
  _mergeSortRec(array) {
    let arrLength = array.length;
    if (arrLength === 1) return array;
    let mid = Math.floor(arrLength / 2),
      left  = array.slice(0, mid),
      right = array.slice(mid, arrLength);
    return this.array = this.merga(this._mergeSortRec(left), this._mergeSortRec(right));
  }

  merga(left, right){
    let result = [], il = 0, ir = 0;
    let leftLength = left.length, rightLength = right.length;

    while (il < leftLength && ir < rightLength) {
      if (left[il] < right[ir]) {
        result.push(left[il++]);
      } else {
        result.push(right[ir++]);
      }
    }

    while (il < leftLength) {
      result.push(left[il++]);
    }

    while (ir < rightLength) {
      result.push(right[ir++]);
    }

    return result;
  }

}

//TestCode
let arr = new ArrayList();
arr.insert(12);
arr.insert(2);
arr.insert(7);
arr.insert(71);
arr.insert(87);
console.log(arr.toString());
arr.mergaSort();
console.log(arr.toString());