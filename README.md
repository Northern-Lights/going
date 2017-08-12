# Going

For those who write Go and need to get going quick!

## Introduction

This repository is a collection of packages designed to reduce the verbosity of everyday tasks.  For example, instead of having to set up an `osfile` and `err` block to read a file, `file.Read(string) ([]byte, error)` takes care of it in a single function call.

## Packages

This repository contains the packages described below.

### crypt

This package simplifies encrypting, decrypting, and hashing of data.  For example, cryptographic hashes, such as MD5 and SHA1, take in a simple `[]byte` and return a simple `[]byte` (as opposed to `[N]byte`, that is, a sized array).  This saves the writer the step of having to save an array as an addressable variable only to convert it to a slice immediately afterward.

### file

This package provides functions for easy reading, writing, and appending data from or to files.
