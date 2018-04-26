(ns sorting.bubble
  (:require [sorting.commons :as c]))

(defn bubble-sort [v]
  (loop [i 1
         j 0
         vect v]
    (if (< i (count v))
      (let [vect2 (if (> (nth vect j) (nth vect (inc j)))
                    (c/swap vect j (inc j))
                    vect)]
        (if (< (inc j) (- (count vect2) i))
          (recur i (inc j) vect2)
          (recur (inc i) 0 vect2)))
      vect)))