package main

import (
	"errors"
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"log"
	"os"
	"os/exec"
)

func main() {

	//se declara el modulo
	moduleHM := ir.NewModule()
	moduleHM.TargetTriple = "x86_64-pc-windows-msvc"
	moduleHM.DataLayout = "e-m:w-i64:64-f80:128-n8:16:32:64-S128"

	formatStr := moduleHM.NewGlobalDef("str", constant.NewCharArrayFromString("Resultado: %d\n\x00"))
	//aquiStr := moduleHM.NewGlobalDef("aquiStr", constant.NewCharArrayFromString("Aqui\n\x00"))

	//puts := moduleHM.NewFunc("puts", types.I32, ir.NewParam("", types.I8Ptr))
	printfInt := moduleHM.NewFunc("printf", types.I32, ir.NewParam("", types.I8Ptr), ir.NewParam("", types.I32))

	//funcion para multiplicar
	funcMult := moduleHM.NewFunc("mult", types.I32, ir.NewParam("x", types.I32), ir.NewParam("y", types.I32))

	//se declaran los bloques de la funcion
	multEntryBlock := funcMult.NewBlock("multEntry")
	multEndBlock := funcMult.NewBlock("multEnd")

	multResult := multEntryBlock.NewAlloca(types.I32)
	multTemp := multEntryBlock.NewMul(funcMult.Params[0], funcMult.Params[1])
	multEntryBlock.NewStore(multTemp, multResult)
	multEntryBlock.NewLoad(types.I32, multResult)
	multEntryBlock.NewBr(multEndBlock)
	multEndBlock.NewRet(multTemp)

	/*
		//bloque de entrada
		multEntryBlock.NewBr(multIFBodyBlock)
		result := multEntryBlock.NewAlloca(types.I32)
		multEntryBlock.NewStore(constant.NewInt(types.I32, 0), result)
		result3 := multEntryBlock.NewLoad(types.I32, result)
		multEntryBlock.NewCall(printfInt, formatStr, funcMult.Params[0])
		multEntryBlock.NewCall(printfInt, formatStr, funcMult.Params[1])

		//bloque de if
		condIF1Mult := multIFBodyBlock.NewICmp(enum.IPredEQ, funcMult.Params[0], constant.NewInt(types.I32, 0))
		condIF2Mult := multIFBodyBlock.NewICmp(enum.IPredEQ, funcMult.Params[1], constant.NewInt(types.I32, 0))
		condIFMult := multIFBodyBlock.NewOr(condIF1Mult, condIF2Mult)
		multIFBodyBlock.NewCondBr(condIFMult, multEndBlock, multElseBlock)

		//se retorna 0 si alguno de los valores es 0
		//multEndBlock.NewRet(constant.NewInt(types.I32, 0))
		//multEndBlock.NewCall(printfInt, formatStr, result3)
		multEndBlock.NewRet(result3)

		//bloque de else
		multResult := multElseBlock.NewAlloca(types.I32)
		multElseBlock.NewStore(constant.NewInt(types.I32, 0), multResult)
		//temp2 := multElseBlock.NewLoad(types.I32, multResult)
		//multElseBlock.NewCall(printfInt, formatStr, temp2)
		multElseBlock.NewBr(multForCondBlock)

		//cuerpo de la condicion del bloque de for
		i := multForCondBlock.NewAlloca(types.I32)
		multForCondBlock.NewStore(constant.NewInt(types.I32, 0), i)
		YMult := multForCondBlock.NewAlloca(types.I32)
		multForCondBlock.NewStore(funcMult.Params[1], YMult)
		valIMult := multForCondBlock.NewLoad(types.I32, i)
		valYMult := multForCondBlock.NewLoad(types.I32, YMult)
		//multForCondBlock.NewCall(printfInt, formatStr, valIMult)
		//multForCondBlock.NewCall(printfInt, formatStr, valYMult)
		condForMult := multForCondBlock.NewICmp(enum.IPredSLT, valIMult, valYMult)
		multForCondBlock.NewCondBr(condForMult, multForBodyBlock, multEndBlock)

		//cuerpo del bucle
		tempResultMult := multForBodyBlock.NewMul(funcMult.Params[0], funcMult.Params[1])
		resultMult := multForBodyBlock.NewMul(tempResultMult, valIMult)
		multForBodyBlock.NewStore(resultMult, multResult)

		//se incrementa i
		multIncrI := multForBodyBlock.NewAdd(valIMult, constant.NewInt(types.I32, 1))
		multForBodyBlock.NewStore(multIncrI, i)
		multForBodyBlock.NewBr(multForCondBlock)

		finalResultMult := multEndBlock.NewLoad(types.I32, multResult)
		multEndBlock.NewRet(finalResultMult)
	*/
	/*






	 */

	//se declara la funcion 2 para el modulo con sus respectivos bloques, variables y operaciones
	funcSub := moduleHM.NewFunc("sub", types.Float, ir.NewParam("x", types.Float))

	//se declaran los bloques de la funcion
	subEntryBlock := funcSub.NewBlock("subEntry")
	subIFBodyBlock := funcSub.NewBlock("subIFBody")
	subElseBlock := funcSub.NewBlock("subElseBody")
	subForBodyBlock := funcSub.NewBlock("subForBody")
	subForCondBlock := funcSub.NewBlock("subForCond")
	subEndBlock := funcSub.NewBlock("subEnd")

	//bloque de entrada
	resultSub := subEntryBlock.NewAlloca(types.I32)
	subEntryBlock.NewStore(funcSub.Params[0], resultSub)
	iSub := subEntryBlock.NewAlloca(types.I32)
	subEntryBlock.NewStore(constant.NewInt(types.I32, 0), iSub)
	XSub := subForCondBlock.NewAlloca(types.I32)
	subForCondBlock.NewStore(funcSub.Params[0], XSub)
	//subEntryBlock.NewCall(printfInt, formatStr, funcSub.Params[0])
	subEntryBlock.NewBr(subIFBodyBlock)

	//bloque de if
	subX := subIFBodyBlock.NewAlloca(types.I32)
	subIFBodyBlock.NewStore(funcSub.Params[0], subX)
	condIFSub := subIFBodyBlock.NewICmp(enum.IPredSGT, subIFBodyBlock.NewLoad(types.I32, subX), constant.NewInt(types.I32, 0))
	subIFBodyBlock.NewCondBr(condIFSub, subForCondBlock, subElseBlock)

	//bloque de else
	subElseBlock.NewBr(subEndBlock)

	//cuerpo de la condicion del bloque de for
	subForCondBlock.NewBr(subForBodyBlock)

	//bloque del cuerpo del for
	//subForBodyBlock.NewCall(printfInt, formatStr, valISub)
	//subForBodyBlock.NewCall(printfInt, formatStr, valXSub)
	condForSub := subForBodyBlock.NewICmp(enum.IPredSLT, subForBodyBlock.NewLoad(types.I32, iSub), subForBodyBlock.NewLoad(types.I32, XSub))
	subForBodyBlock.NewCondBr(condForSub, subForBodyBlock, subEndBlock)

	//cuerpo del bucle
	tempISub := subForBodyBlock.NewLoad(types.I32, iSub)
	//subForBodyBlock.NewCall(printfInt, formatStr, tempISub)
	divSub := subForBodyBlock.NewUDiv(tempISub, constant.NewInt(types.I32, 2))
	//subForBodyBlock.NewCall(printfInt, formatStr, divSub)
	//subForBodyBlock.NewCall(printfInt, formatStr, funcSub.Params[0])
	resultSubTemp := subForBodyBlock.NewCall(funcMult, divSub, funcSub.Params[0])
	//subForBodyBlock.NewCall(printfInt, formatStr, resultSubTemp)
	subForBodyBlock.NewStore(resultSubTemp, resultSub)
	//subForBodyBlock.NewCall(printfInt, formatStr, resultSub)

	//se incrementa i
	subIncrI := subForBodyBlock.NewAdd(tempISub, constant.NewInt(types.I32, 1))
	//subForBodyBlock.NewCall(printfInt, formatStr, subIncrI)
	subForBodyBlock.NewStore(subIncrI, iSub)
	//temp := subForBodyBlock.NewLoad(types.I32, iSub)
	//subForBodyBlock.NewCall(printfInt, formatStr, temp)
	subForBodyBlock.NewBr(subForBodyBlock)

	//bloque final
	subEndBlock.NewBr(subEndBlock)
	subEndBlock.NewCall(printfInt, formatStr, resultSub)
	returnResult := subEndBlock.NewLoad(types.I32, resultSub)
	//subEndBlock.NewCall(printfInt, formatStr, returnResult)
	subEndBlock.NewRet(returnResult)

	//main
	funcMain := moduleHM.NewFunc("main", types.I32)
	mainBlock := funcMain.NewBlock("")
	intValue := mainBlock.NewCall(funcSub, constant.NewInt(types.I32, 5))
	mainBlock.NewCall(printfInt, formatStr, intValue)
	mainBlock.NewRet(intValue)

	fmt.Println(moduleHM.String())

	f, err := os.Create("module.ll")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer f.Close()
	if _, err := moduleHM.WriteTo(f); err != nil {
		fmt.Println("Error al escribir el módulo:", err)
		return
	}

	cmd := exec.Command("clang", "", "module.ll", "-o", "module.exe")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Error al compilar el módulo:", err)
		return
	}

	fmt.Println("El archivo ejecutable .exe ha sido generado correctamente.")

	cmd = exec.Command("module.exe")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		fmt.Println("Error al ejecutar el comando:", err)
		return
	}
}
