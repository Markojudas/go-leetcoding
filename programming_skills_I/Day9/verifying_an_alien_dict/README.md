<!-- markdownlint-disable -->

# Verifying an Alien Dictionary

In an alien language, suprisingly, they also use English lowercase letters, but possibly in a different `order`. the `order` of the alphabet is some permutation of lowercase letters.

Given a sequence of `wrote` written in the alien language, and the `order` of the alphabet, return `true` if and only if the given `words` are sorted lexicographically i the this alien language.
<br>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>          words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
<strong>Output:</strong>         true
<strong>Explanation:</strong>    As 'h' comes before 'l' in this language, then the sequence is sorted.</code></pre>

**Example 2:**

<pre><code><strong>Input:</strong>          words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
<strong>Output:</strong>         false
<strong>Explanation:</strong>    As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.</code></pre>

**Example 3:**

<pre><code><strong>Input:</strong>          words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
<strong>Output:</strong>         false
<strong>Explanation:</strong>    The first three characters "app" match, and the second string is shorter (in size.) According to lexicographical rules "apple" > "app", because 'l' > '∅', where '∅' is defined as the blank character which is less than any other character (<a href="https://en.wikipedia.org/wiki/Lexicographical_order">More info</a>)</code></pre>

<br>
**Constraints:**

<ul>
    <li><code>1 <= words.length <= 100</code></li>
    <li><code>1 <= words[i].length <= 20</code></li>
    <li><code>order.length == 26</code></li>
    <li>All characters in <code>words[i]</code> and <code>order</code> are English lowercase letters</li>

</ul>
