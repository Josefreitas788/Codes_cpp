#include<stdio.h>
#include<stdlib.h>

typedef struct ArvoreBinaria
{
    int num;
    struct ArvoreBinaria *esquerdo_ramo;
    struct ArvoreBinaria *direito_ramo;

}ab;

ab* inicializa(ab *ramo)
{
    ab *root;
    root = ramo;
    root = malloc(sizeof(ab));

    int n;
    scanf("%d",&n);
    root->direito_ramo = NULL;
    root->esquerdo_ramo = NULL;
    root->num = n;
    return root;
}

void nova_arvore(ab *ramo)
{
    ab *arvore, *novo, *ant;
    int n;
    int i,j;

    arvore = ramo;
        scanf("%d",&n);

        while(arvore != NULL)
        {
            ant = arvore;

            if( n > arvore->num )
            {
                arvore = arvore->direito_ramo;
  
            }
            else{
                arvore = arvore->esquerdo_ramo;

            }
        }
        arvore = (ab*)  malloc(sizeof(ab));
        arvore->num = n;
        arvore->direito_ramo = NULL;
        arvore->esquerdo_ramo = NULL;
        if(n > ant->num){
            ant->direito_ramo = arvore;
        }
        else{
            ant->esquerdo_ramo = arvore;
        }

}
void imprime_arvore(ab *arvore, int i)
{   
    if(arvore!=NULL)
    {
        i == 0 ? printf(" %d",arvore->num) : 0 ;
        imprime_arvore(arvore->esquerdo_ramo, i);
        i == 1 ? printf(" %d",arvore->num) : 0;
        imprime_arvore(arvore->direito_ramo, i);
        if(i==2)
        {
            printf(" %d",arvore->num);
            free(arvore);
            arvore = NULL;
        }

        
    }
    
}
int main()
{
    ab *ArvoreBin;
    int qtd_casos, qtd_num;
    int i,j,k;
    
    scanf("%d",&qtd_casos);


    for(i = 0; i < qtd_casos; i++)
    {
        
        scanf("%d",&qtd_num);
        ArvoreBin = inicializa(ArvoreBin);
        for(j = 1; j<qtd_num; j++){
            nova_arvore(ArvoreBin); 
        }
        printf("Case %d:",(i+1));
        for ( k = 0; k < 3; k++)
        {
            k == 0 ? printf("\nPre.:"):0;
            k == 1 ? printf("\nIn..:"):0;
            k == 2 ? printf("\nPost:"):0;
            imprime_arvore(ArvoreBin, k);
        }
        printf("\n\n");
        
    }
    return 0;

}


