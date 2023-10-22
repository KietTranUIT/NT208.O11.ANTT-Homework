#include<stdio.h>
#include<conio.h>

int main(){
	int n;
	float sum =0;
	scanf("%d", &n);
	for(int i = 1; i<=n; i++)
	   sum += 1/float(n*(n+1));
	printf("Tong: %.2f", sum);
	return 0;
	}