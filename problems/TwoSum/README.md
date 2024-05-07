## Two-Sum 問題解決法の説明

この文書では、プログラミングで一般的な「Two-Sum問題」を解く二つのアプローチについて説明します。また、後者のハッシュマップを使用した方法がなぜ初めの方法よりも優れているかを解説します。

### 方法1: ネストされたループを使用するアプローチ

```go
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}
```

#### 説明
この方法では、すべての要素の組み合わせを総当りで調べています。配列 `nums` の各要素 `i` に対して、`i` より後の要素 `j` をループさせ、それぞれの和が `target` と等しいかどうかを確認します。このアプローチは直感的で理解しやすいですが、時間計算量が \(O(n^2)\) となり、大きなデータセットでは非効率的です。

### 方法2: ハッシュマップを使用するアプローチ

```go
package main

import "fmt"

func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i, num := range nums {
        result := target - num
        if _, ok := numMap[result]; ok {
            return []int{numMap[result], i}
        }
        numMap[num] = i
    }
    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    fmt.Println(twoSum(nums, target)) // 出力: [0, 1]
}
```

#### 説明
この方法では、`map` データ構造を利用して各数字とそのインデックスを保存し、計算を進めることで効率よく解を見つけます。各ステップで `target` から現在の数値を引いた結果が `map` に既に存在するかどうかを確認します。この手法では時間計算量が \(O(n)\) と大幅に改善され、大規模なデータに対しても高速に動作する。




