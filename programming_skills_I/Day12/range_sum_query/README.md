<!-- markdownlint-disable -->

# Range Sum Query - Immutable

<p>Given an integer array <code>nums</code>, handle multiple queries of the following type:</p>
<ol>
    <li> Calculate the <strong>sum</strong> of the elements of <code>nums</code> between indices <code>left</code> and <code>right</code> <strong>inclusive</strong> where <code>left <= right</code>.
</ol>
<p>Implement the <code>NumArray</code> class:
<ul>
    <li><code>NumArray(int[] nums)</code> Initializes the object with the integer array <code>nums</code>.
    <li><code>int sumRange(int left, int right)</code> Returns the <strong>sum</strong> of the elements of <code>nums</code> between indices <code>left</code> and <code>right</code> <strong>inclusive</strong> (i.e. <code>nums[left] + nums[left + 1] + ... + nums[right]</code>).
</ul>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
<strong>Output:</strong>
[null, 1, -1, -3]
<strong>Explanation:</strong>
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3</code></pre>
<br>

**Constraints:**

<ul>
    <li><code>1 <= nums.length <= 10<sup>4</sup></code></li>
    <li><code>-10<sup>5</sup> <= nums[i] <= 10<sup>5</sup></code></li>    
    <li><code>0 <= left <= right < nums.length</code></li>
    <li>At most <code>10<sup>4</sup></code> calls will be made to <code>sumRange</code>.</li>    
</ul>
