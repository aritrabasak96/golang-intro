package main

import ("fmt")

func myPointer(){

  // if you do not initialize a pointer to a value
  // then the pointer variable is always be nil
  var check *int;
  fmt.Println("check-----",check);  // it is always nil

  // create pointers way 1
  b := 10;
  a := &b;  // hold the memory address of b

  // fmt.Println(a) => 0xc00000a0c8
  // fmt.Println(*a) => 10    (*a => dreference)

  fmt.Println(*a,b);

  // create pointers way 2
  var val int = 20;
  var point_val *int = &val;

  fmt.Println(val,*point_val);

  // if you want to assign some value
  *point_val = 25;

  // the both val and point_val will reflected
  fmt.Println(val,*point_val);  // o/p => 25 25

}

type myStruct struct{
  foo int
}

func myStructPointer()  {

   var my myStruct
   my = myStruct{foo:23}
   fmt.Println(my);

}

func main()  {

  myPointer()
  myStructPointer()
}
