package main

import (
	"fmt"
	"sync"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (b *UserBalance) Change(amount int) {
	b.Balance += amount
}

func (b *UserBalance) String() string {
	return fmt.Sprintf("Name: %s, Balance: %d", b.Name, b.Balance)
}

func Transfer(sender *UserBalance, receiver *UserBalance, amount int) {
	// Bad implementation on purpose to trigger a deadlock example.
	sender.Lock()
	fmt.Println("Locked sender", "{", sender, "}")
	sender.Change(-amount)

	time.Sleep(1 * time.Second)

	receiver.Lock()
	fmt.Println("Locked receiver", "{", receiver, "}")
	receiver.Change(amount)

	time.Sleep(1 * time.Second)

	sender.Unlock()
	fmt.Println("Unlocked sender", "{", sender, "}")
	receiver.Unlock()
	fmt.Println("Unlocked receiver", "{", receiver, "}")
}

func main() {
	user1 := UserBalance{
		Name:    "Eko",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Budi",
		Balance: 1000000,
	}

	// Notice that "the receives" are never locked in both calls
	// because they are already locked as senders.
	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)
	//Transfer(&user2, &user1, 200000)
	time.Sleep(3 * time.Second)

	fmt.Println("User:", user1.Name+",", "Balance:", user1.Balance)
	fmt.Println("User:", user2.Name+",", "Balance:", user2.Balance)
}
