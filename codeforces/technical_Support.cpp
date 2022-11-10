#include<iostream>
#include<string>
using namespace std;


int main(){
  int countQuestion = 0;
  string message;
  int quant_msg, tam_msg;
  cin >> quant_msg;
  for(int i = 0 ; i < quant_msg ; i++){
    cin >> tam_msg;
    cin >> message; 
    for(int j = 0 ; j < tam_msg ; j++){
      if(message[j] == 'Q'){
        ++countQuestion;
      }
      if(message[j] == 'A'){
        
        --countQuestion;
        
      }
      if(countQuestion < 0){
        countQuestion = 0;
      }
    }
    if(countQuestion == 0){
      cout << "Yes";
    }
    else{
      cout << "No";
    }
    cout << '\n';
    countQuestion = 0; 
    message.erase();
  }
  return 0;
}
