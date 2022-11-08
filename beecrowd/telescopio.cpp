#include<iostream>

using namespace std;

int main(){
  int abertura, num_estrelas, estrela, count = 0;
  int fotonsEstrela = 40000000;
  cin >> abertura;
  cin >> num_estrelas;
  
  for(int i = 0; i < num_estrelas; i++){
    cin >> estrela;
  
    if(fotonsEstrela <= abertura * estrela){
      count++;
    }
  }
  cout << count << '\n';
  return 0;
}
