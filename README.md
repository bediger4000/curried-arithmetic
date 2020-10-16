# Daily Coding Problem: Problem #682 [Medium]

This problem was asked by Squarespace.

Write a function, add_subtract,
which alternately adds and subtracts curried arguments.
Here are some sample operations:

```
add_subtract(7) -> 7

add_subtract(1)(2)(3) -> 1 + 2 - 3 -> 0

add_subtract(-5)(10)(3)(9) -> -5 + 10 - 3 + 9 -> 11
```

## Build and run

```sh
$ go build add_subtract.go
$ ./add_subtract  -5 10 3 9
(+) 11
$
```

`add_subtract` reads numbers sequentially from the command line
and runs `func add_subtract` on them internally.

```go
curried := add_subtract(n)
for _, str := range os.Args[2:] {
    n, _ := strconv.Atoi(str)
    curried = curried(n)
}
curried()  // print calculated number
```

The code ends up calling the final curried function with no
arguments to trigger it to print an answer.
This is a little outside the scope of the problem statement.
Strictly interpreting that statement
doesn't give you a way to output an answer.

## Analysis

This
