#Chapter 2 Exercises

1)How are integers stored on a computer?
**They're stored based on the request by the programmer as either a 32 bit binary value or a 64 bit value.  The Most Sig Bit is reserved for the sign of the number(whether or not it's positive)**
2)We know that (in base 10) the largest one-digit number is 9 and the largest two-digit number is 99.  Given that in binary the largest two-digit number is 11(3), the largest three-digit number is 111(7) and the largest four-digit number is 1111(15), what's the largest eight-digit number?(Hint:10<sup>1</sup>-1 = 9 and 10<sup>2</sup>-1 = 99)
**This was such a basic question that I thought it was a trick.  The answer is, of course, 11111111.  However, I double checked with some python**
>>> bin((2**8)-1)
'0b11111111'
>>> 
3)Although overpowered for the task, you can use Go as a calculator.  Write a program that computes 32,132 x 42,452 and prints it to the terminatl(use the * operator for multiplication).
**see [Exercise3.go](./Exercise3.go)**
4)What is a string?  How do you find its length?
**A string is an address in memory which contains text.  The address in memory(usually in the .data region of the executable) begins with the first letter of the text and ends with a 0 character.  If the user wishes to include a zero-character within the string, this nulchar or termchar must be tokenized. One can find a string's length by calling the _len()_ function.** 
5)What's the value of the expression (true && false) || (false && true) || !(false && false)?
**true.  Due to the last condition being reversed with an exclamation**

