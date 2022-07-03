<!-- markdownlint-disable -->

# Isomorphic Strings

Given two strings `s` and `t`, determine if they are isomorphic.

Two strings `s` and `t` are isomorphic if the characters in `s` can be replaced to get `t`.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.<br>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>   s = "egg", t = "add"
<strong>Output:</strong>  true</code></pre>

**Example 2:**

<pre><code><strong>Input:</strong>   s = "foo", t = "bar"
<strong>Output:</strong>  false</code></pre>

**Example 3:**

<pre><code><strong>Input:</strong>   s = "paper", t = "title"
<strong>Output:</strong>  true</code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>1 <= s.length <= 5 * 10<sup>4</sup></code></li>
    <li><code>t.length == s.length</code></li>
    <li><code>s</code> and <code>t</code> consist of any valid ascii character</li>
</ul>
