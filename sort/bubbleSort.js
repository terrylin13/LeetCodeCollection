'use strict';
const BaseArray = require('./ArrayList.class')
class ArrayList extends BaseArray{

  constructor() {
    super()
  }

  bubbleSort() {
    let arrLength = this.array.length;
    for (let i = 0; i < arrLength; i++){
      for (let j = i + 1; j < arrLength; j++){
        if (this.array[i] > this.array[j]) this.swap(i,j)
      }
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
arr.bubbleSort();
console.log(arr.toString());