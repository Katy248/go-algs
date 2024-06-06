#

## Exchange sort - Сортировка перестановками

```go
func (*ExchangeSorter) Sort(source []int32) (result []int32) {
    result = source
    for i := 0; i < len(result); i++ {
        for j := i; j < len(result); j++ {
            if result[i] > result[j] {
                temp := result[j]
                result[j] = result[i]
                result[i] = temp
            }
        }
    }
    return
}
```

Сложность: $O(n^2)$

## Selection sort - Сортировка выбором

```go
func (*SelectionSorter) Sort(source []int32) (result []int32) {
    result = source
    for i := 0; i < len(result); i++ {
        minIndex := i
        for j := i; j < len(result); j++ {
            if result[j] < result[minIndex] {
                minIndex = j
            }
        }
        result[i], result[minIndex] = result[minIndex], result[i]
    }
    return
}
```

Сложность: $O(2n)$

## Bubble sort - Пузырьковая сортировка

```go
func (*BubbleSorter) Sort(source []int32) (result []int32) {
    result = source
    for i := 0; i < len(result); i++ {
        for j := 0; j < len(result)-1; j++ {
            if result[j+1] < result[j] {
                temp := result[j]
                result[j] = result[j+1]
                result[j+1] = temp
            }
        }
    }
    return
}
```

Сложность: $O(n^2)$

## Insertion sort - Сортировка вставками

```go
func (*InsertionSorter) Sort(source []int32) (result []int32) {
    result = copyArray(source)
    for i := 1; i < len(result); i++ {
        if result[i] < result[i-1] {
            for j := i; j > 0 && result[j] < result[j-1]; j-- {
                result[j-1], result[j]= result[j], result[j-1]
            }
        }
    }
    return
}
```

Сложность: $O(n^2)$

## Shaker sort - Сортировка перемешиванием (шейкерная)

AKA: cocktail sort, bidirectional bubble sort

> Пузырьковая была на столько плоха, что пришлось выпустить <q>Пузырьковая сортировка 2</q>

```go
func (*ShakerSorter) Sort(source []int32) (result []int32) {
    result = copyArray(source)

    left := 1
    right := len(result) - 1

    for left <= right {
        for i := right; i >= left; i-- {
            if result[i-1] > result[i] {
                result[i], result[i-1] = result[i-1], result[i]
            }
        }
        left++

        for i := left; i <= right; i++ {
            if result[i-1] > result[i] {
                result[i], result[i-1] = result[i-1], result[i]
            }
        }
        right--
    }
    return
}
```

Сложность: $O(n^2)$ (На самом деле алгоритм быстрее пузырьковой примерно в $\frac{2}{3}$ раза)

## Comb sort - Сортировка расчёской

Ещё одно усовершенствование пузырьковой сортировки
