## Go TSP solver

Goで書いて、辛くなったので途中でPythonで書き直そうとしてなぜかうまくいかなかったのでGoで再開しました。。

### Usage
Make sure you are in this directory.
```
$ go build
```
makes `my-solver` binary file. Then run:
```
$ ./my-solver -input <input number: 0~6> -solver <solver name>
```
It outputs the length and updates `solution_yours_<input number: 0~6>.csv`.

### Option
- `nn`: Simple solver. (slow ...)
- `2opt`: Swap two edges + nn.
- To be updated...
- (default: Always returns 0.)

### Result
```
$ python solution_verifier.py
```
| Solver | Challenge 0 | Challenge 1 | Challenge 2 | Challenge 3 | Challenge 4 | Challenge 5 | Challenge 6 |
|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|
| `nn` |3418.10|3832.29|5065.58|9276.22|12084.32|24191.66|47822.41|
| `2opt` |3418.10|3832.29|5030.00|8890.56|11382.41|24263.20|47655.98|
| - |-|-|-|-|-|-|-|

- TODO: nnで毎回minを検索しているので遅い（ソートして前から見るべき）
- なぜChallenge5で2optっぽい物を適用したのに距離が伸びてしまったのかは今のところ謎です…