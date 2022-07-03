<!-- markdownlint-disable -->

# is Subsequence

Given two strings `s` and `t`, return `true` if `s` is a **subsequence** of `t`, or `false` otherwise.

A **subsequence** of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., `"ace"` is a subsequence of <code>"<u>a</u>b<u>c</u>d<u>e</u>"</code> while `"aec"` is not).<br>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>    s = "abc", t = "ahbgdc"
<strong>Output:</strong>   true</code></pre>

**Example 2:**

<pre><code><strong>Input:</strong>    s = "axc", t = "ahbgdc"
<strong>Output:</strong>   false</code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>0 <= s.length <= 100</code></li>
    <li><code>0 <= t.length <= 10<sup>4<sup></code></li>
    <li><code>s</code> and <code>t</code> consist only of lowercase English letters.</li>
</ul>
