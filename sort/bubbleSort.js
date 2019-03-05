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
  this.bubbleSort = () => {
    let len = array.length;
    for (let i = 0; i < len; i++)
    {
      for (let j = i + 1; j < len; j++)
      {
        if (array[i] > array[j])
        {
          let temp = array[j];
          array[j] = array[i];
          array[i] = temp;
        }
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