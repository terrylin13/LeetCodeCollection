
function quick (array, left, right){

  var index; //{1}

  if (array.length > 1) { //{2}

    index = partition(array, left, right); //{3}

    if (left < index - 1) {                //{4}
      quick(array, left, index - 1);     //{5}
    }

    if (index < right) {  //{6}
      quick(array, index, right);        //{7}
    }
  }
};

'use strict';

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