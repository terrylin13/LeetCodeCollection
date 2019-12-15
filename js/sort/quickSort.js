'use strict';
const BaseArray = require('./ArrayList.class')
class ArrayList extends BaseArray { 

  quickSort() {
    return this.quick(this.array,)
  }
  quick (array, left, right){

    var index; //{1}
  
    if (array.length > 1) { //{2}
  
      index = this.partition(array, left, right); //{3}
  
      if (left < index - 1) {                //{4}
        quick(array, left, index - 1);     //{5}
      }
  
      if (index < right) {  //{6}
        quick(array, index, right);        //{7}
      }
    }
  }

  partition (array, left, right) {
　
    var pivot = this.array[Math.floor((right + left) / 2)], //{8}
      i = left,                                      //{9}
      j = right;                                     //{10}
      while (i <= j) {                //{11}
      while (array[i] < pivot) {  //{12}
        i++;
      }
      while (array[j] > pivot) {  //{13}
        j--;
      }
      if (i <= j) { //{14}
        this.swap(i, j); //{15}
        i++;
        j--;
      }
    }
    return i; //{16}
  }
}



function ArrayList()
{
  let array = [];
  this.insert = (item) => {
    array.push(item);
  }
  this.toString = () => {
    return array.join();
  }
  this.quickSort =  (array, left, right) => {

    var index; //{1}
  
    if (array.length > 1) { //{2}
  
      index = partition(array, left, right); //{3}
  
      if (left < index - 1) {                //{4}
        quickSort(array, left, index - 1);     //{5}
      }
  
      if (index < right) {  //{6}
        quickSort(array, index, right);        //{7}
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
arr.insertionSort();
console.log(arr.toString());