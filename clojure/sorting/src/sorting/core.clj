(ns sorting.core
  (:gen-class)
  (:require [sorting.bubble :as b]
            [sorting.merge :as m]))

(defn -main
  [& args]
  ;;(def v (b/bubble-sort [3 2 5 4 1]))
  (def v (time (m/merge-sort (shuffle (range 10)))))
  (println v))

