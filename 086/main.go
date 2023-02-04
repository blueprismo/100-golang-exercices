// PROTOBUFFS!
// What is a protocol buffer?
// Please take your time to read, soak and understand a bit what is protobuff
// https://developers.google.com/protocol-buffers
// https://en.wikipedia.org/wiki/Protocol_Buffers
// https://github.com/protocolbuffers/protobuf
// https://protobuf.dev/
// Wow! Much doc!

// Once you have read this, we are going to begin this exercise. The goal will be declaring a `.proto` file

// 1- Install go dependencies for protocol buffers:
// go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

// 2- Set your Gopath bin into your path like this:
// export PATH="$PATH:$(go env GOPATH)/bin"
// 2.1- Ensure you have protoc-gen-go by issuing:
// protoc-gen-go --version

// 3- Create a message inside the "person.proto" file

// 4- That message will be called Person, and it will have 2 fields: 
// 4.1- "name" a string with the tag number '1'
// 4.2- "age" an int32 with the tag number '2'

// That's it! This is what you need to complete this exercise.