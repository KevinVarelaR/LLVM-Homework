target datalayout = "e-m:w-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-windows-msvc"

@str = global [18 x i8] c"Final result: %d\0A\00"
@strMult = global [76 x i8] c"Multiplication result (Substraction result * sub func parameter value): %d\0A\00"
@strSub = global [32 x i8] c"Substraction result (i-10): %d\0A\00"
@strI = global [13 x i8] c"i value: %d\0A\00"

declare i32 @printf(i8* %0, i32 %1)

define i32 @mult(i32 %x, i32 %y) {
multEntry:
	%0 = mul i32 %x, %y
	%1 = alloca i32
	store i32 %0, i32* %1
	%2 = load i32, i32* %1
	%3 = call i32 @printf([76 x i8]* @strMult, i32 %2)
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
	%6 = call i32 @printf([13 x i8]* @strI, i32 %5)
	%7 = sub i32 %5, 10
	%8 = call i32 @printf([32 x i8]* @strSub, i32 %7)
	%9 = call i32 @mult(i32 %7, i32 %x)
	store i32 %9, i32* %0
	%10 = add i32 %5, 1
	store i32 %10, i32* %1
	br label %subForCond

subForCond:
	%11 = load i32, i32* %1
	%12 = load i32, i32* %2
	%13 = icmp slt i32 %11, %12
	br i1 %13, label %subForBody, label %subEnd

subEnd:
	%14 = load i32, i32* %0
	ret i32 %14
}

define i32 @main() {
0:
	%1 = call i32 @sub(i32 5)
	%2 = call i32 @printf([18 x i8]* @str, i32 %1)
	ret i32 %1
}
