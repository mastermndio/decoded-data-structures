# We are implementing a Merge Sort algorithm in Python
def merge_sort(list):
    #Handle a list of a single element
    if len(list) < 2:
        return list

    #So that we can split this array, we need to find the halfway point
    halfway_idx = len(list) // 2

    #Break array in half and save
    left = merge_sort(list[:halfway_idx])
    right= merge_sort(list[halfway_idx:])
 
    #Return left and right sides after merged
    return merge(left, right)

#Create a function to actually do the merging
def merge(left, right):
    #Setup a new array
    sortedList = []

    #setup counters for both loops
    i, j = 0, 0

    
    while (len(sortedList) < len(left) + len(right)):
        if left[i] < right[j]:
            sortedList.append(left[i])
            i += 1
        else:
            sortedList.append(right[j])
            j += 1
        
        #Checking if we are at the end of either the left or the right. If we are at the end, we no longer need to compare. We simply appent the remaining items on the other side into the sorted array
        if i == len(left) or j == len(right):
            sortedList.extend(left[i:] or right[j:])

    return sortedList



A = [100, 45, 2, 56, 2342, 64, 74, 5, 65]

print(merge_sort(A))

