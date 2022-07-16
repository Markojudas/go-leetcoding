<!-- markdownlint-disable -->

# Decode String

Give an encoded string, return its decoded string.

The encoding rule is: `k[encoded_string]`, where `encoded_string` inside the square brackets is being repeated exactly `k` times. Note that `k` is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, `k`. For example, there will not be input like `3a` or `2[4]`.

The test cases are generated so that the length of the output will never exceed <code>10<sup>4</sup></code><br>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>         s = "3[a]2[bc]"
<strong>Output:</strong>        "aaabcbc"</code></pre>
<br>

**Example 2:**

<pre><code><strong>Input:</strong>         s = "3[a2[c]]"
<strong>Output:</strong>        "accaccacc"</code></pre>
<br>

**Example 3:**

<pre><code><strong>Input:</strong>         s = "2[abc]3[cd]ef"
<strong>Output:</strong>        "abcabccdcdcdef"</code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>1 <= s.length <= 30</code></li>
    <li><code>s</code> consists of lowercase English letters, digits, and square brackets <code>'[]'</code>.</li>
    <li><code>s</code> guaranteed to be a <strong>valid</strong> input.</li>
    <li>All the integers in <code>s</code> are in the range <code>[1, 300]</code>.</li>
</ul>
