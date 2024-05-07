## 目次
- [問題文](#問題文)
- [解法](#解法)
  - [Solution1.go（二点間の最大容積を求める双方向アプローチ）](#solution1go二点間の最大容積を求める双方向アプローチ)
  - [Solution2.go（ブルートフォースに近い最適化されたアプローチ）](#solution2goブルートフォースに近い最適化されたアプローチ)
  - [どちらの解法が優れているか](#どちらの解法が優れているか)

### Group Anagrams問題文

##### 日本語

文字列の配列`strs`が与えられたとき、アナグラムをまとめてください。アナグラムとは、別の単語やフレーズの文字を並び替えて形成される単語またはフレーズで、通常は元の文字を全て一度だけ使用します。

### 例

- **例 1:**
  - **入力:** `strs = ["eat","tea","tan","ate","nat","bat"]`
  - **出力:** `[["bat"],["nat","tan"],["ate","eat","tea"]]`

- **例 2:**
  - **入力:** `strs = [""]`
  - **出力:** `[[""]]`

- **例 3:**
  - **入力:** `strs = ["a"]`
  - **出力:** `[["a"]]`

##### 英語
Given an array of strings `strs`, group the anagrams together. An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

### Examples

- **Example 1:**
  - **Input:** `strs = ["eat","tea","tan","ate","nat","bat"]`
  - **Output:** `[["bat"],["nat","tan"],["ate","eat","tea"]]`

- **Example 2:**
  - **Input:** `strs = [""]`
  - **Output:** `[[""]]`

- **Example 3:**
  - **Input:** `strs = ["a"]`
  - **Output:** `[["a"]]`




### Solution 1: 基本的なアプローチの詳細解説

```go
func groupAnagrams(strs []string) [][]string {
    helper := make(map[string][]int)
    for idx, str := range strs {
        runes := []rune(str)
        sort.Slice(runes, func(i, j int) bool {
            return runes[i] < runes[j]
        })
        sortedString := string(runes)
        helper[sortedString] = append(helper[sortedString], idx)
    }
    var res [][]string
    for _, idxs := range helper {
        words := make([]string, len(idxs))
        for i, idx := range idxs {
            words[i] = strs[idx]
        }
        res = append(res, words)
    }
    return res
}
```


このソリューションは、各文字列をソートしてアナグラムを同一視する基本的な手法である。まず、`strs`の各要素をループ処理し、それぞれの文字列についてソートを行う。このソートされた文字列をキーとして使用し、元の文字列のインデックスをマップ`helper`に追加していく。こうすることで、ソートされた文字列が同じである文字列群がキーに対応する値として保存される。最後に、このマップをループして、インデックスに基づき元の文字列を再取得し、結果として返す二次元の文字列配列`res`を構築する。このアプローチは直感的に理解しやすいが、各文字列のソートに伴う時間コストが高く、全体としての実行時間と空間の複雑さが増加する可能性がある。

このソリューションでは、文字列の配列をループして各文字列をソートし、ソートされた文字列をキーとしてインデックスのリストを保持するマップを使用します。これにより、アナグラムがグループ化されます。ただし、文字列の変換とソートにはコストがかかりますので、時間と空間の複雑さが比較的高くなります。

### Solution 2: 改善されたアプローチ

```go
func groupAnagrams(strs []string) [][]string {
    anagramMap := make(map[string][]string)
    var result [][]string
    for _, s := range strs {
        sortedS := sortString(s)
        anagramMap[sortedS] = append(anagramMap[sortedS], s)
    }
    for _, value := range anagramMap {
        result = append(result, value)
    }
    return result
}
```

改善されたこのソリューションでは、インデックスを保持する代わりに、直接元の文字列をマップに追加する。具体的には、各文字列に対してソート処理を行い、その結果得られたソート済み文字列をキーとしてマップ`anagramMap`に文字列を追加する。この手法により、最終的な結果を生成する際のステップが減少し、結果の配列`result`の構築がより直接的かつ効率的になる。また、ソリューション1と比較して、データ構造を利用する過程が簡素化され、コードの可読性が向上している。

### Solution 3: 最先端のアプローチ（最も効率的）

```go
type CharTable [26]byte

func (t *CharTable) Push(value byte) {
    t[value-'a']++
}

func groupAnagrams(strs []string) [][]string {
    group := make(map[CharTable][]string)
    for i := range strs {
        var key CharTable
        for j := range strs[i] {
            key.Push(strs[i][j])
        }
        group[key] = append(group[key], strs[i])
    }
    return getMapValues(group)
}

func getMapValues[K comparable, V any](m map[K]V) []V {
    values := make([]V, 0, len(m))
    for _, value := range m {
        values = append(values, value)
    }
    return values
}
```

最も効率的なこのソリューションでは、文字の出現回数をカウントすることでアナグラムを同一視し、文字列のソートを一切行わずにアナグラムを識別する。具体的には、各文字列に対して、その文字列の各文字に対してカウントを行い、このカウント結果を固定長の配列`CharTable`を用いて記録する。この配列をキーとしてマップ`group`に文字列を追加し、最終的にはこのマップから結果の配列を生成する。この方法では、文字列のソートにかかる時間が削減され、大規模なデータセットに対しても高速に動作するため、計算効率が大幅に向上している。


