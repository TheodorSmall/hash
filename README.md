# hash
A hash generator written in go
___
### Usage:
`hash <algorithm> <password>`

Generate a hash of `<password>` using `<algorithm>`
___
### Algorithms:
- [x] Implemented by default
- [ ] Not implemented by default


- [x] MD4
- [x] MD5
- [x] SHA1
- [x] SHA224
- [x] SHA256
- [x] SHA384
- [x] SHA512
- [ ] MD5SHA1
- [x] RIPEMD160
- [x] SHA3-224
- [x] SHA3-256
- [x] SHA3-384
- [x] SHA3-512
- [x] SHA512-224
- [x] SHA512-256
- [x] BLAKE2s-256
- [x] BLAKE2b-256
- [x] BLAKE2b-384
- [x] BLAKE2b-512
___
### Examples:
#### Valid usage:
- We use algorithm `MD5` and password `myPassword`:  
  `./hash MD5 myPassword`  
  Output: `3rFTb0gEdffVkyGaoa/XTA==`


- Different password or hash algorithm will return another hash:  
  `./hash MD5 testPassword`  
  Output: `/tO2GyYIGEk3gICzTmk9Lg==`


- We use algorithm `SHA3-512` and password `myPassword`:  
  `./hash SHA3-512 myPassword`  
  Output: `e9NoZl7gdhx+atbv1M3I9ROijXEFim5qKa1Rwi3VXp0PiXcdCeCWHq7ptzTod29gpCd7tFWiPT1IRaMazItVOA==`
___

#### Error messages:

- We use an *invalid* algorithm and password `myPassword`:  
  `./hash SHA5 myPassword`  
  Output: `Unknown algorithm`


- We use a *valid* algorithm which is not implemented by default and password `myPassword`:  
  `./hash MD5SHA1 myPassword`  
  Output: `No implementation found`


- We use only one or no command line arguments:
  `./hash`  
  This prints the [usage](#Usage:):  
  Output: `Usage: hash <algorithm> <password>`
___
