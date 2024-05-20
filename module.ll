target datalayout = "e-m:w-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-windows-msvc"

@str = global [15 x i8] c"Resultado: %d\0A\00"

declare i32 @printf(i8* %0, i32 %1)

define i32 @mult(i32 %x, i32 %y) {
multEntry:
	%0 = alloca i32
	store i32 0, i32* %0
	%1 = load i32, i32* %0
	%2 = call i32 @printf([15 x i8]* @str, i32 %x)
	%3 = call i32 @printf([15 x i8]* @str, i32 %y)
	br label %multIFBody

multIFBody:
	%4 = icmp eq i32 %x, 0
	%5 = icmp eq i32 %y, 0
	%6 = or i1 %4, %5
	br i1 %6, label %multEnd, label %multElseBody

multElseBody:
	%7 = alloca i32
	store i32 0, i32* %7
	br label %multForCond

multForBody:
	%8 = mul i32 %x, %y
	%9 = mul i32 %8, %13
	store i32 %9, i32* %7
	%10 = add i32 %13, 1
	store i32 %10, i32* %11
	br label %multForCond

multForCond:
	%11 = alloca i32
	store i32 0, i32* %11
	%12 = alloca i32
	store i32 %y, i32* %12
	%13 = load i32, i32* %11
	%14 = load i32, i32* %12
	%15 = icmp slt i32 %13, %14
	br i1 %15, label %multForBody, label %multEnd

multEnd:
	%16 = load i32, i32* %7
	ret i32 %16
}

define i32 @sub(i32 %x) {
subEntry:
	%0 = alloca i32
	store i32 %x, i32* %0
	%1 = alloca i32
	store i32 0, i32* %1
	br label %subIFBody

subIFBody:
	%2 = alloca i32
	store i32 %x, i32* %2
	%3 = load i32, i32* %2
	%4 = icmp sgt i32 %3, 0
	br i1 %4, label %subForCond, label %subElseBody

subElseBody:
	br label %subEnd

subForBody:
	%5 = load i32, i32* %1
	%6 = load i32, i32* %12
	%7 = icmp slt i32 %5, %6
	%8 = load i32, i32* %1
	%9 = udiv i32 %8, 2
	%10 = call i32 @mult(i32 %9, i32 %x)
	store i32 %10, i32* %0
	%11 = add i32 %8, 1
	store i32 %11, i32* %1
	br label %subForBody

subForCond:
	%12 = alloca i32
	store i32 %x, i32* %12
	br label %subForBody

subEnd:
	%13 = call i32 @printf([15 x i8]* @str, i32* %0)
	%14 = load i32, i32* %0
	ret i32 %14
}

define i32 @main() {
0:
	%1 = call i32 @sub(i32 5)
	%2 = call i32 @printf([15 x i8]* @str, i32 %1)
	ret i32 %1
}
