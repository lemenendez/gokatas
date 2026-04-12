# Strings Basics Study Plan

Strings are arrays with attitude.

## Goals

- Understand matching vs equality vs search
- Know when naive is enough

## Concepts

- Exact match vs substring search
- Prefix and suffix vs contains
- Naive search
- Rabin-Karp (practical)
- Optional: KMP (for understanding, not daily use)

## Implement

- `Index(s, sub)`
- `HasPrefix`
- `HasSuffix`
- `Contains`
- `FindAllOccurrences`

## Problems

- Implement `strStr()` (classic)
- Repeated substring pattern
- Minimum window substring (advanced)

## Tier 1.5 - Strings Basic Katas

These are pure string fundamentals. Do them in order.

1. String length (manual loop, rune-aware)
2. Count vowels and consonants
3. Count spaces and words
4. Reverse a string
5. Check palindrome
6. Case conversion (ASCII manual)
7. Remove extra spaces (trim + collapse internal spaces)
8. Count character frequency
9. First non-repeating character
10. Most frequent character
11. Check anagram (sort-based, then map-based)
12. Check if one string is rotation of another
13. Remove all occurrences of a character
14. Replace all occurrences of a substring (naive)
15. Longest word in a sentence
16. Longest common prefix (array of strings)
17. Run-length encoding basics (`aaabb -> a3b2`)
18. Validate parentheses (basic stack intro)
