(ns rna-transcription)

(defn- nucleotide-translate [nucleotide]
  (case nucleotide
    \G \C
    \C \G
    \T \A
    \A \U
    :else (throw (AssertionError. "Wrong input."))))

(defn to-rna [dna]
  (->> dna
       (map nucleotide-translate)
       (apply str)))
