package main

import "fmt"

// 文末のセミコロン
func implicitSemiCollon() {
	fmt.Println("文法的には文末にセミコロンが必要ですが、コンパイラが暗黙的に付与可能なため省略するのが一般的です。")
}

// 配列(スライス)のセミコロン
func semiCollonOfArray(inputArray []string){
	for i, v := range inputArray {
		fmt.Printf("%d:%s\n", i+1, v)
	}
}

// 型を利用したSwitch
func switchingWithTypes(input interface{}){
	switch input.(type) {
		case bool:
			fmt.Printf("input boolean is : %t\n", input)
		case int, uint:
			fmt.Printf("input number is %d\n", input)
		case string:
			fmt.Printf("input string is : %s\n" , input)
		default:
			fmt.Println("I don't know this input type.\n")
	}
}

func main(){
	implicitSemiCollon()

	// 配列やスライスの最後にはセミコロンを書く必要がある。
	// string型の文字列はダブルクォーテーションで囲む必要がある。
	// シングルクォーテーションの場合はrune型と判断される。
	testArray := []string{
		"りんご",
		"ゴリラ",
		"ラッパ",
	}
	semiCollonOfArray(testArray)

	// 型を利用したSwitch
	typeCheckCatalist := 4
	switchingWithTypes(typeCheckCatalist)

}