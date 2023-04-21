# ai_transfer

1. Compile 'go_send' and copy the binary file to the same level as 'ai_transfer.py'.
2. Start a Deeper-Chain node (https://github.com/deeper-chain/deeper-chain) or another Substrate-based blockchain.
3. Run 'ai_transfer.py'.
4. Input a message such as 'Transfer from Alice to Bob: 999900000000000000000'.
You will receive a response similar to: "

```
Entering new LLMBashChain chain...
transfer from Alice  to Bob  999900000000000000000

```bash
./go_send Alice Bob 999900000000000000000
```['```bash', './go_send Alice Bob 999900000000000000000', '```']

Answer: 2023/04/21 16:34:00 {//Alice 5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY [212 53 147 199 21 253 211 28 97 20 26 189 4 169 159 214 130 44 133 88 133 76 205 227 154 86 132 231 165 109 162 125]} {//Bob 5FHneW46xGXgs5mUiveU4sbTyGBzmstUspZC92UhjJM694ty [142 175 4 21 22 135 115 99 38 201 254 161 126 37 252 82 135 97 54 147 201 18 144 156 178 38 170 71 148 242 106 72]}
2023/04/21 16:34:00 Connecting to ws://127.0.0.1:9944...
0500008eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a4817000016814d68663436002091012b0000000e000000e522e1b46c87f267c6bbfdc587b789b3617eb4e47369877c3df19ea0f0660383e522e1b46c87f267c6bbfdc587b789b3617eb4e47369877c3df19ea0f0660383
 end
2023/04/21 16:34:00 Balance transferred hash  0x19ec13ac6d7b31e6e74e0c3063867736f6e49af0f9e8ca46c33c544b50c50abc

> Finished chain.

```
