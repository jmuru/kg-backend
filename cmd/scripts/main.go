package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kat-generator/KGB/client"
	"net/http"
)

func main() {
	reqUrl := "http://localhost:8080/palette/create"
	accessories := [][]string{
		{"1f2041", "35305a", "4b3f72", "15c2cb"},
		{"fab3a9", "c6ad94", "7fb285", "463239"},
		{"db9d47", "8a1c7c", "da4167", "f0bcd4"},
		{"dbd3c9", "d81159", "8f2d56", "218380"},
		{"f8fcda", "e3e9c2", "f9fbb2", "de3f4c"},
		{"0a2463", "fb3640", "605f5e", "247ba0"},
		{"da2c38", "226f54", "87c38f", "f4f0bb"},
		{"4c191b", "963d5a", "577561", "e3daff"},
		{"4b7f52", "eee0cb", "baa898", "c2847a"},
		{"03b5aa", "037971", "023436", "593959"},
		{"c1dbe3", "c7dfc5", "f6feaa", "fce694"},
		{"5a7faf", "bbc8ca", "b592a0", "9a5b89"},
		{"26532b", "63d471", "63a46c", "775144"},
		{"dda448", "bb342f", "f7a1c4", "6c91c2"},
		{"eee0cb", "baa898", "848586", "c2847a"},
		{"88958d", "606d5d", "bc9cb0", "d3cdd7"},
		{"a4243b", "d8973c", "d8c99b", "efb9cb"},
		{"0fa3b1", "b5e2fa", "f9f7f3", "ff595e"},
		{"3626a7", "02c39a", "00635d", "657ed4"},
		{"f2ccc3", "c94277", "cdeac0", "6d98ba"},
		{"0b132b", "1c2541", "3a506b", "5bc0be"},
		{"157f1f", "4cb963", "f25f5c", "6b2737"},
		{"ba2d0b", "d5f2e3", "73ba9b", "003e1f"},
		{"a76571", "588b8b", "e4dfda", "d4b483"},
		{"d782ba", "e18ad4", "eeb1d5", "efc7e5"},
		{"5597dd", "f8f012", "dd882c", "a40606"},
		{"ede3e4", "bb4430", "214e34", "6f8695"},
		{"ddd1c7", "c2cfb2", "8db580", "7d6d61"},
		{"2b59c3", "253c78", "d36582", "ffeecf"},
		{"1446a0", "67e0cc", "66a5e1", "f2bac9"},
		{"f92a82", "ed7b84", "f5dbcb", "d6d5b3"},
		{"b80c09", "ffca3a", "fbfbff", "1982c4"},
		{"db5375", "4f86c6", "a4d4b4", "ffcf9c"},
		{"880d1e", "dd2d4a", "f26a8d", "f49cbb"},
		{"dd614a", "f48668", "f4a698", "c5c392"},
		{"395b50", "3e2a35", "f79f79", "87b6a7"},
		{"031926", "468189", "77aca2", "9dbebb"},
		{"0e3b43", "357266", "a3bbad", "aa1155"},
		{"069e2d", "058e3f", "04773b", "e53d00"},
		{"916953", "cf8e80", "fcb5b5", "5299d3"},
		{"d4c5e2", "c9d7f8", "83d6d8", "80cfa9"},
		{"699b59", "d16666", "ffa630", "2e5077"},
		{"e95d9c", "090c9b", "3d52d5", "edd9a3"},
		{"590004", "fbd87f", "04a777", "d65780"},
		{"241e4e", "faeedc", "70d6ff", "ff7b9c"},
		{"87b38d", "cc76a1", "4a77bf", "dd9296"},
		{"c41e3d", "ff2c55", "438bc7", "44af69"},
		{"05668d", "db7f8e", "eeeeee", "6cc68c"},
		{"135490", "faf0ca", "f4d35e", "ee964b"},
		{"c14953", "e5dcc5", "1c325f", "3070c5"},
		{"e4fde1", "8acb88", "648381", "575761"},
		{"f4dbd8", "bea8a7", "c09891", "775144"},
		{"d9f0ff", "a3d5ff", "83c9f4", "f9c80e"},
		{"c2f9bb", "e8b9ab", "62c370", "cc3363"},
		{"4281a4", "102542", "e4dfda", "ff674d"},
		{"545f66", "829399", "d0f4ea", "e8fcc2"},
		{"ef767a", "456990", "49dcb1", "ffe66d"},
		{"1b264f", "274690", "576ca8", "d5c9df"},
		{"1c5d99", "8fbfe0", "d65780", "fcca46"},
		{"9b1d20", "6b8f71", "aad2ba", "d9fff5"},
		{"fae8eb", "f6caca", "e4c2c6", "cd9fcc"},
		{"28536b", "c2948a", "7ea8be", "f6f0ed"},
		{"1c7c54", "0b3954", "bfd7ea", "dd7596"},
		{"a2a79e", "a27e8e", "a77464", "88292f"},
		{"749c75", "f9f9f9", "083d77", "e072a4"},
		{"f08700", "f49f0a", "efca08", "00a6a6"},
		{"faf0ca", "f4d35e", "ee964b", "92140c"},
		{"ba1200", "507dbc", "a1c6ea", "bbd1ea"},
		{"e5e6e4", "cfd2cd", "a6a2a2", "d97f6d"},
		{"a49e8d", "504136", "689689", "b2e6d4"},
		{"fc7a57", "fc904f", "fca547", "fbd037"},
		{"acadbc", "9b9ece", "6665dd", "473bf0"},
		{"f4afb4", "ffdc5e", "0d63bf", "94a89a"},
		{"cc3363", "6f9283", "8d9f87", "cdc6a5"},
		{"f46036", "2e294e", "1b998b", "e71d36"},
		{"f0f7ee", "c4d7f2", "afdedc", "91a8a4"},
		{"05668d", "028090", "00a896", "02c39a"},
		{"ffcab1", "69a2b0", "659157", "8b5d33"},
		{"ff2e00", "fffb0a", "06aed5", "ffb5c2"},
		{"ebd4cb", "da9f93", "b6465f", "890620"},
		{"4f6d7a", "c0d6df", "dbe9ee", "4a6fa5"},
		{"e2d4b7", "9c9583", "a1a499", "6d709c"},
		{"bbe2b6", "5da9e9", "003f91", "ffffff"},
		{"6aa08d", "ffd6e0", "2f3061", "ffe66d"},
		{"03256c", "2541b2", "1768ac", "06bee1"},
		{"9aadbf", "6d98ba", "d3b99f", "c17767"},
		{"966b9d", "c98686", "f2b880", "fff4ec"},
		{"c5d1eb", "92afd7", "5a7684", "395b50"},
		{"119da4", "0c7489", "13505b", "040404"},
		{"805d93", "f49fbc", "ffd3ba", "9ebd6e"},
		{"21897e", "3ba99c", "69d1c5", "7ebce6"},
		{"f8ff33", "fe4a49", "3685b5", "0471a6"},
		{"cebebe", "ece2d0", "d5b9b2", "a26769"},
		{"f1e0c5", "c9b79c", "71816d", "342a21"},
		{"f1f7ee", "b0bea9", "92aa83", "e0edc5"},
		{"bbe5ed", "ffde5c", "d11149", "6b16f3"},
		{"f1dac4", "a69cac", "474973", "3bb273"},
		{"820263", "d90368", "eadeda", "2e294e"},
		{"0d1b1e", "7798ab", "c3dbc5", "e8dcb9"},
		{"ff7f7f", "ea4848", "dc1919", "ff206e"},
	}
	for _, accessory := range accessories {
		p := &client.PaletteData{
			Palette: accessory,
			Type:    "accessory",
		}
		fmt.Println("call marshal")
		json, err := json.Marshal(p)
		if err != nil {
			fmt.Println("error in marshal")
			panic(err)
		}
		fmt.Printf("marshalled json %v\n", string(json))

		req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(json))
		if err != nil {
			fmt.Printf("bad request response: %d, err: %v\n", req.Response.StatusCode, err)
			panic(err)
		}
		client := &http.Client{}
		fmt.Println("Do req")
		response, error := client.Do(req)
		if error != nil {
			panic(error)
		}
		defer response.Body.Close()
		fmt.Printf("successfully saved to the db response: %d\n", response.StatusCode)
	}
}
