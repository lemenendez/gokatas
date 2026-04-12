# Linear Arrays and Basics Study Plan

Use this guide to practice array katas from fundamentals to challenging pattern-based problems.

## Tier 1 - Foundational

If these are not automatic, do not move on.

Exercises:

- Find max and min
- Linear search (return index or `-1`)
- Count occurrences of a value
- Check if array is sorted
- Reverse array (new array first, then in-place)
- Sum of elements
- Count evens and odds

Goal:

- No hesitation
- No bugs
- No thinking "how do I start?"

This is like learning to walk before sprinting.

## Tier 2 - Core

This is where growth happens.

These are the best exercises for this topic. They force you to combine
scanning, indexing, and state.

1. Second Largest Element

Find the second largest number without sorting.

Why it is good:

- Forces you to track multiple states (`max`, `secondMax`)
- Teaches careful updates

2. Move Zeros to End (in-place)

`[0,1,0,3,12] -> [1,3,12,0,0]`

Why it is good:

- Introduces in-place transformation
- Prepares you for two pointers later

3. Remove Duplicates (Sorted Array)

`[1,1,2,2,3] -> [1,2,3]`

Why it is good:

- Teaches early slow/fast pointer thinking
- Forces index control

4. Best Time to Buy and Sell Stock (Easy)

`[7,1,5,3,6,4] -> max profit = 5`

Why it is good:

- Introduces running minimum
- A major pattern used in later problems

5. Find First Repeating Element

`[2,5,1,2,3,5] -> 2`

Try:

- Brute force: `O(n^2)`
- Optimized with map: `O(n)`

Why it is good:

- Teaches time vs memory trade-offs

6. Longest Increasing Contiguous Subarray

`[1,2,2,3,4,1] -> longest = 3` (`[2,3,4]`)

Why it is good:

- Teaches streak tracking
- Introduces "reset vs continue"

7. Prefix Sum Array

`[1,2,3,4] -> [1,3,6,10]`

Why it is good:

- Builds a base for advanced problems
- Teaches accumulation thinking

8. Find Missing Number (0 to n)

`[3,0,1] -> missing = 2`

Why it is good:

- Introduces math + array reasoning
- Multiple valid solutions exist

## Tier 3 - Challenging

Insight builders.

These are still array problems, but they require more abstraction and pattern
recognition.

1. Kadane's Algorithm (Maximum Subarray)

`[-2,1,-3,4,-1,2,1,-5,4] -> 6`

Why it is good:

- Teaches optimal substructure
- One of the most important algorithmic patterns

2. Majority Element (`> n/2`)

`[2,2,1,1,2,2] -> 2`

Try:

- Brute force
- Hash map
- Boyer-Moore voting

3. Product of Array Except Self

`[1,2,3,4] -> [24,12,8,6]`

Why it is good:

- Forces prefix and suffix thinking
- No division makes it a deeper challenge

4. Rotate Array (in-place)

`[1,2,3,4,5,6,7], k=3 -> [5,6,7,1,2,3,4]`

Why it is good:

- Strengthens index math
- Builds in-place manipulation mastery

5. Trapping Rain Water (optional stretch)

Why it is good:

- Teaches left-max and right-max reasoning
- Prepares for two-pointer techniques

## How To Approach Each Problem

Do not just solve. Extract knowledge.

For each problem:

1. Brute force first

Even if it is `O(n^2)`.

2. Ask:

- Can I do this in one pass?
- What state do I need?

3. Identify the pattern:

- Running max/min
- Prefix/suffix
- Streak tracking
- Neighbor comparisons

4. Rewrite clean

No hacks. Clean logic.
