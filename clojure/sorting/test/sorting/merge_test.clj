(ns sorting.merge-test
  (:require [clojure.test :refer :all]
            [sorting.merge :refer :all]))

(deftest merge-test
  (testing "Merge sort"
           (is (= (range 10) (merge-sort(shuffle (range 10)))))))