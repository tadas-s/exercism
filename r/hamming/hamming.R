# This is a stub function to take two strings
# and calculate the hamming distance
hamming <- function(strand1, strand2) {
    strand1_characters = unlist(strsplit(strand1, ""))
    strand2_characters = unlist(strsplit(strand2, ""))

    if (length(strand1_characters) != length(strand2_characters)) {
        stop("Lengths of both strands must match")
    }

    if (length(strand1_characters) == 0) {
        return(0)
    }

    distance = 0

    for(index in 1:length(strand1_characters)) {
        if(strand1_characters[index] != strand2_characters[index]) {
            distance = distance + 1
        }
    }

    distance
}
