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
```bash
$ git clone https://github.com/xht308/Diff_Algo_Demo.git
```

Second, change directory to the Go project
```bash
$ cd Diff_Algo_Demo/demo
```

Finally, build the project with Go
```bash
$ go build
```

### Basic Usage
To compare two files:
```bash
$ ./godiff -s path_to_src_file -d path_to_dest_file
```

To compare two strings:
```bash
$ ./godiff -c -s src_string -d dest_string
```

### Commandline Parameters

#### -a: Algorithm
Chose which algorithm (`basic`, `mayers`, `linearspace`) to use.
```bash
-a basic        # Basic Diff Algorithm
-a mayers       # Mayers Diff Algorithm (default)
-a linearspace  # Mayers Diff Algorithm in Linear Space
```

#### -c: Enable Character Mode
Adding this flag will make the program finding shortest edit sequence between src and dest `strings`.

#### -d: The destination file (string)
**Necessary**. The path to the destination file or the destination string if the character mode enabled.

#### -s: The source file (string)
**Necessary**. The path to the source file or the source string if the character mode enabled.

#### -t: Print the processing time
With this flag enabled, the program will append the running time of the algorithm to the output in a new line.

#### -v: Verbose Level
By give different values (0, 1, 2, 3), the program will give different levels of details in the output.
```bash
-v 0    # Disable Output
-v 1    # Basic Output      (default) (e.g. DELETE 2 2)
-v 2    # Verbose Output    (with the content of files)
-v 3    # Fancy Output      (with content and color)
```

### Recommended Usage
Compare two files
```bash
$ ./godiff -s path_to_src_file -d path_to_dest_file -v 3
```

Test Algorithms
```bash
$ ./godiff -s path_to_src_file -d path_to_dest_file -v 0 -t -a algorithm_to_test
```

Pipe output to file
```bash
$ ./godiff -s path_to_src_file -d path_to_dest_file -v 2 > output_file
```