# cs-lab-go
Computer Science Lab in Go

## Sorting Algorithm

### Slice Based

| Algorithm      | Data Structure                                      | Time Complexity |
|----------------|-----------------------------------------------------|-----------------|
| Bubble Sort    | [Slice](pkg/algorithm/sort/slice_bubble_sort.go)    | O(n²)           |
| Cycle Sort     | [Slice](pkg/algorithm/sort/slice_cycle_sort.go)     | O(n²)           |
| Exchange Sort  | [Slice](pkg/algorithm/sort/slice_cycle_sort.go)     | O(n²)           |
| Gnome Sort     | [Slice](pkg/algorithm/sort/slice_gnome_sort.go)     | O(n²)           |
| Insertion Sort | [Slice](pkg/algorithm/sort/slice_insertion_sort.go) | O(n²)           |
| Selection Sort | [Slice](pkg/algorithm/sort/slice_selection_sort.go) | O(n²)           |
| Heap Sort      | [Slice](pkg/algorithm/sort/slice_heap_sort.go)      | O(n log n)      |

### Single Linked List Based Implementation

| Algorithm      | Data Structure                                                                | Time Complexity |
|----------------|-------------------------------------------------------------------------------|-----------------|
| Insertion Sort | [Single Linked List](pkg/algorithm/sort/single_linked_list_insertion_sort.go) | TBD             |

### Double Linked List Based Implementation
 
| Algorithm      | Data Structure                                                                | Time Complexity |
|----------------|-------------------------------------------------------------------------------|-----------------|
| Insertion Sort | [Double Linked List](pkg/algorithm/sort/single_linked_list_insertion_sort.go) | TBD             |

## Search Algorithm

### Slice Based Implementation

| Algorithm      | Data Structure                                       | Time Complexity |
|----------------|------------------------------------------------------|-----------------|
| Bubble Sort    | [Slice](pkg/algorithm/search/slice_binary_search.go) | log n           |