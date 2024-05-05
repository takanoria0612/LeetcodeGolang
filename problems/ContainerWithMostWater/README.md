# Container With Most Water

## 問題文

n個の非負整数 a1, a2, ..., an が与えられます。各整数は座標 (i, ai) における点を表し、n本の垂直線が引かれます。各線の一端は (i, ai)、もう一端は (i, 0) にあります。x軸とこれらの線のうち2本で形成されるコンテナを考え、最も多くの水を含むことができるコンテナを見つけてください。


Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

## 解法

ここでは、`Container With Most Water` 問題に対する二つの解法 `solution1.go` と `solution2.go` を解説し、それぞれの特徴と効率性を評価します。この問題は、一次元の整数配列 `height` が与えられ、`height[i]` が縦の線を表すとき、最も水を多く保持できる二つの線を見つけることです。

### Solution1.go（二点間の最大容積を求める双方向アプローチ）

```go
package ContainerWithMostWater

import "math"

func MaxArea(height []int) int {
	max_area := 0
	l := 0
	r := len(height) - 1
	for l < r {
		h := float64(math.Min(float64(height[l]), float64(height[r])))
		max_area = int(math.Max(float64(max_area), h*float64(r-l)))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max_area
}
```

#### 解法の解説
このアルゴリズムは、配列の最初と最後から始めて双方向に探索を進める方法です。左端 (`l`) と右端 (`r`) のポインタを使用し、それぞれの間の距離と高さの小さい方を使用して水の量を計算します。次に、より低い側のポインタを中央に向かって移動させ、これを繰り返しながら最大の水の量を探索します。この方法は効率的で、計算量は \(O(n)\) です。

### Solution2.go（ブルートフォースに近い最適化されたアプローチ）

```go
package ContainerWithMostWater

func MaxArea2(height []int) int {
	retVal := 0
	for i := 0; i < len(height); i++ {
		localMax := 0
		for j := len(height) - 1; j >= i; j-- {
			maxContainerPossible := (j - i) * height[i]
			if maxContainerPossible < retVal || maxContainerPossible < localMax {
				break
			}
			containerSize := (j - i) * min(height[i], height[j])
			localMax = max(localMax, containerSize)
		}
		retVal = max(retVal, localMax)
	}
	return retVal
}
```

#### 解法の解説
このアルゴリズムは、各位置 `i` で可能な最大の水の量を評価し、それよりも大きな可能性がなくなれば内側のループを抜けるという最適化を施しています。内部ループでは `j` を `i` から配列の末尾まで減少させながら、可能な最大容積を計算し、その容積が現在の最大値 `retVal` より小さくなればループを抜けます。これにより不要な計算を省略していますが、最悪のケースでは \(O(n^2)\) の計算量がかかります。

### どちらの解法が優れているか

## 実行速度
- **型変換の削減：** 元のコードでは `int` から `float64` への型変換が必要でしたが、代替案ではその必要がありません。型変換は実行時間に微妙ながらも追加のコストをかける可能性があります。特にループの中で繰り返し行われる場合、この時間は無視できない場合があります。

- **関数呼び出し：** `math.Max` 関数を使用することも、関数呼び出しのオーバーヘッドを引き起こします。代替案では単純な条件式を使って最大値を求めており、これにより追加の関数コールが削減され、パフォーマンスが向上します。

## メモリ消費
- **メモリアロケーション：** `float64` 型は `int` 型よりも多くのメモリ（通常 `int` は4または8バイト、`float64` は常に8バイト）を消費するため、型変換によって一時的なメモリ使用量が増加する可能性があります。ただし、これは非常に小さな違いであり、実際のアプリケーションのパフォーマンスに顕著な影響を与えることは少ないです。

- **スタック使用量：** 整数の計算は、通常、浮動小数点数の計算よりもCPUが扱いやすいため、スタック上の処理が少し効率的になる可能性があります。

### まとめ

`solution1.go` のアプローチは、全体的に高効率で \(O(n)\) の計算量で解を求めることができます。一方で `solution2.go` は、最悪のケースでは非効率的ですが、早期終了の条件により多くのケースで計算量を削減できるため、一部のシナリオでは高速に動作する可能性があります。

