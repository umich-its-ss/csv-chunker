# csv-chunker

A simple utility to slice a CSV into smaller CSVs. Copies the header line on every additional CSV.

Example use:

```
csv-chunker a_huge.csv
```

Will produce `a_huge.csv-split-1.csv`, `a_huge.csv-split-2.csv`, and so on, until the end. Each will have 1,000 lines plus the original header.
