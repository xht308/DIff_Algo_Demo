# GoDiff

> Wenzhou-Kean University - Fall 2023
> 
> CPS 3410 Applied Alg. & Data Structures
> 
> Individual Project - Code Part

A simple commandline diff2 tool written in Go.

## Algorithms
A total of **three** diff algorithms were implemented and used in this project to compare the two files and propose the minimum edit operations needed.

*For more details about the algorithms implemented, please check the **report** of the project in the same repo.*

### Basic Diff Algorithm
Implementation: `basic_diff.go`

Time Complexity: `O(mn)`

Space Complexity: `O(mn)`

Inspired by the dynamic programming solution of the Longest Common Subsequence (LCS) problem, the basic diff algorithm is simply a modification of the LCS algorithm counting the minimum number of mismatches instead of matches and tracing the shortest edit operations sequence along the same path as the longest subsequence.

### Mayers Diff Algorithm
Implementation: `mayers_diff.go`

Time Complexity: `O(M+N+D^2)` expected / `O((M+N)D)` worest

Space Complexity: `O(D^2)`

Firstly proposed by Eugene W. Myers in 1986 in the paper titled [AnO(ND) difference algorithm and its variations](https://doi.org/10.1007/BF01840446), the Mayers Diff Algorithm searches the shortest edit sequence in a way similar to breadth-first search and might be the most widely-used diff algorithm today. It offers better performance in common use cases (Version Control, Genetic Sequence Differentiation, etc.)

### Mayers Diff Algorithm in Linear Space
Implementation: `mayers_linear_space.go`

Time Complexity: `O((M+N)D)`

Space Complexity: `O(M+N)`

As the `O(D^2)` space complexity of the basic mayers diff algorithm is not practical for the use cases involving the processing of large files, Eugene W. Myers proposed the space-optimized version of the algorithm in the same [paper](https://doi.org/10.1007/BF01840446) utilizing the divide-and-conquer and recursive method to elimiate the need to keep historical records and make the space complexity linear. Similar to quick sort, the algorithm tends to find the matching sequence (middle snake) at the middle of edit sequence and divide the problem into two smaller ones in each iteration.

## Usage
*Only **windows** executable is provided in the releases. You may need to build your own binary if using other platforms.*

### Build Binary
Prerequisites
```
Go >= 1.21.0
```

First, clone the git repository to the local machine
```
$ git clone 
```