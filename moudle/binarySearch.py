class Solution:
    def __init__(self, nums, target):
        self.nums = nums
        self.target = target

    def left_mid_right(self):
        left = 0
        right = len(self.nums) - 1
        mid = (left + right) // 2
        return left, mid, right
    
    '''基础的二分查找
    >>>nums = [1, 2, 2, 2, 3]
    >>>target = 2
    ans = 2
    '''
    def binary_search(self):
        left, mid, right = self.left_mid_right()
        while left <= right:
            if self.nums[mid] > self.target:
                right = mid - 1
            elif self.nums[mid] < self.target:
                left = mid + 1
            elif self.nums[mid] == self.target:
                return mid
            mid = (left + right) // 2
        return -1

    '''左边界的二分查找
        >>>nums = [1, 2, 2, 2, 3]
        >>>target = 2
        ans = 1
        '''
    def left_search(self):
        left, mid, right = self.left_mid_right()
        while left <= right:
            if self.nums[mid] >= self.target:
                right = mid - 1
            elif self.nums[mid] < self.target:
                left = mid + 1
            mid = (left + right) // 2
        if left == len(self.nums):
            return -1
        return left if self.nums[left] == self.target else -1

    '''右边界的二分查找
        >>>nums = [1, 2, 2, 2, 3]
        >>>target = 2
        ans = 3
        '''
    def right_search(self):
        left, mid, right = self.left_mid_right()
        while left <= right:
            if self.nums[mid] > self.target:
                right = mid - 1
            elif self.nums[mid] <= self.target:
                left = mid + 1
            mid = (left + right) // 2
        if right < 0:
            return -1
        return right if self.nums[right] == self.target else -1

    def searchRange(self):
        b = self.binary_search()
        l = self.left_search()
        r = self.right_search()
        return b, l, r


nums = list(map(int, input().split()))
target = int(input())
# a = [1, 2, 2, 2, 3]
ans = Solution(nums=nums, target=target)

print(ans.searchRange())
