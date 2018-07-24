# This is a stub function to take two strings
# and calculate the hamming distance
hamming <- function(strand1, strand2) {
    if (nchar(strand1) != nchar(strand2)) {
        stop("Lengths of both strands must match")
    }

    sum(strsplit(strand1, "")[[1]] != strsplit(strand2, "")[[1]])
}
