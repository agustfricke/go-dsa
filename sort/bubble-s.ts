function bubble_sort(arr: number[]): number[] {

    for (let i = 0; i < arr.length; i++) {
        for (let j = 0; j < arr.length - 1 - i; j++) {
            if (arr[j] > arr[j + 1]) {
                const tmp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = tmp;
                console.log(tmp, arr[j]) 
                console.log(arr)
            }
        }
    }
    return arr;
}

const arr = [9, 3, 7, 4, 69, 420, 42];
console.log(arr)
console.log(bubble_sort(arr));

