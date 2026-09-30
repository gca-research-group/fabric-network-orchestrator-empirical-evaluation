# FNO Experimental Evaluation — Detection of Invalid Configurations

This evaluation examines FNO's capacity to detect and report invalid network configurations. Mutation testing was applied to the valid [`seed.yaml`](seed.yaml) configuration, producing datasets with one to five simultaneous mutations per scenario (`mc1`–`mc5`).

## Results

| Mutation order (*k*) | Configurations | Introduced violations | Detected violations | Missed violations | Unexpected reports |
| ---: | ---: | ---: | ---: | ---: | ---: |
| 1 | 58 | 58 | 58 | 0 | 0 |
| 2 | 1,575 | 3,150 | 3,150 | 0 | 0 |
| 3 | 26,827 | 80,481 | 80,481 | 0 | 0 |
| 4 | 323,287 | 1,293,148 | 1,293,148 | 0 | 0 |
| 5 | 2,943,011 | 14,715,055 | 14,715,055 | 0 | 0 |
| **Total** | **3,294,758** | **16,091,892** | **16,091,892** | **0** | **0** |

A scenario passes when FNO reports every validation rule expected from its injected mutations. In this experiment, all generated invalid configurations passed this criterion.

## Repository Structure

- `seed.yaml`: valid FNO configuration used as the mutation seed.
- `mc1/` to `mc5/`: results grouped by the number of simultaneous mutations.
  - `scenarios.parquet`: generated scenarios and their applied mutation rules.
  - `results.parquet`: expected, reported, and missing rules, together with validation status and errors.
  - `metadata.json`: generation and validation parameters, counts, timings, and execution status.
