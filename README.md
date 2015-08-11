Go Tomita Parser wrapper
========================

Небольшой враппер для удобной работы с Томита Парсер от Яндекса в Go.

Пример
------

```go
import (
	"fmt"
	"github.com/makhov/go-tomita"
)

func main() {
	p, _ := tomita.New("/home/user/tomita-parser", "/home/user/config.proto")
	p.SetDebug(true)
	output, err := p.Run("This is text to parse")
	if err != nil {
		...	
	}
}

```

Важно
-----

- config.proto не должен содержать дескрипторов File (ввод/вывод осуществляется через STDIN/STDOUT)


За пример спасибо автору [poor-python-yandex-tomita-parser](https://github.com/vas3k/poor-python-yandex-tomita-parser).
