/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
const containsNearbyDuplicate = (nums, k) => {
    for (let i = 0; i < nums.length -1; i +=1) {
        for(let j = i + 1; j < nums.length && j - i <= k; j +=1) {
            if (nums[i] === nums[j]) {
                return true;
            }
        }
    }
    return false;
};
