## How Pointer Works?

```go
    
   // This two variable has a type which is 'int', Has a value and also has a memory address 
   i, j := 34,45

   // 'p' is now a pointer which holds the memory address of i and 'p' also has a type which is  ' *int '
   p := &i


  // ' *p ' is for dereferencing. *p gives you the value that p is pointing to
   fmt.Println(*p)
   

```

### Stack and Heap for Memory Allocation

> when you execute the main function a goroutine is created and each goroutine get a stack of memory

