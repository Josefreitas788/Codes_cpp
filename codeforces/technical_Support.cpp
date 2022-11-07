#include<iostream>
#include<string>
using namespace std;


int main(){
  int countQuestion = 0, countAnswer = 0;
  string message;
  int quant_msg, tam_msg;
  cin >> quant_msg;
  for(int i = 0 ; i < quant_msg ; i++){
    cin >> tam_msg;
    cin >> message; 
    if(message[0] == 'A'){
      cout << "No" << endl;
    }
    else if(message[message.size()-1] == 'Q'){
      cout << "No" << endl;
    }
    else{
    for(int j = 0 ; j < tam_msg ; j++){
      if(message[j] == 'Q'){
        countQuestion++;
      }
      if(message[j] == 'A'){
        countAnswer++;
      }
    }
    if(countAnswer >= countQuestion){
      cout << "Yes";
    }
    else{
      cout << "No";
    }
    cout << endl;
    countQuestion = 0;
    countAnswer = 0;
    message.erase();
    }
  }
  return 0;
}
