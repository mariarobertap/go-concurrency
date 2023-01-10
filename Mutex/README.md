## Mutex 



Mutex permite que você lide com situação de race condition (concorrência)
Você pode dar lock em um dado enquanto estiver usando, e após terminar é só dar unlock.

Dessa maneira as goroutines irão acessar o dado com segurança.
Caso o dado esteja em lock, e outra goroutine tente acessar, ela irá esperar até que o dado esteja disponivel.

## Exemplo 

```go
func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func main() {

	var mutex sync.Mutex

	wg.Add(2)
	updateMessage("Message 2", &mutex)
	updateMessage("Message 1", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
```


