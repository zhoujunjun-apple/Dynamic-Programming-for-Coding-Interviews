#include <stdio.h>
#include "factorial.c"

typedef struct TestSuit {
    unsigned int inputN;
    unsigned int expOutput;
}TestIO;

void main(void) {
    const int SIZE = 5;
    TestIO t1 = {1, 1};
    TestIO t2 = {2, 2};
    TestIO t3 = {3, 6};
    TestIO t4 = {4, 24};
    TestIO t5 = {5, 120};

    TestIO allTests[SIZE];
    allTests[0] = t1;
    allTests[1] = t2;
    allTests[2] = t3;
    allTests[3] = t4;
    allTests[4] = t5;

    int idx = 0;
    unsigned int getOut, expInp, expOut;
    for(; idx < SIZE; idx++){
        expInp = allTests[idx].inputN;
        expOut = allTests[idx].expOutput;
        getOut = factorialFunc(expInp);
        if(getOut != expOut){
            printf("input=%d, expected=%d, get=%d\n", expInp, expOut, getOut);
        }
    }
    printf("test finish.\n");
}