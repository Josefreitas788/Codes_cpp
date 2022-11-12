#include<iostream>
#include<string>
using namespace std;

int main()
{
  string word;
  string aux;

  cin >> word;
  aux.push_back('.');
  for(int i = 0 ; i < word.length() ; i++)
  {
    word[i] = tolower(word[i]);
    if(word[i] != 'a' && word[i] != 'e' && word[i] != 'i' && word[i] != 'o' && word[i] != 'u' && word[i] != 'y'){
      aux.push_back(word[i]);
      aux.push_back('.');
    }
  
  }
  aux.pop_back();
  cout << aux << '\n';
  return 0;
}
