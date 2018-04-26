(ns sorting.commons)

(defn swap [v i j]
  (assoc v i (nth v j) j (nth v i)))