# this is a stub function that takes a natural_number
# and should return the difference-of-squares as described
# in the README.md
difference_of_squares <- function(natural_number) {
    sequence <- 0:natural_number
    sum(sequence) ^ 2 - sum(sequence ^ 2)
}
