# Unhandled Error in Go Calculation

This repository demonstrates a common error in Go programs: failing to properly handle errors returned by functions.

The `bug.go` file contains code that attempts to divide by zero. While the `calculate` function correctly returns an error in this case, the `main` function fails to check for and handle this error, potentially leading to unexpected behavior or crashes.

The `bugSolution.go` file shows how to correctly handle the error, using a standard `if err != nil` check.