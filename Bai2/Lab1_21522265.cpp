// Lab 1
// MSSV: 21522265
// Lớp: NT209.011.1

#include <iostream>
#include<bitset>
using namespace std;


int isPositive(int);
// Support function
void PrintBits(unsigned int x)
{
	int i;
	for (i = 8 * sizeof(x) - 1; i >= 0; i--)
	{
		if (x & (1 << i))
			cout << '1';
		else
			cout << '0';
	}
	cout << "\n";
}
void PrintBitsOfByte(unsigned int x)
{
	int i;
	for (i = 7; i >= 0; i--)
	{
		if (x & (1 << i))
			cout << '1';
		else
			cout << '0';
	}
	cout << "\n";
}

// 1.1
int bitAnd(int x, int y)
{
    return ~(~x | ~y);
}
// 1.2
int negative(int x)
{
    return (1 + ~(x));
}
// 1.3
int setBit(int x, int n)
{
    int y = 1;
    y = y << n;
	return x | y;
}
// 1.4
int getnBit(int x, int n)
{
	int mask = ~(2147483647 << n);
    return x & mask;
}
// 1.5
int most1bit(unsigned int x)
{
	int count = 0;
    x = x >> 1;
	count += isPositive(x);
    x = x >> 1;
	count += isPositive(x);

    x = x >> 1;
	count += isPositive(x);

    x = x >> 1;
	count += isPositive(x);

    x = x >> 1;
	count += isPositive(x);

    x = x >> 1;
	count += isPositive(x);

    x = x >> 1;
	count += isPositive(x);

    x = x >> 1;
	count += isPositive(x);

	return count;
}

// 2.1
int isSameSign(int x, int y)
{
    return (((x >> 31) & 1) ^ ((y >> 31) & 1)) ^ 1;	
}
// 2.2
int is4x(int x)
{
	int modulo = x & ((1 << 2) - 1);
    int result = !(modulo & 1);
    return result;
}
// 2.3
int isPositive(int x)
{
	return(!(x >> 31) ^ !x);
}

// 2.4
int isGE(int x, int y)
{
	return !((x + negative(y)) >> 31);
}

int main() {
	float score = 0;
	// check ket qua cac ham dua tren cac phep thu
	// 1.1
	cout << "=== PART 1\n";
	cout << "1.1 - bitAnd()\t\t";
	if (bitAnd(3, -9) == 3 & -9)
	{
		cout <<"Pass.\n";
		score += 1;
	}
	else
		cout <<"Failed. Please check again.\n";

	// 1.2
	cout <<"1.2 - negative()\t";
	if (negative(0) == 0 && negative(9) == -9 && negative(-5) == 5)
	{
		cout <<"Pass.\n";
		score += 1;
	}
	else
		cout <<"Failed. Please check again.\n";

	// 1.3
	cout <<"1.3 - setBit()\t\t";
	if (setBit(0,2) == 4 && setBit(5,1) == 7 && setBit(6,2) == 6)
	{
		cout <<"Pass.\n";
		score += 2;
	}
	else
		cout <<"Failed. Please check again.\n";

	// 1.4
	cout <<"1.4 - getnBit()\t\t";
	if (getnBit(0x11223344,2) == 0x0 && getnBit(0x11223344,8) == 0x44 && getnBit(0x11223344, 9) == 0x144)
	{
		cout <<"Pass.\n";
		score += 2;
	}
	else
		cout <<"Failed. Please check again.\n";

	// 1.5
	cout <<"1.5 - most1bit()\t";
	if (most1bit(2) == 1 && most1bit(31) == 4 && most1bit(255) == 7)
	{
		cout <<"Pass.\n";
		score += 4;
	}
	else
		cout <<"Failed. Please check again.\n";

	// 2.1
	cout <<"\n=== PART 2\n";
	cout <<"2.1 - isSameSign()\t";
	if (isSameSign(2,10) && !isSameSign(5,-4) && isSameSign(-21,-8))
	{
		cout <<"Pass.\n";
		score += 2;
	}
	else
		cout <<"Failed. Please check again.\n";

	// 2.2
	cout <<"2.2 - is4x()\t\t";
	if (is4x(4) && !is4x(15) && is4x(0))
	{
		cout <<"Pass.\n";
		score += 2;
	}
	else
		cout <<"Failed. Please check again.\n";

	// 2.3
	cout <<"2.3 - isPositive()\t";
	if (isPositive(10) && !isPositive(-5) && !isPositive(0))
	{
		cout <<"Pass.\n";
		score += 3;
	}
	else
		cout <<"Failed. Please check again.\n";

	// 2.4
	cout <<"2.4 - isGE()\t\t";
	if (isGE(8,7) && isGE(10,10) && !isGE(-5,9))
	{
		cout << "Pass.\n";
		score += 3;
	}
	else
		cout << "Failed. Please check again.\n";

	score = (float)score / 2;

	cout << "\n===\nYour final score: " << score << "\n";
	if (score == 10)
		cout << "Exellent!! You've got a genious mind and solved all problems. Want more? :)))\n";
	else if (score < 10 && score >= 8)
		cout << "Good job! You're so closed, try more to solve failed problem(s).";
	else if (score < 8 && score >= 5)
		cout << "Nice try! Think more carefully about your solutions.";
	else
		cout << "Got any trouble in solving these problems? Contact your lecturer for some hints :)";
	//cin >> c;

}