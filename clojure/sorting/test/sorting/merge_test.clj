(ns sorting.merge-test
  (:require [clojure.test :refer :all]
            [sorting.merge :refer :all]))

(deftest merge-test
  (testing "Merge sort"
           (is (= (range 10) (merge-sort(shuffle (range 10)))))))

(deftest merge-cond-test
  (testing "Merge sort using cond"
           (is (= (range 10) (merge-sort-cond(shuffle (range 10)))))))

(deftest merge-destruct-test
  (testing "Merge sort argument destruction"
           (is (= (range 10) (merge-sort-desctruct(shuffle (range 10)))))))