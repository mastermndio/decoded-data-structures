#Iterate through entire Array,starting at second item, comparing index with all items before it to determine where to insert it
nums = [21, 6, 13, 1, 105, 74, 37]

for i in range(1, len(nums)):
    current = nums[i]

    j = i - 1

    while j >= 0 and current < nums[j]:
        nums[j + 1] = nums[j]

        j -= 1

    nums[j+1] = current

print(nums)





