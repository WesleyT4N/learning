package main

func Reduce[T, R any](col []T, reducer func(R, T) R, init R) R {
	res := init
	for _, el := range col {
		res = reducer(res, el)
	}
	return res
}

func Sum(numbers []int) int {
	sum := func(total int, number int) int {
		return total + number
	}
	return Reduce(numbers, sum, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	sumAll := func(sums []int, numbers []int) []int {
		return append(sums, Sum(numbers))
	}
	return Reduce(numbersToSum, sumAll, []int{})
}

func SumAllTails(numbersToSum ...[]int) []int {
	sumAllTails := func(sums []int, numbers []int) []int {
		if len(numbers) == 0 {
			return append(sums, 0)
		} else {
			return append(sums, Sum(numbers[1:]))
		}
	}

	return Reduce(numbersToSum, sumAllTails, []int{})
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
    return Transaction{from.Name, to.Name, sum}
}

type Account struct {
    Name string
    Balance float64
}

func NewBalanceFor(a Account, t []Transaction) Account {
    applyTransaction := func(a Account, t Transaction) Account {
		switch a.Name {
		case t.From:
            a.Balance -= t.Sum
		case t.To:
            a.Balance += t.Sum
		}
        return a
    }
    return Reduce(t, applyTransaction, a)
}

func Find[T any](items []T, finder func(T) bool) (val T, found bool) {
    for _, v := range items {
        if finder(v) {
            return v, true
        }
    }
    return
}
