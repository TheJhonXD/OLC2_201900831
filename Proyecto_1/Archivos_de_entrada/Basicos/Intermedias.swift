var a = 909

print("=======================================================================")
print("==================================IF===================================")
print("=======================================================================")

if a > 50 {
    print("IF CORRECTO")
} else if a == 56 {
    print("IF INCORRECTO")
} else {
    print("IF INCORRECTO")
}

print("")
print("=======================================================================")
print("=============================IFs ANIDADOS==============================")
print("=======================================================================")
var aux = 10
if aux > 0 {
    print("PRIMER IF CORRECTO")
    if true && (aux == 1) {
        print("SEGUNDO IF INCORRECTO")
    } else if aux > 10 {
        print("SEGUNDO IF INCORRECTO")
    } else {
        print("SEGUNDO IF CORRECTO")
    }
} else if aux <= 3 {
    print("PRIMER IF INCORRECTO")
    if true && (aux == 1) {
        print("SEGUNDO IF INCORRECTO")
    } else if aux > 10 {
        print("SEGUNDO IF INCORRECTO")
    } else {
        print("SEGUNDO IF CORRECTO")
    }
} else if aux == a {
    print("PRIMER IF INCORRECTO")
    if true && (aux == 1) {
        print("SEGUNDO IF INCORRECTO")
    } else if aux > 10 {
        print("SEGUNDO IF INCORRECTO")
    } else {
        print("SEGUNDO IF CORRECTO")
    }
}

print("")
print("=======================================================================")
print("=================================WHILE=================================")
print("=======================================================================")
var index = 0
while index >= 0 {
    if index == 0 {
        index = index + 100
    } else if index > 50 {
        index = index / 2 - 25
    } else {
        index = (index / 2) - 1
    }
    print(index)
}

print("")
print("=======================================================================")
print("==============================WHILE-2===================================")
print("=======================================================================")
index = -2
index = index + 1
while index != 12 {
    index = index + 1
    if index == 0 || index == 1 || index == 11 || index == 12 {
        print("*********************************************************************************************************")
    } else if index == 2 {
        print("**********  ***************  ******                 ******                 ******              **********")
    } else if index >= 3 && index <= 5 {
        print("**********  ***************  ******  *********************  *************  ******  **********************")
    } else if index == 6 {
        print("**********  ***************  ******                 ******                 ******  **********************")
    } else if index >= 7 && index <= 9 {
        print("**********  ***************  ********************   ******  *************  ******  **********************")
    } else if index == 10 {
        print("**********                   ******                 ******  *************  ******              **********")
    }
}

print("")
print("=======================================================================")
print("=============================TRANSFERENCIA=============================")
print("=======================================================================")
a = -1
while a < 5 {
    a = a + 1
    if a == 3 {
        print("a")
        continue
    } else if a == 4 {
        print("b")
        break
    }
    print("El valor de a es:", a)
}

print("Se debió imprimir")

print("")
print("=======================================================================")
print("==================================FOR===================================")
print("=======================================================================")
for i in 0...9 {
    var output = ""
    output = ""
    for j in 0...(10 - i) {
        output = output + " "
    }
    for k in 0...i {
        output = output + "* "
    }
    print(output)
}

/*
=======================================================================
==================================IF===================================
=======================================================================
IF CORRECTO

=======================================================================
=============================IFs ANIDADOS==============================
=======================================================================
PRIMER IF CORRECTO
SEGUNDO IF CORRECTO

=======================================================================
=================================WHILE=================================
=======================================================================
100
25
11
4
1
-1

=======================================================================
==============================WHILE-2===================================
=======================================================================
*********************************************************************************************************
*********************************************************************************************************
**********  ***************  ******                 ******                 ******              **********
**********  ***************  ******  *********************  *************  ******  **********************
**********  ***************  ******  *********************  *************  ******  **********************
**********  ***************  ******  *********************  *************  ******  **********************
**********  ***************  ******                 ******                 ******  **********************
**********  ***************  ********************   ******  *************  ******  **********************
**********  ***************  ********************   ******  *************  ******  **********************
**********  ***************  ********************   ******  *************  ******  **********************
**********                   ******                 ******  *************  ******              **********
*********************************************************************************************************
*********************************************************************************************************

=======================================================================
=============================TRANSFERENCIA=============================
=======================================================================
El valor de a es: 0
El valor de a es: 1
El valor de a es: 2
a
b
Se debió imprimir

=======================================================================
==================================FOR===================================
=======================================================================
           * 
          * * 
         * * * 
        * * * * 
       * * * * * 
      * * * * * * 
     * * * * * * * 
    * * * * * * * * 
   * * * * * * * * * 
  * * * * * * * * * * 
*/