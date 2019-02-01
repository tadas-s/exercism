(ns armstrong-numbers)

(defn digits [number] (map #(long (Character/digit % 10)) (str number)))

(defn armstrong? [num]
  (let [all_digits (digits num)]
      ( = num (long (reduce + (map #(Math/pow % (count all_digits)) all_digits))))
    )
  )
