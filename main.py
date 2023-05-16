import time

start = time.time()

for i in range(1, 10000001):
    print(i)

elapsed = time.time() - start
print("Tiempo transcurrido:", elapsed, "segundos")
