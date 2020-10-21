#Bubble Sort compares adjacent elements and ensures that they are sorted appropriately

nums = [21, 6, 13, 1, 105, 74, 37]

for i in range(len(nums) - 1):
    isSorted = True
    for j in range(len(nums) - 1):
        if nums[j] > nums[j + 1]:
            nums[j], nums[j + 1] = nums[j + 1], nums[j]
            isSorted = False

    if isSorted == True:
        print(nums)
        break

        
        
#            tempSwap = nums[j]
#            nums[j] = nums[j + 1]
#            nums[j + 1] = tempSwap
#
