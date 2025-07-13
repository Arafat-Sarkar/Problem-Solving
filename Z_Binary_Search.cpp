#include<bits/stdc++.h>
using namespace std;

int main() {
    int N, Q;
    cin >> N >> Q;
    int A[N];
    
    
    for(int i = 0; i < N; i++) {
        cin >> A[i];
    }

    
    while(Q--) {
        int X;
        cin >> X;
        bool found = false;

       
        for(int i = 0; i < N; i++) {
            if(A[i] == X) {
                found = true;
                break;
            }
        }

        if(found) {
            cout << "found" << endl;
        } else {
            cout << "not found" << endl;
        }
    }

    return 0;
}
