# DEchan
Decentralized Chan-like board system using the DERO blockchain.

## Usage
To use a DEchan board, one must have the SCID of the board, as well as a client software to read, parse, and format messages correctly.

## Standard
A valid DEchan board Smart Contract contains a few things. 
First, is a variable named "IsBoard" with the value "True". This variable of course just tells the client that this is supposed to be a valid DEchan board.
Secondly, the messages need to be formatted correctly. A correct key for a message is as follows: `TIMESTAMP:rREPLY`, where the `TIMESTAMP` is the output of `BLOCK_TIMESTAMP()` and `REPLY` is either 0 for a thread, or the `TIMESTAMP` of a thread for a reply to that thread.
