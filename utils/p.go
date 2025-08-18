package utils

import (
	"fmt"

	C "github.com/ultrazg/xyz/constant"
)

func P(p string) {
	appLogo := `:::    ::: :::   ::: ::::::::: 
:+:    :+: :+:   :+:      :+:  
 +:+  +:+   +:+ +:+      +:+   
  +#++:+     +#++:      +#+    
 +#+  +#+     +#+      +#+     
#+#    #+#    #+#     #+#      
###    ###    ###    ######### `

	fmt.Println("\033[H\033[2J" + appLogo + "v" + C.Version + "\n")
	fmt.Println("API 文档：http://localhost:" + p + "/docs" + "\n")
}
