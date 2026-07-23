# Go Design Patterns

22 classic design patterns implemented in Go. All examples use composition over inheritance, embracing Go's idioms.

## Categories

### Creational Patterns

| Pattern | Description |
|---|---|
| [Singleton](#singleton) | One instance, globally accessible |
| [Factory](#factory-method) | Create objects without specifying exact class |
| [Abstract Factory](#abstract-factory) | Create families of related objects |
| [Builder](#builder) | Construct complex objects step by step |
| [Prototype](#prototype) | Clone existing objects |

### Structural Patterns

| Pattern | Description |
|---|---|
| [Adapter](#adapter) | Bridge incompatible interfaces |
| [Bridge](#bridge) | Decouple abstraction from implementation |
| [Composite](#composite) | Tree structure of objects |
| [Decorator](#decorator) | Add behavior dynamically |
| [Facade](#facade) | Simplified interface to complex subsystem |
| [Flyweight](#flyweight) | Share objects to reduce memory |
| [Proxy](#proxy) | Control access to another object |

### Behavioral Patterns

| Pattern | Description |
|---|---|
| [Chain of Responsibility](#chain-of-responsibility) | Pass request along a chain |
| [Command](#command) | Encapsulate a request as an object |
| [Iterator](#iterator) | Sequential access without exposing internals |
| [Mediator](#mediator) | Centralized communication between objects |
| [Memento](#memento) | Capture and restore object state |
| [Observer](#observer) | Notify dependents of state changes |
| [State](#state) | Change behavior based on internal state |
| [Strategy](#strategy) | Swap algorithms at runtime |
| [Template Method](#template-method) | Define skeleton, let subclasses fill in |
| [Visitor](#visitor) | Add operations without modifying classes |

---

## Singleton

Ensures only one instance exists. Uses **double-checked locking** with `sync.Mutex`.

```go
package instance

import (
    "fmt"
    "sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func GetInstance() *single {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        if singleInstance == nil {
            fmt.Println("Creating single instance now.")
            singleInstance = &single{}
        }
    }
    return singleInstance
}

func main() {
    for i := 0; i < 30; i++ {
        go GetInstance()
    }
    fmt.Scanln()
}
```

---

## Factory Method

Creates objects through a factory function, decoupling client from concrete types.

```go
type IGun interface {
    setName(name string)
    getName() string
    setPower(power int)
    getPower() int
}

type Gun struct {
    name  string
    power int
}

type Ak47 struct{ Gun }     // NewAk47() sets name="AK47 gun", power=4
type musket struct{ Gun }   // NewMusket() sets name="Musket gun", power=1

func Create(gunType string) (IGun, error) {
    if gunType == "ak47"  { return NewAk47(), nil }
    if gunType == "musket" { return NewMusket(), nil }
    return nil, fmt.Errorf("wrong gun type passed")
}

// Usage
ak47, _ := Create("ak47")
musket, _ := Create("musket")
```

---

## Abstract Factory

Creates **families** of related objects (e.g., Nike shoes + Nike shirts).

```go
type ISportsFactory interface {
    makeShoe() IShoe
    makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
    if brand == "adidas" { return &Adidas{}, nil }
    if brand == "nike"   { return &Nike{}, nil }
    return nil, fmt.Errorf("invalid brand")
}

type IShoe interface  { getLogo() string; getSize() int }
type IShirt interface { getLogo() string; getSize() int }

// Adidas and Nike implement ISportsFactory, returning branded shoes/shirts
```

---

## Builder

Constructs complex objects step-by-step with a fluent interface.

```go
type builder interface {
    TopSpeed(int) builder
    Paint(string) builder
    Build() Vehicle
}

type CarBuilder struct {
    speedOption int
    color       string
}

func (cb *CarBuilder) TopSpeed(speed int) builder {
    cb.speedOption = speed
    return cb   // return self for chaining
}

func Builder() builder { return &CarBuilder{} }

// Usage
car := car.Builder().TopSpeed(50).Paint("blue").Build()
fmt.Println(car.Drive()) // "Driving at speed: 50"
fmt.Println(car.Stop())  // "Stopping a blue car"
```

---

## Prototype

Creates new objects by cloning existing ones (deep copy).

```go
type Inode interface {
    print(string)
    clone() Inode
}

type File struct{ name string }
func (f *File) clone() Inode { return &File{name: f.name + "_clone"} }

type Folder struct {
    children []Inode
    name     string
}
func (f *Folder) clone() Inode {
    cloneFolder := &Folder{name: f.name + "_clone"}
    for _, i := range f.children {
        cloneFolder.children = append(cloneFolder.children, i.clone())
    }
    return cloneFolder
}
```

---

## Adapter

Makes incompatible interfaces work together.

```go
type Computer interface { InsertIntoLightningPort() }

type Mac struct{}
func (m *Mac) InsertIntoLightningPort() { fmt.Println("Lightning connector plugged in.") }

type Windows struct{}
func (w *Windows) insertIntoUSBPort() { fmt.Println("USB connector plugged in.") }

// Adapter wraps Windows to work with Lightning port
type WindowsAdapter struct { windowMachine *Windows }
func (w *WindowsAdapter) InsertIntoLightningPort() {
    fmt.Println("Adapter converts Lightning signal to USB.")
    w.windowMachine.insertIntoUSBPort()
}

// Usage
client.InsertLightningConnectorIntoComputer(mac)                    // works natively
client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)  // via adapter
```

---

## Bridge

Decouples an abstraction from its implementation so both can vary independently.

```go
type Computer interface { Print(); SetPrinter(Printer) }
type Printer  interface { PrintFile() }

type Mac     struct{ printer Printer }
type Windows struct{ printer Printer }
type Hp      struct{}   // Printer implementation
type Epson   struct{}   // Printer implementation

// Usage: swap printers at runtime
macComputer.SetPrinter(hpPrinter)
macComputer.Print()
macComputer.SetPrinter(epsonPrinter)
macComputer.Print()
```

---

## Composite

Treats individual objects and compositions uniformly (tree structure).

```go
type Component interface { Search(string) }

type File struct{ name string }
func (f *File) Search(keyword string) { fmt.Printf("Searching %s in file %s\n", keyword, f.name) }

type Folder struct {
    components []Component
    name       string
}
func (f *Folder) Search(keyword string) {
    for _, c := range f.components { c.Search(keyword) }
}
func (f *Folder) Add(components ...Component) {
    f.components = append(f.components, components...)
}
```

---

## Decorator

Adds behavior to objects dynamically by wrapping them.

```go
type IPizza interface { getPrice() int }

type VeggeMania struct{}          // base price: 15
func (p *VeggeMania) getPrice() int { return 15 }

type CheeseTopping struct{ pizza IPizza }  // +10
func (c *CheeseTopping) getPrice() int { return c.pizza.getPrice() + 10 }

type TomatoTopping struct{ pizza IPizza }  // +7
func (c *TomatoTopping) getPrice() int { return c.pizza.getPrice() + 7 }

// Usage: stack decorators
pizza := &VeggeMania{}
pizzaWithCheese := &CheeseTopping{pizza: pizza}
pizzaWithCheeseAndTomato := &TomatoTopping{pizza: pizzaWithCheese}
fmt.Println(pizzaWithCheeseAndTomato.getPrice()) // 32
```

---

## Facade

Provides a simplified interface to a complex subsystem.

```go
// Complex subsystem
type Account      struct{}  // checkAccount()
type SecurityCode struct{}  // checkCode()
type Wallet       struct{}  // creditBalance(), debitBalance()
type Notification struct{}  // sendWalletCreditNotification(), sendWalletDebitNotification()
type Ledger       struct{}  // makeEntry()

// Facade hides complexity behind two simple methods
type WalletFacade struct {
    account, wallet, securityCode, notification, ledger
}

func New(accountID string, code int) *WalletFacade { /* init all subsystems */ }

func (w *WalletFacade) Add(accountID string, code int, amount int) error {
    w.account.checkAccount(accountID)
    w.securityCode.checkCode(code)
    w.wallet.creditBalance(amount)
    w.notification.sendWalletCreditNotification()
    w.ledger.makeEntry(accountID, "credit", amount)
    return nil
}

// Usage
wallet := wallet.New("Abhishek Ghosh", 1234)
wallet.Add("Abhishek Ghosh", 1234, 10)
wallet.Deduct("Abhishek Ghosh", 1234, 5)
```

---

## Flyweight

Shares common state across many objects to reduce memory usage.

```go
// Intrinsic (shared) state
type Dress interface { getColor() string }
type TerroristDress struct{ color string }        // red
type CounterTerroristDress struct{ color string } // green

// Flyweight factory: reuses dress objects by type
type DressFactory struct { dressMap map[string]Dress }

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
    if d.dressMap[dressType] != nil { return d.dressMap[dressType], nil }
    // create if not exists
}

// Extrinsic (unique) state
type Player struct {
    dress      Dress    // shared flyweight
    playerType string   // "T" or "CT"
    lat, long  int      // unique position
}
```

---

## Proxy

Controls access to another object — here, an Nginx-like rate limiter.

```go
type Server interface { handleRequest(url, method string) (int, string) }
type Application struct{}  // real server

type Nginx struct {
    server            Server
    maxAllowedRequest int
    rateLimiter       map[string]int
}

func (n *Nginx) handleRequest(url, method string) (int, string) {
    if !n.checkRateLimiting(url) { return 403, "Not Allowed" }
    return n.server.handleRequest(url, method)
}

// Usage: max 2 requests per URL
nginxServer := newNginxServer(&Application{})
nginxServer.handleRequest("/app/status", "GET")  // 200
nginxServer.handleRequest("/app/status", "GET")  // 200
nginxServer.handleRequest("/app/status", "GET")  // 403
```

---

## Chain of Responsibility

Passes a request along a chain of handlers. Each handler decides to process or pass.

```go
type Department interface {
    execute(*Patient)
    setNext(Department) Department   // returns next for fluent chaining
}

type Reception struct{ next Department }
func (r *Reception) execute(p *Patient) {
    if p.registrationDone { r.next.execute(p); return }
    fmt.Println("Reception registering patient")
    p.registrationDone = true
    r.next.execute(p)
}

// Chain: Reception → Doctor → Medical → Cashier
reception.setNext(doctor).setNext(medical).setNext(cashier)
reception.execute(&Patient{name: "Abhishek Ghosh"})
```

---

## Command

Encapsulates a request as an object, enabling parameterization and queuing.

```go
type Command interface { execute() }
type Device  interface { on(); off() }

type OnCommand  struct{ device Device }
type OffCommand struct{ device Device }

type Tv struct{ isRunning bool }
func (t *Tv) on()  { t.isRunning = true; fmt.Println("Turning tv on") }
func (t *Tv) off() { t.isRunning = false; fmt.Println("Turning tv off") }

type Button struct{ command Command }
func (b *Button) press() { b.command.execute() }

// Usage
onButton  := &Button{command: &OnCommand{device: tv}}
offButton := &Button{command: &OffCommand{device: tv}}
onButton.press()
offButton.press()
```

---

## Iterator

Provides sequential access to elements without exposing underlying representation.

```go
type Collection[T any] interface { Iterator() Iterator[T]; Add(data ...*T) }
type Iterator[T any]   interface { HasNext() bool; Next() *T }

type List[T any] struct{ datas []*T }

type ListIterator[T any] struct {
    index int
    datas []*T
}

func (li *ListIterator[T]) HasNext() bool { return li.index < len(li.datas) }
func (li *ListIterator[T]) Next() *T {
    if li.HasNext() { user := li.datas[li.index]; li.index++; return user }
    return nil
}

// Usage
userList := list.New(user1, user2, user3)
iterator := userList.Iterator()
for iterator.HasNext() { user := iterator.Next(); user.Print() }
```

---

## Mediator

Centralizes communication between objects to reduce coupling.

```go
type Mediator interface {
    canArrive(Train) bool
    notifyAboutDeparture()
}

type Train interface {
    arrive(); depart(); permitArrival()
}

// StationManager coordinates trains
type StationManager struct {
    isPlatformFree bool
    trainQueue     []Train
}

func (s *StationManager) canArrive(t Train) bool {
    if s.isPlatformFree { s.isPlatformFree = false; return true }
    s.trainQueue = append(s.trainQueue, t)
    return false
}

func (s *StationManager) notifyAboutDeparture() {
    s.isPlatformFree = true
    if len(s.trainQueue) > 0 {
        next := s.trainQueue[0]
        s.trainQueue = s.trainQueue[1:]
        next.permitArrival()
    }
}
```

---

## Memento

Captures and restores an object's internal state without violating encapsulation.

```go
type Memento struct{ state string }

type Originator struct{ state string }
func (e *Originator) createMemento() *Memento { return &Memento{state: e.state} }
func (e *Originator) restoreMemento(m *Memento) { e.state = m.getSavedState() }

type Caretaker struct{ mementoArray []*Memento }
func (c *Caretaker) addMemento(m *Memento)      { c.mementoArray = append(c.mementoArray, m) }
func (c *Caretaker) getMemento(index int) *Memento { return c.mementoArray[index] }

// Usage
caretaker.addMemento(originator.createMemento())   // save state "A"
originator.setState("B")
caretaker.addMemento(originator.createMemento())   // save state "B"
originator.restoreMemento(caretaker.getMemento(0)) // restore to "A"
```

---

## Observer

Notifies dependents automatically when state changes (pub/sub).

```go
type Observer interface { update(string); getID() string }
type Subject  interface { register(Observer); deregister(Observer); notifyAll() }

type Item struct {
    observerList []Observer
    name         string
    inStock      bool
}

func (i *Item) updateAvailability() {
    i.inStock = true
    i.notifyAll()  // notify all observers
}

type Customer struct{ id string }
func (c *Customer) update(itemName string) {
    fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

// Usage
shirtItem.register(observerFirst)
shirtItem.register(observerSecond)
shirtItem.updateAvailability()  // both get notified
```

---

## State

Changes object behavior when its internal state changes.

```go
type State interface {
    addItem(int) error; requestItem() error; insertMoney(int) error; dispenseItem() error
}

// Concrete states: NoItemState, HasItemState, ItemRequestedState, HasMoneyState

type VendingMachine struct {
    hasItem, itemRequested, hasMoney, noItem State
    currentState State
    itemCount, itemPrice int
}

func (v *VendingMachine) requestItem() error  { return v.currentState.requestItem() }
func (v *VendingMachine) insertMoney(m int) error { return v.currentState.insertMoney(m) }
func (v *VendingMachine) dispenseItem() error { return v.currentState.dispenseItem() }

// State transitions: hasItem → itemRequested → hasMoney → hasItem/noItem
```

---

## Strategy

Defines a family of algorithms and makes them interchangeable.

```go
type EvictionAlgo interface { evict(c *Cache) }

type Lfu struct{}   // least frequently used
type Lru struct{}   // least recently used
type Fifo struct{}  // first in first out

type Cache struct {
    storage      map[string]string
    evictionAlgo EvictionAlgo
    capacity     int
    maxCapacity  int
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) { c.evictionAlgo = e }

// Usage: swap strategy at runtime
cache := initCache(&Lfu{})
cache.add("a", "1"); cache.add("b", "2"); cache.add("c", "3")  // LFU eviction
cache.setEvictionAlgo(&Lru{})
cache.add("d", "4")  // LRU eviction
```

---

## Template Method

Defines a skeleton algorithm in a base struct, letting substructs override specific steps.

```go
type IOtp interface {
    genRandomOTP(int) string; saveOTPCache(string)
    getMessage(string) string; sendNotification(string) error
}

type Otp struct{ iOtp IOtp }

func (o *Otp) genAndSendOTP(otpLength int) error {
    otp := o.iOtp.genRandomOTP(otpLength)
    o.iOtp.saveOTPCache(otp)
    message := o.iOtp.getMessage(otp)
    return o.iOtp.sendNotification(message)
}

type Sms struct{}   // implements IOtp for SMS
type Email struct{} // implements IOtp for Email

// Usage
smsOTP := &Otp{iOtp: &Sms{}}
smsOTP.genAndSendOTP(4)  // SMS flow
emailOTP := &Otp{iOtp: &Email{}}
emailOTP.genAndSendOTP(4) // Email flow
```

---

## Visitor

Adds new operations to objects without modifying their classes.

```go
type Shape interface { getType() string; accept(Visitor) }
type Visitor interface {
    visitForSquare(*Square); visitForCircle(*Circle); visitForrectangle(*Rectangle)
}

type Square    struct{ side int }
type Circle    struct{ radius int }
type Rectangle struct{ l, b int }

func (s *Square) accept(v Visitor)    { v.visitForSquare(s) }
func (c *Circle) accept(v Visitor)    { v.visitForCircle(c) }
func (r *Rectangle) accept(v Visitor) { v.visitForrectangle(r) }

// Two different visitors
type AreaCalculator struct{ area int }
type MiddleCoordinates struct{ x, y int }

// Usage: same shapes, different operations
areaCalculator := &AreaCalculator{}
square.accept(areaCalculator)
middleCoordinates := &MiddleCoordinates{}
square.accept(middleCoordinates)
```
