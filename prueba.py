def resta(x):
    resultado = x
    i = 0
    if x > 0:
        for i in range(x):
            sub = i-10
            resultado = multiplicacion(sub, x)
            print("Resultado de cada multiplicacion: ", resultado)
        print("resultado final:", resultado)
    else:
        print("resultado final:", resultado)

def multiplicacion(x, y):
    resultado = 0
    resultado = x * y
    return resultado

resta(5)
resta(0)