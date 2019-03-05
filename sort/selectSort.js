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
  this.selectSort = () => {
    let len = array.length,indexMin;
    for (let i = 0; i < len-1; i++)
    {
      indexMin = i;
      for (let j = i + 1; j < len; j++)
      {
        if (array[indexMin] > array[j])
        {
          indexMin = j;
        }
      }
      if (i !== indexMin) {
        let temp = array[i];
        array[i] = array[indexMin];
        array[indexMin] = temp;
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
arr.selectSort();
console.log(arr.toString());