#include<iostream>
#include<string>
#include<vector>
using namespace std;
int main(){
  
  int n;
  string palavra;
  string abreviacao;
  vector<string> palavras, abreviacoes;
  cin >> n;
  for(int i = 0; i < n; i++){
    
    cin >> palavra;
    palavras.push_back(palavra);
  }
  for(int i = 0; i < n; i++){
    if(palavras[i].length() > 10){
      
      abreviacao  = palavras[i][0];
      abreviacao += to_string(palavras[i].length() - 2);
      abreviacao += palavras[i][palavras[i].length() - 1];
      
      abreviacoes.push_back(abreviacao);
    
    }else{

      abreviacoes.push_back(palavras[i]);
    
    }
  }
  for(int i = 0; i < n; i++){
    
    cout << abreviacoes[i] << endl;
  }
return 0;

}
