# go-bytes

## Проблема
- есть битовый слайс `[][]bool`
- надо получить слайс байт `[]byte`
- каждая последовательность бит преобразуется в последовательность байт, например:
  `[]bool{true, true, false, false, true, true, false, false}` => `[]byte{204}`
- если количество бит больше восьми, то в итоговом слайсе больше одного байта, например:
  `[]bool{false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, true}` => `[]byte{1, 1}`
- если бит недостаточно, чтобы заполнить байт, то правые разряды байта заполняются нулями, например:
  `[]bool{true, false, true, false}` => `[]byte{160}`
- каждый изисходных слайсов преобразуется отдельно, а затем результат конкатенируется в одну последовательность байт
