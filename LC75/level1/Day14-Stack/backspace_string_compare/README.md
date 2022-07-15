<!-- markdownlint-disable -->

# Backspace String Compare

Given two strings `s` and `t`, return `true` if they are equal when both are typed into empty text editors. `'#'` means a backspace character.

Note that after backspacing an empty text, the text will continue empty.<br>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>         s = "ab#c", t = "ad#c"
<strong>Output:</strong>        true
<strong>Explanation:</strong>   Both s and t become "ac".</code></pre>
<br>

**Example 2:**

<pre><code><strong>Input:</strong>         s = "ab##", t = "c#d#"
<strong>Output:</strong>        true
<strong>Explanation:</strong>   Both s and t become "".</code></pre>
<br>

**Example 3:**

<pre><code><strong>Input:</strong>         s = "a#c", t = "b"
<strong>Output:</strong>        false
<strong>Explanation:</strong>   s becomes "c" while t becomes "b".</code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>1 <= s.length, t.length <= 200</code></li>
    <li><code>s</code> and <code>t</code> only contain lowercase letters and <code>'#'</code> characters.</li>
</ul>
