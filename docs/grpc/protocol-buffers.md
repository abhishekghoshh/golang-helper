# Protocol Buffers (Protobuf)

Protocol Buffers is a language-neutral, platform-neutral mechanism for serializing structured data — smaller, faster, and simpler than XML/JSON.

## Setup

Install `protoc` compiler and Go plugins:

```bash
brew install protobuf

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export PATH="$PATH:$(go env GOPATH)/bin"
```

Module: `proto-example` using `google.golang.org/protobuf v1.32.0`

## Compiling Proto Files

```bash
# Compile entire folder
protoc --go_out=./app/proto proto/*.proto

# Compile specific file
protoc --go_out=./app/proto proto/address.proto

# With go_package option in proto, no need to specify output path
protoc --go_out=. proto/*.proto
```

Makefile workflow:
```makefile
build: generate
	go build -o ${BIN} .
generate:
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. ${PROTO_DIR}/*.proto
```

## Proto File Structure

```protobuf
syntax = "proto3";

// Optionally specify output directory for each language
option go_package = "/app/proto";
option java_package = "/src/java";

// Messages are like structs
message FullAddress {
  string buildingName = 1;        // field type, name, tag number
  string streetName = 2;
  int32 pinCode = 8;
  optional string googleLocation = 9;  // optional field (no payload if empty)
  repeated string landMarks = 11;      // array
  map<string, string> otherInformations = 12; // map
}
```

**Rules:**
- Each field has a unique tag number (used in binary encoding)
- `optional` fields are not serialized when empty → no payload overhead
- `oneof` and `map` cannot be `repeated`
- `proto.Message` is the interface all generated proto structs implement

## Scalar Types

```protobuf
message ScalarTypes {
    int32   one = 1;    // 32-bit signed integer
    int64   two = 2;    // 64-bit signed integer
    uint32  three = 3;  // unsigned 32-bit
    uint64  four = 4;   // unsigned 64-bit
    sint32  five = 5;   // signed (zigzag encoding, efficient for negatives)
    sint64  six = 6;
    fixed32 seven = 7;  // always 4 bytes
    fixed64 eight = 8;  // always 8 bytes
    sfixed32 nine = 9;  // signed fixed
    sfixed64 ten = 10;
    double  eleven = 11;
    float   twelve = 12;
    bool    thirteen = 13;
    bytes   fourteen = 14;
    string  fifteen = 15;
}
```

## Enums

```protobuf
enum AccountType {
    ACCOUNT_TYPE_UNSPECIFIED = 0;  // required: first value must be 0
    ACCOUNT_TYPE_SAVINGS = 1;
    ACCOUNT_TYPE_CREDIT = 2;
    ACCOUNT_TYPE_BOTH = 3;
}

message Account {
    uint32 id = 1;
    string name = 2;
    bytes thumbnail = 3;
    bool is_verified = 4;
    float height = 5;
    repeated string phones = 6;
    AccountType acountType = 7;
}
```

## OneOf (Union/Sum Type)

Only one field within a `oneof` can be set at a time.

```protobuf
message Success { bytes mssg = 1; }
message Error   { bytes mssg = 1; repeated string stackTrace = 2; }

message Response {
    oneof response {
        Success success = 1;
        Error error = 2;
    }
    int32 code = 3;
    map<string, string> headers = 4;
}
```

**Go usage with type switch:**
```go
func IsSuccess(response *oneof.Response) (bool, error) {
    switch response.Response.(type) {
    case *oneof.Response_Success: return true, nil
    case *oneof.Response_Error:   return false, nil
    default:                      return false, errors.New("response is empty")
    }
}
```

## Nesting Messages

```protobuf
message OuterMessage {
    message InnerMessage { string name = 1; }
    enum InnerEnum {
        INNER_TYPE_UNSPECIFIED = 0;
        INNER_TYPE_FIRST = 1;
        INNER_TYPE_SECOND = 2;
    }
    InnerMessage student = 1;
    InnerEnum type = 2;
}
```

## Message Composition Patterns

### Compact (all in one file)
```protobuf
message City     { string name = 1; string zip_code = 2; string country_name = 3; }
message Street   { string name = 1; City city = 2; }
message Building { string name = 1; uint32 number = 2; Street street = 3; }
message Address  { City city = 1; Street street = 2; Building building = 3; }
```

### Segregated with Package
Each message in its own file, with a package declaration:
```protobuf
// city.proto
package city;
message City { string name = 1; string zip_code = 2; string country_name = 3; }

// street.proto
package street;
import "city.proto";
message Street { string name = 1; city.City city = 2; }

// building.proto
package building;
import "street.proto";
message Building { string name = 1; uint32 number = 2; street.Street street = 3; }

// address.proto
package address;
import "building.proto"; import "city.proto"; import "street.proto";
message Address { city.City city = 1; street.Street street = 2; building.Building building = 3; }
```

### Hierarchical (Nested, Top-Down)
```protobuf
message City {
    string name = 1; string zip_code = 2; string country_name = 3;
    message Street {
        string name = 1; City city = 2;
        message Building { string name = 1; uint32 number = 2; Street street = 3; }
    }
}
message Address {
    City city = 1;
    City.Street street = 2;
    City.Street.Building building = 3;
}
```

### Hierarchical (Nested, Bottom-Up)
```protobuf
message Building {
    message Street {
        message City { string name = 1; string zip_code = 2; string country_name = 3; }
        string name = 1; City city = 2;
    }
    string name = 1; uint32 number = 2; Street street = 3;
}
message Address {
    Building.Street.City city = 1;
    Building.Street street = 2;
    Building building = 3;
}
```

## Importing Proto Files

```protobuf
// packaging.proto
package example;
message PackagedMessage { string name = 1; }

// importing.proto
import "full_address.proto";
import "package_example.proto";

message Human { FullAddress address = 1; }                          // no-package import
message ImportedPackagedMessage {
    example.PackagedMessage imported_packaged_message = 1;            // with-package import
}
```

## protoc CLI Tools

```bash
# Decode raw binary to readable format
cat simple.bin | protoc --decode_raw

# Decode binary to text with schema
cat simple.bin | protoc --decode=Simple simple.proto > simple.txt

# Encode text format to binary
cat simple.txt | protoc --encode=Simple simple.proto > simple.pb
```

## Go: Working with Protobuf

### Creating Messages
```go
func NewFullAddress() *sample.FullAddress {
    googleLocation := "sample google location"
    return &sample.FullAddress{
        BuildingName:   "sample building",
        StreetName:     "sample street",
        PinCode:        123456,
        GoogleLocation: &googleLocation,          // optional field = pointer
        LandMarks:      []string{"landmark 1"},
        OtherInformations: map[string]string{"1": "sample 1"},
        AddressType:    sample.AddressType_ADDRESS_TYPE_URBAN,
    }
}
```

### Serialize to Binary File
```go
import "google.golang.org/protobuf/proto"

func WriteToFile(fname string, pb proto.Message) error {
    out, err := proto.Marshal(pb)
    if err != nil { return err }
    return os.WriteFile(fname, out, 0644)
}

func ReadFromFile(fname string, pb proto.Message) error {
    in, err := os.ReadFile(fname)
    if err != nil { return err }
    return proto.Unmarshal(in, pb)
}
```

### Serialize to/from JSON
```go
import "google.golang.org/protobuf/encoding/protojson"

func ToJSON(pb proto.Message) string {
    option := &protojson.MarshalOptions{Multiline: true}
    out, _ := option.Marshal(pb)
    return string(out)
}

func FromJSON(in string, pb proto.Message) {
    option := protojson.UnmarshalOptions{DiscardUnknown: true}
    option.Unmarshal([]byte(in), pb)
}
```

**Key point:** When reading from a file, pass an **empty** message instance — `proto.Unmarshal` populates it in place.

## Dockerfile for Docs

A `Dockerfile.docs` exists at the project root for building the MkDocs documentation site:

```dockerfile
# Check Dockerfile.docs at the project root for MkDocs build setup
```

## References

- [Official Proto 3 Guide](https://protobuf.dev/programming-guides/proto3/)
- [Style Guide](https://protobuf.dev/programming-guides/style/)
- [Well-Known Types](https://protobuf.dev/reference/protobuf/google.protobuf/#index)
- [All Proto Options](https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/descriptor.proto)
