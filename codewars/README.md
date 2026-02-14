# Codewars Katas (Go)

This folder contains small Go solutions to various Codewars katas. Each kata lives in its own package directory with its own tests.

## Run tests

Run all tests across all kata packages:

```sh
make test
```

Or directly via Go:

```sh
go test ./...
```

## Katas

| Folder                                                       | Package                    | Function                                          | What it does                                                 |
| ------------------------------------------------------------ | -------------------------- | ------------------------------------------------- | ------------------------------------------------------------ |
| [bool_to_word](./bool_to_word)                               | `booltoword`               | `BoolToWord(word bool) string`                    | Converts a boolean into `"Yes"` or `"No"`.                   |
| [count_the_monkeys](./count_the_monkeys)                     | `countthemonkeys`          | `monkeyCount(n int) []int`                        | Returns the sequence `1..n` (empty for `n <= 0`).            |
| [multiply](./multiply)                                       | `multiply`                 | `Multiply(a, b int) int`                          | Multiplies two integers.                                     |
| [opposite_number](./opposite_number)                         | `oppositenumber`           | `Opposite(number int) int`                        | Returns the additive inverse of an integer.                  |
| [remove_first_last_character](./remove_first_last_character) | `removefirstlastcharacter` | `RemoveChar(word string) string`                  | Removes the first and last character (empty if length <= 2). |
| [return_negative](./return_negative)                         | `returnnegative`           | `MakeNegative(x int) int`                         | Ensures a number is negative (or zero).                      |
| [reversed_strings](./reversed_strings)                       | `reversedstrings`          | `Solution(word string) string`                    | Reverses a string (Unicode-safe via runes).                  |
| [square_sum](./square_sum)                                   | `squaresum`                | `SquareSum(numbers []int) int`                    | Returns the sum of the squares of all numbers in a slice.    |
| [string_repeat](./string_repeat)                             | `stringrepeat`             | `RepeatStr(repetitions int, value string) string` | Repeats a string `repetitions` times.                        |
| [smallest_integer_finder](./smallest_integer_finder)         | `smallestintegerfinder`    | `SmallestIntegerFinder(numbers []int) int`        | Returns the smallest integer in a slice (0 if empty).        |
| [sum_of_positive](./sum_of_positive)                         | `sumofpositive`            | `PositiveSum(numbers []int) int`                  | Sums only the positive numbers in a slice.                   |

## Notes

- Some katas use the standard Go `testing` package; `count_the_monkeys` uses Ginkgo/Gomega (already included in the module dependencies).
