# Go-Katas

Multiple algorithms and data-structure practice exercises from multiple sources.

> Each file is a **standalone kata program** with `//go:build ignore` and is meant to be run individually, for example:
>
> ```bash
> go run 01-linear-tier-1/006-search.go
> ```

## Index

### Tier 1 — Linear Fundamentals

| # | Name | TC |
|---|------|----|
| 001 | [Linear Table](01-linear-tier-1/001-table.go) | O(n) |
| 002 | [Find Max / Min Value](01-linear-tier-1/002-max-min.go) | O(n) |
| 003 | [Find Max Value by Swapping](01-linear-tier-1/003-max-swap.go) | O(n) |
| 004 | [Find Min Value by Swapping](01-linear-tier-1/004-min-swap.go) | O(n) |
| 005 | [Find Max / Min Value (Generics)](01-linear-tier-1/005-max-min-gen.go) | O(n) |
| 006 | [Linear Search](01-linear-tier-1/006-search.go) | O(n) |
| 007 | [Linear Search (Generics)](01-linear-tier-1/007-search-gen.go) | O(n) |
| 008 | [Count Occurrences](01-linear-tier-1/008-count.go) | O(n) |
| 009 | [Count Occurrences (Generics)](01-linear-tier-1/009-count-gen.go) | O(n) |
| 010 | [Count Evens and Odds](01-linear-tier-1/010-count-even-odd.go) | O(n) |
| 011 | [Sum](01-linear-tier-1/011-sum.go) | O(n) |
| 012 | [Append](01-linear-tier-1/012-append.go) | O(n) |
| 013 | [Insert At Index](01-linear-tier-1/013-insert.go) | O(n) |
| 014 | [Delete At Index](01-linear-tier-1/014-delete.go) | O(n) |
| 015 | [Is Sorted](01-linear-tier-1/015-is-sorted.go) | O(n) |
| 016 | [Reverse In Place](01-linear-tier-1/016-reverse.go) | O(n) |
| 100 | [DS: Array (capstone)](01-linear-tier-1/100-ds-array-ops.go) | O(n) per op |

### Tier 2 — Write-Pointer / Compaction

| # | Name | TC |
|---|------|----|
| 001 | [Collect Non-Zeros Into New Array](02-linear-tier-2/001-collect-non-zeros-new-array.go) | O(n) |
| 002 | [Compact In Place Without Fill](02-linear-tier-2/002-compact-in-place-no-fill.go) | O(n) |
| 003 | [Compact In Place With Fill](02-linear-tier-2/003-compact-in-place-with-fill.go) | O(n) |
| 004 | [Move Zeros (Swap Write Pointer)](02-linear-tier-2/004-move-zeros-swap-write-pointer.go) | O(n) |
| 005 | [Move Zeros (Two-Index Window)](02-linear-tier-2/005-move-zeros-two-index-window.go) | O(n) |
| 006 | [Move Zeros Left (Mirror)](02-linear-tier-2/006-move-zeros-left-mirror.go) | O(n) |
| 007 | [Move Any Target To End](02-linear-tier-2/007-move-target-to-end.go) | O(n) |
| 008 | [Linear Filter Pattern](02-linear-tier-2/008-filters.go) | O(n) |
| 009 | [Remove Duplicates (Sorted Invariant / Dictionary)](02-linear-tier-2/009-remove-dups-invariant.go) | O(n) |