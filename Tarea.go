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
	formatStrMult := moduleHM.NewGlobalDef("strMult", constant.NewCharArrayFromString("Resultado multiplicacion: %d\n\x00"))
	//aquiStr := moduleHM.NewGlobalDef("aquiStr", constant.NewCharArrayFromString("Aqui\n\x00"))

	//puts := moduleHM.NewFunc("puts", types.I32, ir.NewParam("", types.I8Ptr))
	printfInt := moduleHM.NewFunc("printf", types.I32, ir.NewParam("", types.I8Ptr), ir.NewParam("", types.I32))

	//funcion para multiplicar
	funcMult := moduleHM.NewFunc("mult", types.I32,
		ir.NewParam("x", types.I32),
		ir.NewParam("y", types.I32),
	)

	//se declaran los bloques de la funcion
	multEntryBlock := funcMult.NewBlock("multEntry")
	multEndBlock := funcMult.NewBlock("multEnd")

	//bloque de entrada
	multTemp := multEntryBlock.NewMul(funcMult.Params[0], funcMult.Params[1])
	multResult := multEntryBlock.NewAlloca(types.I32)
	multEntryBlock.NewStore(multTemp, multResult)
	multFinal := multEntryBlock.NewLoad(types.I32, multResult)

	multEntryBlock.NewCall(printfInt, formatStrMult, multFinal)

	multEntryBlock.NewBr(multEndBlock)
	multEndBlock.NewRet(multFinal)

	//se declara la funcion 2 para el modulo con sus respectivos bloques, variables y operaciones
	funcSub := moduleHM.NewFunc("sub", types.I32,
		ir.NewParam("x", types.I32))

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

	subX := subEntryBlock.NewAlloca(types.I32)
	subEntryBlock.NewStore(funcSub.Params[0], subX)
	subIFX := subEntryBlock.NewLoad(types.I32, subX)

	subEntryBlock.NewBr(subIFBodyBlock)

	//bloque de if
	condIFSub := subIFBodyBlock.NewICmp(enum.IPredSGT, subIFX, constant.NewInt(types.I32, 0))
	subIFBodyBlock.NewCondBr(condIFSub, subForCondBlock, subElseBlock)

	//bloque de else'
	subElseBlock.NewBr(subEndBlock)

	//bloque del cuerpo del for
	subParam1 := subForCondBlock.NewLoad(types.I32, iSub)
	subParam2 := subForCondBlock.NewLoad(types.I32, subX)
	condForSub := subForCondBlock.NewICmp(enum.IPredSLT, subParam1, subParam2)
	subForCondBlock.NewCondBr(condForSub, subForBodyBlock, subEndBlock)

	//cuerpo del bucle
	tempISub := subForBodyBlock.NewLoad(types.I32, iSub)
	sub := subForBodyBlock.NewSub(tempISub, constant.NewInt(types.I32, 10))
	resultSubTemp := subForBodyBlock.NewCall(funcMult, sub, funcSub.Params[0])
	subForBodyBlock.NewStore(resultSubTemp, resultSub)

	//se incrementa i
	subIncrI := subForBodyBlock.NewAdd(tempISub, constant.NewInt(types.I32, 1))
	subForBodyBlock.NewStore(subIncrI, iSub)
	subForBodyBlock.NewBr(subForCondBlock)

	//bloque final
	subEndBlock.NewBr(subEndBlock)
	returnResult := subEndBlock.NewLoad(types.I32, resultSub)
	subEndBlock.NewCall(printfInt, formatStr, returnResult)
	subEndBlock.NewRet(returnResult)

	//main
	funcMain := moduleHM.NewFunc("main", types.I32)
	mainBlock := funcMain.NewBlock("")
	intValue := mainBlock.NewCall(funcSub, constant.NewInt(types.I32, 5))
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
