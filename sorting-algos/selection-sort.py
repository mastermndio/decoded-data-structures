#Iterate through entire Array, selecting the lowest value and replacing that value with the current iteration value
nums = [21, 6, 13, 1, 105, 74, 37]

for i in range(len(nums) -1):
    low_val_idx = i
    for j in range(i, len(nums) -1):
        if nums[low_val_idx] > nums[j + 1]:
            low_val_idx = j+1

    nums[i], nums[low_val_idx] = nums[low_val_idx], nums[i]

print(nums)
