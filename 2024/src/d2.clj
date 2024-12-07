(ns aoc-2024.d2
	(:require [clojure.string :as str])
	(:import (java.lang Integer)))

(def sample-input
	"7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9")

(defn get-input
	[]
	(slurp "src/d2.txt"))

(defn process-input
	[input]
	(->> input
			 str/split-lines
			 (map #(str/split % #" "))
			 (map #(map Integer/parseInt %))))

(defn decrease-within?
	[val compare-val]
	(let [diff (- val compare-val)]
		(and (> diff 0) (< diff 4))))

(defn increase-within?
	[val compare-val]
	(let [diff (- val compare-val)]
		(and (< diff 0) (> diff -4))))

(defn increasing-slowly?
	[coll]
	(let [first (first coll)
				second (second coll)
				rest (rest coll)]
		(cond
			(nil? first) false
			(empty? rest) true
			(and (some? second) (not (increase-within? first second))) false
			:else (increasing-slowly? rest))))

(defn decreasing-slowly?
	[coll]
	(let [first (first coll)
				second (second coll)
				rest (rest coll)]
		(cond
			(nil? first) false
			(and (some? second) (not (decrease-within? first second))) false
			(empty? rest) true
			:else (decreasing-slowly? rest))))

(defn safe?
	[coll]
	(or (decreasing-slowly? coll) (increasing-slowly? coll)))

(defn part1
	[input]
	(count (filter safe? (process-input input))))

(defn part2
	[input])
