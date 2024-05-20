def resta(x):
    resultado = x
    i = 0
    if x > 0:
        for i in range(x):
            div = i/2
            resultado = multiplicacion(div, x)
            print("resultado: ", resultado)
    else:
        print("resultado: ", resultado)

def multiplicacion(x, y):
    resultado = 0
    resultado = x * y
    return resultado

resta(5)