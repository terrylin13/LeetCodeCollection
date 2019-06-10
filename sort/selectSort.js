'use strict';
const BaseArray = require('./ArrayList.class')
class ArrayList extends BaseArray {
  constructor() {
    super();
  }
  selectSort() {
    let arrLength = this.array.length, indexMin;
    for (let i = 0; i < arrLength - 1; i++){
      indexMin = i;
      for (let j = i + 1; j < arrLength; j++){
        if (this.array[indexMin] > this.array[j]) indexMin = j;
      }
      if (i != indexMin) this.swap(i, indexMin);
    }
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
arr.selectSort();
console.log(arr.toString());