VDF (verified delay function) implementation in Golang

This is the VDF based on  Benjanmin Wesolowski's paper "Efficient verifiable delay functions"(https://eprint.iacr.org/2018/623.pdf). In this implementation, the VDF function takes 32 bytes as seed and an integer as difficulty.

Please note that only 2048 integer size for class group variables are supported now.

