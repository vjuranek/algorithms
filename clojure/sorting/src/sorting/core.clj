(ns sorting.core
  (:gen-class))

(defn swap [v i j]
  (assoc v i (nth v j) j (nth v i)))

(defn bubble-sort [v]
  (loop [i 1
         j 0
         vect v]
    (if (< i (count v))
      (let [vect2 (if (> (nth vect j) (nth vect (inc j)))
                    (swap vect j (inc j))
                    vect)]
        (if (< (inc j) (- (count vect2) i))
          (recur i (inc j) vect2)
          (recur (inc i) 0 vect2)))
      vect)))

(defn -main
  [& args]
  (def v (bubble-sort [3 2 5 4 1]))
  (println v))
