<!-- markdownlint-disable -->

# Goal Parser Interpretation

You own a **Goal Parser** that can interpret a string `command`. The `command` consists of an alphabet of `"G"`, `"()"`, and/or `"(al)"` in some order. The Goal Parser will interpret `"G"` as the string `"G"`, `"()"` as the string `"o"`, and `"(al)"` as the string `"al"`. The interpreted strings are then concatenated in the original order.

Given the string `comand`, return the **Goal Parser**'s interpretation of `comand`.
<br>

<br>

**Example 1:**

<pre><code><strong>Input:</strong>          command = "G()(al)"
<strong>Output:</strong>         "Goal"
<strong>Explanation:</strong>    The Goal Parser interprets the command as follows:
G -> G
() -> o
(al) -> al
The final concatenated result is "Goal".</code></pre>

**Example 2:**

<pre><code><strong>Input:</strong>          command = "G()()()()(al)"
<strong>Output:</strong>         "Gooooal"</code></pre>

**Example 3**:

<pre><code><strong>Input:</strong>          command = "(al)G(al)()()G"
<strong>Output:</strong>         "alGalooG"</code></pre>

**Constraints:**

<ul>
<li><code>1 <= command.length <= 100</code></li>
<li><code>command</code> consists of <code>"G"</code>, <code>"()"</code>, and/or <code>"(al)"</code> in some order.</li>
</ul>
