# this is a stub function that takes a natural_number
# and should return the difference-of-squares as described
# in the README.md
difference_of_squares <- function(natural_number) {
    square_of_sum = Reduce('+', 1:natural_number, 0) ^ 2
    sum_of_squares = Reduce('+', Map(function(x) x^2, 1:natural_number), 0)

    square_of_sum - sum_of_squares
}
