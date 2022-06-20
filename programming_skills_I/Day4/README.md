<!-- markdownlint-disable -->

# Sign of the Product of an Array

There is a function `signFunc(x)` that returns:

<ul>
    <li><code>1</code> if <code>x</code> is positive</li>
    <li><code>-1</code> if <code>x</code> is negative</li>
    <li><code>0</code> if <code>x</code> is equal to 0.
</ul>

You are given an integer array `nums`. Let `product` be the product of all values in the array `nums`.

Return `signFunc(product)`.

**Example 1:**

<pre><code><strong>Input:          nums = [-1,-2,-3,-4,3,2,1]</strong>
<strong>Output:</strong>         1
<strong>Explanation:</strong>    The product of all values in the array is 144, and signFunc(144) = 1</code></pre>

**Example 2:**

<pre><code><strong>Input:          nums = [1,5,0,2,-3]</strong>
<strong>Output:</strong>         0
<strong>Explanation:</strong>    The product of all values in the array is 0, and signFunc(0) = 0</code></pre>

**Example 3:**

<pre><code><strong>Input:          nums = [-1,1,-1,1,-1]</strong>
<strong>Output:</strong>         -1
<strong>Explanation:</strong>    The product of all values in the array is -1, and signFunc(-1) = -1</code></pre>

**Constraints:**

<ul>
<li><code>1 <= nums.length <= 1000</code></li>
<li><code>-100 <= nums[i] <= 100</code></li>
</ul>
