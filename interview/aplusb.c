#include<stdio.h>
#include<stdlib.h>

int abPlus();

int main(int argc, char* argv[],char** envp)
{
    int r;

    int i;
    printf("there are %d input param\n",argc);
    printf("they are :\n");
    for(i = 0 ; i < argc ; i++)
    {
        printf("%s\n",argv[i]);
    }
    r = abPlus(atoi(argv[1]), atoi(argv[2]));
    printf("%d", r);
    return r;
}

int abPlus(int a, int b){
    int tp;
	if (a == 0) {
		return b;
	}

	if (b == 0) {
		return a;
	}

    int x = a ^b ;
    int y = a&b << 1;
    return abPlus(x, y);
}
