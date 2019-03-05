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
  this.insertionSort = () => {
    let len = array.length,j,temp;
    for (let i = 0; i < len; i++)
    {
      j = i;
      temp = array[i];
      while (j > 0 && array[j - 1] > temp)
      {
        array[j] = array[j-1];        //{6}
        j--;
      }
      array[j] = temp;
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