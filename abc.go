package main

import "fmt"

func main() {
   /* local variable definition */
   var inputi int
   var istVal int
   var secVal int
   var ret int

   /* calling a function to get max value */
  fmt.Printf( "Enter Function Please : \n");
  fmt.Printf( "1. Maximum Function \n");
  fmt.Printf( "2. Minimum Function \n");
  fmt.Printf( "3. Sum Function \n");
  fmt.Printf( "4. Multiply Function \n");

  fmt.Scan(&inputi)

  fmt.Printf( "Enter First Value   :  \n");
  fmt.Scan(&istVal)
  fmt.Printf( "Enter Second Value : \n");
  fmt.Scan(&secVal)

  if(inputi==1){
    ret = max(istVal,secVal);
  }
  if(inputi==2){
    ret = min(istVal,secVal);
  }
  if(inputi==3){
    ret = sum(istVal,secVal);
  }
  if(inputi==4){
    ret = multiply(istVal,secVal);
  }

   fmt.Printf( "Max value is : %d\n", ret )
}

/* function returning the max between two numbers */
func max(num1, num2 int) int {
   /* local variable declaration */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}

func min(num1, num2 int) int {
   /* local variable declaration */
   var result int

   if (num1 < num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}

func sum(num1, num2 int) int {
   /* local variable declaration */
   var result int
   result = num1 + num2;
   return result
}

func multiply(num1, num2 int) int {
   /* local variable declaration */
   var result int
   result = num1*num2;
   return result
}
