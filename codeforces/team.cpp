#include<iostream>

using namespace std;

int main(){
  int quant_questions, a,b,c;
  int count = 0;

  cin>>quant_questions;

  for(int i = 0 ; i < quant_questions ; i++){
    cin>>a>>b>>c;
    if(a+b+c >= 2){
      count++;
    }
  }
  cout<<count<<'\n';
  return 0;
}
