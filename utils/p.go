package utils

import (
	"fmt"
	C "github.com/ultrazg/xyz/constant"
	"runtime"
)

func P(p string) {
	appLogo := `:::    ::: :::   ::: ::::::::: 
:+:    :+: :+:   :+:      :+:  
 +:+  +:+   +:+ +:+      +:+   
  +#++:+     +#++:      +#+    
 +#+  +#+     +#+      +#+     
#+#    #+#    #+#     #+#      
###    ###    ###    ######### `

	fmt.Println(appLogo + "v" + C.Version + " built with " + runtime.Version() + "\n")
	fmt.Println("API 文档：http://localhost:" + p + "/docs")
}
