package or

import (
	"sync"
)

// Or объединяет любое количество done каналов в один, который закрывается в случае, если любой из этих каналов закроется.
// Закрытие объединенного канала происходит безопасно: после закрытия первого из множества каналов остальные можно закрыть
// без паники.
func Or(channels ...<-chan interface{}) <-chan interface{} {
	aggregated := make(chan interface{})

	mutex := sync.Mutex{}

	for _, ch := range channels {
		go func(c <-chan interface{}) {
			<- c
			// Провека, закрыт ли уже объединенный канал
			mutex.Lock()
			defer mutex.Unlock()
			select {
			case <- aggregated:
				return			// уже закрыт
			default:
				close(aggregated)
			}
		}(ch)
	}


	return aggregated
}