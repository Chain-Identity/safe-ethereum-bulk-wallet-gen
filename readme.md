# Bulk ethereum wallet generator

<br>
<h3 align="center">
  Bulk Ethereum Wallet generator
</h3>
<br/>


> Minimal code bulk address generator based on the official go-ethereum generator ⚡️ <br> thousands of addresses in a matter of seconds

## **SAFE AND FAST WALLET GENERATOR **

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
