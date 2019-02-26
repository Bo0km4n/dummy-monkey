let fizzbuzz = fn(n) {
    for (let i=1;i<n+1;++i) {
        switch {
            case i % 3 == 0 && i % 5 == 0:
                puts("FizzBuzz");
                break;
            case i % 3 == 0:
                puts("Fizz");
                break;
            case i % 5 == 0:
                puts("Buzz");
                break;
        }
    }
};

fizzbuzz(15);
