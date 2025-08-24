#include <bits/stdc++.h>
using namespace std;

int main() {
    int X, Y;
    cin >> X >> Y;
    int result = (X + Y - 1) % 12 + 1;

    cout << result << endl;
    return 0;
}
