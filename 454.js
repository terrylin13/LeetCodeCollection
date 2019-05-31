/**
 * @param {number[]} A
 * @param {number[]} B
 * @param {number[]} C
 * @param {number[]} D
 * @return {number}
 */
var fourSumCount = function(A, B, C, D) {
  let res = 0, cache = new Map();
 
     for (let i = 0; i < A.length; i++) {
         for (let j = 0; j < B.length; j++) {
             let missing = (A[i] + B[j]) * -1;
             let occur = cache.get(missing) || 0;
             
             cache.set(missing, ++occur);
         }
     }
     
     for (let k = 0; k < C.length; k++) {
         for (let l = 0; l < D.length; l++) {
             let num = C[k] + D[l];
             if (cache.has(num)) {
                 res += cache.get(num);
             }
         }
     }
     
     return res;
 };