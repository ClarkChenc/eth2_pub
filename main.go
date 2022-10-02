package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"study/go_study/greetings"
	"study/go_study/htmlparser"
	"syscall"
	"time"

	"github.com/holiman/uint256"
	"github.com/shopspring/decimal"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Unit2String(data uint256.Int) (ret string) {
	zero := uint256.NewInt(0)
	ten := uint256.NewInt(10)

	p_data := &data
	for {
		if data.Cmp(zero) == 0 {
			break
		}

		cur_digit := uint256.NewInt(0)
		cur_digit.Mod(p_data, ten)
		ret += strconv.FormatInt(cur_digit.ToBig().Int64(), 10)
		p_data.Div(&data, ten)
	}
	ret = Reverse(ret)

	if len(ret) == 0 {
		ret = "0"
	}

	return
}

func TestDecimal() {
	data, _ := uint256.FromHex("0xd2f13f7789f0000")
	fmt.Println(Unit2String(*data))

	decimal, err := decimal.NewFromString(Unit2String(*data))
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(decimal)
}

type StringPrinter struct{}

func (p *StringPrinter) String() string {
	return StructToString(p)
}

func StructToString(args interface{}) string {
	b, err := json.Marshal(args)
	if err != nil {
		return fmt.Sprintf("%+v", args)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", args)
	}
	return out.String()
}

type Data struct {
	StringPrinter
	Number int    `json:"number"`
	Name   string `json:"name"`
}

func TestStruct2String() {
	d := Data{}
	fmt.Println(StructToString(d))

	fmt.Print(d.String())
}

func TestTime() {
	query := "{\"num\":" + strconv.FormatUint(1575597387, 10) + "}"
	fmt.Println(query)
	unixTime := time.Unix(1575597387000, 0)

	fmt.Println(unixTime)
}

func (d *Data) PtrAssign() {
	tmp := &Data{}
	tmp.Name = "tmp"
	*d = *tmp
}

func TestPtr() {
	d := &Data{}
	d.PtrAssign()
	fmt.Println(StructToString(d))
}

func TestTimeCompute() {
	const timeInterval time.Duration = 24 * time.Second
	taskDelay := time.Duration(1) * timeInterval

	fmt.Println(taskDelay)
}

func TestSelect() {
	signal := make(chan int, 1)
	select {
	case data := <-signal:
		fmt.Println(data)
	}
}

func TestInt() {
	a := big.NewInt(10)
	fmt.Println(a)
	b := big.NewInt(0).Sub(a, big.NewInt(1))

	fmt.Println(a, b)
}

func TestMap() {
	mp := map[string]string{}
	mp["a"] = "a"
	mp["b"] = "b"
	mp["c"] = "c"

	for i := 0; i < 10; i++ {
		fmt.Println("round:", i)
		for k, v := range mp {
			fmt.Println("key", k, "value", v)
		}
	}
}

func handleTick(signal chan int) {
	fmt.Println("handle tick", "now", time.Now())

	num := rand.Int() % 10
	fmt.Println("random interval", num)
	signal <- num
}

func handleTick2() {
	fmt.Println("---tick start", "time", time.Now())
	time.Sleep(11 * time.Second)
	fmt.Println("\ttick end", "time", time.Now())
}

func TestTicker() {
	interval := 10 * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	signal := make(chan int)
	fmt.Println("start", "now", time.Now())
	for {
		select {
		case <-ticker.C:
			ticker.Reset(interval)
			handleTick2()

		case num := <-signal:
			ticker.Reset(time.Duration(num * int(time.Second)))
		}
	}
}

func TestGreeting() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{
		"Gladys",
		"Samantha",
		"Darrin",
	}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}

func TestArrary() {
	datas := []Data{
		Data{Number: 1},
		Data{Number: 2},
		Data{Number: 3},
	}

	for _, data := range datas {
		fmt.Printf("data %v, address %x\n", data, &data)
		data.Number = 10
	}
	for _, data := range datas {
		fmt.Printf("data %v\n", data)
	}
}

func TestSkip() {
	arr := []int{}

	start := 0
	size := 1000
	for i := 0; i < size; i++ {
		start += 16*60 + 50
		start = start % (30 * 60)

		if start <= 48 {
			arr = append(arr, start)
		}
	}

	fmt.Println(arr)

	ratio := float64(len(arr)) / float64(size)
	fmt.Println(ratio)
}

func TestReg() {
	pattern := regexp.MustCompile(`https://(.*?)/`)
	match := pattern.FindStringSubmatch("https://filfox.info/en/address/f1xsq7i5dm53l7xq5jqrw7exwciz6vdqro2w5kaey")
	fmt.Println(match)
}

func TestBlockChair() {
	start := time.Now().UTC().Unix()
	resp, err := http.Get("https://blockchair.com/bitcoin-sv/address/17ve2EPbtvUaQykXvTBhvKHX9e9uS2kFi5")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(body)

	pattern, _ := regexp.Compile(`field="balance"[\s\S]+?>([\d|\.|\,]+?)<`)

	match := pattern.FindStringSubmatch(content)

	fmt.Println("used time: ", time.Now().UTC().Unix()-start)

	fmt.Println(len(match))
	fmt.Println(match[1])
}

func TestFilScan() {
	start := time.Now().UTC().Unix()
	resp, err := http.Get("https://filfox.info/en/address/f1rmrlyjzym3jgv6tt77kbbbwbz6huxbimlt5va7y")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(body)

	pattern, _ := regexp.Compile(`Balance[\s\S]+?> ([\d|\.|\,]+?) FIL`)

	match := pattern.FindStringSubmatch(content)

	fmt.Println("used time: ", time.Now().UTC().Unix()-start)

	fmt.Println(len(match))
	fmt.Println(match[1])
}

func TestXTZ() {
	// start := time.Now().UTC().Unix()
	// resp, err := http.Get("https://api.tzkt.io/v1/accounts/tz1gMsMPNBnq4AqnjrAgG536DP12uVLFgRbt/balance")
	resp, err := http.Get("https://blockchain.elastos.org/address/EXoZzV2qMvCG2je8Yt7Rs3ruDMF9FAMUsp")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(body)
	fmt.Println(content)

	// pattern, _ := regexp.Compile(`(.*)`)

	// match := pattern.FindStringSubmatch(content)

	// fmt.Println("used time: ", time.Now().UTC().Unix()-start)

	// fmt.Println(len(match))
	// fmt.Println(match[1])
}

func TestCatchKillSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		<-c
		fmt.Println("elgant exit")
		os.Exit(1)
	}()

	for {
		fmt.Println("program is running...")
		time.Sleep(1 * time.Second)
	}
}

func TestCloseChan() {
	c := make(chan struct{})

	go func() {
		time.Sleep(10 * time.Second)
		close(c)
	}()
	<-c
	fmt.Println("get")
}

func main() {
	// TestDecimal()
	// TestStruct2String()
	// TestTime()
	// TestPtr()

	// TestTimeCompute()
	// TestSelect()

	// fmt.Println(decimal.NewFromInt(int64(math.Pow10(10))))
	// TestInt()

	// eth.EthGasFee()
	// copy.CopyTest()

	// parser.TestParser()
	// ticker.TestAdjustTicker()
	// TestTicker()

	// TestGreeting()

	// TestMap()
	// TestArrary()
	// TestSkip()

	// algorithm.TestRandom()
	// TestBlockChair()
	// TestXTZ()

	// TestReg()
	// operator.TestSigOperator()

	// TestCatchKillSignal()
	// TestCloseChan()

	htmlparser.GetETH2AccountBalance()
}
