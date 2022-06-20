<!-- markdownlint-disable -->

# Can Make Arithmetic Progressive From Sequence

A sequence of numbers is called an **arithmetic progressive** if the difference between any two consecutive elements is the same.

Given an array of numbers `arr`, return `true` if the array can be rearranged to form an **arithmetic progression**. Otherwise, return `false`.

**Example 1:**

<pre><code><strong>Input:          arr = [3,5,1]</strong>
<strong>Output:</strong>         true
<strong>Explanation:</strong>    We can reorder the elements as [1,3,5] or [5,3,1] with differences 2 and -2 respectively, between each consecutive elements.</code></pre>

**Example 2:**

<pre><code><strong>Input:          arr = [1,2,4]</strong>
<strong>Output:</strong>         false
<strong>Explanation:</strong>    There is no way to reorder the elements to obtain an arithmetic progression.</code></pre>

**Constraints:**

<ul>
<li><code>2 <= arr.length <= 1000</code></li>
<li><code>-10<sup>6</sup> <= arr[i] <= 10<sup>6</sup></code></li>
</ul>
