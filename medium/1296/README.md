## 1296. Divide Array in Sets of K Consecutive Numbers

source: [https://leetcode-cn.com/problems/divide-array-in-sets-of-k-consecutive-numbers](https://leetcode-cn.com/problems/divide-array-in-sets-of-k-consecutive-numbers)

### Describe

Given an array of integers nums and a positive integer k, check whether it is possible to divide this array into sets of k consecutive numbers.

Return true if it is possible. Otherwise, return false.

### Example

```
Input: nums = [1,2,3,3,4,4,5,6], k = 4
Output: true
Explanation: Array can be divided into [1,2,3,4] and [3,4,5,6].
```
```
Input: nums = [3,2,1,2,3,4,3,4,5,9,10,11], k = 3
Output: true
Explanation: Array can be divided into [1,2,3] , [2,3,4] , [3,4,5] and [9,10,11].
```
```
Input: nums = [1,2,3,4], k = 3
Output: false
Explanation: Each array should be divided in subarrays of size 3.
```

### Constraints

- 1 <= k <= nums.length <= 10^5
- 1 <= nums[i] <= 10^9