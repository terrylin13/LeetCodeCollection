/**
  * @param {number[]} nums
  * @return {number[]}
  * 功能：合并 k 个链表
  * 边界条件：
  * 1）判断数组是否为空
  * 2）判断数组长度为 1 时
  * 3）判断数组长度为 2 时
  * 4）判断数组长度大于 2 时
  */
 var mergeKLists = function(lists) {
  // 当 lists 中有一个链表时
  if(lists.length == 0){
      return null;
  }else if(lists.length == 1){
      // 判断数组长度为 1 时
      return lists[0];
  }else if(lists.length == 2){
      // 判断数组长度为 2 时
      return mergeTwoLists(lists[0],lists[1]);
  }else{
      // 判断数组长度大于 2 时
      // 取数组的中部坐标
      let middle = Math.floor(lists.length/2);
      // 取左右两边数组
      let leftList = lists.slice(0,middle);
      let rightList = lists.slice(middle);
      // 递归、分割、合并
      return mergeTwoLists(mergeKLists(leftList),mergeKLists(rightList));
  }       
};
//两个链表合并
var mergeTwoLists = function(l1, l2) {
  let result = null;

  //终止条件
  if(l1 == null) return l2;
  if(l2 == null) return l1;

  //判断数值大小递归
  if(l1.val < l2.val){
      result = l1;
      result.next = mergeTwoLists(l1.next,l2);
  }else{
      result = l2;
      result.next = mergeTwoLists(l2.next,l1);
  }
  
  //返回结果
  return result;
};   