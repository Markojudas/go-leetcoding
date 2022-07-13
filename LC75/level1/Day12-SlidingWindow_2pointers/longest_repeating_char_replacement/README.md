<!-- markdownlint-disable -->

# Longest Repeating Character Replacement

You are given a string `s` and an integer `k`. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most `k` times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.<br>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>         s = "ABAB", k = 2
<strong>Output:</strong>        4        
<strong>Explanation:</strong>   Replace the two 'A's with two 'B's or vice versa.</code></code></pre>
<br>

**Example 2:**

<pre><code><strong>Input:</strong>         s = "AABABBA", k = 1
<strong>Output:</strong>        4        
<strong>Explanation:</strong>   Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4./code></code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>1 <= s.length <= 10<sup>5</sup></code></li>
    <li><code>s</code> consist of only uppercase English letters</li>
    <li><code>0 <= k <= s.length</code></li>
</ul>
