# XDR for Stellar developers

This article offers an introduction to `XDR`, which is a standard for
describing and encoding data. The goal of this article is to explore
the standard while analyzing how the XDR golang library works. Any
kind of developer can benefit from reading this article, but it will
be oriented towards Stellar developers. By the end of the article,
you'll be familiar with `XDR`, the golang `XDR` library and understand
how `XDR` is used in the Stellar protocol (ecosystem)?.

## History

Chances are you have never heard of `XDR` and you might even be
surprised that there is something different from `JSON` to send
messages between computers!

While `JSON` is today the predominant data format for sending messages
around the web, before that, there was `XML` (do not confuse with
`XLM`) and even before that, people were using other formats to send
messages between machines.

When you are sending messages, you need a way to describe the
information you are sending. Without any context, what are you
supposed to do if you receive the following bytes?

```golang Sequence of bytes in hex format
bytes := [0x0, 0x0, 0x0, 0xb, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x0]
```

In the mid 1980s Sun Microsystems came up with a proposal which would
allow you not only to send data in `bytes` but also describe the shape
of that data, this format was called External Data Representation
Standard or `XDR`.

Following the `XDR` spec, the `bytes` above represent the string
`hello world`, where the first `4 bytes` tells you the number of
`bytes` that a string might contain (`11` in this case), and the the
rest are the UTF-8 bytes representantion of the string. [^1]

XDR is used by a variety of systems like Stellar Payment Network, ZFS
File System and The SpiderMonkey JavaScript engine.

`XDR` can be compared to Google's [protocol buffers](https://developers.google.com/protocol-buffers/).

In the following chapters you'll see a visual representation of an XDR
structure, learn about the basic data types and then explore each data
type while playing with the golang library.

## Stellar and XDR

>XDR, also known as External Data Representation, is used throughout
>the Stellar network and protocol. The ledger, transactions, results,
>history, and even the messages passed between computers running
>stellar-core are encoded using XDR.
>-https://www.stellar.org/developers/guides/concepts/xdr.html

You can develop applications on top of Stellar without worrying about
how XDR works, the SDKs includes helpers to help you interact with the
network and convert data back and forth in XDR.

The page where the the quote above was taken, list some reasons on why XDR was chosen, like:

 - It is very compact, so it can be transmitted quickly and stored with minimal disk space.
 - Data encoded in XDR is reliably and predictably stored. Fields are always in the same order, which makes cryptographically signing and verifying XDR messages simple.
 - XDR definitions include rich descriptions of data types and structures, which is not possible in simpler formats like JSON, TOML, or YAML.

Nicolas Barry also expands on the subject in the following stack
 overflow question where someone asked ["Why did the project settle on
 XDR for external data
 serialisation?"](https://stellar.stackexchange.com/a/284/1066). It
 mentions similar points, but adds emphasis on the extra security
 added by using the protocol, specifically as the protocol evolves.

## Representing data with XDR

```xdr
struct PaymentOp
{
    AccountID destination; // recipient of the payment
    Asset asset;           // what they end up with
    int64 amount;          // amount they end up with
};
```

## Assumptions

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
minim veniam, quis nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat. Duis aute irure dolor in
reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.

## Data types

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
minim veniam, quis nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat. Duis aute irure dolor in
reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.


## Reimplementing go-xdr
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
minim veniam, quis nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat. Duis aute irure dolor in
reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.

## Integer

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
minim veniam, quis nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat. Duis aute irure dolor in
reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.

```go
  func (enc *Encoder) EncodeInt(v int32) (int, error) {
          var b [4]byte
          b[0] = byte(v >> 24)
          b[1] = byte(v >> 16)
          b[2] = byte(v >> 8)
          b[3] = byte(v)

          n, err := enc.w.Write(b[:])
          if err != nil {
                  msg := fmt.Sprintf(errIOEncode, err.Error(), 4)
                  err := marshalError("EncodeInt", ErrIO, msg, b[:n], err)
                  return n, err
          }

          return n, nil
  }
```

## Unsigned Integer

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
minim veniam, quis nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat. Duis aute irure dolor in
reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.


## Footnotes

- [^1] The spec talks about ASCII bytes but current implementations use UTF-8 which has replaced ASCII.
