Universal publishing json resource hashing
==========================================

For a variety of reasons, we with to store a hash of a json resource, typically for comparison with other hashes as a simple way to establish whether a resource is likely to have changed.

Here is a description of how JSON resources should be serialised for UP usage.

The JSON must be encoded as UTF8. No whitespace, newlines, or carriage returns should be included.  The properties should be sorted by alpha-numeric key.

This JSON must be hashed using the x64, 128 bit variant of MurmurHash3.

The resulting hash must be encoded as a hex string, with alpha characters in lower case.

Example inputs and hash outputs
-------------------------------

```
{}
```
```466e20057851c2d220882a78996617be```

```
    {      }
```
```466e20057851c2d220882a78996617be```

```
{   "uuid":  "f86e77a4-0dbd-4e13-98b5-c5c97c34611a" }
```
```b0720da3fb340cee79559f8bc9233d0d```

```
{"b":"b", "c":"c", "a":"a"}
```
```4a3ab9c9f7b795e09b6d2a822f5552ba```

```
{"x":1,
"y":2}
```
```4a2f0489c0b8a54e32804edf876a6df4```
