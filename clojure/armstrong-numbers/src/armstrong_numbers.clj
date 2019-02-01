(ns armstrong-numbers)

(defn digits [number] (map #(long (Character/digit % 10)) (str number)))

;; Integers only base^exponent operation
(defn pow [base exponent]
  (reduce * (repeat exponent base)))

(defn armstrong? [num]
  (let [all_digits (digits num) number_of_digits (count all_digits)]
    (= num (reduce + (map #(pow % number_of_digits) all_digits)))))
