package main

import (
	"fmt"

	"github.com/tarantool/go-tarantool"
	"gopkg.in/vmihailenco/msgpack.v2"
)

type User struct {
	ID uint64
	Name string
}

func (u *User) EncodeMsgpack(e *msgpack.Encoder) error {
	e.EncodeArrayLen(2)
	if err := e.EncodeUint64(u.ID); err != nil {
		return err
	}

	if err := e.EncodeString(u.Name); err != nil {
		return err
	}

	return nil
}

func (u *User) DecodeMsgpack(d *msgpack.Decoder) error {
	l, err := d.DecodeArrayLen()
	if err != nil {
		return err
	}

	ll := 1
	if u.ID, err = d.DecodeUint64(); err != nil || l == ll {
		return err
	}

	ll++
	if u.Name, err = d.DecodeString(); err != nil || l == ll {
		return err
	}

	return nil
}

func main() {
	conn, err := tarantool.Connect("127.0.0.1:3301", tarantool.Opts{})
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	users := []User{
		{
			ID:   1,
			Name: "Ivan",
		},
		{
			ID:   2,
			Name: "Bob",
		},
	}

	for i := range users {
		_, err = conn.Replace("user", &users[i])
		if err != nil {
			panic(err)
		}
	}

	resp := make([]User, 0)
	err = conn.Call17Typed("get_user", []interface{}{2}, &resp)
	if err != nil {
		panic(err)
	}

	if len(resp) == 0 {
		panic("")
	}

	fmt.Printf("%v\n", resp[0])

	topResp := make([][]User, 0)
	err = conn.Call17Typed("get_top10_users", []interface{}{}, &topResp)
	if err != nil {
		panic(err)
	}

	if len(resp) == 0 {
		panic("")
	}

	fmt.Printf("%v\n", topResp[0])

}
