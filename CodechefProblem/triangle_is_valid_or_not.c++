#include<bits/stdc++.h>
using namespace std;
int main()
{
   int a,b,c;
   cin>>a>>b>>c;
   int sum = a+b+c;
   if(sum==180 && a>0 && b>0 && c>0)
   {
     cout<<"Triangle is valid."<<endl;
   } 
   else
   {
    cout<<"Triangle is not valid"<<endl;
   }
    return 0;
}