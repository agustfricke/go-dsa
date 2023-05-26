import time

start = time.time()

def linear_search(haystack, needle):
    for i in range(len(haystack)):
        print(i)
        if haystack[i] == needle:
            return "Found"
    return "Not found"

arr = list(range(1, 100001))
q = 23010

print(linear_search(arr, q))

elapsed = time.time() - start
print("Tiempo transcurrido:", elapsed, "segundos")


