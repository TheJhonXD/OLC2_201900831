var arr: [Int] = [8, 4, 6, 2]
    var arr2: [Int] = [40, 21, 1, 3, 14, 4]
    var arr3: [Int] = [90, 3, 40, 10, 8, 5]

    // Función para intercambiar dos elementos en un arreglo
    func intercambiar(_ name: String, _ i: Int, _ j: Int) {
        if name == "arr"{
            let aux = arr[i]
            arr[i] = arr[j]
            arr[j] = aux
        }else if name == "arr2"{
            let aux = arr2[i]
            arr2[i] = arr2[j]
            arr2[j] = aux
        }else if name == "arr3"{
            let aux = arr3[i]
            arr3[i] = arr3[j]
            arr3[j] = aux
        }
    }

    // Algoritmo de ordenamiento por intercambio (Bubble Sort)
    func ordIntercambio(_ name: String) {
        var i = 0
        var j = 0
        
        while i < (arr.count - 1) {
            j = i + 1
            if name == "arr" {
                while j < arr.count {
                    if arr[i] > arr[j] {
                        intercambiar("arr", i, j)
                    }
                    j += 1
                }
            }else if name == "arr2" {
                while j < arr2.count {
                    if arr2[i] > arr2[j] {
                        intercambiar("arr2", i, j)
                    }
                    j += 1
                }
            }else if name == "arr3" {
                while j < arr3.count {
                    if arr3[i] > arr3[j] {
                        intercambiar("arr3", i, j)
                    }
                    j += 1
                }
            }
            i += 1
        }
    }

    // Algoritmo de ordenamiento por selección (Selection Sort)
    func ordSeleccion(_ name: String) {
        var i = 0
        var j = 0
        var indiceMenor = 0
        var n = arr.count
        var nombre = "arr"

        if name == "arr2" {
            n = arr2.count
        }else if name == "arr3" {
            n = arr3.count
        }
        
        while i < (n - 1) {
            indiceMenor = i
            j = i + 1
            while j < n {
                if name == "arr" {
                    if arr[j] < arr[indiceMenor] {
                        indiceMenor = j
                    }
                    nombre = "arr"
                }else if name == "arr2" {
                    if arr2[j] < arr2[indiceMenor] {
                        indiceMenor = j
                    }
                    nombre = "arr2"
                }else if name == "arr3" {
                    if arr3[j] < arr3[indiceMenor] {
                        indiceMenor = j
                    }
                    nombre = "arr3"
                }
                
                j += 1
            }
            
            if i != indiceMenor {
                intercambiar(nombre, i, indiceMenor)
            }
            i += 1
        }
    }

    // Función para imprimir un arreglo
    func printArray(_ msg: String, _ name: String) {
        var out = ".["
        if name == "arr"{
            for i in 0...arr.count-1 {
                if i == arr.count - 1 {
                    out += String(arr[i])
                } else {
                    out += String(arr[i]) + ", "
                }
            }
        }else if name == "arr2" {
            for i in 0...arr2.count-1 {
                if i == arr2.count - 1 {
                    out += String(arr2[i])
                } else {
                    out += String(arr2[i]) + ", "
                }
            }
        }else if name == "arr3" {
            for i in 0...arr3.count-1 {
                if i == arr3.count - 1 {
                    out += String(arr3[i])
                } else {
                    out += String(arr3[i]) + ", "
                }
            }
        }
        out += "]."
        print(msg + out)
    }
    print("=============================================")
    print("===============ORDENAMIENTO==================")
    print("=============================================")
    print("INTERCAMBIO:")
    printArray("entrada: ", "arr")
    ordIntercambio("arr")
    printArray("salida: ", "arr")
    print("SELECCIÓN:")
    printArray("entrada: ", "arr2")
    ordSeleccion("arr2")
    printArray("salida: ", "arr2")

    print("=============================================")
    print("================FUNCIONES====================")
    print("=============================================")

    printArray("arr1: ", "arr")
    arr.append(9)
    printArray("arr1: ", "arr")

    printArray("arr2: ", "arr2")
    arr2.removeLast()
    printArray("arr2: ", "arr2")

    printArray("arr2: ", "arr3")
    arr3.remove(3)
    printArray("arr2: ", "arr3")

    print("arr1 vacío: ", arr.isEmpty)
    print("arr1 cantidad: ", arr.count)


    /*
    =============================================
    ===============ORDENAMIENTO==================
    =============================================
    INTERCAMBIO:
    entrada: .[8, 4, 6, 2].
    salida: .[2, 4, 6, 8].
    SELECCIÓN:
    entrada: .[40, 21, 1, 3, 14, 4].
    salida: .[1, 3, 4, 14, 21, 40].
    =============================================
    ================FUNCIONES====================
    =============================================
    arr1: .[2, 4, 6, 8].
    arr1: .[2, 4, 6, 8, 9].
    arr2: .[1, 3, 4, 14, 21, 40].
    arr2: .[1, 3, 4, 14, 21].
    arr2: .[90, 3, 40, 10, 8, 5].
    arr2: .[90, 3, 40, 8, 5].
    arr1 vacío:  false
    arr1 cantidad:  5
    */
