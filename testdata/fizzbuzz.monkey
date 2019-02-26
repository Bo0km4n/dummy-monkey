let fizzbuzz = fn(n) {
    for (let i=0;i<n;++i) {
        switch {
            case i % 3 == 0 && case i % 5 == 0:
                puts("FizzBuzz")
            case i % 3 == 0:
                puts("Fizz");
                break;
            case i % 5 == 0:
                puts("Buzz");
                break;
        }
    }
};

fizzbuzz(10);
