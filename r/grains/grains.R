square <- function(n) {
    if(n >= 1 && n <= 64) {
        2 ^ (n - 1)
    } else {
        stop(sprintf("Argument n must be in range 1..64, but %d given", n))
    }
}

total <- function() {
    sum(unlist(Map(square, 1:64)))
}
