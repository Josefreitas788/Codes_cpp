#include<iostream>
#include<vector>
using namespace std;

int main(){
  int n, aux, tamanho;
  vector<int> nums, aux_vet, aux_vet2;
  cin >> n;

  for(int i = 0; i < n; i++){
    cin >> aux;
    nums.push_back(aux);

  }
  for(int i = 0; i < n; i++){
    for(int j = 1; j <= nums[i]; j++){
      aux_vet.push_back(j);
    }
    tamanho = aux_vet.size();
        if(tamanho %2 == 0){
          for(int k = 1; k <= tamanho /2; k++){
            aux_vet2.push_back(aux_vet[(tamanho/2)-k]);
            aux_vet2.push_back(aux_vet[tamanho-k]);
          
          }
        }
        else{
          aux_vet2.push_back(aux_vet[0]);
          for(int k = 1; k <= tamanho/2; k++){
                    
            aux_vet2.push_back(aux_vet[(tamanho/2+1)-k]);
            aux_vet2.push_back(aux_vet[tamanho-k]);
          }

        }
    for(int j = 0; j < aux_vet2.size(); j++){
      cout << aux_vet2[j] << " ";
    }
    cout << endl;
    aux_vet.clear();
    aux_vet2.clear();

  }
}


