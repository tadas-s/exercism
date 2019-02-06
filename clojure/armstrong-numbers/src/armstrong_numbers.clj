(ns armstrong-numbers)

(defn- pow [base exponent]
  (reduce *
          (repeat exponent base)))

(defn- digits [num]
  (+ 1 (int (Math/floor (Math/log10 num)))))

(defn- dividers [num]
  (take-while #(> num (rem num %))
              (iterate #(* 10 %) 1)))

(defn armstrong? [num]
  (let [total_digits (digits num)]
    (->> (dividers num)
         (map #(quot (rem num (* 10 %)) %))
         (map #(pow % total_digits))
         (reduce +)
         (= num))))
