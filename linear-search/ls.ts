// Big O : O(n)

function linear_search(haystack: number[], needle: number): boolean {
    for (let i = 0; i < haystack.length; i++) {
        if (haystack[i] === needle) {
            return true;
        }
    }
    return false;
}

const arr: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 0];
const q: number = 6;

console.log(linear_search(arr, q));
