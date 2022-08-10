define i64 @main() {
0:
	%1 = alloca i64
	store i64 1, i64* %1
	%2 = load i64, i64* %1
	%3 = alloca i64
	store i64 5, i64* %3
	%4 = load i64, i64* %3
	%5 = add i64 %2, %4
	ret i64 %5
}
