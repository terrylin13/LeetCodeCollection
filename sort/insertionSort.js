'use strict';
const BaseArray = require('./ArrayList.class')
class ArrayList extends BaseArray { 
  insertionSort() {
    let arrLength = this.array.length;
    for (let i = 0; i < arrLength; i++) {
      let temp = this.array[i], j = i;
      while (j > 0 && this.array[j - 1] > temp) {
        this.array[j] = this.array[j - 1]
        j--;
      }
      this.array[j] = temp;
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
arr.insertionSort();
console.log(arr.toString());