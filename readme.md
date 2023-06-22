# Bulk ethereum wallet generator

<br>
<h3 align="center">
  Bulk Ethereum Wallet generator
</h3>
<br/>


> Minimal code bulk address generator based on the official go-ethereum generator ⚡️ <br> thousands of addresses in a matter of seconds

## **SAFE AND FAST WALLET GENERATOR **

![ethereum and crypto wallets generated] 

- Speed: 100K wallets per second
- SQLite support
- Minimal code, low deps, so you can ensure the code is 100% safe


## Installation


```console
$ go install github.com/Chain-Identity/safe-ethereum-bulk-wallet-gen@latest
```

## Output options

There are three options for output

- **CSV** - DEFAULT. Outputs wallets to a csv file 
- **DB** - Outputs wallets to a local sqlite db
- **In memory** - Outputs wallets to memory

## Usage
	number := flag.Int("n", 100, "set number of generate times (not number of result wallet) (set number to 0 for Infinite loop ∞)")
	limit := flag.Int("limit", 100, "set limit number of result wallet. stop generate when result of vanity wallet reach the limit (set number to 0 for no limit)")
	dbPath := flag.String("db", "", "set sqlite output name that will be created in /db)")
	csvPath := flag.String("csv name", "../result.csv", "csv filename")
	concurrency := flag.Int("c", 1, "set concurrency value")
```console
Flags
  -n          int    number of wallets generated
  -db         sqlite output db path
  -csv        csv output path
  -c          int    concurrent goroutines (tl;dr speed multiplicator, c=10 is 10x times faster than c=1)
```

## Examples

### **Generate 10k addressses**

```console
$ safe-ethereum-bulk-wallet-gen -n 10000 -db ~/result.csv
```

## Thanks to

- [Planxnx](https://github.com/Planxnx/) - used some of his code here
