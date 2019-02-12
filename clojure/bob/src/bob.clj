(ns bob)

(use '[clojure.string :only (upper-case trim ends-with?)])

(defn- silence? [s]
  (= (trim s) ""))

(defn- asking? [s]
  (ends-with? s "?"))

(defn- shouting? [s]
  (and (= s (upper-case s))
       (some? (re-find (re-matcher #"[a-zA-Z]+" s)))))

(defn response-for [s]
  (cond
    (silence? s) "Fine. Be that way!"
    (and (asking? s) (shouting? s)) "Calm down, I know what I'm doing!"
    (asking? s) "Sure."
    (shouting? s) "Whoa, chill out!"
    :else "Whatever."))
