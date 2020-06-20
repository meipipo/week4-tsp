## Go TSP solver

### Usage
Make sure you are in this directory.
```
$ go build
```
makes `my-solver` binary file. Then run:
```
$ ./my-solver
<input number: 0~6>
<option>
```
It outputs the length and updates `solution_yours_<input number: 0~6>.csv`.

### Option
- `greedy`: Greedy solver. (It works only input 0 for now.)
- To be updated...
- (default: Always returns 0.)

### Result
| Solver | Challenge 0 | Challenge 1 | Challenge 2 | Challenge 3 | Challenge 4 | Challenge 5 | Challenge 6 |
|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|
| `greedy` |3418.10|-|-|-|-|-|-|
| - |-|-|-|-|-|-|-|