def resta(x):
    resultado = x
    i = 0
    if x > 0:
        for i in range(x):
            sub = i-10
            resultado = multiplicacion(sub, x)
            ##print("resultado: ", resultado)
        print("resultado: ", resultado)
    else:
        print("resultado: ", resultado)

def multiplicacion(x, y):
    resultado = 0
    resultado = x * y
    return resultado

resta(5)