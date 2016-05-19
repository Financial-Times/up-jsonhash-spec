## Universal publishing json resource hashing

For a variety of reasons, we with to store a hash of a json resource, typically for comparison with other hashes as a simple way to establish whether a resource is likely to have changed.

Here is a description of how JSON resources should be serialised for UP usage.

The JSON must be endoced as UTF8. No whitespace, newlines, or carriage returns should be included.  The properties should be sorted by alpha-numeric key.

This JSON must be hashed using the 128 bit variant of MurmurHash3.

The resulting hash must be encoded as a hex string, with alpha characters in lower case.
