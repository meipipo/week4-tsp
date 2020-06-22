## Go TSP solver

### Usage
Make sure you are in this directory.
```
$ go build
```
makes `my-solver` binary file. Then run:
```
$ ./my-solver -input <input number: 0~6> - <solver name>
```
It outputs the length and updates `solution_yours_<input number: 0~6>.csv`.

### Option
- `nn`: Simple solver. (slow ...)
- To be updated...
- (default: Always returns 0.)

### Result
| Solver | Challenge 0 | Challenge 1 | Challenge 2 | Challenge 3 | Challenge 4 | Challenge 5 | Challenge 6 |
|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|
| `nn` |3418.10|3832.29|5065.58|9276.22|12084.32|298497.41|47822.41|
| - |-|-|-|-|-|-|-|