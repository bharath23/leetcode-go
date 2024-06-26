#### Regular Expression Matching

Given an input string (`s`) and a pattern (`p`), implement regular expression
matching with support for '`.`' and '`*`' where: 

* '`.`' Matches any single character.
* '`*`' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).

**Example 1**:
<pre><code><b>Input</b>: s = "aa", p = "a"
<b>Output</b>: false
<b>Explanation</b>: "a" does not match the entire string "aa".
</code></pre>

**Example 2**:
<pre><code><b>Input</b>: s = "aa", p = "a*"
<b>Output</b>: true
<b>Explanation</b>: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating
            'a' once, it becomes "aa". 
</code></pre>

**Example 3**:
<pre><code><b>Input</b>: s = "ab", p = ".*"
<b>Output</b>: true
<b>Explanation</b>: ".*" means "zero or more (*) of any character (.)".
</code></pre>

**Example 4**:
<pre><code><b>Input</b>: s = "aab", p = "c*a*b"
<b>Output</b>: true
<b>Explanation</b>: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches
            "aab".
</code></pre>

**Example 5**:
<pre><code><b>Input</b>: s = "mississippi", p = "mis*is*p*."
<b>Output</b>: false
</code></pre>

**Constraints**:
* `0 <= s.length <= 20`
* `0 <= p.length <= 30`
* `s` contains only lowercase English letters.
* `p` contains only lowercase English letters, '`.`', and '`*`'.
* It is guaranteed for each appearance of the character '`*`', there will be a
previous valid character to match.
