# run to simulate
```bash
git clone git@github.com:silagadzeshota/Validator.git
cd Validator
export GO111MODULE=on
export DUTY_WEBSOCKET_URL=ws://127.0.0.1:5000
go mod tidy
go run main.go
// start websocket AFTER starting dutyprocessor
```
<details>
<summary>Expected output</summary>
<div class="highlight highlight-source-shell"><pre>
2023/06/18 20:55:41 listening for incoming duties to process
Validator  5  created and started listening for incoming duties
2023/06/18 20:55:44 Validator  5 : Received new duty  PROPOSER  for the height  0
2023/06/18 20:55:44 Validator  5 : Processed duty  PROPOSER  for the height  0
2023/06/18 20:55:47 Validator  5 : Received new duty  ATTESTER  for the height  0
2023/06/18 20:55:47 Validator  5 : Processed duty  ATTESTER  for the height  0
Validator  1  created and started listening for incoming duties
2023/06/18 20:55:50 Validator  1 : Received new duty  PROPOSER  for the height  0
2023/06/18 20:55:50 Validator  1 : Processed duty  PROPOSER  for the height  0
2023/06/18 20:55:53 Validator  5 : Received new duty  AGGREGATOR  for the height  0
2023/06/18 20:55:53 Validator  5 : Processed duty  AGGREGATOR  for the height  0
2023/06/18 20:55:56 Validator  5 : Received new duty  SYNC_COMMITTEE  for the height  0
2023/06/18 20:55:56 Validator  5 : Processed duty  SYNC_COMMITTEE  for the height  0
2023/06/18 20:55:56 Validator  5  moved to height  1
2023/06/18 20:55:59 Validator  1 : Received new duty  ATTESTER  for the height  0
2023/06/18 20:55:59 Validator  1 : Processed duty  ATTESTER  for the height  0
Validator  3  created and started listening for incoming duties
2023/06/18 20:56:02 Validator  3 : Received new duty  PROPOSER  for the height  0
2023/06/18 20:56:02 Validator  3 : Processed duty  PROPOSER  for the height  0
Validator  4  created and started listening for incoming duties
2023/06/18 20:56:05 Validator  4 : Received new duty  PROPOSER  for the height  0
2023/06/18 20:56:05 Validator  4 : Processed duty  PROPOSER  for the height  0
Validator  2  created and started listening for incoming duties
2023/06/18 20:56:08 Validator  2 : Received new duty  PROPOSER  for the height  0
2023/06/18 20:56:08 Validator  2 : Processed duty  PROPOSER  for the height  0
2023/06/18 20:56:11 Validator  2 : Received new duty  ATTESTER  for the height  0
2023/06/18 20:56:11 Validator  2 : Processed duty  ATTESTER  for the height  0
2023/06/18 20:56:14 Validator  3 : Received new duty  ATTESTER  for the height  0
2023/06/18 20:56:14 Validator  3 : Processed duty  ATTESTER  for the height  0
Validator  6  created and started listening for incoming duties
2023/06/18 20:56:17 Validator  6 : Received new duty  PROPOSER  for the height  0
2023/06/18 20:56:17 Validator  6 : Processed duty  PROPOSER  for the height  0
2023/06/18 20:56:20 Validator  2 : Received new duty  AGGREGATOR  for the height  0
2023/06/18 20:56:20 Validator  2 : Processed duty  AGGREGATOR  for the height  0
2023/06/18 20:56:23 Validator  4 : Received new duty  ATTESTER  for the height  0
2023/06/18 20:56:23 Validator  4 : Processed duty  ATTESTER  for the height  0
2023/06/18 20:56:26 Validator  4 : Received new duty  AGGREGATOR  for the height  0
2023/06/18 20:56:26 Validator  4 : Processed duty  AGGREGATOR  for the height  0
2023/06/18 20:56:29 Validator  2 : Received new duty  SYNC_COMMITTEE  for the height  0
2023/06/18 20:56:29 Validator  2 : Processed duty  SYNC_COMMITTEE  for the height  0
2023/06/18 20:56:29 Validator  2  moved to height  1
2023/06/18 20:56:32 Validator  3 : Received new duty  AGGREGATOR  for the height  0
2023/06/18 20:56:32 Validator  3 : Processed duty  AGGREGATOR  for the height  0
2023/06/18 20:56:35 Validator  6 : Received new duty  ATTESTER  for the height  0
2023/06/18 20:56:35 Validator  6 : Processed duty  ATTESTER  for the height  0
2023/06/18 20:56:38 Validator  5 : Received new duty  PROPOSER  for the height  1
2023/06/18 20:56:38 Validator  5 : Processed duty  PROPOSER  for the height  1
2023/06/18 20:56:41 Validator  1 : Received new duty  AGGREGATOR  for the height  0
2023/06/18 20:56:41 Validator  1 : Processed duty  AGGREGATOR  for the height  0
2023/06/18 20:56:44 Validator  2 : Received new duty  PROPOSER  for the height  1
2023/06/18 20:56:44 Validator  2 : Processed duty  PROPOSER  for the height  1
2023/06/18 20:56:47 Validator  2 : Received new duty  ATTESTER  for the height  1
2023/06/18 20:56:47 Validator  2 : Processed duty  ATTESTER  for the height  1
2023/06/18 20:56:50 Validator  6 : Received new duty  AGGREGATOR  for the height  0
2023/06/18 20:56:50 Validator  6 : Processed duty  AGGREGATOR  for the height  0
2023/06/18 20:56:53 Validator  3 : Received new duty  SYNC_COMMITTEE  for the height  0
2023/06/18 20:56:53 Validator  3 : Processed duty  SYNC_COMMITTEE  for the height  0
2023/06/18 20:56:53 Validator  3  moved to height  1
2023/06/18 20:56:56 Validator  2 : Received new duty  AGGREGATOR  for the height  1
2023/06/18 20:56:56 Validator  2 : Processed duty  AGGREGATOR  for the height  1</pre></div>
</details>

## test
```bash
go test ./...
```
