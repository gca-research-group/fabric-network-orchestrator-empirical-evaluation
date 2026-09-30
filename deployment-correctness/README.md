# FNO Experimental Evaluation — Deployment Correctness

This evaluation examines whether FNO correctly deploys the networks described by its configurations. Five configurations were deployed, and their container state, component logs, and available functional-test outputs were preserved as runtime evidence.

## Scenarios

| Scenario | Evaluated characteristic |
| --- | --- |
| `s1` | Two organizations, one peer per organization, one orderer, and one channel |
| `s2` | Five organizations, two peers per organization, one orderer, and one channel |
| `s3` | Three organizations, three channels, and two channel profiles |
| `s4` | Mixed Fabric peer, orderer, and CA versions |
| `s5` | Three chaincodes, an endorsement policy, and a private-data collection |

## Evidence

Each scenario directory contains:

- `config.yml`: network configuration supplied to FNO.
- `containers.txt`: captured state, image, and version of the deployed containers.
- `container-logs/`: logs captured from the deployed Fabric components and available chaincode containers.

Scenario `s5` additionally contains:

- `chaincodes/`: chaincode source code and the private-data collection configuration used by the functional checks.
- `*-invoke.txt` and `*-query.txt`: commands and successful transaction or query outputs for the Asset, Product, and PrivateAgreement chaincodes.

The captured artifacts show that FNO instantiated the requested network topologies. Scenario `s5` additionally demonstrates successful ledger writes and reads for a regular chaincode, a chaincode with an endorsement policy, and a chaincode using private data.
