function bs_list(haystack: number[], needle: number): boolean {

    let low = 0
    let hi = haystack.length

    do {
        const mid = Math.floor(low + (hi - low) / 2);
        const value = haystack[mid]
        console.log(value)

        if (value === needle) {
            return true
        } else if(value > needle) {
            hi = mid;
        } else {
            low = mid + 1;
        }

    } while (low < hi); 
        return false
}

const orderedList: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
console.log(bs_list(orderedList, 2))
