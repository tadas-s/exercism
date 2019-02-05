(ns armstrong-numbers)

(defn- pow [base exponent]
  (reduce * 
          (repeat exponent base)))

(defn- digits [num]
  (+ 1 (int (Math/floor (Math/log10 num)))))

(defn- dividers [num]
  (take-while #(> num (mod num %)) 
              (iterate #(* 10 %) 1)))

(defn armstrong? [num]
  (let [total_digits (digits num)]
    (= num
       (reduce +
               (map #(pow % total_digits)
                    (map #(quot (mod num (* 10 %)) %)
                         (dividers num)))))))
