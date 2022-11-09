#include<iostream>
#include<string>
#include<algorithm>
#include<set>


int main(){
  int count = 0, odd = 0;
  std::set<char> check;
  std::string username;

  std::cin >> username;
  for(int i = 0; i < username.length(); i++){
    check.insert(username[i]);
  }
 
  if(check.size() % 2 == 0){
    std::cout << "CHAT WITH HER!";
  }else{
    std::cout << "IGNORE HIM!";
  }
  std::cout << '\n';
  return 0;
}

