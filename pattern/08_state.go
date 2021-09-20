// Реализация паттерна "состояние" на примере игрока из некоторой игры
package pattern

import (
	"fmt"
)

// Player - игрок
type Player struct {
	state State
	x, y float64
}

// NewPlayer - конструктор Player
func NewPlayer() *Player {
	var player = &Player{}
	player.setState(NewHealthyState(player))
	return player
}

// setState устанавливает состояние игрока
func (p *Player) setState(s State){
	p.state = s
}

// Move делегирует передвижение состоянию
func (p *Player) Move(dx, dy float64) {
	p.state.Move(dx, dy)
}

// Action делегирует действие состоянию
func (p *Player) Action() {
	p.state.Action()
}

// Injure ранит игрока, делигирует состоянию
func (p *Player) Injure() {
	p.state.Injure()
}

// Heal лечит игрока, делигирует состоянию
func (p *Player) Heal() {
	p.state.Heal()
}

// Интерфейс состояния
type State interface {
	Move(dx, dy float64)
	Action()
	Injure()
	Heal()
}

// HealthyState - состояние здорового игрока
type HealthyState struct {
	player *Player
}

// NewHealthyState - конструктор HealthyState
func NewHealthyState(p *Player) *HealthyState {
	return &HealthyState{p}
}

// Move передвигает игрока быстро
func (hs *HealthyState) Move(dx, dy float64) {
	hs.player.x += dx
	hs.player.y += dy
	fmt.Println("Healthy player moves quickly")
}

// Action совершает полноценное действие
func (hs *HealthyState) Action() {
	fmt.Println("Healthy player perfroms a full action")
}

// Injure ранит игрока, переходит в раненое состояние
func (hs *HealthyState) Injure() {
	hs.player.setState(NewInjuredState(hs.player))
	fmt.Println("Healthy player is injured now")
}

// Heal ничего не делает, так как игрок уже здоров
func (hs *HealthyState) Heal() {
	fmt.Println("The player is already healthy!")
}



// InjuredState - состояние раненого игрока
type InjuredState struct {
	player *Player
}

// NewInjuredState - конструктор InjuredState
func NewInjuredState(p *Player) *InjuredState {
	return &InjuredState{p}
}

// Move передвигает игрока медленно
func (is *InjuredState) Move(dx, dy float64) {
	is.player.x += dx / 2
	is.player.y += dy / 2
	fmt.Println("Injured player moves slowly")
}

// Action совершает ограниченное действие
func (is *InjuredState) Action() {
	fmt.Println("Injured player perfroms a limited action")
}

// Injure ранит игрока, переводит его в умирающее состояние
func (is *InjuredState) Injure() {
	is.player.setState(NewDyingStateState(is.player))
	fmt.Println("Injured player is now dying")
}

// Heal лечит игрока, переводит его в здоровое состояние
func (is *InjuredState) Heal() {
	is.player.setState(NewHealthyState(is.player))
	fmt.Println("Injured player is now healed")
}





// DyingState - состояние умирающего игрока
type DyingState struct {
	player *Player
}

// NewDyingStateState - конструктор DyingState
func NewDyingStateState(p *Player) *DyingState {
	return &DyingState{p}
}

// Move ничего не делает, так как игрок умирает
func (ds *DyingState) Move(dx, dy float64) {
	fmt.Println("Dying player can't move!")
}

// Action пытается поднять умирающего игрока
func (ds *DyingState) Action() {
	fmt.Println("Dying player tries to get up")
}

// Heal лечит игрока, переводит его в здоровое состояние
func (ds *DyingState) Heal() {
	ds.player.setState(NewInjuredState(ds.player))
	fmt.Println("Dying player is up and merely injured")
}

// Injure ничего не делает, так как игрок уже умирает
func (ds *DyingState) Injure() {
	fmt.Println("The player is already dying")
}


// Пример использования
/* func main() {
	var player = NewPlayer()

	player.Move(5, 5)
	player.Action()
	
	player.Injure()
	player.Move(5, 5)
	player.Action()

	player.Injure()
	player.Move(5, 5)
	player.Action()

	player.Heal()
	player.Move(5, 5)
	player.Action()

	player.Heal()
	player.Move(5, 5)
	player.Action()
}
*/