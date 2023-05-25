package main

import (
	"fmt"
	"strings"
)

/*
Olhando o exemplo de uma string solta parece um pouco forçado, mas usando dentro de uma estrutura eu diria que é bem
comum encontrar.
*/
type Taxpayer struct {
	FiscalCondition string
}

func main() {

	var taxpayer Taxpayer

	taxpayer.FiscalCondition = "Monotributo"

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(taxpayer.FiscalCondition)
	fmt.Println(strings.Repeat("-", 30))
}

/*
Modelando como uma simples string...
*/

/*
O problema dessa abordagem é que praticamente não existe tipagem, assim como coloquei Monotributo
poderia ter colocado Murilo. Ou seja, é um modelo que diz muito pouco a respeito do que é uma condição fiscal e
mais importante, não garante consistência
*/

/*
Aqui vale mencionar um ponto, creio que a principal DIFERENÇA ENTRE CLEAN ARCH E DDD é que clean arch, como o próprio
nome já diz tem relação com a arquitetura e organização das camadas, foca bastante nos casos de uso. Não que casos de
uso não sejam importantes, eles definitivamente são, mas O CASO DE USO TEM QUE FAZER USO DOS MODELOS DE DOMÍNIO QUE
DEVEM GARANTIR A SUA PRÓPRIA CONSISTÊNCIA, caso contrário a manipulação das informações das entidades fica sobre a
responsabilidade dos diversos possíveis casos de uso e desta forma o PRINCÍPIO DO ENCAPSULAMENTO SE PERDE.
Para ENTENDER UMA CONDIÇÃO FISCAL TERIA QUE CONHECER O QUE TODOS CASOS DE USO fazem com a string fora o risco de
inconsistência.

*/
