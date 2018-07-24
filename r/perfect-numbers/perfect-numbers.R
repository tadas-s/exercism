number_type <- function(n) {
    if (n <= 0) {
        stop(sprintf("n must be > 0, but %d was given", n))
    }

    aliquot_sum = sum(
        Filter(function(x) { n %% x == 0 }, 1:(n/2))
    )

    if (n == 1 || aliquot_sum < n) {
        return("deficient")
    } else if (aliquot_sum == n) {
        return("perfect")
    } else if (aliquot_sum > n) {
        return("abundant")
    }
}
