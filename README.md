# stbp - an implementation of СТБ П 34.101.31-2007 in Golang

This project was done as part of a term paper on cryptography. Repository contains the implementation of the core block cipher, its modes of operation, as defined in СТБ П 34.101.31-2007:
- ECB (Electronic CodeBook) mode,
- CBC (Cipher Block Chaining) mode,
- CFB (Cipher FeedBack) mode,
- CTR (CounTeR) mode,
as well as data authentication through MAC (Message Authentication Codes), hashing and various helper functions.

Code was tested according to examples, provided in СТБ 34.101.31-2011 (CBC mode was slightly altered in the later edition, so the testcases were changed to fit).
