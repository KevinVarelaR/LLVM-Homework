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

	//module declarated
	moduleHM := ir.NewModule()
	moduleHM.TargetTriple = "x86_64-pc-windows-msvc"
	moduleHM.DataLayout = "e-m:w-i64:64-f80:128-n8:16:32:64-S128"

	//global variables declarated
	formatStr := moduleHM.NewGlobalDef("str", constant.NewCharArrayFromString("Final result: %d\n\x00"))
	formatStrMult := moduleHM.NewGlobalDef("strMult", constant.NewCharArrayFromString("Multiplication result (Substraction result * sub func parameter value): %d\n\x00"))
	formatStrSub := moduleHM.NewGlobalDef("strSub", constant.NewCharArrayFromString("Substraction result (i-10): %d\n\x00"))
	formatStrI := moduleHM.NewGlobalDef("strI", constant.NewCharArrayFromString("i value: %d\n\x00"))
	printfInt := moduleHM.NewFunc("printf", types.I32, ir.NewParam("", types.I8Ptr), ir.NewParam("", types.I32))

	//multiplication function
	funcMult := moduleHM.NewFunc("mult", types.I32,
		ir.NewParam("x", types.I32),
		ir.NewParam("y", types.I32),
	)

	//multiplication func blocks declarated
	multEntryBlock := funcMult.NewBlock("multEntry")
	multEndBlock := funcMult.NewBlock("multEnd")

	//multiplication entry block
	multTemp := multEntryBlock.NewMul(funcMult.Params[0], funcMult.Params[1])
	multResult := multEntryBlock.NewAlloca(types.I32)
	multEntryBlock.NewStore(multTemp, multResult)
	multFinal := multEntryBlock.NewLoad(types.I32, multResult)

	//multiplication result printed
	multEntryBlock.NewCall(printfInt, formatStrMult, multFinal)

	//multiplication final block
	multEntryBlock.NewBr(multEndBlock)
	multEndBlock.NewRet(multFinal)

	//substraction func
	funcSub := moduleHM.NewFunc("sub", types.I32,
		ir.NewParam("x", types.I32))

	//substraction func blocks declarated
	subEntryBlock := funcSub.NewBlock("subEntry")

	subIFBodyBlock := funcSub.NewBlock("subIFBody")
	subElseBlock := funcSub.NewBlock("subElseBody")

	subForBodyBlock := funcSub.NewBlock("subForBody")
	subForCondBlock := funcSub.NewBlock("subForCond")

	subEndBlock := funcSub.NewBlock("subEnd")

	//substraction entry block
	resultSub := subEntryBlock.NewAlloca(types.I32)
	subEntryBlock.NewStore(funcSub.Params[0], resultSub)

	iSub := subEntryBlock.NewAlloca(types.I32)
	subEntryBlock.NewStore(constant.NewInt(types.I32, 0), iSub)

	subX := subEntryBlock.NewAlloca(types.I32)
	subEntryBlock.NewStore(funcSub.Params[0], subX)
	subIFX := subEntryBlock.NewLoad(types.I32, subX)

	//if called
	subEntryBlock.NewBr(subIFBodyBlock)

	//if block
	condIFSub := subIFBodyBlock.NewICmp(enum.IPredSGT, subIFX, constant.NewInt(types.I32, 0))
	subIFBodyBlock.NewCondBr(condIFSub, subForCondBlock, subElseBlock)

	//else block
	subElseBlock.NewBr(subEndBlock)

	//for condition block
	subParam1 := subForCondBlock.NewLoad(types.I32, iSub)
	subParam2 := subForCondBlock.NewLoad(types.I32, subX)
	condForSub := subForCondBlock.NewICmp(enum.IPredSLT, subParam1, subParam2)
	subForCondBlock.NewCondBr(condForSub, subForBodyBlock, subEndBlock)

	//CONDTITION BODY
	tempISub := subForBodyBlock.NewLoad(types.I32, iSub)
	subForBodyBlock.NewCall(printfInt, formatStrI, tempISub)

	//substraction
	sub := subForBodyBlock.NewSub(tempISub, constant.NewInt(types.I32, 10))
	subForBodyBlock.NewCall(printfInt, formatStrSub, sub)

	//multiplication func called
	resultSubTemp := subForBodyBlock.NewCall(funcMult, sub, funcSub.Params[0])
	subForBodyBlock.NewStore(resultSubTemp, resultSub)

	//increment i
	subIncrI := subForBodyBlock.NewAdd(tempISub, constant.NewInt(types.I32, 1))
	subForBodyBlock.NewStore(subIncrI, iSub)

	//for condition block called
	subForBodyBlock.NewBr(subForCondBlock)

	//substraction final block
	subEndBlock.NewBr(subEndBlock)
	returnResult := subEndBlock.NewLoad(types.I32, resultSub)
	subEndBlock.NewRet(returnResult)

	//MAIN
	funcMain := moduleHM.NewFunc("main", types.I32)
	mainBlock := funcMain.NewBlock("")

	//substraction func called
	intValue := mainBlock.NewCall(funcSub, constant.NewInt(types.I32, 5))

	//substraction result printed
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
