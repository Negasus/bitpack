# Bitpack 
Simple library for pack slice of **unique** Int values to bitmask and unpack back

Example: 
```
    buf := make([]uint64, 0)

    // Base 10
    buf = PackBase(10, []int{1, 11, 21}, buf)
    
    fmt.Printf("%d", buf)
    
    // Result decimal (three numbers): 1 1 1
    //
    // Binary representation (three numbers by 10 bits)
    //         21         11         01 bits
    // 0000000001 0000000001 0000000001 

    // Base 64
    buf = PackBase(64, []int{1, 11, 21}, buf)
    
    fmt.Printf("%d", buf)
    
    // Result decimal: 1049601
    //
    // Binary representation (one number up to 64 bit)
    //         21        11        01 bits
    // 000000000100000000010000000001 
```