<!-- markdownlint-disable -->

# Find Pivot Index

Given an array of integer `nums`, calculate the **pivot index** of this array.

The **pivot index** is the index where the sum of all the numbners **strickly** to the left of the index is equal to the sum of all the numbers **strickly** to the index's right.

If the index is on the left edge of the array, then the left sum is `0` because there are no elements to the left. This also applies to the right ede of the array.

Return the **leftmost pivot index**. If no such index exists, return -1.<br>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>          nums = [1,7,3,6,5,6]
<strong>Output:</strong>         3
<strong>Explanation:</strong>
The pivot index is 3.
Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
Right sum = nums[4] + nums[5] = 5 + 6 = 11</code></pre>

**Example 2:**

<pre><code><strong>Input:</strong>          nums = [1,2,3]
<strong>Output:</strong>         -1
<strong>Explanation:</strong>
There is no index that satisfies the conditions in the problem statement.</code></pre>

**Example 3:**

<pre><code><strong>Input:</strong>          nums = [2,1,-1]
<strong>Output:</strong>         0
<strong>Explanation:</strong>
The pivot index is 0.
Left sum = 0 (no elements to the left of index 0)
Right sum = nums[1] + nums[2] = 1 + -1 = 0</code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>1 <= nums.length <= 10<sup>4</sup></code></li>
    <li><code>-1000 <= nums[i] <= 1000</code></li>
</ul>
