(ns run-length-encoding)

(defn- join
  [a b]
  (let [[count-a char-a] (take-last 2 a)
        [count-b char-b] b]
    (if (= char-a char-b)
      (concat (drop-last 2 a) (list (+ count-a count-b) char-a))
      (concat a b))))

; TODO: rewrite using: (partition-by identity "aaaaarrrrggghhhhhpsh") 
(defn run-length-encode
  "encodes a string with run-length-encoding"
  [plain-text]
  (if (= plain-text "")
    plain-text
    (->> (seq plain-text)
         (map #(list 1 %1))
         (reduce join)
         (remove #(and (int? %1) (= 1 %1)))
         (apply str))))

(defn- digit-character? [c] (contains? (set (seq "0123456789")) c))

(defn run-length-decode
  "decodes a run-length-encoded string"
  [cipher-text]
  (->> (seq cipher-text)
       (partition-by #(if (digit-character? %1) true (identity %1)))
       (map #(if (digit-character? (first %1)) (Integer/parseInt (apply str %1)) %1))
       (partition 2 1 (list nil))
       (map #(if (int? (first %1)) (repeat (- (first %1) 1) (last %1)) (first %1)))
       (flatten)
       (apply str)))
