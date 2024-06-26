#### Flatten a Multilevel Doubly Linked List
You are given a doubly linked list, which contains nodes that have a next pointer, a previous pointer, and an additional  **child pointer**. This child pointer may or may not point to a separate doubly linked list, also containing these special nodes. These child lists may have one or more children of their own, and so on, to produce a  **multilevel data structure**  as shown in the example below.

Given the  `head`  of the first level of the list,  **flatten**  the list so that all the nodes appear in a single-level, doubly linked list. Let  `curr`  be a node with a child list. The nodes in the child list should appear  **after**  `curr`  and  **before**  `curr.next`  in the flattened list.

Return  _the_ `head` _of the flattened list. The nodes in the list must have  **all**  of their child pointers set to_ `null`.

**Example 1**:

![](example_1.jpg)
<pre><code><b>Input</b>: head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
<b>Output</b>: [1,2,3,7,8,11,12,9,10,4,5,6]
<b>Explanation</b>: The multilevel linked list in the input is shown.
After flattening the multilevel linked list it becomes:
</code>
<img src="example_1_flat.jpg" /></pre>

**Example 2**:

![](example_2.jpg)
<pre><code><b>Input</b>: head = [1,2,null,3]
<b>Output</b>: [1,3,2]
<b>Explanation</b>: The multilevel linked list in the input is shown.
After flattening the multilevel linked list it becomes:
</code>
<img src="example_2_flat.jpg" /></pre>

**Example 3**:
<pre><code><b>Input</b>: head = []
<b>Output</b>: []
<b>Explanation</b>: There could be empty list in the input.
</code></pre>

**Constraints**:
* The number of Nodes will not exceed `1000`.
* <code>1 <= Node.val <= 10<sup>5</sup></code>

**How the multilevel linked list is represented in test cases:**
We use the multilevel linked list from  **Example 1**  above:
<pre>
<code> 1---2---3---4---5---6--NULL
         |
         7---8---9---10--NULL
             |
             11--12--NULL
</code></pre>

The serialization of each level is as follows:
<pre><code>[1,2,3,4,5,6,null]
[7,8,9,10,null]
[11,12,null]
</code></pre>
To serialize all levels together, we will add nulls in each level to signify no node connects to the upper node of the previous level. The serialization becomes:
<pre><code>[1,    2,    3, 4, 5, 6, null]
             |
[null, null, 7,    8, 9, 10, null]
                   |
[            null, 11, 12, null]
</code></pre>
Merging the serialization of each level and removing trailing nulls we obtain:
<pre><code>[1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
</code></pre>
