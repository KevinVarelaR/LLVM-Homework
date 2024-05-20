target datalayout = "e-m:w-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-windows-msvc"

@str = global [15 x i8] c"Resultado: %d\0A\00"
@strMult = global [30 x i8] c"Resultado multiplicacion: %d\0A\00"

declare i32 @printf(i8* %0, i32 %1)

define i32 @mult(i32 %x, i32 %y) {
multEntry:
	%0 = mul i32 %x, %y
	%1 = alloca i32
	store i32 %0, i32* %1
	%2 = load i32, i32* %1
	%3 = call i32 @printf([30 x i8]* @strMult, i32 %2)
	br label %multEnd

multEnd:
	ret i32 %2
}

define i32 @sub(i32 %x) {
subEntry:
	%0 = alloca i32
	store i32 %x, i32* %0
	%1 = alloca i32
	store i32 0, i32* %1
	%2 = alloca i32
	store i32 %x, i32* %2
	%3 = load i32, i32* %2
	br label %subIFBody

subIFBody:
	%4 = icmp sgt i32 %3, 0
	br i1 %4, label %subForCond, label %subElseBody

subElseBody:
	br label %subEnd

subForBody:
	%5 = load i32, i32* %1
	%6 = sub i32 %5, 10
	%7 = call i32 @mult(i32 %6, i32 %x)
	store i32 %7, i32* %0
	%8 = add i32 %5, 1
	store i32 %8, i32* %1
	br label %subForCond

subForCond:
	%9 = load i32, i32* %1
	%10 = load i32, i32* %2
	%11 = icmp slt i32 %9, %10
	br i1 %11, label %subForBody, label %subEnd

subEnd:
	%12 = load i32, i32* %0
	%13 = call i32 @printf([15 x i8]* @str, i32 %12)
	ret i32 %12
}

define i32 @main() {
0:
	%1 = call i32 @sub(i32 5)
	ret i32 %1
}
