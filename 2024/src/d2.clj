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

(defn pairs-pred?
	[pred coll]
	(let [first (first coll)
				second (second coll)
				rest (rest coll)]
		(cond
			(nil? first) false
			(empty? rest) true
			(and (some? second) (not (pred first second))) false
			:else (pairs-pred? pred rest))))
;; Update after research on other solutions:
;; (partition 2 1 coll) could have been used. It returns every possible pair
;; in the collection (e.g. for (1 2 3 4) => ((1 2) (2 3) (3 4))) and therefore
;; allows to just calculate each difference. After that it would have been enough
;; to check every (with (every?)) difference to be in range as well as all positive
;; or all negative. This should also be more performant that the recursion I came
;; up with

(defn safe?
	[coll]
	(or (pairs-pred? increase-within? coll) (pairs-pred? decrease-within? coll)))

(defn part1
	[input]
	(count (filter safe? (process-input input))))

;; Part 2 - Dampened differences
(defn possible-combinations
	[coll]
	(let [indexes (range (count coll))]
		(map (fn [index]
					 (keep-indexed #(when (not (= %1 index)) %2) coll)) indexes)))

(defn some-combination-safe?
	[coll]
	(let [some-is-safe (some safe? (possible-combinations coll))]
		(if (= true some-is-safe) true false)))

(defn part2
	[input]
	(count (filter some-combination-safe? (process-input input))))
