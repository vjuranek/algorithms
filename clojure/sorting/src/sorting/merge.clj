(ns sorting.merge
  (:require [sorting.commons :as c]))

(defn do-merge [v w]
  (let [cv (count v)
        cw (count w)]
    (loop [i 0
           j 0
           m []]
      (if (and (< i cv) (< j cw))
        (if (< (nth v i) (nth w j))
          (recur (inc i) j (conj m (nth v i)))
          (recur i (inc j) (conj m (nth w j))))
        (if (= i cv)
          (concat m (subvec w j))
          (concat m (subvec v i)))))))

(defn do-merge-cond [v w]
  (let [cv (count v)
        cw (count w)]
    (loop [i 0
           j 0
           m []]
      (cond
        (= i cv) (concat m (subvec w j))
        (= j cw) (concat m (subvec v i))
        :else (if (< (nth v i) (nth w j))
                (recur (inc i) j (conj m (nth v i)))
                (recur i (inc j) (conj m (nth w j))))))))

(defn do-merge-destruct [v w m]
    (let [[v1 & vn] v
          [w1 & wn] w]
      (cond
        (nil? v1) (concat m w)
        (nil? w1) (concat m v)
        (< v1 w1) (recur vn w (conj m v1))
        :else (recur v wn (conj m w1)))))

(defn merge-sort [v]
  (if (<= (count v) 1)
    v
    (let [ half (int (/(count v) 2))
           split (split-at half v)
           v1 (merge-sort (first split))
           v2 (merge-sort (second split))]
      ;(do-merge-cond (into [] v1) (into [] v2))
      (do-merge-destruct (into [] v1) (into [] v2) []))))