(ns reverse-string)

(defn reverse-string [s]
  (reduce #(str %2 %1) "" s))
