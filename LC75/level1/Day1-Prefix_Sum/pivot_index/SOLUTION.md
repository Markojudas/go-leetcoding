<!-- markdownlint-disable -->

# Solution

We need the sum of the values to the left and the right of every index.

Let's say we knew `S` as the sum of the numbers, and we are at index `i`. If we knew the sum of numbers `leftsum` that are to the left of index `i`, then the other sum to the right of the index would be `S - nums[i] - leftsum`.

As we iterate through candidate indexes `i`, we will maintain the correct value of `leftsum`.
