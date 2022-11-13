package main

// import the "testing" library
import (
    "testing"
    "encoding/hex"
    "bytes"
)

// Again, like the Test or Benchmark signature, go also accepts Fuzzing (https://en.wikipedia.org/wiki/Fuzzing)
// The syntax of the Fuzzer function  name will be FuzzXxxxxx and the signature (f *testing.F)
// Create a fuzzer function named FuzzHex
func FuzzHex(f *testing.F) {
    // now we will seed the fuzzer with the function *f.Add() and the seed as the parameter
    for _, seed := range [][]byte{{}, {0}, {9}, {0xa}, {0xf}, {1, 2, 3, 4}} {
      //here!
      
    }

    // now we will call the *f.Fuzz() function, this is considered the target. 
    // documentation: A fuzz target must accept a *T parameter, followed by one or more parameters for random inputs. 
    f.Fuzz(func(t *testing.T, in []byte) {
      // here we will use enc (hex.EncodeToString(in))
      
      // and here out := hex.DecodeString(enc)
      

      // and now... fuzz!
      if err != nil {
        t.Fatalf("%v: decode: %v", in, err)
      }
      if !bytes.Equal(in, out) {
        t.Fatalf("%v: not equal after round trip: %v", in, out)
      }
    })
}