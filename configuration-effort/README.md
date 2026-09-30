# FNO Experimental Evaluation — Configuration Effort

This repository contains the experimental evaluation and benchmark scenarios for the **Fabric Network Orchestrator (FNO)**, a domain-specific model-driven CLI tool designed to simplify the deployment, configuration, and validation of Hyperledger Fabric networks. 

By defining the desired network architecture in a single, high-level YAML specification, FNO handles cryptographic identity generation, docker-compose configuration orchestration, channel creation, and chaincode deployment, while performing rigorous syntactic and semantic checks.

---

## Benchmark Results: Lines of Code (LoC)

To evaluate the efficiency of FNO, we compare the required Lines of Code (LoC) to deploy identical networks under three environments:
1. **Baseline**: Manual deployment using raw Docker Compose files, `configtx.yaml` profiles, and custom shell scripting.
2. **Fablo**: A popular open-source generator tool for Hyperledger Fabric networks.
3. **FNO (Ours)**: Our proposed model-driven orchestration approach.

The following table summarizes the comparative LoC for five standard deployment scenarios:

| Scenario | Network Characteristics | Baseline (LoC) | Fablo (LoC) | FNO (LoC) |
| :--- | :--- | :---: | :---: | :---: |
| **Scenario s1** | 2 Orgs, 1 Channel, 1 Peer/Org, 1 Orderer (Org1) | 594 | 31 | **29** |
| **Scenario s2** | 5 Orgs, 1 Channel, 2 Peers/Org, 1 Orderer (Org1) | 1,518 | **52** | 57 |
| **Scenario s3** | 3 Orgs, 3 Channels, 1 Peer/Org, 1 Orderer (Org1) | 851 | 52 | **44** |
| **Scenario s4** | 3 Orgs, 1 Channel, Exposed Ports, 2 Chaincodes, Endorsement Policy | 594 | *N/A (Unsupported)* | **36** |
| **Scenario s5** | 3 Orgs, 1 Channel, Exposed Ports, Private Data Collections, 3 Chaincodes | 822 | 60 | **58** |

### Key Takeaways

> [!TIP]
>
> * **FNO substantially reduces deployment-specific configuration compared with manual deployment.** Across the five benchmark scenarios, FNO requires approximately **95% fewer Lines of Code (LoC)** than the baseline composed of Docker Compose files, Fabric configuration artefacts, and deployment scripts. This result indicates that a high-level declarative specification can considerably reduce the amount of deployment-specific configuration that developers must prepare and maintain.
>
> * **FNO and Fablo provide similarly concise specifications for the scenarios supported by both approaches, while differing in configuration coverage.** FNO requires fewer LoC in Scenarios **s1, s3, and s5**, whereas Fablo requires fewer LoC in **s2**. The main distinction is therefore not simply specification size: **Scenario s4** could not be represented with the evaluated Fablo configuration because of its required endorsement-policy configuration, whereas FNO supports the scenario within the same declarative specification model.

---


## Repository Structure

```
├── baseline/                        # Manual deployment files & helper scripts
│   └── s1 to s5/
├── fablo/                           # Fablo configuration specs
│   └── s1 to s5/
├── fabric-network-orchestrator/     # Fabric Network Orchestrator (FNO) workspace
│   ├── fno.exe                      # Compiled FNO CLI executable
│   └── s1 to s5/                    # Benchmark scenario configurations for FNO
├── cloc-2.08.exe                    # CLOC executable used to measure Lines of Code
├── cloc.sh                          # Shell script to automatically run the CLOC benchmark
└── README.md                        # Project documentation (this file)
```

---

## Getting Started & Execution

### 1. Run the LoC Comparison Benchmark
To compute the LoC comparison table on a Bash environment:
```bash
./cloc.sh
```

*(Note: Ensure line endings on the script are LF or run it via PowerShell).*
