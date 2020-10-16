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
$ cd $GOPATH/src
$ git clone https://github.com/bediger4000/curried-arithmetic.git
$ cd curried-arithmetic
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
This is outside the scope of the problem statement.
Strictly interpreting that statement
doesn't give you a way to output an answer.

## Analysis

This is an unusually tricky problem for a "medium" rating.
Curried functions are a staple of functional programming,
but are very rare outside that regime.

I did this in Go with two, similar inter-related functions:

```go
func subtract_add(n ...int) fn {
    return func(a ...int) fn {
        if len(a) == 0 {
            fmt.Printf("(+) %d\n", n[0])
            return nil
        }
        return add_subtract(n[0] - a[0])
    }
}

func add_subtract(n ...int) fn {
    return func(a ...int) fn {
        if len(a) == 0 {
            fmt.Printf("(-) %d\n", n[0])
            return nil
        }
        return subtract_add(n[0] + a[0])
    }
}
```

They differ only in calling each other,
and the operation invoked on their integer arguments.

The type `fn` is important: `type fn func(...int) fn`

That is, objects of type `fn` are functions that can be invoked
with zero or more integer arguments, and return an object of type `fn`.

I figured this out by a lot of experimentation.
I am not a functional programming by nature.

My thoughts and programming process went something like this:

1. The function `add_subtract` takes an integer argument.
2. The example call, `add_subtract(1)(2)(3)` implies that
the function `add_subtract` has to return a function,
further a function that takes an integer argument.
3. The function returned by `add_subtract` has to return
a function taking an integer as an argument.
4. Since the example doesn't specify a limiting number
of chained invocations,
the function from (3) has to return a similar function,
with an integer argument, and a function return.
5. Each function of type `fn` has to return a function of type `fn`.
6. Something like the mechanism in a [self-replicating program](https://github.com/bediger4000/Self-replicating-go)
has to exist for a function of type `fn` to return another function of type `fn`.
7. That's too hard,
what if I have two functions of type `fn` that return functions
of type `fn`, one doing subtraction, one doing addition,
and each one creates an anonymous function that invokes the other.
8. I'll just tinker around with these two functions until I get it right.

I wrote a number of throwaway experimental programs to try various
implementations of type `fn`.

There's no good way to get a curried function to print out a final answer. 
I did this by having the anonymous functions check argument count,
and if invoked with no formal arguments, print the final arithmetic value calculated.

As an interview question,
this seems really hard.
The interviewer should not expect much if whiteboarding.
This problem is better-suited for a take-home problem.
Even then, it's not clear to me what you'd get back.

Based on the solutions I found around the web,
the interviewer should be prepared for just about anything.
I don't think this is a very good interview question for a number
of reasons, including the obscurity of currying,
and the difficulty of designing recursive types.
These difficulties lead to the rainbow of solutions,
most of which work around the recursive type.

Once again, Daily Coding Problem gets the rating wrong,
reinforcing my belief that a job candidate cannot feel bad about
getting skunked by a programming problem posed by an interviewer.
There's so much variation in the difficulty of the problems versus
the ratings, that a candidate will regularly get baffled.

## Other Solutions

I found a [Go version](https://github.com/vaskoz/dailycodingproblem-go/tree/master/day363)
that uses a struct with a method,
seperating that object + method into the add_subtract "function".

There's a [Medium article](https://medium.com/@Charles_Stover/variable-length-currying-in-javascript-7f7bb7bdad8b)
discussing the problem.
The author has the same problem with no good way to show you
what's been calculated, much of the article exists to show a
weird JavaScript way to get an answer.

The Medium article references a [Reddit post](https://www.reddit.com/r/javascript/comments/9vxdkx/was_asked_this_js_interview_question_is_this_even/)
that is supposedly the public source of this problem.
The reddit thread has a number of... unusual... solutions.
