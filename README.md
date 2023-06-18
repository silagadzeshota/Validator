<h2><a id="user-content-whats-tron" class="anchor" aria-hidden="true" href="#whats-tron"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>What is Validator</h2>
<p>Validator processes incoming duty requests for each validator. The service starts listening to the websocket on the address exported in DUTY_WEBSOCKET_URL.
After receiving the duty checks if the validator with given ID exists, if not it creates a new validator, registers it into the duty processor and starts it. Each started validator is waiting for the incoming duties. After receiving duty checks if the duty received or even processed already, if so - does nothing.</p>
<p>If the duty is not received registers the duty for the corresponding height and if the height is the current height that it is processing starts processing the duty immediately. After the validator processes all four types of duties for the specific height it starts waiting for and processing duties for the next height. Note: validator starts processing duties from the height 0, if the first duties are received for the higher heights it will just store those as requests but will not process them until it has processed all the duty types for all the lower heights</p>

# run to simulate
```bash
git clone git@github.com:silagadzeshota/Validator.git
cd Validator
export GO111MODULE=on
export DUTY_WEBSOCKET_URL=ws://127.0.0.1:5000
go mod tidy
go run main.go
```

For simulation - after starting the validator start python script for generating duties
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

## Testing
```bash
go test ./...
```
