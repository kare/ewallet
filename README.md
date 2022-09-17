# kkn.fi/cmd/ewallet

## Description
`ewallet` is an experimental Ethereum command line utility that converts an
Ethereum private key to address or public key and to converts an Ethereum adress
to mixed-case checksum address encoding as defined in
[EIP-55](https://eips.ethereum.org/EIPS/eip-55).

## Install

```bash
go install kkn.fi/cmd/ewallet
```

## Command line examples

### Generate new private key
```bash
ewallet new
```
#### Output
```
7488b71626c424bfd2b0c9b3c6a83a00c1de6312453f016d5f40dfa86a96409c
```

### Convert private key to address
```bash
ewallet address 7488b71626c424bfd2b0c9b3c6a83a00c1de6312453f016d5f40dfa86a96409c
```
#### Output
```
0x655A7A1f0E1819e3395723DDD2a9D900fAff5cFB
```

### Convert private key to public key
```bash
ewallet public 7488b71626c424bfd2b0c9b3c6a83a00c1de6312453f016d5f40dfa86a96409c
```
#### Output
```
ef3561edac05d3f0b5becc5d851686fffee6b8565ae471510008b3124b000cd5e7f3b2f58ad7e8912f153f04444d26bbb2de1aa13c1691a9d2a04c0b3c111216
```

### Convert address to checksum case
```bash
ewallet checksum 0X655a7a1F0e1819E3395723ddd2A9d900FaFF5Cfb
```
#### Output
```
0x655A7A1f0E1819e3395723DDD2a9D900fAff5cFB
```
