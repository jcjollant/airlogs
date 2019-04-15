# Overview
## AirLogs : Usable Aviation Logbooks.
In a nutshell, we are bringing blockchain to aviation records.

### Why?
Logbooks, such as Aircraft Maintenance, are prime candidates to be lost, stolen, destroyed or withheld. We can remediate this.
General Aviation is a fantastic community. It deserves the right tools.

### What?
This project will bring Distributed Ledgers (aka Blockchains) efficiency to General Aviation.
The current objective is to address aircraft maintenance records.
Our first deliverable is a public demo environment with transferable aircraft and associated maintenance records.

### How?
Our underlying Distributed Ledger is [HyperLedger Fabric](https://www.hyperledger.org/projects/fabric "HyperLedger Fabric") .
This open source project will provide all necessary software and documentation to operate the Ledger (inluding Settings, Smart Contract, Membership Service Provider, ...).

### Business models
This project is and shall remain a nonprofit venture.

# QuickStart
Instanciate fabric-samples first-network using ./byfn.sh up (using 1.4)
Build chaincode with 'go build'
From the cli docker (docker exec -it bash) Instanciate chaincode using 'peer'
Execute a query using 'peer'
