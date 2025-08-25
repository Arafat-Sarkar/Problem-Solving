#include<bits/stdc++.h>
using namespace std;
int main()
{
   char ch;
   cin>>ch;
   if((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z'))
   {
    cout<<ch <<" "<<" Is character "<<endl;
   }
   else if(ch >= '0' && ch <= '9')
   {
     cout<<ch <<" "<<" Is Digit "<<endl;
   }
   else
   {
     cout<<ch <<" "<<" Is Special Symbol "<<endl;
   }
    return 0;
}