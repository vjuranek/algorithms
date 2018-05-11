(ns sorting.bubble-test
  (:require [clojure.test :refer :all]
            [sorting.bubble :refer :all]))

(deftest bubble-test
  (testing "Bubble sort"
           (is (= (range 10) (bubble-sort(shuffle (range 10)))))))