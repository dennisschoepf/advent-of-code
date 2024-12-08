(ns aoc-2024.d3
	(:import (java.lang Integer))
	(:require [clojure.string :as str]))

(def sample-input
	"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")

(defn get-input
	[]
	(slurp "src/d3.txt"))

(defn parse-and-multiply-items
	[coll]
	(->> coll
			 (map Integer/parseInt)
			 (apply *)))

(defn part1
	[text]
	(->> text
			 (re-seq #"mul\((\d+),(\d+)\)")
			 (map rest)
			 (map parse-and-multiply-items)
			 (apply +)))
