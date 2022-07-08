<!-- markdownlint-disable -->

# First Bad Version

You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have `n` versions `[1, 2, ..., n]` and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool `isBadVersion(version)` which returns whether `version` is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.<br>
<br>

<pre><code>/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */</code></pre>

**Example 1:**

<pre><code><strong>Input:</strong>     n = 5, bad = 4
<strong>Output:</strong>    4
<strong>Explanation:</strong>
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version.</code></pre>
<br>

**Example 2:**

<pre><code><strong>Input:</strong>     n = 1, bad = 1
<strong>Output:</strong>    1</code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>1 <= bad <= n <= 2<sup>31</sup> - 1</code></li>
</ul>
