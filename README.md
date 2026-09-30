# FNO Experimental Evaluation

This repository contains the experimental evaluation of the **Fabric Network Orchestrator (FNO)**, a model-driven command-line tool for configuring, validating, and deploying Hyperledger Fabric networks.

The evaluation examines the configuration effort required by FNO, its capacity to detect invalid configurations, and the correctness of the generated deployments.

## Experiments

| Experiment | Description |
| --- | --- |
| [Configuration Effort](configuration-effort/) | Compares the lines of configuration code required by FNO, Fablo, and manual deployment across five network scenarios. |
| [Detection of Invalid Configurations](detection-of-invalid-configurations/) | Evaluates whether FNO detects and reports the expected validation rules in configurations containing one to five simultaneous mutations. |
| [Deployment Correctness](deployment-correctness/) | Evaluates five deployed network configurations using captured container states, component logs, and chaincode transaction outputs. |

Each experiment directory contains its configurations, results or runtime evidence, and a README describing the evaluation artifacts.
