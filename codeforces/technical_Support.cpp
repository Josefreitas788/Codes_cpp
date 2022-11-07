#include<iostream>
#include<string>
using namespace std;


int main(){
  int countQuestion = 0, countAnswer = 0;
  string messages;
  int quant_msg, tam_msg;
  cin >> quant_msg;
  for(int i = 0 ; i < quant_msg ; i++){
    getline(cin, messages);
    cin >> tam_msg;
    for(int j = 0 ; j < tam_msg ; j++){
      if(messages[j] == 'Q'){
        countQuestion++;
      }
      if(messages[j] == 'A'){
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
  }
  return 0;
}
