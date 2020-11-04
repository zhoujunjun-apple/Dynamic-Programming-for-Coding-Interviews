unsigned int factorialFunc(unsigned int n) {
    if(n<=0){
        return 0;
    }
    if(n==1){
        return 1;
    }
    return n * factorialFunc(n-1);
}