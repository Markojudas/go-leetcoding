# Average Salary Excluding the Minimum and Maximum Salary

<!-- markdownlint-disable -->

You are given an array of <strong>unique</strong> integers `salary` where `salary[i]` is the salary of the <code>1<sup>th</sup></code> employee.

Return the average salary of employees excluding the minimum and maximum salary. Answers within <code>10<sup>-5</sup></code> of the actual answer will be acceped.

<strong>Example 1:</strong>

<pre><code><strong>Input:</strong>         salary = [4000,3000,1000,2000]
<strong>Output:</strong>        2500.00000
<strong>Explanation:</strong>   Minimum salary and maximum salary are 1000 and 4000 respectively.
Average salary excluding minimum and maximum salary is (2000+3000) / 2 = 2500
</code></pre>

<strong>Example 2:</strong>

<pre><code><strong>Input:</strong>         salary = [1000,2000,3000]
<strong>Output:</strong>        2000.00000
<strong>Explanation:</strong>   Minimum salary and maximum salary are 1000 and 3000 respectively.
Average salary excluding minimum and maximum salary is (2000) / 1 = 2000
</code></pre>

<strong>Constraints:</strong>

<pre><code><ul><li>3 <= salary.length <= 100</li><li>1000 <= salary[i] <= 10<sup>6</sup></li><li>All the integers of <em>salary</em> are <strong>unique</strong></li></ul></code></pre>
