<!-- markdownlint-disable -->

# Binary Search

Given an array of integers `nums` which is sorted in ascending order, and an integer ` target``, write a function to search  `target`in`nums`. if `target`exists, then return its index. Otherwise, return`-1```.

You must write an algorithm with `O(log n)` runtime complexity.
<br>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>          nums = [-1,0,3,5,9,12], target = 9
<strong>Output:</strong>         4
<strong>Explanation:</strong>    9 exists in nums and its index is 4</code></pre>
<br>

**Example 2:**

<pre><code><strong>Input:</strong>          nums = [-1,0,3,5,9,12], target = 2
<strong>Output:</strong>         -1
<strong>Explanation:</strong>    2 does not exist in nums so return -1/code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>1 <= nums.length <= 10<sup>4</sup></code></li>
    <li><code>-10<sup>4</sup> < nums[i], target < 10<sup>4</sup></code></li>
    <li>All the integers in <code>nums</code> are <strong>unique</strong>.</li>
    <li><code>nums</code> is sorted in ascending order.</li>
</ul>
