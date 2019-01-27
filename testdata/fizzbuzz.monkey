let fizzbuzz = fn(n) {
    for (let i=0;i<n;++i) {
        switch {
            case i % 3 == 0:
                puts("fizz");
                break;
            case i % 5 == 0:
                puts("buzz")
                break;
        }
    }
};

fizzbuzz(10);