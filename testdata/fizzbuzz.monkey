let fizzbuzz = fn(n) {
    for (let i=0;i<n;i = i + 1) {
        if i % 3 == 0 {
            puts("fizz")
        } else if (i % 5 == 0) {
            puts(buzz)
        }
    }
};

fizzbuzz(100);