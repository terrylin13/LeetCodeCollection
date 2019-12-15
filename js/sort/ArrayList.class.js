'use strict';
class ArrayList {

  constructor() {
    this.array = [];
  }

  insert(item) {
    this.array.push(item);
  }

  swap(index1,index2) {
    let temp = this.array[index1];
    this.array[index1] = this.array[index2];
    this.array[index2] = temp;
  }
  toString() {
    return this.array.join(',')
  }
}

module.exports = ArrayList;