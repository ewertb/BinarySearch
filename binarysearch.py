def get_middle(start, end):
    return int(start + (end - start) / 2)

def search(numbers, find):
    numbers = sorted(numbers)
    start = 0
    end = len(numbers) - 1
    mid = get_middle(start, end)
    found = False

    while found == False:
        if end < start:
            break

        if numbers[mid] < find:
            start = mid + 1
            mid = get_middle(start, end)

        if numbers[mid] > find:
            end = mid - 1
            mid = get_middle(start, end)

        if numbers[mid] == find:
            found = True

    return found

numbers = [1,3,9,7,5,4,6,2,8]
find = 7
print(search(numbers, find))