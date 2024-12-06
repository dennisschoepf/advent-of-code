(ns aoc-2024.d1
	(:import (java.lang Integer)))

(def sample-input
	"3   4
4   3
2   5
1   3
3   9
3   3")

(defn get-input
	[]
	(slurp "src/d1.txt"))

(defn keep-odd
	[coll]
	(keep-indexed #(when (odd? %1) %2) coll))

(defn keep-even
	[coll]
	(keep-indexed #(when (even? %1) %2) coll))

(defn process-into-list
	[text]
	(let [numbers (->> text
										 (#(str/split % #"\s+"))
										 (map Integer/parseInt))]
		(list (sort (keep-even numbers)) (sort (keep-odd numbers)))))

(defn map-similarity
	[ids freqs]
	(map
	 (fn [id]
		 (* id (get freqs id 0)))
	 ids))

(defn part1
	[input]
	(->> input
			 process-into-list
			 (#(map - (first %) (second %)))
			 (map abs)
			 (apply +)))

(defn part2
	[input]
	(let [left (first (process-into-list input))
				right (second (process-into-list input))
				freqs (frequencies right)
				similarities (map-similarity left freqs)]
		(reduce + similarities)))
