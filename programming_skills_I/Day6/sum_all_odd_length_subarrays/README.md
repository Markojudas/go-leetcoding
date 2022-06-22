<!-- markdownlint-disable -->

# Sum of All Odd Length Subarrays

Given an array of positive integers `arr`, return <em>the sum of all possible</em>**odd-length subarrays** of `arr`.

A **subarray** is a contiguous subsequence of the array.
<br>

**Example 1:**

<pre><code><strong>Input:</strong>          arr = [1,4,2,5,3]
<strong>Output:</strong>         58
<strong>Explanation:</strong>    The odd-length subarrays of arr and their sums are:
[1] = 1
[4] = 4
[2] = 2
[5] = 5
[3] = 3
[1,4,2] = 7
[4,2,5] = 11
[2,5,3] = 10
[1,4,2,5,3] = 15
If we add all these together we get 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58</code></pre>

**Example 2:**

<pre><code><strong>Input:</strong>          arr = [1,2]
<strong>Output:</strong>         3
<strong>Explanation:</strong>    There are only 2 subarrays of odd length, [1] and [2]. Their sum is 3.</code></pre>

**Example 3**:

<pre><code><strong>Input:</strong>          arr = [10,11,12]
<strong>Output:</strong>         66</code></pre>

**Constraints:**

<ul>
<li><code>1 <= arr.length <= 100</code></li>
<li><code>1 <= arr[i] <= 1000</code></li>
</ul>
