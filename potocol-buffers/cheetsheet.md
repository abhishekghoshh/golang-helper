# cheetsheet for Protocol buffer


#### This is a sample proto file
first we havr to mention the syntax then we will create a proto messge, which is similar to struct </br>
a message contains fields, which each field consists of [ field type , field name = tag; ] </br>
```
syntax = "proto3";

message Account{
    uint32 id = 1;
    string name = 2;
    bool is_verified = 3;
}
```

when we do not set any field then the field will not be serialized </br>
we will not get any payload overhead for any empty data, though we can use some default value </br>