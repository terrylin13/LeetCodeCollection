/**
 *  @file greedyAlgorithm implement with JavaScript
 */
function coinChange(amount,coins){
  let change = [],
      total = 0;
  coins = coins.sort((a,b)=>{
    return a - b;
  })
  let len = coins.length
  for(let i = len; i>=0; i--){
    let coin = coins[i];
    while( total + coin <= amount ){
      change.push(coin);
      total += coin;
    }
  }
  console.log(change)
  return change;
}

coinChange(38,[1,5,20,10])
