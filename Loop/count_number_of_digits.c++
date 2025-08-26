#include<bits/stdc++.h>
using namespace std;
int main()
{
   long long int n;
   cin>>n;
   int count = 0;
   if(n==0)
   {
     cout<<1<<endl;
   } 
   else
   {
      while (n!=0)
      {
         count++;
         n /=10;
      }
      
   }
   cout<<count<<endl;
    return 0;
}