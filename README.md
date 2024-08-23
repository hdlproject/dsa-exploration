# Data Structure and Algorithm Exploration

This repository contains any code that is related to Data Structure & Algorithm
and its applications in problem-solving.

## Table of Contents

Here is the table of contents:

### Basic

- [Queue](./queue): a data structure for storing data in FIFO (First In First Out) order.
- [Stack](./stack): a data structure for storing data in LIFO (Last In First Out) order.

### Advanced
- [Tree](./tree): a data structure for storing data hierarchically.
    - [Binary Search Tree](./tree/bst.go): a data structure for storing sorted hierarchical data with each node having
      at most two children. The left one is smaller from the parent and the right one is bigger.
        - [Breadth First Search](./tree/bst_bfs.go): an algorithm to traverse the tree elements begin from all the root
          node adjacent moving downward.
        - [Depth First Search](./tree/bst_dfs.go): an algorithm to traverse the tree elements as far as possible from
          the root starting from the most left adjacent then moving back to the parent.
    - [Balanced Binary Search Tree](./tree/balanced_bst.go): a special type of Binary Search Tree which maintains the
      height in each alteration (insert/update/delete). Balanced Binary Search Tree in this repository implements AVL
      Tree. The core concept of this data structure is doing checking in every alteration and rotating the branch if the
      alteration cause imbalance.
- [Shunting Yard](./shunting-yard): ???
