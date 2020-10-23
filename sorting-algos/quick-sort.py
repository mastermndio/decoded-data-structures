def start_sort(list):
    quick_sort(list, 0, len(list) - 1)

def quick_sort(list, l, r):
    #check if left and right are the same because if they are, we likely have a sorted list
    if l >= r:
        return

    p = partition(list, l, r)

    #Call quicksort on beginning of array up to pivot
    quick_sort(list, l, p - 1)

    #call quicksort again on array from pivot to end
    quick_sort(list, p + 1, r)

def partition(list, l, r):
    #Grab last item in list to use as point of reference
    pivot = list[r]
    i = l - 1

    for j in range(l, r):
        if list[j] < pivot:
            i += 1
            list[i], list[j] = list[j], list[i]
        
    #After partitioning, put pivot in the right place
    list[i + 1], list[r] = list[r], list[i + 1]
        
    print(list)
    return i + 1

ourList = [34,2,56,4,8,17,389,8,3,0]

start_sort(ourList)
